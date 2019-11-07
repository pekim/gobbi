package generate

import "github.com/pekim/jennifer/jen"

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
		Id("ToC").
		Params(). // no params
		ParamsFunc(r.generateReturnDeclaration).
		BlockFunc(r.generateBody)

	g.Line()
}

func (r *RecordToCFunc) generateReturnDeclaration(g *jen.Group) {
	g.Qual("unsafe", "Pointer")
}

func (r *RecordToCFunc) generateBody(g *jen.Group) {
	if !r.FieldsPrivate {
		r.generateFieldsAssignment(g)
	}

	g.Line()

	g.
		Return().
		Parens(jen.Qual("unsafe", "Pointer")).
		Parens(jen.Id("recv").Op(".").Id("native"))
}

func (r *RecordToCFunc) generateFieldsAssignment(g *jen.Group) {
	for _, f := range r.Fields {
		if supported, _ := f.supported(); !supported {
			continue
		}

		if supported, _ := f.Type.generator.isSupportedAsField(); !supported {
			continue
		}

		goReference := jen.
			Id("recv").
			Op(".").
			Id(f.goVarName)

		g.
			Id("recv").
			Op(".").
			Id("native").
			Op(".").
			Id(f.Name).
			Op("=")
		f.Type.generator.generateGoToC(g, goReference)
	}
}
