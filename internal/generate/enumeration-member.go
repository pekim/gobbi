package generate

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/dave/jennifer/jen"
)

type Member struct {
	Namespace *Namespace

	Name        string `xml:"name,attr"`
	Blacklist   bool   `xml:"blacklist,attr"`
	Value       int    `xml:"value,attr"`
	CIdentifier string `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	GlibNick    string `xml:"http://www.gtk.org/introspection/glib/1.0 nick,attr"`
	Doc         *Doc   `xml:"doc"`

	ns     *Namespace
	goName string
}

func (m *Member) init(ns *Namespace, namePrefix string) {
	m.ns = ns

	name := strings.TrimPrefix(m.CIdentifier, namePrefix)
	firstRune, _ := utf8.DecodeRuneInString(name)
	if unicode.IsDigit(firstRune) {
		// Make name a legal Go identifier (cannot start with a digit).
		name = "_" + name
	}

	m.goName = name
}

func (m *Member) mergeAddenda(addenda *Member) {
	m.Blacklist = addenda.Blacklist
}

func (m *Member) generate(g *jen.Group, goTypeName string) {
	if m.Blacklist {
		g.Commentf("Blacklisted member : %s", m.CIdentifier)
		return
	}

	// declare a member constant
	g.
		Id(m.goName).
		Id(goTypeName).
		Op("=").
		Lit(m.Value)
}
