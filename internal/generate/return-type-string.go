package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

type ReturnTypeString struct {
	returnValue   *ReturnValue
	cTypeName     string
	indirectLevel int
}

func ReturnTypeStringNew(rv *ReturnValue) *ReturnTypeString {
	cTypeParts := strings.Split(rv.Type.CType, " ")
	cType := cTypeParts[len(cTypeParts)-1]
	cTypeName := strings.TrimRight(cType, "*")
	indirectLevel := len(cType) - len(cTypeName)

	return &ReturnTypeString{
		returnValue:   rv,
		cTypeName:     cTypeName,
		indirectLevel: indirectLevel,
	}
}

func (rt *ReturnTypeString) isSupported() (supported bool, reason string) {
	if rt.indirectLevel > 1 {
		return false, fmt.Sprintf("return for string with indirection level of %d", rt.indirectLevel)
	}

	return true, ""
}

func (rt *ReturnTypeString) generateFunctionDeclaration(g *jen.Group) {
	g.Id(rt.returnValue.goType)
}

func (rt *ReturnTypeString) generateCToGo(g *jen.Group, cVarName string) {
	g.
		Qual("C", "GoString").
		Call(jen.Id(cVarName))

	if rt.returnValue.TransferOwnership != "none" {
		// the library will be responsible for the memory
		return
	}

	g.
		Defer().
		Qual("C", "free").
		Call(jen.
			Qual("unsafe", "Pointer").
			Call(jen.Id(cVarName)))
}
