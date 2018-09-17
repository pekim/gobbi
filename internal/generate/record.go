package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type Record struct {
	Namespace *Namespace

	Name           string         `xml:"name,attr"`
	Blacklist      bool           `xml:"blacklist,attr"`
	GoName         string         `xml:"goname,attr"`
	Version        string         `xml:"version,attr"`
	CSymbolPrefix  string         `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefix,attr"`
	CType          string         `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	ParentName     string         `xml:"parent,attr"`
	GlibTypeName   string         `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GlibGetType    string         `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	GlibTypeStruct string         `xml:"http://www.gtk.org/introspection/glib/1.0 type-struct,attr"`
	Doc            *Doc           `xml:"doc"`
	Fields         Fields         `xml:"field"`
	Constructors   []*Constructor `xml:"constructor"`
	Methods        Methods        `xml:"method"`

	newFromCFuncName string
}

func (r *Record) init(ns *Namespace) {
	r.Namespace = ns
	r.GoName = makeExportedGoName(r.Name)
	r.newFromCFuncName = fmt.Sprintf("%sNewFromC", makeGoName(r.Name))

	for _, ctor := range r.Constructors {
		ctor.init(ns)
	}

	for _, method := range r.Methods {
		method.init(ns)
	}

	for _, field := range r.Fields {
		field.init(ns)
	}
}

func (r *Record) version() string {
	return r.Version
}

func (r *Record) blacklisted() (bool, string) {
	return r.Blacklist, r.CType
}

func (r *Record) supported() (supported bool, reason string) {

	return true, ""
}

func (r *Record) mergeAddenda(addenda *Record) {
	r.Blacklist = addenda.Blacklist
}

func (r *Record) generate(g *jen.Group, version *Version) {
	r.generateType(g)
	r.generateNewFromCFunc(g)
}

func (r *Record) generateType(g *jen.Group) {
	g.Commentf("%s is a wrapper around the C record %s.", r.GoName, r.CType)

	g.
		Type().
		Id(r.GoName).
		StructFunc(func(g *jen.Group) {
			g.
				Id("native").
				Op("*").
				Qual("C", r.CType)

			r.Fields.generate(g)
		})
	g.Line()
}

func (r *Record) generateNewFromCFunc(g *jen.Group) {
	g.
		Func().
		Id(r.newFromCFuncName).
		Params(jen.
			Id("c").
			Op("*").
			Qual("C", r.CType),

			jen.
				Id("finalizeFree").
				Id("bool")).
		Params(jen.
			Op("*").
			Id(r.GoName)).
		BlockFunc(func(g *jen.Group) {
			// If the C parameter is nil, do not create a Go struct.
			g.If(
				jen.
					Id("c").
					Op("==").
					Nil()).
				Block(jen.Return(jen.Nil()))
			g.Line()

			g.
				Id("g").
				Op(":=").
				Op("&").
				Id(r.GoName).
				Values(
					jen.DictFunc(func(d jen.Dict) {
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
					}))
			g.Line()

			g.
				If(jen.Id("finalizeFree")).
				BlockFunc(func(g *jen.Group) {
					g.Qual("runtime", "SetFinalizer").
						Call(
							jen.Id("g"),
							jen.
								Func().
								Params(
									jen.
										Id("obj").
										Id("interface").
										Op("{}")).
								BlockFunc(func(g *jen.Group) {
									g.
										Qual("C", "g_free").
										Call(jen.
											Parens(jen.Qual("C", "gpointer")).
											Parens(jen.Id("c")))
								}))
				})

			g.Line()

			g.
				Return().
				Id("g")
		})

	g.Line()
}

type Constructor struct {
	*Function
}
