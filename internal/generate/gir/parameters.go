package gir

type Parameters []*Parameter

func (pp Parameters) init(ns *Namespace) {
	for _, param := range pp {
		param.init(ns)
	}
}
