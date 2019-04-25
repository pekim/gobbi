package generate

import (
	"github.com/dave/jennifer/jen"
)

type Enumeration struct {
	Namespace *Namespace

	Name         string    `xml:"name,attr"`
	Blacklist    bool      `xml:"blacklist,attr"`
	GoTypeName   string    `xml:"goname,attr"` // used in addenda files
	Version      string    `xml:"version,attr"`
	CType        string    `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	GlibTypeName string    `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GlibGetType  string    `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	Doc          *Doc      `xml:"doc"`
	Members      Members   `xml:"member"`
	Functions    Functions `xml:"function"`

	goTypeName string
}

func (e *Enumeration) init(ns *Namespace) {
	e.Namespace = ns

	e.goTypeName = e.Name
	if e.GoTypeName != "" {
		e.goTypeName = e.GoTypeName
	}

	memberNamePrefix := e.Namespace.CIdentifierPrefixes + "_"
	for _, member := range e.Members {
		member.init(ns, memberNamePrefix)
	}

	e.Functions.init(ns, e.Name)
}

func (e *Enumeration) version() string {
	return e.Version
}

func (e *Enumeration) mergeAddenda(addenda *Enumeration) {
	e.Blacklist = addenda.Blacklist
	e.GoTypeName = addenda.GoTypeName
	if addenda.CType != "" {
		e.CType = addenda.CType
	}
	if addenda.Version != "" {
		e.Version = addenda.Version
	}
	e.Members.mergeAddenda(addenda.Members)
	e.Functions.mergeAddenda(addenda.Functions)
}

func (e *Enumeration) blacklisted() (bool, string) {
	return e.Blacklist, e.CType
}

func (e *Enumeration) supported() (supported bool, reason string) {
	return true, ""
}

func (e *Enumeration) generate(g *jen.Group, version *Version) {
	if !supportedByVersion(e, version) {
		return
	}

	// define the type
	g.
		Type().
		Id(e.goTypeName).
		Int()
		//Qual("C", e.CType)

	// define members
	g.Const().DefsFunc(func(g *jen.Group) {
		for _, member := range e.Members {
			member.generate(g, e.goTypeName)
		}
	})

	e.Functions.generate(g, version)
}
