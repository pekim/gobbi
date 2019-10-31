package generate

import (
	"github.com/dave/jennifer/jen"
)

const gobjectClassGoTypeMapVarName = "GobjectClassGoTypeMap"

func (ns *Namespace) generateGobjectClassGoTypeMap(file *jen.File, version Version) {
	if !ns.GenerateGobjectclassGotypeMap {
		return
	}

	file.
		Var().
		Id(gobjectClassGoTypeMapVarName).
		Op("=").
		Map(
			jen.String()).
		Index().
		Qual("reflect", "Type").
		Values(
			jen.DictFunc(func(d jen.Dict) {
				for _, class := range ns.Classes {
					class.generateAddToGobjectClassGoTypeMap(d, version)
				}
			}))

	file.Line()
}

func (c *Class) generateAddToGobjectClassGoTypeMap(d jen.Dict, version Version) {
	//ns := c.Record.Namespace
	blacklisted, _ := c.blacklisted()
	supported, _ := c.supported()

	if blacklisted || !supported || !supportedByVersion(c, &version) {
		return
	}

	reflectType := jen.
		Qual("reflect", "TypeOf").
		Call(jen.
			Parens(jen.
				Op("*").
				Id(c.GoName)).
			//Id(c.GlibTypeName)).
			Parens(jen.Nil()))

	d[jen.Lit(c.GlibTypeName)] =
		jen.
			Index().
			Qual("reflect", "Type").
			Values(reflectType)
}
