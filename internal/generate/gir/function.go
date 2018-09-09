package gir

import "github.com/dave/jennifer/jen"

type Function struct {
	Namespace *Namespace

	Name              string       `xml:"name,attr"`
	Blacklist         bool         `xml:"blacklist,attr"`
	GoName            string       `xml:"goname,attr"`
	Version           string       `xml:"version,attr"`
	CIdentifier       string       `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	Deprecated        int          `xml:"deprecated,attr"`
	DeprecatedVersion string       `xml:"deprecated-version,attr"`
	Doc               *Doc         `xml:"doc"`
	InstanceParameter *Parameter   `xml:"parameters>instance-parameter"`
	Parameters        Parameters   `xml:"parameters>parameter"`
	ReturnValue       *ReturnValue `xml:"return-value"`
	Throws            int          `xml:"throws,attr"`
	Introspectable    string       `xml:"introspectable,attr"`
}

func (f *Function) init(ns *Namespace) {
	f.Namespace = ns
	f.Parameters.init(ns)

	if f.ReturnValue != nil {
		f.ReturnValue.init(ns)
	}
}

func (f *Function) blacklisted() bool {
	return f.Blacklist
}

func (f *Function) introspectable() bool {
	return f.Introspectable != "0"
}

func (f *Function) version() string {
	return f.Version
}

func (f *Function) mergeAddenda(addenda *Function) {
	f.Blacklist = addenda.Blacklist
	f.GoName = addenda.GoName
}

func (f *Function) generate(g *jen.Group, version *Version) {
	if !supportedByVersion(f, version) {
		return
	}

	if f.Blacklist {
		g.Commentf("Blacklisted function: %s", f.CIdentifier)
		g.Line()
		return
	}

}
