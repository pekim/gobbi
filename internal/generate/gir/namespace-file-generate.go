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

func (ns *Namespace) generateFile(name string, generateContent func(f *jen.File)) {
	file := jen.NewFile(ns.goPackageName)

	generateContent(file)

	filepath := path.Join(ns.libDir, fmt.Sprintf("%s.go", name))
	err := file.Save(filepath)
	if err != nil {
		panic(err)
	}
}

// generatePackageFile generates a file with cgo pkg-config comments
// for the namespace's packages.
func (ns *Namespace) generatePackageFile() {
	ns.generateFile("package", func(f *jen.File) {
		for _, pkg := range ns.repo.Packages {
			f.CgoPreamble(fmt.Sprintf("// #cgo pkg-config: %s", pkg.Name))
		}

		/*
		 * Suppress C compiler warnings about deprecated functions.
		 *
		 * There are api functions that are deprecated from various
		 * library versions. The compilers warnings are noisy as
		 * they will be emitted regardless of whether such functions
		 * are used or not by an application.
		 */
		f.CgoPreamble("// #cgo CFLAGS: -Wno-deprecated-declarations")
	})
}
