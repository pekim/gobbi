package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type Record struct {
	Namespace *Namespace

	Blacklist     bool `xml:"blacklist,attr"`
	FieldsPrivate bool `xml:"fieldsPrivate,attr"`

	Name           string       `xml:"name,attr"`
	GoName         string       `xml:"goname,attr"`
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
	Signals        Signals      `xml:"http://www.gtk.org/introspection/glib/1.0 signal"`

	newFromCFuncName string
}

func (r *Record) init(ns *Namespace) {
	r.Namespace = ns
	r.GoName = r.Name
	r.newFromCFuncName = fmt.Sprintf("%sNewFromC", r.Name)

	r.Constructors.init(ns, r)
	r.Functions.init(ns, r.GoName)
	r.Methods.init(ns, r)
	r.Fields.init(ns)
	r.Signals.init(ns, r)
}

func (r *Record) version() string {
	return r.Version
}

func (r *Record) blacklisted() (bool, string) {
	return r.Blacklist, r.CType
}

func (r *Record) supported() (supported bool, reason string) {
	if r.CType == "" {
		return false, fmt.Sprintf("%s : no CType", r.Name)
	}

	return true, ""
}

func (r *Record) mergeAddenda(addenda *Record) {
	r.Blacklist = addenda.Blacklist
	r.FieldsPrivate = addenda.FieldsPrivate

	if addenda.Version != "" {
		r.Version = addenda.Version
	}

	r.Constructors.mergeAddenda(addenda.Constructors)
	r.Functions.mergeAddenda(addenda.Functions)
	r.Methods.mergeAddenda(addenda.Methods)
}

func (r *Record) generate(g *jen.Group, version *Version) {
	if supportedByVersion(r, version) {
		//r.generateType(g)
		//(&RecordNewFromCFunc{r}).generate(g)
		//(&RecordToCFunc{r}).generate(g)
		(&RecordEqualFunc{r}).generate(g)
		r.generateUpcasts(g)
		r.generateDowncast(g)
	}

	if r.Version == "" || version.GTE(VersionNew(r.Version)) {
		r.Signals.generate(g, version, r.Version)
		r.Constructors.generate(g, version)
		r.Functions.generate(g, version)
		r.Methods.generate(g, version)
	}
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

			if r.FieldsPrivate {
				g.Comment("All fields are private")
			} else {
				r.Fields.generate(g)
			}
		})
	g.Line()
}

func (r *Record) root() *Record {
	if r.ParentName == "" {
		return r
	}

	qname := QNameNew(r.Namespace, r.ParentName)
	parent, found := qname.namespace.recordOrClassRecordForName(qname.name)
	if !found {
		panic(fmt.Sprintf("Failed to find parent %s for %s", r.ParentName, r.Name))
	}
	return parent.root()
}

func (r *Record) isDerivedFrom(goPackageName, goTypeName string) bool {
	rec := r

	for {
		if rec.Namespace.goPackageName == goPackageName && rec.GoName == goTypeName {
			return true
		}

		if rec.ParentName == "" {
			return false
		}

		qname := QNameNew(rec.Namespace, rec.ParentName)
		parent, found := qname.namespace.recordOrClassRecordForName(qname.name)
		if !found {
			panic(fmt.Sprintf("Failed to find parent %s for %s", rec.ParentName, rec.Name))
		}

		rec = parent
	}
}

func (r *Record) generateC(g *jen.Group, version *Version) {
	if supportedByVersion(r, version) {
		r.generateType(g)
		(&RecordNewFromCFunc{r}).generate(g)
		(&RecordToCFunc{r}).generate(g)
		//(&RecordEqualFunc{r}).generate(g)
		//r.generateUpcasts(g)
		//r.generateDowncast(g)
	}
}
