package generate

import (
	"github.com/dave/jennifer/jen"
)

type TypeGeneratorArgcArgv struct {
	typ *Type
}

func TypeGeneratorArgcArgvNew(typ *Type) *TypeGeneratorArgcArgv {
	return &TypeGeneratorArgcArgv{
		typ: typ,
	}
}

func (t *TypeGeneratorArgcArgv) isSupportedAsParam(direction string) (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorArgcArgv) isSupportedAsField() (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorArgcArgv) isSupportedAsReturnValue() (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorArgcArgv) generateDeclaration(g *jen.Group, goVarName string) {
	g.
		Id(goVarName).
		Op("[]").
		String()
}

func (t *TypeGeneratorArgcArgv) generateParamCallArgument(g *jen.Group, cVarName string) {
	g.
		Op("&").
		Id("cArgc")

	//g.Id(cVarName)
}

func (t *TypeGeneratorArgcArgv) generateParamOutCallArgument(g *jen.Group, cVarName string) {
}

func (t *TypeGeneratorArgcArgv) generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string) {
	g.
		Var().
		Id("cArgc").
		Qual("C", "gint").
		Op("=").
		Len(jen.Id(goVarName))
}

func (t *TypeGeneratorArgcArgv) generateParamOutCVar(g *jen.Group, cVarName string) {
	g.
		Var().
		Id(cVarName).
		Qual("C", "gboolean")
}

func (t *TypeGeneratorArgcArgv) generateReturnFunctionDeclaration(g *jen.Group) {
	g.Op("[]").String()
}

func (t *TypeGeneratorArgcArgv) generateReturnCToGo(g *jen.Group, isParam bool,
	cVarName string, goVarName string, pkg string, transferOwnership string) {

	g.
		Id(goVarName).
		Op(":=").
		Id(cVarName).
		Op("==").
		Qual("C", "TRUE")
}

func (t *TypeGeneratorArgcArgv) generateCToGo(pkg string, cVarReference *jen.Statement) *jen.Statement {
	return cVarReference.
		Op("==").
		Qual("C", "TRUE")
}

func (t *TypeGeneratorArgcArgv) generateGoToC(g *jen.Group, goVarReference *jen.Statement) {
	t.generateCallBoolToGboolean(g, goVarReference)
}

func (t *TypeGeneratorArgcArgv) generateCallBoolToGboolean(g *jen.Group, goVarReference *jen.Statement) {
	g.
		Id("boolToGboolean").
		Call(goVarReference)
}
