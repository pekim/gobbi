package gir

import "github.com/dave/jennifer/jen"

type Alias struct {
	Namespace *Namespace

	Name      string `xml:"name,attr"`
	Blacklist bool   `xml:"blacklist,attr"`
	CType     string `xml:"type,attr"`
	Type      Type   `xml:"type"`
	Doc       *Doc   `xml:"doc"`
}

func (a *Alias) fixup(ns *Namespace) {
	a.Namespace = ns
	a.Type.Namespace = ns
}

func (a *Alias) blacklisted() bool {
	return a.Blacklist
}

func (a Alias) generate(g *jen.Group) {
	if a.blacklisted() {
		g.Commentf("Blacklisted alias : %s", a.CType)
		g.Line()
		return
	}

	// goDoc{file.Group}.linesFromDoc(a.gir.Doc)

	g.
		Type().
		Id(a.Name).
		Qual("C", a.CType)

	g.Line()
}
