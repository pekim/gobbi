package generate

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Gir struct {
	repo        *Repository
	addendaRepo *Repository
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

func (g *Gir) LoadFile(filename string, required bool) *Repository {
	filepath := projectFilepath("internal", "gir-files", filename)
	source, err := ioutil.ReadFile(filepath)
	if err != nil {
		if os.IsNotExist(err) && !required {
			return &Repository{}
		}

		panic(err)
	}

	girRepo := &Repository{}
	err = xml.Unmarshal(source, girRepo)
	if err != nil {
		panic(fmt.Errorf("Failed to parse %s : %s", filepath, err))
	}
	girRepo.Init()

	return girRepo
}
