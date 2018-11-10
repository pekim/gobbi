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

	if t.CType == "GtkIconSize" && t.Name == "gint" {
		t.Name = "IconSize"
	}

	t.qname, t.generator = t.qnameAndGenerator(t)
}

// fullGoPackageName returns a package name if the type's qname is
// for another package. Or empty string if in the same package as
// the type.
func (t *Type) fullGoPackageName() string {
	if t.qname.namespace != t.Namespace {
		return t.qname.namespace.fullGoPackageName
	}
	return ""
}

// goTypeAndGenerator determines the
//		package qualified Go name
//		generator
// that are specific to the type of the Type.
func (t *Type) qnameAndGenerator(targetType *Type) (*QName, TypeGenerator) {
	if t.Name == "argcargv" {
		return QNameNew(t.Namespace, ""), TypeGeneratorArgcArgvNew(t)
	}

	if t.Name == "ignore" {
		return QNameNew(t.Namespace, ""), TypeGeneratorIgnoreNew(t)
	}

	if t.Name == "gboolean" {
		return QNameNew(t.Namespace, "bool"), TypeGeneratorBooleanNew(targetType)
	}

	if t.Name == "utf8" || t.Name == "filename" {
		return QNameNew(t.Namespace, "string"), TypeGeneratorStringNew(targetType)
	}

	if t.Name == "GType" {
		return QNameNew(t.Namespace, "GObject.Type"), TypeGeneratorIntegerNew(targetType)
	}

	goType, isInteger := integerCTypeMap[t.cTypeName]
	if isInteger {
		return QNameNew(t.Namespace, goType), TypeGeneratorIntegerNew(targetType)
	}

	qname := QNameNew(t.Namespace, t.Name)

	alias, found := qname.namespace.aliasForName(qname.name)
	if found {
		// Use a generator for the alias' Type rather than
		// this Type.
		alias.Type.init(t.Namespace)
		_, typeGenerator := alias.Type.qnameAndGenerator(t)
		return qname, typeGenerator
	}

	enum, found := qname.namespace.bitfieldForName(qname.name)
	if found {
		if enum.goTypeName != "" {
			qname.name = enum.goTypeName
		}
		return qname, TypeGeneratorEnumerationNew(targetType, enum)
	}

	enum, found = qname.namespace.enumForName(qname.name)
	if found {
		if enum.goTypeName != "" {
			qname.name = enum.goTypeName
		}
		return qname, TypeGeneratorEnumerationNew(targetType, enum)
	}

	record, found := qname.namespace.recordOrClassRecordForName(qname.name)
	if found {
		return qname, TypeGeneratorRecordNew(targetType, record)
	}

	iface, found := qname.namespace.interfaceForName(qname.name)
	if found {
		return qname, TypeGeneratorInterfaceNew(targetType, iface)
	}

	return nil, nil
}
