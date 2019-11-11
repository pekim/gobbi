package generate

type Namespace struct {
	Blacklist           bool   `xml:"blacklist,attr"`
	Name                string `xml:"name,attr"`
	Version             string `xml:"version,attr"`
	SharedLibrary       string `xml:"shared-library,attr"`
	CDocPath            string `xml:"c-doc-path,attr"`
	CIdentifierPrefixes string `xml:"http://www.gtk.org/introspection/c/1.0 identifier-prefixes,attr"`
	CSymbolPrefixes     string `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefixes,attr"`
	//	Aliases                       Aliases      `xml:"alias"`
	//	Bitfields                     Enumerations `xml:"bitfield"`
	//	Callbacks                     Callbacks    `xml:"callback"`
	//	Classes                       Classes      `xml:"class"`
	//	Constants                     Constants    `xml:"constant"`
	//	Enumerations                  Enumerations `xml:"enumeration"`
	//	Functions                     Functions    `xml:"function"`
	//	Records                       Records      `xml:"record"`
	//	Interfaces                    Interfaces   `xml:"interface"`
	//	GenerateGobjectclassGotypeMap bool         `xml:"generate-gobjectclass-gotype-map,attr"`
}
