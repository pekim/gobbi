package gir

type Alias struct {
	Namespace *Namespace

	Name      string `xml:"name,attr"`
	Blacklist bool   `xml:"blacklist,attr"`
	CType     string `xml:"type,attr"`
	Type      Type   `xml:"type"`
	Doc       *Doc   `xml:"doc"`
}

func (a *Alias) fixup(ns *Namespace) {
	a.Namespace = ns
	a.Type.Namespace = ns
}
