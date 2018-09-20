package generate

import "github.com/dave/jennifer/jen"

type Fields []*Field

func (cc Fields) init(ns *Namespace) {
	for _, f := range cc {
		f.init(ns)
	}
}

func (ff Fields) generate(g *jen.Group) {
	for _, f := range ff {
		f.generate(g)
	}
}
