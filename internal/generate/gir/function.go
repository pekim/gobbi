package gir

type Function struct {
	Namespace *Namespace

	Name              string       `xml:"name,attr"`
	Blacklist         bool         `xml:"blacklist,attr"`
	Version           string       `xml:"version,attr"`
	CIdentifier       string       `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	Deprecated        int          `xml:"deprecated,attr"`
	DeprecatedVersion string       `xml:"deprecated-version,attr"`
	Doc               *Doc         `xml:"doc"`
	InstanceParameter *Parameter   `xml:"parameters>instance-parameter"`
	Parameters        []*Parameter `xml:"parameters>parameter"`
	ReturnValue       *ReturnValue `xml:"return-value"`
	Throws            int          `xml:"throws,attr"`
	Introspectable    string       `xml:"introspectable,attr"`
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

func (f *Function) blacklisted() bool {
	return f.Blacklist
}

func (f *Function) introspectable() bool {
	return f.Introspectable != "0"
}

func (f *Function) version() string {
	return f.Version
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
