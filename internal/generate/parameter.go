package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

type Parameter struct {
	Namespace *Namespace

	Name              string    `xml:"name,attr"`
	Direction         string    `xml:"direction,attr"`
	TransferOwnership string    `xml:"transfer-ownership,attr"`
	Nullable          bool      `xml:"nullable,attr"`
	AllowNone         bool      `xml:"allow-none,attr"`
	Doc               *Doc      `xml:"doc"`
	Type              *Type     `xml:"type"`
	Array             *Array    `xml:"array"`
	Varargs           *struct{} `xml:"varargs"`

	goVarName       string
	cVarName        string
	arrayLengthFor  *Parameter
	stringLengthFor *Parameter
}

func (p *Parameter) init(ns *Namespace) {
	p.Namespace = ns
	p.goVarName = makeGoName(p.Name)
	p.cVarName = "c_" + p.Name

	if p.Type != nil {
		p.Type.init(ns)
	}

	if p.Array != nil && p.Array.Type != nil {
		p.Array.init(ns)
	}
}

func (p *Parameter) isSupported() (bool, string) {
	if p.Varargs != nil {
		return false, "varargs"
	}

	if p.Type != nil {
		if p.Type.generator == nil {
			return false, fmt.Sprintf("no type generator for %s (%s) for param %s",
				p.Type.Name, p.Type.CType, p.Name)
		}

		if supported, reason := p.Type.generator.isSupportedAsParam(p.Direction); !supported {
			return false, reason
		}

		if p.arrayLengthFor != nil && strings.HasSuffix(p.Type.CType, "*") {
			return false, fmt.Sprintf("array length param %s is pointer (%s)",
				p.Name, p.Type.CType)
		}

		return true, ""
	}

	if p.Array != nil {
		if p.Direction == "out" {
			return false, fmt.Sprintf("output array param %s", p.Name)
		}

		if p.Array.Type.generator == nil {
			return false, fmt.Sprintf("no type generator for %s (%s) for array param %s",
				p.Array.Type.Name, p.Array.Type.CType, p.Name)
		}

		if supported, reason := p.Array.Type.generator.isSupportedAsArrayParam(p.Direction); !supported {
			return false, reason
		}

		return true, ""
	}

	return false, "no param type or array"
}

func (p *Parameter) isSupportedC() (bool, string) {
	if p.Type != nil {
		if p.Type.generator == nil {
			return false, fmt.Sprintf("no type generator for %s, %s", p.Type.Name, p.Type.CType)
		}

		if supported, reason := p.Type.generator.isSupportedAsParamC(); !supported {
			return false, fmt.Sprintf("type %s : %s", p.Type.Name, reason)
		}

		return true, ""
	}

	if p.Array != nil {
		if p.Array.Type.generator == nil {
			return false, fmt.Sprintf("no type generator for %s (%s) for array param %s",
				p.Array.Type.Name, p.Array.Type.CType, p.Name)
		}

		if supported, reason := p.Array.Type.generator.isSupportedAsArrayParamC(p.Direction); !supported {
			return false, reason
		}

		return true, ""
	}

	return false, "no param type or array"
}

func (p *Parameter) generateFunctionDeclaration(g *jen.Group) {
	if p.Direction == "out" {
		return
	}

	if p.stringLengthFor != nil {
		return
	}

	if p.Array != nil {
		p.Array.generateDeclaration(g, p.goVarName)
	} else {
		if p.arrayLengthFor != nil {
			return
		}

		p.Type.generator.generateDeclaration(g, p.goVarName)
	}
}

func (p *Parameter) generateFunctionDeclarationCtype(g *jen.Group) {
	if p.Array != nil {
		p.Array.generateDeclarationC(g, p.cVarName)
	} else {
		p.Type.generator.generateDeclarationC(g, p.cVarName)
	}
}

func (p *Parameter) generateCVar(g *jen.Group) {
	if p.Direction == "out" {
		p.Type.generator.generateParamOutCVar(g, p.cVarName)
	} else if p.Array != nil {
		p.Array.generateParamCVar(g, p.cVarName, p.goVarName, p.TransferOwnership)
	} else if p.arrayLengthFor != nil {
		p.Array.generateArrayLenParamCVar(g, p.cVarName, p.arrayLengthFor.goVarName, p.Type.CType)
	} else if p.stringLengthFor != nil {
		p.generateCVarForStringLength(g)
	} else {
		p.Type.generator.generateParamCVar(g, p.cVarName, p.goVarName, p.TransferOwnership)
	}

	g.Line()
}

func (p *Parameter) generateCVarForStringLength(g *jen.Group) {
	g.
		Id(p.cVarName).
		Op(":=").
		Parens(jen.Qual("C", p.Type.CType)).
		Parens(jen.Len(jen.Id(p.stringLengthFor.goVarName)))
}

func (p *Parameter) generateGoVar(g *jen.Group) {
	if p.arrayLengthFor != nil {
		return
	}

	if p.Array != nil {
		p.Array.generateParamGoVar(g, p)
	} else {
		p.Type.generator.generateParamGoVar(g, p.goVarName, p.cVarName, p.Type.fullGoPackageName())
	}

	g.Line()
}

func (p *Parameter) generateCallArgument(g *jen.Group) {
	if p.Direction == "out" {
		p.Type.generator.generateParamOutCallArgument(g, p.cVarName)
	} else {
		if p.Array != nil {
			p.Array.generateParamCallArgument(g, p.cVarName)
		} else {
			p.Type.generator.generateParamCallArgument(g, p.cVarName)
		}
	}
}

func (p *Parameter) generateOutputParamGoVar(g *jen.Group) {
	if p.Direction == "out" || (p.Direction == "inout" && p.Type != nil && p.Type.Name == "argcargv") {
		p.Type.generator.generateReturnCToGo(g, true,
			p.cVarName, p.goVarName, p.Type.fullGoPackageName(),
			p.TransferOwnership, false)
		g.Line()
	}
}

func (p *Parameter) generateOutputParamReturnDeclaration(g *jen.Group) {
	if p.Direction == "out" || (p.Direction == "inout" && p.Type != nil && p.Type.Name == "argcargv") {
		p.Type.generator.generateReturnFunctionDeclaration(g)
	}
}

func (p *Parameter) generateOutputParamReturn(g *jen.Group) {
	if p.Direction == "out" || (p.Direction == "inout" && p.Type != nil && p.Type.Name == "argcargv") {
		if p.Type.Name != "ignore" {
			g.Id(p.goVarName)
		}
	}
}
