package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

type Parameter struct {
	Name      string `xml:"name,attr"`
	Direction string `xml:"direction,attr"`
	Argument
	AllowNone bool      `xml:"allow-none,attr"`
	Doc       *Doc      `xml:"doc"`
	Varargs   *struct{} `xml:"varargs"`

	goVarName string
	namespace *Namespace

	lengthParam    *Parameter
	lengthForParam *Parameter
}

func (p *Parameter) init(ns *Namespace) {
	p.namespace = ns

	p.goVarName = makeUnexportedGoName(p.Name)
	if p.goVarName == "object" {
		p.goVarName += "_"
	}

	if p.Direction == "" {
		p.Direction = "in"
	}

	if p.Type != nil {
		//// If the parameter's type is an alias, replace it with the aliased type
		//alias, found := ns.aliasForName(p.Type.Name)
		//if found {
		//	p.initAlias(ns, alias)
		//}

		p.Type.init(ns)
	}

}

func (p Parameter) supported() (bool, string) {
	if !p.isSupported() {
		return false, fmt.Sprintf("parameter '%s' of type '%s' not supported", p.Name, p.Type)
	}

	return true, ""
}

func (p *Parameter) isIn() bool {
	return p.Direction == "in" || p.Direction == "inout"
}

func (p *Parameter) isOut() bool {
	return p.Direction == "out" || p.Direction == "inout"
}

func (p *Parameter) generateInDeclaration(g *jen.Group) {
	if p.Array != nil && strings.HasSuffix(p.Array.CType, "gchar*") {
		g.Id(p.goVarName).String()
		return
	}

	goType, err := p.Type.jenGoType()
	if err != nil {
		panic(err)
	}

	g.Id(p.goVarName).Add(goType)
}

func (p Parameter) generateInArg(g *jen.Group, index int) {
	goVar := jen.Id(p.goVarName)

	if p.Array != nil && strings.HasSuffix(p.Array.CType, "gchar*") {
		// inArgs[<index>].Set...(<goVar>)
		g.
			Id("inArgs").
			Index(jen.Lit(index)).
			Dot("SetString").
			Call(goVar)

		return
	}

	if p.Type.isAlias() {
		typ := p.Type.resolvedType()

		goVar = jen.
			Add(jenGoTypes[typ.Name]).
			Parens(goVar)
	}

	if p.Type.isBitfield() || p.Type.isEnumeration() {
		goVar = jen.
			Add(jenGoTypes["int"]).
			Parens(goVar)
	}

	if p.Type.isRecord() || p.Type.isInterface() {
		goVar = goVar.Dot(nativeAccessorName).Call()
	}

	value := p.Type.createFromInArgument(goVar)

	// inArgs[<index>].Set...(<goVar>)
	g.
		Id("inArgs").
		Index(jen.Lit(index)).
		Dot(p.Type.argumentValueSetFunctionName()).
		Call(value)
}

func (p Parameter) generateOutValue(g *jen.Group, varNamePrefix string, index int) {
	arg := jen.
		Id("outArgs").
		Index(jen.Lit(index))

	name := jen.Id(fmt.Sprintf("%s%d", varNamePrefix, index))

	p.generateValue(g, name, arg)
}

func (p *Parameter) transferOwnershipJen(g *jen.Group) {
	if !p.Type.isString() {
		return
	}

	g.Lit(p.TransferOwnership == "full")
}
