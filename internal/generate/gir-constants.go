package generate

import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
)

type Constants []*Constant

func (cc Constants) init(ns *Namespace) {
	for _, constant := range cc {
		constant.init(ns)
	}
}

func (cc Constants) generateLib(f *jen.File, version semver.Version) {
	for _, c := range cc {
		c.generateLib(f, version)
	}

	f.Line()
}
