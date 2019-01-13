package generate

import "C"
import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

type Array struct {
	Namespace *Namespace

	Type           *Type  `xml:"type"`
	CType          string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	FixedSize      *int   `xml:"fixed-size,attr"`
	Length         *int   `xml:"length,attr"`
	ZeroTerminated string `xml:"zero-terminated,attr"`

	lengthParam *Parameter
}

func (a *Array) init(ns *Namespace) {
	a.Namespace = ns

	// Some array's Type has a Name but no CType.
	// In all observed cases this is an integer type, usually 'guint8'.
	// For these cases use the Type's Name as the CType.
	_, isInteger := integerCTypeMap[a.Type.Name]
	if isInteger && a.Type.CType == "" {
		a.Type.CType = a.Type.Name
	}

	a.Type.init(ns)
}

func (a *Array) generateDeclaration(g *jen.Group, goVarName string) {
	a.Type.generator.generateArrayDeclaration(g, goVarName)
}

func (a *Array) generateReturnDeclaration(g *jen.Group) {
	a.Type.generator.generateArrayDeclaration(g, "")
}

func (a *Array) generateDeclarationC(g *jen.Group, cVarName string) {
	g.
		Id(cVarName).
		Qual("C", "gpointer")
}

func (a *Array) generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string) {
	cVarArrayName := cVarName + "_array"

	// c_?_array := make([]C..., len(?)+1, len(?)+1)
	g.
		Id(cVarArrayName).
		Op(":=").
		Make(jen.
			Index().
			Op(strings.Repeat("*", a.Type.indirectLevel)).
			Qual("C", a.Type.cTypeName),
			jen.Len(jen.Id(goVarName)).Op("+").Lit(1),
			jen.Len(jen.Id(goVarName)).Op("+").Lit(1),
		)

	cVarArrayNamePtr := cVarArrayName + "Ptr"

	g.
		// for i, item := range goVarName {
		For(
			jen.
				Id("i").
				Op(",").
				Id("item").
				Op(":=").
				Range().
				Id(goVarName),
		).
		BlockFunc(func(g *jen.Group) {
			a.Type.generator.generateParamCVar(g, "c", "item", transferOwnership)

			g.
				Id(cVarArrayName).
				Index(jen.Id("i")).
				Op("=").
				Id("c")
		})

	// terminate array with a nil or 0
	if strings.HasSuffix(a.CType, "**") {
		g.
			Id(cVarArrayName).
			Index(jen.Len(jen.Id(goVarName))).
			Op("=").
			Nil()
	} else {
		g.
			Id(cVarArrayName).
			Index(jen.Len(jen.Id(goVarName))).
			Op("=").
			Lit(0)
	}

	// c_?_arrayPtr := &c_?_array[0]
	g.
		Id(cVarArrayNamePtr).
		Op(":=").
		Op("&").
		Id(cVarArrayName).
		Index(jen.Lit(0))

	if a.CType == "void*" {
		g.
			Id(cVarName).
			Op(":=").
			Parens(jen.
				Qual("unsafe", "Pointer").
				Call(jen.Id(cVarArrayNamePtr)))
	} else {
		cType := strings.TrimRight(a.CType, "*")
		indirectLevel := len(a.CType) - len(cType)
		if indirectLevel == 0 && !strings.HasSuffix(a.CType, "pointer") {
			indirectLevel = 1
		}
		indirect := strings.Repeat("*", indirectLevel)

		// c_? := (*C...)(unsafe.Pointer(c_?_arrayPtr))
		g.
			Id(cVarName).
			Op(":=").
			Parens(jen.Op(indirect).Qual("C", cType)).
			Parens(jen.
				Qual("unsafe", "Pointer").
				Call(jen.Id(cVarArrayNamePtr)))
	}
}

func (a *Array) generateArrayLenParamCVar(g *jen.Group, cVarName string, arrayGoVarName string, ctype string) {
	g.
		Id(cVarName).
		Op(":=").
		Parens(jen.Qual("C", ctype)).
		Parens(jen.Len(jen.Id(arrayGoVarName)))
}

func (a *Array) generateParamGoVar(g *jen.Group, param *Parameter) {
	qname := QNameNew(a.Type.Namespace, a.Type.Name)
	iface, found := qname.namespace.interfaceForName(qname.name)
	if !found {
		fmt.Println(qname.namespace.Name, qname.name)
		panic(fmt.Sprintf("Failed to find record for %s", a.Type.Name))
	}

	g.
		Id(param.goVarName).
		Op(":=").
		Make(
			jen.Index().Op("*").Id(param.Array.Type.qname.name),
			jen.Id("int").Parens(jen.Id(param.Array.lengthParam.cVarName)),
			jen.Id("int").Parens(jen.Id(param.Array.lengthParam.cVarName)))

	g.
		// for i := 0; i < int(c_n_files); i++ {
		For(
			jen.Id("i").Op(":=").Lit(0),
			jen.Id("i").Op("<").Int().Call(jen.Id(param.Array.lengthParam.cVarName)),
			jen.Id("i").Op("++"),
		).
		BlockFunc(func(g *jen.Group) {
			// _item := FileNewFromC(unsafe.Pointer(*(*C.gpointer)(c_files)))
			g.
				Id("_item").
				Op(":=").
				Id(iface.newFromCFuncName).
				Call(jen.Qual("unsafe", "Pointer").Parens(jen.
					Op("*").
					Parens(jen.Op("*").Qual("C", "gpointer")).
					Parens(jen.Id(param.cVarName))))

			// files[i] = _item
			g.
				Id(param.goVarName).
				Index(jen.Id("i")).
				Op("=").
				Id("_item")

			// c_files = C.gpointer(uintptr(c_files) + uintptr(C.sizeof_gpointer))
			g.
				Id(param.cVarName).
				Op("=").
				Qual("C", "gpointer").
				Parens(jen.
					Id("uintptr").Parens(jen.Id(param.cVarName)).
					Op("+").
					Id("uintptr").Parens(jen.Qual("C", "sizeof_gpointer")))
		})
}

func (a *Array) generateParamCallArgument(g *jen.Group, cVarName string) {
	g.Id(cVarName)
}
