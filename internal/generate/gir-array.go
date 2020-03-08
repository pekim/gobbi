package generate

import (
	"github.com/dave/jennifer/jen"
)

type Array struct {
	Type           *Type  `xml:"type"`
	CType          string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	FixedSize      *int   `xml:"fixed-size,attr"`
	Length         *int   `xml:"length,attr"`
	ZeroTerminated bool   `xml:"zero-terminated,attr"`

	namespace   *Namespace
	lengthParam *Parameter
	cType       cType
}

func (a *Array) init(ns *Namespace, params Parameters) {
	if a == nil {
		return
	}

	a.namespace = ns
	a.Type.init(ns)
	a.cType = parseCtype(a.CType)

	if params != nil && a.Length != nil {
		a.lengthParam = params[*a.Length]
	}
}

func (a *Array) isSupported() (bool, string) {
	if a.Type == nil {
		return false, "array has no type"
	}

	return true, ""
}

func (a *Array) sysParamGoType() *jen.Statement {
	return jen.Index().Add(a.Type.sysParamGoPlainType())
}
