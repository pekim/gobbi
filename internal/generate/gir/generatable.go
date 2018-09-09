package gir

import (
	"github.com/dave/jennifer/jen"
)

type Generatable interface {
	Versioned
	generate(g *jen.Group, version *Version)
}

type Generatables interface {
	VersionLister
	entities() []Generatable
}
