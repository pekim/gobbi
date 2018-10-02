package generate

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

type Parameters []*Parameter

func (pp Parameters) init(ns *Namespace) {
	for _, param := range pp {
		param.init(ns)
	}

	pp.fixupArgcArgv()
}

func (pp Parameters) fixupArgcArgv() {
	var previous *Parameter

	for i, p := range pp {
		if previous != nil && previous.Name == "argc" && p.Name == "argv" {
			pp.replaceArgcArgv(i - 1)
			break
		}

		previous = p
	}
}

func (pp Parameters) replaceArgcArgv(index int) {
	pp[index] = &Parameter{
		Name: "args",
		Type: &Type{
			Name: "argcargv",
		},
	}

	pp[index+1] = &Parameter{
		Type: &Type{
			Name: "ignore",
		},
	}
}

func (pp Parameters) generateFunctionDeclaration(g *jen.Group) {
	for _, p := range pp {
		p.generateFunctionDeclaration(g)
	}
}

func (pp Parameters) generateCVars(g *jen.Group) {
	for _, p := range pp {
		p.generateCVar(g)
	}
}

func (pp Parameters) generateOutputParamsReturnDeclaration(g *jen.Group) {
	for _, p := range pp {
		p.generateOutputParamReturnDeclaration(g)
	}
}

func (pp Parameters) generateOutputParamsReturns(g *jen.Group) {
	for _, p := range pp {
		p.generateOutputParamReturn(g)
	}
}

func (pp Parameters) generateOutputParamsGoVars(g *jen.Group) {
	for _, p := range pp {
		p.generateOutputParamGoVar(g)
	}
}

func (pp Parameters) generateCallArguments(g *jen.Group) {
	for _, p := range pp {
		p.generateCallArgument(g)
	}
}

func (pp Parameters) allSupported() (bool, string) {
	for _, p := range pp {
		if supported, reason := p.isSupported(); !supported {
			return supported, fmt.Sprintf("unsupported parameter %s : %s", p.Name, reason)
		}
	}

	return true, ""
}
