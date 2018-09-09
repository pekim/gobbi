package gir

import "github.com/dave/jennifer/jen"

type Aliases []*Alias

func (aa Aliases) fixup(ns *Namespace) {
	for _, alias := range aa {
		alias.fixup(ns)
	}
}

func (aa Aliases) versionList() Versions {
	return Versions{}
}

func (aa Aliases) generate(file *jen.File) {
	// for _, alias := range aa {
	// 	alias.generate(file.Group)
	// }
}
