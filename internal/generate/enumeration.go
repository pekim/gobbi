package generate

type Enumeration struct {
	Name         string  `xml:"name,attr"`
	Version      string  `xml:"version,attr"`
	CType        string  `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	GlibTypeName string  `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GlibGetType  string  `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	Members      Members `xml:"member"`
	//Doc          *Doc      `xml:"doc"`
	//Functions Functions `xml:"function"`

	namespace  *Namespace
	goTypeName string
}

func (e *Enumeration) init(ns *Namespace) {
	e.namespace = ns
	e.goTypeName = makeExportedGoName(e.Name)

	//e.goTypeName = e.Name
	//if e.GoTypeName != "" {
	//	e.goTypeName = e.GoTypeName
	//}
	//
	//memberNamePrefix := e.Namespace.CIdentifierPrefixes + "_"
	//for _, member := range e.Members {
	//	member.init(ns, memberNamePrefix)
	//}
	//
	//e.Functions.init(ns, e.Name)
}

func (e *Enumeration) generate(f *file) {
	f.docForC(e.goTypeName, e.Name)
	f.
		Type().
		Id(e.goTypeName).
		Int()
}

//
//func (e *Enumeration) version() string {
//	return e.Version
//}
//
//func (e *Enumeration) mergeAddenda(addenda *Enumeration) {
//	e.Blacklist = addenda.Blacklist
//	e.GoTypeName = addenda.GoTypeName
//	if addenda.CType != "" {
//		e.CType = addenda.CType
//	}
//	if addenda.Version != "" {
//		e.Version = addenda.Version
//	}
//	e.Members.mergeAddenda(addenda.Members)
//	e.Functions.mergeAddenda(addenda.Functions)
//}
//
//func (e *Enumeration) blacklisted() (bool, string) {
//	return e.Blacklist, e.CType
//}
//
//func (e *Enumeration) supported() (supported bool, reason string) {
//	return true, ""
//}
//
//func (e *Enumeration) generate(g *jen.Group, version *Version) {
//	if !supportedByVersion(e, version) {
//		return
//	}
//
//	// define the type
//	g.
//		Type().
//		Id(e.goTypeName).
//		Qual("C", e.CType)
//
//	// define members
//	g.Const().DefsFunc(func(g *jen.Group) {
//		for _, member := range e.Members {
//			member.generate(g, e.goTypeName)
//		}
//	})
//
//	e.Functions.generate(g, version)
//}
