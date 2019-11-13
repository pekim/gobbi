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

	// Deep initialise all namespaces.
	// In particular provide descendants with their namespace.
	for _, r := range rr {
		r.Namespace.init(namespaces)
	}

	// Generate files for all namespaces
	for _, r := range rr {
		r.Namespace.generate()
	}
}
