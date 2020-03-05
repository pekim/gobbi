package generate

import "C"
import (
	"fmt"
)

type Parameter struct {
	context   *context
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

func (p *Parameter) init(ns *Namespace, parentContext *context) {
	p.context = newContext(parentContext, "Parameter", p.Name)
	p.namespace = ns
	p.Type.init(ns)
	p.Array.init(ns)
	p.goVarName = makeUnexportedGoName(p.Name)
}

func (p Parameter) isType() bool {
	return p.Type != nil
}

func (p Parameter) isArray() bool {
	return p.Array != nil
}

func (p Parameter) isVarargs() bool {
	return p.Varargs != nil
}

func (p Parameter) isSupported() (bool, string) {
	if !(p.isType() || p.isArray() || p.isVarargs()) {
		return false, "not a Type, Array, or Varargs"
	}

	if p.isType() && p.Type.isCallback() {
		return false, "is callback"
	}

	if p.isType() && p.Type.isLongDouble() {
		return false, "is long double"
	}

	if p.isType() && p.Type.cType.indirectionCount > 1 {
		return false, "is non array with indirect count > 1	"
	}

	if p.isArray() && p.Array.cType.indirectionCount == 0 {
		if !p.Array.Type.isString() {
			return false, fmt.Sprintf("is array parameter with indirection of %d", p.Array.cType.indirectionCount)
		}
	}

	if p.isArray() {
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
	return (p.isType() && p.Type.isVaList()) || p.isVarargs()
}
