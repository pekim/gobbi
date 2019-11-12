package generate

type Constant struct {
	Name string `xml:"name,attr"`
	//Blacklist bool   `xml:"blacklist,attr"`
	Value   string `xml:"value,attr"`
	Version string `xml:"version,attr"`
	CType   string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	//Doc       *Doc   `xml:"doc"`
	Type *Type `xml:"type"`

	Namespace *Namespace
}

func (c *Constant) init(ns *Namespace) {
	c.Namespace = ns
}

func (c Constant) generate(f *file) {
	goName := makeExportedGoName(c.Name)

	value, err := c.Type.jenValue(c.Value)
	if err != nil {
		f.unsupported(c.Name, err.Error())
		return
	}

	f.docForC(goName, c.Name)
	f.
		Const().
		Id(goName).
		Op("=").
		Add(value)
	f.Line()
}
