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

	r.generateLibType(f)
	r.generateToC(f)
}

func (r *Record) generateLibType(f *jen.File) {
	f.Commentf("%s is a representation of the C record %s.", r.goName, r.CType)
	docVersion(f, r.Version)

	f.Type().Id(r.goName).Struct(
		jen.Id("native").Add(jenUnsafePointer()),
	)

	f.Line()
}

func (r *Record) generateToC(f *jen.File) {
	f.
		Func().
		Parens(jen.Id("recv").Op("*").Id(r.goName)).
		Id("ToC").
		Params().
		Add(jenUnsafePointer()).
		Block(jen.Return().Id("recv").Dot("native"))
}
