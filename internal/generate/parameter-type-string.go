package generate

import "github.com/dave/jennifer/jen"

type ParameterTypeString struct {
	param *Parameter
}

func ParameterTypeStringNew(param *Parameter) *ParameterTypeString {
	return &ParameterTypeString{param: param}
}

func (pt *ParameterTypeString) generateFunctionDeclaration(g *jen.Group) {
	g.
		Id(pt.param.goVarName).
		Id(pt.param.goType)
}

func (pt *ParameterTypeString) generateCallArgument(g *jen.Group) {
	g.Id(pt.param.cVarName)
}

func (pt *ParameterTypeString) generateCVar(g *jen.Group) {
	g.
		Id(pt.param.cVarName).
		Op(":=").
		Qual("C", "CString").
		Call(jen.Id(pt.param.goVarName))

	g.
		Defer().
		Qual("C", "free").
		Call(jen.
			Qual("unsafe", "Pointer").
			Call(jen.Id(pt.param.cVarName)))
}

func (pt *ParameterTypeString) generateOutCVar(g *jen.Group) {
	//g.
	//	Var().
	//	Id(pt.param.cVarName).
	//	Qual("C", pt.param.Type.CType)
}
