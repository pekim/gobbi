package generate

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"gobbi2/internal/generate/gir"
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
	// f := jen.NewFile("main")
	// f.Func().Id("main").Params().Block(
	// 	jen.Qual("fmt", "Println").Call(jen.Lit("hw")),
	// )
	// fmt.Printf("%#v", f)
}

func (g *Gir) LoadFile() *gir.Repository {
	source, err := ioutil.ReadFile(g.filepath)
	if err != nil {
		panic(err)
	}

	girRepo := &gir.Repository{}
	err = xml.Unmarshal(source, girRepo)
	if err != nil {
		panic(fmt.Errorf("Failed to parse %s : %s", g.filepath, err))
	}
	// girRepo.Fixup()

	fmt.Println(girRepo)
	return girRepo
}
