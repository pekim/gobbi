package generate

import "github.com/dave/jennifer/jen"

type Doc struct {
	Text string `xml:",chardata"`
}

func generateDoc(d *Doc, g *jen.Group) {
	if d != nil {
		d.generate(g)
	}
}

func (d *Doc) generate(g *jen.Group) {
	if d.Text == "" {
		return
	}

	g.Comment(d.Text)
}
