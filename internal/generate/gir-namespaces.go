package generate

import (
	"fmt"
	"strings"
)

type namespaces map[string]*Namespace

func (nn namespaces) byName(name string) (*Namespace, bool) {
	for _, ns := range nn {
		if ns.Name == name {
			return ns, true
		}
	}

	return nil, false
}

func (nn namespaces) analyseName(name string) (bool, *Namespace, string) {
	parts := strings.Split(name, ".")

	if len(parts) == 1 {
		return false, nil, ""
	}

	if len(parts) != 2 {
		panic(fmt.Sprintf("Unsupported Type name %s", name))
	}

	ns, found := nn.byName(parts[0])
	if !found {
		panic(fmt.Sprintf("Namespace %s referenced in Type name %s not found", parts[0], name))
	}
	foreignNamespace := ns
	foreignName := parts[1]

	return true, foreignNamespace, foreignName
}
