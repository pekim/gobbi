package generate

import "github.com/dave/jennifer/jen"

func (ns *Namespace) generateGobjectClassGoTypeMap(file *jen.File) {
	file.
		Var().
		Id("gobjectClassGoTypeMap").
		Op("=").
		Make(
			jen.Map(jen.String()).Qual("reflect", "Type"))

	file.Line()
}
