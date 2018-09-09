package gir

import (
	"github.com/dave/jennifer/jen"
)

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
	f.GoName = makeExportedGoName(f.Name)
	f.Parameters.init(ns)

	if f.ReturnValue != nil {
		f.ReturnValue.init(ns)
	}
}

func (f *Function) version() string {
	return f.Version
}

func (f *Function) mergeAddenda(addenda *Function) {
	f.Blacklist = addenda.Blacklist

	if addenda.GoName != "" {
		f.GoName = addenda.GoName
	}
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

	if supported, reason := f.Parameters.allSupported(); !supported {
		g.Commentf("Unsupported function: %s : %s", f.CIdentifier, reason)
		g.Line()
		return
	}

	g.Commentf("%s is a wrapper around the C function %s.", f.GoName, f.CIdentifier)

	g.
		Func().
		Id(f.GoName).
		ParamsFunc(f.Parameters.generateFunctionDeclaration).
		BlockFunc(func(g *jen.Group) {
		}).
		Line()
}
