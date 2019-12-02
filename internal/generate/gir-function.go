package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/pekim/gobbi/internal/gi"
)

const receiverName = "recv"

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

	namespace *Namespace
	goName    string
	record    *Record
	receiver  bool

	funcInfoGoName        string
	funcInfoOnceGoName    string
	funcInfoSetFuncGoName string
}

func (f *Function) init(ns *Namespace, record *Record, receiver bool) {
	f.namespace = ns
	f.Parameters.init(ns)
	f.ReturnValue.init(ns)
	f.record = record
	f.receiver = receiver

	f.goName = makeExportedGoName(f.Name)
	if f.namespace.Name == "GObject" && f.goName == "ParamSpecOverride" {
		// Avoid name clash with a class of the same name.
		f.goName += "_"
	}

	recordName := ""
	if record != nil {
		recordName = record.Name
	}
	f.funcInfoGoName = makeUnexportedGoName(fmt.Sprintf("%s%sFunction",
		makeExportedGoName(recordName),
		makeExportedGoName(f.Name)))
	f.funcInfoOnceGoName = fmt.Sprintf("%s_Once", f.funcInfoGoName)
	f.funcInfoSetFuncGoName = fmt.Sprintf("%s_Set", f.funcInfoGoName)
}

func (f *Function) supported() (bool, string) {
	if f.Name == "" && f.MovedTo != "" {
		return false, fmt.Sprintf("moved to %s", f.MovedTo)
	}

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

	// func somefunctionFuncSet() error {
	//   ...
	// }
	fi.
		Func().
		Id(f.funcInfoSetFuncGoName).
		Params().
		List(jen.Id("error")).
		BlockFunc(func(g *jen.Group) {
			// var err error
			g.Var().Id("err").Id("error")

			// somefunctionFuncOnce.Do(func() {
			//   ...
			// })
			g.
				Id(f.funcInfoOnceGoName).
				Dot("Do").
				Call(jen.
					Func().
					Params().
					BlockFunc(func(g *jen.Group) {
						if f.record != nil {
							// err = someStruct_Set()
							g.
								Id("err").
								Op("=").
								Id(f.record.structInfoSetFuncGoName).
								Call()

							// if err != nil {
							//   return
							// }
							g.
								If(jen.Id("err").Op("!=").Nil()).
								Block(jen.Return())

							// somefunctionFunc, err = someStruct.InvokerNew("SomeFunction")
							g.
								Id(f.funcInfoGoName).
								Op(",").
								Id("err").
								Op("=").
								Id(f.record.structInfoGoName).
								Dot("InvokerNew").
								Call(jen.Lit(f.Name))
						} else {
							// somefunctionFunc, err = gi.FunctionInvokerNew("Gdk", "SomeFunction")
							g.
								Id(f.funcInfoGoName).
								Op(",").
								Id("err").
								Op("=").
								Qual(gi.PackageName, "FunctionInvokerNew").
								Call(jen.Lit(f.namespace.Name),
									jen.Lit(f.Name))
						}
					}))

			g.Return().Id("err")
		})
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

	s.Add(generateReceiverParams(f.record))
}

func (f *Function) generateBody(g *jen.Group) {
	f.Parameters.generateInArgs(g, f.receiver)
	g.Line()
	f.Parameters.generateOutArgs(g)
	f.generateReturnArg(g)
	g.Line()
	f.generateInitialiseInvoker(g)
	f.generateCallFunction(g)
	g.Line()
	f.generateReturnVar(g)
	f.generateOutArgValues(g)
	g.Line()
	f.generateReturn(g)
}

func (f *Function) generateInitialiseInvoker(g *jen.Group) {
	g.
		Id("err").
		Op(":=").
		Id(f.funcInfoSetFuncGoName).
		Call()
}

func (f *Function) generateReturnArg(g *jen.Group) {
	if f.ReturnValue.isVoid() {
		return
	}

	g.
		Var().
		Id("ret").
		Qual(gi.PackageName, "Argument")
}

func (f *Function) generateCallFunction(g *jen.Group) {
	g.
		If(jen.Id("err").Op("==").Nil()).
		BlockFunc(func(g *jen.Group) {
			g.Do(func(s *jen.Statement) {
				if !f.ReturnValue.isVoid() {
					s.
						Id("ret").
						Op("=")
				}
			}).
				Id(f.funcInfoGoName).
				Dot("Invoke").
				CallFunc(func(i *jen.Group) {
					f.Parameters.generateCallParams(i, f.receiver)
				})
		})
}

func (f *Function) generateReturnVar(g *jen.Group) {
	if f.ReturnValue.isVoid() {
		return
	}

	value := f.ReturnValue.generateValue(jen.Id("ret"))

	g.
		Id("retGo").
		Op(":=").
		Add(value)
}

func (f *Function) generateOutArgValues(g *jen.Group) {
	f.Parameters.generateOutValues(g, "out")
}

func (f *Function) generateReturn(g *jen.Group) {
	g.ReturnFunc(func(g *jen.Group) {
		if !f.ReturnValue.isVoid() {
			g.Id("retGo")
		}

		for i := 0; i < f.Parameters.outCount(); i++ {
			g.Id(fmt.Sprintf("out%d", i))
		}
	})
}
