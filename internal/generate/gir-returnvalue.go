package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type ReturnValue struct {
	Namespace *Namespace

	TransferOwnership string `xml:"transfer-ownership,attr"`
	Nullable          bool   `xml:"nullable,attr"`
	//Doc               *Doc   `xml:"doc"`
	Type *Type `xml:"type"`
	//Array             *Array `xml:"array"`
}

func (r *ReturnValue) init(ns *Namespace) {
	r.Namespace = ns

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

func (r *ReturnValue) transferOwnershipJen(g *jen.Group) {
	if r.Type.Name != "utf8" {
		return
	}

	g.Lit(r.TransferOwnership == "full")
}

func (r *ReturnValue) supported() (bool, string) {
	if r.Type.supportedAsReturnValue() {
		return true, ""
	}
	return false, fmt.Sprintf("return type '%s' not supported", r.Type.Name)
}
