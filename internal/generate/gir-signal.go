package generate

type Signal struct {
	Namespace *Namespace

	Name        string       `xml:"name,attr"`
	When        string       `xml:"when,attr"`
	Version     string       `xml:"version,attr"`
	Doc         *Doc         `xml:"doc"`
	Parameters  Parameters   `xml:"parameters>parameter"`
	ReturnValue *ReturnValue `xml:"return-value"`

	record *Record
}

func (s *Signal) init(ns *Namespace, record *Record) {
	s.Namespace = ns
	s.record = record

	s.Parameters.init(ns)

	if s.ReturnValue != nil {
		s.ReturnValue.init(ns)
	}
}

func (s *Signal) generate(f *file) {
}

//func (recv *Widget) ConnectShow(handler func(widget *Widget)) int {
//	return callback.ConnectSignal(recv.Native(), "show",
//		func(return_value *callback.Value, paramValues []callback.Value) {
//			value0 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
//			widget := WidgetNewFromNative(value0.GetObject().Native())
//
//			handler(widget)
//		},
//	)
//}
