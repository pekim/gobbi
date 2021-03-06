package generate

import (
	"fmt"
	"github.com/pekim/jennifer/jen"
)

type Constant struct {
	Namespace *Namespace

	Name      string `xml:"name,attr"`
	Blacklist bool   `xml:"blacklist,attr"`
	Value     string `xml:"value,attr"`
	Version   string `xml:"version,attr"`
	CType     string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	Doc       *Doc   `xml:"doc"`
	Type      *Type  `xml:"type"`

	goTypeName string
}

func (c *Constant) init(ns *Namespace) {
	c.Namespace = ns

	switch c.Type.Name {
	case "gboolean":
		c.goTypeName = "bool"
	case "utf8":
		c.goTypeName = "string"
	default:
		if goTypeName, found := numberCTypeMap[c.Type.Name]; found {
			c.goTypeName = goTypeName
		}
	}
}

func (c *Constant) version() string {
	return c.Version
}

func (c *Constant) mergeAddenda(addenda *Constant) {
	c.Blacklist = addenda.Blacklist
}

func (c *Constant) blacklisted() (bool, string) {
	return c.Blacklist, c.Name
}

func (c *Constant) supported() (supported bool, reason string) {
	if c.goTypeName != "" {
		return true, ""
	} else {
		return false, fmt.Sprintf("type %s for %s", c.Type.Name, c.Name)
	}
}

func (c *Constant) generate(g *jen.Group, version *Version) {
	if !supportedByVersion(c, version) {
		return
	}

	if c.Type.Name == "gboolean" {
		g.
			Const().
			Id(c.Name).
			Id(c.goTypeName).
			Op("=").
			Lit(c.Value == "true").
			Commentf("C.%s", c.CType)

		return
	}

	g.
		Const().
		Id(c.Name).
		Id(c.goTypeName).
		Op("=").
		Qual("C", c.CType)
}
