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
