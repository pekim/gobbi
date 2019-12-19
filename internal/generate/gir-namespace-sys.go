package generate

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
	"path/filepath"
)

func (n *Namespace) generateC() {
	filename := filepath.Join(n.cDir, "package")
	n.generateFile(filename, n.generatePackageFile)

	for _, version := range n.versions.versions {
		filename := filepath.Join(n.cDir, "v-"+versionString(version))

		n.generateFile(filename, func(f *jen.File) {
			n.generateSysFile(f, version)
		})
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

func (ns *Namespace) generateSysFile(f *jen.File, version semver.Version) {
	if !versionIsNone(version) {
		buildTag := fmt.Sprintf("+build %s_%s", ns.goPackageName, versionString(version))
		f.HeaderComment(buildTag)
		f.Line()
	}

	ns.repository.CIncludes.generate(f)

	ns.Bitfields.generateSys(f, "bitfields")
	ns.Enumerations.generateSys(f, "enumerations")
}
