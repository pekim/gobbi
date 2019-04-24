package generate

import (
	"github.com/dave/jennifer/jen"
)

// TypeGenerator is an interface where implementors provide generator
// support for a type or types.
type TypeGenerator interface {
	isSupportedAsField() (supported bool, reason string)
	isSupportedAsParam(direction string) (supported bool, reason string)
	isSupportedAsParamC() (supported bool, reason string)
	isSupportedAsArrayParam(direction string) (supported bool, reason string)
	isSupportedAsArrayParamC(direction string) (supported bool, reason string)
	isSupportedAsArrayReturnValue() (supported bool, reason string)
	isSupportedAsReturnValue() (supported bool, reason string)
	isSupportedAsReturnCValue() (supported bool, reason string)

	generateDeclaration(g *jen.Group, goVarName string)
	generateDeclarationC(g *jen.Group, goVarName string)
	generateArrayDeclaration(g *jen.Group, goVarName string)
	generateArrayDeclarationC(g *jen.Group, cVarName string)
	generateCToGo(pkg string, cVarReference *jen.Statement) *jen.Statement
	generateGoToC(g *jen.Group, goVarReference *jen.Statement)

	generateCallReturnType() string

	generateParamCallArgument(g *jen.Group, cVarName string)
	generateParamOutCallArgument(g *jen.Group, cVarName string)
	generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string)
	generateParamOutCVar(g *jen.Group, cVarName string)
	generateParamGoVar(g *jen.Group, goVarName string, cVarName string, pkg string)

	generateReturnFunctionDeclaration(g *jen.Group)
	generateReturnFunctionDeclarationCtype(g *jen.Group)
	generateReturnCToGo(g *jen.Group, isParam bool, cVarName string, goVarName string, pkg string,
		transferOwnership string, nullable bool)
	generateArrayReturnCToGo(g *jen.Group, isParam bool, cVarName string, goVarName string, pkg string,
		transferOwnership string, nullable bool)
}

type TypeGeneratorPanic struct{}

func (p TypeGeneratorPanic) isSupportedAsField() (supported bool, reason string) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) isSupportedAsParam(direction string) (supported bool, reason string) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) isSupportedAsParamC() (supported bool, reason string) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) isSupportedAsArrayParam(direction string) (supported bool, reason string) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) isSupportedAsArrayParamC(direction string) (supported bool, reason string) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) isSupportedAsArrayReturnValue() (supported bool, reason string) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) isSupportedAsReturnValue() (supported bool, reason string) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) isSupportedAsReturnCValue() (supported bool, reason string) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) generateDeclaration(g *jen.Group, goVarName string) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) generateDeclarationC(g *jen.Group, goVarName string) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) generateArrayDeclaration(g *jen.Group, goVarName string) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) generateArrayDeclarationC(g *jen.Group, cVarName string) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) generateCToGo(pkg string, cVarReference *jen.Statement) *jen.Statement {
	panic("not implmented")
}

func (p TypeGeneratorPanic) generateGoToC(g *jen.Group, goVarReference *jen.Statement) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) generateCallReturnType() string {
	panic("not implmented")
}

func (p TypeGeneratorPanic) generateParamCallArgument(g *jen.Group, cVarName string) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) generateParamOutCallArgument(g *jen.Group, cVarName string) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) generateParamOutCVar(g *jen.Group, cVarName string) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) generateParamGoVar(g *jen.Group, goVarName string, cVarName string, pkg string) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) generateReturnFunctionDeclaration(g *jen.Group) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) generateReturnFunctionDeclarationCtype(g *jen.Group) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) generateReturnCToGo(g *jen.Group, isParam bool, cVarName string, goVarName string, pkg string, transferOwnership string, nullable bool) {
	panic("not implmented")
}

func (p TypeGeneratorPanic) generateArrayReturnCToGo(g *jen.Group, isParam bool, cVarName string, goVarName string, pkg string, transferOwnership string, nullable bool) {
	panic("not implmented")
}
