package generate

import "github.com/dave/jennifer/jen"

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
			// If the C parameter is nil, do not create a Go struct.
			// Instead, return nil.
			g.If(jen.Id("c").Op("==").Nil()).
				Block(jen.Return(jen.Nil()))
			g.Line()

			r.generateCreateGoStruct(g)
			g.Line()

			g.If(jen.Id("finalizeFree")).
				BlockFunc(r.generateFinalizeFree)
			g.Line()

			g.
				Return().
				Id("g")
		})

	g.Line()
}

func (r *RecordNewFromCFunc) generateParams(g *jen.Group) {
	g.
		Id("c").
		Op("*").
		Qual("C", r.CType)

	g.
		Id("finalizeFree").
		Id("bool")
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

	for _, f := range r.Fields {
		if supported, _ := f.supported(); !supported {
			continue
		}

		cValue := jen.
			Id("c").
			Op(".").
			Id(f.Name)

		d[jen.Id(f.goVarName)] = f.Type.generator.generateCToGo(cValue)
	}
}

func (r *RecordNewFromCFunc) generateFinalizeFree(g *jen.Group) {
	g.Qual("runtime", "SetFinalizer").
		Call(
			jen.Id("g"),
			jen.
				Func().
				Params(jen.Id("obj").Id("interface").Op("{}")).
				BlockFunc(func(g *jen.Group) {
					g.
						Qual("C", "g_free").
						Call(jen.
							Parens(jen.Qual("C", "gpointer")).
							Parens(jen.Id("c")))
				}))
}
