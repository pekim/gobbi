package generate

import "github.com/pekim/jennifer/jen"

type Signals []*Signal

func (ss Signals) init(ns *Namespace, record *Record) {
	for _, signal := range ss {
		signal.init(ns, record)
	}
}

func (ss Signals) generate(g *jen.Group, version *Version, parentVersion string) {
	for _, signal := range ss {
		signal.generate(g, version, parentVersion)
	}
}
