package generate

import "github.com/dave/jennifer/jen"

type Alias struct {
	Namespace *Namespace

	Name      string `xml:"name,attr"`
	Blacklist bool   `xml:"blacklist,attr"`
	CType     string `xml:"type,attr"`
	Type      Type   `xml:"type"`
	Doc       *Doc   `xml:"doc"`
}

func (a *Alias) init(ns *Namespace) {
	a.Namespace = ns
	a.Type.Namespace = ns
}

func (a *Alias) version() string {
	return ""
}

func (a *Alias) mergeAddenda(addenda *Alias) {
	a.Blacklist = addenda.Blacklist
}

func (a Alias) generate(g *jen.Group, version *Version) {
	if a.Blacklist {
		g.Commentf("Blacklisted alias : %s", a.CType)
		g.Line()
		return
	}

	// goDoc{file.Group}.linesFromDoc(a.gir.Doc)

	g.
		Type().
		Id(a.Name).
		Id(a.Type.Name)

	g.Line()
}
