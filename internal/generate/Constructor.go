package generate

import (
	"github.com/dave/jennifer/jen"
)

type Constructor struct {
	*Function

	//record *Record
	//goName string
}

func (c *Constructor) init(ns *Namespace, record *Record) {
	c.Function.init(ns)
	c.GoName = record.GoName + makeExportedGoName(c.Name)
}

func (c *Constructor) generate(g *jen.Group, version *Version) {
	supported, reason := c.supported()

	if supported {
		c.Function.generate(g, version)
	} else {
		g.Commentf("Unsupported : %s", reason)
		g.Line()
	}
}
