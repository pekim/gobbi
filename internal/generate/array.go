package generate

import (
	"github.com/dave/jennifer/jen"
	"strings"
)

type Array struct {
	Namespace *Namespace

	Type           *Type  `xml:"type"`
	CType          string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	FixedSize      *int   `xml:"fixed-size,attr"`
	Length         *int   `xml:"length,attr"`
	ZeroTerminated string `xml:"zero-terminated,attr"`
}

func (a *Array) init(ns *Namespace) {
	a.Namespace = ns
	a.Type.init(ns)
}

func (a *Array) generateDeclaration(g *jen.Group, goVarName string) {
	a.Type.generator.generateArrayDeclaration(g, goVarName)
}

func (a *Array) generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string) {
	g.
		Id(cVarName).
		Op(":=").
		Op("&").
		Id(goVarName).
		Index(jen.Lit(0))
}

func (a *Array) generateArrayLenParamCVar(g *jen.Group, cVarName string, arrayGoVarName string, ctype string) {
	//fmt.Println(a)
	g.
		Id(cVarName).
		Op(":=").
		Parens(jen.Qual("C", ctype)).
		Parens(jen.Len(jen.Id(arrayGoVarName)))
}

func (a *Array) generateParamCallArgument(g *jen.Group, cVarName string) {
	cType := strings.TrimRight(a.CType, "*")
	indirectLevel := len(a.CType) - len(cType)
	if indirectLevel == 0 {
		indirectLevel = 1
	}
	indirect := strings.Repeat("*", indirectLevel)

	g.
		Parens(jen.Op(indirect).Qual("C", cType)).
		Parens(jen.Qual("unsafe", "Pointer").Call(jen.Id(cVarName)))
}
