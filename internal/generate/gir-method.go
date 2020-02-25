package generate

type Method struct {
	*Function
}

func (m *Method) init(context *context, ns *Namespace, record *Record) {
	m.Function.init(context, ns, record, true)
}
