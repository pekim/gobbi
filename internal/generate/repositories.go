package generate

import "fmt"

type RepositorySpec struct {
	Version string
	Name    string
}

func Generate(specs []RepositorySpec) {
	for _, spec := range specs {
		r := repositoryFromFile(spec)
		fmt.Println(r.Namespace.Version, r.Namespace.Name)
	}
}
