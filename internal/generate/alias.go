package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type Alias struct {
	Namespace *Namespace

	Name      string `xml:"name,attr"`
	Blacklist bool   `xml:"blacklist,attr"`
	CType     string `xml:"type,attr"`
	Type      *Type  `xml:"type"`
	Doc       *Doc   `xml:"doc"`

	goName string
}

func (a *Alias) init(ns *Namespace) {
	a.Namespace = ns
	//a.goName = makeExportedGoName(a.Name)
	a.goName = a.Name

	if a.Type != nil {
		a.Type.init(ns)
	}
}

func (a Alias) version() string {
	return ""
}

func (a *Alias) blacklisted() (bool, string) {
	return a.Blacklist, a.CType
}

func (a *Alias) supported() (supported bool, reason string) {
	if a.Type == nil {
		return false, "alias has no param type"
	}
	if a.Type.generator == nil {
		return false, fmt.Sprintf("alias has no type generator for %s, %s", a.Type.Name, a.Type.CType)
	}

	return true, ""
}

func (a *Alias) mergeAddenda(addenda *Alias) {
	a.Blacklist = addenda.Blacklist
}

func (a Alias) generate(g *jen.Group, version *Version) {
	if !supportedByVersion(a, version) {
		return
	}

	g.Commentf("%s is a representation of the C alias of the same name.", a.Name)

	g.Type()
	a.Type.generator.generateDeclaration(g, a.goName)

	g.Line()
}
