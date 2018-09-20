package generate

type Classes []*Class

func (cc Classes) init(ns *Namespace) {
	for _, class := range cc {
		class.init(ns)
	}
}

func (cc Classes) versionList() Versions {
	var versions Versions

	for _, c := range cc {
		if c.Version != "" {
			versions = append(versions, VersionNew(c.Version))
		}
	}

	return versions
}

func (cc Classes) entities() []Generatable {
	var generatables []Generatable

	for _, class := range cc {
		generatables = append(generatables, class)
	}

	return generatables
}

func (cc Classes) forName(name string) *Class {
	for _, class := range cc {
		if class.Name == name {
			return class
		}
	}

	return nil
}

func (cc Classes) mergeAddenda(addenda Classes) {
	for _, addendaClass := range addenda {
		if class := cc.forName(addendaClass.Name); class != nil {
			class.mergeAddenda(addendaClass)
		}
	}
}
