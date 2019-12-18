package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Namespace struct {
	Name                string `xml:"name,attr"`
	Version             string `xml:"version,attr"`
	CIdentifierPrefixes string `xml:"http://www.gtk.org/introspection/c/1.0 identifier-prefixes,attr"`
	CSymbolPrefixes     string `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefixes,attr"`
	//Aliases             Aliases      `xml:"alias"`
	//Bitfields           Enumerations `xml:"bitfield"`
	//	Callbacks                     Callbacks    `xml:"callback"`
	//Classes      Classes      `xml:"class"`
	//Constants    Constants    `xml:"constant"`
	//Enumerations Enumerations `xml:"enumeration"`
	//Functions    Functions    `xml:"function"`
	//Records      Records      `xml:"record"`
	//Interfaces   Interfaces   `xml:"interface"`

	repository        *repository
	libDir            string
	cDir              string
	namespaces        namespaces
	goPackageName     string
	goFullPackageName string
	cSymbolPrefixes   []string
	unsupportedCount  int
}

func (n *Namespace) init(repository *repository, namespaces namespaces) {
	n.repository = repository
	n.namespaces = namespaces
	n.cSymbolPrefixes = strings.Split(n.CSymbolPrefixes, ",")
	n.goPackageName = strings.ToLower(n.Name)
	n.goFullPackageName = "github.com/pekim/gobbi/lib/" + n.goPackageName

	//n.Aliases.init(n)
	//n.Bitfields.init(n)
	//n.Constants.init(n)
	//n.Enumerations.init(n)
	//n.Functions.init(n)
	//n.Records.init(n)
	//n.Classes.init(n)
	//n.Interfaces.init((n))
}

func (n *Namespace) generate() {
	libDir := projectFilepath("..", "lib")
	n.libDir = filepath.Join(libDir, n.goPackageName)
	n.cDir = filepath.Join(libDir, "internal", "c", n.goPackageName)

	n.mkDirs()

	n.generateFile(filepath.Join(n.cDir, "sys"), n.generateSys)

	//n.generateFile("alias", n.Aliases.generate)
	//n.generateFile("bitfield", n.Bitfields.generate)
	//n.generateFile("constant", n.Constants.generate)
	//n.generateFile("enumeration", n.Enumerations.generate)
	//n.generateFile("function", n.Functions.generate)
	//n.generateFile("record", n.Records.generate)
	//n.generateFile("class", n.Classes.generate)
	//n.generateFile("interface", n.Interfaces.generate)

	n.setUnsupportedCount()
}

func (n *Namespace) mkDirs() {
	err := os.MkdirAll(n.libDir, 0775)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(n.cDir, 0775)
	if err != nil {
		panic(err)
	}
}

func (n *Namespace) generateFile(name string, generateContent func(f *jen.File)) {
	f := jen.NewFile(n.goPackageName)

	// Use a standard generated file comment format.
	// https://github.com/golang/go/issues/13560#issuecomment-288457920
	f.HeaderComment(("Code generated - DO NOT EDIT."))
	f.Line()

	generateContent(f)

	filepath := fmt.Sprintf("%s.go", name)
	err := f.Save(filepath)
	if err != nil {
		panic(err)
	}
}

func (n *Namespace) setUnsupportedCount() {
	re := regexp.MustCompile(`// UNSUPPORTED `)

	infos, err := ioutil.ReadDir(n.libDir)
	if err != nil {
		panic(err)
	}

	for _, info := range infos {
		path := filepath.Join(n.libDir, info.Name())
		fi, _ := os.Stat(path)
		if fi.IsDir() {
			continue
		}

		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		content := string(bytes)
		matches := re.FindAllString(content, -1)
		n.unsupportedCount += len(matches)
	}
}
