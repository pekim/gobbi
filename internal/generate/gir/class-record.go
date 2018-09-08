package gir

type Record struct {
	Namespace *Namespace

	Name           string         `xml:"name,attr"`
	Blacklist      bool           `xml:"blacklist,attr"`
	Version        string         `xml:"version,attr"`
	CSymbolPrefix  string         `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefix,attr"`
	CType          string         `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	ParentName     string         `xml:"parent,attr"`
	GlibTypeName   string         `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GlibGetType    string         `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	GlibTypeStruct string         `xml:"http://www.gtk.org/introspection/glib/1.0 type-struct,attr"`
	Doc            *Doc           `xml:"doc"`
	Constructors   []*Constructor `xml:"constructor"`
	Methods        []*Method      `xml:"method"`
}

func (r *Record) fixup(ns *Namespace) {
	r.Namespace = ns

	for _, ctor := range r.Constructors {
		ctor.fixup(ns)
	}

	for _, method := range r.Methods {
		method.fixup(ns)
	}
}

func (r *Record) blacklisted() bool {
	return r.Blacklist
}

type Records []*Record

func (rs Records) fixup(ns *Namespace) {
	for _, record := range rs {
		record.fixup(ns)
	}
}

type Class struct {
	*Record
}

type Classes Records

func (cs Classes) fixup(ns *Namespace) {
	for _, class := range cs {
		class.fixup(ns)
	}
}

type Constructor struct {
	*Function
}

type Method struct {
	*Function
}
