package generate

type Implementss []*Implements

func (ii Implementss) mergeAddenda(addenda Implementss) {
	for _, addendaImplements := range addenda {
		if implements := ii.forName(addendaImplements.Name); implements != nil {
			implements.mergeAddenda(addendaImplements)
		}
	}
}

func (ii Implementss) forName(name string) *Implements {
	for _, implements := range ii {
		if implements.Name == name {
			return implements
		}
	}

	return nil
}

// -------------------

type Implements struct {
	Name    string `xml:"name,attr"`
	Version string `xml:"version,attr"`
}

func (i *Implements) version() string {
	return i.Version
}

func (i *Implements) mergeAddenda(addenda *Implements) {
	if addenda.Version != "" {
		i.Version = addenda.Version
	}
}
