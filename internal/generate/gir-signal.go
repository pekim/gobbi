package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/pekim/gobbi/internal/cgo/callback"
)

type Signal struct {
	Namespace *Namespace

	Name        string       `xml:"name,attr"`
	When        string       `xml:"when,attr"`
	Version     string       `xml:"version,attr"`
	Doc         *Doc         `xml:"doc"`
	Parameters  Parameters   `xml:"parameters>parameter"`
	ReturnValue *ReturnValue `xml:"return-value"`

	goName string
	record *Record
}

func (s *Signal) init(ns *Namespace, record *Record) {
	s.Namespace = ns
	s.record = record
	s.goName = makeExportedGoName(s.Name)

	s.Parameters.init(ns)

	if s.ReturnValue != nil {
		s.ReturnValue.init(ns)
	}
}

func (s *Signal) supported() (bool, string) {
	if supported, reason := s.Parameters.supported(); !supported {
		return false, reason
	}

	if s.Parameters.outCount() > 0 {
		return false, "has out params"
	}

	if supported, reason := s.ReturnValue.supported(); !supported {
		return false, reason
	}

	if s.ReturnValue.Type.isQualifiedName() && s.ReturnValue.Type.foreignNamespace.Name == "cairo" {
		return false, "return value is from package cairo"
	}

	return true, ""
}

func (s *Signal) generate(f *file) {
	if supported, reason := s.supported(); !supported {
		f.unsupported(s.Name, reason)
		return
	}

	methodName := "Connect" + s.goName

	f.Commentf(`%s connects a callback to the '%s' signal of the %s.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
`, methodName, s.Name, s.record.goName)

	marshalName := "marshal"

	// GEN: func (recv *Device) ConnectShow() int {...}
	s.record.methodPrelude(f, methodName).
		ParamsFunc(s.generateHandlerParameter).
		Params(jen.Int()).
		BlockFunc(func(g *jen.Group) {
			s.generateMarshalFunc(g, marshalName)

			g.Line()

			// GEN: return callback.ConnectSignal(recv.Native(), "show", marshal)
			g.
				Return().
				Qual(callback.PackageName, "ConnectSignal").
				Call(
					jen.Id(receiverName).Dot("Native").Call(),
					jen.Lit(s.Name),
					jen.Id(marshalName),
				)
		})

	f.Line()
}

func (s *Signal) generateHandlerParameter(g *jen.Group) {
	g.
		Id("handler").
		Func().
		ParamsFunc(func(g *jen.Group) {
			// GEN: instance *SomeRecord
			g.Id("instance").Op("*").Id(s.record.Name)
			s.Parameters.generateInDeclarations(g)
		}).
		ParamsFunc(func(g *jen.Group) {
			s.ReturnValue.generateDeclaration(g)       // return value
			s.Parameters.generateReturnDeclarations(g) // out params
		})
}

func (s *Signal) generateMarshalFunc(g *jen.Group, marshalName string) {
	g.
		Id(marshalName).
		Op(":=").
		Func().
		Params(
			jen.Id("returnValue").Op("*").Qual(callback.PackageName, "Value"),
			jen.Id("paramValues").Index().Qual(callback.PackageName, "Value"),
		).
		BlockFunc(s.generateMarshalBody)
}

func (s *Signal) generateMarshalBody(g *jen.Group) {
	//if s.ReturnValue.Type.Name != "none" {
	//	g.Commentf("has return : %s", s.ReturnValue.Type.Name)
	//	return
	//}

	s.generateMarshalBodyInstanceParam(g)
	s.generateMarshalBodyInParams(g)
	s.generateMarshalBodyCallHandler(g)
}

func (s *Signal) generateMarshalBodyInstanceParam(g *jen.Group) {
	objectVarName := "objectInstance"

	gobjectNs, _ := s.Namespace.namespaces.byName("GObject")
	var valueNewFromNative *jen.Statement
	if s.Namespace == gobjectNs {
		valueNewFromNative = jen.Id("ValueNewFromNative")
	} else {
		valueNewFromNative = jen.Qual(gobjectNs.goFullPackageName, "ValueNewFromNative")
	}

	// GEN: value0 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
	g.
		Id(objectVarName).
		Op(":=").
		Add(valueNewFromNative).
		Call(jen.
			Qual("unsafe", "Pointer").
			Call(jen.Op("&").Id("paramValues").Index(jen.Lit(0))),
		)

	// GEN: WidgetNewFromNative(value0.GetObject().Native())
	g.
		Id("argInstance").Op(":=").
		Id(s.record.newFromNativeName).
		Call(jen.
			Id(objectVarName).Dot("GetObject").Call().
			Dot("Native").Call())

	g.Line()
}

func (s *Signal) generateMarshalBodyInParams(g *jen.Group) {
	gobjectNs, _ := s.Namespace.namespaces.byName("GObject")
	var valueNewFromNative *jen.Statement
	if s.Namespace == gobjectNs {
		valueNewFromNative = jen.Id("ValueNewFromNative")
	} else {
		valueNewFromNative = jen.Qual(gobjectNs.goFullPackageName, "ValueNewFromNative")
	}

	for p, param := range s.Parameters {
		if !param.isIn() {
			continue
		}

		objectVarName := fmt.Sprintf("object%d", p+1)
		argVarName := fmt.Sprintf("arg%d", p+1)

		// GEN: object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		g.
			Id(objectVarName).
			Op(":=").
			Add(valueNewFromNative).
			Call(jen.
				Qual("unsafe", "Pointer").
				Call(jen.Op("&").Id("paramValues").Index(jen.Lit(p + 1))),
			)

		// GEN: arg1 := ...
		g.Id(argVarName).Op(":=").Do(func(s *jen.Statement) {
			param.generateValueFromObject(s, objectVarName)
		})

		g.Line()
	}
}

func (s *Signal) generateMarshalBodyCallHandler(g *jen.Group) {
	g.
		Do(func(st *jen.Statement) {
			if s.ReturnValue.Type.Name != "none" {
				st.Id("retGo").Op(":=")
			}
		}).
		Id("handler").
		CallFunc(func(g *jen.Group) {
			g.Id("argInstance")

			for p, param := range s.Parameters {
				if !param.isIn() {
					continue
				}

				g.Id(fmt.Sprintf("arg%d", p+1))
			}
		})

	if s.ReturnValue.Type.Name != "none" {
		gobjectNs, _ := s.Namespace.namespaces.byName("GObject")
		var valueNewFromNative *jen.Statement
		if s.Namespace == gobjectNs {
			valueNewFromNative = jen.Id("ValueNewFromNative")
		} else {
			valueNewFromNative = jen.Qual(gobjectNs.goFullPackageName, "ValueNewFromNative")
		}

		// GEN: value0 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		g.
			Id("returnObject").
			Op(":=").
			Add(valueNewFromNative).
			Call(jen.
				Qual("unsafe", "Pointer").
				Call(jen.Id("returnValue")),
			)

		s.ReturnValue.generateObjectFromValue(g, "returnObject", "retGo")
	}
}
