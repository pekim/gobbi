package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

type QName struct {
	referringNamespace *Namespace
	namespace          *Namespace
	sameNamespace      bool
	name               string
	originalName       string
}

func QNameNew(referingNamespace *Namespace, name string) *QName {
	parts := strings.Split(name, ".")

	switch len(parts) {
	case 1:
		// A name in the same namespace as the referring namespace.
		return &QName{
			referringNamespace: referingNamespace,
			namespace:          referingNamespace,
			sameNamespace:      true,
			name:               name,
			originalName:       name,
		}
	case 2:
		// A name in a differenct namespace to the referring namespace.
		ns := referingNamespace.get(parts[0])
		if ns == nil {
			panic(fmt.Sprintf("Failed to find namespace for '%s'", name))
		}

		return &QName{
			referringNamespace: referingNamespace,
			namespace:          ns,
			sameNamespace:      ns == referingNamespace,
			name:               parts[1],
			originalName:       name,
		}
	default:
		panic(fmt.Sprintf("Unsupported type name '%s'", name))
	}

}

func (q *QName) generate(s *jen.Statement) {
	if q.sameNamespace {
		s.Id(q.name)
	} else {
		s.Qual(q.namespace.fullGoPackageName, q.name)
	}
}

func (q *QName) String() string {
	return fmt.Sprintf("%s.%s (%s)", q.namespace.goPackageName, q.name, q.originalName)
}
