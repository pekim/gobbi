package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
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
}

func (p *Parameter) init(ns *Namespace) {
	p.namespace = ns
	p.goVarName = makeUnexportedGoName(p.Name)

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
	if p.Type == nil {
		return false, fmt.Sprintf("parameter '%s' has no type", p.Name)
	}

	if !p.isSupported() {
		return false, fmt.Sprintf("parameter '%s' of type '%s' not supported", p.Name, p.Type.Name)
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
	goType, err := p.Type.jenGoType()
	if err != nil {
		panic(err)
	}

	g.
		Id(p.goVarName).
		Add(goType)
}

func (p Parameter) generateInArg(g *jen.Group, index int) {
	goVar := jen.Id(p.goVarName)

	if p.Type.isAlias() {
		typ := p.Type.resolvedType()

		goVar = jen.
			Add(jenGoTypes[typ.Name]).
			Parens(goVar)
	}

	if p.Type.isRecord() {
		goVar = goVar.Dot("native")
	}

	g.
		Id("inArgs").
		Index(jen.Lit(index)).
		Dot(p.Type.argumentValueSetFunctionName()).
		Call(goVar)
}

func (p Parameter) generateOutValue(g *jen.Group, varNamePrefix string, index int) {
	arg := jen.
		Id("outArgs").
		Index(jen.Lit(index))

	g.
		Id(fmt.Sprintf("%s%d", varNamePrefix, index)).
		Op(":=").
		Do(func(s *jen.Statement) {
			p.generateValue(s, arg)
		})
}

func (p *Parameter) transferOwnershipJen(g *jen.Group) {
	if !p.Type.isString() {
		return
	}

	g.Lit(p.TransferOwnership == "full")
}
