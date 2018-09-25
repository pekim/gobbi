package generate

import (
	"github.com/dave/jennifer/jen"
)

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

func (mm Methods) forCIdentifier(cidentifier string) *Method {
	for _, method := range mm {
		if method.CIdentifier == cidentifier {
			return method
		}
	}

	return nil
}

func (mm Methods) mergeAddenda(addenda Methods) {
	for _, addendaMethod := range addenda {
		if method := mm.forCIdentifier(addendaMethod.CIdentifier); method != nil {
			method.mergeAddenda(addendaMethod.Function)
		}
	}
}
