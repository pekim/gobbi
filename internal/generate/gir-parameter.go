package generate

import "C"
import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type Parameter struct {
	Name      string `xml:"name,attr"`
	Direction string `xml:"direction,attr"`
	Argument
	AllowNone bool `xml:"allow-none,attr"`
	//Doc       *Doc      `xml:"doc"`
	Varargs *struct{} `xml:"varargs"`

	goVarName string
	namespace *Namespace

	lengthParam    *Parameter
	lengthForParam *Parameter
	argcParam      *Parameter
	argvParam      *Parameter
}

func (p *Parameter) init(ns *Namespace) {
	p.namespace = ns
	p.Type.init(ns)
	p.Array.init(ns)
	p.goVarName = makeUnexportedGoName(p.Name)
}

func (p *Parameter) isIn() bool {
	return p.Direction == "" || p.Direction == "in" || p.Direction == "inout"
}

func (p *Parameter) isOut() bool {
	return p.Direction == "out" || p.Direction == "inout"
}

func (p *Parameter) isInOut() bool {
	return p.Direction == "inout"
}

func (p *Parameter) sysParamGoType() *jen.Statement {
	if p.Type != nil {
		return p.Type.sysParamGoType()
	}

	if p.Array != nil {
		return p.Array.sysParamGoType()
	}

	panic(fmt.Sprintf("Parameter is not a type or an array: %s", p.Name))
}

func (p *Parameter) generateSysCArg(g *jen.Group, goVarName string, cVarName string) {
	if p.Type.isString() {
		p.generateSysCArgString(g, goVarName, cVarName)
		return
	}

	cValue := p.generateSysCValue(goVarName)
	g.Id(cVarName).Op(":=").Add(cValue)
}

func (p *Parameter) generateSysCValue(goVarName string) *jen.Statement {
	goValue := jen.Id(goVarName)

	if p.Type.CType == "gboolean" {
		return jen.Id("toCBool").Call(goValue)
	}

	if p.Type.cIndirectionCount > 0 {
		goValue = jenUnsafePointer().Call(goValue)
	}

	// Atoms are really pointers underneath.
	if p.Type.CType == "GdkAtom" {
		goValue = jenUnsafePointer().Call(goValue)
	}

	return jen.Parens(p.Type.jenGoCType()).Parens(goValue)
}

func (p *Parameter) generateSysCArgString(g *jen.Group, goVarName string, cVarName string) {
	if p.Type.cIndirectionCount == 1 {
		p.generateSysCArgStringSimple(g, goVarName, cVarName)
		return
	}

	if p.Type.cIndirectionCount == 2 {
		p.generateSysCArgStringPointer(g, goVarName, cVarName)
		return
	}

	panic(fmt.Sprintf("Unsupported indirection count (%d) for string param ", p.Type.cIndirectionCount))
}

func (p *Parameter) generateSysCArgStringSimple(g *jen.Group, goVarName string, cVarName string) {
	cValue := jen.Parens(p.Type.jenGoCType()).Parens(jen.Qual("C", "CString").Call(jen.Id(goVarName)))
	g.Id(cVarName).Op(":=").Add(cValue)

	if p.transferNone() {
		g.Defer().Qual("C", "free").Call(jenUnsafePointer().Call(jen.Id(cVarName)))
	}
}

func (p *Parameter) generateSysCArgStringPointer(g *jen.Group, goVarName string, cVarName string) {
	cStringVarName := cVarName + "String"

	// var cValue0String *C.gchar
	g.Var().Id(cStringVarName).Op("*").Qual("C", "gchar")

	if p.isIn() {
		cValue := jen.Parens(jen.Op("*").Qual("C", "gchar")).Parens(jen.Qual("C", "CString").Call(jen.Op("*").Id(goVarName)))
		g.Id(cStringVarName).Op("=").Add(cValue)
	}

	g.Id(cVarName).Op(":=").Op("&").Id(cStringVarName)

	if p.isIn() && p.transferNone() {
		g.Defer().Qual("C", "free").Call(jenUnsafePointer().Call(jen.Id(cVarName)))
	}
}

func (p *Parameter) generateSysCArgOut(g *jen.Group, goVarName string, cVarName string) {
	if !p.isOut() {
		return
	}

	if p.Type.isString() && p.Type.cIndirectionCount == 2 {
		p.generateSysCArgStringPointerOut(g, goVarName, cVarName)
	}
}

func (p *Parameter) generateSysCArgStringPointerOut(g *jen.Group, goVarName string, cVarName string) {
	cStringVarName := cVarName + "String"
	goStringVarName := goVarName + "String"

	g.Id(goStringVarName).Op(":=").Qual("C", "GoString").Call(jen.Id(cStringVarName))
	g.Op("*").Id(goVarName).Op("=").Id(goStringVarName)
}
