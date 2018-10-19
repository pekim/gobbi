package generate

import (
	"github.com/dave/jennifer/jen"
)

type TypeGeneratorIgnore struct {
}

func TypeGeneratorIgnoreNew(typ *Type) *TypeGeneratorIgnore {
	return &TypeGeneratorIgnore{}
}

func (t *TypeGeneratorIgnore) isSupportedAsParam(direction string) (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorIgnore) isSupportedAsParamC() (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorIgnore) isSupportedAsField() (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorIgnore) isSupportedAsReturnValue() (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorIgnore) isSupportedAsReturnCValue() (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorIgnore) generateDeclaration(g *jen.Group, goVarName string) {
}

func (t *TypeGeneratorIgnore) generateDeclarationC(g *jen.Group, goVarName string) {
}

func (t *TypeGeneratorIgnore) generateParamCallArgument(g *jen.Group, cVarName string) {
}

func (t *TypeGeneratorIgnore) generateParamOutCallArgument(g *jen.Group, cVarName string) {
}

func (t *TypeGeneratorIgnore) generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string) {
}

func (t *TypeGeneratorIgnore) generateParamGoVar(g *jen.Group, goVarName string, cVarName string, transferOwnership string) {
}

func (t *TypeGeneratorIgnore) generateParamOutCVar(g *jen.Group, cVarName string) {
}

func (t *TypeGeneratorIgnore) generateReturnFunctionDeclaration(g *jen.Group) {
}

func (t *TypeGeneratorIgnore) generateReturnFunctionDeclarationCtype(g *jen.Group) {
}

func (t *TypeGeneratorIgnore) generateReturnCToGo(g *jen.Group, isParam bool,
	cVarName string, goVarName string, pkg string, transferOwnership string) {
}

func (t *TypeGeneratorIgnore) generateCToGo(pkg string, cVarReference *jen.Statement) *jen.Statement {
	return jen.Do(func(s *jen.Statement) {})
}

func (t *TypeGeneratorIgnore) generateGoToC(g *jen.Group, goVarReference *jen.Statement) {
}

func (t *TypeGeneratorIgnore) generateCallBoolToGboolean(g *jen.Group, goVarReference *jen.Statement) {
}
