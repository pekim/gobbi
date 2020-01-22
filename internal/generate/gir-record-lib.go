package generate

import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
)

func (r *Record) generateLib(f *jen.File, version semver.Version) {
	if r.blacklist {
		f.Commentf("UNSUPPORTED : %s : blacklisted", r.Name)
		f.Line()
		return
	}

	if !generateEntityForVersion(version, r.version) {
		return
	}

	generateLibStructType(f, "record", r.Name, r.CType, r.Version)
	generateLibToC(f, r.Name, r.CType)
	generateLibNewFromC(f, r.Name, r.CType)
}
