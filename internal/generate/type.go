package generate

import (
	"github.com/dave/jennifer/jen"
	"strings"
)

type Type struct {
	Namespace *Namespace

	Name  string `xml:"name,attr"`
	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	generator     TypeGenerator
	goType        string
	cTypeName     string // the C type, such as "gchar"
	indirectLevel int    // the level of pointer indirection
}

func (t *Type) init(ns *Namespace) {
	t.Namespace = ns

	cTypeParts := strings.Split(t.CType, " ")
	cType := cTypeParts[len(cTypeParts)-1]

	t.cTypeName = strings.TrimRight(cType, "*")
	t.indirectLevel = len(cType) - len(t.cTypeName)

	// set goType and a TypeGenerator
	goType, isInteger := integerCTypeMap[t.CType]
	if isInteger {
		t.goType = goType
		t.generator = TypeIntegerNew(t)
	}

	if t.Name == "utf8" || t.Name == "filename" {
		t.goType = "string"
		t.generator = TypeStringNew(t)
	}
}

// TypeGenerator is an interface where implementors provide generator
// support for a type or types.
type TypeGenerator interface {
	isSupportedAsParam(direction string) (supported bool, reason string)
	isSupportedAsReturnValue() (supported bool, reason string)

	generateParamFunctionDeclaration(g *jen.Group, goVarName string)
	generateParamCallArgument(g *jen.Group, cVarName string)
	generateParamOutCallArgument(g *jen.Group, cVarName string)
	generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string)
	generateParamOutCVar(g *jen.Group, cVarName string)

	generateReturnFunctionDeclaration(g *jen.Group)
	generateReturnCToGo(g *jen.Group, cVarName string, transferOwnership string)
}
