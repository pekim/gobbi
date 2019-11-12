package generate

type Alias struct {
	Name string `xml:"name,attr"`
	//Blacklist bool   `xml:"blacklist,attr"`
	CType string `xml:"type,attr"`
	Type  *Type  `xml:"type"`
	//Doc       *Doc   `xml:"doc"`

	namespace *Namespace
	goName    string
}

func (a *Alias) init(ns *Namespace) {
	a.namespace = ns
	a.goName = makeExportedGoName(a.Name)
}

func (a *Alias) generate(f *file) {
}
