package gir

import (
	"github.com/dave/jennifer/jen"
	"strings"
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
	f.setGoName()
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

func (f *Function) setGoName() {
	cParts := strings.Split(f.Name, "_")
	goParts := []string{}

	for _, cPart := range cParts {
		goPart := strings.Title(strings.ToLower(cPart))
		goParts = append(goParts, goPart)
	}

	f.GoName = strings.Join(goParts, "")
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

	g.
		Func().
		Id(f.GoName).
		ParamsFunc(func(g *jen.Group) {
		}).
		BlockFunc(func(g *jen.Group) {
		}).
		Line()
}
