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
	t.goType, t.generator = t.goTypeAndGenerator(t)
}

func (t *Type) goTypeAndGenerator(targetType *Type) (string, TypeGenerator) {
	goType, isInteger := integerCTypeMap[t.CType]
	if isInteger {
		return goType, TypeGeneratorIntegerNew(targetType)
	}

	if t.Name == "utf8" || t.Name == "filename" {
		return "string", TypeGeneratorStringNew(targetType)
	}

	alias, found := t.Namespace.aliasForName(t.Name)
	if found {
		// Use a generator for the alias' Type rather than
		// this Type.
		_, typeGenerator := alias.Type.goTypeAndGenerator(t)
		return t.Name, typeGenerator
	}

	_, found = t.Namespace.enumForName(t.Name)
	if found {
		return t.Name, TypeGeneratorEnumerationNew(targetType)
	}

	record, found := t.Namespace.recordForName(t.Name)
	if found {
		return t.Name, TypeGeneratorRecordNew(targetType, record)
	}

	return "", nil
}
