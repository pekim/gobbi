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

	paramType ParameterType
	goVarName string
	goType    string
	cVarName  string
}

func (p *Parameter) init(ns *Namespace) {
	p.Namespace = ns
	p.goVarName = makeGoName(p.Name)
	p.cVarName = "c_" + p.Name

	if p.Type != nil {
		p.Type.Namespace = ns

		goType, isInteger := integerCTypeMap[p.Type.CType]
		if isInteger {
			p.goType = goType
			p.paramType = ParameterTypeIntegerNew(p)
		} else if p.Type.Name == "utf8" {
			p.goType = "string"
			p.paramType = ParameterTypeStringNew(p)
		}
	}
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

func (p Parameter) generateFunctionDeclaration(g *jen.Group) {
	if p.Direction == "out" {
		return
	}

	p.paramType.generateFunctionDeclaration(g)
}

func (p Parameter) generateCVar(g *jen.Group) {
	if p.Direction == "out" {
		p.paramType.generateOutCVar(g)
	} else {
		p.paramType.generateCVar(g)
	}

	g.Line()
}

func (p Parameter) generateCallArgument(g *jen.Group) {
	p.paramType.generateCallArgument(g)
}
