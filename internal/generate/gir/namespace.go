package gir

import (
	"fmt"
)

type Namespace struct {
	Name                string         `xml:"name,attr"`
	Version             string         `xml:"version,attr"`
	SharedLibrary       string         `xml:"shared-library,attr"`
	CIdentifierPrefixes string         `xml:"http://www.gtk.org/introspection/c/1.0 identifier-prefixes,attr"`
	CSymbolPrefixes     string         `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefixes,attr"`
	Aliases             []*Alias       `xml:"alias"`
	Callbacks           []*Callback    `xml:"callback"`
	Classes             []*Class       `xml:"class"`
	Bitfields           []*Enumeration `xml:"bitfield"`
	Enumerations        []*Enumeration `xml:"enumeration"`
	Functions           []*Function    `xml:"function"`
	Constants           []*Constant    `xml:"constant"`
	Records             []*Record      `xml:"record"`
}

func (ns *Namespace) fixup() {
	ns.fixupAliases()
}

func (ns *Namespace) mergeAddenda(addenda *Namespace) {
	ns.mergeAddendaFunctions(addenda)
}

func (ns *Namespace) mergeAddendaFunctions(addenda *Namespace) {
	for _, function := range addenda.Functions {
		fmt.Println("addenda", function)
	}
}

func (ns *Namespace) fixupAliases() {
	for _, alias := range ns.Aliases {
		alias.fixup(ns)
	}

	for _, callback := range ns.Callbacks {
		callback.fixup(ns)
	}

	for _, class := range ns.Classes {
		class.fixup(ns)
	}

	for _, bitfield := range ns.Bitfields {
		bitfield.fixup(ns)
	}

	for _, enum := range ns.Enumerations {
		enum.fixup(ns)
	}

	for _, function := range ns.Functions {
		function.fixup(ns)
	}

	for _, constant := range ns.Constants {
		constant.fixup(ns)
	}

	for _, record := range ns.Records {
		record.fixup(ns)
	}

	for _, class := range ns.Classes {
		class.fixup(ns)
	}
}
