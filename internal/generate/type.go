package generate

import (
	"errors"
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

func (t *Type) jenValue(stringValue string) (*jen.Statement, error) {
	if t == nil {
		return nil, errors.New("missing Type")
	}

	var lit *jen.Statement
	intValue, err := strconv.Atoi(stringValue)

	switch t.CType {
	case "gint8":
		lit = jen.Lit(int8(intValue))
	case "gshort", "gint16":
		lit = jen.Lit(int16(intValue))
	case "int", "gint", "gint32":
		lit = jen.Lit(int32(intValue))
	case "glong", "gint64":
		lit = jen.Lit(int64(intValue))

	case "guchar", "guint8":
		lit = jen.Lit(uint8(intValue))
	case "gushort", "guint16":
		lit = jen.Lit(uint8(intValue))
	case "guint", "guint32":
		lit = jen.Lit(uint8(intValue))
	case "gulong", "guint64":
		lit = jen.Lit(uint8(intValue))
	default:
		return nil, fmt.Errorf("Cannot generate literal value for '%s'\n", t.CType)
	}

	if err != nil {
		return nil, err
	}
	return lit, nil
}
