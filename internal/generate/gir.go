package generate

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"gobbi2/internal/generate/file"
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
	g.repo.MergeAddenda(g.addendaRepo)
	return g
}

func (g *Gir) Generate() {
	g.repo.Generate()
}

func (g *Gir) LoadFile(filename string, required bool) *gir.Repository {
	filepath := file.ProjectFilepath("internal", "gir-files", filename)
	source, err := ioutil.ReadFile(filepath)
	if err != nil {
		if os.IsNotExist(err) && !required {
			fmt.Println(err)
			return &gir.Repository{}
		}

		panic(err)
	}

	girRepo := &gir.Repository{}
	err = xml.Unmarshal(source, girRepo)
	if err != nil {
		panic(fmt.Errorf("Failed to parse %s : %s", filepath, err))
	}
	girRepo.Fixup()

	return girRepo
}
