package generate

import "github.com/dave/jennifer/jen"

// argValuegenerator can generate a value from a *gi.Argument.
type argValuegenerator interface {
	canGenerateValue() bool
	generateArgValue(arg *jen.Statement)
}

type outParameterGenerator interface {
	supportedAsOutParameter() bool
	generateDeclaration() *jen.Statement
	argumentGetFunctionName() string
	argumentSetFunctionName() string

	createFromArgument(g *jen.Group, arg *jen.Statement)
}
