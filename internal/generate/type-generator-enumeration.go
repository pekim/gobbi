package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type TypeGeneratorEnumeration struct {
	TypeGeneratorPanic
	typ  *Type
	enum *Enumeration
}

func TypeGeneratorEnumerationNew(typ *Type, enum *Enumeration) *TypeGeneratorEnumeration {
	return &TypeGeneratorEnumeration{
		typ:  typ,
		enum: enum,
	}
}

func (t *TypeGeneratorEnumeration) isSupportedAsParam(direction string) (supported bool, reason string) {
	if t.typ.indirectLevel > 0 {
		return false, fmt.Sprintf("%s with indirection level of %d",
			t.typ.CType, t.typ.indirectLevel)
	}

	return true, ""
}

func (t *TypeGeneratorEnumeration) isSupportedAsArrayParam(direction string) (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorEnumeration) isSupportedAsArrayParamC(direction string) (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorEnumeration) isSupportedAsParamC() (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorEnumeration) isSupportedAsField() (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorEnumeration) isSupportedAsArrayReturnValue() (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorEnumeration) isSupportedAsReturnValue() (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorEnumeration) isSupportedAsReturnCValue() (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorEnumeration) generateDeclaration(g *jen.Group, goVarName string) {
	g.
		Id(goVarName).
		Do(t.typ.qname.generate)
}

func (t *TypeGeneratorEnumeration) generateArrayDeclaration(g *jen.Group, goVarName string) {
	g.
		Id(goVarName).
		Index().
		Do(t.typ.qname.generate)
}

func (t *TypeGeneratorEnumeration) generateArrayDeclarationC(g *jen.Group, cVarName string) {
}

func (t *TypeGeneratorEnumeration) generateDeclarationC(g *jen.Group, cVarName string) {
	g.Id(cVarName).Qual("C", t.enum.CType)
}

func (t *TypeGeneratorEnumeration) generateParamCallArgument(g *jen.Group, cVarName string) {
	g.Id(cVarName)
}

func (t *TypeGeneratorEnumeration) generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string) {
	g.
		Id(cVarName).
		Op(":=").
		Parens(jen.Qual("C", t.typ.CType)).
		Parens(jen.Id(goVarName))
}

func (t *TypeGeneratorEnumeration) generateParamGoVar(g *jen.Group, goVarName string, cVarName string, pkg string) {
	g.
		Id(goVarName).
		Op(":=").
		Do(t.typ.qname.generate).
		Parens(jen.Id(cVarName))
}

func (t *TypeGeneratorEnumeration) generateParamOutCVar(g *jen.Group, cVarName string) {
	g.
		Var().
		Id(cVarName).
		Qual("C", t.typ.CType)
}

func (t *TypeGeneratorEnumeration) generateReturnFunctionDeclaration(g *jen.Group) {
	g.Do(t.typ.qname.generate)
}

func (t *TypeGeneratorEnumeration) generateReturnFunctionDeclarationCtype(g *jen.Group) {
}

func (t *TypeGeneratorEnumeration) generateReturnCToGo(g *jen.Group, isParam bool,
	cVarName string, goVarName string, pkg string, transferOwnership string, nullable bool) {
	g.
		Id(goVarName).
		Op(":=").
		Parens(jen.Do(t.typ.qname.generate)).
		Parens(jen.Id(cVarName))
}

func (t *TypeGeneratorEnumeration) generateCToGo(pkg string, cVarReference *jen.Statement) *jen.Statement {
	return jen.
		Parens(jen.Do(t.typ.qname.generate)).
		Params(cVarReference)
}

func (t *TypeGeneratorEnumeration) generateGoToC(g *jen.Group, goVarReference *jen.Statement) {
	ctype := t.typ.cTypeName
	if ctype == "" {
		ctype = t.enum.CType
	}

	g.
		Parens(jen.Qual("C", ctype)).
		Parens(goVarReference)
}
