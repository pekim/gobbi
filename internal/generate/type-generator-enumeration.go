package generate

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

type TypeGeneratorEnumeration struct {
	typ *Type
}

func TypeGeneratorEnumerationNew(typ *Type) *TypeGeneratorEnumeration {
	return &TypeGeneratorEnumeration{
		typ: typ,
	}
}

func (t *TypeGeneratorEnumeration) isSupportedAsParam(direction string) (supported bool, reason string) {
	if t.typ.indirectLevel > 0 {
		return false, fmt.Sprintf("%s with indirection level of %d",
			t.typ.CType, t.typ.indirectLevel)
	}

	return true, ""
}

func (t *TypeGeneratorEnumeration) isSupportedAsReturnValue() (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorEnumeration) generateDeclaration(g *jen.Group, goVarName string) {
	g.
		Id(goVarName).
		Id(t.typ.goType)
}

func (t *TypeGeneratorEnumeration) generateParamCallArgument(g *jen.Group, cVarName string) {
	g.Id(cVarName)
}

func (t *TypeGeneratorEnumeration) generateParamOutCallArgument(g *jen.Group, cVarName string) {
	panic(fmt.Sprintf("call argument for an enum out param, not supported : %s", cVarName))
}

func (t *TypeGeneratorEnumeration) generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string) {
	g.
		Id(cVarName).
		Op(":=").
		Parens(jen.Qual("C", t.typ.CType)).
		Parens(jen.Id(goVarName))
}

func (t *TypeGeneratorEnumeration) generateParamOutCVar(g *jen.Group, cVarName string) {
	g.
		Var().
		Id(cVarName).
		Qual("C", t.typ.CType)
}

func (t *TypeGeneratorEnumeration) generateReturnFunctionDeclaration(g *jen.Group) {
	g.Id(t.typ.goType)
}

func (t *TypeGeneratorEnumeration) generateReturnCToGo(g *jen.Group, cVarName string, goVarName string, transferOwnership string) {
	g.
		Id(goVarName).
		Op(":=").
		Parens(jen.Id(t.typ.goType)).
		Parens(jen.Id(cVarName))
}

func (t *TypeGeneratorEnumeration) generateCToGo(cVarReference *jen.Statement) *jen.Statement {
	return jen.
		Parens(jen.Id(t.typ.goType)).
		Params(cVarReference)
}
