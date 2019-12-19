package generate

import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
)

type Enumerations []*Enumeration

func (ee Enumerations) init(ns *Namespace) {
	for _, enum := range ee {
		enum.init(ns)
	}
}

func (ee Enumerations) generateSys(f *jen.File, version semver.Version, typ string) {
	f.Comment(typ)

	for _, enum := range ee {
		enum.generateSys(f, version)
	}

	f.Line()
}
