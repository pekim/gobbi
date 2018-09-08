package gir

type Callbacks []*Callback

func (cs Callbacks) fixup(ns *Namespace) {
	for _, callback := range cs {
		callback.fixup(ns)
	}
}
