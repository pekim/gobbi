package generate

import "github.com/dave/jennifer/jen"

type ReturnTypeInteger struct {
	returnValue *ReturnValue
}

func ReturnTypeIntegerNew(rv *ReturnValue) *ReturnTypeInteger {
	return &ReturnTypeInteger{returnValue: rv}
}

func (rt *ReturnTypeInteger) isSupported() (supported bool, reason string) {
	return true, ""
}

func (rt *ReturnTypeInteger) generateFunctionDeclaration(g *jen.Group) {
	g.Id(rt.returnValue.goType)
}

func (rt *ReturnTypeInteger) generateCToGo(g *jen.Group, cVarName string) {
	g.
		Parens(jen.Id(rt.returnValue.goType)).
		Parens(jen.Id(cVarName))
}
