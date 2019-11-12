package generate

import "fmt"

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
	fmt.Println(a.Name, a.Type.Name, a.Type.CType, a.CType)

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
