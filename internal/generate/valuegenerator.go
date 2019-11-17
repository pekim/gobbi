package generate

import "github.com/dave/jennifer/jen"

// argValuegenerator can generate a value from a *gi.Argument.
type argValuegenerator interface {
	canGenerateValue() bool
	generateArgValue(arg *jen.Statement)
}
