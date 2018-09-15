package generate

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

type TypeGeneratorInteger struct {
	typ *Type
}

func TypeGeneratorIntegerNew(typ *Type) *TypeGeneratorInteger {
	return &TypeGeneratorInteger{
		typ: typ,
	}
}

func (t *TypeGeneratorInteger) isSupportedAsParam(direction string) (supported bool, reason string) {
	if t.typ.indirectLevel > 0 {
		return false, fmt.Sprintf("%s with indirection level of %d",
			t.typ.CType, t.typ.indirectLevel)
	}

	return true, ""
}

func (t *TypeGeneratorInteger) isSupportedAsReturnValue() (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorInteger) generateDeclaration(g *jen.Group, goVarName string) {
	g.
		Id(goVarName).
		Id(t.typ.goType)
}

func (t *TypeGeneratorInteger) generateParamCallArgument(g *jen.Group, cVarName string) {
	g.Id(cVarName)
}

func (t *TypeGeneratorInteger) generateParamOutCallArgument(g *jen.Group, cVarName string) {
	panic(fmt.Sprintf("call argument for an integer out param, not supported : %s", cVarName))
}

func (t *TypeGeneratorInteger) generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string) {
	g.
		Id(cVarName).
		Op(":=").
		Parens(jen.Qual("C", t.typ.CType)).
		Parens(jen.Id(goVarName))
}

func (t *TypeGeneratorInteger) generateParamOutCVar(g *jen.Group, cVarName string) {
	g.
		Var().
		Id(cVarName).
		Qual("C", t.typ.CType)
}

func (t *TypeGeneratorInteger) generateReturnFunctionDeclaration(g *jen.Group) {
	g.Id(t.typ.goType)
}

func (t *TypeGeneratorInteger) generateReturnCToGo(g *jen.Group, cVarName string, transferOwnership string) {
	g.
		Parens(jen.Id(t.typ.goType)).
		Parens(jen.Id(cVarName))
}
