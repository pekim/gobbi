package generate

type Classes []*Class

func (cc Classes) init(ns *Namespace) {
	for _, class := range cc {
		class.init(ns)
	}
}

func (cc Classes) generate(f *file) {
	for _, c := range cc {
		c.generate(f)
	}
}

func (cc Classes) byName(name string) (*Class, bool) {
	for _, c := range cc {
		if c.Name == name {
			return c, true
		}
	}

	return nil, false
}
