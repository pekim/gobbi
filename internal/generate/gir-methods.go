package generate

type Methods []*Method

func (mm Methods) init(context *context, ns *Namespace, record *Record) {
	for _, m := range mm {
		m.init(context, ns, record)
	}
}

func (mm Methods) generate(f *file) {
	for _, method := range mm {
		method.generate(f)
	}
}
