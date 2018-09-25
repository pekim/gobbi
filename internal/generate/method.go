package generate

import (
	"github.com/dave/jennifer/jen"
)

type Method struct {
	*Function
}

func (m *Method) init(ns *Namespace, record *Record) {
	m.Function.init(ns, record)

	if record.Version != "" && m.Version == "" {
		m.Version = record.Version
	}

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

	if blacklisted, detail := m.blacklisted(); blacklisted {
		g.Commentf("Blacklisted : %s", detail)
		g.Line()
		return
	}

	m.Function.generate(g, version)
}
