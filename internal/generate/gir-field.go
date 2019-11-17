package generate

type Field struct {
	Name     string `xml:"name,attr"`
	Writable int    `xml:"writable,attr"`
	Bits     int    `xml:"bits,attr"`
	Private  bool   `xml:"private,attr"`
	Doc      *Doc   `xml:"doc"`
	Type     *Type  `xml:"type"`

	goVarName string
	namespace *Namespace
}

func (f *Field) init(ns *Namespace) {
	f.namespace = ns
	f.goVarName = makeExportedGoName(f.Name)

	if f.Type != nil {
		f.Type.init(ns)
	}
}

func (f Field) generate(g *group) {
	if f.Private {
		return
	}

	goType, err := f.Type.jenGoType()
	if err != nil {
		g.unsupported(f.Name, err.Error())
		return
	}

	g.Id(f.goVarName).Add(goType)
}
