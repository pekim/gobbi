package generate

import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
)

type Aliases []*Alias

func (aa Aliases) init(ns *Namespace) {
	for _, alias := range aa {
		alias.init(ns)
	}
}

func (aa Aliases) byName(name string) (*Alias, bool) {
	for _, alias := range aa {
		if alias.Name == name {
			return alias, true
		}
	}

	return nil, false
}

func (aa Aliases) generateLib(f *jen.File, version semver.Version) {
	for _, a := range aa {
		a.generateLib(f, version)
	}
}
