package generate

import "github.com/dave/jennifer/jen"

type ParameterType interface {
	generateFunctionDeclaration(g *jen.Group)
	generateCallArgument(g *jen.Group)
	generateCVar(g *jen.Group)
	generateOutCVar(g *jen.Group)
}
