package generate

type Class struct {
	*Record
	//Implements Implementss `xml:"implements"`
}

func (c *Class) init(ns *Namespace) {
	c.namespace = ns
	c.applyAddenda()
	c.Record.init(ns)
}
