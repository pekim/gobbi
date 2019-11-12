package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type file struct {
	*jen.File
}

func (f *file) unsupported(cName string, format string, args ...interface{}) {
	if format != "" {
		text := fmt.Sprintf(format, args...)
		f.Commentf("// UNSUPPORTED : C value '%s' : %s", cName, text)
	} else {
		f.Commentf("// UNSUPPORTED : C value '%s'", cName)
	}

	f.Line()
}

func (f *file) docForC(goName, cName string) {
	f.Commentf("// %s is a representation of the C type %s.", goName, cName)
}
