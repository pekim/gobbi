package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/pekim/gobbi/internal/gi"
	"os"
	"path"
	"strings"
	"unicode"
)

type Namespace struct {
	//Blacklist bool   `xml:"blacklist,attr"`
	Name    string `xml:"name,attr"`
	Version string `xml:"version,attr"`
	//SharedLibrary       string `xml:"shared-library,attr"`
	//CDocPath            string `xml:"c-doc-path,attr"`
	CIdentifierPrefixes string       `xml:"http://www.gtk.org/introspection/c/1.0 identifier-prefixes,attr"`
	CSymbolPrefixes     string       `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefixes,attr"`
	Aliases             Aliases      `xml:"alias"`
	Bitfields           Enumerations `xml:"bitfield"`
	//	Callbacks                     Callbacks    `xml:"callback"`
	//	Classes                       Classes      `xml:"class"`
	Constants    Constants    `xml:"constant"`
	Enumerations Enumerations `xml:"enumeration"`
	Functions    Functions    `xml:"function"`
	Records      Records      `xml:"record"`
	//	Interfaces                    Interfaces   `xml:"interface"`
	//	GenerateGobjectclassGotypeMap bool         `xml:"generate-gobjectclass-gotype-map,attr"`

	libDir          string
	namespaces      namespaces
	goPackageName   string
	cSymbolPrefixes []string
}

func (n *Namespace) init(namespaces namespaces) {
	n.namespaces = namespaces
	n.cSymbolPrefixes = strings.Split(n.CSymbolPrefixes, ",")

	n.Aliases.init(n)
	n.Bitfields.init(n)
	n.Constants.init(n)
	n.Enumerations.init(n)
	n.Functions.init(n)
	n.Records.init(n)
}

func (n *Namespace) generate() {
	n.goPackageName = strings.ToLower(n.Name)

	n.libDir = projectFilepath("..", "lib", n.goPackageName)
	n.generateLibDir()

	n.generateFile("gi", n.giFile)
	n.generateFile("alias", n.Aliases.generate)
	n.generateFile("bitfield", n.Bitfields.generate)
	n.generateFile("constant", n.Constants.generate)
	n.generateFile("enumeration", n.Enumerations.generate)
	n.generateFile("function", n.Functions.generate)
	n.generateFile("record", n.Records.generate)
}

func (n *Namespace) generateLibDir() {
	err := os.MkdirAll(n.libDir, 0775)
	if err != nil {
		panic(err)
	}
}

func (n *Namespace) generateFile(name string, generateContent func(f *file)) {
	jenFile := jen.NewFile(n.goPackageName)
	f := &file{jenFile}

	// Use a standard generated file comment format.
	// https://github.com/golang/go/issues/13560#issuecomment-288457920
	f.HeaderComment(("Code generated - DO NOT EDIT."))
	f.Line()

	generateContent(f)

	filepath := path.Join(n.libDir, fmt.Sprintf("%s.go", name))
	err := f.Save(filepath)
	if err != nil {
		panic(err)
	}
}

func (n *Namespace) haveType(typeName string) bool {
	if _, found := n.Aliases.byName(typeName); found {
		return true
	}
	if _, found := n.Constants.findByName(typeName); found {
		return true
	}
	return false
}

func (n *Namespace) jenGoTypeForTypeName(typeName string) (*jen.Statement, bool) {
	// If not starts with upper case then not a reference to a type.
	// (Referenceing first byte as a rune should be fine for type names.)
	if !unicode.IsUpper(rune(typeName[0])) {
		return nil, false
	}

	parts := strings.Split(typeName, ".")

	if len(parts) == 1 {
		if n.haveType(parts[0]) {
			return jen.Id(parts[0]), true
		}
	}

	if len(parts) == 2 {
		// TODO qualified reference; find in namespaces
	}

	return nil, false
}

func (n *Namespace) giFile(f *file) {
	f.
		Func().
		Id("init").
		Params().
		Block(
			jen.
				Qual(gi.PackageName, "Require").
				Call(jen.Lit(n.Name), jen.Lit(n.Version)))
}

func (n *Namespace) outParameterGeneratorByName(name string) (outParameterGenerator, bool) {
	if record, found := n.Records.byName(name); found {
		return record, true
	}

	return nil, false
}
