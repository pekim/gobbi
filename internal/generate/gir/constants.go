package gir

type Constants []*Constant

func (cs Constants) fixup(ns *Namespace) {
	for _, constant := range cs {
		constant.fixup(ns)
	}
}
