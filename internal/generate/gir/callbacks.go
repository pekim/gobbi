package gir

type Callbacks []*Callback

func (cs Callbacks) init(ns *Namespace) {
	for _, callback := range cs {
		callback.init(ns)
	}
}
