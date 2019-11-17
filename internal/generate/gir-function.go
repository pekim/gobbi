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
	Doc               *Doc   `xml:"doc"`
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
	f.Parameters.init(ns)
	f.ReturnValue.init(ns)
	f.goName = makeExportedGoName(f.Name)
	f.invokerVarName = makeUnexportedGoName(f.Name, false) + "Invoker"
}

func (f *Function) supported() (bool, string) {
	if supported, reason := f.Parameters.supported(); !supported {
		return false, reason
	}

	if supported, reason := f.ReturnValue.supported(); !supported {
		return false, reason
	}

	return true, ""
}

func (f *Function) generate(fi *file) {
	if supported, reason := f.supported(); !supported {
		fi.unsupported(f.CIdentifier, reason)
		return
	}

	f.generateInvokerVar(fi)
	f.generateFunction(fi)
}

func (f *Function) generateInvokerVar(fi *file) {
	fi.
		Var().
		Id(f.invokerVarName).
		Op("*").
		Qual(gi.PackageName, "Function")
	fi.Line()
}

func (f *Function) generateFunction(fi *file) {
	fi.docForC(f.goName, f.CIdentifier)
	fi.
		Func().                                               // "func"
		Id(f.goName).                                         // function-name
		Add(paramsFunc(f.Parameters.generateInDeclarations)). // (in params)
		Add(paramsFunc(func(g *group) {                       // (return value, out params)
			f.ReturnValue.generateDeclaration(g)    // return value
			f.Parameters.generateOutDeclarations(g) // out params
		})).
		Add(blockFunc(f.generateBody)) // { body }
	fi.Line()
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

	f.Parameters.generateInArgs(g)
	f.Parameters.generateOutArgs(g)

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
		CallFunc(f.Parameters.generateCallParams)
	g.Line()

	// marshall and return any return value
	if !f.ReturnValue.isVoid() || f.Parameters.outCount() > 0 {
		g.ReturnFunc(f.generateReturnValues)
	}
}

func (f *Function) generateReturnValues(g *jen.Group) {
	if !f.ReturnValue.isVoid() {
		g.
			Id("ret").
			Add(f.ReturnValue.generateValue())
	}

	f.Parameters.generateOutValues(g)
}
