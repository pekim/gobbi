package generate

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

func (r *Record) init(ns *Namespace) {
	r.Namespace = ns

	for _, ctor := range r.Constructors {
		ctor.init(ns)
	}

	for _, method := range r.Methods {
		method.init(ns)
	}
}

func (r *Record) blacklisted() bool {
	return r.Blacklist
}

type Class struct {
	*Record
}

type Constructor struct {
	*Function
}

type Method struct {
	*Function
}
