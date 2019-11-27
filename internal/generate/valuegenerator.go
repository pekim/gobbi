package generate

import "github.com/dave/jennifer/jen"

type outParameterGenerator interface {
	supportedAsOutParameter() bool
	generateDeclaration() *jen.Statement
	argumentGetFunctionName() string
	argumentSetFunctionName() string

	createFromArgument(s *jen.Statement, arg *jen.Statement)
}
