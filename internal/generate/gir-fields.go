package generate

type Fields []*Field

func (cc Fields) init(ns *Namespace, record *Record) {
	for _, f := range cc {
		f.init(ns, record)
	}
}

func (ff Fields) generate(fi *file) {
	for _, f := range ff {
		f.generate(fi)
	}
}
