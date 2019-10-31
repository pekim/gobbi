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
		Qual("reflect", "Value").
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

	d[jen.Lit(c.GlibTypeName)] = jen.
		Qual("reflect", "ValueOf").
		Call(jen.Id(c.GoName + "NewFromC"))
}
