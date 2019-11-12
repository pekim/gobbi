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

	switch t.CType {
	case "gint8", "gshort", "gint16", "int", "gint", "gint32", "glong", "gint64":
		intValue, err := strconv.ParseInt(stringValue, 10, 64)
		if err != nil {
			return nil, err
		}

		switch t.CType {
		case "gint8":
			return jen.Lit(int8(intValue)), nil
		case "gshort", "gint16":
			return jen.Lit(int16(intValue)), nil
		case "int", "gint", "gint32":
			return jen.Lit(int32(intValue)), nil
		case "glong", "gint64":
			return jen.Lit(int64(intValue)), nil
		}
	case "guchar", "guint8", "gushort", "guint16", "guint", "guint32", "gulong", "guint64":
		uintValue, err := strconv.ParseUint(stringValue, 10, 64)
		if err != nil {
			return nil, err
		}

		switch t.CType {
		case "guchar", "guint8":
			return jen.Lit(uint8(uintValue)), nil
		case "gushort", "guint16":
			return jen.Lit(uint16(uintValue)), nil
		case "guint", "guint32":
			return jen.Lit(uint32(uintValue)), nil
		case "gulong", "guint64":
			return jen.Lit(uint64(uintValue)), nil
		}
	case "gdouble":
		value, err := strconv.ParseFloat(stringValue, 64)
		return jen.Lit(value), err
	case "gboolean":
		return jen.Lit(stringValue == "true"), nil
	case "gchar*":
		return jen.Lit(stringValue), nil
	}

	return nil, fmt.Errorf("Cannot generate literal value for '%s'\n", t.CType)
}
