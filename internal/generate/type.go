package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type Type struct {
	Name  string `xml:"name,attr"`
	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	Namespace *Namespace
}

func (t *Type) jenGoType() (*jen.Statement, string) {
	if t == nil {
		return nil, "missing Type"
	}
	if t.CType == "" {
		return nil, "missing Type.CType"
	}

	goType, ok := numberCTypeMap[t.CType]
	if !ok {
		//panic(fmt.Sprintf("No Go type for %s", t.CType))
		return nil, fmt.Sprintf("No Go type for '%s'\n", t.CType)
	}

	return jen.Id(goType), ""
}
