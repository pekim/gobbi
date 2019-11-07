package generate

import (
	"fmt"
	"github.com/pekim/jennifer/jen"
)

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
		Id(m.ancestorsTypeName).
		Index().
		Struct(
			jen.Id("className").String(),
			jen.Id("methodName").String(),
		)
}

func (m gobjectClassToGoTypeMetaMap) generateMap() {
	m.file.
		Var().
		Id("gobjectClassToGoTypeMetaMap").
		Op("=").
		Map(
			jen.String()).
		Struct(
			jen.Id("ctor").Qual("reflect", "Value"),
			jen.Id("interfaceMethodNames").Index().String(),
			jen.Id("ancestors").Id(m.ancestorsTypeName)).
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
		jen.Id("ctor"): jen.
			Qual("reflect", "ValueOf").
			Call(jen.Id(class.GoName + "NewFromC")),

		jen.Id("ancestors"): jen.
			Id(m.ancestorsTypeName).
			ValuesFunc(func(g *jen.Group) {
				m.generateClassAncestors(class, g)
			}),

		jen.Id("interfaceMethodNames"): jen.
			Index().
			String().
			ValuesFunc(func(g *jen.Group) {
				m.generateInterfaces(class, g)
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

		// Currently the map is only created for the gtk package.
		// Once an ancestor from a different package is reached, stop.
		// This is probably okay, as all widgets are descended from a
		// gobject or atk class.
		if qname.namespace.goPackageName != m.ns.goPackageName {
			break
		}

		g.Values(jen.DictFunc(func(d jen.Dict) {
			d[jen.Id("className")] = jen.Lit(ancestor.CType)
			d[jen.Id("methodName")] = jen.Lit(qname.name)
		}))

		previousAncestor := ancestor
		ancestorName = previousAncestor.ParentName
	}
}

func (m gobjectClassToGoTypeMetaMap) generateInterfaces(class *Class, g *jen.Group) {
	for _, iface := range class.Implements {
		if !supportedByVersion(iface, m.version) {
			return
		}

		qname := QNameNew(m.ns, iface.Name)
		if qname.namespace.goPackageName != m.ns.goPackageName {
			continue
		}

		g.Lit(qname.name)
	}
}
