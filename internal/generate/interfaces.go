package generate

import (
	"github.com/dave/jennifer/jen"
)

type Interfaces []*Interface

func (ii Interfaces) init(ns *Namespace) {
	for _, i := range ii {
		i.init(ns)
	}
}

func (ii Interfaces) versionList() Versions {
	var versions Versions

	for _, i := range ii {
		if i.Version != "" {
			versions = append(versions, VersionNew(i.Version))
		}

		for _, ctor := range i.Constructors {
			if ctor.Version != "" {
				versions = append(versions, VersionNew(ctor.Version))
			}
		}
		for _, m := range i.Methods {
			if m.Version != "" {
				versions = append(versions, VersionNew(m.Version))
			}
		}
	}

	return versions
}

func (ii Interfaces) entities() []Generatable {
	var generatables []Generatable

	for _, record := range ii {
		generatables = append(generatables, record)
	}

	return generatables
}

func (ii Interfaces) generate(g *jen.Group, version *Version) {
	for _, i := range ii {
		i.generate(g, version)
	}
}

func (ii Interfaces) forName(name string) *Interface {
	for _, iface := range ii {
		if iface.Name == name {
			return iface
		}
	}

	return nil
}

func (ii Interfaces) mergeAddenda(addenda Interfaces) {
	for _, addendaInterface := range addenda {
		if iface := ii.forName(addendaInterface.Name); iface != nil {
			iface.Record.mergeAddenda(addendaInterface.Record)
		}
	}
}
