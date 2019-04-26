package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type ReturnValue struct {
	Namespace *Namespace

	TransferOwnership string `xml:"transfer-ownership,attr"`
	Nullable          bool   `xml:"nullable,attr"`
	Doc               *Doc   `xml:"doc"`
	Type              *Type  `xml:"type"`
	Array             *Array `xml:"array"`
}

func (r *ReturnValue) init(ns *Namespace) {
	r.Namespace = ns

	if r.Type != nil {
		r.Type.init(ns)
	}

	if r.Array != nil && r.Array.Type != nil {
		r.Array.init(ns)
	}
}

func (r *ReturnValue) mergeAddenda(addenda *ReturnValue) {
	if addenda.TransferOwnership != "" {
		r.TransferOwnership = addenda.TransferOwnership
	}

	if addenda.Type != nil {
		if addenda.Type.Name != "" {
			r.Type.Name = addenda.Type.Name
		}
		if addenda.Type.CType != "" {
			r.Type.CType = addenda.Type.CType
		}
	}
}

func (r *ReturnValue) isSupported() (bool, string) {
	if r.Array != nil {
		if r.Array.Type == nil {
			return false, "no type for array return"
		}
		if r.Array.Type.generator == nil {
			return false, fmt.Sprintf("no type generator for %s (%s) for array return",
				r.Array.Type.Name, r.Array.Type.CType)
		}

		if supported, reason := r.Array.Type.generator.isSupportedAsArrayReturnValue(); !supported {
			return false, fmt.Sprintf("array return type : %s", reason)
		}

		return true, ""
	}

	if r.Type == nil {
		return false, "no return type"
	}

	if r.Type.Name == "none" {
		return true, ""
	}

	if r.Type.generator == nil {
		return false, "no return generator"
	}

	if supported, reason := r.Type.generator.isSupportedAsReturnValue(); !supported {
		return false, fmt.Sprintf("return type : %s", reason)
	}

	return true, ""
}

func (r *ReturnValue) generateFunctionDeclaration(g *jen.Group) {
	if r.Type != nil && r.Type.Name == "none" {
		return
	}

	if r.Array != nil {
		r.Array.generateReturnDeclaration(g)
	} else {
		r.Type.generator.generateReturnFunctionDeclaration(g)
	}
}

func (r *ReturnValue) generateFunctionDeclarationCtype(g *jen.Group) {
	if r.Type.Name == "none" {
		return
	}

	r.Type.generator.generateReturnFunctionDeclarationCtype(g)
}

func (r *ReturnValue) generateCToGo(g *jen.Group, cVarName string, goVarName string) {
	if r.Type.Name == "none" {
		return
	}

	r.Type.generator.generateReturnCToGo(g, false, cVarName, goVarName,
		r.Type.fullGoPackageName(),
		r.TransferOwnership, r.Nullable)
}

func (r *ReturnValue) generateArrayCToGo(g *jen.Group, cVarName string, goVarName string) {
	r.Array.Type.generator.generateArrayReturnCToGo(g, false, cVarName, goVarName,
		r.Array.Type.fullGoPackageName(),
		r.TransferOwnership, r.Nullable)
}

func (r *ReturnValue) generatePopulateCallData(callData jen.Dict) {
	callData[jen.Id("Return")] = jen.
		Qual(callPkg, "Return").
		Values(jen.DictFunc(func(d jen.Dict) {
			if r.Type.generator != nil {
				d[jen.Id("Type")] = jen.Qual(callPkg, r.Type.generator.generateCallReturnType())
			} else {
				d[jen.Id("Type")] = jen.Qual(callPkg, "RT_VOID")
			}
		}))
}
