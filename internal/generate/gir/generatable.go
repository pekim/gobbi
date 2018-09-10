package gir

import (
	"github.com/dave/jennifer/jen"
)

type Generatable interface {
	Versioned
	blacklisted() (bool, string)
	generate(g *jen.Group, version *Version)
	supported() (supported bool, reason string)
}

type Generatables interface {
	VersionLister
	entities() []Generatable
}
