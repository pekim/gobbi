package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"path/filepath"
)

func (n *Namespace) generateSys() {
	n.generateFile(filepath.Join(n.cDir, "package"), n.generatePackageFile)
}

// generatePackageFile generates a file with cgo pkg-config comments
// for the namespace's packages.
func (ns *Namespace) generatePackageFile(f *jen.File) {
	// pkg-config
	for _, pkg := range ns.repository.Packages {
		f.CgoPreamble(fmt.Sprintf("// #cgo pkg-config: %s", pkg.Name))
	}
}
