package generate

import (
	"strings"
)

type Type struct {
	Namespace *Namespace

	Name  string `xml:"name,attr"`
	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	generator     TypeGenerator
	qname         *QName
	cTypeName     string // the C type, such as "gchar"
	indirectLevel int    // the level of pointer indirection
}

func (t *Type) init(ns *Namespace) {
	t.Namespace = ns

	cTypeParts := strings.Split(t.CType, " ")
	cType := cTypeParts[len(cTypeParts)-1]

	t.cTypeName = strings.TrimRight(cType, "*")
	t.indirectLevel = len(cType) - len(t.cTypeName)

	t.qname, t.generator = t.qnameAndGenerator(t)
}

// goTypeAndGenerator determines the
//		package qualified Go name
//		generator
// that are specific to the type of the Type.
func (t *Type) qnameAndGenerator(targetType *Type) (*QName, TypeGenerator) {
	if t.Name == "gboolean" {
		return QNameNew(t.Namespace, "bool"), TypeGeneratorBooleanNew(targetType)
	}

	if t.Name == "utf8" || t.Name == "filename" {
		return QNameNew(t.Namespace, "string"), TypeGeneratorStringNew(targetType)
	}

	goType, isInteger := integerCTypeMap[t.cTypeName]
	if isInteger {
		return QNameNew(t.Namespace, goType), TypeGeneratorIntegerNew(targetType)
	}

	qname := QNameNew(t.Namespace, t.Name)

	alias, found := qname.ns.aliasForName(qname.name)
	if found {
		// Use a generator for the alias' Type rather than
		// this Type.
		alias.Type.init(t.Namespace)
		_, typeGenerator := alias.Type.qnameAndGenerator(t)
		return qname, typeGenerator
	}

	enum, found := qname.ns.bitfieldForName(qname.name)
	if found {
		if enum.goTypeName != "" {
			qname.name = enum.goTypeName
		}
		return qname, TypeGeneratorEnumerationNew(targetType, enum)
	}

	enum, found = qname.ns.enumForName(qname.name)
	if found {
		if enum.goTypeName != "" {
			qname.name = enum.goTypeName
		}
		return qname, TypeGeneratorEnumerationNew(targetType, enum)
	}

	record, found := qname.ns.recordForName(qname.name)
	if found {
		return qname, TypeGeneratorRecordNew(targetType, record)
	}

	class, found := qname.ns.classForName(qname.name)
	if found {
		return qname, TypeGeneratorRecordNew(targetType, class.Record)
	}

	return nil, nil
}
