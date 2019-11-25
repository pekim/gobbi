package generate

import (
	"github.com/dave/jennifer/jen"
	"strings"
)

type Argument struct {
	TransferOwnership string `xml:"transfer-ownership,attr"`
	Nullable          bool   `xml:"nullable,attr"`
	Type              *Type  `xml:"type"`
	//Array             *Array `xml:"array"`
}

func (a *Argument) generateValue(s *jen.Statement, argVar *jen.Statement) {
	typ := a.Type.resolvedType()

	argValue := argVar.
		Dot(typ.argumentValueGetFunctionName()).
		CallFunc(a.transferOwnershipJen)

	if a.Type.isAlias() {
		argValue = jen.
			Id(a.Type.Name).
			Parens(argValue)
	}

	createFromArgument := typ.createFromArgumentFunction()

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

	typ := a.Type.resolvedType()
	if typ != a.Type {
		// do not yet support pointers here
		if strings.HasSuffix(a.Type.CType, "*") {
			return false
		}
	}

	if _, ok := argumentGetFunctionNames[typ.Name]; ok {
		return true
	}

	if _, ok := typ.namespace.outParameterGeneratorByName(typ.Name); ok {
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
