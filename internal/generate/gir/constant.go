package gir

type Constant struct {
	Namespace *Namespace

	Name    string `xml:"name,attr"`
	Value   string `xml:"value,attr"`
	Version string `xml:"version,attr"`
	CType   string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	Doc     *Doc   `xml:"doc"`
	Type    *Type  `xml:"type"`
}

func (c *Constant) init(ns *Namespace) {
	c.Namespace = ns
	c.Type.Namespace = ns
}
