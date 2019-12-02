package generate

import (
	"github.com/dave/jennifer/jen"
)

type Class struct {
	*Record
	//Implements Implementss `xml:"implements"`
}

func (c *Class) generate(f *file) {
	c.Record.generate(f)
	//c.generateImplementss(f)
}

func (c *Class) generateImplementss(g *jen.Group) {
	//for _, implements := range c.Implements {
	//		c.generateImplements(g, implements)
	//}
}

//func (c *Class) generateImplements(g *jen.Group, implements *Implements) {
//}
