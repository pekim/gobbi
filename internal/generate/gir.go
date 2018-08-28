package generate

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"gobbi2/internal/generate/gir"
)

type Gir struct {
	repo        *gir.Repository
	addendaRepo *gir.Repository
}

func GirNew(name string) *Gir {

	g := &Gir{}

	g.repo = g.LoadFile(name+".gir", true)
	g.addendaRepo = g.LoadFile(name+"-addenda.gir", false)

	fmt.Println(g)
	return g
}

func (g *Gir) Generate() {
	// f := jen.NewFile("main")
	// f.Func().Id("main").Params().Block(
	// 	jen.Qual("fmt", "Println").Call(jen.Lit("hw")),
	// )
	// fmt.Printf("%#v", f)
}

func (g *Gir) LoadFile(filename string, required bool) *gir.Repository {
	filepath := projectFilepath("internal", "gir-files", filename)
	source, err := ioutil.ReadFile(filepath)
	if err != nil {
		if os.IsNotExist(err) && !required {
			return &gir.Repository{}
		}

		panic(err)
	}

	girRepo := &gir.Repository{}
	err = xml.Unmarshal(source, girRepo)
	if err != nil {
		panic(fmt.Errorf("Failed to parse %s : %s", filepath, err))
	}
	// girRepo.Fixup()

	return girRepo
}
