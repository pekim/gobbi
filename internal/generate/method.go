package generate

import "github.com/dave/jennifer/jen"

type Method struct {
	*Function
}

func (m *Method) init(ns *Namespace, record *Record) {
	m.Function.init(ns, record)
}

func (m *Method) generate(g *jen.Group, version *Version) {
	supported, reason := m.supported()
	if !supported {
		g.Commentf("Unsupported : %s", reason)
		g.Line()
		return
	}

	if !supportedByVersion(m, version) {
		return
	}

	m.Function.generate(g, version)
}
