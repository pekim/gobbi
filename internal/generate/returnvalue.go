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
	goType    string
}

func (r *ReturnValue) init(ns *Namespace) {
	r.Namespace = ns

	if r.Type != nil {
		r.Type.Namespace = ns

		goType, isInteger := integerCTypeMap[r.Type.CType]
		if isInteger {
			r.goType = goType
			r.returnType = ReturnTypeIntegerNew(r)
		//} else if p.Type.Name == "utf8" {
		//	p.goType = "string"
		//	p.paramType = ParameterTypeStringNew(p)
		}
	}
}

func (r *ReturnValue) isSupported() (bool, string) {
	if r.returnType == nil {
		return false, "no return type"
	}

	if supported, reason := r.returnType.isSupported(); !supported {
		return false, fmt.Sprintf("return type : %s",reason)
	}

	return true, ""
}

func (r *ReturnValue) generateFunctionDeclaration(g *jen.Group) {
	r.returnType.generateFunctionDeclaration(g)
}

func (r *ReturnValue) generateCToGo(g *jen.Group, cVarName string) {
	r.returnType.generateCToGo(g, cVarName)
}
