package generate

import (
	"fmt"
	"sync"
)

type RepositorySpec struct {
	Version string
	Name    string
}

func Generate(specs []RepositorySpec) {
	fmt.Print("\u001B[?25l")       // hide cursor
	defer fmt.Print("\u001B[?25h") // show cursor

	var wg sync.WaitGroup

	namespaces := make(namespaces)
	rr := []*repository{}

	progressLineLen := 2 * len(specs)
	prepareProgressLine(progressLineLen)

	var mu sync.Mutex
	for _, spec := range specs {
		wg.Add(1)

		go func(spec RepositorySpec) {
			defer wg.Done()

			r := repositoryFromFile(spec)

			mu.Lock()
			defer mu.Unlock()

			incrementProgress()
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

			mu.Lock()
			defer mu.Unlock()
			incrementProgress()

			wg.Done()
		}()
	}
	wg.Wait()

	clearProgressLine(progressLineLen)
}

func lineOfChars(char string, len int) {
	fmt.Print("\u001B[1G") // start of line

	for i := 0; i < len; i++ {
		fmt.Print(char)
	}

	fmt.Print("\u001B[1G") // start of line
}

func prepareProgressLine(len int) {
	lineOfChars("\u2591", len) // light shade block
}

func clearProgressLine(len int) {
	lineOfChars(" ", len)
}

func incrementProgress() {
	fmt.Print("\u2592") // medium shade block
}
