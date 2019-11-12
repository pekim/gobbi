package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strconv"
)

type Type struct {
	Name  string `xml:"name,attr"`
	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	Namespace *Namespace
}

func (t *Type) jenGoType() (*jen.Statement, string) {
	if t == nil {
		return nil, "missing Type"
	}
	if t.CType == "" {
		return nil, "missing Type.CType"
	}

	switch t.CType {
	case "gint8":
		return jen.Int8(), ""
	case "gshort", "gint16":
		return jen.Int16(), ""
	case "int", "gint", "gint32":
		return jen.Int32(), ""
	case "glong", "gint64":
		return jen.Int64(), ""

	case "guchar", "guint8":
		return jen.Uint8(), ""
	case "gushort", "guint16":
		return jen.Uint16(), ""
	case "guint", "guint32":
		return jen.Uint32(), ""
	case "gulong", "guint64":
		return jen.Uint64(), ""
	}

	return nil, fmt.Sprintf("No Go type for '%s'\n", t.CType)
}

func (t *Type) jenValue(valueText string) (*jen.Statement, string) {
	if t == nil {
		return nil, "missing Type"
	}

	var lit *jen.Statement
	var value int
	var err error

	switch t.CType {
	case "gint8":
		value, err = strconv.Atoi(valueText)
		lit = jen.Lit(int8(value))
	case "gshort", "gint16":
		value, err = strconv.Atoi(valueText)
		lit = jen.Lit(int16(value))
	case "int", "gint", "gint32":
		value, err = strconv.Atoi(valueText)
		lit = jen.Lit(int32(value))
	case "glong", "gint64":
		value, err = strconv.Atoi(valueText)
		lit = jen.Lit(int64(value))

	case "guchar", "guint8":
		value, err = strconv.Atoi(valueText)
		lit = jen.Lit(uint8(value))
	case "gushort", "guint16":
		value, err = strconv.Atoi(valueText)
		lit = jen.Lit(uint8(value))
	case "guint", "guint32":
		value, err = strconv.Atoi(valueText)
		lit = jen.Lit(uint8(value))
	case "gulong", "guint64":
		value, err = strconv.Atoi(valueText)
		lit = jen.Lit(uint8(value))
	default:
		return nil, fmt.Sprintf("Cannot generate literal value for '%s'\n", t.CType)
	}

	if err != nil {
		return nil, err.Error()
	}
	return lit, ""
}
