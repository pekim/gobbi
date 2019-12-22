package generate

import "github.com/dave/jennifer/jen"

type Array struct {
	Type           *Type  `xml:"type"`
	CType          string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	FixedSize      *int   `xml:"fixed-size,attr"`
	Length         *int   `xml:"length,attr"`
	ZeroTerminated string `xml:"zero-terminated,attr"`

	namespace   *Namespace
	lengthParam *Parameter
}

func (a *Array) init(ns *Namespace) {
	if a == nil {
		return
	}

	a.namespace = ns
	a.Type.init(ns)
}

func (a *Array) sysParamGoType() *jen.Statement {
	return jen.
		Op("*"). // an extra level of indirection for array
		Add(a.Type.sysParamGoType())
}
