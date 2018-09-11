package generate

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

type ParameterTypeString struct {
	param         *Parameter
	cTypeName     string
	indirectLevel int
}

func ParameterTypeStringNew(param *Parameter) *ParameterTypeString {
	cTypeParts := strings.Split(param.Type.CType, " ")
	cType := cTypeParts[len(cTypeParts)-1]
	cTypeName := strings.TrimRight(cType, "*")
	indirectLevel := len(cType) - len(cTypeName)

	return &ParameterTypeString{
		param:         param,
		cTypeName:     cTypeName,
		indirectLevel: indirectLevel,
	}
}

func (pt *ParameterTypeString) isSupported() (supported bool, reason string) {
	if pt.param.Direction != "out" && pt.indirectLevel > 1 {
		return false, fmt.Sprintf("in for string with indirection level of %d", pt.indirectLevel)
	}

	return true, ""
}

func (pt *ParameterTypeString) generateFunctionDeclaration(g *jen.Group) {
	g.
		Id(pt.param.goVarName).
		Id(pt.param.goType)
}

func (pt *ParameterTypeString) generateCallArgument(g *jen.Group) {
	g.Id(pt.param.cVarName)
}

func (pt *ParameterTypeString) generateOutCallArgument(g *jen.Group) {
	g.
		Op("&").
		Id(pt.param.cVarName)
}

func (pt *ParameterTypeString) generateCVar(g *jen.Group) {
	g.
		Id(pt.param.cVarName).
		Op(":=").
		Qual("C", "CString").
		Call(jen.Id(pt.param.goVarName))

	g.
		Defer().
		Qual("C", "free").
		Call(jen.
			Qual("unsafe", "Pointer").
			Call(jen.Id(pt.param.cVarName)))
}

func (pt *ParameterTypeString) generateOutCVar(g *jen.Group) {
	g.
		Var().
		Id(pt.param.cVarName).
		Op(strings.Repeat("*", pt.indirectLevel-1)).
		Qual("C", pt.cTypeName)
}
