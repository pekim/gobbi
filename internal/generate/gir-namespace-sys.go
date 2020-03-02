package generate

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
	"path/filepath"
	"sync"
)

func (n *Namespace) generateSys() {
	filename := filepath.Join(n.cDir, "package")
	n.generateFile(filename, n.generateSysPackageFile)

	var wg sync.WaitGroup
	for _, version := range n.versions.versions {
		wg.Add(1)

		go func(version semver.Version) {
			defer wg.Done()

			filename := filepath.Join(n.cDir, "v-"+versionString(version))
			n.generateFile(filename, func(f *jen.File) {
				n.generateSysFile(f, version)
			})
		}(version)
	}
	wg.Wait()
}

// generateSysPackageFile generates a file with cgo pkg-config comments
// for the namespace's packages.
func (ns *Namespace) generateSysPackageFile(f *jen.File) {
	// pkg-config
	for _, pkg := range ns.repository.Packages {
		f.CgoPreamble(fmt.Sprintf("// #cgo pkg-config: %s", pkg.Name))

		/*
		 * Suppress C compiler warnings about deprecated functions.
		 *
		 * There are api functions that are deprecated from various
		 * library versions. The compiler warnings are noisy as
		 * they will be emitted regardless of whether such functions
		 * are used or not by an application.
		 */
		f.CgoPreamble("#cgo CFLAGS: -Wno-deprecated-declarations")
	}
}

func (ns *Namespace) generateSysFile(f *jen.File, version semver.Version) {
	ns.generateFileBuildTags(f, version)
	ns.repository.CIncludes.generate(f)

	ns.generateSysToCBoolFunction(f)

	ns.Records.generateSys(f, version)
	ns.Functions.generateSys(f, version)
	ns.Classes.generateSys(f, version)
	ns.Interfaces.generateSys(f, version)
}

func (ns *Namespace) generateSysToCBoolFunction(f *jen.File) {
	if ns.Name == "xlib" {
		return
	}

	f.
		Func().
		Id("toCBool").
		Params(jen.Id("b").Bool()).
		Params(jen.Qual("C", "gboolean")).
		Block(jen.
			If(jen.Id("b")).Block(jen.Return(jen.Qual("C", "TRUE"))),
			jen.Return(jen.Qual("C", "FALSE")),
		)

	f.
		Func().
		Id("toGoBool").
		Params(jen.Id("b").Qual("C", "gboolean")).
		Params(jen.Bool()).
		Block(jen.
			Return().Id("b").Op("==").Qual("C", "TRUE"))
}