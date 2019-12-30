package generate

import (
	"github.com/dave/jennifer/jen"
)

var simpleSysParamGoTypes = map[string]*jen.Statement{
	"char":          jen.Int8(),
	"gchar":         jen.Int8(),
	"guchar":        jen.Uint8(),
	"gunichar":      jen.Rune(),
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
	"gfloat":        jen.Float32(),
	"double":        jen.Float64(),
	"gdouble":       jen.Float64(),
	"gsize":         jen.Uint64(),
	"gssize":        jen.Uint64(),
	"gboolean":      jen.Bool(),
	"gpointer":      jenUnsafePointer(),
	"gconstpointer": jenUnsafePointer(),
	"goffset":       jen.Int64(),
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

func (t *Type) isQualifiedName() bool {
	return t.foreignNamespace != nil && t.foreignNamespace != t.namespace
}

func (t *Type) jenGoCType() *jen.Statement {
	return jen.Parens(jen.Op(t.cType.stars).Qual("C", t.cType.typ))
}
