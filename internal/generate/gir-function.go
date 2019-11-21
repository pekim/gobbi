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

	funcInfoGoName        string
	funcInfoOnceGoName    string
	funcInfoSetFuncGoName string
}

func (f *Function) init(ns *Namespace, record *Record, receiver bool) {
	f.namespace = ns
	f.Parameters.init(ns)
	f.ReturnValue.init(ns)
	f.goName = makeExportedGoName(f.Name)
	f.record = record
	f.receiver = receiver

	recordName := ""
	if record != nil {
		recordName = record.Name
	}
	f.funcInfoGoName = makeUnexportedGoName(fmt.Sprintf("%s%sFunction",
		makeExportedGoName(recordName),
		makeExportedGoName(f.Name)))
	f.funcInfoOnceGoName = fmt.Sprintf("%s_Once", f.funcInfoGoName)
	f.funcInfoSetFuncGoName = fmt.Sprintf("%s_Set", f.funcInfoGoName)

	if record != nil {
		f.invokerVarName = makeUnexportedGoName(f.Name + record.Name + "Invoker")
	} else {
		f.invokerVarName = makeUnexportedGoName(f.Name + "Invoker")
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

	f.generateFuncInfo(fi)
	f.generateInvokerVar(fi)
	f.generateFunction(fi)
}

func (f *Function) generateFuncInfo(fi *file) {
	//var somefunctionFunc *gi.Function
	fi.
		Var().
		Id(f.funcInfoGoName).
		Op("*").
		Qual(gi.PackageName, "Function")

	// var somefunctionFuncOnce sync.Once
	fi.
		Var().
		Id(f.funcInfoOnceGoName).
		Qual("sync", "Once")

	// func somefunctionFuncSet() {
	//   ...
	// }
	fi.
		Func().
		Id(f.funcInfoSetFuncGoName).
		Params().
		BlockFunc(func(g *jen.Group) {
			//   somefunctionFuncOnce.Do(func() {
			//     somefunctionFunc = gi.FunctionInvokerNew("Gdk", "SomeFunction")
			//   })
			g.
				Id(f.funcInfoOnceGoName).
				Dot("Do").
				Call(jen.
					Func().
					Params().
					Block(jen.
						Id(f.funcInfoGoName).
						Op("=").
						Qual(gi.PackageName, "FunctionInvokerNew").
						Call(jen.Lit(f.namespace.Name),
							jen.Lit(f.Name))))
		})
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
		Id(f.funcInfoSetFuncGoName).
		Call()
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
		Id(f.funcInfoGoName).
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
