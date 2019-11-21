package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type ReturnValue struct {
	Argument
	Doc *Doc `xml:"doc"`

	namespace *Namespace
}

func (r *ReturnValue) init(ns *Namespace) {
	r.namespace = ns

	if r.Type != nil {
		r.Type.init(ns)
	}

	//if r.Array != nil && r.Array.Type != nil {
	//	r.Array.init(ns)
	//}
}

func (r *ReturnValue) isVoid() bool {
	return r.Type == nil || r.Type.Name == "none"
}

func (r *ReturnValue) supported() (bool, string) {
	if r.supportedAsOutParameter() {
		return true, ""
	}

	return false, fmt.Sprintf("return type '%s' not supported", r.Type.Name)
}

func (r *ReturnValue) generateDeclaration(g *jen.Group) {
	if r.isVoid() {
		return
	}

	r.generateReturnDeclaration(g)
}
