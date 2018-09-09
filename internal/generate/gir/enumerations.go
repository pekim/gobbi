package gir

type Enumerations []*Enumeration

func (ee Enumerations) fixup(ns *Namespace) {
	for _, enum := range ee {
		enum.fixup(ns)
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
