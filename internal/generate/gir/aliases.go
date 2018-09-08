package gir

type Aliases []*Alias

func (aa Aliases) fixup(ns *Namespace) {
	for _, alias := range aa {
		alias.fixup(ns)
	}
}

func (aa Aliases) versionList() Versions {
	return Versions{}
}
