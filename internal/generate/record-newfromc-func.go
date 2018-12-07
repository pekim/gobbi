package generate

import (
	"github.com/dave/jennifer/jen"
)

type RecordNewFromCFunc struct {
	*Record
}

func (r *RecordNewFromCFunc) generate(g *jen.Group) {
	g.
		Func().
		Id(r.newFromCFuncName).
		ParamsFunc(r.generateParams).
		ParamsFunc(r.generateReturnDeclaration).
		BlockFunc(func(g *jen.Group) {
			// convert unsafe.Pointer arg to C type
			g.
				Id("c").
				Op(":=").
				Parens(jen.Id("*").Qual("C", r.CType)).
				Parens(jen.Id("u"))

			// If the C parameter is nil, do not create a Go struct.
			// Instead, return nil.
			g.If(jen.Id("c").Op("==").Nil()).
				Block(jen.Return(jen.Nil()))
			g.Line()

			r.generateCreateGoStruct(g)
			g.Line()

			g.
				Return().
				Id("g")
		})

	g.Line()
}

func (r *RecordNewFromCFunc) generateParams(g *jen.Group) {
	g.
		Id("u").
		Qual("unsafe", "Pointer")
}

func (r *RecordNewFromCFunc) generateReturnDeclaration(g *jen.Group) {
	g.
		Op("*").
		Id(r.GoName)
}

func (r *RecordNewFromCFunc) generateCreateGoStruct(g *jen.Group) {
	g.
		Id("g").
		Op(":=").
		Op("&").
		Id(r.GoName).
		Values(jen.DictFunc(r.generateStructValues))
}

func (r *RecordNewFromCFunc) generateStructValues(d jen.Dict) {
	d[jen.Id("native")] = jen.Id("c")

	if r.FieldsPrivate {
		return
	}

	for _, f := range r.Fields {
		if supported, _ := f.supported(); !supported {
			continue
		}

		cValue := jen.
			Id("c").
			Op(".").
			Id(f.Name)

		d[jen.Id(f.goVarName)] = f.Type.generator.generateCToGo(
			f.Namespace.fullGoPackageName, cValue)
	}
}
