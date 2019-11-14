package generate

type Parameter struct {
	Namespace *Namespace

	Name              string `xml:"name,attr"`
	Direction         string `xml:"direction,attr"`
	TransferOwnership string `xml:"transfer-ownership,attr"`
	Nullable          bool   `xml:"nullable,attr"`
	AllowNone         bool   `xml:"allow-none,attr"`
	//Doc               *Doc      `xml:"doc"`
	Type *Type `xml:"type"`
	//Array             *Array    `xml:"array"`
	Varargs *struct{} `xml:"varargs"`

	goVarName string
}

func (p *Parameter) init(ns *Namespace) {
	p.Namespace = ns
	p.goVarName = makeExportedGoName(p.Name)

	if p.Type != nil {
		//// If the parameter's type is an alias, replace it with the aliased type
		//alias, found := ns.aliasForName(p.Type.Name)
		//if found {
		//	p.initAlias(ns, alias)
		//}

		p.Type.init(ns)
	}

}
