package generate

import (
	"path/filepath"
)

// ProjectFilepath creates an absolute file path to a project-relative
// path provided by any number of path elements.
func projectFilepath(pathElements ...string) string {
	basePath, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}

	basePathElements := filepath.SplitList(basePath)
	paths := append(basePathElements, pathElements...)

	return filepath.Join(paths...)
}
