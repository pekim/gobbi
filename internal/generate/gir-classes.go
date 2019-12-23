package generate

import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
)

type Classes []*Class

func (cc Classes) init(ns *Namespace) {
	for _, class := range cc {
		class.init(ns)
	}
}

func (cc Classes) byName(name string) (*Class, bool) {
	for _, c := range cc {
		if c.Name == name {
			return c, true
		}
	}

	return nil, false
}

func (cc Classes) generateSys(f *jen.File, version semver.Version) {
	for _, c := range cc {
		c.generateSys(f, version)
	}

	f.Line()
}
