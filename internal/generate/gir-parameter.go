package generate

import "C"
import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type Parameter struct {
	Name      string `xml:"name,attr"`
	Direction string `xml:"direction,attr"`
	Argument
	AllowNone bool `xml:"allow-none,attr"`
	//Doc       *Doc      `xml:"doc"`
	Varargs *struct{} `xml:"varargs"`

	goVarName string
	namespace *Namespace

	lengthParam    *Parameter
	lengthForParam *Parameter
	argcParam      *Parameter
	argvParam      *Parameter
}

func (p *Parameter) init(ns *Namespace) {
	p.namespace = ns
	p.Type.init(ns)
	p.Array.init(ns)
	p.goVarName = makeUnexportedGoName(p.Name)
}

func (p *Parameter) isIn() bool {
	return p.Direction == "" || p.Direction == "in" || p.Direction == "inout"
}

func (p *Parameter) isOut() bool {
	return p.Direction == "out" || p.Direction == "inout"
}

func (p *Parameter) isInOut() bool {
	return p.Direction == "inout"
}

func (p *Parameter) sysParamGoType() *jen.Statement {
	if p.Type != nil {
		return p.Type.sysParamGoType()
	}

	if p.Array != nil {
		return p.Array.sysParamGoType()
	}

	panic(fmt.Sprintf("Parameter is not a type or an array: %s", p.Name))
}

func (p *Parameter) generateSysCArg(g *jen.Group, goVarName string, cVarName string) {
	if p.Array != nil {
		p.generateSysCArgArray(g, goVarName, cVarName)
		return
	}

	if p.Type.isString() {
		p.generateSysCArgString(g, goVarName, cVarName)
		return
	}

	cValue := p.generateSysCValue(goVarName)
	g.Id(cVarName).Op(":=").Add(cValue)
}

func (p *Parameter) generateSysCValue(goVarName string) *jen.Statement {
	goValue := jen.Id(goVarName)

	if p.Type.CType == "gboolean" {
		return jen.Id("toCBool").Call(goValue)
	}

	if p.Type.cType.indirectionCount > 0 {
		goValue = jenUnsafePointer().Call(goValue)
	}

	// Atoms are really pointers underneath.
	if p.Type.CType == "GdkAtom" {
		goValue = jenUnsafePointer().Call(goValue)
	}

	return jen.Parens(p.Type.jenGoCType()).Parens(goValue)
}

func (p *Parameter) generateSysCArgString(g *jen.Group, goVarName string, cVarName string) {
	if p.Type.cType.indirectionCount == 1 {
		p.generateSysCArgStringSimple(g, goVarName, cVarName)
		return
	}

	if p.Type.cType.indirectionCount == 2 {
		p.generateSysCArgStringPointer(g, goVarName, cVarName)
		return
	}

	panic(fmt.Sprintf("Unsupported indirection count (%d) for string param ", p.Type.cType.indirectionCount))
}

func (p *Parameter) generateSysCArgStringSimple(g *jen.Group, goVarName string, cVarName string) {
	cValue := jen.Parens(p.Type.jenGoCType()).Parens(jen.Qual("C", "CString").Call(jen.Id(goVarName)))
	g.Id(cVarName).Op(":=").Add(cValue)

	if p.transferNone() {
		g.Defer().Qual("C", "free").Call(jenUnsafePointer().Call(jen.Id(cVarName)))
	}
}

func (p *Parameter) generateSysCArgStringPointer(g *jen.Group, goVarName string, cVarName string) {
	cStringVarName := cVarName + "String"

	// var cValue0String *C.gchar
	g.Var().Id(cStringVarName).Op("*").Qual("C", "gchar")

	if p.isIn() {
		cValue := jen.Parens(jen.Op("*").Qual("C", "gchar")).Parens(jen.Qual("C", "CString").Call(jen.Op("*").Id(goVarName)))
		g.Id(cStringVarName).Op("=").Add(cValue)
	}

	g.Id(cVarName).Op(":=").Op("&").Id(cStringVarName)

	if p.isIn() && p.transferNone() {
		g.Defer().Qual("C", "free").Call(jenUnsafePointer().Call(jen.Id(cVarName)))
	}
}

func (p *Parameter) generateSysCArgOut(g *jen.Group, goVarName string, cVarName string) {
	if !p.isOut() {
		return
	}

	if p.Type != nil && p.Type.isString() && p.Type.cType.indirectionCount == 2 {
		p.generateSysCArgStringPointerOut(g, goVarName, cVarName)
	}
}

func (p *Parameter) generateSysCArgStringPointerOut(g *jen.Group, goVarName string, cVarName string) {
	cStringVarName := cVarName + "String"
	goStringVarName := goVarName + "String"

	g.Id(goStringVarName).Op(":=").Qual("C", "GoString").Call(jen.Id(cStringVarName))
	g.Op("*").Id(goVarName).Op("=").Id(goStringVarName)
}

func (p *Parameter) generateSysCArgArray(g *jen.Group, goVarName string, cVarName string) {
	if p.Array.Type.isString() {
		if p.Array.cType.indirectionCount == 2 {
			p.generateSysCArgArrayString(g, goVarName, cVarName)
			return
		}

		if p.Array.cType.indirectionCount == 3 {
			p.generateSysCArgArrayStringPointer(g, goVarName, cVarName)
			return
		}

		panic(fmt.Sprintf("Unsupported indirection count (%d) for array string param", p.Array.Type.cType.indirectionCount))
	}
}

func (p *Parameter) generateSysCArgArrayString(g *jen.Group, goVarName string, cVarName string) {
	// Convert Go slice of strings to C array of null terminated strings.
	goSliceVarName := p.generateGoArrayStringToC(g, goVarName, cVarName)

	// cValue2 := &cValue2Array[0]
	g.Id(cVarName).Op(":=").Op("&").Id(goSliceVarName).Index(jen.Lit(0))
}

func (p *Parameter) generateSysCArgArrayStringPointer(g *jen.Group, goVarName string, cVarName string) {
	goVarIndirectedName := goVarName + "Indirected"
	// param2Indirected := *param2
	g.Id(goVarIndirectedName).Op(":=").Op("*").Id(goVarName)

	// Convert Go slice of strings to C array of null terminated strings.
	goCPtrSliceVarName := p.generateGoArrayStringToC(g, goVarIndirectedName, cVarName)

	cArrayPointerVarName := cVarName + "ArrayPointer"

	// var cValue2 ***C.gchar
	g.Var().Id(cVarName).Op("***").Qual("C", "gchar")

	// var cValue2ArrayPointer **C.gchar
	g.Var().Id(cArrayPointerVarName).Op("**").Qual("C", "gchar")

	// if len(param2Slice) > 0 {
	//   cValue2ArrayPointer = &param2Slice[0]
	// }
	g.If(jen.Len(jen.Id(goCPtrSliceVarName)).Op(">").Lit(0)).
		Block(jen.Id(cArrayPointerVarName).Op("=").Op("&").Id(goCPtrSliceVarName).Index(jen.Lit(0)))

	// cValue2 = &cValue2ArrayPointer
	g.Id(cVarName).Op("=").Op("&").Id(cArrayPointerVarName)
}

// generateGoArrayStringToC converts a Go slice of strings to a C array of null terminated strings.
func (p *Parameter) generateGoArrayStringToC(g *jen.Group, goVarName string, cVarName string) string {
	// cValue2Array := C.malloc((C.ulong)(len(param2)) * C.sizeof_gpointer)
	// param2Slice := (*[1 << 28]C.gchar)(unsafe.Pointer(cParam2Array))[:param2Len:param2Len]
	//
	// for param2i, str := range param2 {
	//     cValue2Array[param2i] = (*C.gchar)(C.CString(param2[param2i]))
	// }

	lenVarName := goVarName + "Len"
	cArrayVarName := cVarName + "Array"
	goSliceVarName := goVarName + "Slice"

	// param2Len := len(param2)
	g.Id(lenVarName).Op(":=").Len(jen.Id(goVarName))

	// cValue2Array := C.malloc((C.ulong)(len(param2)) * C.sizeof_gpointer)
	count := jen.
		Parens(jen.Qual("C", "ulong")).
		Parens(jen.Id(lenVarName))
	g.Id(cArrayVarName).Op(":=").Qual("C", "malloc").Call(
		count.Op("*").Qual("C", "sizeof_gpointer"),
	)

	// defer C.free(unsafe.Pointer(cValue2Array))
	g.Defer().Qual("C", "free").Call(jenUnsafePointer().Call(jen.Id(cArrayVarName)))

	// param2Slice := (*[1 << 28]C.gchar)(unsafe.Pointer(cParam2Array))[:param2Len:param2Len]
	g.
		// param2Slice :=
		Id(goSliceVarName).Op(":=").
		// (*[1 << 28]C.gchar)
		Parens(
			jen.Op("*").Index(jen.Lit(1).Op("<<").Lit(30)).Parens(jen.Op("*").Qual("C", "gchar"))).
		// (unsafe.Pointer(cParam2Array))
		Parens(jenUnsafePointer().Call(jen.Id(cArrayVarName))).
		// [:param2Len:param2Len]
		Index(jen.Op(":").Id(lenVarName).Op(":").Id(lenVarName))

	indexVarName := goVarName + "i"
	stringVarName := goVarName + "String"
	// for param2i, param2String := range param2 {
	//     cValue2Array[param2i] = (*C.gchar)(C.CString(param2String))
	//     defer C.free(unsafe.Pointer(cValue2Array[param2i]))
	// }
	g.
		For(jen.Id(indexVarName).Op(",").Id(stringVarName).Op(":=").Range().Id(goVarName)).
		Block(
			jen.
				Id(goSliceVarName).Index(jen.Id(indexVarName)).Op("=").
				Parens(jen.Op("*").Qual("C", "gchar")).
				Parens(jen.Qual("C", "CString").Call(jen.Id(stringVarName))),

			// defer C.free(unsafe.Pointer(cValue2Array[param2i]))
			jen.Defer().Qual("C", "free").Call(jenUnsafePointer().Call(jen.Id(goSliceVarName).Index(jen.Id(indexVarName)))),
		)

	return goSliceVarName
}
