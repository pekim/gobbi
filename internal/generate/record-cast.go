package generate

import (
	"fmt"
	"github.com/pekim/jennifer/jen"
)

func (r *Record) generateUpcasts(g *jen.Group) {
	if r.ParentName == "" {
		return
	}

	methodNames := make(map[string]bool)

	qname := QNameNew(r.Namespace, r.ParentName)
	methodNames[qname.name] = true

	parent, found := qname.namespace.recordOrClassRecordForName(qname.name)
	if !found {
		panic(fmt.Sprintf("Failed to find parent %s for %s", r.ParentName, r.Name))
	}

	if parent.Blacklist {
		return
	}

	r.generateParentUpcast(g, parent, qname)

	parentQname := qname
	previousAncestor := parent
	ancestorName := previousAncestor.ParentName

	for ancestorName != "" {
		qname = QNameNew(qname.namespace, ancestorName)

		if _, nameUsedPreviously := methodNames[qname.name]; nameUsedPreviously {
			// A method for this ancestor would have the same name as an earlier one.
			// So skip it, and all further classAncestors.
			// This will affect relatively few methods, notably those around Atk.Object.
			return
		}
		methodNames[qname.name] = true

		ancestor, found := qname.namespace.recordOrClassRecordForName(qname.name)
		if !found {
			panic(fmt.Sprintf("Failed to find parent %s for %s", ancestorName, previousAncestor.Name))
		}

		if ancestor.Blacklist {
			return
		}

		r.generateAncestorUpcast(g, parentQname, qname)

		previousAncestor := ancestor
		ancestorName = previousAncestor.ParentName
	}
}

func (r *Record) generateParentUpcast(g *jen.Group, parent *Record, qname *QName) {
	g.Commentf("%s upcasts to *%s", qname.name, qname.name)

	g.
		Func().
		Parens(jen.Id("recv").Op("*").Id(r.GoName)). //  receiver declaration
		Id(qname.name).                              // func name
		Params().                                    // params
		ParamsFunc(func(g *jen.Group) {              // return value
			g.Do(func(s *jen.Statement) {
				s.Op("*")
				qname.generate(s)
			})
		}).
		BlockFunc(func(g *jen.Group) { // body
			nativeAsUnsafePointer := jen.
				Qual("unsafe", "Pointer").
				Call(jen.Id("recv").Op(".").Id("native"))

			if qname.sameNamespace {
				g.
					Return().
					Id(parent.newFromCFuncName).
					Call(nativeAsUnsafePointer)
			} else {
				g.
					Return().
					Qual(qname.namespace.fullGoPackageName, parent.newFromCFuncName).
					Call(nativeAsUnsafePointer)
			}
		}).
		Line()
}

func (r *Record) generateAncestorUpcast(g *jen.Group, parentQName *QName, qname *QName) {
	g.Commentf("%s upcasts to *%s", qname.name, qname.name)

	g.
		Func().
		Parens(jen.Id("recv").Op("*").Id(r.GoName)). //  receiver declaration
		Id(qname.name).                              // func name
		Params().                                    // params
		ParamsFunc(func(g *jen.Group) {              // return value
			if qname.namespace.Name == r.Namespace.Name {
				g.Op("*").Id(qname.name)
			} else {
				g.Op("*").Qual(qname.namespace.fullGoPackageName, qname.name)
			}
		}).
		BlockFunc(func(g *jen.Group) { // body
			g.
				Return().
				Id("recv").
				Op(".").
				Id(parentQName.name).
				Call().
				Op(".").
				Id(qname.name).
				Call()
		}).
		Line()
}

func (r *Record) generateDowncast(g *jen.Group) {
	if r.root().Name != "Object" {
		return
	}

	g.Commentf("CastToWidget down casts any arbitrary Object to %s.", r.Name)
	g.Commentf("Exercise care, as this is a potentially dangerous function if the Object is not a %s.", r.Name)

	g.
		Func().
		Id("CastTo" + r.Name).
		ParamsFunc(func(g *jen.Group) {
			if r.Namespace.Name == "GObject" {
				g.Id("object").Op("*").Id("Object")
			} else {
				g.Id("object").Op("*").Qual("github.com/pekim/gobbi/lib/gobject", "Object")
			}
		}).
		Params(jen.Op("*").Id(r.Name)).
		Block(jen.
			Return().
			Id(r.Name + "NewFromC").
			Call(jen.
				Id("object").
				Op(".").
				Id("ToC").
				Call()))

	g.Line()
}
