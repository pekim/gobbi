package generate

import (
	"strings"
)

type Type struct {
	Namespace *Namespace

	Name  string `xml:"name,attr"`
	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	generator TypeGenerator
	qname     *QName
	//goType        string
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
	goType, isInteger := integerCTypeMap[t.CType]
	if isInteger {
		return QNameNew(t.Namespace, goType), TypeGeneratorIntegerNew(targetType)
	}

	if t.Name == "utf8" || t.Name == "filename" {
		return QNameNew(t.Namespace, "string"), TypeGeneratorStringNew(targetType)
	}

	qname := QNameNew(t.Namespace, t.Name)

	alias, found := qname.ns.aliasForName(qname.name)
	if found {
		// Use a generator for the alias' Type rather than
		// this Type.
		_, typeGenerator := alias.Type.qnameAndGenerator(t)
		return qname, typeGenerator
	}

	_, found = qname.ns.bitfieldForName(qname.name)
	if found {
		return qname, TypeGeneratorEnumerationNew(targetType)
	}

	_, found = qname.ns.enumForName(qname.name)
	if found {
		return qname, TypeGeneratorEnumerationNew(targetType)
	}

	record, found := qname.ns.recordForName(qname.name)
	if found {
		return qname, TypeGeneratorRecordNew(targetType, record)
	}

	return nil, nil
}

//func (t *Type) nameParts() (*Namespace, string) {
//	parts := strings.Split(t.Name, ".")
//
//	switch len(parts) {
//	case 1:
//		return t.Namespace, t.Name
//	case 2:
//		ns := t.Namespace.get(parts[0])
//		if ns == nil {
//			panic(fmt.Sprintf("Failed to find namespace for '%s'", t.Name))
//		}
//		return ns, parts[1]
//	default:
//		panic(fmt.Sprintf("Unsupported type name '%s'", t.Name))
//	}
//}
