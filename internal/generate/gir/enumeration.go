package gir

type Enumeration struct {
	Namespace *Namespace

	Name         string    `xml:"name,attr"`
	GoTypeName   string    `xml:"goname,attr"`
	Version      string    `xml:"version,attr"`
	CType        string    `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	GlibTypeName string    `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GlibGetType  string    `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	Doc          *Doc      `xml:"doc"`
	Members      []*Member `xml:"member"`
}

func (e *Enumeration) fixup(ns *Namespace) {
	e.Namespace = ns

	for _, member := range e.Members {
		member.Namespace = ns
	}
}

type Member struct {
	Namespace *Namespace

	Name        string `xml:"name,attr"`
	Value       int    `xml:"value,attr"`
	CIdentifier string `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	GlibNick    string `xml:"http://www.gtk.org/introspection/glib/1.0 nick,attr"`
	Doc         *Doc   `xml:"doc"`
}
