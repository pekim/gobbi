package generate

import (
	"github.com/dave/jennifer/jen"
)

type Constructors []*Constructor

func (cc Constructors) init(ns *Namespace, record *Record) {
	for _, ctor := range cc {
		ctor.init(ns, record)
	}
}

func (cc Constructors) forCIdentifier(cidentifier string) *Constructor {
	for _, ctor := range cc {
		if ctor.CIdentifier == cidentifier {
			return ctor
		}
	}

	return nil
}

func (cc Constructors) mergeAddenda(addenda Constructors) {
	for _, addendaConstructor := range addenda {
		if ctor := cc.forCIdentifier(addendaConstructor.CIdentifier); ctor != nil {
			ctor.mergeAddenda(addendaConstructor.Function)
		}
	}
}

func (cc Constructors) generate(g *jen.Group, version *Version) {
	for _, ctor := range cc {
		ctor.generate(g, version)
	}
}
