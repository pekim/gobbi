package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type ReturnValue struct {
	Namespace *Namespace

	TransferOwnership string `xml:"transfer-ownership,attr"`
	Nullable          string `xml:"nullable,attr"`
	Doc               *Doc   `xml:"doc"`
	Type              *Type  `xml:"type"`
	Array             *Array `xml:"array"`

	returnType ReturnType
	goType     string
}

func (r *ReturnValue) init(ns *Namespace) {
	r.Namespace = ns

	if r.Type != nil {
		r.Type.init(ns)
		r.goType, r.returnType = returnType(r)
	}
}

func (r *ReturnValue) isSupported() (bool, string) {
	if r.returnType == nil {
		return false, "no return type"
	}

	if supported, reason := r.returnType.isSupported(); !supported {
		return false, fmt.Sprintf("return type : %s", reason)
	}

	return true, ""
}

func (r *ReturnValue) generateFunctionDeclaration(g *jen.Group) {
	r.returnType.generateFunctionDeclaration(g)
}

func (r *ReturnValue) generateCToGo(g *jen.Group, cVarName string) {
	r.returnType.generateCToGo(g, cVarName)
}
