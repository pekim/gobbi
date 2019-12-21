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

func (ee Enumerations) byName(name string) (*Enumeration, bool) {
	for _, e := range ee {
		if e.Name == name {
			return e, true
		}
	}

	return nil, false
}

func (ee Enumerations) generateSys(f *jen.File, version semver.Version, typ string) {
	f.Comment(typ)

	for _, e := range ee {
		e.generateSys(f, version)
	}

	f.Line()
}
