package generate

import (
	"github.com/dave/jennifer/jen"
)

func (p *Parameter) libParamGoType() *jen.Statement {

	if p.Type != nil {
		return jen.Add(p.Type.libParamGoType(p.isOut()))
	}

	//star := ""
	//if p.isOut() {
	//	star = "*"
	//}
	if p.Array != nil {
		//return jen.
		//	Op(star).
		//	Add(p.Array.sysParamGoType())
	}

	panic("TODO")
}

func (p *Parameter) generateLibArg(g *jen.Group, varName string) {
	if p.Type != nil {
		argValue := jen.Id(p.goVarName)

		if p.Type.isStruct() {
			// A few strange cases in GObject.
			if p.Type.CType != "gpointer" {
				argValue = jen.Id(p.goVarName).Dot("ToC").Call()
			}
		}

		if p.Type.isBitfield() || p.Type.isEnumeration() {
			argValue = jen.Parens(jen.Op(p.Type.cType.stars).Int()).Parens(argValue)
		}

		g.Id(varName).Op(":=").Add(argValue)
	}

	//if p.Array != nil {
	//	p.generateSysCArgArray(g, goVarName, cVarName)
	//	return
	//}
}
