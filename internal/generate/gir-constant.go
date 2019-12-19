package generate

import "github.com/blang/semver"

type Constant struct {
	Name      string `xml:"name,attr"`
	Blacklist bool   `xml:"blacklist,attr"`
	Value     string `xml:"value,attr"`
	Version   string `xml:"version,attr"`
	CType     string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	//Doc       *Doc   `xml:"doc"`
	//Type      *Type  `xml:"type"`

	namespace  *Namespace
	version    semver.Version
	goTypeName string
}

func (c *Constant) init(ns *Namespace) {
	c.namespace = ns

	c.version = versionNew(c.Version)
	c.namespace.versions.add(c.version)

	//
	//switch c.Type.Name {
	//case "gboolean":
	//	c.goTypeName = "bool"
	//case "utf8":
	//	c.goTypeName = "string"
	//default:
	//	if goTypeName, found := numberCTypeMap[c.Type.Name]; found {
	//		c.goTypeName = goTypeName
	//	}
	//}
}
