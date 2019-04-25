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

		for _, ctor := range c.Constructors {
			if ctor.Version != "" {
				versions = append(versions, VersionNew(ctor.Version))
			}
		}
		for _, m := range c.Methods {
			if m.Version != "" {
				versions = append(versions, VersionNew(m.Version))
			}
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

func (cc Classes) generateDocs(ns *Namespace, typeName string) {}

func (cc Classes) needCgo() bool {
	return true
}
