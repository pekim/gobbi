package generate

import (
	"strings"
)

type Namespace struct {
	Blacklist           bool         `xml:"blacklist,attr"`
	Name                string       `xml:"name,attr"`
	Version             string       `xml:"version,attr"`
	SharedLibrary       string       `xml:"shared-library,attr"`
	CIdentifierPrefixes string       `xml:"http://www.gtk.org/introspection/c/1.0 identifier-prefixes,attr"`
	CSymbolPrefixes     string       `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefixes,attr"`
	Aliases             Aliases      `xml:"alias"`
	Bitfields           Enumerations `xml:"bitfield"`
	Callbacks           Callbacks    `xml:"callback"`
	Classes             Classes      `xml:"class"`
	Constants           Constants    `xml:"constant"`
	Enumerations        Enumerations `xml:"enumeration"`
	Functions           Functions    `xml:"function"`
	Records             Records      `xml:"record"`

	repo          *Repository
	goPackageName string
	allVersions   Versions
	versionDebug  bool
	libDir        string
}

func (ns *Namespace) init(repo *Repository) {
	ns.repo = repo
	ns.goPackageName = strings.ToLower(ns.Name)
	ns.libDir = projectFilepath("lib", ns.goPackageName)

	ns.Aliases.init(ns)
	ns.Bitfields.init(ns)
	ns.Callbacks.init(ns)
	ns.Classes.init(ns)
	ns.Constants.init(ns)
	ns.Enumerations.init(ns)
	ns.Functions.init(ns)
	ns.Records.init(ns)

	ns.setAllVersions()
}

func (ns *Namespace) mergeAddenda(addenda *Namespace) {
	if addenda != nil {
		ns.Aliases.mergeAddenda(addenda.Aliases)
		ns.Bitfields.mergeAddenda(addenda.Bitfields)
		ns.Constants.mergeAddenda(addenda.Constants)
		ns.Enumerations.mergeAddenda(addenda.Enumerations)
		ns.Functions.mergeAddenda(addenda.Functions)
		ns.Records.mergeAddenda(addenda.Records)
	}
}

func (ns *Namespace) blacklisted() bool {
	return ns.Blacklist
}

func (ns *Namespace) generate() {
	ns.generateLibDir()
	ns.generatePackageFile()
	ns.generateGeneratables("alias", ns.Aliases)
	ns.generateGeneratables("bitfield", ns.Bitfields)
	ns.generateGeneratables("constant", ns.Constants)
	ns.generateGeneratables("enum", ns.Enumerations)
	ns.generateGeneratables("function", ns.Functions)
}
