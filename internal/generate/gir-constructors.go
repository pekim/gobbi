package generate

import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
)

type Constructors []*Constructor

func (cc Constructors) init(ns *Namespace, record *Record) {
	for _, ctor := range cc {
		ctor.init(ns, record)
	}
}

func (cc Constructors) generateSys(f *jen.File, version semver.Version) {
	for _, ctor := range cc {
		ctor.generateSys(f, version)
	}
}
