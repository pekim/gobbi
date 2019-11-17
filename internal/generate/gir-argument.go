package generate

import "github.com/dave/jennifer/jen"

type Argument struct {
	TransferOwnership string `xml:"transfer-ownership,attr"`
	Nullable          bool   `xml:"nullable,attr"`
	Type              *Type  `xml:"type"`
	//Array             *Array `xml:"array"`
}

func (a *Argument) generateValue() *jen.Statement {
	return jen.
		Dot(a.Type.argumentValueGetFunctionName()).
		CallFunc(a.transferOwnershipJen)
}

func (a *Argument) transferOwnershipJen(g *jen.Group) {
	if a.Type.Name != "utf8" {
		return
	}

	g.Lit(a.TransferOwnership == "full")
}
