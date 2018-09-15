package generate

import (
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
}

func (r *Record) init(ns *Namespace) {
	r.Namespace = ns
	r.GoName = makeExportedGoName(r.Name)

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
		}).
		Line()
}

type Constructor struct {
	*Function
}
