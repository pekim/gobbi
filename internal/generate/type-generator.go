package generate

import (
	"github.com/dave/jennifer/jen"
)

// TypeGenerator is an interface where implementors provide generator
// support for a type or types.
type TypeGenerator interface {
	isSupportedAsParam(direction string) (supported bool, reason string)
	isSupportedAsReturnValue() (supported bool, reason string)

	generateParamFunctionDeclaration(g *jen.Group, goVarName string)
	generateParamCallArgument(g *jen.Group, cVarName string)
	generateParamOutCallArgument(g *jen.Group, cVarName string)
	generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string)
	generateParamOutCVar(g *jen.Group, cVarName string)

	generateReturnFunctionDeclaration(g *jen.Group)
	generateReturnCToGo(g *jen.Group, cVarName string, transferOwnership string)
}
