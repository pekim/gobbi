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

	if r.version.GT(version) {
		return
	}

	f.Commentf("%s is a representation of the C record %s.", r.goName, r.CType)

	f.Type().Id(r.goName).Struct(
		jen.Id("native").Add(jenUnsafePointer()),
	)

	f.Line()
}
