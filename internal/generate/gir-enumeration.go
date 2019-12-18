package generate

import (
	"github.com/dave/jennifer/jen"
)

type Enumeration struct {
	Namespace *Namespace

	Name         string `xml:"name,attr"`
	Blacklist    bool   `xml:"blacklist,attr"`
	GoTypeName   string `xml:"goname,attr"` // used in addenda files
	Version      string `xml:"version,attr"`
	CType        string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	GlibTypeName string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GlibGetType  string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	//Doc          *Doc      `xml:"doc"`
	//Members      Members   `xml:"member"`
	//Functions    Functions `xml:"function"`

	goTypeName string
	version    *version
}

func (e *Enumeration) init(ns *Namespace) {
	e.Namespace = ns
	e.version = versionNew(e.Version)
}

func (e Enumeration) generateSys(f *jen.File) {
	f.
		Type().
		Id(e.Name).
		Qual("C", e.CType)
}
