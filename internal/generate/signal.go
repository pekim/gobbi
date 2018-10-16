package generate

type Signal struct {
	//Namespace *Namespace

	//Blacklist     bool `xml:"blacklist,attr"`

	Name        string       `xml:"name,attr"`
	When        string       `xml:"when,attr"`
	Version     string       `xml:"version,attr"`
	Doc         *Doc         `xml:"doc"`
	Parameters  Parameters   `xml:"parameters>parameter"`
	ReturnValue *ReturnValue `xml:"return-value"`
}

//func (r *Record) init(ns *Namespace) {
//	r.Namespace = ns
//	r.GoName = r.Name
//	r.newFromCFuncName = fmt.Sprintf("%sNewFromC", r.Name)
//
//	r.Constructors.init(ns, r)
//	r.Methods.init(ns, r)
//	r.Fields.init(ns)
//}
//
//func (r *Record) version() string {
//	return r.Version
//}
//
//func (r *Record) blacklisted() (bool, string) {
//	return r.Blacklist, r.CType
//}
//
//func (r *Record) supported() (supported bool, reason string) {
//	if r.CType == "" {
//		return false, fmt.Sprintf("%s : no CType", r.Name)
//	}
//
//	return true, ""
//}
//
//func (r *Record) mergeAddenda(addenda *Record) {
//	r.Blacklist = addenda.Blacklist
//	r.FieldsPrivate = addenda.FieldsPrivate
//
//	if addenda.Version != "" {
//		r.Version = addenda.Version
//	}
//
//	r.Methods.mergeAddenda(addenda.Methods)
//}
//
//func (r *Record) generate(g *jen.Group, version *Version) {
//	if supportedByVersion(r, version) {
//		r.generateType(g)
//		(&RecordNewFromCFunc{r}).generate(g)
//		(&RecordToCFunc{r}).generate(g)
//		r.generateUpcasts(g)
//		r.generateDowncast(g)
//	}
//
//	r.Constructors.generate(g, version)
//	r.Methods.generate(g, version)
//}
//
//func (r *Record) generateType(g *jen.Group) {
//	g.Commentf("%s is a wrapper around the C record %s.", r.GoName, r.CType)
//
//	g.
//		Type().
//		Id(r.GoName).
//		StructFunc(func(g *jen.Group) {
//			g.
//				Id("native").
//				Op("*").
//				Qual("C", r.CType)
//
//			if r.FieldsPrivate {
//				g.Comment("All fields are private")
//			} else {
//				r.Fields.generate(g)
//			}
//		})
//	g.Line()
//}
//
//func (r *Record) root() *Record {
//	if r.ParentName == "" {
//		return r
//	}
//
//	qname := QNameNew(r.Namespace, r.ParentName)
//	parent, found := qname.ns.recordOrClassRecordForName(qname.name)
//	if !found {
//		panic(fmt.Sprintf("Failed to find parent %s for %s", r.ParentName, r.Name))
//	}
//	return parent.root()
//}
