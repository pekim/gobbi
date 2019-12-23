package generate

import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
)

type Record struct {
	Name           string `xml:"name,attr"`
	Version        string `xml:"version,attr"`
	CSymbolPrefix  string `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefix,attr"`
	CType          string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	GlibTypeName   string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GlibGetType    string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	GlibTypeStruct string `xml:"http://www.gtk.org/introspection/glib/1.0 type-struct,attr"`
	//Doc            *Doc         `xml:"doc"`
	ParentName string `xml:"parent,attr"`
	//Fields         Fields       `xml:"field"`
	Constructors Constructors `xml:"constructor"`
	Functions    Functions    `xml:"function"`
	Methods      Methods      `xml:"method"`
	//Signals        Signals      `xml:"http://www.gtk.org/introspection/glib/1.0 signal"`

	goName    string
	blacklist bool
	namespace *Namespace
	version   semver.Version
	//newFromNativeName string
	//parentNamespace   *Namespace
	//parent            *Class
}

func (r *Record) init(ns *Namespace) {
	r.namespace = ns
	r.applyAddenda()
	r.version = versionNew(r.Version)
	r.namespace.versions.add(r.version)
	r.Constructors.init(ns, r)
	r.Methods.init(ns, r)
	r.Functions.init(ns)

	r.goName = r.Name
}

func (r Record) generateSysType(f *jen.File, version semver.Version) {
	if r.blacklist {
		f.Commentf("UNSUPPORTED : %s : blacklisted", r.Name)
		return
	}

	if r.version.GT(version) {
		return
	}

	// GEN: type SomeRecord SomeCType
	f.Type().Id(r.Name).Qual("C", r.CType)
}

func (r *Record) generateSys(f *jen.File, version semver.Version) {
	r.Constructors.generateSys(f, version)
	r.Methods.generateSys(f, version)
	r.Functions.generateSys(f, version)
}
