package generate

import (
	"fmt"
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

	namespace      *Namespace
	goName         string
	invokerVarName string
}

func (f *Function) init(ns *Namespace /*receiver *Record, namePrefix string*/) {
	f.namespace = ns
	f.ReturnValue.init(ns)
	f.goName = makeExportedGoName(f.Name)
	f.invokerVarName = makeUnexportedGoName(f.Name, false) + "Invoker"
}

func (f *Function) generate(fi *file) {
	if len(f.Parameters) > 0 {
		fi.unsupported(f.CIdentifier, "has parameters")
		return
	}

	if !f.ReturnValue.Type.supportedAsReturnValue() {
		fi.unsupported(f.CIdentifier, "return type '%s' not supported", f.ReturnValue.Type.Name)
		return
	}

	f.generateInvokerVar(fi)

	fi.docForC(f.goName, f.CIdentifier)
	fi.
		Func().
		Id(f.goName).
		Add(paramsFunc(f.generateParamsDeclaration)).
		Add(paramsFunc(f.generateReturnDeclaration)).
		Add(blockFunc(f.generateBody))
}

func (f *Function) generateInvokerVar(fi *file) {
	fi.
		Var().
		Id(f.invokerVarName).
		Op("*").
		Qual(gi.PackageName, "Function")
}

func (f *Function) generateParamsDeclaration(g *group) {
}

func (f *Function) generateReturnDeclaration(g *group) {
	if f.ReturnValue.isVoid() {
		return
	}

	goType, err := f.ReturnValue.Type.jenGoType()
	if err != nil {
		fmt.Println(err)
		return
	}
	g.Add(goType)
}

func (f *Function) generateBody(g *group) {
	// initialise invoker
	g.
		If(jen.Id(f.invokerVarName).Op("==").Nil()).
		Block(jen.
			Id(f.invokerVarName).
			Op("=").
			Qual(gi.PackageName, "FunctionInvokerNew").
			Call(
				jen.Lit(f.namespace.Name),
				jen.Lit(f.Name),
			))
	g.Line()

	// call the function
	g.
		Do(func(s *jen.Statement) {
			if !f.ReturnValue.isVoid() {
				s.
					Id("ret").
					Op(":=")
			}
		}).
		Id(f.invokerVarName).
		Dot("Invoke").
		Call()

	// marshall and return any return value
	if !f.ReturnValue.isVoid() {
		g.Return(
			jen.
				Id("ret").
				Dot(f.ReturnValue.Type.returnValueExtractFunctionName()).
				Call())
	}
}
