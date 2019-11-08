package generate

import (
	"github.com/pekim/jennifer/jen"
)

type Signals []*Signal

func (ss Signals) init(ns *Namespace, record *Record) {
	for _, signal := range ss {
		signal.init(ns, record)
	}
}

func (ss Signals) generate(g *jen.Group, version *Version, parentVersion string) {
	for _, signal := range ss {
		signal.generate(g, version, parentVersion)
	}
}

func (ss Signals) forName(name string) *Signal {
	for _, signal := range ss {
		if signal.Name == name {
			return signal
		}
	}

	return nil
}

func (ss Signals) mergeAddenda(addenda Signals) {
	for _, addendaSignal := range addenda {
		if field := ss.forName(addendaSignal.Name); field != nil {
			field.mergeAddenda(addendaSignal)
		}
	}
}
