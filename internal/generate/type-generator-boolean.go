package generate

import (
	"github.com/dave/jennifer/jen"
)

type TypeGeneratorBoolean struct {
	typ *Type
}

func TypeGeneratorBooleanNew(typ *Type) *TypeGeneratorBoolean {
	return &TypeGeneratorBoolean{
		typ: typ,
	}
}

func (t *TypeGeneratorBoolean) isSupportedAsParam(direction string) (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorBoolean) isSupportedAsField() (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorBoolean) isSupportedAsReturnValue() (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorBoolean) generateDeclaration(g *jen.Group, goVarName string) {
	g.
		Id(goVarName).
		Id("bool")
}

func (t *TypeGeneratorBoolean) generateParamCallArgument(g *jen.Group, cVarName string) {
	g.Id(cVarName)
}

func (t *TypeGeneratorBoolean) generateParamOutCallArgument(g *jen.Group, cVarName string) {
	g.
		Op("&").
		Id(cVarName)
}

func (t *TypeGeneratorBoolean) generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string) {
	g.Id(cVarName).Op(":=")
	t.generateCallBoolToGboolean(g, jen.Id(goVarName))
}

func (t *TypeGeneratorBoolean) generateParamOutCVar(g *jen.Group, cVarName string) {
}

func (t *TypeGeneratorBoolean) generateReturnFunctionDeclaration(g *jen.Group) {
	g.Id("bool")
}

func (t *TypeGeneratorBoolean) generateReturnCToGo(g *jen.Group,
	cVarName string, goVarName string, pkg string, transferOwnership string) {

	g.
		Id(goVarName).
		Op(":=").
		Id(cVarName).
		Op("==").
		Qual("C", "TRUE")
}

func (t *TypeGeneratorBoolean) generateCToGo(pkg string, cVarReference *jen.Statement) *jen.Statement {
	return cVarReference.
		Op("==").
		Qual("C", "TRUE")
}

func (t *TypeGeneratorBoolean) generateGoToC(g *jen.Group, goVarReference *jen.Statement) {
	t.generateCallBoolToGboolean(g, goVarReference)
}

func (t *TypeGeneratorBoolean) generateCallBoolToGboolean(g *jen.Group, goVarReference *jen.Statement) {
	g.
		Id("boolToGboolean").
		Call(goVarReference)
}