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
	if t.typ.indirectLevel != 1 {
		return false, fmt.Sprintf("record with indirection of %d", t.typ.indirectLevel)
	}

	return true, ""
}

func (t *TypeGeneratorRecord) isSupportedAsParam(direction string) (supported bool, reason string) {
	if t.record.Blacklist {
		return false, fmt.Sprintf("Blacklisted record : %s", t.record.CType)
	}

	//if direction != "out" && t.typ.indirectLevel > 1 {
	//	return false, fmt.Sprintf("in string with indirection level of %d",
	//		t.typ.indirectLevel)
	//}

	return false, "record param - coming soon"
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

func (t *TypeGeneratorRecord) generateDeclaration(g *jen.Group, goVarName string) {
	g.
		Id(goVarName).
		Op("*").
		Id(t.typ.goType)
}

func (t *TypeGeneratorRecord) generateParamCallArgument(g *jen.Group, cVarName string) {
}

func (t *TypeGeneratorRecord) generateParamOutCallArgument(g *jen.Group, cVarName string) {
	g.
		Op("&").
		Id(cVarName)
}

func (t *TypeGeneratorRecord) generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string) {
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
		Id(t.typ.goType)
}

func (t *TypeGeneratorRecord) generateReturnCToGo(g *jen.Group, cVarName string, goVarName string,
	transferOwnership string) {
	g.
		Id(goVarName).
		Op(":=").
		Id(t.record.newFromCFuncName).
		Call(jen.Id(cVarName))
}

func (t *TypeGeneratorRecord) generateCToGo(cVarReference *jen.Statement) *jen.Statement {
	return jen.
		Id(t.record.newFromCFuncName).
		Call(cVarReference)
}
