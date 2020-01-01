package generate

import "github.com/dave/jennifer/jen"

func (p *Parameter) libParamGoType() *jen.Statement {

	if p.Type != nil {
		return jen.Add(p.Type.libParamGoType(false))
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
		argValue := p.Type.sysParamGoType(false).Parens(jen.Id(p.goVarName))
		g.Id(varName).Op(":=").Add(argValue)
	}

	//if p.Array != nil {
	//	p.generateSysCArgArray(g, goVarName, cVarName)
	//	return
	//}
}
