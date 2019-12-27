package generate

type Unions []*Union

func (uu Unions) init(ns *Namespace) {
	for _, union := range uu {
		union.init(ns)
	}
}

func (uu Unions) byName(name string) (*Union, bool) {
	for _, u := range uu {
		if u.Name == name {
			return u, true
		}
	}

	return nil, false
}
