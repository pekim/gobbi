package generate

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

type Parameters []*Parameter

func (pp Parameters) init(ns *Namespace) {
	pp.fixupArgcArgv()
	pp.fixupStringLengthParams()
	pp.fixupFormatArgs()

	for _, param := range pp {
		param.init(ns)

		if param.Array != nil {
			if param.Array.Type != nil {
				param.Array.Type.init(ns)
				if param.Array.Type.generator == nil && param.Name == "files" {
					fmt.Println(param.Name, param.Array.Type.CType)
					panic("xxx")
				}
			}

			if param.Array.Length != nil {
				// Provide an array length param with a reference to its array param.
				paramIndex := *param.Array.Length
				pp[paramIndex].arrayLengthFor = param

				// And the reverse.
				param.Array.lengthParam = pp[paramIndex]
			}
		}
	}
}

// fixupStringLengthParams attempts to find parameters that are
// used to provide a length for a string parameter.
//
// Any such length parameter is updated with a reference to the
// string parameter.
func (pp Parameters) fixupStringLengthParams() {
	for i, param := range pp {
		if i == 0 {
			continue
		}

		if param.Name != "length" || param.Type == nil || param.Type.Name != "gssize" || param.Direction == "out" {
			continue
		}

		prevParam := pp[i-1]
		if prevParam.Type == nil || prevParam.Type.Name != "utf8" {
			continue
		}

		param.stringLengthFor = prevParam
	}
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
	p1 := pp[index]
	p1.Name = "args"
	p1.Type.Name = "argcargv"

	p2 := pp[index+1]
	p2.Array = nil
	p2.Type = &Type{
		Name: "ignore",
	}
}

func (pp Parameters) hasFormatArgs() bool {
	if len(pp) < 2 {
		return false
	}

	paramFormatString := pp[len(pp)-2]
	paramArgs := pp[len(pp)-1]

	if !strings.HasPrefix(paramFormatString.Name, "format") &&
		!strings.HasSuffix(paramFormatString.Name, "format") {
		return false
	}

	if paramFormatString.Type == nil || paramFormatString.Type.Name != "utf8" {
		return false
	}

	if paramArgs.Varargs == nil {
		return false
	}

	return true
}

func (pp Parameters) fixupFormatArgs() {
	if !pp.hasFormatArgs() {
		return
	}

	paramFormatString := pp[len(pp)-2]
	paramArgs := pp[len(pp)-1]

	paramFormatString.formatString = true
	paramArgs.formatArgs = true
}

func (pp Parameters) generateFunctionDeclaration(g *jen.Group) {
	for _, p := range pp {
		if p.ignore {
			continue
		}

		p.generateFunctionDeclaration(g)
	}
}

func (pp Parameters) generateFunctionDeclarationCtypes(g *jen.Group) {
	for _, p := range pp {
		if p.ignore {
			continue
		}

		p.generateFunctionDeclarationCtype(g)
	}
}

func (pp Parameters) generateCVars(g *jen.Group) {
	for _, p := range pp {
		if p.ignore {
			continue
		}

		p.generateCVar(g)
	}
}

func (pp Parameters) generateGoVars(g *jen.Group) {
	for _, p := range pp {
		if p.ignore {
			continue
		}

		p.generateGoVar(g)
	}
}

func (pp Parameters) generateOutputParamsReturnDeclaration(g *jen.Group) {
	for _, p := range pp {
		if p.ignore {
			continue
		}

		p.generateOutputParamReturnDeclaration(g)
	}
}

func (pp Parameters) generateOutputParamsReturns(g *jen.Group) {
	for _, p := range pp {
		if p.ignore {
			continue
		}

		p.generateOutputParamReturn(g)
	}
}

func (pp Parameters) generateOutputParamsGoVars(g *jen.Group) {
	for _, p := range pp {
		if p.ignore {
			continue
		}

		p.generateOutputParamGoVar(g)
	}
}

func (pp Parameters) generateCallArguments(g *jen.Group) {
	for _, p := range pp {
		if p.ignore {
			continue
		}

		p.generateCallArgument(g)
	}
}

func (pp Parameters) allSupported() (bool, string) {
	for _, p := range pp {
		if p.ignore {
			continue
		}

		if supported, reason := p.isSupported(); !supported {
			return supported, fmt.Sprintf("unsupported parameter %s : %s", p.Name, reason)
		}
	}

	return true, ""
}

func (pp Parameters) allSupportedC() (bool, string) {
	for _, p := range pp {
		if p.Direction == "out" {
			return false, fmt.Sprintf("unsupported parameter %s : direction is 'out'", p.Name)
		}

		if supported, reason := p.isSupportedC(); !supported {
			return supported, fmt.Sprintf("unsupported parameter %s : %s", p.Name, reason)
		}
	}

	return true, ""
}

func (pp Parameters) ignoreVarargs() {
	for _, p := range pp {
		if p.Varargs != nil {
			p.ignore = true
		}
	}
}
