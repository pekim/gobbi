package generate

import (
	"github.com/dave/jennifer/jen"
	"strings"
)

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
	if a.Type.isString() && strings.HasSuffix(a.CType, "***") {
		// special case for "*[]string"
		return jen.
			Op("*").
			Index().
			String()
	}

	return jen.
		Index().
		Add(a.Type.sysParamGoType())
}
