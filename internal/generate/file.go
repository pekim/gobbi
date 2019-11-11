package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type file struct {
	*jen.File
}

func (f *file) unsupported(format string, args ...interface{}) {
	text := fmt.Sprintf(format, args...)
	f.Commentf("UNSUPPORTED : %s", text)
}
