package generate

import (
	"github.com/dave/jennifer/jen"
	"regexp"
)

var cTypeRegex = regexp.MustCompile(" *(const)? *([a-zA-Z0-9 ]+) *(\\**)? *")

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
	p.Type.init(ns)
	p.Array.init(ns)
	p.goVarName = makeUnexportedGoName(p.Name)
}

func (p *Parameter) isIn() bool {
	return p.Direction == "" || p.Direction == "in" || p.Direction == "inout"
}

func (p *Parameter) isOut() bool {
	return p.Direction == "out" || p.Direction == "inout"
}

func (p *Parameter) isInOut() bool {
	return p.Direction == "inout"
}

func (p *Parameter) sysParamGoType() *jen.Statement {
	if p.Type != nil {
		return p.Type.sysParamGoType()
	}

	return jen.Qual("github.com/pekim/gobbi/lib/internal/c", "UndefinedParamType")
}
