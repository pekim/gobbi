package generate

type Constant struct {
	Name string `xml:"name,attr"`
	//Blacklist bool   `xml:"blacklist,attr"`
	Value   string `xml:"value,attr"`
	Version string `xml:"version,attr"`
	CType   string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	Doc     *Doc   `xml:"doc"`
	Type    *Type  `xml:"type"`

	namespace *Namespace
	goName    string
}

func (c *Constant) init(ns *Namespace) {
	c.namespace = ns

	// Unlike most generate Go names, for constants do not
	// transform using the makeExportedGoName function.
	// Instead keep the upper case snake names that are use for C.
	//
	// One reason is to avoid some name clashes, such as
	// SourceRemove in the glib package.
	c.goName = c.Name

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
