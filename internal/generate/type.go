package generate

import (
	"strings"
)

type Type struct {
	Namespace *Namespace

	Name  string `xml:"name,attr"`
	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	cTypeName     string // the C type, such as "gchar"
	indirectLevel int    // the level of pointer indirection
}

func (t *Type) init(ns *Namespace) {
	t.Namespace = ns

	cTypeParts := strings.Split(t.CType, " ")
	cType := cTypeParts[len(cTypeParts)-1]

	t.cTypeName = strings.TrimRight(cType, "*")
	t.indirectLevel = len(cType) - len(t.cTypeName)
}
