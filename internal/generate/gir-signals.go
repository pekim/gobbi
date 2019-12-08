package generate

type Signals []*Signal

func (ss Signals) init(ns *Namespace, record *Record) {
	for _, signal := range ss {
		signal.init(ns, record)
	}
}

func (ss Signals) generate(f *file) {
	for _, signal := range ss {
		signal.generate(f)
	}
}

func (ss Signals) byName(name string) *Signal {
	for _, signal := range ss {
		if signal.Name == name {
			return signal
		}
	}

	return nil
}
