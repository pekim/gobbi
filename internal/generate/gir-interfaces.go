package generate

import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
)

type Interfaces []*Interface

func (ii Interfaces) init(ns *Namespace) {
	for _, i := range ii {
		i.init(ns)
	}
}

func (ii Interfaces) byName(name string) (*Interface, bool) {
	for _, iface := range ii {
		if iface.Name == name {
			return iface, true
		}
	}

	return nil, false
}

func (ii Interfaces) generateSys(f *jen.File, version semver.Version) {
	f.Comment("interfaces")

	for _, i := range ii {
		i.generateSys(f, version)
	}

	f.Line()
}
