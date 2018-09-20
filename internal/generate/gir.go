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

func GirNewRoot(name string, version string) *Gir {
	return GirNew(name, version, map[string]*Gir{})
}

func GirNew(name string, version string, girs map[string]*Gir) *Gir {
	fullname := name + "-" + version

	if g, haveGir := girs[fullname]; haveGir {
		return g
	}

	g := &Gir{}
	g.repo = g.LoadFile(fullname+".gir", true)
	g.addendaRepo = g.LoadFile(fullname+"-addenda.gir", false)
	g.repo.MergeAddenda(g.addendaRepo)

	girs[fullname] = g

	for _, i := range g.repo.Includes {
		GirNew(i.Name, i.Version, girs)
	}

	return g
}

func (g *Gir) Generate() {
	g.repo.Generate()
}

func (g *Gir) LoadFile(filename string, required bool) *Repository {
	fmt.Println(filename)
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
