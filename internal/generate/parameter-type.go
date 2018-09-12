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

// parameterType gets the Go type and a ParameterType for a Parameter.
func parameterType(param *Parameter) (string, ParameterType) {
	goType, isInteger := integerCTypeMap[param.Type.CType]
	if isInteger {
		return goType, ParameterTypeIntegerNew(param)
	}

	if param.Type.Name == "utf8" || param.Type.Name == "filename" {
		return "string", ParameterTypeStringNew(param)
	}

	return "", nil
}
