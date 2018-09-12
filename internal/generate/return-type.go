package generate

import (
	"github.com/dave/jennifer/jen"
)

type ReturnType interface {
	isSupported() (supported bool, reason string)

	generateFunctionDeclaration(g *jen.Group)
	generateCToGo(g *jen.Group, cVarName string)
}

// returnType gets the Go type and a ReturnType for a ReturnValue.
func returnType(rv *ReturnValue) (string, ReturnType) {
	goType, isInteger := integerCTypeMap[rv.Type.CType]
	if isInteger {
		return goType, ReturnTypeIntegerNew(rv)
	}

	if rv.Type.Name == "utf8" {
		return "string", ReturnTypeStringNew(rv)
	}

	return "", nil
}
