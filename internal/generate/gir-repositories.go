package generate

import (
	"sync"
)

type RepositorySpec struct {
	Version string
	Name    string
}

func Generate(specs []RepositorySpec) {
	var wg sync.WaitGroup

	namespaces := make(namespaces)
	rr := []*repository{}

	var mu sync.Mutex
	for _, spec := range specs {
		wg.Add(1)

		go func(spec RepositorySpec) {
			defer wg.Done()

			r := repositoryFromFile(spec)

			mu.Lock()
			defer mu.Unlock()
			namespaces[r.Namespace.Name] = r.Namespace
			rr = append(rr, r)
		}(spec)
	}

	wg.Wait()

	// Deep initialise all namespaces.
	// In particular provide descendants with their namespace.
	for _, r := range rr {
		r.Namespace.init(namespaces)
	}

	// Generate files for all namespaces
	for _, r := range rr {
		wg.Add(1)

		go func() {
			r.Namespace.generate()
			wg.Done()
		}()
	}
	wg.Wait()
}
