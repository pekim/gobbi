package generate

type Records []*Record

func (rr Records) init(ns *Namespace) {
	for _, record := range rr {
		record.init(ns, false)
	}
}

func (rr Records) generate(f *file) {
	for _, r := range rr {
		r.generate(f)
	}
}

func (rr Records) byName(name string) (*Record, bool) {
	for _, r := range rr {
		if r.Name == name {
			return r, true
		}
	}

	return nil, false
}
