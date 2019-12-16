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
	if a.Array != nil && strings.HasSuffix(a.Array.CType, "gchar*") {
		g.Add(argName).Op(":=").Add(argVar).Dot("String").Call(jen.Lit(*a.transferOwnership()))
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

	return a.isSupported()
}

func (a Argument) isSupported() bool {
	if a.Array != nil && strings.HasSuffix(a.Array.CType, "gchar*") {
		return true
	}

	if !a.Type.isValid() {
		return false
	}

	if _, err := a.Type.jenGoType(); err == nil {
		return true
	}

	return false
}

func (a Argument) generateReturnDeclaration(g *jen.Group) {
	if a.Array != nil && strings.HasSuffix(a.Array.CType, "gchar*") {
		g.String()
		return
	}

	goType, err := a.Type.jenGoType()
	if err != nil {
		panic(err)
	}
	g.Add(goType)
}
