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
	if len(ns.allVersions) == 0 {
		return
	}

	var tags string
	if version.value != "" {
		tags = ns.versionBuildTag(version)
	} else {
		allBuildTags := make([]string, len(ns.allVersions))
		for i, version := range ns.allVersions {
			allBuildTags[i] = "!" + ns.versionBuildTag(version)
		}
		tags = strings.Join(allBuildTags, ",")
	}

	file.HeaderComment(fmt.Sprintf("+build %s", tags))
}

func (ns *Namespace) versionBuildTag(version Version) string {
	return fmt.Sprintf("%s_%s", ns.goPackageName, version.value)
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
