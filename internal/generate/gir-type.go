package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"regexp"
)

var cTypeRegex = regexp.MustCompile(" *(const |volatile )* *([a-zA-Z0-9 ]+) *(\\**)? *")

var simpleSysParamGoTypes = map[string]*jen.Statement{
	"gchar":         jen.Int8(),
	"guchar":        jen.Uint8(),
	"gunichar":      jen.Rune(),
	"gunichar2":     jen.Uint16(),
	"int":           jen.Int(),
	"gint":          jen.Int(),
	"gint8":         jen.Int8(),
	"gint16":        jen.Int16(),
	"gint32":        jen.Int32(),
	"gint64":        jen.Int64(),
	"guint":         jen.Uint(),
	"guint8":        jen.Uint8(),
	"guint16":       jen.Uint16(),
	"guint32":       jen.Uint32(),
	"guint64":       jen.Uint64(),
	"glong":         jen.Int64(),
	"gulong":        jen.Uint64(),
	"gloat":         jen.Float32(),
	"gfloat":        jen.Float32(),
	"double":        jen.Float64(),
	"long double":   jen.Float64(), // not ideal, but Go has nothing better
	"gdouble":       jen.Float64(),
	"gsize":         jen.Uint64(),
	"gssize":        jen.Uint64(),
	"gboolean":      jen.Bool(),
	"gpointer":      jen.Qual("unsafe", "Pointer"),
	"gconstpointer": jen.Qual("unsafe", "Pointer"),
	"goffset":       jen.Int64(),
}

type Type struct {
	Name  string `xml:"name,attr"`
	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	namespace        *Namespace
	foreignNamespace *Namespace
	foreignName      string

	cType             string
	cStars            string
	cIndirectionCount int
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
	t.parseCtype()
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

func (t *Type) parseCtype() {
	parts := cTypeRegex.FindStringSubmatch(t.CType)
	t.cType = parts[2]
	t.cStars = parts[3]
	t.cIndirectionCount = len(t.cStars)
	if t.cType == "" {
		panic(fmt.Sprintf("Failed to parse type ; '%s'", t.CType))
	}
}

func (t *Type) sysParamGoType() *jen.Statement {
	if t.Name == "utf8" || t.Name == "filename" {
		return jen.String()
	}

	if t.cType == "void" && t.cIndirectionCount == 1 {
		return jen.Qual("unsafe", "Pointer")
	}

	if simpleGoType, ok := simpleSysParamGoTypes[t.cType]; ok {
		return jen.Op(t.cStars).Add(simpleGoType)
	}

	if t.isAlias() ||
		t.isBitfield() ||
		t.isEnumeration() ||
		t.isClass() ||
		t.isRecord() ||
		t.isInterface() {

		if t.isQualifiedName() {
			return jen.Op(t.cStars).Qual(t.foreignNamespace.goFullSysPackageName, t.foreignName)
		}
		return jen.Op(t.cStars).Id(t.Name)
	}

	fmt.Printf("type : %s : '%s' : '%s'\n", t.namespace.Name, t.CType, t.Name)
	return jen.Qual("github.com/pekim/gobbi/lib/internal/c", "UndefinedParamType")
}

func (t *Type) isQualifiedName() bool {
	return t.foreignNamespace != nil && t.foreignNamespace != t.namespace
}
