package generate

import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
)

type Functions []*Function

func (ff Functions) init(ns *Namespace /*, namePrefix string*/) {
	for _, function := range ff {
		function.init(ns, nil, false)
	}
}

func (ff Functions) generateSys(f *jen.File, version semver.Version) {
	for _, fn := range ff {
		fn.generateSys(f, version)
	}
}

func (ff Functions) generateLib(f *jen.File, version semver.Version) {
	for _, fn := range ff {
		fn.generateLib(f, version)
	}
}
