package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type file struct {
	*jen.File
}

func (f *file) unsupported(cname string, format string, args ...interface{}) {
	if format != "" {
		text := fmt.Sprintf(format, args...)
		f.Commentf("UNSUPPORTED : C value '%s' : %s", cname, text)
	} else {
		f.Commentf("UNSUPPORTED : C value '%s'", cname)
	}
}
