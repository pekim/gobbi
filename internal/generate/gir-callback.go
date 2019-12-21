package generate

type Callback struct {
	Name    string `xml:"name,attr"`
	Version string `xml:"version,attr"`
	CType   string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	namespace *Namespace
}

func (c *Callback) init(ns *Namespace) {
	c.namespace = ns
}
