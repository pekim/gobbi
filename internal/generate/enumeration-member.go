package generate

import (
	"github.com/dave/jennifer/jen"
)

type Member struct {
	Name        string `xml:"name,attr"`
	Value       int    `xml:"value,attr"`
	CIdentifier string `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	//Doc         *Doc   `xml:"doc"`

	namespace *Namespace
}

func (m *Member) init(ns *Namespace) {
	m.namespace = ns
}

func (m *Member) generate(g *jen.Group, goTypeName string) {
	goName := makeExportedGoName(m.CIdentifier)

	g.
		Id(goName).
		Id(goTypeName).
		Op("=").
		Lit(m.Value)
}
