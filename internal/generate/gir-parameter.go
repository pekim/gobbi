package generate

import "C"
import (
	"fmt"
)

type Parameter struct {
	Name      string `xml:"name,attr"`
	Direction string `xml:"direction,attr"`
	Argument
	AllowNone bool `xml:"allow-none,attr"`
	//Doc       *Doc      `xml:"doc"`
	Varargs *struct{} `xml:"varargs"`

	goVarName string
	namespace *Namespace

	lengthParam     *Parameter
	lengthParamN    int
	lengthForParam  *Parameter
	lengthForParamN int
}

func (p *Parameter) init(ns *Namespace) {
	p.namespace = ns
	p.Type.init(ns)
	p.Array.init(ns)
	p.goVarName = makeUnexportedGoName(p.Name)
}

func (p Parameter) isSupported() (bool, string) {
	if p.Type != nil && p.Type.isCallback() {
		return false, "is callback"
	}

	if p.Type != nil && p.Type.isLongDouble() {
		return false, "is long double"
	}

	if p.Array != nil && p.lengthParam == nil {
		return false, "is array parameter without length parameter"
	}

	if p.Array != nil && p.Array.cType.indirectionCount == 0 {
		if !p.Array.Type.isString() {
			return false, fmt.Sprintf("is array parameter with indirection of %d", p.Array.cType.indirectionCount)
		}
	}

	if p.Array != nil {
		if supported, reason := p.Array.isSupported(); !supported {
			return supported, reason
		}
	}

	return true, ""
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

func (p *Parameter) isVarargsOrValist() bool {
	return (p.Type != nil && p.Type.isVaList()) || p.Varargs != nil
}
