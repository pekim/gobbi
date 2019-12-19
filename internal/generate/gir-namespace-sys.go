package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"path/filepath"
)

func (n *Namespace) generateC() {
	n.generateFile(filepath.Join(n.cDir, "package"), n.generatePackageFile)

	for _, version := range n.versions.versions {
		n.generateFile(filepath.Join(n.cDir, "v-"+version.String()), n.generateSysFile)
	}
}

// generatePackageFile generates a file with cgo pkg-config comments
// for the namespace's packages.
func (ns *Namespace) generatePackageFile(f *jen.File) {
	// pkg-config
	for _, pkg := range ns.repository.Packages {
		f.CgoPreamble(fmt.Sprintf("// #cgo pkg-config: %s", pkg.Name))
	}
}

func (ns *Namespace) generateSysFile(f *jen.File) {
	ns.repository.CIncludes.generate(f)

	ns.Bitfields.generateSys(f, "bitfields")
	ns.Enumerations.generateSys(f, "enumerations")
}
