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

func (ee Enumerations) byName(name string) (*Enumeration, bool) {
	for _, enum := range ee {
		if enum.Name == name {
			return enum, true
		}
	}

	return nil, false
}

func (ee Enumerations) byGoTypeName(name string) (*Enumeration, bool) {
	for _, enum := range ee {
		if enum.goTypeName == name {
			return enum, true
		}
	}

	return nil, false
}
