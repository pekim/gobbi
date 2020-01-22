package generate

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
	"strconv"
)

type Constant struct {
	Name      string `xml:"name,attr"`
	Blacklist bool   `xml:"blacklist,attr"`
	Value     string `xml:"value,attr"`
	Version   string `xml:"version,attr"`
	CType     string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	//Doc       *Doc   `xml:"doc"`
	Type *Type `xml:"type"`

	namespace  *Namespace
	version    semver.Version
	goName     string
	goTypeName string
}

var constantIntBits = map[string]int{
	"gint8":   8,
	"gint16":  16,
	"gint32":  32,
	"gint64":  64,
	"guint8":  8,
	"guint16": 16,
	"guint32": 32,
	"guint64": 64,
}

func (c *Constant) init(ns *Namespace) {
	c.namespace = ns

	c.version = versionNew(c.Version)
	c.namespace.versions.add(c.version)
	c.Type.init(ns)

	// Unlike most generate Go names, for constants do not
	// transform using the makeExportedGoName function.
	// Instead keep the upper case snake names that are use for C.
	//
	// One reason is to avoid some name clashes, such as
	// SourceRemove in the glib package.
	c.goName = c.Name
}

func (c *Constant) generateLib(f *jen.File, version semver.Version) {
	if !generateEntityForVersion(version, c.version) {
		return
	}

	if c.Type.isQualifiedName() {
		f.Commentf("// UNSUPPORTED the alias %s is qualified, '%s'", c.Name, c.Type.Name)
		f.Line()
		return
	}

	typeName := c.Type.Name

	if c.Type.isAlias() {
		if c.Type.isQualifiedName() {
			alias, _ := c.Type.foreignNamespace.Aliases.byName(c.Type.foreignName)
			typeName = alias.Type.Name
		} else {
			alias, _ := c.Type.namespace.Aliases.byName(c.Type.Name)
			typeName = alias.Type.Name
		}
	}

	var value *jen.Statement
	switch typeName {

	case "gboolean":
		v, err := strconv.ParseBool(c.Value)
		if err != nil {
			panic(fmt.Sprintf("Cannot convert '%s' to a bool", c.Value))
		}
		value = jen.Lit(v)

	case "gdouble":
		v, err := strconv.ParseFloat(c.Value, 64)
		if err != nil {
			panic(fmt.Sprintf("Cannot convert '%s' to a float", c.Value))
		}
		value = jen.Lit(v)

	case "gchar":
		value = jen.LitRune([]rune(c.Value)[0])

	case "gint":
		v, err := strconv.Atoi(c.Value)
		if err != nil {
			panic(fmt.Sprintf("Cannot convert '%s' to an integer", c.Value))
		}
		value = jen.Lit(v)

	case "gint8", "gint16", "gint32", "gint64":
		bits := constantIntBits[typeName]
		v, err := strconv.ParseInt(c.Value, 10, bits)
		if err != nil {
			panic(fmt.Sprintf("Cannot convert '%s' to an integer", c.Value))
		}
		value = jen.Lit(v)

	case "guint":
		v, err := strconv.Atoi(c.Value)
		if err != nil {
			panic(fmt.Sprintf("Cannot convert '%s' to an integer", c.Value))
		}
		value = jen.Lit(v)

	case "guint8", "guint16", "guint32", "guint64":
		bits := constantIntBits[typeName]
		v, err := strconv.ParseUint(c.Value, 10, bits)
		if err != nil {
			panic(fmt.Sprintf("Cannot convert '%s' to an integer", c.Value))
		}
		value = jen.Lit(v)

	case "utf8":
		value = jen.Lit(c.Value)

	default:
		panic(fmt.Sprintf("Do not know how to create literal value for %s %s of type %s", c.namespace.Name, c.Name, c.Type.Name))
	}

	f.Commentf("%s is a representation of the C constant %s.", c.goName, c.CType)
	docVersion(f, c.Version)

	f.
		Const().
		Id(c.goName).
		Op("=").
		Add(value)

	f.Line()
}
