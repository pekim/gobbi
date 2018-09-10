package generate

import (
	"fmt"

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

	goName   string
	goType   string
	cVarName string
}

func (p *Parameter) init(ns *Namespace) {
	p.Namespace = ns
	p.goName = makeGoName(p.Name)
	p.cVarName = "c_" + p.Name

	if p.Type != nil {
		p.Type.Namespace = ns

		goType, isPrimitive := primitiveCTypeMap[p.Type.CType]
		if isPrimitive {
			p.goType = goType
		} else if p.Type.Name == "utf8" {
			p.goType = "string"
		}
	}
}

func (p Parameter) generateFunctionDeclaration(g *jen.Group) {
	g.
		Id(p.goName).
		Id(p.goType)
}

func (p Parameter) generateAssignmentToCVar(g *jen.Group) {
	if p.Direction == "out" {
		return
	}

	if p.Type.Name == "utf8" {
		p.generateAssignmentToCVarString(g)
		return
	}
	p.generateAssignmentToCVarPrimitive(g)
}

func (p Parameter) generateAssignmentToCVarString(g *jen.Group) {
	g.
		Id(p.cVarName).
		Op(":=").
		Qual("C", "CString").
		Call(jen.Id(p.goName))

	g.
		Defer().
		Qual("C", "free").
		Call(jen.
			Qual("unsafe", "Pointer").
			Call(jen.Id(p.cVarName)))

	g.Line()
}

func (p Parameter) generateAssignmentToCVarPrimitive(g *jen.Group) {
	g.
		Id(p.cVarName).
		Op(":=").
		Parens(jen.Qual("C", p.Type.CType)).
		Parens(jen.Id(p.goName))

	g.Line()
}

func (p Parameter) generateCallArgument(g *jen.Group) {
	g.Id(p.cVarName)
}

func (p Parameter) isSupported() (bool, string) {
	if p.Varargs != nil {
		return false, "varargs"
	}

	if p.goType == "" {
		if p.Type != nil {
			return false, fmt.Sprintf("type %s, %s", p.Type.Name, p.Type.CType)
		} else {
			return false, "no type"
		}
	}

	return true, ""
}
