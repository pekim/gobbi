package gir

type Constants []*Constant

func (cs Constants) init(ns *Namespace) {
	for _, constant := range cs {
		constant.init(ns)
	}
}
