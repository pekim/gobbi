package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"os"
	"strings"
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
	case "gboolean", "gint", "utf8":
		return true, ""
	default:
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
			Bool().
			Op("=").
			Lit(c.Value == "true").
			Commentf("C.%s", c.CType)

		return
	}

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

func (c *Constant) generateDocs(file *os.File) {
	doc := ""
	if c.Doc != nil {
		doc = strings.Replace(c.Doc.Text, "\n#", "\n&num;", -1)
	}

	_, err := file.WriteString(fmt.Sprintf(
		"## `%s`\n\n%s\n\nC - `%s`\n\n",
		c.Name,
		doc,
		c.CType,
	))
	if err != nil {
		panic(err)
	}
}
