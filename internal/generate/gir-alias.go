package generate

type Alias struct {
	Name string `xml:"name,attr"`
	//Blacklist bool   `xml:"blacklist,attr"`
	CType string `xml:"type,attr"`
	Type  *Type  `xml:"type"`
	Doc   *Doc   `xml:"doc"`

	namespace *Namespace
	goName    string
}

func (a *Alias) init(ns *Namespace) {
	a.namespace = ns
	a.goName = makeExportedGoName(a.Name)
	a.Type.init(ns)
}

func (a *Alias) generate(f *file) {
	goType, err := a.Type.jenGoType()
	if err != nil {
		f.unsupported(a.Name, err.Error())
		return
	}

	f.docForC(a.goName, a.Name)
	f.
		Type().
		Id(a.goName).
		Add(goType)
	f.Line()

}
