package generate

type Constructors []*Constructor

func (cc Constructors) init(ns *Namespace, record *Record) {
	for _, ctor := range cc {
		ctor.init(ns, record)
	}
}

func (cc Constructors) generate(f *file) {
	for _, ctor := range cc {
		ctor.generate(f)
	}
}
