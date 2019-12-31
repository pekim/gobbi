package generate

import (
	"github.com/dave/jennifer/jen"
)

func docVersion(f *jen.File, version string) {
	if version == "" {
		return
	}

	f.Commentf("//\n// since %s", version)
}
