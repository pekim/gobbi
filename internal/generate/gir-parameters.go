package generate

import "fmt"

type Parameters []*Parameter

func (pp Parameters) init(ns *Namespace, parentContext *context) {
	pp.pairUpArrayLengthParams()

	for _, param := range pp {
		param.init(ns, parentContext)
	}
}

func (pp Parameters) pairUpArrayLengthParams() {
	for n, param := range pp {
		if param.Array == nil || param.Array.Length == nil {
			continue
		}
		arrayParam := param
		lengthParam := pp[*param.Array.Length]

		// Mutually reference the array and length params.
		lengthParam.lengthForParam = arrayParam
		lengthParam.lengthForParamN = n
		arrayParam.lengthParam = lengthParam
		arrayParam.lengthParamN = *param.Array.Length
	}
}

func (pp Parameters) byName(name string) (*Parameter, int, bool) {
	for n, param := range pp {
		if param.Name == name {
			return param, n, true
		}
	}

	return nil, -1, false
}

func (pp Parameters) allSupported() (bool, string) {
	for _, p := range pp {
		if supported, reason := p.isSupported(); !supported {
			return supported, fmt.Sprintf("parameter '%s' %s", p.Name, reason)
		}
	}

	return true, ""
}

func (pp Parameters) hasVarargs() bool {
	for _, p := range pp {
		if p.Varargs != nil {
			return true
		}
	}

	return false
}

func (pp Parameters) hasVaList() bool {
	for _, p := range pp {
		if p.Type != nil && p.Type.isVaList() {
			return true
		}
	}

	return false
}
