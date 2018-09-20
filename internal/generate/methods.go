package generate

import "github.com/dave/jennifer/jen"

type Methods []*Method

func (mm Methods) init(ns *Namespace, record *Record) {
	for _, m := range mm {
		m.init(ns, record)
	}
}

func (mm Methods) generate(g *jen.Group, version *Version) {
	for _, method := range mm {
		method.generate(g, version)
	}
}
