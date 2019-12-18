package generate

type CInclude struct {
	Name    string `xml:"name,attr"`
	Version string `xml:"version,attr"`
}

type CIncludes []*CInclude
