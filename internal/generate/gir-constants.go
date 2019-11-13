package generate

type Constants []*Constant

func (cc Constants) init(ns *Namespace) {
	for _, constant := range cc {
		constant.init(ns)
	}
}

func (cc Constants) generate(f *file) {
	for _, constant := range cc {
		constant.generate(f)
	}
}

func (cc Constants) findByName(name string) (*Constant, bool) {
	for _, constant := range cc {
		if constant.Name == name {
			return constant, true
		}
	}

	return nil, false
}
