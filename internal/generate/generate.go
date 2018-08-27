package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type Gir struct {

}

func GirNew() *Gir {
	return &Gir{}
}

func (g *Gir) Generate() {
	f := jen.NewFile("main")
	f.Func().Id("main").Params().Block(
		jen.Qual("fmt", "Println").Call(jen.Lit("Hello, world")),
	)
	fmt.Printf("%#v", f)
}
