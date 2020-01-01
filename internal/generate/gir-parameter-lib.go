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
