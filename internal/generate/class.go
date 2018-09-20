package generate

type Class struct {
	*Record
}

func (c *Class) mergeAddenda(addenda *Class) {
	c.Record.mergeAddenda(addenda.Record)
}
