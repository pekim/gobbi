package generate

import (
	"github.com/dave/jennifer/jen"
	"strings"
)

type Argument struct {
	TransferOwnership string `xml:"transfer-ownership,attr"`
	Nullable          bool   `xml:"nullable,attr"`
	Type              *Type  `xml:"type"`
	Array             *Array `xml:"array"`
}

func (a *Argument) generateValue(g *jen.Group, argName *jen.Statement, argVar *jen.Statement) {
	if a.Array != nil && strings.HasSuffix(a.Array.CType, "char*") {
		g.Add(argName).Op(":=").Add(argVar).Dot("String").Call(jen.Lit(*a.transferOwnership()))
		return
	}

	if a.Array != nil && strings.HasSuffix(a.Array.CType, "char***") {
		g.Add(argName).Op(":=").Add(argVar).Dot("StringArray").Call(jen.Lit(*a.transferOwnership()))
		return
	}

	a.Type.generateOutArgValue(g, argName, argVar, a.transferOwnership())
}

func (a *Argument) transferOwnership() *bool {
	if a.Type != nil && !a.Type.isString() {
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

	return a.isSupported(false, true)
}

func (a Argument) isSupported(in bool, out bool) bool {
	if a.Array != nil && a.Array.isSupported(in, out) {
		return true
	}

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
	if a.Array != nil && strings.HasSuffix(a.Array.CType, "char*") {
		g.String()
		return
	}

	if a.Array != nil && strings.HasSuffix(a.Array.CType, "char***") {
		g.Index().String()
		return
	}

	goType, err := a.Type.jenGoType()
	if err != nil {
		panic(err)
	}
	g.Add(goType)
}
