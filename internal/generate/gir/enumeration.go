package gir

import (
	"strings"
	"unicode"
	"unicode/utf8"

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
	Members      []*Member `xml:"member"`

	goTypeName string
}

func (e *Enumeration) fixup(ns *Namespace) {
	e.Namespace = ns

	e.goTypeName = e.Name
	if e.GoTypeName != "" {
		e.goTypeName = e.GoTypeName
	}

	for _, member := range e.Members {
		member.Namespace = ns
	}
}

func (e *Enumeration) version() string {
	return e.Version
}

func (e *Enumeration) generate(g *jen.Group, version *Version) {
	if !supportedByVersion(e, version) {
		return
	}

	if e.Blacklist {
		g.Commentf("Blacklisted enum : %s", e.CType)
		g.Line()
		return
	}

	memberNamePrefix := e.Namespace.CIdentifierPrefixes + "_"

	// define the type
	g.
		Type().
		Id(e.goTypeName).
		Qual("C", e.CType)

	g.Const().DefsFunc(func(g *jen.Group) {
		for _, member := range e.Members {
			if member.Blacklist {
				g.Commentf("Blacklisted member : %s", member.CIdentifier)
				return
			}

			name := strings.TrimPrefix(member.CIdentifier, memberNamePrefix)
			firstRune, _ := utf8.DecodeRuneInString(name)
			if unicode.IsDigit(firstRune) {
				// Make name a legal Go identifier (cannot start with a digit).
				name = "_" + name
			}

			// declare a member constant
			g.
				Id(name).
				Id(e.goTypeName).
				Op("=").
				Lit(member.Value)
		}
	})

	g.Line()
}

type Member struct {
	Namespace *Namespace

	Name        string `xml:"name,attr"`
	Blacklist   bool   `xml:"blacklist,attr"`
	Value       int    `xml:"value,attr"`
	CIdentifier string `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	GlibNick    string `xml:"http://www.gtk.org/introspection/glib/1.0 nick,attr"`
	Doc         *Doc   `xml:"doc"`
}

func (m *Member) blacklisted() bool {
	return m.Blacklist
}
