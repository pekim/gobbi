package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strconv"
)

func (p *Parameter) sysParamGoType(decrementIndirectionCount bool) *jen.Statement {
	if p.Type != nil {
		// Atoms are really pointers underneath.
		if p.Type.CType == "GdkAtom" {
			return jenUnsafePointer()
		}

		return jen.
			Add(p.Type.sysParamGoType(decrementIndirectionCount))
	}

	star := ""
	if p.isOut() {
		star = "*"
	}
	if p.Array != nil {
		return jen.
			Op(star).
			Add(p.Array.sysParamGoType())
	}

	panic(fmt.Sprintf("Parameter is not a type or an array: %s", p.Name))
}

func (p *Parameter) sysReturnGoType() *jen.Statement {
	if p.Type != nil {
		// Atoms are really pointers underneath.
		if p.Type.CType == "GdkAtom" {
			return jenUnsafePointer()
		}

		return jen.
			Add(p.Type.sysParamGoType(true))
	}

	if p.Array != nil {
		panic("TODO")
	}

	panic(fmt.Sprintf("Out parameter is not a type or an array: %s", p.Name))
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

	if p.Array != nil {
		if p.lengthParam == nil {
			panic(fmt.Sprintf("No length param for %s", p.Name))
		}

		if p.Array.Type.isString() && p.Array.cType.indirectionCount == 3 {
			p.generateSysCArgArrayStringPointerOut(g, goVarName, cVarName)
		} else {
			p.generateSysCArgArrayPointerOut(g, goVarName, cVarName)
		}
	}
}

func (p *Parameter) generateSysCArgStringPointerOut(g *jen.Group, goVarName string, cVarName string) {
	cStringVarName := cVarName + "String"
	goStringVarName := goVarName + "String"

	g.Line()

	g.Id(goStringVarName).Op(":=").Qual("C", "GoString").Call(jen.Id(cStringVarName))
	g.Op("*").Id(goVarName).Op("=").Id(goStringVarName)
}

func (p *Parameter) generateSysCArgArrayStringPointerOut(g *jen.Group, goVarName string, cVarName string) {
	g.Line()

	outVarName := goVarName + "Out"
	lenVarName := outVarName + "Len"
	cSliceVarName := outVarName + "CSlice"
	cArrayVarName := cVarName + "ArrayPointer"
	lengthParamName := "cValue" + strconv.Itoa(p.lengthParamN)

	// param2OutLen := (int)(*cValue?)
	g.Id(lenVarName).Op(":=").Int().Parens(jen.Op("*").Id(lengthParamName))

	// param2Out := make([]string, param2OutLen, param2OutLen)
	g.Id(outVarName).Op(":=").Make(
		jen.Index().String(),
		jen.Id(lenVarName),
		jen.Id(lenVarName),
	)

	// if (param2OutLen > 0) {...}
	g.If(jen.Id(lenVarName).Op(">").Lit(0)).BlockFunc(func(g *jen.Group) {
		// param2OutCSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue2ArrayPointer))[:param2OutLen:param2OutLen]
		g.Id(cSliceVarName).Op(":=").
			// (*[1 << 30]C.gchar)
			Parens(
				jen.Op("*").Index(jen.Lit(1).Op("<<").Lit(30)).Parens(jen.Op("*").Qual("C", "gchar"))).
			// (unsafe.Pointer(cParam2Array))
			Parens(jenUnsafePointer().Call(jen.Id(cArrayVarName))).
			// [:param2Len:param2Len]
			Index(jen.Op(":").Id(lenVarName).Op(":").Id(lenVarName))

		indexVarName := outVarName + "i"
		cStringVarName := outVarName + "CString"
		// for param2Outi, param2OutCString := range param2OutCSlice {
		//    param2Out[param2Outi] = C.GoString(param2OutCString)
		// }
		g.For(jen.Id(indexVarName).Op(",").Id(cStringVarName).Op(":=").Range().Id(cSliceVarName)).
			Block(jen.Id(outVarName).Index(jen.Id(indexVarName)).Op("=").Qual("C", "GoString").Call(jen.Id(cStringVarName)))
	})

	// *param2 = param2Out
	g.Op("*").Id(goVarName).Op("=").Id(outVarName)
}

func (p *Parameter) generateSysCArgArrayPointerOut(g *jen.Group, goVarName string, cVarName string) {
	g.Line()

	outVarName := goVarName + "Out"
	lenVarName := outVarName + "Len"
	cArrayVarName := cVarName + "ArrayPointer"
	lengthParamName := "cValue" + strconv.Itoa(p.lengthParamN)

	// param2OutLen := (int)(*cValue?)
	g.Id(lenVarName).Op(":=").Int().Parens(jen.Op("*").Id(lengthParamName))

	// param2Out := make([]SomeType, param2OutLen, param2OutLen)
	g.Id(outVarName).Op(":=").Make(
		jen.Add(p.Array.sysParamGoType()),
		jen.Id(lenVarName),
		jen.Id(lenVarName),
	)

	// if (param2OutLen > 0) {...}
	g.If(jen.Id(lenVarName).Op(">").Lit(0)).BlockFunc(func(g *jen.Group) {
		// param2OutCSlice := (*[1 << 30](*C.Sometype))(unsafe.Pointer(cValue2ArrayPointer))[:param2OutLen:param2OutLen]
		g.Id(outVarName).Op("=").
			// (*[1 << 30]C.SomeType)
			Parens(
				jen.Op("*").Index(jen.Lit(1).Op("<<").Lit(30)).Parens(p.Array.Type.sysParamGoPlainType())).
			// (unsafe.Pointer(cParam2Array))
			Parens(jenUnsafePointer().Call(jen.Id(cArrayVarName))).
			// [:param2Len:param2Len]
			Index(jen.Op(":").Id(lenVarName).Op(":").Id(lenVarName))
	})

	// *param2 = param2Out
	g.Op("*").Id(goVarName).Op("=").Id(outVarName)
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

		panic(fmt.Sprintf("Unsupported indirection count (%d) for array string param %s", p.Array.cType.indirectionCount, p.Name))
	}

	if p.isOut() {
		p.generateSysCArgArrayNonStringOut(g, goVarName, cVarName)
		return
	} else {
		p.generateSysCArgArrayNonString(g, goVarName, cVarName)
		return
	}

	panic(fmt.Sprintf("Unsupported indirection count (%d) for array param %s", p.Array.cType.indirectionCount, p.Name))
}

func (p *Parameter) generateSysCArgArrayString(g *jen.Group, goVarName string, cVarName string) {
	// Convert Go slice of strings to C array of null terminated strings.
	goSliceVarName := p.generateGoArrayStringToC(g, goVarName, cVarName)

	// cValue2 := &cValue2Array[0]
	g.Id(cVarName).Op(":=").Op("&").Id(goSliceVarName).Index(jen.Lit(0))
}

func (p *Parameter) generateSysCArgArrayStringPointer(g *jen.Group, goVarName string, cVarName string) {
	cArrayPointerVarName := cVarName + "ArrayPointer"

	// var cValue2ArrayPointer **C.gchar
	g.Var().Id(cArrayPointerVarName).Op("**").Qual("C", "gchar")

	// cValue2 = &cValue2ArrayPointer
	g.Id(cVarName).Op(":=").Op("&").Id(cArrayPointerVarName)

	if p.isIn() {
		goVarIndirectedName := goVarName + "Indirected"
		// param2Indirected := *param2
		g.Id(goVarIndirectedName).Op(":=").Op("*").Id(goVarName)

		// Convert Go slice of strings to C array of null terminated strings.
		goCPtrSliceVarName := p.generateGoArrayStringToC(g, goVarIndirectedName, cVarName)

		// if len(param2Slice) > 0 {
		//   cValue2ArrayPointer = &param2Slice[0]
		// }
		g.If(jen.Len(jen.Id(goCPtrSliceVarName)).Op(">").Lit(0)).
			Block(jen.Id(cArrayPointerVarName).Op("=").Op("&").Id(goCPtrSliceVarName).Index(jen.Lit(0)))
	}
}

// generateGoArrayStringToC converts a Go slice of strings to a C array of null terminated strings.
func (p *Parameter) generateGoArrayStringToC(g *jen.Group, goVarName string, cVarName string) string {
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

	// param2Slice := (*[1 << 30]C.gchar)(unsafe.Pointer(cParam2Array))[:param2Len:param2Len]
	g.
		// param2Slice :=
		Id(goSliceVarName).Op(":=").
		// (*[1 << 30]C.gchar)
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

func (p *Parameter) generateSysCArgArrayNonString(g *jen.Group, goVarName string, cVarName string) {
	cType := jen.Op(p.Array.cType.stars).Qual("C", p.Array.cType.typ)
	if p.Array.cType.typ == "void" && p.Array.cType.indirectionCount == 1 {
		cType = jenUnsafePointer()
	}

	// cValue2 := ([*]C.SomeType)(unsafe.Pointer(&cParam2[0]))
	g.
		Id(cVarName).
		Op(":=").
		Parens(cType).
		Parens(jenUnsafePointer().Call(jen.Op("&").Id(goVarName).Index(jen.Lit(0))))
}

func (p *Parameter) generateSysCArgArrayNonStringOut(g *jen.Group, goVarName string, cVarName string) {
	cArrayPointerVarName := cVarName + "ArrayPointer"

	// var cValue2ArrayPointer *[*]C.SomeType
	g.Var().Id(cArrayPointerVarName).Parens(jen.Op(p.Array.Type.cType.stars).Qual("C", p.Array.Type.cType.typ))

	// cValue2 = &cValue2ArrayPointer
	g.Id(cVarName).Op(":=").Op("&").Id(cArrayPointerVarName)

	if p.isIn() {
		panic(fmt.Sprintf("Unsupported inout array param %s", p.Name))
	}
}
