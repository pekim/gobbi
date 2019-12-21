package generate

type Parameter struct {
	Name      string `xml:"name,attr"`
	Direction string `xml:"direction,attr"`
	Argument
	AllowNone bool `xml:"allow-none,attr"`
	//Doc       *Doc      `xml:"doc"`
	Varargs *struct{} `xml:"varargs"`

	goVarName string
	namespace *Namespace

	lengthParam    *Parameter
	lengthForParam *Parameter
	argcParam      *Parameter
	argvParam      *Parameter
}

func (p *Parameter) init(ns *Namespace) {
	p.namespace = ns
	p.Array.init(ns)
	p.goVarName = makeUnexportedGoName(p.Name)
}
