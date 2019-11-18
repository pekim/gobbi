package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
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
	//Methods        Methods      `xml:"method"`
	//Signals        Signals      `xml:"http://www.gtk.org/introspection/glib/1.0 signal"`

	goName           string `xml:"goname,attr"`
	namespace        *Namespace
	newFromCFuncName string
}

func (r *Record) init(ns *Namespace) {
	r.namespace = ns
	r.goName = r.Name
	r.newFromCFuncName = fmt.Sprintf("%sNewFromC", r.Name)

	r.Constructors.init(ns, r)
	r.Functions.init(ns /*r.GoName*/)
	//r.Methods.init(ns, r)
	r.Fields.init(ns)
	//r.Signals.init(ns, r)
}

func (r *Record) generate(f *file) {
	r.generateType(f)
	r.Constructors.generate(f)
}

func (r *Record) generateType(f *file) {
	f.
		Type().
		Id(r.goName).
		Add(structFunc(r.generateFields))

	f.Line()
}

func (r *Record) generateFields(g *group) {
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

func (r *Record) createFromArgument(g *jen.Group, argValue *jen.Statement) {
	g.
		Op("&").
		Id(r.goName).
		Values(jen.Dict{
			jen.Id("native"): argValue,
		})
}
