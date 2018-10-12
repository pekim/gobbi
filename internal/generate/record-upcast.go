package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

func (r *Record) generateUpcasts(g *jen.Group) {
	if r.ParentName == "" {
		return
	}

	methodNames := make(map[string]bool)

	qname := QNameNew(r.Namespace, r.ParentName)
	methodNames[qname.name] = true

	parent, found := qname.ns.recordOrClassRecordForName(qname.name)
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
		qname = QNameNew(qname.ns, ancestorName)

		if _, nameUsedPreviously := methodNames[qname.name]; nameUsedPreviously {
			// A method for this ancestor would have the same name as an earlier one.
			// So skip it, and all further ancestors.
			// This will affect relatively few methods, notably those around Atk.Object.
			return
		}
		methodNames[qname.name] = true

		ancestor, found := qname.ns.recordOrClassRecordForName(qname.name)
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
			if qname.sameNamespace {
				g.Op("*").Id(qname.name)
			} else {
				g.Op("*").Qual(qname.ns.fullGoPackageName, qname.name)
			}
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
					Qual(qname.ns.fullGoPackageName, parent.newFromCFuncName).
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
			if qname.ns.Name == r.Namespace.Name {
				g.Op("*").Id(qname.name)
			} else {
				g.Op("*").Qual(qname.ns.fullGoPackageName, qname.name)
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