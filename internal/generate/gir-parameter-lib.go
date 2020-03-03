package generate

import (
	"github.com/dave/jennifer/jen"
)

func (p *Parameter) libParamGoType() *jen.Statement {

	if p.isType() {
		star := ""
		if p.Nullable && !p.isOut() && !p.Type.isStruct() && !p.Type.isPointer() {
			// nullable simple type, so make it a pointer
			star = "*"
		}

		return jen.Op(star).Add(p.Type.libParamGoType(false))
	}

	//star := ""
	//if p.isOut() {
	//	star = "*"
	//}
	if p.isArray() {
		//return jen.
		//	Op(star).
		//	Add(p.Array.sysParamGoType())
	}

	panic("TODO")
}

func (p *Parameter) generateLibArg(g *jen.Group, varName string) {
	if p.isType() {
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

	//if p.isArray() {
	//	p.generateSysCArgArray(g, goVarName, cVarName)
	//	return
	//}
}
