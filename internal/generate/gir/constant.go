package gir

import (
	"fmt"

	"github.com/dave/jennifer/jen"
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
}

func (c *Constant) init(ns *Namespace) {
	c.Namespace = ns
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
	switch c.Type.Name {
	case "gint", "utf8":
		return true, ""
	default:
		return false, fmt.Sprintf("type %s for %s", c.Type.Name, c.Name)
	}
}

func (c *Constant) generate(g *jen.Group, version *Version) {
	var goTypeName string
	switch c.Type.Name {
	case "gint":
		goTypeName = "int"
	case "utf8":
		goTypeName = "string"
	}

	g.
		Const().
		Id(c.Name).
		Id(goTypeName).
		Op("=").
		Qual("C", c.CType)
}
