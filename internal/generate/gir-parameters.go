package generate

type Parameters []*Parameter

func (pp Parameters) init(ns *Namespace) {
	//pp.pairUpArrayLengthParams()
	//pp.pairUpArgcArgvParams()

	for _, param := range pp {
		param.init(ns)
	}
}

func (pp Parameters) hasVarargs() bool {
	for _, p := range pp {
		if p.Varargs != nil {
			return true
		}
	}

	return false
}
