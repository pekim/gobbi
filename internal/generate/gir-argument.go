package generate

import (
	"github.com/dave/jennifer/jen"
)

type Argument struct {
	TransferOwnership string `xml:"transfer-ownership,attr"`
	Nullable          bool   `xml:"nullable,attr"`
	Type              *Type  `xml:"type"`
	//Array             *Array `xml:"array"`
}

func (a *Argument) generateValue(s *jen.Statement, argVar *jen.Statement) {
	argValue := argVar.
		Dot(a.Type.argumentValueGetFunctionName()).
		CallFunc(a.transferOwnershipJen)

	createFromArgument := a.Type.createFromArgumentFunction()

	if createFromArgument != nil {
		createFromArgument(s, argValue)
	} else {
		s.Add(argValue)
	}
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

	if _, ok := a.Type.namespace.outParameterGeneratorByName(a.Type.Name); ok {
		return true
	}

	return false
}

func (a Argument) supportedAsInParameter() bool {
	typ := a.Type.resolvedType()

	if typ == nil || typ.Name == "" {
		return false
	}

	if _, ok := argumentSetFunctionNames[typ.Name]; ok {
		return true
	}

	return false
}

func (a Argument) generateReturnDeclaration(g *jen.Group) {
	goType, err := a.Type.jenGoType()
	if err != nil {
		panic(err)
	}
	g.Add(goType)
}
