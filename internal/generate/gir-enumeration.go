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
	Members Members `xml:"member"`
	//Doc          *Doc      `xml:"doc"`

	blacklist  bool
	namespace  *Namespace
	version    semver.Version
	goTypeName string
}

type Members []*Member

type Member struct {
	Name        string `xml:"name,attr"`
	Value       int    `xml:"value,attr"`
	CIdentifier string `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	//Doc         *Doc   `xml:"doc"`

	namespace *Namespace
}

func (e *Enumeration) init(ns *Namespace) {
	e.namespace = ns
	e.applyAddenda()
	e.version = versionNew(e.Version)
	e.namespace.versions.add(e.version)
	e.goTypeName = makeExportedGoName(e.Name)
}

func (e *Enumeration) generateLib(f *jen.File, version semver.Version, typeName string) {
	if e.blacklist {
		f.Commentf("UNSUPPORTED : %s : blacklisted", e.Name)
		return
	}

	if !generateEntityForVersion(version, e.version) {
		return
	}

	f.Commentf("%s is a representation of the C %s %s.", e.Name, typeName, e.CType)
	f.Type().Id(e.Name).Int()

	for _, member := range e.Members {
		name := e.Name + "_" + member.Name

		f.Commentf("%s is a representation of the C %s member %s.", name, typeName, member.CIdentifier)
		f.Const().Id(name).Op("=").Id(e.Name).Parens(jen.Lit(member.Value))
	}
}
