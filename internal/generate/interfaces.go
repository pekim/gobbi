package generate

import "github.com/dave/jennifer/jen"

type Interfaces []*Interface

func (ii Interfaces) init(ns *Namespace) {
	for _, i := range ii {
		i.init(ns)
	}
}

func (ii Interfaces) generate(g *jen.Group, version *Version) {
	for _, i := range ii {
		i.generate(g, version)
	}
}

func (ii Interfaces) forName(name string) *Interface {
	for _, iface := range ii {
		if iface.Name == name {
			return iface
		}
	}

	return nil
}
