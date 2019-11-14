package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/pekim/gobbi/internal/gi"
)

type Function struct {
	Name              string `xml:"name,attr"`
	Blacklist         bool   `xml:"blacklist,attr"`
	Version           string `xml:"version,attr"`
	MovedTo           string `xml:"moved-to,attr"`
	CIdentifier       string `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	Deprecated        int    `xml:"deprecated,attr"`
	DeprecatedVersion string `xml:"deprecated-version,attr"`
	//Doc               *Doc         `xml:"doc"`
	//InstanceParameter *Parameter   `xml:"parameters>instance-parameter"`
	Parameters     Parameters   `xml:"parameters>parameter"`
	ReturnValue    *ReturnValue `xml:"return-value"`
	Throws         int          `xml:"throws,attr"`
	Introspectable string       `xml:"introspectable,attr"`

	namespace *Namespace
	goName    string
}

func (f *Function) init(ns *Namespace /*receiver *Record, namePrefix string*/) {
	f.namespace = ns
	f.goName = makeExportedGoName(f.Name)
}

func (f *Function) generate(fi *file) {
	if len(f.Parameters) > 0 || f.ReturnValue != nil {
		fi.unsupported(f.CIdentifier, "non trivial function")
	}

	fi.docForC(f.goName, f.CIdentifier)

	fi.
		Func().
		Id(f.goName).
		Add(paramsFunc(f.generateParams)).
		Add(blockFunc(f.generateBody))
}

func (f *Function) generateParams(g *group) {
}

func (f *Function) generateBody(g *group) {
	g.
		Id("invoker").
		Op(":=").
		Qual(gi.PackageName, "FunctionInvokerNew").
		Call(
			jen.Lit(f.namespace.Name),
			jen.Lit(f.Name),
		)

	g.
		Qual("fmt", "Println").
		Call(jen.Id("invoker"))
}
