package generate

import "github.com/dave/jennifer/jen"

type Signals []*Signal

func (ss Signals) init(ns *Namespace, record *Record) {
	for _, signal := range ss {
		signal.init(ns, record)
	}
}

func (ss Signals) generate(g *jen.Group, version *Version) {
	for _, signal := range ss {
		signal.generate(g, version)
	}
}
