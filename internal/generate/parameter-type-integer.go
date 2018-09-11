package generate

import (
	"github.com/dave/jennifer/jen"
)

type ParameterTypeInteger struct {
	param *Parameter
}

func ParameterTypeIntegerNew(param *Parameter) *ParameterTypeInteger {
	return &ParameterTypeInteger{param: param}
}

func (pt *ParameterTypeInteger) generateFunctionDeclaration(g *jen.Group) {
	g.
		Id(pt.param.goVarName).
		Id(pt.param.goType)
}

func (pt *ParameterTypeInteger) generateCallArgument(g *jen.Group) {
	g.Id(pt.param.cVarName)
}

func (pt *ParameterTypeInteger) generateCVar(g *jen.Group) {
	g.
		Id(pt.param.cVarName).
		Op(":=").
		Parens(jen.Qual("C", pt.param.Type.CType)).
		Parens(jen.Id(pt.param.goVarName))
}

func (pt *ParameterTypeInteger) generateOutCVar(g *jen.Group) {
	g.
		Var().
		Id(pt.param.cVarName).
		Qual("C", pt.param.Type.CType)
}
