package generate

import (
	"github.com/dave/jennifer/jen"
)

type Constructor struct {
	*Function
}

func (m *Constructor) init(ns *Namespace, record *Record) {
	m.Function.init(ns, nil)
	m.GoName = record.GoName + makeExportedGoName(m.Name)
}

func (m *Constructor) generate(g *jen.Group, version *Version) {
	supported, reason := m.supported()

	if supported {
		m.Function.generate(g, version)
	} else {
		g.Commentf("Unsupported : %s", reason)
		g.Line()
	}
}
