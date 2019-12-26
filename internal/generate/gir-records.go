package generate

import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
)

type Records []*Record

func (rr Records) init(ns *Namespace) {
	for _, record := range rr {
		record.init(ns)
	}
}

func (rr Records) byName(name string) (*Record, bool) {
	for _, r := range rr {
		if r.Name == name {
			return r, true
		}
	}

	return nil, false
}

func (rr Records) generateSys(f *jen.File, version semver.Version) {
	for _, r := range rr {
		r.generateSysType(f, version)
	}

	f.Line()
}