package generate

import "C"
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
			r.generateObjectRefManagement(g)
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

func (r *RecordNewFromCFunc) generateObjectRefManagement(g *jen.Group) {
	if !r.isDerivedFrom("gobject", "Object") {
		return
	}

	if r.Name == "Object" {
		return
	}

	g.Line()

	g.
		// ug := (C.gpointer)(u)
		Id("ug").
		Op(":=").
		Parens(jen.Qual("C", "gpointer")).
		Parens(jen.Id("u"))

	g.
		//if C.g_object_is_floating(ug) == C.TRUE {
		If(jen.
			Qual("C", "g_object_is_floating").
			Call(jen.Id("ug")).
			Op("==").
			Qual("C", "TRUE")).
		BlockFunc(func(g *jen.Group) {
			//	C.g_object_ref_sink(ug)
			g.
				Qual("C", "g_object_ref_sink").
				Call(jen.Id("ug"))
		}).
		Else().
		BlockFunc(func(g *jen.Group) {
			//	C.g_object_ref(ug)
			g.
				Qual("C", "g_object_ref").
				Call(jen.Id("ug"))
		})

	g.
		//runtime.SetFinalizer(
		Qual("runtime", "SetFinalizer").
		Call(
			// g, func(o *SomeGoType) {
			jen.Id("g"),
			jen.
				Func().
				Params(jen.Id("o").Op("*").Id(r.GoName)).
				BlockFunc(func(g *jen.Group) {
					//g.
					//	Qual("fmt", "Println").
					//	Call(jen.Lit("finalizer run"), jen.Lit(r.GoName))

					//	C.g_object_unref(o.native)
					g.
						Qual("C", "g_object_unref").
						Call(jen.
							Parens(jen.Qual("C", "gpointer")).
							Parens(jen.Id("o").Dot("native")))

				}),
		)
}
