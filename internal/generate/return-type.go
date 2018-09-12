package generate

import "github.com/dave/jennifer/jen"

type ReturnType interface {
	isSupported() (supported bool, reason string)

	generateFunctionDeclaration(g *jen.Group)
	generateCToGo(g *jen.Group, cVarName string)
}
