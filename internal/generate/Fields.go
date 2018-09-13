package generate

import "github.com/dave/jennifer/jen"

type Fields []*Field

func (ff Fields) generate(g *jen.Group) {
	for _, f := range ff {
		f.generate(g)
	}
}
