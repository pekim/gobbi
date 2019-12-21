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

/*
analyseName takes a, possibly qualified ("ns.name"), name and returns its
namespace and name.

The returned bool is true if the name is well formed and if qualified it
namespace was found.
The returned namespace is the matching namespace for a qualified name.
The returned string is the name withing the namespace.
*/
func (nn namespaces) analyseName(name string) (bool, *Namespace, string) {
	parts := strings.Split(name, ".")

	if len(parts) == 1 {
		return false, nil, parts[0]
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
