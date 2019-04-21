package generate

import "C"
import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

type TypeGeneratorString struct {
	typ *Type
}

func TypeGeneratorStringNew(typ *Type) *TypeGeneratorString {
	return &TypeGeneratorString{
		typ: typ,
	}
}

func (t *TypeGeneratorString) isSupportedAsField() (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorString) isSupportedAsParam(direction string) (supported bool, reason string) {
	if direction != "out" && t.typ.indirectLevel > 1 {
		return false, fmt.Sprintf("in string with indirection level of %d",
			t.typ.indirectLevel)
	}

	return true, ""
}

func (t *TypeGeneratorString) isSupportedAsArrayParam(direction string) (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorString) isSupportedAsArrayParamC(direction string) (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorString) isSupportedAsParamC() (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorString) isSupportedAsArrayReturnValue() (supported bool, reason string) {
	return false, ""
	//return true, ""
}

func (t *TypeGeneratorString) isSupportedAsReturnValue() (supported bool, reason string) {
	return false, ""
	//if t.typ.indirectLevel > 1 {
	//	return false, fmt.Sprintf("string with indirection level of %d",
	//		t.typ.indirectLevel)
	//}
	//
	//return true, ""
}

func (t *TypeGeneratorString) isSupportedAsReturnCValue() (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorString) generateDeclaration(g *jen.Group, goVarName string) {
	g.
		Id(goVarName).
		Do(t.typ.qname.generate)
}

func (t *TypeGeneratorString) generateArrayDeclaration(g *jen.Group, goVarName string) {
	if goVarName != "" {
		g.
			Id(goVarName).
			Index().
			Do(t.typ.qname.generate)
	} else {
		g.
			Index().
			Do(t.typ.qname.generate)
	}
}

func (t *TypeGeneratorString) generateArrayDeclarationC(g *jen.Group, cVarName string) {
}

func (t *TypeGeneratorString) generateDeclarationC(g *jen.Group, cVarName string) {
	g.
		Id(cVarName).
		Op(strings.Repeat("*", t.typ.indirectLevel)).
		Qual("C", t.typ.cTypeName)
}

func (t *TypeGeneratorString) generateParamCallArgument(g *jen.Group, cVarName string) {
	g.Id(cVarName)
}

func (t *TypeGeneratorString) generateParamOutCallArgument(g *jen.Group, cVarName string) {
	g.
		Op("&").
		Id(cVarName)
}

func (t *TypeGeneratorString) generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string) {
	g.
		Id(cVarName).
		Op(":=").
		Qual("C", "CString").
		Call(jen.Id(goVarName))

	if transferOwnership != "none" {
		// ownership is transferred (to the library) so we should not
		// free the string memory
		return
	}

	g.
		Defer().
		Qual("C", "free").
		Call(jen.
			Qual("unsafe", "Pointer").
			Call(jen.Id(cVarName)))
}

func (t *TypeGeneratorString) generateParamGoVar(g *jen.Group, goVarName string, cVarName string, pkg string) {
	g.
		Id(goVarName).
		Op(":=").
		Qual("C", "GoString").
		Call(jen.Id(cVarName))
}

func (t *TypeGeneratorString) generateParamOutCVar(g *jen.Group, cVarName string) {
	g.
		Var().
		Id(cVarName).
		Op(strings.Repeat("*", t.typ.indirectLevel-1)).
		Qual("C", t.typ.cTypeName)
}

func (t *TypeGeneratorString) generateReturnFunctionDeclaration(g *jen.Group) {
	g.Do(t.typ.qname.generate)
}

func (t *TypeGeneratorString) generateReturnFunctionDeclarationCtype(g *jen.Group) {
}

func (t *TypeGeneratorString) generateReturnCToGo(g *jen.Group, isParam bool,
	cVarName string, goVarName string, pkg string, transferOwnership string, nullable bool) {
	g.
		Id(goVarName).
		Op(":=").
		Qual("C", "GoString").
		Call(jen.Id(cVarName))

	if transferOwnership == "none" {
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

func (t *TypeGeneratorString) generateArrayReturnCToGo(g *jen.Group, isParam bool,
	cVarName string, goVarName string, pkg string, transferOwnership string, nullable bool,
) {
	g.
		// retGo := []string{}
		Id(goVarName).
		Op(":=").
		Index().
		String().
		Values()

	g.
		// for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		For(
			jen.Id("p").Op(":=").Id("retC"),
			jen.Op("*").Id("p").Op("!=").Nil(),
			jen.
				Id("p").
				Op("=").
				Parens(jen.
					Op("**").
					Qual("C", "char")).
				Parens(jen.Qual("C", "gpointer").Parens(jen.
					Parens(jen.
						Id("uintptr").Parens(jen.Qual("C", "gpointer").Parens(jen.Id("p"))).
						Op("+").
						Id("uintptr").Parens(jen.Qual("C", "sizeof_gpointer"))))),
		).
		BlockFunc(func(g *jen.Group) {
			g.
				// s := C.GoString(*p)
				Id("s").
				Op(":=").
				Qual("C", "GoString").
				Call(jen.Op("*").Id("p"))

			g.
				// retGo = append(retGo, s)
				Id("retGo").
				Op("=").
				Append(
					jen.Id("retGo"),
					jen.Id("s"),
				)
		})

	if transferOwnership == "none" {
		// the library will be responsible for the memory
		return
	}

	if transferOwnership == "full" {
		g.
			Defer().
			Qual("C", "g_strfreev").
			Call(jen.Id(cVarName))
	}
}

func (t *TypeGeneratorString) generateCToGo(pkg string, cVarReference *jen.Statement) *jen.Statement {
	return jen.
		Qual("C", "GoString").
		Call(cVarReference)
}

func (t *TypeGeneratorString) generateGoToC(g *jen.Group, goVarReference *jen.Statement) {
	g.
		Qual("C", "CString").
		Call(goVarReference)
}
