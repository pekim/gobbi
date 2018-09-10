package generate

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

// getCollectionVersions gets an ordered list of versions
// from a collection that has elements with conditional
// version support.
func (ns *Namespace) getCollectionVersions(versionLister VersionLister) Versions {
	versions := versionLister.versionList()
	versions = versions.dedupe()
	versions.sort()

	return versions
}

// generateVersionDebugFunction conditionally generates
// a package "init" function that will print the package
// name and library version constraint for the file.
func (ns *Namespace) generateVersionDebugFunction(file *jen.File, version string) {
	if !ns.versionDebug {
		return
	}

	text := fmt.Sprintf("version : %s %s", ns.goPackageName, version)

	file.Func().Id("init").Params().BlockFunc(func(g *jen.Group) {
		g.Qual("fmt", "Println").Call(jen.Lit(text))
	})
}

func (ns *Namespace) buildConstraintsForVersion(file *jen.File, version Version) {
	if version.value == "" {
		return
	}

	tags := ns.constraintsForVersion(version)

	if tags != "" {
		file.HeaderComment(fmt.Sprintf("+build %s", tags))
	}
}

// constraintsForVersion builds a build tag string such that the generated
// file will be included when building for a specific library version
// or a later version.
//
// An example.
// If the target version is 4, then the file will NOT be included for
// versions 1, 2, or 3, but will for versions 4, 5, or 6.
//
// The string is built in the form "+build 4 5 6", which is
// interpreted as 4 OR 5 OR 6.
func (ns *Namespace) constraintsForVersion(version Version) string {
	eligibleVersions := ns.allVersions.versionStringsGreaterThanOrEqual(version)

	for i, eligibleVersion := range eligibleVersions {
		eligibleVersions[i] = fmt.Sprintf("%s_%s", ns.goPackageName, eligibleVersion)
	}

	if len(eligibleVersions) == 0 {
		return ""
	}

	return strings.Join(eligibleVersions, " ")
}

// allVersions gets the list of versions across all generatables.
func (ns *Namespace) setAllVersions() {
	versions := Versions{}
	versions = append(versions, ns.getCollectionVersions(ns.Bitfields)...)
	versions = append(versions, ns.getCollectionVersions(ns.Classes)...)
	versions = append(versions, ns.getCollectionVersions(ns.Enumerations)...)
	versions = append(versions, ns.getCollectionVersions(ns.Functions)...)
	versions = append(versions, ns.getCollectionVersions(ns.Records)...)

	versions = versions.dedupe()
	versions.sort()

	ns.allVersions = versions
}
