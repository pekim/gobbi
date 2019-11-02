package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

//type classAncestors []struct {
//	methodName           string
//	interfaceMethodNames []string
//}
//
//var GobjectClassGoTypeMap = map[string]struct {
//	ctor                 reflect.Value
//	interfaceMethodNames []string
//	ancestors            classAncestors
//}{
//	"GtkAboutDialog": {
//		ctor:                 reflect.ValueOf(gtk.AboutDialogNewFromC),
//		interfaceMethodNames: []string{},
//		ancestors: classAncestors{
//			{
//				methodName:           "Widget",
//				interfaceMethodNames: []string{},
//			},
//		},
//	},
//}

type gobjectClassToGoTypeMetaMap struct {
	ancestorsTypeName string
	ns                *Namespace
	file              *jen.File
	version           *Version
}

func (ns *Namespace) generateGobjectClassToGoTypeMetaMap(file *jen.File, version Version) {
	gobjectClassToGoTypeMetaMap{
		ancestorsTypeName: "classAncestors",
		ns:                ns,
		file:              file,
		version:           &version,
	}.generate()
}

func (m gobjectClassToGoTypeMetaMap) generate() {
	if !m.ns.GenerateGobjectclassGotypeMap {
		return
	}

	m.generateClassAncestorsType()
	m.generateMap()
	m.file.Line()
}

func (m gobjectClassToGoTypeMetaMap) generateClassAncestorsType() {
	m.file.
		Type().
		Id("classAncestors").
		Index().
		Struct(
			jen.Id("MethodName").String(),
			jen.Id("InterfaceMethodNames").Index().String(),
		)
}

func (m gobjectClassToGoTypeMetaMap) generateMap() {
	m.file.
		Var().
		Id("GobjectClassToGoTypeMetaMap").
		Op("=").
		Map(
			jen.String()).
		Struct(
			jen.Id("Ctor").Qual("reflect", "Value"),
			jen.Id("InterfaceMethodNames").Index().String(),
			jen.Id("Ancestors").Id("classAncestors")).
		Values(
			jen.DictFunc(func(d jen.Dict) {
				for _, class := range m.ns.Classes {
					m.generateClass(class, d)
				}
			}))
}

func (m gobjectClassToGoTypeMetaMap) generateClass(class *Class, d jen.Dict) {
	blacklisted, _ := class.blacklisted()
	supported, _ := class.supported()

	if blacklisted || !supported || !supportedByVersion(class, m.version) {
		return
	}

	d[jen.Lit(class.GlibTypeName)] = jen.Values(jen.Dict{
		jen.Id("Ctor"): jen.
			Qual("reflect", "ValueOf").
			Call(jen.Id(class.GoName + "NewFromC")),

		jen.Id("Ancestors"): jen.
			Id(m.ancestorsTypeName).
			ValuesFunc(func(g *jen.Group) {
				m.generateClassAncestors(class, g)
			}),
	})
}

func (m gobjectClassToGoTypeMetaMap) generateClassAncestors(class *Class, g *jen.Group) {
	qname := QNameNew(class.Namespace, class.ParentName)

	parent, found := qname.namespace.recordOrClassRecordForName(qname.name)
	if !found {
		panic(fmt.Sprintf("Failed to find parent %s for %s", class.ParentName, class.Name))
	}

	if parent.Blacklist {
		return
	}

	previousAncestor := parent
	ancestorName := previousAncestor.ParentName

	for ancestorName != "" {
		qname = QNameNew(qname.namespace, ancestorName)

		ancestor, found := qname.namespace.recordOrClassRecordForName(qname.name)
		if !found {
			panic(fmt.Sprintf("Failed to find parent %s for %s", ancestorName, previousAncestor.Name))
		}

		if ancestor.Blacklist {
			return
		}

		if qname.namespace.goPackageName != m.ns.goPackageName {
			break
		}

		g.Values(jen.DictFunc(func(d jen.Dict) {
			d[jen.Id("MethodName")] = jen.Lit(qname.name)
		}))

		previousAncestor := ancestor
		ancestorName = previousAncestor.ParentName
	}
}
