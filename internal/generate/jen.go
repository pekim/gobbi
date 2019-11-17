package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type file struct {
	*jen.File
	group *group
}

func (f *file) unsupported(cName string, format string, args ...interface{}) {
	f.group.unsupported(cName, format, args...)
}

func (f *file) docForC(goName, cName string) {
	f.group.docForC(goName, cName)
}

func (f *file) docVersion(version string) {
	f.group.docVersion(version)
}

type group struct {
	*jen.Group
}

func (g *group) unsupported(cName string, format string, args ...interface{}) {
	if format != "" {
		text := fmt.Sprintf(format, args...)
		g.Commentf("// UNSUPPORTED : C value '%s' : %s", cName, text)
	} else {
		g.Commentf("// UNSUPPORTED : C value '%s'", cName)
	}

	g.Line()
}

func (g *group) docForC(goName, cName string) {
	g.Commentf("// %s is a representation of the C type %s.", goName, cName)
}

func (g *group) docVersion(version string) {
	if version == "" {
		return
	}

	g.Comment("//")
	g.Commentf("// since %s", version)
}

func blockFunc(fn func(g *group)) *jen.Statement {
	return jen.BlockFunc(func(jg *jen.Group) {
		fn(&group{jg})
	})
}

func defsFunc(fn func(g *group)) *jen.Statement {
	return jen.DefsFunc(func(jg *jen.Group) {
		fn(&group{jg})
	})
}

func paramsFunc(fn func(g *group)) *jen.Statement {
	return jen.ParamsFunc(func(jg *jen.Group) {
		fn(&group{jg})
	})
}

func structFunc(fn func(g *group)) *jen.Statement {
	return jen.StructFunc(func(jg *jen.Group) {
		fn(&group{jg})
	})
}
