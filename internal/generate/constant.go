package generate

import "strings"

type Constant struct {
	Name string `xml:"name,attr"`
	//Blacklist bool   `xml:"blacklist,attr"`
	Value   string `xml:"value,attr"`
	Version string `xml:"version,attr"`
	CType   string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	//Doc       *Doc   `xml:"doc"`
	Type *Type `xml:"type"`

	namespace *Namespace
	goName    string
}

func (c *Constant) init(ns *Namespace) {
	c.namespace = ns

	if c.namespace.Name == "Gdk" && strings.HasPrefix(c.Name, "KEY_") {
		// Special case, to avoid duplicate names.
		// Do not transform.
		c.goName = c.Name
	} else if c.namespace.Name == "GLib" && strings.HasPrefix(c.Name, "CSET_") {
		// Special case, to avoid duplicate names.
		// Do not transform.
		c.goName = c.Name
	} else {
		c.goName = makeExportedGoName(c.Name)
	}

	c.Type.init(ns)
}

func (c *Constant) generate(f *file) {
	value, err := c.Type.jenValue(c.Value)
	if err != nil {
		f.unsupported(c.Name, err.Error())
		return
	}

	f.docForC(c.goName, c.Name)
	f.docVersion(c.Version)
	f.
		Const().
		Id(c.goName).
		Op("=").
		Add(value)
	f.Line()
}
