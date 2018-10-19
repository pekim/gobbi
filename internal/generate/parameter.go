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

	goVarName string
	cVarName  string
}

func (p *Parameter) init(ns *Namespace) {
	p.Namespace = ns
	p.goVarName = makeGoName(p.Name)
	p.cVarName = "c_" + p.Name

	if p.Type != nil {
		p.Type.init(ns)
	}
}

func (p *Parameter) isSupported() (bool, string) {
	if p.Varargs != nil {
		return false, "varargs"
	}

	if p.Type == nil {
		return false, "no param type"
	}
	if p.Type.generator == nil {
		return false, fmt.Sprintf("no type generator for %s, %s", p.Type.Name, p.Type.CType)
	}

	if supported, reason := p.Type.generator.isSupportedAsParam(p.Direction); !supported {
		return false, reason
	}

	return true, ""
}

func (p *Parameter) isSupportedC() (bool, string) {
	if p.Type == nil {
		return false, "no param type"
	}
	if p.Type.generator == nil {
		return false, fmt.Sprintf("no type generator for %s, %s", p.Type.Name, p.Type.CType)
	}

	if supported, reason := p.Type.generator.isSupportedAsParamC(); !supported {
		return false, fmt.Sprintf("type %s : %s", p.Type.Name, reason)
	}

	return true, ""
}

func (p *Parameter) generateFunctionDeclaration(g *jen.Group) {
	if p.Direction == "out" {
		return
	}

	p.Type.generator.generateDeclaration(g, p.goVarName)
}

func (p *Parameter) generateFunctionDeclarationCtype(g *jen.Group) {
	p.Type.generator.generateDeclarationC(g, p.cVarName)
}

func (p *Parameter) generateCVar(g *jen.Group) {
	if p.Direction == "out" {
		p.Type.generator.generateParamOutCVar(g, p.cVarName)
	} else {
		p.Type.generator.generateParamCVar(g, p.cVarName, p.goVarName, p.TransferOwnership)
	}

	g.Line()
}

func (p *Parameter) generateCallArgument(g *jen.Group) {
	if p.Direction == "out" {
		p.Type.generator.generateParamOutCallArgument(g, p.cVarName)
	} else {
		p.Type.generator.generateParamCallArgument(g, p.cVarName)
	}
}

func (p *Parameter) generateOutputParamGoVar(g *jen.Group) {
	pkg := ""
	if p.Namespace != p.Type.qname.ns {
		pkg = p.Type.qname.ns.fullGoPackageName
	}

	if p.Direction == "out" || (p.Direction == "inout" && p.Type.Name == "argcargv") {
		p.Type.generator.generateReturnCToGo(g, true,
			p.cVarName, p.goVarName, pkg, p.TransferOwnership)
		g.Line()
	}
}

func (p *Parameter) generateOutputParamReturnDeclaration(g *jen.Group) {
	if p.Direction == "out" || (p.Direction == "inout" && p.Type.Name == "argcargv") {
		p.Type.generator.generateReturnFunctionDeclaration(g)
	}
}

func (p *Parameter) generateOutputParamReturn(g *jen.Group) {
	if p.Direction == "out" || (p.Direction == "inout" && p.Type.Name == "argcargv") {
		if p.Type.Name != "ignore" {
			g.Id(p.goVarName)
		}
	}
}
