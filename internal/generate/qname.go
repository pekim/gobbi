package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

type QName struct {
	sameNamespace bool
	ns            *Namespace
	name          string
}

func QNameNew(referingNamespace *Namespace, name string) *QName {
	parts := strings.Split(name, ".")

	switch len(parts) {
	case 1:
		return &QName{
			sameNamespace: true,
			ns:            referingNamespace,
			name:          name,
		}
	case 2:
		ns := referingNamespace.get(parts[0])
		if ns == nil {
			panic(fmt.Sprintf("Failed to find namespace for '%s'", name))
		}

		return &QName{
			sameNamespace: ns == referingNamespace,
			ns:            ns,
			name:          parts[1],
		}
	default:
		panic(fmt.Sprintf("Unsupported type name '%s'", name))
	}

}

func (q *QName) generate(s *jen.Statement) {
	if q.sameNamespace {
		s.Id(q.name)
	} else {
		s.Qual(q.ns.fullGoPackageName, q.name)
	}
}
