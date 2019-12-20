package generate

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
	"path/filepath"
	"strings"
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
	ns.generateSysFileBuildTags(f, version)
	ns.repository.CIncludes.generate(f)

	ns.Bitfields.generateSys(f, version, "bitfields")
	ns.Enumerations.generateSys(f, version, "enumerations")
	ns.Records.generateSys(f, version)
}

func (ns *Namespace) generateSysFileBuildTags(f *jen.File, version semver.Version) {
	var buildTags string

	if versionIsNone(version) {
		versions := []string{}
		for _, v := range ns.versions.versions {
			if versionIsNone(v) {
				continue
			}

			versions = append(versions, "!"+ns.goPackageName+"_"+versionString(v))
		}

		// !pkg_1.20,!pkg_1.22,!pgk_1.30,..
		buildTags = strings.Join(versions, ",")
	} else {
		buildTags = fmt.Sprintf("%s_%s", ns.goPackageName, versionString(version))
	}

	f.HeaderComment(fmt.Sprintf("+build %s", buildTags))
	f.Line()
}
