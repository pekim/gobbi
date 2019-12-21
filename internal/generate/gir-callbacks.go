package generate

type Callbacks []*Callback

func (cc Callbacks) init(ns *Namespace) {
	for _, callback := range cc {
		callback.init(ns)
	}
}

func (cc Callbacks) byName(name string) (*Callback, bool) {
	for _, c := range cc {
		if c.Name == name {
			return c, true
		}
	}

	return nil, false
}
