package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type Class struct {
	*Record
	//Implements Implementss `xml:"implements"`
}

func (c *Class) init(ns *Namespace) {
	c.Record.init(ns)
	c.setParent()
}

func (c *Class) setParent() {
	if c.ParentName == "" {
		return
	}

	isForeign, foreignNamespace, foreignName := c.namespace.namespaces.analyseName(c.ParentName)
	if isForeign {
		c.parentNamespace = foreignNamespace

		parent, found := foreignNamespace.Classes.byName(foreignName)
		if !found {
			panic(fmt.Sprintf("Failed to find parent class %s in namespace %s", foreignName, foreignNamespace.Name))
		}
		c.parent = parent
	} else {
		parent, found := c.namespace.Classes.byName(c.ParentName)
		if !found {
			panic(fmt.Sprintf("Failed to find parent class %s in namespace %s", c.ParentName, c.namespace.Name))
		}
		c.parent = parent
	}
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
