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

func (p *Parameter) generateSysCValue(goVarName string) *jen.Statement {
	if p.Type.isString() {
		return jen.Lit(42)
		//return p.generateSysStringCValue(goVarName)
	}

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

//func (p *Parameter) generateSysStringCValue(goVarName string) *jen.Statement {
//	if p.Type.cIndirectionCount > 0 {
//		return jen.Lit(42)
//	}
//
//}
