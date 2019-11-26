package generate

type namespaces map[string]*Namespace

func (nn namespaces) byName(name string) (*Namespace, bool) {
	for _, ns := range nn {
		if ns.Name == name {
			return ns, true
		}
	}

	return nil, false
}
