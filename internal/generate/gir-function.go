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
	Doc               *Doc   `xml:"doc"`
	//InstanceParameter *Parameter   `xml:"parameters>instance-parameter"`
	Parameters     Parameters   `xml:"parameters>parameter"`
	ReturnValue    *ReturnValue `xml:"return-value"`
	Throws         int          `xml:"throws,attr"`
	Introspectable string       `xml:"introspectable,attr"`

	namespace      *Namespace
	goName         string
	record         *Record
	invokerVarName string
	receiver       bool
}

func (f *Function) init(ns *Namespace, record *Record, receiver bool) {
	f.namespace = ns
	f.Parameters.init(ns)
	f.ReturnValue.init(ns)
	f.goName = makeExportedGoName(f.Name)
	f.record = record
	f.receiver = receiver

	if record != nil {
		f.invokerVarName = makeUnexportedGoName(f.Name, false) + record.Name + "Invoker"
	} else {
		f.invokerVarName = makeUnexportedGoName(f.Name, false) + "Invoker"
	}
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
		Func().                                          // "func"
		Do(f.generateReceiver).                          // (recv)
		Id(f.goName).                                    // function-name
		ParamsFunc(f.Parameters.generateInDeclarations). // (in params)
		ParamsFunc(func(g *jen.Group) {                  // (return value, out params)
			f.ReturnValue.generateDeclaration(g)       // return value
			f.Parameters.generateReturnDeclarations(g) // out params
		}).
		BlockFunc(f.generateBody) // { body }
	fi.Line()
}

func (f *Function) generateReceiver(s *jen.Statement) {
	if !f.receiver {
		return
	}

	s.Params(jen.
		Id("recv").
		Op("*").
		Id(f.record.goName),
	)
}

func (f *Function) generateBody(g *jen.Group) {
	f.generateInitialiseInvoker(g)
	g.Line()
	f.Parameters.generateInArgs(g, f.receiver)
	g.Line()
	f.Parameters.generateOutArgs(g)
	g.Line()
	f.generateCallFunction(g)
	g.Line()
	f.generateReturnVar(g)
	f.generateOutArgValues(g)
	g.Line()
	f.generateReturn(g)
}

func (f *Function) generateInitialiseInvoker(g *jen.Group) {
	g.
		If(jen.Id(f.invokerVarName).Op("==").Nil()).
		Block(jen.
			Id(f.invokerVarName).
			Op("=").
			Do(func(s *jen.Statement) {
				if f.record != nil {
					s.Qual(gi.PackageName, "StructFunctionInvokerNew").
						Call(
							jen.Lit(f.namespace.Name),
							jen.Lit(f.record.Name),
							jen.Lit(f.Name),
						)
				} else {
					s.Qual(gi.PackageName, "FunctionInvokerNew").
						Call(
							jen.Lit(f.namespace.Name),
							jen.Lit(f.Name),
						)
				}
			}))
}

func (f *Function) generateCallFunction(g *jen.Group) {
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
		CallFunc(func(i *jen.Group) {
			f.Parameters.generateCallParams(i, f.receiver)
		})
}

func (f *Function) generateReturnVar(g *jen.Group) {
	if f.ReturnValue.isVoid() {
		return
	}

	g.
		Id("retGo").
		Op(":=").
		Do(func(s *jen.Statement) {
			f.ReturnValue.generateValue(s, jen.Id("ret"))
		})
}

func (f *Function) generateOutArgValues(g *jen.Group) {
	f.Parameters.generateOutValues(g, "out")
}

func (f *Function) generateReturn(g *jen.Group) {
	if f.ReturnValue.isVoid() && f.Parameters.outCount() == 0 {
		return
	}

	g.ReturnFunc(func(g *jen.Group) {
		if !f.ReturnValue.isVoid() {
			g.Id("retGo")
		}

		for i := 0; i < f.Parameters.outCount(); i++ {
			g.Id(fmt.Sprintf("out%d", i))
		}
	})
}
