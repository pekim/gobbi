package generate

type Aliases []*Alias

func (aa Aliases) init(ns *Namespace) {
	for _, alias := range aa {
		alias.init(ns)
	}
}

func (aa Aliases) versionList() Versions {
	return Versions{}
}

func (aa Aliases) entities() []Generatable {
	var generatables []Generatable

	for _, alias := range aa {
		generatables = append(generatables, alias)
	}

	return generatables
}

func (aa Aliases) forName(name string) *Alias {
	for _, alias := range aa {
		if alias.Name == name {
			return alias
		}
	}

	return nil
}

func (aa Aliases) mergeAddenda(addenda Aliases) {
	for _, addendaAlias := range addenda {
		if alias := aa.forName(addendaAlias.Name); alias != nil {
			alias.mergeAddenda(addendaAlias)
		}
	}
}

func (aa Aliases) generateDocs(ns *Namespace, typeName string) {}
