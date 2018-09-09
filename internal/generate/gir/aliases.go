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

func (aa Aliases) entities() []Generatable {
	var generatables []Generatable

	// DO NOT GENERATE aliases until they and their use is better understood.

	// for _, function := range aa {
	// 	generatables = append(generatables, function)
	// }

	return generatables
}
