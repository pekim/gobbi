package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

type Namespace struct {
	Blacklist           bool         `xml:"blacklist,attr"`
	Name                string       `xml:"name,attr"`
	Version             string       `xml:"version,attr"`
	SharedLibrary       string       `xml:"shared-library,attr"`
	CDocPath            string       `xml:"c-doc-path,attr"`
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
	Interfaces          Interfaces   `xml:"interface"`

	repo              *Repository
	callFile          *CallFile
	jenFile           *jen.File
	goPackageName     string
	fullGoPackageName string
	allVersions       Versions
	versionDebug      bool
	libDir            string
	namespaces        map[string]*Namespace
}

func (ns *Namespace) init(repo *Repository, callFile *CallFile) {
	ns.repo = repo
	ns.callFile = callFile
	ns.goPackageName = strings.ToLower(ns.Name)
	ns.fullGoPackageName = fmt.Sprintf("github.com/pekim/gobbi/lib/%s", ns.goPackageName)
	ns.libDir = projectFilepath("lib", ns.goPackageName)

	if ns.CDocPath == "" {
		ns.CDocPath = ns.goPackageName
	}

	ns.Aliases.init(ns)
	ns.Bitfields.init(ns)
	ns.Callbacks.init(ns)
	ns.Classes.init(ns)
	ns.Constants.init(ns)
	ns.Enumerations.init(ns)
	ns.Functions.init(ns, "")
	ns.Records.init(ns)
	ns.Interfaces.init(ns)

	ns.setAllVersions()
}

func (ns *Namespace) mergeAddenda(addenda *Namespace) {
	if addenda != nil {
		ns.Blacklist = addenda.Blacklist

		if addenda.CDocPath != "" {
			ns.CDocPath = addenda.CDocPath
		}

		ns.Aliases.mergeAddenda(addenda.Aliases)
		ns.Bitfields.mergeAddenda(addenda.Bitfields)
		ns.Classes.mergeAddenda(addenda.Classes)
		ns.Constants.mergeAddenda(addenda.Constants)
		ns.Enumerations.mergeAddenda(addenda.Enumerations)
		ns.Functions.mergeAddenda(addenda.Functions)
		ns.Interfaces.mergeAddenda(addenda.Interfaces)
		ns.Records.mergeAddenda(addenda.Records)
	}
}

func (ns *Namespace) blacklisted() bool {
	return ns.Blacklist
}

func (ns *Namespace) generate() {
	if ns.Blacklist {
		return
	}

	fmt.Printf("%-10s %s\n", ns.Name, ns.Version)

	ns.generateLibDir()

	//ns.generatePackageFile()
	//ns.generateTemplatedFiles()

	//ns.generateBooleanFile()
	ns.generateGeneratables("alias", ns.Aliases)
	ns.generateGeneratables("bitfield", ns.Bitfields)
	//ns.generateGeneratables("class", ns.Classes)
	ns.generateGeneratables("constant", ns.Constants)
	ns.generateGeneratables("enum", ns.Enumerations)
	ns.generateGeneratables("function", ns.Functions)
	//ns.generateGeneratables("interface", ns.Interfaces)
	ns.generateGeneratables("record", ns.Records)
}

func (ns *Namespace) aliasForName(name string) (*Alias, bool) {
	alias := ns.Aliases.forName(name)
	return alias, alias != nil
}

func (ns *Namespace) bitfieldForName(name string) (*Enumeration, bool) {
	bitfield := ns.Bitfields.forName(name)
	return bitfield, bitfield != nil
}

func (ns *Namespace) enumForName(name string) (*Enumeration, bool) {
	enum := ns.Enumerations.forName(name)
	return enum, enum != nil
}

func (ns *Namespace) recordForName(name string) (*Record, bool) {
	record := ns.Records.forName(name)
	return record, record != nil
}

func (ns *Namespace) classForName(name string) (*Class, bool) {
	class := ns.Classes.forName(name)
	return class, class != nil
}

func (ns *Namespace) recordOrClassRecordForName(name string) (*Record, bool) {
	if record, found := ns.recordForName(name); found {
		return record, found
	}

	class, found := ns.classForName(name)
	if found {
		return class.Record, found
	}

	return nil, false
}

func (ns *Namespace) interfaceForName(name string) (*Interface, bool) {
	iface := ns.Interfaces.forName(name)
	return iface, iface != nil
}

func (ns *Namespace) get(name string) *Namespace {
	for _, namespace := range ns.namespaces {
		if namespace.Name == name {
			return namespace
		}
	}

	return nil
}
