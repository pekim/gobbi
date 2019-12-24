package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type CInclude struct {
	Name    string `xml:"name,attr"`
	Version string `xml:"version,attr"`
}

type CIncludes []*CInclude

func (cc CIncludes) generate(f *jen.File) { //, version Version) {
	for _, cInclude := range cc {
		f.CgoPreamble(fmt.Sprintf("#include <%s>", cInclude.Name))
	}

	f.CgoPreamble(fmt.Sprintf("#include <%s>", "stdlib.h"))
}
