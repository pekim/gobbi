package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type Implements struct {
	Name string `xml:"name,attr"`
}

type Class struct {
	*Record
	Implements []*Implements `xml:"implements"`
}

func (c *Class) mergeAddenda(addenda *Class) {
	c.Record.mergeAddenda(addenda.Record)
}

func (c *Class) generate(g *jen.Group, version *Version) {
	c.Record.generate(g, version)

	if supportedByVersion(c, version) {
		for _, implements := range c.Implements {
			fmt.Println(c.Namespace.Name, c.Name, implements.Name)
		}
	}
}
