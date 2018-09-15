package generate

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

type TypeInteger struct {
	typ *Type
}

func TypeIntegerNew(typ *Type) *TypeInteger {
	return &TypeInteger{
		typ: typ,
	}
}

func (t *TypeInteger) isSupportedAsParam(direction string) (supported bool, reason string) {
	if t.typ.indirectLevel > 0 {
		return false, fmt.Sprintf("%s with indirection level of %d",
			t.typ.CType, t.typ.indirectLevel)
	}

	return true, ""
}

func (t *TypeInteger) isSupportedAsReturnValue() (supported bool, reason string) {
	return true, ""
}

func (t *TypeInteger) generateParamFunctionDeclaration(g *jen.Group, goVarName string) {
	g.
		Id(goVarName).
		Id(t.typ.goType)
}

func (t *TypeInteger) generateParamCallArgument(g *jen.Group, cVarName string) {
	g.Id(cVarName)
}

func (t *TypeInteger) generateParamOutCallArgument(g *jen.Group, cVarName string) {
	panic(fmt.Sprintf("call argument for an integer out param, not supported : %s", cVarName))
}

func (t *TypeInteger) generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string) {
	g.
		Id(cVarName).
		Op(":=").
		Parens(jen.Qual("C", t.typ.CType)).
		Parens(jen.Id(goVarName))
}

func (t *TypeInteger) generateParamOutCVar(g *jen.Group, cVarName string) {
	g.
		Var().
		Id(cVarName).
		Qual("C", t.typ.CType)
}

func (t *TypeInteger) generateReturnFunctionDeclaration(g *jen.Group) {
	g.Id(t.typ.goType)
}

func (t *TypeInteger) generateReturnCToGo(g *jen.Group, cVarName string, transferOwnership string) {
	g.
		Parens(jen.Id(t.typ.goType)).
		Parens(jen.Id(cVarName))
}
