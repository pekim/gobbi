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
	ns.fixupBitfields()
	ns.fixupCallbacks()
	ns.fixupClasses()
	ns.fixupConstants()
	ns.fixupEnumerations()
	ns.fixupFunctions()
	ns.fixupRecords()
}

func (ns *Namespace) fixupAliases() {
	for _, alias := range ns.Aliases {
		alias.fixup(ns)
	}
}

func (ns *Namespace) fixupBitfields() {
	for _, bitfield := range ns.Bitfields {
		bitfield.fixup(ns)
	}
}

func (ns *Namespace) fixupCallbacks() {
	for _, callback := range ns.Callbacks {
		callback.fixup(ns)
	}
}

func (ns *Namespace) fixupClasses() {
	for _, class := range ns.Classes {
		class.fixup(ns)
	}
}

func (ns *Namespace) fixupConstants() {
	for _, constant := range ns.Constants {
		constant.fixup(ns)
	}
}

func (ns *Namespace) fixupEnumerations() {
	for _, enum := range ns.Enumerations {
		enum.fixup(ns)
	}
}

func (ns *Namespace) fixupFunctions() {
	for _, function := range ns.Functions {
		function.fixup(ns)
	}
}

func (ns *Namespace) fixupRecords() {
	for _, record := range ns.Records {
		record.fixup(ns)
	}
}

func (ns *Namespace) mergeAddenda(addenda *Namespace) {
	ns.mergeAddendaFunctions(addenda)
}

func (ns *Namespace) mergeAddendaFunctions(addenda *Namespace) {
	for _, function := range addenda.Functions {
		fmt.Println("addenda", function)
	}
}
