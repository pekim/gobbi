package generate

import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
	"path/filepath"
	"sync"
)

func (n *Namespace) generateLib() {
	filename := filepath.Join(n.libDir, "package")
	n.generateFile(filename, n.generateSysPackageFile)

	var wg sync.WaitGroup
	for _, version := range n.versions.versions {
		wg.Add(1)

		go func(version semver.Version) {
			defer wg.Done()

			filename := filepath.Join(n.libDir, "v-"+versionString(version))
			n.generateFile(filename, func(f *jen.File) {
				n.generateLibFile(f, version)
			})
		}(version)
	}
	wg.Wait()
}

func (ns *Namespace) generateLibFile(f *jen.File, version semver.Version) {
	ns.generateFileBuildTags(f, version)

	ns.Aliases.generateLib(f, version)
	//ns.Records.generateLib(f, version)
}
