package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"os"
	"path"
	"strings"
)

type Namespace struct {
	//Blacklist bool   `xml:"blacklist,attr"`
	Name    string `xml:"name,attr"`
	Version string `xml:"version,attr"`
	//SharedLibrary       string `xml:"shared-library,attr"`
	//CDocPath            string `xml:"c-doc-path,attr"`
	//CIdentifierPrefixes string `xml:"http://www.gtk.org/introspection/c/1.0 identifier-prefixes,attr"`
	//CSymbolPrefixes     string `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefixes,attr"`
	//	Aliases                       Aliases      `xml:"alias"`
	//	Bitfields                     Enumerations `xml:"bitfield"`
	//	Callbacks                     Callbacks    `xml:"callback"`
	//	Classes                       Classes      `xml:"class"`
	Constants Constants `xml:"constant"`
	//	Enumerations                  Enumerations `xml:"enumeration"`
	//	Functions                     Functions    `xml:"function"`
	//	Records                       Records      `xml:"record"`
	//	Interfaces                    Interfaces   `xml:"interface"`
	//	GenerateGobjectclassGotypeMap bool         `xml:"generate-gobjectclass-gotype-map,attr"`

	libDir        string
	namespaces    namespaces
	goPackageName string
}

func (n *Namespace) init() {
	n.Constants.init(n)
}

func (n *Namespace) generate() {
	n.goPackageName = strings.ToLower(n.Name)

	n.libDir = projectFilepath("..", "lib", n.goPackageName)
	n.generateLibDir()

	n.generateFile("constant", n.Constants.generate)
}

func (n *Namespace) generateLibDir() {
	err := os.MkdirAll(n.libDir, 0775)
	if err != nil {
		panic(err)
	}
}

func (n *Namespace) generateFile(name string, generateContent func(f *file)) {
	f := &file{
		jen.NewFile(n.goPackageName),
	}

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
