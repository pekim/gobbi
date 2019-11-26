package generate

import (
	"errors"
	"fmt"
	"github.com/dave/jennifer/jen"
	"strconv"
	"strings"
)

var jenGoTypes = map[string]*jen.Statement{
	// signed
	"gchar":  jen.Int8(),
	"gint8":  jen.Int8(),
	"gshort": jen.Int16(),
	"gint16": jen.Int16(),
	"int":    jen.Int32(),
	"gint":   jen.Int32(),
	"gint32": jen.Int32(),
	"gssize": jen.Int32(),
	"glong":  jen.Int64(),
	"gint64": jen.Int64(),

	// unsigned
	"guchar":  jen.Uint8(),
	"guint8":  jen.Uint8(),
	"gushort": jen.Uint16(),
	"guint16": jen.Uint16(),
	"guint":   jen.Uint32(),
	"guint32": jen.Uint32(),
	"gulong":  jen.Uint64(),
	"guint64": jen.Uint64(),
	"gsize":   jen.Uint64(),

	// floats
	"gfloat":  jen.Float32(),
	"gdouble": jen.Float64(),

	// misc.
	"gboolean": jen.Bool(),
	"utf8":     jen.String(),
}

type Type struct {
	Name  string `xml:"name,attr"`
	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	namespace        *Namespace
	foreignNamespace *Namespace
	foreignName      string
}

func (t *Type) init(ns *Namespace) {
	t.namespace = ns
	t.analyseName()
}

func (t *Type) analyseName() {
	parts := strings.Split(t.Name, ".")

	if len(parts) == 1 {
		return
	}

	if len(parts) != 2 {
		panic(fmt.Sprintf("Unsupported Type name %s", t.Name))
	}

	ns, found := t.namespace.namespaces.byName(parts[0])
	if !found {
		panic(fmt.Sprintf("Namespace % referenced in Type name %s not found", parts[0], t.Name))
	}
	t.foreignNamespace = ns
	t.foreignName = parts[1]
}

func (t *Type) isQualifiedName() bool {
	return t.foreignNamespace != nil
}

func (t *Type) jenGoTypeForTypeName() (*jen.Statement, bool) {
	if t == nil {
		return nil, false
	}

	if t.isQualifiedName() {
		return nil, false
	}

	if jenType, found := jenGoTypes[t.Name]; found {
		return jenType, true
	}

	if t.namespace.haveType(t.Name) {
		return jen.Id(t.Name), true
	}

	return nil, false
}

func (t *Type) jenGoType() (*jen.Statement, error) {
	if t == nil {
		return nil, errors.New("missing Type")
	}

	if jenType, found := jenGoTypes[t.Name]; found {
		return jenType, nil
	}

	goType, ok := t.jenGoTypeForTypeName()
	if ok {
		return goType, nil
	}

	if generator, ok := t.namespace.outParameterGeneratorByName(t.Name); ok {
		return generator.generateDeclaration(), nil
	}

	return nil, fmt.Errorf("no Go type for '%s'", t.Name)
}

func (t *Type) jenValue(stringValue string) (*jen.Statement, error) {
	if t == nil {
		return nil, errors.New("missing Type")
	}

	switch t.Name {
	case "gchar", "gint8", "gshort", "gint16", "int", "gint", "gint32", "gssize", "glong", "gint64":
		return t.jenValueInt(stringValue)
	case "guchar", "guint8", "gushort", "guint16", "guint", "guint32", "gulong", "guint64":
		return t.jenValueUint(stringValue)
	case "gfloat":
		value, err := strconv.ParseFloat(stringValue, 32)
		return jen.Lit(float32(value)), err
	case "gdouble":
		value, err := strconv.ParseFloat(stringValue, 64)
		return jen.Lit(float64(value)), err
	case "gboolean":
		return jen.Lit(stringValue == "true"), nil
	case "utf8":
		return jen.Lit(stringValue), nil
	}

	return nil, fmt.Errorf("Cannot generate literal value for '%s'\n", t.Name)
}

func (t *Type) jenValueInt(stringValue string) (*jen.Statement, error) {
	value, err := strconv.ParseInt(stringValue, 10, 64)
	if err != nil {
		return nil, err
	}

	switch t.Name {
	case "gchar", "gint8":
		return jen.Lit(int8(value)), nil
	case "gshort", "gint16":
		return jen.Lit(int16(value)), nil
	case "int", "gint", "gint32", "gssize":
		return jen.Lit(int32(value)), nil
	case "glong", "gint64":
		return jen.Lit(int64(value)), nil
	}

	return nil, fmt.Errorf("Unknown type'%s'\n", t.Name)
}

func (t *Type) jenValueUint(stringValue string) (*jen.Statement, error) {
	value, err := strconv.ParseUint(stringValue, 10, 64)
	if err != nil {
		return nil, err
	}

	switch t.Name {
	case "guchar", "guint8":
		return jen.Lit(uint8(value)), nil
	case "gushort", "guint16":
		return jen.Lit(uint16(value)), nil
	case "guint", "guint32":
		return jen.Lit(uint32(value)), nil
	case "gulong", "guint64":
		return jen.Lit(uint64(value)), nil
	}

	return nil, fmt.Errorf("Unknown type'%s'\n", t.Name)
}

func (t *Type) argumentValueGetFunctionName() string {
	if getFunctionName, ok := argumentGetFunctionNames[t.Name]; ok {
		return getFunctionName
	}

	if generator, ok := t.namespace.outParameterGeneratorByName(t.Name); ok {
		return generator.argumentGetFunctionName()
	}

	panic(fmt.Sprintf("Cannot determine argumentGetFunctionName for %s", t.Name))
}

func (t *Type) resolvedType() *Type {
	alias, isAlias := t.namespace.Aliases.byName(t.Name)
	if isAlias {
		return alias.Type
	}
	return t
}

func (t *Type) argumentValueSetFunctionName() string {
	name := t.resolvedType().Name

	if setFunctionName, ok := argumentSetFunctionNames[name]; ok {
		return setFunctionName
	}

	if generator, ok := t.namespace.outParameterGeneratorByName(name); ok {
		return generator.argumentSetFunctionName()
	}

	panic(fmt.Sprintf("Cannot determine argumentValueSetFunctionName for %s", t.Name))
}

func (t *Type) createFromArgumentFunction() func(s *jen.Statement, arg *jen.Statement) {
	if generator, ok := t.namespace.outParameterGeneratorByName(t.Name); ok {
		return generator.createFromArgument
	}

	return nil
}

func (t *Type) isAlias() bool {
	_, found := t.namespace.Aliases.byName(t.Name)
	return found
}
