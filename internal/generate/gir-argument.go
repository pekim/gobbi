package generate

import "github.com/dave/jennifer/jen"

type Argument struct {
	TransferOwnership string `xml:"transfer-ownership,attr"`
	Nullable          bool   `xml:"nullable,attr"`
	Type              *Type  `xml:"type"`
	//Array             *Array `xml:"array"`
}

func (a *Argument) generateValue(arg *jen.Statement) {
	arg.
		Dot(a.Type.argumentValueGetFunctionName()).
		CallFunc(a.transferOwnershipJen)
}

func (a *Argument) transferOwnershipJen(g *jen.Group) {
	if a.Type.Name != "utf8" {
		return
	}

	g.Lit(a.TransferOwnership == "full")
}

func (a Argument) supportedAsOutParameter() bool {
	if a.Type == nil || a.Type.Name == "none" {
		// return type is void
		return true
	}

	if _, ok := argumentGetFunctionNames[a.Type.Name]; ok {
		return true
	}

	return false
}

func (a Argument) supportedAsInParameter() bool {
	if a.Type == nil || a.Type.Name == "" {
		return false
	}

	if _, ok := argumentSetFunctionNames[a.Type.Name]; ok {
		return true
	}

	return false
}
