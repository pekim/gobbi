package gir

type Callback struct {
	*Function
	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
}

type Callbacks []*Callback

func (cs Callbacks) fixup(ns *Namespace) {
	for _, callback := range cs {
		callback.fixup(ns)
	}
}
