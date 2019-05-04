package generate

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

type Gir struct {
	repo        *Repository
	addendaRepo *Repository
}

func FromRoot(name, version string, noFormat bool) {
	girs := girNewRoot(name, version)

	for _, gir := range girs {
		gir.repo.Namespace.noFormat = noFormat
		gir.generate()
	}
}

func girNewRoot(name string, version string) []*Gir {
	girsMap := map[string]*Gir{}
	girNew(name, version, girsMap)

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

	girs := []*Gir{}
	for _, gir := range girsMap {
		girs = append(girs, gir)
	}
	sort.Slice(girs, func(i, j int) bool {
		return girs[i].repo.Namespace.goPackageName <
			girs[j].repo.Namespace.goPackageName
	})

	return girs
}

func girNew(name string, version string, girs map[string]*Gir) {
	fullname := name + "-" + version

	if _, haveGir := girs[fullname]; haveGir {
		return
	}

	g := &Gir{}
	g.repo = g.LoadFile(fullname+".gir", true)
	g.addendaRepo = g.LoadFile(fullname+"-addenda.gir", false)
	g.repo.MergeAddenda(g.addendaRepo)

	girs[fullname] = g

	for _, i := range g.repo.Includes {
		girNew(i.Name, i.Version, girs)
	}

	return
}

func (g *Gir) generate() {
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

	return girRepo
}
