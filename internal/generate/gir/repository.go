package gir

import (
	"encoding/xml"
)

type Repository struct {
	XMLName   xml.Name   `xml:"repository"`
	Version   string     `xml:"version,attr"`
	CIncludes []CInclude `xml:"http://www.gtk.org/introspection/c/1.0 include"`
	Includes  []Include  `xml:"http://www.gtk.org/introspection/core/1.0 include"`
	Packages  []Package  `xml:"package"`
	Namespace *Namespace `xml:"namespace"`
}

func (r *Repository) Fixup() {
	r.Namespace.fixup()
}

func (r *Repository) MergeAddenda(addenda *Repository) {
	r.Namespace.mergeAddenda(addenda.Namespace)
}

type CInclude struct {
	Name string `xml:"name,attr"`
}

type Package struct {
	Name string `xml:"name,attr"`
}
