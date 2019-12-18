package generate

type Constants []*Constant

func (cc Constants) init(ns *Namespace) {
	for _, constant := range cc {
		constant.init(ns)
	}
}
