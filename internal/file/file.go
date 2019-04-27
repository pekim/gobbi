package file

import (
	"fmt"
	"go/format"
	"log"
	"sort"
)

const basePkg = "github.com/pekim/gobbi"

type File struct {
	pkg      string
	filepath string
	imports  map[string]bool
	contents string
}

func New(pkg string, filepath string) *File {
	return &File{
		pkg:      pkg,
		filepath: filepath,
		imports:  make(map[string]bool),
	}
}

func (f *File) imprt(pkg string) {
	f.imports[pkg] = true
}

func (f *File) gobbiImprt(pkg string) {
	f.imprt(basePkg + "/" + pkg)
}

func (f *File) importsSrc() string {
	// get package names, and sort them
	importNames := []string{}
	for imprt, _ := range f.imports {
		importNames = append(importNames, imprt)
	}
	sort.Strings(importNames)

	// create a string of import statements, each on a line
	imports := ""
	for _, imprt := range importNames {
		imports += fmt.Sprintf("import \"%s\"\n", imprt)
	}

	return imports + "\n"
}

func (f *File) Write() {

}

func (f *File) packageSrc() string {
	return fmt.Sprintf("package %s\n\n", f.pkg)
}

func (f *File) src() string {
	return "" +
		f.packageSrc() +
		f.importsSrc() +
		f.contents
}

func (f *File) formattedSrc() string {
	src, err := format.Source([]byte(f.src()))
	if err != nil {
		log.Fatalf("Failed to format source for %s : %s", f.filepath, err)
	}

	return string(src)
}
