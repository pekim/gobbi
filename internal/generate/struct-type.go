package generate

import (
	"github.com/dave/jennifer/jen"
)

func generateLibStructType(f *jen.File, typ string, name string, cType string, version string) {
	f.Commentf("%s is a representation of the C %s %s.", name, typ, cType)
	docVersion(f, version)

	f.Type().Id(name).Struct(
		jen.Id("native").Add(jenUnsafePointer()),
	)

	f.Line()
}

func generateLibToC(f *jen.File, name string, cType string) {
	f.Commentf("ToC returns a pointer to the C %s that represents the %s.", cType, name)

	f.
		Func().
		Parens(jen.Id("recv").Op("*").Id(name)).
		Id("ToC").
		Params().
		Add(jenUnsafePointer()).
		Block(jen.Return().Id("recv").Dot("native"))
}

func generateLibNewFromC(f *jen.File, name string, cType string) {
	f.Commentf("%sNewFromC creates a new %s from a pointer to the C %s that represents the %s.",
		name, name, cType, name)

	f.
		Func().
		Id(name + "NewFromC").
		Params(jen.Id("native").Add(jenUnsafePointer())).
		Op("*").Id(name).
		Block(jen.
			Return().Op("&").Id(name).Values(jen.Dict{
			jen.Id("native"): jen.Id("native"),
		}))
}
