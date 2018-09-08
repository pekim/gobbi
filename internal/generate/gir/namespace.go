package gir

type Namespace struct {
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
}

func (ns *Namespace) fixup() {
	ns.Aliases.fixup(ns)
	ns.Bitfields.fixup(ns)
	ns.Callbacks.fixup(ns)
	ns.Classes.fixup(ns)
	ns.Constants.fixup(ns)
	ns.Enumerations.fixup(ns)
	ns.Functions.fixup(ns)
	ns.Records.fixup(ns)
}

func (ns *Namespace) mergeAddenda(addenda *Namespace) {
	if addenda != nil {
		ns.mergeAddendaFunctions(addenda)
	}
}

func (ns *Namespace) mergeAddendaFunctions(addenda *Namespace) {
	// for _, function := range addenda.Functions {
	// 	fmt.Println("addenda", function)
	// }
}
