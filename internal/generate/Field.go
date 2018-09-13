package generate

import (
	"github.com/dave/jennifer/jen"
)

type Field struct {
	Name     string `xml:"name,attr"`
	Writable int    `xml:"writable,attr"`
	Bits     int    `xml:"bits,attr"`
	Doc      *Doc   `xml:"doc"`
	Type     *Type  `xml:"type"`
}

func (f Field) generate(g *jen.Group) {
	if f.Bits > 0 {
		g.Commentf("Bitfield not supported : %2d %s", f.Bits, f.Name)
		return
	}

	g.
		Id(makeExportedGoName(f.Name)).
		Id("int")
}
