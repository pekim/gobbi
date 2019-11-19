package generate

type Method struct {
	*Function
}

func (m *Method) init(ns *Namespace, record *Record) {
	m.Function.init(ns, record, true)
}
