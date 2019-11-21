package generate

import "github.com/dave/jennifer/jen"

type Enumeration struct {
	Name         string  `xml:"name,attr"`
	Version      string  `xml:"version,attr"`
	CType        string  `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	GlibTypeName string  `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GlibGetType  string  `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	Members      Members `xml:"member"`
	Doc          *Doc    `xml:"doc"`
	//Functions Functions `xml:"function"`

	namespace      *Namespace
	goTypeName     string
	goMemberPrefix string
}

func (e *Enumeration) init(ns *Namespace) {
	e.namespace = ns
	e.goTypeName = makeExportedGoName(e.Name)
	e.goMemberPrefix = e.goTypeName

	for _, member := range e.Members {
		member.init(ns)
	}
}

func (e *Enumeration) generate(f *file) {
	f.docForC(e.goTypeName, e.CType)
	f.docVersion(e.Version)

	// define the type
	f.
		Type().
		Id(e.goTypeName).
		Int()

	// generate the member consts
	f.
		Const().
		DefsFunc(func(g *jen.Group) {
			for _, member := range e.Members {
				member.generate(g, e.goTypeName, e.goMemberPrefix)
			}
		})
}
