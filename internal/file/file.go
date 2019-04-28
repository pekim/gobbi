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

	return imports
}

func (f *File) generatedFileComment() string {
	return "// This is a generated file - DO NOT EDIT\n"
}

func (f *File) buildTagsSrc() string {
	return ""
}

func (f *File) Write() {
	err := ioutil.WriteFile(f.filepath, f.formattedSrc(), 0644)
	if err != nil {
		log.Fatalf("Failed to write file %s : %s", f.filepath, err)
	}
}

func (f *File) packageSrc() string {
	return fmt.Sprintf("package %s\n", f.pkg)
}

func (f *File) src() string {
	return "" +
		f.generatedFileComment() +
		f.buildTagsSrc() +
		"\n" +
		f.packageSrc() +
		"\n" +
		f.importsSrc() +
		"\n" +
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

func (f *File) Line(text string) {
	f.contents += text + "\n"
}

func (f *File) Linef(format string, args ...interface{}) {
	f.Line(fmt.Sprintf(format, args...))
}

func (f *File) AddConstInt(name string, value int) {
	f.Linef("const %s int = %d", name, value)
}

func (f *File) AddConstString(name string, value string) {
	f.Linef("const %s string = \"%s\"", name, value)
}

func (f *File) AddConsts(typ string, cb func(add func(name string, value int))) {
	f.Line("const (")

	usedType := false
	cb(func(name string, value int) {
		if !usedType {
			usedType = true
			f.Linef("\t%s %s = %d", name, typ, value)
		} else {
			f.Linef("\t%s = %d", name, value)
		}
	})

	f.Line(")")
	f.Line("")
}
