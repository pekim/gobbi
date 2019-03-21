package generate

import (
	"github.com/dave/jennifer/jen"
)

type TypeGeneratorGpointer struct {
	typ *Type
}

func TypeGeneratorGpointerNew(typ *Type) *TypeGeneratorGpointer {
	return &TypeGeneratorGpointer{
		typ: typ,
	}
}

func (t *TypeGeneratorGpointer) isSupportedAsParam(direction string) (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorGpointer) isSupportedAsArrayParam(direction string) (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorGpointer) isSupportedAsArrayParamC(direction string) (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorGpointer) isSupportedAsParamC() (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorGpointer) isSupportedAsField() (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorGpointer) isSupportedAsArrayReturnValue() (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorGpointer) isSupportedAsReturnValue() (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorGpointer) isSupportedAsReturnCValue() (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorGpointer) generateDeclaration(g *jen.Group, goVarName string) {
	panic("unsupported")
}

func (t *TypeGeneratorGpointer) generateArrayDeclaration(g *jen.Group, goVarName string) {
	panic("unsupported")
}

func (t *TypeGeneratorGpointer) generateArrayDeclarationC(g *jen.Group, cVarName string) {
	panic("unsupported")
}

func (t *TypeGeneratorGpointer) generateDeclarationC(g *jen.Group, cVarName string) {
	g.Id(cVarName).Qual("C", "gpointer")
}

func (t *TypeGeneratorGpointer) generateParamCallArgument(g *jen.Group, cVarName string) {
	panic("unsupported")
}

func (t *TypeGeneratorGpointer) generateParamOutCallArgument(g *jen.Group, cVarName string) {
	panic("unsupported")
}

func (t *TypeGeneratorGpointer) generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string) {
	panic("unsupported")
}

func (t *TypeGeneratorGpointer) generateParamGoVar(g *jen.Group, goVarName string, cVarName string, pkg string) {
	panic("unsupported")
}

func (t *TypeGeneratorGpointer) generateParamOutCVar(g *jen.Group, cVarName string) {
	panic("unsupported")
}

func (t *TypeGeneratorGpointer) generateReturnFunctionDeclaration(g *jen.Group) {
	panic("unsupported")
}

func (t *TypeGeneratorGpointer) generateReturnFunctionDeclarationCtype(g *jen.Group) {
	panic("unsupported")
}

func (t *TypeGeneratorGpointer) generateReturnCToGo(g *jen.Group, isParam bool,
	cVarName string, goVarName string, pkg string, transferOwnership string, nullable bool) {
}

func (t *TypeGeneratorGpointer) generateArrayReturnCToGo(g *jen.Group, isParam bool,
	cVarName string, goVarName string, pkg string, transferOwnership string, nullable bool) {
	panic("unsupported")
}

func (t *TypeGeneratorGpointer) generateCToGo(pkg string, cVarReference *jen.Statement) *jen.Statement {
	panic("unsupported")
}

func (t *TypeGeneratorGpointer) generateGoToC(g *jen.Group, goVarReference *jen.Statement) {
	panic("unsupported")
}
