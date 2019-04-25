package generate

type Enumerations []*Enumeration

func (ee Enumerations) init(ns *Namespace) {
	for _, enum := range ee {
		enum.init(ns)
	}
}

func (ee Enumerations) versionList() Versions {
	var versions Versions

	for _, e := range ee {
		if e.Version != "" {
			versions = append(versions, VersionNew(e.Version))
		}
	}

	return versions
}

func (ee Enumerations) entities() []Generatable {
	var generatables []Generatable

	for _, enum := range ee {
		generatables = append(generatables, enum)
	}

	return generatables
}

func (ee Enumerations) forName(name string) *Enumeration {
	for _, enum := range ee {
		if enum.Name == name {
			return enum
		}
	}

	return nil
}

func (ee Enumerations) mergeAddenda(addenda Enumerations) {
	for _, addendaEnum := range addenda {
		if enum := ee.forName(addendaEnum.Name); enum != nil {
			enum.mergeAddenda(addendaEnum)
		}
	}
}

func (ee Enumerations) needCgo() bool {
	return false
}
