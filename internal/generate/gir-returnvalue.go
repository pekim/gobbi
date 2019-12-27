package generate

import (
	"github.com/dave/jennifer/jen"
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

	if r.Array != nil {
		return false, "has array return"
	}

	if r.Type.isCallback() {
		return false, "has callback return "
	}

	return true, ""
}

func (r *ReturnValue) generateSysGoValue(cVarName string) *jen.Statement {
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

	return jen.Parens(r.Type.sysParamGoType()).Parens(jen.Id("ret"))
}
