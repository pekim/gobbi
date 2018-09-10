package generate

type Callbacks []*Callback

func (cc Callbacks) init(ns *Namespace) {
	for _, callback := range cc {
		callback.init(ns)
	}
}
