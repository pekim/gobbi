package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

var simpleSysParamGoTypes = map[string]*jen.Statement{
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
}

type Type struct {
	Name  string `xml:"name,attr"`
	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	namespace        *Namespace
	foreignNamespace *Namespace
	foreignName      string

	isConst          bool
	cType            string
	stars            string
	indirectionCount int
}

func (t *Type) init(ns *Namespace) {
	if t == nil {
		return
	}

	t.namespace = ns
	t.parseCtype()
}

func (t *Type) parseCtype() {
	parts := cTypeRegex.FindStringSubmatch(t.CType)
	t.isConst = parts[1] != ""
	t.cType = parts[2]
	t.stars = parts[3]
	t.indirectionCount = len(t.stars)
	if t.cType == "" {
		panic(fmt.Sprintf("Failed to parse type ; '%s'", t.CType))
	}
}

func (t *Type) sysParamGoType() *jen.Statement {
	if t.Name == "utf8" || t.Name == "filename" {
		return jen.String()
	}

	if simpleGoType, ok := simpleSysParamGoTypes[t.cType]; ok {
		return jen.Op(t.stars).Add(simpleGoType)
	}
	//fmt.Printf("type '%s' : '%s'\n", t.CType, t.Name)

	return jen.Qual("github.com/pekim/gobbi/lib/internal/c", "UndefinedParamType")
}
