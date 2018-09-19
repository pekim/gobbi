package generate

import "github.com/dave/jennifer/jen"

type Constructors []*Constructor

func (cc Constructors) init(ns *Namespace, record *Record) {
	for _, ctor := range cc {
		ctor.init(ns, record)
	}
}

func (cc Constructors) generate(g *jen.Group, version *Version) {
	for _, ctor := range cc {
		ctor.generate(g, version)
	}
}
