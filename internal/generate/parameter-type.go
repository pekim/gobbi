package generate

import "github.com/dave/jennifer/jen"

type ParameterType interface {
	isSupported() (supported bool, reason string)

	generateFunctionDeclaration(g *jen.Group)
	generateCallArgument(g *jen.Group)
	generateOutCallArgument(g *jen.Group)
	generateCVar(g *jen.Group)
	generateOutCVar(g *jen.Group)
}
