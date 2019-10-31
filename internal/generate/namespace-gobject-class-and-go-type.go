package generate

import "github.com/dave/jennifer/jen"

func (ns *Namespace) generateGobjectClassGoTypeMap(file *jen.File) {
	if !ns.GenerateGobjectclassToGotypeMap {
		return
	}

	file.
		Var().
		Id("gobjectClassGoTypeMap").
		Op("=").
		Make(
			jen.Map(jen.String()).Qual("reflect", "Type"))

	file.Line()
}

func (c *Class) generateAddToGobjectClassGoTypeMap(g *jen.Group, version *Version) {
	if !c.Record.Namespace.GenerateGobjectclassToGotypeMap {
		return
	}
	if !supportedByVersion(c, version) {
		return
	}

	g.Commentf("AddToGobjectClassGoTypeMap : %s", c.GlibTypeName)

	g.Line()
}
