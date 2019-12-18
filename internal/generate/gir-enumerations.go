package generate

import "github.com/dave/jennifer/jen"

type Enumerations []*Enumeration

func (ee Enumerations) init(ns *Namespace) {
	for _, enum := range ee {
		enum.init(ns)
	}
}

func (ee Enumerations) generateSys(f *jen.File, typ string) {
	f.Comment(typ)

	for _, enum := range ee {
		enum.generateSys(f)
	}

	f.Line()
}
