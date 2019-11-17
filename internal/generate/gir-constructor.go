package generate

type Constructor struct {
	*Function
}

func (c *Constructor) init(ns *Namespace, record *Record) {
	// Some constructors defined in gir files reference a subclass
	// of their real type.
	// Ensure that all constructors return a pointer to their
	// record/class type.
	c.ReturnValue.Type.CType = record.CType + "*"
	c.ReturnValue.Type.Name = record.Name

	c.Function.init(ns /*nil, ""*/)
	c.goName = record.goName + makeExportedGoName(c.Name)
	//c.Function.ctorRecord = record

	if record.Version != "" && c.Version == "" {
		c.Version = record.Version
	}
}

func (c *Constructor) generate(f *file) {
	c.Function.generate(f)
}
