package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/pekim/gobbi/internal/gi"
)

type Record struct {
	Name           string       `xml:"name,attr"`
	Version        string       `xml:"version,attr"`
	CSymbolPrefix  string       `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefix,attr"`
	CType          string       `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	ParentName     string       `xml:"parent,attr"`
	GlibTypeName   string       `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GlibGetType    string       `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	GlibTypeStruct string       `xml:"http://www.gtk.org/introspection/glib/1.0 type-struct,attr"`
	Doc            *Doc         `xml:"doc"`
	Fields         Fields       `xml:"field"`
	Constructors   Constructors `xml:"constructor"`
	Functions      Functions    `xml:"function"`
	Methods        Methods      `xml:"method"`
	//Signals        Signals      `xml:"http://www.gtk.org/introspection/glib/1.0 signal"`

	goName           string
	namespace        *Namespace
	newFromCFuncName string

	structInfoGoName        string
	structInfoOnceGoName    string
	structInfoSetFuncGoName string
}

func (r *Record) init(ns *Namespace) {
	r.namespace = ns
	r.goName = r.Name
	r.newFromCFuncName = fmt.Sprintf("%sNewFromC", r.Name)

	r.structInfoGoName = fmt.Sprintf("%sStruct", makeUnexportedGoName(r.Name))
	r.structInfoOnceGoName = fmt.Sprintf("%sOnce", r.structInfoGoName)
	r.structInfoSetFuncGoName = fmt.Sprintf("%sSet", r.structInfoGoName)

	r.Constructors.init(ns, r)
	r.Functions.init(ns /*r.GoName*/)
	r.Methods.init(ns, r)
	r.Fields.init(ns)
	//r.Signals.init(ns, r)
}

func (r *Record) generate(f *file) {
	r.generateStructInfo(f)
	r.generateType(f)
	r.Constructors.generate(f)
	r.Methods.generate(f)
}

func (r *Record) generateStructInfo(f *file) {
	// var colorStruct *gi.Struct
	f.
		Var().
		Id(r.structInfoGoName).
		Op("*").
		Qual(gi.PackageName, "Struct")

	// var colorStructOnce sync.Once
	f.
		Var().
		Id(r.structInfoOnceGoName).
		Qual("sync", "Once")

	// func colorStructSet() {
	//   ...
	// }
	f.
		Func().
		Id(r.structInfoSetFuncGoName).
		Params().
		BlockFunc(func(g *jen.Group) {
			//   colorStructOnce.Do(func() {
			//     colorStruct = gi.StructNew("Gdk", "Color")
			//   })
			g.
				Id(r.structInfoOnceGoName).
				Dot("Do").
				Call(jen.
					Func().
					Params().
					Block(jen.
						Id(r.structInfoGoName).
						Op("=").
						Qual(gi.PackageName, "StructNew").
						Call(jen.Lit(r.namespace.Name),
							jen.Lit(r.Name))))
		})
}

func (r *Record) generateType(f *file) {
	f.
		Type().
		Id(r.goName).
		StructFunc(r.generateFields)

	f.Line()
}

func (r *Record) generateFields(g *jen.Group) {
	g.Id("native").Uintptr()
	r.Fields.generate(g)
}

func (r *Record) supportedAsOutParameter() bool {
	return true
}

func (r *Record) generateDeclaration() *jen.Statement {
	return jen.Op("*").Id(r.goName)
}

func (r *Record) argumentGetFunctionName() string {
	return "Pointer"
}

func (r *Record) argumentSetFunctionName() string {
	return "TODOsfn"
}

func (r *Record) createFromArgument(s *jen.Statement, argValue *jen.Statement) {
	s.
		Op("&").
		Id(r.goName).
		Values(jen.Dict{
			jen.Id("native"): argValue,
		})
}
