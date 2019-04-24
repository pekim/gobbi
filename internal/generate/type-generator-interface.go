package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

type TypeGeneratorInterface struct {
	TypeGeneratorPanic
	typ   *Type
	iface *Interface
}

func TypeGeneratorInterfaceNew(typ *Type, iface *Interface) *TypeGeneratorInterface {
	return &TypeGeneratorInterface{
		typ:   typ,
		iface: iface,
	}
}

func (t *TypeGeneratorInterface) isSupportedAsField() (supported bool, reason string) {
	return false, "interface"

	//if t.typ.indirectLevel != 1 {
	//	return false, fmt.Sprintf("interface with indirection of %d", t.typ.indirectLevel)
	//}
	//
	//return true, ""
}

func (t *TypeGeneratorInterface) isSupportedAsParam(direction string) (supported bool, reason string) {
	if t.iface.Blacklist {
		return false, fmt.Sprintf("Blacklisted interface : %s", t.iface.CType)
	}

	if t.typ.indirectLevel > 2 {
		return false, fmt.Sprintf("interface with indirection level of %d",
			t.typ.indirectLevel)
	}

	return true, ""
}

func (t *TypeGeneratorInterface) isSupportedAsArrayParam(direction string) (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorInterface) isSupportedAsArrayParamC(direction string) (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorInterface) isSupportedAsParamC() (supported bool, reason string) {
	if t.iface.Blacklist {
		return false, fmt.Sprintf("Blacklisted interface : %s", t.iface.CType)
	}

	return true, ""
}

func (t *TypeGeneratorInterface) isSupportedAsArrayReturnValue() (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorInterface) isSupportedAsReturnValue() (supported bool, reason string) {
	return false, ""
	//if t.iface.Blacklist {
	//	return false, fmt.Sprintf("Blacklisted interface : %s", t.iface.CType)
	//}
	//
	//if t.typ.indirectLevel > 1 {
	//	return false, fmt.Sprintf("interface with indirection level of %d",
	//		t.typ.indirectLevel)
	//}
	//
	//return true, ""
}

func (t *TypeGeneratorInterface) isSupportedAsReturnCValue() (supported bool, reason string) {
	if t.iface.Blacklist {
		return false, fmt.Sprintf("Blacklisted interface : %s", t.iface.CType)
	}

	return true, ""
}

func (t *TypeGeneratorInterface) generateDeclaration(g *jen.Group, goVarName string) {
	g.
		Id(goVarName).
		Op("*").
		Do(t.typ.qname.generate)
}

func (t *TypeGeneratorInterface) generateArrayDeclaration(g *jen.Group, goVarName string) {
	g.
		Id(goVarName).
		Index().
		Op("*").
		Do(t.typ.qname.generate)
}

func (t *TypeGeneratorInterface) generateArrayDeclarationC(g *jen.Group, cVarName string) {
	g.
		Id(cVarName).
		Op("*").
		Qual("C", t.iface.CType)
}

func (t *TypeGeneratorInterface) generateDeclarationC(g *jen.Group, cVarName string) {
	g.
		Id(cVarName).
		Op("*").
		Qual("C", t.iface.CType)
}

func (t *TypeGeneratorInterface) generateParamCallArgument(g *jen.Group, cVarName string) {
	g.Id(cVarName)
}

func (t *TypeGeneratorInterface) generateParamOutCallArgument(g *jen.Group, cVarName string) {
	g.
		Op("&").
		Id(cVarName)
}

func (t *TypeGeneratorInterface) generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string) {
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

func (t *TypeGeneratorInterface) generateParamGoVar(
	g *jen.Group, goVarName string, cVarName string, pkg string,
) {
	g.
		Id(goVarName).
		Op(":=").
		Do(func(s *jen.Statement) {
			if pkg != "" {
				s.Qual(pkg, t.iface.newFromCFuncName)
			} else {
				s.Id(t.iface.newFromCFuncName)
			}
		}).
		Call(jen.
			Qual("unsafe", "Pointer").
			Call(jen.Id(cVarName)))
}

func (t *TypeGeneratorInterface) generateParamOutCVar(g *jen.Group, cVarName string) {
	g.
		Var().
		Id(cVarName).
		Op(strings.Repeat("*", t.typ.indirectLevel-1)).
		Qual("C", t.typ.cTypeName)
}

func (t *TypeGeneratorInterface) generateReturnFunctionDeclaration(g *jen.Group) {
	indirectLevel := t.typ.indirectLevel
	if indirectLevel == 2 {
		indirectLevel = 1
	}

	g.
		Op(strings.Repeat("*", indirectLevel)).
		Do(t.typ.qname.generate)
}

func (t *TypeGeneratorInterface) generateReturnFunctionDeclarationCtype(g *jen.Group) {
	g.
		Op(strings.Repeat("*", t.typ.indirectLevel+1)).
		Qual("C", t.iface.CType)

}

func (t *TypeGeneratorInterface) generateReturnCToGo(g *jen.Group, isParam bool,
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
							s.Qual(pkg, t.iface.newFromCFuncName)
						} else {
							s.Id(t.iface.newFromCFuncName)
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
				if pkg != "" {
					s.Qual(pkg, t.iface.newFromCFuncName)
				} else {
					s.Id(t.iface.newFromCFuncName)
				}
			}).
			Call(jen.
				Qual("unsafe", "Pointer").
				Call(cVarRef))
	}
}

func (t *TypeGeneratorInterface) generateArrayReturnCToGo(g *jen.Group, isParam bool,
	cVarName string, goVarName string, pkg string, transferOwnership string, nullable bool) {
	panic("unsupported")
}

func (t *TypeGeneratorInterface) generateCToGo(pkg string, cVarReference *jen.Statement) *jen.Statement {
	return jen.
		Do(func(s *jen.Statement) {
			if pkg != "" {
				s.Qual(pkg, t.iface.newFromCFuncName)
			} else {
				s.Id(t.iface.newFromCFuncName)
			}
		}).
		Call(cVarReference)
}

func (t *TypeGeneratorInterface) generateGoToC(g *jen.Group, goVarReference *jen.Statement) {
	g.
		Parens(jen.Op("*").Qual("C", t.iface.CType)).
		Parens(goVarReference.Dot("ToC").Call())
}
