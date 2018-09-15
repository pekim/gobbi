package generate

import (
	"strings"
)

type Type struct {
	Namespace *Namespace

	Name  string `xml:"name,attr"`
	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	generator     TypeGenerator
	goType        string
	cTypeName     string // the C type, such as "gchar"
	indirectLevel int    // the level of pointer indirection
}

func (t *Type) init(ns *Namespace) {
	t.Namespace = ns

	cTypeParts := strings.Split(t.CType, " ")
	cType := cTypeParts[len(cTypeParts)-1]

	t.cTypeName = strings.TrimRight(cType, "*")
	t.indirectLevel = len(cType) - len(t.cTypeName)

	t.initTypeSpecific()
}

// initTypeSpecific initialises field that are specific to the type of the Type.
//
// fields set
//		goType
//		generator
func (t *Type) initTypeSpecific() {
	goType, isInteger := integerCTypeMap[t.CType]
	if isInteger {
		t.goType = goType
		t.generator = TypeGeneratorIntegerNew(t)
		return
	}

	if t.Name == "utf8" || t.Name == "filename" {
		t.goType = "string"
		t.generator = TypeGeneratorStringNew(t)
		return
	}

	record, found := t.Namespace.recordForName(t.Name)

	if !found || t.Name != "Error" {
		return
	}
	t.goType = t.Name
	t.generator = TypeGeneratorRecordNew(t, record)
}
