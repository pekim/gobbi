package generate

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

type CInclude struct {
	Name string `xml:"name,attr"`
}

type CIncludes []*CInclude

func (cc CIncludes) generate(f *jen.File) {
	for _, cInclude := range cc {
		f.CgoPreamble(fmt.Sprintf("#include <%s>", cInclude.Name))
	}
}
