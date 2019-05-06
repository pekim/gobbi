package generate

import "github.com/dave/jennifer/jen"

type TypeGeneratorPanic struct {
	typ *Type
}

func (t *TypeGeneratorPanic) isSupportedAsField() (supported bool, reason string) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) isSupportedAsParam(direction string) (supported bool, reason string) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) isSupportedAsParamC() (supported bool, reason string) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) isSupportedAsArrayParam(direction string) (supported bool, reason string) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) isSupportedAsArrayParamC(direction string) (supported bool, reason string) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) isSupportedAsArrayReturnValue() (supported bool, reason string) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) isSupportedAsReturnValue() (supported bool, reason string) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) isSupportedAsReturnCValue() (supported bool, reason string) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) generateDeclaration(g *jen.Group, goVarName string) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) generateDeclarationC(g *jen.Group, goVarName string) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) generateArrayDeclaration(g *jen.Group, goVarName string) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) generateArrayDeclarationC(g *jen.Group, cVarName string) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) generateCToGo(pkg string, cVarReference *jen.Statement) *jen.Statement {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) generateGoToC(g *jen.Group, goVarReference *jen.Statement) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) generateParamCallArgument(g *jen.Group, cVarName string) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) generateParamOutCallArgument(g *jen.Group, cVarName string) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) generateParamOutCVar(g *jen.Group, cVarName string) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) generateParamGoVar(g *jen.Group, goVarName string, cVarName string, pkg string) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) generateReturnFunctionDeclaration(g *jen.Group) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) generateReturnFunctionDeclarationCtype(g *jen.Group) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) generateReturnCToGo(g *jen.Group, isParam bool, cVarName string, goVarName string, pkg string, transferOwnership string, nullable bool) {
	panic("not supported by type generator")
}

func (t *TypeGeneratorPanic) generateArrayReturnCToGo(g *jen.Group, isParam bool, cVarName string, goVarName string, pkg string, transferOwnership string, nullable bool) {
	panic("not supported by type generator")
}
