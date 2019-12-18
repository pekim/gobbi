package generate

import "github.com/dave/jennifer/jen"

type Constant struct {
	Namespace *Namespace

	Name      string `xml:"name,attr"`
	Blacklist bool   `xml:"blacklist,attr"`
	Value     string `xml:"value,attr"`
	Version   string `xml:"version,attr"`
	CType     string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	//Doc       *Doc   `xml:"doc"`
	//Type      *Type  `xml:"type"`

	goTypeName string
}

func (c *Constant) init(ns *Namespace) {
	c.Namespace = ns
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

func (c Constant) generateSys(f *jen.File) {
	f.Commentf("constant : %s", c.Name)

}
