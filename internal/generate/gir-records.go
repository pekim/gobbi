package generate

type Records []*Record

func (rr Records) init(ns *Namespace) {
	for _, record := range rr {
		record.init(ns)
	}
}

func (rr Records) generate(f *file) {
	for _, r := range rr {
		r.generate(f)
	}
}
