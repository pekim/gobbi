package generate

import "github.com/dave/jennifer/jen"

type RecordToCFunc struct {
	*Record
}

func (r *RecordToCFunc) generate(g *jen.Group) {
	g.
		Func().
		Parens(jen.
			Id("recv").
			Op("*").
			Id(r.GoName)).
		Id("toC").
		Params(). // no params
		ParamsFunc(r.generateReturnDeclaration).
		BlockFunc(r.generateBody)

	g.Line()
}

func (r *RecordToCFunc) generateReturnDeclaration(g *jen.Group) {
	g.
		Op("*").
		Qual("C", r.CType)
}

func (r *RecordToCFunc) generateBody(g *jen.Group) {
	g.
		Comment("TODO marshall fields to native").
		Line()

	g.
		Return().
		Id("recv").Op(".").Id("native")
}
