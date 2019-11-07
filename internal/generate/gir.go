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

func FromRoot(name, version string, noFormat bool, addToTotal func(n int), oneDone func()) {
	girsMap := map[string]*Gir{}
	gir := girNew(name, version, girsMap)

	// create a map of namespaces
	nn := make(map[string]*Namespace)
	for _, g := range girsMap {
		ns := g.repo.Namespace
		ns.repo = g.repo
		nn[ns.Name] = ns
	}

	// make all namespaces available all namespaces
	for _, ns := range nn {
		ns.namespaces = nn
	}

	// initialise each namespace
	for _, ns := range nn {
		ns.repo.Init()
	}

	gir.repo.Namespace.noFormat = noFormat
	if gir.repo.Namespace.Name == name && gir.repo.Namespace.Version == version {
		gir.repo.Generate(addToTotal, oneDone)
	}
}

func girNew(name string, version string, girs map[string]*Gir) *Gir {
	fullname := name + "-" + version

	if gir, haveGir := girs[fullname]; haveGir {
		return gir
	}

	g := &Gir{}
	g.repo = g.LoadFile(fullname+".gir", true)
	g.addendaRepo = g.LoadFile(fullname+"-addenda.gir", false)
	g.repo.MergeAddenda(g.addendaRepo)

	girs[fullname] = g

	for _, i := range g.repo.Includes {
		girNew(i.Name, i.Version, girs)
	}

	return g
}

func (g *Gir) LoadFile(filename string, required bool) *Repository {
	filepath := projectFilepath("gir-files", filename)
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

	return girRepo
}
