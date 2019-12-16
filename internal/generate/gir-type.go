package generate

import (
	"errors"
	"fmt"
	"github.com/dave/jennifer/jen"
	"strconv"
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
	"filename": jen.String(),
	"utf8":     jen.String(),
	"GType":    jen.Int64(),
}

type Type struct {
	Name  string `xml:"name,attr"`
	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	namespace        *Namespace
	foreignNamespace *Namespace
	foreignName      string
}

func (t *Type) String() string {
	if t == nil {
		return "nil"
	}

	return t.Name
}

func (t *Type) init(ns *Namespace) {
	t.namespace = ns
	t.analyseName()
}

func (t *Type) analyseName() {
	isForeign, foreignNamespace, foreignName := t.namespace.namespaces.analyseName(t.Name)
	if isForeign {
		t.foreignNamespace = foreignNamespace
		t.foreignName = foreignName
	}
}

func (t *Type) isQualifiedName() bool {
	return t.foreignNamespace != nil && t.foreignNamespace != t.namespace
}

func (t *Type) isValid() bool {
	return t != nil && t.Name != ""
}

func (t *Type) jenGoTypeForTypeName() (*jen.Statement, bool) {
	if t == nil {
		return nil, false
	}

	//if t.isQualifiedName() {
	//	if _, ok := t.foreignNamespace.Aliases.byName(t.foreignName); ok {
	//		return jen.Qual(t.foreignNamespace.goFullPackageName, t.foreignName), true
	//	}
	//}

	if jenType, found := jenGoTypes[t.Name]; found {
		return jenType, true
	}

	if t.namespace.haveType(t.Name) {
		return jen.Id(t.Name), true
	}

	return nil, false
}

func (t *Type) jenGoType() (*jen.Statement, error) {
	if !t.isValid() {
		return nil, errors.New("missing Type")
	}

	if jenType, found := jenGoTypes[t.Name]; found {
		return jenType, nil
	}

	if t.Name == "gpointer" {
		return jen.Qual("unsafe", "Pointer"), nil
	}

	goType, ok := t.jenGoTypeForTypeName()
	if ok {
		return goType, nil
	}

	if t.isClass() {
		if t.isQualifiedName() {
			class, _ := t.foreignNamespace.Classes.byName(t.foreignName)
			return jen.Op("*").Qual(t.foreignNamespace.goFullPackageName, class.goName), nil
		}

		class, _ := t.namespace.Classes.byName(t.Name)
		return jen.Op("*").Id(class.goName), nil
	}

	if t.isRecord() {
		if t.isQualifiedName() {
			record, _ := t.foreignNamespace.Records.byName(t.foreignName)
			return jen.Op("*").Qual(t.foreignNamespace.goFullPackageName, record.goName), nil
		}

		record, _ := t.namespace.Records.byName(t.Name)
		return jen.Op("*").Id(record.goName), nil
	}

	if t.isInterface() {
		if t.isQualifiedName() {
			iface, _ := t.foreignNamespace.Interfaces.byName(t.foreignName)
			return jen.Op("*").Qual(t.foreignNamespace.goFullPackageName, iface.goName), nil
		}

		iface, _ := t.namespace.Interfaces.byName(t.Name)
		return jen.Op("*").Id(iface.goName), nil
	}

	if t.isBitfield() {
		if t.isQualifiedName() {
			bitfield, _ := t.foreignNamespace.Bitfields.byName(t.foreignName)
			return jen.Qual(t.foreignNamespace.goFullPackageName, bitfield.goTypeName), nil
		}

		bitfield, _ := t.namespace.Bitfields.byName(t.Name)
		return jen.Id(bitfield.goTypeName), nil
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
	case "filename", "utf8":
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

	if t.Name == "gpointer" {
		return "Pointer"
	}

	if t.isBitfield() {
		return argumentGetFunctionNames["int"]
	}

	if t.isEnumeration() {
		return argumentGetFunctionNames["int"]
	}

	//if t.isQualifiedName() {
	//	if alias, ok := t.foreignNamespace.Aliases.byName(t.foreignName); ok {
	//		if getFunctionName, ok := argumentGetFunctionNames[alias.Type.Name]; ok {
	//			return getFunctionName
	//		}
	//	}
	//}

	if t.isRecord() || t.isInterface() {
		return "Pointer"
	}

	panic(fmt.Sprintf("Cannot determine argumentGetFunctionName for %s", t.Name))
}

func (t *Type) resolvedTypeX() *Type {
	if t == nil {
		return t
	}

	alias, isAlias := t.namespace.Aliases.byName(t.Name)
	if isAlias {
		return alias.Type
	}
	return t
}

func (t *Type) argumentValueSetFunctionName() string {
	if setFunctionName, ok := argumentSetFunctionNames[t.Name]; ok {
		return setFunctionName
	}

	if t.Name == "gpointer" {
		return "SetPointer"
	}

	if t.isBitfield() {
		return argumentSetFunctionNames["int"]
	}

	if t.isEnumeration() {
		return argumentSetFunctionNames["int"]
	}

	//if alias, ok := t.namespace.Aliases.byName(t.Name); ok {
	//	fmt.Println(alias)
	//	if setFunctionName, ok := argumentSetFunctionNames[alias.Type.Name]; ok {
	//		return setFunctionName
	//	}
	//}

	//if t.isQualifiedName() {
	//	if alias, ok := t.foreignNamespace.Aliases.byName(t.foreignName); ok {
	//		if setFunctionName, ok := argumentSetFunctionNames[alias.Type.Name]; ok {
	//			return setFunctionName
	//		}
	//	}
	//}

	if t.isRecord() || t.isInterface() {
		return "SetPointer"
	}

	panic(fmt.Sprintf("Cannot determine argumentValueSetFunctionName for %s %s", t.namespace.Name, t.Name))
}

func (t *Type) createFromOutArgument(g *jen.Group, argName *jen.Statement, argValue *jen.Statement) {
	if t.isClass() {
		if t.isQualifiedName() {
			class, _ := t.foreignNamespace.Classes.byName(t.foreignName)
			class.createFromArgument(g, t.foreignNamespace, argName, argValue)
		} else {
			class, _ := t.namespace.Classes.byName(t.Name)
			class.createFromArgument(g, nil, argName, argValue)
		}

		return
	}

	if t.isRecord() {
		if t.isQualifiedName() {
			record, _ := t.foreignNamespace.Records.byName(t.foreignName)
			record.createFromArgument(g, t.foreignNamespace, argName, argValue)
		} else {
			record, _ := t.namespace.Records.byName(t.Name)
			record.createFromArgument(g, nil, argName, argValue)
		}

		return
	}

	if t.isInterface() {
		if t.isQualifiedName() {
			iface, _ := t.foreignNamespace.Interfaces.byName(t.foreignName)
			iface.createFromArgument(g, t.foreignNamespace, argName, argValue)
		} else {
			iface, _ := t.namespace.Interfaces.byName(t.Name)
			iface.createFromArgument(g, nil, argName, argValue)
		}

		return
	}

	//if t.isQualifiedName() {
	//	if _, ok := t.foreignNamespace.Aliases.byName(t.foreignName); ok {
	//		g.
	//			Add(argName).
	//			Op(":=").
	//			Qual(t.foreignNamespace.goFullPackageName, t.foreignName).
	//			Parens(argValue)
	//
	//		return
	//	}
	//}

	g.
		Add(argName).
		Op(":=").
		Add(argValue)
}

func (t *Type) createFromInArgument(arg *jen.Statement) *jen.Statement {
	//if t.isRecord() {
	//	record, _ := t.namespace.Records.byName(t.Name)
	//	return record.createFromArgument
	//}

	//if t.isQualifiedName() {
	//	if alias, ok := t.foreignNamespace.Aliases.byName(t.foreignName); ok {
	//		return jen.
	//			Add(jenGoTypes[alias.Type.Name]).
	//			Parens(arg)
	//	}
	//}

	return arg
}

// generateOutArgValue generates a statement that transforms an out
// argument value to the Type's Go type.
func (t *Type) generateOutArgValue(g *jen.Group,
	argName *jen.Statement, argVar *jen.Statement, transferOwnership *bool,
) {
	var to *jen.Statement
	if transferOwnership != nil {
		to = jen.Lit(*transferOwnership)
	}

	argValue := argVar.
		Dot(t.argumentValueGetFunctionName()).
		Call(to)

	if t.isBitfield() {
		bitfield, _ := t.namespace.Bitfields.byName(t.Name)
		argValue = jen.
			Id(bitfield.goTypeName).
			Parens(argValue)
	}

	if t.isEnumeration() {
		argValue = jen.
			Id(t.Name).
			Parens(argValue)
	}

	t.createFromOutArgument(g, argName, argValue)
}

func (t *Type) jenNewFromNative() *jen.Statement {
	if t.isClass() {
		if t.isQualifiedName() {
			class, _ := t.foreignNamespace.Classes.byName(t.foreignName)
			return jen.Qual(t.foreignNamespace.goFullPackageName, class.newFromNativeName)
		} else {
			class, _ := t.namespace.Classes.byName(t.Name)
			return jen.Id(class.newFromNativeName)
		}

	}

	if t.isRecord() {
		if t.isQualifiedName() {
			record, _ := t.foreignNamespace.Records.byName(t.foreignName)
			return jen.Qual(t.foreignNamespace.goFullPackageName, record.newFromNativeName)
		} else {
			record, _ := t.namespace.Records.byName(t.Name)
			return jen.Id(record.newFromNativeName)
		}
	}

	if t.isInterface() {
		if t.isQualifiedName() {
			iface, _ := t.foreignNamespace.Interfaces.byName(t.foreignName)
			return jen.Qual(t.foreignNamespace.goFullPackageName, iface.newFromNativeName)
		} else {
			iface, _ := t.namespace.Interfaces.byName(t.Name)
			return jen.Id(iface.newFromNativeName)
		}
	}

	panic(fmt.Sprintf("Cannot provide newFromNative for %s", t.Name))
}

func (t *Type) isString() bool {
	return t.Name == "utf8" || t.Name == "filename"
}

func (t *Type) isAliasX() bool {
	_, found := t.namespace.Aliases.byName(t.Name)
	return found
}

func (t *Type) isBitfield() bool {
	_, found := t.namespace.Bitfields.byName(t.Name)
	return found
}

func (t *Type) isEnumeration() bool {
	_, found := t.namespace.Enumerations.byName(t.Name)
	return found
}

func (t *Type) isClass() bool {
	if t.isQualifiedName() {
		_, found := t.foreignNamespace.Classes.byName(t.foreignName)
		return found
	}

	_, found := t.namespace.Classes.byName(t.Name)
	return found
}

func (t *Type) isRecord() bool {
	if t.isClass() {
		return true
	}

	if t.isQualifiedName() {
		_, found := t.foreignNamespace.Records.byName(t.foreignName)
		return found
	}

	_, found := t.namespace.Records.byName(t.Name)
	return found
}

func (t *Type) isInterface() bool {
	if t.isQualifiedName() {
		_, found := t.foreignNamespace.Interfaces.byName(t.foreignName)
		return found
	}

	_, found := t.namespace.Interfaces.byName(t.Name)
	return found
}
