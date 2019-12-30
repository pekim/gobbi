package generate

import (
	"github.com/dave/jennifer/jen"
)

type Array struct {
	Type           *Type  `xml:"type"`
	CType          string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	FixedSize      *int   `xml:"fixed-size,attr"`
	Length         *int   `xml:"length,attr"`
	ZeroTerminated string `xml:"zero-terminated,attr"`

	namespace *Namespace
	cType     cType
}

func (a *Array) init(ns *Namespace) {
	if a == nil {
		return
	}

	a.namespace = ns
	a.Type.init(ns)
	a.cType = parseCtype(a.CType)
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
