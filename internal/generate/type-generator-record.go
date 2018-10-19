package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

type TypeGeneratorRecord struct {
	typ    *Type
	record *Record
}

func TypeGeneratorRecordNew(typ *Type, record *Record) *TypeGeneratorRecord {
	return &TypeGeneratorRecord{
		typ:    typ,
		record: record,
	}
}

func (t *TypeGeneratorRecord) isSupportedAsField() (supported bool, reason string) {
	return false, "record"

	//if t.typ.indirectLevel != 1 {
	//	return false, fmt.Sprintf("record with indirection of %d", t.typ.indirectLevel)
	//}
	//
	//return true, ""
}

func (t *TypeGeneratorRecord) isSupportedAsParam(direction string) (supported bool, reason string) {
	if t.record.Blacklist {
		return false, fmt.Sprintf("Blacklisted record : %s", t.record.CType)
	}

	if t.typ.indirectLevel > 1 {
		return false, fmt.Sprintf("record with indirection level of %d",
			t.typ.indirectLevel)
	}

	return true, ""
}

func (t *TypeGeneratorRecord) isSupportedAsParamC() (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorRecord) isSupportedAsReturnValue() (supported bool, reason string) {
	if t.record.Blacklist {
		return false, fmt.Sprintf("Blacklisted record : %s", t.record.CType)
	}

	if t.typ.indirectLevel > 1 {
		return false, fmt.Sprintf("record with indirection level of %d",
			t.typ.indirectLevel)
	}

	return true, ""
}

func (t *TypeGeneratorRecord) isSupportedAsReturnCValue() (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorRecord) generateDeclaration(g *jen.Group, goVarName string) {
	g.
		Id(goVarName).
		Op("*").
		Do(t.typ.qname.generate)
}

func (t *TypeGeneratorRecord) generateDeclarationC(g *jen.Group, goVarName string) {
}

func (t *TypeGeneratorRecord) generateParamCallArgument(g *jen.Group, cVarName string) {
	g.Id(cVarName)
}

func (t *TypeGeneratorRecord) generateParamOutCallArgument(g *jen.Group, cVarName string) {
	g.
		Op("&").
		Id(cVarName)
}

func (t *TypeGeneratorRecord) generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string) {
	g.
		Id(cVarName).
		Op(":=").
		Parens(jen.
			Op(strings.Repeat("*", t.typ.indirectLevel)).
			Qual("C", t.typ.cTypeName)).
		Parens(jen.
			Id(goVarName).
			Op(".").
			Id("ToC").
			Call())
}

func (t *TypeGeneratorRecord) generateParamOutCVar(g *jen.Group, cVarName string) {
	g.
		Var().
		Id(cVarName).
		Op(strings.Repeat("*", t.typ.indirectLevel-1)).
		Qual("C", t.typ.cTypeName)
}

func (t *TypeGeneratorRecord) generateReturnFunctionDeclaration(g *jen.Group) {
	g.
		Op(strings.Repeat("*", t.typ.indirectLevel)).
		Do(t.typ.qname.generate)
}

func (t *TypeGeneratorRecord) generateReturnFunctionDeclarationCtype(g *jen.Group) {
	g.
		Op(strings.Repeat("*", t.typ.indirectLevel+1)).
		Qual("C", t.record.CType)

}

func (t *TypeGeneratorRecord) generateReturnCToGo(g *jen.Group, isParam bool,
	cVarName string, goVarName string, pkg string,
	transferOwnership string) {

	cVarRef := jen.Id(cVarName)
	if isParam && t.typ.indirectLevel == 1 {
		cVarRef = jen.Op("&").Id(cVarName)
	}

	g.
		Id(goVarName).
		Op(":=").
		Do(func(s *jen.Statement) {
			if pkg != "" {
				s.Qual(pkg, t.record.newFromCFuncName)
			} else {
				s.Id(t.record.newFromCFuncName)
			}
		}).
		Call(jen.
			Qual("unsafe", "Pointer").
			Call(cVarRef))
}

func (t *TypeGeneratorRecord) generateCToGo(pkg string, cVarReference *jen.Statement) *jen.Statement {
	return jen.
		Do(func(s *jen.Statement) {
			if pkg != "" {
				s.Qual(pkg, t.record.newFromCFuncName)
			} else {
				s.Id(t.record.newFromCFuncName)
			}
		}).
		Call(cVarReference)
}

func (t *TypeGeneratorRecord) generateGoToC(g *jen.Group, goVarReference *jen.Statement) {
}
