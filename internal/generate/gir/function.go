package gir

type Function struct {
	Namespace *Namespace

	Name              string       `xml:"name,attr"`
	Whitelist         bool         `xml:"whitelist,attr"`
	Version           string       `xml:"version,attr"`
	CIdentifier       string       `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	Deprecated        int          `xml:"deprecated,attr"`
	DeprecatedVersion string       `xml:"deprecated-version,attr"`
	Doc               *Doc         `xml:"doc"`
	InstanceParameter *Parameter   `xml:"parameters>instance-parameter"`
	Parameters        []*Parameter `xml:"parameters>parameter"`
	ReturnValue       *ReturnValue `xml:"return-value"`
	Throws            int          `xml:"throws,attr"`
}

func (f *Function) fixup(ns *Namespace) {
	f.Namespace = ns

	for _, param := range f.Parameters {
		param.fixup(ns)
	}

	if f.ReturnValue != nil {
		f.ReturnValue.fixup(ns)
	}
}

type Functions []*Function

func (fs Functions) fixup(ns *Namespace) {
	for _, function := range fs {
		function.fixup(ns)
	}
}

type Parameter struct {
	Namespace *Namespace

	Name              string    `xml:"name,attr"`
	Direction         string    `xml:"direction,attr"`
	TransferOwnership string    `xml:"transfer-ownership,attr"`
	Doc               *Doc      `xml:"doc"`
	Type              *Type     `xml:"type"`
	Array             *Array    `xml:"array"`
	Varargs           *struct{} `xml:"varargs"`
}

func (p *Parameter) fixup(ns *Namespace) {
	p.Namespace = ns

	if p.Type != nil {
		p.Type.Namespace = ns
	}
}

type ReturnValue struct {
	Namespace *Namespace

	TransferOwnership string `xml:"transfer-ownership,attr"`
	Nullable          string `xml:"nullable,attr"`
	Doc               *Doc   `xml:"doc"`
	Type              *Type  `xml:"type"`
	Array             *Array `xml:"array"`
}

func (r *ReturnValue) fixup(ns *Namespace) {
	r.Namespace = ns

	if r.Type != nil {
		r.Type.Namespace = ns
	}
}
