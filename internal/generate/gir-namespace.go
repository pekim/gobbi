package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/pekim/gobbi/internal/cgo/gi"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
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
	Classes      Classes      `xml:"class"`
	Constants    Constants    `xml:"constant"`
	Enumerations Enumerations `xml:"enumeration"`
	Functions    Functions    `xml:"function"`
	Records      Records      `xml:"record"`
	//	Interfaces                    Interfaces   `xml:"interface"`
	//	GenerateGobjectclassGotypeMap bool         `xml:"generate-gobjectclass-gotype-map,attr"`

	libDir            string
	namespaces        namespaces
	goPackageName     string
	goFullPackageName string
	cSymbolPrefixes   []string
	unsupportedCount  int
}

func (n *Namespace) init(namespaces namespaces) {
	n.namespaces = namespaces
	n.cSymbolPrefixes = strings.Split(n.CSymbolPrefixes, ",")
	n.goPackageName = strings.ToLower(n.Name)
	n.goFullPackageName = "github.com/pekim/gobbi/lib/" + n.goPackageName

	n.Aliases.init(n)
	n.Bitfields.init(n)
	n.Constants.init(n)
	n.Enumerations.init(n)
	n.Functions.init(n)
	n.Records.init(n)
	n.Classes.init(n)
}

func (n *Namespace) generate() {
	n.libDir = projectFilepath("..", "lib", n.goPackageName)
	n.generateLibDir()

	n.generateFile("gi", n.giFile)
	n.generateFile("alias", n.Aliases.generate)
	n.generateFile("bitfield", n.Bitfields.generate)
	n.generateFile("constant", n.Constants.generate)
	n.generateFile("enumeration", n.Enumerations.generate)
	n.generateFile("function", n.Functions.generate)
	n.generateFile("record", n.Records.generate)
	n.generateFile("class", n.Classes.generate)

	n.setUnsupportedCount()
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

	if _, found := n.Constants.byName(typeName); found {
		return true
	}

	if _, found := n.Enumerations.byName(typeName); found {
		return true
	}

	return false
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

func (n *Namespace) setUnsupportedCount() {
	re := regexp.MustCompile(`\/\/ UNSUPPORTED `)

	infos, err := ioutil.ReadDir(n.libDir)
	if err != nil {
		panic(err)
	}

	for _, info := range infos {
		path := filepath.Join(n.libDir, info.Name())
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		content := string(bytes)
		matches := re.FindAllString(content, -1)
		n.unsupportedCount += len(matches)
	}
}
