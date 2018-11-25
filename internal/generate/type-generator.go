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
	isSupportedAsReturnValue() (supported bool, reason string)
	isSupportedAsReturnCValue() (supported bool, reason string)

	generateDeclaration(g *jen.Group, goVarName string)
	generateDeclarationC(g *jen.Group, goVarName string)
	generateArrayDeclaration(g *jen.Group, goVarName string)
	generateCToGo(pkg string, cVarReference *jen.Statement) *jen.Statement
	generateGoToC(g *jen.Group, goVarReference *jen.Statement)

	generateParamCallArgument(g *jen.Group, cVarName string)
	generateParamOutCallArgument(g *jen.Group, cVarName string)
	generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string)
	generateParamOutCVar(g *jen.Group, cVarName string)
	generateParamGoVar(g *jen.Group, goVarName string, cVarName string, pkg string)

	generateReturnFunctionDeclaration(g *jen.Group)
	generateReturnFunctionDeclarationCtype(g *jen.Group)
	generateReturnCToGo(g *jen.Group, isParam bool, cVarName string, goVarName string, pkg string,
		transferOwnership string, nullable bool)
}
