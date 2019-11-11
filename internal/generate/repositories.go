package generate

type RepositorySpec struct {
	Version string
	Name    string
}

func Generate(specs []RepositorySpec) {
	namespaces := make(namespaces)
	rr := []*repository{}

	for _, spec := range specs {
		r := repositoryFromFile(spec)
		namespaces[r.Namespace.Name] = r.Namespace
		rr = append(rr, r)
	}

	for _, r := range rr {
		r.Namespace.namespaces = namespaces
	}

	for _, r := range rr {
		r.Namespace.generate()
	}
}
