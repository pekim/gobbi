package generate

type Aliases []*Alias

func (aa Aliases) init(ns *Namespace) {
	for _, alias := range aa {
		alias.init(ns)
	}
}

func (aa Aliases) generate(f *file) {
	for _, alias := range aa {
		alias.generate(f)
	}
}

func (aa Aliases) findByName(name string) (*Alias, bool) {
	for _, alias := range aa {
		if alias.Name == name {
			return alias, true
		}
	}

	return nil, false
}
