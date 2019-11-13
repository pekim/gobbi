package generate

type Enumerations []*Enumeration

func (ee Enumerations) init(ns *Namespace) {
	for _, enum := range ee {
		enum.init(ns)
	}
}

func (ee Enumerations) generate(f *file) {
	for _, enum := range ee {
		enum.generate(f)
	}
}
