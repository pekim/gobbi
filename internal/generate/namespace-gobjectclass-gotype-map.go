package generate

import (
	"github.com/dave/jennifer/jen"
)

const gobjectClassGoTypeMapVarName = "GobjectClassGoTypeMap"

//type classAncestors struct {
//	methodName           string
//	interfaceMethodNames []string
//}
//
//var GobjectClassGoTypeMap = map[string]struct {
//	ctor      reflect.Value
//	ancestors []classAncestors
//}{
//	"GtkAboutDialog": {
//		ctor: reflect.ValueOf(gtk.AboutDialogNewFromC),
//		ancestors: []classAncestors{
//			{
//				methodName:           "Widget",
//				interfaceMethodNames: []string{},
//			},
//		},
//	},
//}

func (ns *Namespace) generateGobjectClassGoTypeMap(file *jen.File, version Version) {
	if !ns.GenerateGobjectclassGotypeMap {
		return
	}

	ancestorsType := jen.
		Index().
		Struct(
			jen.Id("methodName").String(),
			jen.Id("interfaceMethodNames").Index().String(),
		)

	file.
		Var().
		Id(gobjectClassGoTypeMapVarName).
		Op("=").
		Map(
			jen.String()).
		Struct(
			jen.Id("ctor").Qual("reflect", "Value"),
			jen.Id("ancestors").Add(ancestorsType)).
		Values(
			jen.DictFunc(func(d jen.Dict) {
				for _, class := range ns.Classes {
					class.generateAddToGobjectClassGoTypeMap(d, version)
				}
			}))

	file.Line()
}

//	"GtkAboutDialog": {
//		ctor: reflect.ValueOf(gtk.AboutDialogNewFromC),
//		ancestors: []classAncestors{
//			{
//				methodName:           "Widget",
//				interfaceMethodNames: []string{},
//			},
//		},
//	},
func (c *Class) generateAddToGobjectClassGoTypeMap(d jen.Dict, version Version) {
	//ns := c.Record.Namespace
	blacklisted, _ := c.blacklisted()
	supported, _ := c.supported()

	if blacklisted || !supported || !supportedByVersion(c, &version) {
		return
	}

	d[jen.Lit(c.GlibTypeName)] = jen.
		Values(
			jen.Dict{
jen.Lit("ctor"):jen.
	Qual("reflect", "ValueOf").
	Call(jen.Id(c.GoName + "NewFromC")),
			})

			jen.
				Id("ctor").
		Qual("reflect", "ValueOf").
		Call(jen.Id(c.GoName + "NewFromC")),

		jen.
				Id("ancestors")
			)
		//Qual("reflect", "ValueOf").
		//Call(jen.Id(c.GoName + "NewFromC"))
}
