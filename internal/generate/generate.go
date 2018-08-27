package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type Gir struct {
	filepath string
}

func GirNew(name string) *Gir {
	filepath := projectFilepath("internal", "gir-files", name+".gir")

	return &Gir{
		filepath: filepath,
	}
}

func (g *Gir) Generate() {
	fmt.Println((g.filepath))

	f := jen.NewFile("main")
	f.Func().Id("main").Params().Block(
		jen.Qual("fmt", "Println").Call(jen.Lit("Hello, world")),
	)
	fmt.Printf("%#v", f)
}
