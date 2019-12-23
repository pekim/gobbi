package generate

import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
)

type Class struct {
	*Record
	//Implements Implementss `xml:"implements"`
}

func (c *Class) init(ns *Namespace) {
	c.namespace = ns
	c.applyAddenda()
	c.Record.init(ns)
}

func (c *Class) generateSys(f *jen.File, version semver.Version) {
	c.Constructors.generateSys(f, version)
}
