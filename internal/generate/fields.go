package generate

import "github.com/pekim/jennifer/jen"

type Fields []*Field

func (cc Fields) init(ns *Namespace) {
	for _, f := range cc {
		f.init(ns)
	}
}

func (ff Fields) generate(g *jen.Group) {
	for _, f := range ff {
		f.generate(g)
	}
}

func (ff Fields) forName(name string) *Field {
	for _, field := range ff {
		if field.Name == name {
			return field
		}
	}

	return nil
}

func (ff Fields) mergeAddenda(addenda Fields) {
	for _, addendaField := range addenda {
		if field := ff.forName(addendaField.Name); field != nil {
			field.mergeAddenda(addendaField)
		}
	}
}
