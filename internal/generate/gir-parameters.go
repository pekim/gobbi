package generate

type Parameters []*Parameter

func (pp Parameters) init(ns *Namespace) {
	pp.pairUpArrayLengthParams()
	pp.pairUpArgcArgvParams()

	for _, param := range pp {
		param.init(ns)
	}
}

func (pp Parameters) pairUpArrayLengthParams() {
	for _, param := range pp {
		if param.Array == nil || param.Array.Length == nil {
			continue
		}
		arrayParam := param
		lengthParam := pp[*param.Array.Length]

		// Mutually reference the array and length params.
		lengthParam.lengthForParam = arrayParam
		arrayParam.lengthParam = lengthParam
	}
}

func (pp Parameters) pairUpArgcArgvParams() {
	argcParam, foundArgc := pp.byName("argc")
	argvParam, foundArgv := pp.byName("argv")
	if !foundArgc || !foundArgv {
		return
	}

	// Mutually reference the argc and argv params.
	argcParam.argvParam = argvParam
	argvParam.argcParam = argcParam
}

func (pp Parameters) byName(name string) (*Parameter, bool) {
	for _, param := range pp {
		if param.Name == name {
			return param, true
		}
	}

	return nil, false
}

func (pp Parameters) allSupported() (bool, string) {
	if pp.hasVarargs() {
		return false, "has varargs"
	}

	if pp.hasVaList() {
		return false, "has va_list"
	}

	if pp.hasCallback() {
		return false, "has callback"
	}

	if pp.hasLongDouble() {
		return false, "has long double"
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

func (pp Parameters) hasCallback() bool {
	for _, p := range pp {
		if p.Type != nil && p.Type.isCallback() {
			return true
		}
	}

	return false
}

func (pp Parameters) hasLongDouble() bool {
	for _, p := range pp {
		if p.Type != nil && p.Type.isLongDouble() {
			return true
		}
	}

	return false
}
