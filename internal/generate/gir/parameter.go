package gir

import (
	"github.com/dave/jennifer/jen"
)

type Parameter struct {
	Namespace *Namespace

	Name              string    `xml:"name,attr"`
	Direction         string    `xml:"direction,attr"`
	TransferOwnership string    `xml:"transfer-ownership,attr"`
	Doc               *Doc      `xml:"doc"`
	Type              *Type     `xml:"type"`
	Array             *Array    `xml:"array"`
	Varargs           *struct{} `xml:"varargs"`

	goName string
}

func (p *Parameter) init(ns *Namespace) {
	p.Namespace = ns
	p.goName = makeGoName(p.Name)

	if p.Type != nil {
		p.Type.Namespace = ns
	}
}

func (p Parameter) generateFunctionDeclaration(g *jen.Group) {
	g.
		Id(p.goName).
		Id("int")
}

func (p Parameter) isSupported() (bool, string) {
	if p.Varargs != nil {
		return false, "varargs"
	}

	return true, ""
}
