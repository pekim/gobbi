package file

import (
	"fmt"
	"go/format"
	"io/ioutil"
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
	if pkg == f.pkg {
		return
	}

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
	err := ioutil.WriteFile(f.filepath, f.formattedSrc(), 0644)
	if err != nil {
		log.Fatalf("Failed to write file %s : %s", f.filepath, err)
	}
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

func (f *File) formattedSrc() []byte {
	src, err := format.Source([]byte(f.src()))
	if err != nil {
		log.Fatalf("Failed to format source for %s : %s", f.filepath, err)
	}

	return src
}

func (f *File) Qualified(pkg string, id string) string {
	f.imprt(pkg)
	return pkg + "." + "id"
}

func (f *File) GobbiQualified(pkg string, id string) string {
	f.gobbiImprt(pkg)
	return pkg + "." + "id"
}

func (f *File) AddConstInt(name string, value int) {
	f.contents += fmt.Sprintf("const %s int = %d\n", name, value)
}

func (f *File) AddConstString(name string, value string) {
	f.contents += fmt.Sprintf("const %s string = \"%s\"\n", name, value)
}

func (f *File) AddConsts(typ string, cb func(add func(name string, value int))) {
	f.contents += "const (\n"

	usedType := false
	cb(func(name string, value int) {
		if !usedType {
			usedType = true
			f.contents += fmt.Sprintf("\t%s %s = %d\n", name, typ, value)
		} else {
			f.contents += fmt.Sprintf("\t%s = %d\n", name, value)
		}
	})

	f.contents += ")\n\n"
}
