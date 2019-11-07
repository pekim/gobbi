package generate

import "github.com/pekim/jennifer/jen"

type RecordEqualFunc struct {
	*Record
}

func (r *RecordEqualFunc) generate(g *jen.Group) {
	g.Commentf(
		"Equals compares this %s with another %s, and returns true if they represent the same GObject.",
		r.GoName, r.GoName)

	g.
		Func().
		Parens(jen.
			Id("recv").
			Op("*").
			Id(r.GoName)).
		Id("Equals").
		Params(jen.
			Id("other").
			Op("*").Id(r.GoName)).
		Params(jen.Bool()).
		BlockFunc(r.generateBody)

	g.Line()
}

func (r *RecordEqualFunc) generateBody(g *jen.Group) {
	g.
		Return().
		Id("other").Dot("ToC").Call().
		Op("==").
		Id("recv").Dot("ToC").Call()
}
