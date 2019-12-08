package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

type Class struct {
	*Record
	Implements Implementss `xml:"implements"`
}

func (c *Class) init(ns *Namespace) {
	c.Record.init(ns, "Object")
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
	c.generateImplementss(f)
}

func (c *Class) generateImplementss(f *file) {
	for _, implements := range c.Implements {
		c.generateImplements(f, implements)
	}
}

func (c *Class) generateImplements(f *file, implements *Implements) {
	methodName := strings.Replace(implements.Name, ".", "", 1)

	f.Commentf("%s returns the %s interface implemented by %s",
		methodName, implements.Name, c.goName)

	isForeign, foreignNamespace, foreignName := c.namespace.namespaces.analyseName(implements.Name)

	var interface_ *jen.Statement
	if isForeign {
		interface_ = jen.Qual(foreignNamespace.goFullPackageName, foreignName)
	} else {
		interface_ = jen.Id(implements.Name)
	}

	var ctor *jen.Statement
	if isForeign {
		ctor = jen.Qual(foreignNamespace.goFullPackageName, foreignName+"NewFromNative")
	} else {
		ctor = jen.Id(implements.Name + "NewFromNative")
	}

	c.receiverFunc(f, methodName).
		Params().
		Params(jen.Op("*").Add(interface_)).
		Block(jen.Return().Add(ctor).Call(jen.Id(receiverName).Dot("Native").Call()))

	/*
		func (recv *MyClass) MyInterface() *MyInterface {
			return MyInterfaceNewFromNative(recv.Native())
		}
	*/
}
