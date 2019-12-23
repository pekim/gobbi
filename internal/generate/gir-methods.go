package generate

import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
)

type Methods []*Method

func (mm Methods) init(ns *Namespace, record *Record) {
	for _, m := range mm {
		m.init(ns, record)
	}
}

func (mm Methods) generateSys(f *jen.File, version semver.Version) {
	for _, m := range mm {
		m.generateSys(f, version)
	}
}
