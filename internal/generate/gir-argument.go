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

func (a *Argument) generateValue(g *jen.Group, argName *jen.Statement, argVar *jen.Statement) {
	a.Type.generateOutArgValue(g, argName, argVar, a.transferOwnership())
}

func (a *Argument) transferOwnership() *bool {
	if !a.Type.isString() {
		return nil
	}

	to := a.TransferOwnership == "full"
	return &to
}

func (a Argument) supportedAsOutParameter() bool {
	if a.Type == nil || a.Type.Name == "none" {
		// return type is void
		return true
	}

	if a.Type.isAlias() {
		alias, _ := a.Type.namespace.Aliases.byName(a.Type.Name)
		if !strings.HasPrefix(alias.Type.Name, "g") {
			// Alias that is not a simple type is not yet supported.
			return false
		}
	}

	return a.isSupported()
}

func (a Argument) isSupported() bool {
	if !a.Type.isValid() {
		return false
	}

	typ := a.Type.resolvedType()

	if !typ.isValid() {
		return false
	}

	if _, err := typ.jenGoType(); err == nil {
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
