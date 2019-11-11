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

func (t *Type) jenGoType() *jen.Statement {
	if t == nil || t.CType == "" {
		fmt.Println("No Ctype for", t)
		return nil
	}

	goType, ok := numberCTypeMap[t.CType]
	if !ok {
		//panic(fmt.Sprintf("No Go type for %s", t.CType))
		fmt.Printf("No Go type for %s\n", t.CType)
		return nil
	}

	return jen.Id(goType)
}
