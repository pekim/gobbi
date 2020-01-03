package generate

import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
)

type Union struct {
	Name          string `xml:"name,attr"`
	CSymbolPrefix string `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefix,attr"`
	CType         string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	GlibTypeName  string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GlibGetType   string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	//Doc            *Doc         `xml:"doc"`
	//Fields         Fields       `xml:"field"`

	namespace *Namespace
	blacklist bool
}

func (u *Union) init(ns *Namespace) {
	u.namespace = ns
	u.applyAddenda()
}

func (u *Union) generateLib(f *jen.File, version semver.Version) {
	if u.blacklist {
		f.Commentf("UNSUPPORTED : %s : blacklisted", u.Name)
		return
	}

	generateLibStructType(f, "union", u.Name, u.CType, "")
	generateLibToC(f, u.Name, u.CType)
}
