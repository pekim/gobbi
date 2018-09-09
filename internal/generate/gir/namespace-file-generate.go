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

func (ns *Namespace) cgoPreambleHeaders(file *jen.File) {
	for _, cInclude := range ns.repo.CIncludes {
		file.CgoPreamble(fmt.Sprintf("#include <%s>", cInclude.Name))
	}

	file.CgoPreamble("#include <stdlib.h>")
	// file.CgoPreamble("#include \"callback.h\"")
}

func (ns *Namespace) generateGeneratables(typeName string, generatables Generatables) {
	versions := ns.getCollectionVersions(generatables)

	ns.generateFile(typeName, func(f *jen.File) {
		ns.cgoPreambleHeaders(f)
		ns.generateVersionDebugFunction(f, "")

		for _, entity := range generatables.entities() {
			entity.generate(f.Group, nil)
		}
	})

	for _, version := range versions {
		ns.generateFile(typeName+"-"+version.value, func(f *jen.File) {
			ns.buildConstraintsForVersion(f, version)
			ns.cgoPreambleHeaders(f)
			ns.generateVersionDebugFunction(f, version.value)

			for _, entity := range generatables.entities() {
				entity.generate(f.Group, &version)
			}
		})
	}
}
