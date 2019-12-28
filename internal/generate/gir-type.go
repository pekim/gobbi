package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

var simpleSysParamGoTypes = map[string]*jen.Statement{
	"char":            jen.Int8(),
	"gchar":           jen.Int8(),
	"guchar":          jen.Uint8(),
	"gunichar":        jen.Rune(),
	"gunichar2":       jen.Uint16(),
	"gatomicrefcount": jen.Int(),
	"grefcount":       jen.Int(),
	"pid_t":           jen.Int(),
	"int":             jen.Int(),
	"gint":            jen.Int(),
	"gint8":           jen.Int8(),
	"gint16":          jen.Int16(),
	"gint32":          jen.Int32(),
	"gint64":          jen.Int64(),
	"uid_t":           jen.Uint(),
	"guint":           jen.Uint(),
	"guint8":          jen.Uint8(),
	"guint16":         jen.Uint16(),
	"guint32":         jen.Uint32(),
	"guint64":         jen.Uint64(),
	"time_t":          jen.Int64(),
	"glong":           jen.Int64(),
	"gulong":          jen.Uint64(),
	"gloat":           jen.Float32(),
	"gfloat":          jen.Float32(),
	"double":          jen.Float64(),
	"gdouble":         jen.Float64(),
	"gsize":           jen.Uint64(),
	"gssize":          jen.Uint64(),
	"gboolean":        jen.Bool(),
	"gpointer":        jenUnsafePointer(),
	"gconstpointer":   jenUnsafePointer(),
	"goffset":         jen.Int64(),
	"cairo_format_t":  jen.Int(),
}

type Type struct {
	Name  string `xml:"name,attr"`
	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	namespace        *Namespace
	foreignNamespace *Namespace
	foreignName      string
	cType            cType
}

func (t *Type) init(ns *Namespace) {
	if t == nil {
		return
	}

	t.namespace = ns

	// Strangely some GType references are not consistent.
	if t.Name == "GType" && t.namespace.Name != "GLib" {
		t.Name = "GLib.Type"
	}

	t.analyseName()

	if t.CType != "" {
		t.cType = parseCtype(t.CType)
	} else {
		t.cType = parseCtype(t.Name)
	}
}

func (t *Type) analyseName() {
	isForeign, foreignNamespace, foreignName := t.namespace.namespaces.analyseName(t.Name)

	if !isForeign {
		return
	}

	if foreignNamespace == t.namespace {
		// It's a qualified reference, but in this namespace.
		// So it's not really foreign.

		t.Name = foreignName
		return
	}

	t.foreignNamespace = foreignNamespace
	t.foreignName = foreignName

}

func (t *Type) sysParamGoPlainType() *jen.Statement {
	if t.isString() {
		return jen.String()
	}

	//if t.cType.typ == "void" && t.cType.indirectionCount == 1 {
	//	return jenUnsafePointer()
	//}
	if t.cType.typ == "void" {
		return nil
	}

	// pango specific
	if t.cType.typ == "FILE" && t.cType.indirectionCount == 1 {
		return jenUnsafePointer()
	}

	if simpleGoType, ok := simpleSysParamGoTypes[t.cType.typ]; ok {
		return jen.Add(simpleGoType)
	}

	if t.isAlias() {
		if t.isQualifiedName() {
			alias, _ := t.foreignNamespace.Aliases.byName(t.foreignName)
			return jen.Add(alias.Type.sysParamGoType(false))
		}

		alias, _ := t.namespace.Aliases.byName(t.Name)
		return jen.Add(alias.Type.sysParamGoType(false))
	}

	if t.isBitfield() || t.isEnumeration() {
		return jen.Int()
	}

	if t.isRecord() && t.cType.indirectionCount == 0 {
		if t.isQualifiedName() {
			return jen.Qual(t.foreignNamespace.goFullSysPackageName, t.foreignName)
		}
		return jen.Id(t.Name)
	}

	if t.isClass() ||
		t.isRecord() ||
		t.isInterface() ||
		t.isUnion() {

		return jenUnsafePointer()
	}

	panic(fmt.Sprintf("Unsupported type : %s %s (%s)", t.namespace.Name, t.Name, t.CType))
}

func (t *Type) sysParamGoType(decrementIndirectionCount bool) *jen.Statement {
	if t.isString() {
		stars := ""
		if t.cType.indirectionCount > 0 {
			stars = strings.Repeat("*", t.cType.indirectionCount-1)
		}
		return jen.Op(stars).String()
	}

	if t.cType.typ == "void" && t.cType.indirectionCount == 1 {
		return jenUnsafePointer()
	}

	// pango specific
	if t.cType.typ == "FILE" && t.cType.indirectionCount == 1 {
		return jenUnsafePointer()
	}

	stars := t.cType.stars
	if decrementIndirectionCount && t.cType.indirectionCount > 0 {
		stars = strings.Repeat("*", t.cType.indirectionCount-1)
	}

	if simpleGoType, ok := simpleSysParamGoTypes[t.cType.typ]; ok {
		return jen.Op(stars).Add(simpleGoType)
	}

	if t.isAlias() {
		if t.isQualifiedName() {
			alias, _ := t.foreignNamespace.Aliases.byName(t.foreignName)
			return jen.Op(stars).Add(alias.Type.sysParamGoType(false))
		}

		alias, _ := t.namespace.Aliases.byName(t.Name)
		return jen.Op(stars).Add(alias.Type.sysParamGoType(false))
	}

	if t.isBitfield() || t.isEnumeration() {
		return jen.Op(stars).Int()
	}

	if t.isRecord() && t.cType.indirectionCount == 0 {
		if t.isQualifiedName() {
			return jen.Qual(t.foreignNamespace.goFullSysPackageName, t.foreignName)
		}
		return jen.Op(stars).Id(t.Name)
	}

	if t.isClass() ||
		t.isRecord() ||
		t.isInterface() ||
		t.isUnion() {

		if t.cType.indirectionCount == 0 {
			return jenUnsafePointer()
		}

		stars := strings.Repeat("*", t.cType.indirectionCount-1)
		return jen.Op(stars).Add(jenUnsafePointer())
	}

	panic(fmt.Sprintf("Unsupported type : %s %s (%s)", t.namespace.Name, t.Name, t.CType))
}

func (t *Type) isQualifiedName() bool {
	return t.foreignNamespace != nil && t.foreignNamespace != t.namespace
}

func (t *Type) jenGoCType() *jen.Statement {
	return jen.Parens(jen.Op(t.cType.stars).Qual("C", t.cType.typ))
}
