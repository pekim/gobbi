package generate

import (
	"fmt"

	"github.com/pekim/jennifer/jen"
)

type TypeGeneratorNumber struct {
	TypeGeneratorPanic
	typ *Type
}

func TypeGeneratorIntegerNew(typ *Type) *TypeGeneratorNumber {
	return &TypeGeneratorNumber{
		typ: typ,
	}
}

func (t *TypeGeneratorNumber) isSupportedAsParam(direction string) (supported bool, reason string) {
	if t.typ.indirectLevel > 1 {
		return false, fmt.Sprintf("%s with indirection level of %d",
			t.typ.CType, t.typ.indirectLevel)
	}

	return true, ""
}

func (t *TypeGeneratorNumber) isSupportedAsArrayParam(direction string) (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorNumber) isSupportedAsArrayParamC(direction string) (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorNumber) isSupportedAsParamC() (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorNumber) isSupportedAsField() (supported bool, reason string) {
	if t.typ.indirectLevel > 0 {
		return false, fmt.Sprintf("%s with indirection level of %d",
			t.typ.CType, t.typ.indirectLevel)
	}

	return true, ""
}

func (t *TypeGeneratorNumber) isSupportedAsArrayReturnValue() (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorNumber) isSupportedAsReturnValue() (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorNumber) isSupportedAsReturnCValue() (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorNumber) generateDeclaration(g *jen.Group, goVarName string) {
	g.
		Id(goVarName).
		Do(t.typ.qname.generate)
}

func (t *TypeGeneratorNumber) generateArrayDeclaration(g *jen.Group, goVarName string) {
	g.
		Id(goVarName).
		Index().
		Do(t.typ.qname.generate)
}

func (t *TypeGeneratorNumber) generateArrayDeclarationC(g *jen.Group, cVarName string) {
}

func (t *TypeGeneratorNumber) generateDeclarationC(g *jen.Group, cVarName string) {
	g.Id(cVarName).Qual("C", t.typ.CType)
}

func (t *TypeGeneratorNumber) generateParamCallArgument(g *jen.Group, cVarName string) {
	g.
		Do(func(s *jen.Statement) {
			if t.typ.indirectLevel == 1 {
				s.Op("&")
			}
		}).
		Id(cVarName)
}

func (t *TypeGeneratorNumber) generateParamOutCallArgument(g *jen.Group, cVarName string) {
	g.
		Op("&").
		Id(cVarName)
}

func (t *TypeGeneratorNumber) generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string) {
	g.
		Id(cVarName).
		Op(":=").
		Parens(jen.Qual("C", t.typ.cTypeName)).
		Parens(jen.Id(goVarName))
}

func (t *TypeGeneratorNumber) generateParamGoVar(g *jen.Group, goVarName string, cVarName string, pkg string) {
	g.
		Id(goVarName).
		Op(":=").
		Id(numberCTypeMap[t.typ.CType]).
		Call(jen.Id(cVarName))

}

func (t *TypeGeneratorNumber) generateParamOutCVar(g *jen.Group, cVarName string) {
	g.
		Var().
		Id(cVarName).
		Qual("C", t.typ.cTypeName)
}

func (t *TypeGeneratorNumber) generateReturnFunctionDeclaration(g *jen.Group) {
	g.Do(t.typ.qname.generate)
}

func (t *TypeGeneratorNumber) generateReturnFunctionDeclarationCtype(g *jen.Group) {
	g.Qual("C", t.typ.CType)
}

func (t *TypeGeneratorNumber) generateReturnCToGo(g *jen.Group, isParam bool,
	cVarName string, goVarName string, pkg string, transferOwnership string, nullable bool) {
	g.
		Id(goVarName).
		Op(":=").
		Parens(jen.Do(t.typ.qname.generate)).
		Parens(jen.Do(func(s *jen.Statement) {
			if t.typ.Name == "gpointer" {
				s.
					Qual("unsafe", "Pointer").
					CallFunc(func(g *jen.Group) {
						if t.typ.indirectLevel == 1 {
							g.
								Op("&").
								Id(cVarName)
						} else {
							g.Id(cVarName)
						}
					})
			} else {
				s.Id(cVarName)
			}
		}))
}

func (t *TypeGeneratorNumber) generateCToGo(pkg string, cVarReference *jen.Statement) *jen.Statement {
	return jen.
		Parens(jen.Do(func(s *jen.Statement) {
			if t.typ.indirectLevel == 1 {
				s.Op("*")
			}
			t.typ.qname.generate(s)
		})).
		Parens(jen.Do(func(s *jen.Statement) {
			if t.typ.indirectLevel == 1 {
				s.Op("&")
			}
			s.Add(cVarReference)
		}))
}

func (t *TypeGeneratorNumber) generateGoToC(g *jen.Group, goVarReference *jen.Statement) {
	g.
		Parens(jen.Qual("C", t.typ.cTypeName)).
		Parens(goVarReference)
}

func (t *TypeGeneratorNumber) isSupportedByVersion(version *Version) bool {
	return true
}
