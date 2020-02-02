package generate

import (
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
	a.Type.init(ns)

	//// Some array's Type has a Name but no CType.
	//// In all observed cases this is an integer type, usually 'guint8'.
	//// For these cases use the Type's Name as the CType.
	//_, isInteger := numberCTypeMap[a.Type.Name]
	//if isInteger && a.Type.CType == "" {
	//	a.Type.CType = a.Type.Name
	//}
	//
	//a.Type.init(ns)
}

func (a *Array) isSupported(in bool, out bool) bool {
	if strings.HasSuffix(a.CType, "char*") {
		// a simple string
		return true
	}

	if strings.HasSuffix(a.CType, "char***") && a.Type.isString() {
		// an array of strings
		return true
	}

	return false
}
