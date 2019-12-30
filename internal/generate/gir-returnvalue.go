package generate

import (
	"github.com/dave/jennifer/jen"
	"strconv"
)

type ReturnValue struct {
	Namespace *Namespace

	TransferOwnership string `xml:"transfer-ownership,attr"`
	Nullable          bool   `xml:"nullable,attr"`
	//Doc               *Doc   `xml:"doc"`
	Type  *Type  `xml:"type"`
	Array *Array `xml:"array"`
}

func (r *ReturnValue) init(ns *Namespace) {
	r.Namespace = ns
	r.Type.init(ns)
	r.Array.init(ns)
}

func (r *ReturnValue) isVoid() bool {
	if r == nil {
		return true
	}

	if r.Type == nil && r.Array == nil {
		return true
	}

	if r.Type != nil && r.Type.Name == "none" {
		return true
	}

	return false
}

func (r *ReturnValue) isSupported() (bool, string) {
	if r.isVoid() {
		return true, ""
	}

	if r.Type != nil && r.Type.isCallback() {
		return false, "has callback return "
	}

	if r.Array != nil {
		if supported, reason := r.Array.isSupported(); !supported {
			return supported, reason
		}

		if r.Array.Length == nil {
			return false, "no array length"
		}
	}

	return true, ""
}

func (r *ReturnValue) generateSysGoValue(g *jen.Group, cVarName string) *jen.Statement {
	if r.Type != nil {
		return r.generateSysGoTypeValue(cVarName)
	}

	if r.Array != nil {
		return r.generateSysGoArrayValue(g, cVarName)
	}

	panic("Unsupported return value; no Type or Array")
}

func (r *ReturnValue) generateSysGoTypeValue(cVarName string) *jen.Statement {
	if r.Type.CType == "GdkAtom" {
		if r.Type.isQualifiedName() {
			return jen.Qual(r.Type.foreignNamespace.goFullSysPackageName, r.Type.foreignName).Parens(jenUnsafePointer().Call(jen.Id(cVarName)))
		}
		return jen.Id(r.Type.Name).Parens(jen.Id(cVarName))
	}

	if r.Type.isString() {
		if r.Type.cType.typ == "guchar" {
			return jen.Qual("C", "GoString").Call(
				jen.Parens(jen.Id("*").Qual("C", "char")).Parens(jenUnsafePointer().Call(jen.Id(cVarName))))
		}
		return jen.Qual("C", "GoString").Call(jen.Id(cVarName))
	}

	if r.Type.isClass() || r.Type.isRecord() || r.Type.isInterface() || r.Type.isUnion() {
		return jenUnsafePointer().Call(jen.Id(cVarName))
	}

	if r.Type.CType == "gboolean" {
		return jen.Id("toGoBool").Call(jen.Id(cVarName))
	}

	return jen.Parens(r.Type.sysParamGoType(false)).Parens(jen.Id("ret"))
}

func (r *ReturnValue) generateSysGoArrayValue(g *jen.Group, cVarName string) *jen.Statement {
	retVarName := "retGo"
	lenVarName := "retLen"
	sliceVarName := "retSlice"
	lengthParamName := "cValue" + strconv.Itoa(*r.Array.Length)

	// retLen := (int)(*cValue?)
	g.Id(lenVarName).Op(":=").Int().Parens(jen.Op("*").Id(lengthParamName))

	// retGo := make([]SomeType, param2OutLen, param2OutLen)
	g.Id(retVarName).Op(":=").Make(
		jen.Add(r.Array.sysParamGoType()),
		jen.Id(lenVarName),
		jen.Id(lenVarName),
	)

	if r.Array.Type.isString() {
		// param2Slice := (*[1 << 30]C.gchar)(unsafe.Pointer(cParam2Array))[:param2Len:param2Len]
		g.
			// param2Slice :=
			Id(sliceVarName).Op(":=").
			// (*[1 << 30]C.gchar)
			Parens(
				jen.Op("*").Index(jen.Lit(1).Op("<<").Lit(30)).Parens(jen.Op("*").Qual("C", "gchar"))).
			// (unsafe.Pointer(cParam2Array))
			Parens(jenUnsafePointer().Call(jen.Id(cVarName))).
			// [:param2Len:param2Len]
			Index(jen.Op(":").Id(lenVarName).Op(":").Id(lenVarName))

		indexVarName := retVarName + "i"
		cStringVarName := retVarName + "String"
		// for retGoi, retGoCString := range retSlice {
		//    retGo[retGoi] = C.GoString(retGoCString)
		// }
		g.For(jen.Id(indexVarName).Op(",").Id(cStringVarName).Op(":=").Range().Id(sliceVarName)).
			Block(jen.Id(retVarName).Index(jen.Id(indexVarName)).Op("=").Qual("C", "GoString").Call(jen.Id(cStringVarName)))
	} else {
		// if (retLen > 0) {...}
		g.If(jen.Id(lenVarName).Op(">").Lit(0)).BlockFunc(func(g *jen.Group) {
			// retGo := (*[1 << 30](*C.Sometype))(unsafe.Pointer(ret))[:retLen:retLen]
			g.Id(retVarName).Op("=").
				// (*[1 << 30]C.SomeType)
				Parens(
					jen.Op("*").Index(jen.Lit(1).Op("<<").Lit(30)).Parens(r.Array.Type.sysParamGoPlainType())).
				// (unsafe.Pointer(ret))
				Parens(jenUnsafePointer().Call(jen.Id(cVarName))).
				// [:retLen:retLen]
				Index(jen.Op(":").Id(lenVarName).Op(":").Id(lenVarName))
		})
	}

	return jen.Id(retVarName)
}

func (r *ReturnValue) sysParamGoType() *jen.Statement {
	if r.Type != nil {
		return jen.Add(r.Type.sysParamGoType(false))
	}

	if r.Array != nil {
		return jen.Add(r.Array.sysParamGoType())
	}

	panic("ReturnValue is not a type or an array")
}
