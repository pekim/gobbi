package generate

import "github.com/dave/jennifer/jen"

type Enumerations []*Enumeration

func (ee Enumerations) init(ns *Namespace) {
	for _, enum := range ee {
		enum.init(ns)
	}
}

func (ee Enumerations) generateSys(f *jen.File) {
	for _, enum := range ee {
		enum.generateSys(f)
	}
}
