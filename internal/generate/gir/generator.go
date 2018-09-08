package gir

import (
	"github.com/dave/jennifer/jen"
)

type Generator interface {
	// Versioned
	generate(file *jen.File, version *Version)
}

type Generators interface {
	VersionLister
	entities() []Generator
}
