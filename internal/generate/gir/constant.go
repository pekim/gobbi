package gir

import (
	"fmt"
	"strconv"

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

func (c *Constant) generate(g *jen.Group, version *Version) {
	if !supportedByVersion(c, version) {
		return
	}

	if c.Blacklist {
		g.Commentf("Blacklisted constant : %s", c.Name)
		return
	}

	goTypeName := "int"
	goValue, err := strconv.Atoi(c.Value)
	if err != nil {
		panic(fmt.Errorf("Failed to parse int constant for %s : %s", c.Name, err))
	}

	g.
		Const().
		Id(c.Name).
		Id(goTypeName).
		Op("=").
		Lit(goValue)
}
