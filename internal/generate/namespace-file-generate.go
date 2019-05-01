package generate

import "C"
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
	ns.jenFile = file

	file.HeaderComment(("This is a generated file - DO NOT EDIT"))
	file.Line()

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
		// pkg-config
		for _, pkg := range ns.repo.Packages {
			f.CgoPreamble(fmt.Sprintf("// #cgo pkg-config: %s", pkg.Name))
		}
	})
}

//func (ns *Namespace) generateBooleanFile() {
//	ns.generateFile("boolean", func(f *jen.File) {
//		f.
//			Func().
//			Id("boolToGboolean").
//			Params(jen.Id("b").Id("bool")).
//			Parens(jen.Qual("C", "gboolean")).
//			BlockFunc(func(g *jen.Group) {
//				g.
//					If(jen.Id("b")).
//					Block(jen.Return(jen.Qual("C", "TRUE")))
//
//				g.Return(jen.Qual("C", "FALSE"))
//			})
//	})
//}

func (ns *Namespace) cgoPreambleHeaders(file *jen.File, version Version) {
	/*
	 * Suppress C compiler warnings about deprecated functions.
	 *
	 * There are api functions that are deprecated from various
	 * library versions. The compilers warnings are noisy as
	 * they will be emitted regardless of whether such functions
	 * are used or not by an application.
	 */
	file.CgoPreamble("#cgo CFLAGS: -Wno-deprecated-declarations")

	// Suppress C compiler warnings from format function wrappers.
	file.CgoPreamble("#cgo CFLAGS: -Wno-format-security")
	file.CgoPreamble("#cgo CFLAGS: -Wno-incompatible-pointer-types")

	ns.repo.CIncludes.generate(file, version)

	file.CgoPreamble("#include <stdlib.h>")
}

func (ns *Namespace) generateGeneratables(typeName string, generatables Generatables) {
	// file for non version-specific entities
	if generatables.generatesC() {
		ns.generateEntityVersionedFileC(typeName+"-c", Version{}, generatables)
	}
	ns.generateEntityVersionedFile(typeName, Version{}, generatables)

	// files for version-specific entities
	versions := ns.getCollectionVersions(generatables)
	for _, version := range versions {
		if generatables.generatesC() {
			ns.generateEntityVersionedFileC(typeName+"-c-"+version.value, version, generatables)
		}

		ns.generateEntityVersionedFile(typeName+"-"+version.value, version, generatables)
	}
}

func (ns *Namespace) generateEntityVersionedFileC(filename string, version Version, generatables Generatables) {
	if !ns.someGeneratablesSupportVersion(version, generatables) {
		return
	}

	ns.generateFile(filename, func(f *jen.File) {
		ns.buildConstraintsForVersion(f, version)
		ns.cgoPreambleHeaders(f, version)
		ns.generateVersionDebugFunction(f, version.value)

		for _, entity := range generatables.entities() {
			if supported, reason := entity.supported(); !supported {
				if !supportedByVersion(entity, &version) {
					continue
				}

				f.Commentf("Unsupported : %s", reason)
				f.Line()
				continue
			}

			if blacklisted, detail := entity.blacklisted(); blacklisted {
				if !supportedByVersion(entity, &version) {
					continue
				}

				f.Commentf("Blacklisted : %s", detail)
				f.Line()
				continue
			}

			entity.generateC(f.Group, &version)
		}
	})
}

// generateEntityVersionedFile generates a file for Generatables that
// meet the version criterion.
func (ns *Namespace) generateEntityVersionedFile(filename string, version Version, generatables Generatables) {
	if !ns.someGeneratablesSupportVersion(version, generatables) {
		return
	}

	ns.generateFile(filename, func(f *jen.File) {
		ns.buildConstraintsForVersion(f, version)
		//ns.cgoPreambleHeaders(f, version)
		ns.generateVersionDebugFunction(f, version.value)

		for _, entity := range generatables.entities() {
			if supported, reason := entity.supported(); !supported {
				if !supportedByVersion(entity, &version) {
					continue
				}

				f.Commentf("Unsupported : %s", reason)
				f.Line()
				continue
			}

			if blacklisted, detail := entity.blacklisted(); blacklisted {
				if !supportedByVersion(entity, &version) {
					continue
				}

				f.Commentf("Blacklisted : %s", detail)
				f.Line()
				continue
			}

			entity.generate(f.Group, &version)
		}
	})
}

func (ns *Namespace) someGeneratablesSupportVersion(version Version, generatables Generatables) bool {
	for _, entity := range generatables.entities() {
		blacklisted, _ := entity.blacklisted()

		if supportedByVersion(entity, &version) && !blacklisted {
			return true
		}
	}

	return false
}
