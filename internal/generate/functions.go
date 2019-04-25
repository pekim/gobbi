package generate

import (
	"github.com/dave/jennifer/jen"
)

type Functions []*Function

func (ff Functions) init(ns *Namespace, namePrefix string) {
	for _, function := range ff {
		function.init(ns, nil, namePrefix)
	}
}

func (ff Functions) versionList() Versions {
	var versions Versions

	for _, f := range ff {
		if f.Version != "" {
			versions = append(versions, VersionNew(f.Version))
		}
	}

	return versions
}

func (ff Functions) entities() []Generatable {
	var generatables []Generatable

	for _, function := range ff {
		generatables = append(generatables, function)
	}

	return generatables
}

func (ff Functions) forCIdentifier(cidentifier string) *Function {
	for _, function := range ff {
		if function.CIdentifier == cidentifier {
			return function
		}
	}

	return nil
}

func (ff Functions) mergeAddenda(addenda Functions) {
	for _, addendaFunction := range addenda {
		if function := ff.forCIdentifier(addendaFunction.CIdentifier); function != nil {
			function.mergeAddenda(addendaFunction)
		}
	}
}

func (ff Functions) generate(g *jen.Group, version *Version) {
	for _, fn := range ff {
		fn.generate(g, version)
	}
}

func (ff Functions) generateDocs(ns *Namespace, typeName string) {}

func (ff Functions) needCgo() bool {
	return true
}
