package generate

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Namespace struct {
	Name                string       `xml:"name,attr"`
	Version             string       `xml:"version,attr"`
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
	Interfaces   Interfaces   `xml:"interface"`
	Callbacks    Callbacks    `xml:"callback"`
	Unions       Unions       `xml:"union"`

	repository           *repository
	libDir               string
	cDir                 string
	namespaces           namespaces
	goPackageName        string
	goFullPackageName    string
	goFullSysPackageName string
	cSymbolPrefixes      []string
	cUnsupportedCount    int
	goUnsupportedCount   int
	versions             *versions
}

func (n *Namespace) init(repository *repository, namespaces namespaces) {
	n.repository = repository
	n.namespaces = namespaces
	n.cSymbolPrefixes = strings.Split(n.CSymbolPrefixes, ",")
	n.goPackageName = strings.ToLower(n.Name)
	n.goFullPackageName = "github.com/pekim/gobbi/lib/" + n.goPackageName
	n.goFullSysPackageName = "github.com/pekim/gobbi/lib/internal/c/" + n.goPackageName
	n.versions = &versions{}

	n.Aliases.init(n)
	n.Bitfields.init(n)
	n.Constants.init(n)
	n.Enumerations.init(n)
	n.Functions.init(n)
	n.Records.init(n)
	n.Classes.init(n)
	n.Interfaces.init(n)
	n.Callbacks.init(n)
	n.Unions.init(n)

	n.versions.sort()
}

func (n *Namespace) generate() {
	libDir := projectFilepath("..", "lib")
	n.libDir = filepath.Join(libDir, n.goPackageName)
	n.cDir = filepath.Join(libDir, "internal", "c", n.goPackageName)

	n.mkDirs()

	n.generateSys()
	n.generateLib()

	//n.generateFile("alias", n.Aliases.generate)
	//n.generateFile("bitfield", n.Bitfields.generate)
	//n.generateFile("constant", n.Constants.generate)
	//n.generateFile("enumeration", n.Enumerations.generate)
	//n.generateFile("function", n.Functions.generate)
	//n.generateFile("record", n.Records.generate)
	//n.generateFile("class", n.Classes.generate)
	//n.generateFile("interface", n.Interfaces.generate)

	n.cUnsupportedCount = n.getUnsupportedCount(n.cDir)
	n.goUnsupportedCount = n.getUnsupportedCount(n.libDir)
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

func (n *Namespace) getUnsupportedCount(dir string) int {
	re := regexp.MustCompile(`// UNSUPPORTED (.*)`)

	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	uniqUnsupported := make(map[string]bool)

	for _, info := range infos {
		path := filepath.Join(dir, info.Name())
		fi, _ := os.Stat(path)
		if fi.IsDir() {
			continue
		}

		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		content := string(bytes)
		matches := re.FindAllStringSubmatch(content, -1)

		for _, m := range matches {
			uniqUnsupported[m[0]] = true
		}
	}

	return len(uniqUnsupported)
}

func (ns *Namespace) generateFileBuildTagsSys(f *jen.File, version semver.Version) {
	var buildTags string

	if versionIsNone(version) {
		versions := []string{}
		for _, v := range ns.versions.versions {
			if versionIsNone(v) {
				continue
			}

			versions = append(versions, "!"+ns.goPackageName+"_"+versionString(v))
		}

		// !pkg_1.20,!pkg_1.22,!pgk_1.30,..
		buildTags = strings.Join(versions, ",")
	} else {
		buildTags = fmt.Sprintf("%s_%s", ns.goPackageName, versionString(version))
	}

	if buildTags == "" {
		return
	}

	f.HeaderComment(fmt.Sprintf("+build %s", buildTags))
	f.Line()
}

func (ns *Namespace) generateFileBuildTagsLib(f *jen.File, version semver.Version) {
	if versionIsNone(version) {
		return
	}

	versions := []string{}
	for _, v := range ns.versions.versions {
		if versionIsNone(v) {
			continue
		}

		if v.GE(version) {
			versions = append(versions, ns.goPackageName+"_"+versionString(v))
		}
	}

	// pkg_1.20 pkg_1.22 pgk_1.30,...
	//
	// Which will be used like this when building.
	//     pkg_1.20 OR pkg_1.22 OR pgk_1.30 OR ...
	buildTags := strings.Join(versions, " ")

	f.HeaderComment(fmt.Sprintf("+build %s", buildTags))
	f.Line()
}
