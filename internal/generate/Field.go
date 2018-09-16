package generate

import (
	"github.com/dave/jennifer/jen"
)

type Field struct {
	Namespace *Namespace

	Name     string `xml:"name,attr"`
	Writable int    `xml:"writable,attr"`
	Bits     int    `xml:"bits,attr"`
	Doc      *Doc   `xml:"doc"`
	Type     *Type  `xml:"type"`

	goVarName string
}

func (f *Field) init(ns *Namespace) {
	f.Namespace = ns
	f.goVarName = makeExportedGoName(f.Name)

	if f.Type != nil {
		f.Type.init(ns)
	}
}

func (f Field) generate(g *jen.Group) {
	if f.Bits > 0 {
		g.Commentf("Bitfield not supported : %2d %s", f.Bits, f.Name)
		return
	}

	if f.Type == nil {
		g.Commentf("no type for %s", f.Name)
		return
	}

	if f.Type.generator == nil {
		g.Commentf("%s : no type generator for %s, %s", f.Name, f.Type.Name, f.Type.CType)
		return
	}

	f.Type.generator.generateDeclaration(g, f.goVarName)
}
