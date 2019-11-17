package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type Parameter struct {
	Namespace *Namespace

	Name              string `xml:"name,attr"`
	Direction         string `xml:"direction,attr"`
	TransferOwnership string `xml:"transfer-ownership,attr"`
	Nullable          bool   `xml:"nullable,attr"`
	AllowNone         bool   `xml:"allow-none,attr"`
	//Doc               *Doc      `xml:"doc"`
	Type *Type `xml:"type"`
	//Array             *Array    `xml:"array"`
	Varargs *struct{} `xml:"varargs"`

	goVarName string
}

func (p *Parameter) init(ns *Namespace) {
	p.Namespace = ns
	p.goVarName = makeUnexportedGoName(p.Name, false)

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

	if !p.Type.supportedAsParameter() {
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

func (p *Parameter) generateInDeclaration(g *group) {
	goType, err := p.Type.jenGoType()
	if err != nil {
		panic(err)
	}

	g.
		Id(p.goVarName).
		Add(goType)
}

func (p *Parameter) generateOutDeclaration(g *group) {
	goType, err := p.Type.jenGoType()
	if err != nil {
		panic(err)
	}

	g.Add(goType)
}

func (p Parameter) generateInArg(g *group, index int) {
	g.
		Id("inArgs").
		Index(jen.Lit(index)).
		Dot(p.Type.argumentSetFunctionName()).
		Call(jen.Id(p.goVarName))
}

func (p Parameter) generateOutValue(g *jen.Group, index int) {
	g.
		Id("outArgs").
		Index(jen.Lit(index)).
		Dot(p.Type.argumentValueGetFunctionName()).
		CallFunc(p.transferOwnershipJen)
}

func (p *Parameter) transferOwnershipJen(g *jen.Group) {
	if p.Type.Name != "utf8" {
		return
	}

	g.Lit(p.TransferOwnership == "full")
}
