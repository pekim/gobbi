package generate

import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
)

type Enumeration struct {
	Name         string `xml:"name,attr"`
	Blacklist    bool   `xml:"blacklist,attr"`
	GoTypeName   string `xml:"goname,attr"` // used in addenda files
	Version      string `xml:"version,attr"`
	CType        string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	GlibTypeName string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GlibGetType  string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`

	//Functions    Functions `xml:"function"`
	//Members      Members   `xml:"member"`
	//Doc          *Doc      `xml:"doc"`

	blacklist  bool
	namespace  *Namespace
	version    semver.Version
	goTypeName string
}

func (e *Enumeration) init(ns *Namespace) {
	e.namespace = ns
	e.applyAddenda()
	e.version = versionNew(e.Version)
	e.namespace.versions.add(e.version)
}

func (e *Enumeration) generateSys(f *jen.File, version semver.Version) {
	if e.blacklist {
		f.Commentf("UNSUPPORTED : %s : blacklisted", e.Name)
		return
	}

	if e.version.GT(version) {
		return
	}

	// GEN: type SomeEnum SomeCType
	f.Type().Id(e.Name).Qual("C", e.CType)
}
