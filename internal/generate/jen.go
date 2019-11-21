package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type file struct {
	*jen.File
}

func (f *file) unsupported(cName string, format string, args ...interface{}) {
	f.Comment(unsupportedf(cName, format, args...))
	f.Line()
}

func (f *file) docForC(goName, cName string) {
	f.Comment(docForC(goName, cName))
}

func (f *file) docVersion(version string) {
	f.Comment(docVersion(version))
}

func unsupportedf(cName string, format string, args ...interface{}) string {
	text := fmt.Sprintf(format, args...)
	return fmt.Sprintf("// UNSUPPORTED : C value '%s' : %s", cName, text)
}

func docForC(goName, cName string) string {
	return fmt.Sprintf("// %s is a representation of the C type %s.", goName, cName)
}

func docVersion(version string) string {
	if version == "" {
		return ""
	}

	return fmt.Sprintf("//\n// since %s", version)
}
