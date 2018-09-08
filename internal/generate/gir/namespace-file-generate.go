package gir

import (
	"fmt"
	"os"
	"path"

	"github.com/dave/jennifer/jen"
)

func (ns *Namespace) generateLibDir() {
	err := os.MkdirAll(ns.libDir, 0775)
	if err != nil {
		panic(err)
	}
}

func (ns *Namespace) generateFile(name string, generateContent func(g *jen.File)) {
	file := jen.NewFile(ns.goPackageName)

	generateContent(file)

	filepath := path.Join(ns.libDir, fmt.Sprintf("%s.go", name))
	err := file.Save(filepath)
	if err != nil {
		panic(err)
	}
}
