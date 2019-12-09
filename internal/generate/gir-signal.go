package generate

import (
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

	if supported, reason := s.ReturnValue.supported(); !supported {
		return false, reason
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
		ParamsFunc(s.generateHandlerParam).
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

func (s *Signal) generateHandlerParam(g *jen.Group) {
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
}

//func (recv *Widget) ConnectShow(handler func(widget *Widget)) int {
//	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
//		value0 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
//		widget := WidgetNewFromNative(value0.GetObject().Native())
//
//		handler(widget)
//	}
//
//	return callback.ConnectSignal(recv.Native(), "show", marshal)
//}
