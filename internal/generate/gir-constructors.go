package generate

type Constructors []*Constructor

func (cc Constructors) init(context *context, ns *Namespace, record *Record) {
	for _, ctor := range cc {
		ctor.init(context, ns, record)
	}
}

func (cc Constructors) generate(f *file) {
	for _, ctor := range cc {
		ctor.generate(f)
	}
}
