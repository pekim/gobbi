package generate

import "fmt"

type Record struct {
	Name           string `xml:"name,attr"`
	Version        string `xml:"version,attr"`
	CSymbolPrefix  string `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefix,attr"`
	CType          string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	ParentName     string `xml:"parent,attr"`
	GlibTypeName   string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GlibGetType    string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	GlibTypeStruct string `xml:"http://www.gtk.org/introspection/glib/1.0 type-struct,attr"`
	Doc            *Doc   `xml:"doc"`
	//Fields         Fields       `xml:"field"`
	//Constructors   Constructors `xml:"constructor"`
	Functions Functions `xml:"function"`
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

	//r.Constructors.init(ns, r)
	r.Functions.init(ns /*r.GoName*/)
	//r.Methods.init(ns, r)
	//r.Fields.init(ns)
	//r.Signals.init(ns, r)
}

func (r *Record) generate(f *file) {
}
