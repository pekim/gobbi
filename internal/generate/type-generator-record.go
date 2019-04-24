package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

type TypeGeneratorRecord struct {
	TypeGeneratorPanic
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

	if t.typ.indirectLevel > 2 {
		return false, fmt.Sprintf("record with indirection level of %d",
			t.typ.indirectLevel)
	}

	return true, ""
}

func (t *TypeGeneratorRecord) isSupportedAsArrayParam(direction string) (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorRecord) isSupportedAsArrayParamC(direction string) (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorRecord) isSupportedAsParamC() (supported bool, reason string) {
	if t.record.Blacklist {
		return false, fmt.Sprintf("Blacklisted record : %s", t.record.CType)
	}

	return true, ""
}

func (t *TypeGeneratorRecord) isSupportedAsArrayReturnValue() (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorRecord) isSupportedAsReturnValue() (supported bool, reason string) {
	return false, ""
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
	if t.record.Blacklist {
		return false, fmt.Sprintf("Blacklisted record : %s", t.record.CType)
	}

	return true, ""
}

func (t *TypeGeneratorRecord) generateDeclaration(g *jen.Group, goVarName string) {
	g.
		Id(goVarName).
		Op("*").
		Do(t.typ.qname.generate)
}

func (t *TypeGeneratorRecord) generateArrayDeclaration(g *jen.Group, goVarName string) {
	g.
		Id(goVarName).
		Index().
		Op("*").
		Do(t.typ.qname.generate)
}

func (t *TypeGeneratorRecord) generateArrayDeclarationC(g *jen.Group, cVarName string) {
}

func (t *TypeGeneratorRecord) generateDeclarationC(g *jen.Group, cVarName string) {
	g.
		Id(cVarName).
		Op("*").
		Qual("C", t.record.CType)
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
			Qual("C", "NULL"))

	g.
		If(jen.Id(goVarName).Op("!=").Nil()).
		BlockFunc(func(g *jen.Group) {
			g.
				Id(cVarName).
				Op("=").
				Parens(jen.
					Op(strings.Repeat("*", t.typ.indirectLevel)).
					Qual("C", t.typ.cTypeName)).
				Parens(jen.
					Id(goVarName).
					Op(".").
					Id("ToC").
					Call())
		})
}

func (t *TypeGeneratorRecord) generateParamGoVar(
	g *jen.Group, goVarName string, cVarName string, pkg string,
) {
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
			Call(jen.Id(cVarName)))
}

func (t *TypeGeneratorRecord) generateParamOutCVar(g *jen.Group, cVarName string) {
	g.
		Var().
		Id(cVarName).
		Op(strings.Repeat("*", t.typ.indirectLevel-1)).
		Qual("C", t.typ.cTypeName)
}

func (t *TypeGeneratorRecord) generateReturnFunctionDeclaration(g *jen.Group) {
	indirectLevel := t.typ.indirectLevel
	if indirectLevel == 2 {
		indirectLevel = 1
	}

	g.
		Op(strings.Repeat("*", indirectLevel)).
		Do(t.typ.qname.generate)
}

func (t *TypeGeneratorRecord) generateReturnFunctionDeclarationCtype(g *jen.Group) {
	g.
		Op(strings.Repeat("*", t.typ.indirectLevel+1)).
		Qual("C", t.record.CType)

}

func (t *TypeGeneratorRecord) generateReturnCToGo(g *jen.Group, isParam bool,
	cVarName string, goVarName string, pkg string,
	transferOwnership string, nullable bool) {

	cVarRef := jen.Id(cVarName)
	if isParam && t.typ.indirectLevel == 1 {
		cVarRef = jen.Op("&").Id(cVarName)
	}

	if t.typ.indirectLevel == 1 && nullable {
		g.
			Var().
			Id(goVarName).
			ParamsFunc(t.generateReturnFunctionDeclaration)

		g.
			If(jen.Id(cVarName).Op("==").Nil()).
			Block(jen.Id(goVarName).Op("=").Nil()).
			Else().
			BlockFunc(func(g *jen.Group) {
				g.
					Id(goVarName).
					Op("=").
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
			})
	} else {
		g.
			Id(goVarName).
			Op(":=").
			Do(func(s *jen.Statement) {
				if t.typ.indirectLevel == 0 {
					s.Op("*")
				}
			}).
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
}

func (t *TypeGeneratorRecord) generateArrayReturnCToGo(g *jen.Group, isParam bool,
	cVarName string, goVarName string, pkg string, transferOwnership string, nullable bool) {
	panic("unsupported")
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
	g.
		Parens(jen.Op("*").Qual("C", t.record.CType)).
		Parens(goVarReference.Dot("ToC").Call())
}
