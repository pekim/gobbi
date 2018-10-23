package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type Class struct {
	*Record
	Implements Implementss `xml:"implements"`
}

func (c *Class) mergeAddenda(addenda *Class) {
	c.Record.mergeAddenda(addenda.Record)
	c.Implements.mergeAddenda(addenda.Implements)
}

func (c *Class) generate(g *jen.Group, version *Version) {
	c.Record.generate(g, version)
	c.generateImplementss(g, version)
}

func (c *Class) generateImplementss(g *jen.Group, version *Version) {
	if supportedByVersion(c, version) {
		for _, implements := range c.Implements {
			if supportedByVersion(implements, version) {
				c.generateImplements(g, implements)
			}
		}
	}
}

func (c *Class) generateImplements(g *jen.Group, implements *Implements) {
	qname := QNameNew(c.Namespace, implements.Name)

	iface, found := qname.ns.interfaceForName(qname.name)
	if !found {
		panic(fmt.Sprintf("Failed to interface %s for %s", implements.Name, c.Name))
	}

	/*
		func (recv *MyClass) MyInterface() *MyInterface {
			return MyInterfaceNewFromC(recv.ToC())
		}
	*/
	g.
		Func().
		Params(jen.Id("recv").Op("*").Id(c.GoName)).
		Id(qname.name).
		Params().
		ParamsFunc(func(g *jen.Group) {
			if qname.sameNamespace {
				g.Op("*").Id(qname.name)
			} else {
				g.Op("*").Qual(qname.ns.fullGoPackageName, qname.name)
			}
		}).
		BlockFunc(func(g *jen.Group) {
			g.
				Return().
				Do(func(s *jen.Statement) {
					if qname.sameNamespace {
						s.Id(iface.newFromCFuncName)
					} else {
						s.Qual(qname.ns.fullGoPackageName, iface.newFromCFuncName)
					}

				}).
				Call(jen.Id("recv").Dot("ToC").Call())
		})

	g.Line()
}
