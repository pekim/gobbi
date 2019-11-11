package generate

import "github.com/dave/jennifer/jen"

type file struct {
	*jen.File
}

func (f *file) unsupported(s string) {
	f.Commentf("UNSUPPORTED : %s", s)
}
