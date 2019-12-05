package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/pekim/gobbi/internal/gi"
)

type Record struct {
	Name           string       `xml:"name,attr"`
	Version        string       `xml:"version,attr"`
	CSymbolPrefix  string       `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefix,attr"`
	CType          string       `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	GlibTypeName   string       `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GlibGetType    string       `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	GlibTypeStruct string       `xml:"http://www.gtk.org/introspection/glib/1.0 type-struct,attr"`
	Doc            *Doc         `xml:"doc"`
	ParentName     string       `xml:"parent,attr"`
	Fields         Fields       `xml:"field"`
	Constructors   Constructors `xml:"constructor"`
	Functions      Functions    `xml:"function"`
	Methods        Methods      `xml:"method"`
	//Signals        Signals      `xml:"http://www.gtk.org/introspection/glib/1.0 signal"`

	isClass           bool
	goName            string
	newFromNativeName string
	namespace         *Namespace
	parentNamespace   *Namespace
	parent            *Class

	giInfotype          string
	giInfoGoName        string
	giInfoOnceGoName    string
	giInfoSetFuncGoName string
}

func (r *Record) init(ns *Namespace, isClass bool) {
	r.isClass = isClass
	r.namespace = ns

	r.goName = r.Name
	if r.namespace.Name == "GObject" && r.goName == "SignalQuery" {
		// avoid name clash with a function of the same name
		r.goName = "SignalQuery_"
	}

	r.newFromNativeName = fmt.Sprintf("%sNewFromNative", r.goName)

	r.giInfotype = "Struct"
	if r.isClass {
		r.giInfotype = "Object"
	}
	r.giInfoGoName = fmt.Sprintf("%s%s", makeUnexportedGoName(r.Name), r.giInfotype)
	r.giInfoOnceGoName = fmt.Sprintf("%s_Once", r.giInfoGoName)
	r.giInfoSetFuncGoName = fmt.Sprintf("%s_Set", r.giInfoGoName)

	r.Constructors.init(ns, r)
	r.Functions.init(ns /*r.GoName*/)
	r.Methods.init(ns, r)
	r.Fields.init(ns, r)
	//r.Signals.init(ns, r)
}

func (r *Record) generate(f *file) {
	r.generateGiInfo(f)
	r.generateType(f)
	r.generateNewFromNative(f)
	r.generateAncestorAccessors(f)
	r.Fields.generate(f)
	r.Constructors.generate(f)
	r.Methods.generate(f)

	if len(r.Constructors) == 0 {
		r.generateStructConstructor(f)
		r.generateStructFinalizer(f)
	}

}

func (r *Record) generateGiInfo(f *file) {
	// var colorStruct *gi.Struct
	f.
		Var().
		Id(r.giInfoGoName).
		Op("*").
		Qual(gi.PackageName, r.giInfotype)

	// var colorStructOnce sync.Once
	f.
		Var().
		Id(r.giInfoOnceGoName).
		Qual("sync", "Once")

	// func colorStructSet() error {
	//   ...
	// }
	f.
		Func().
		Id(r.giInfoSetFuncGoName).
		Params().
		List(jen.Id("error")).
		BlockFunc(func(g *jen.Group) {
			// var err error
			g.Var().Id("err").Id("error")

			newFuncName := "StructNew"
			if r.isClass {
				newFuncName = "ObjectNew"
			}

			//   colorStructOnce.Do(func() {
			//     colorStruct, err = gi.StructNew("Gdk", "Color")
			//   })
			g.
				Id(r.giInfoOnceGoName).
				Dot("Do").
				Call(jen.
					Func().
					Params().
					Block(jen.
						Id(r.giInfoGoName).
						Op(",").
						Id("err").
						Op("=").
						Qual(gi.PackageName, newFuncName).
						Call(jen.Lit(r.namespace.Name),
							jen.Lit(r.Name))))

			// return err
			g.Return(jen.Id("err"))
		})
}

func (r *Record) generateType(f *file) {
	f.
		Type().
		Id(r.goName).
		Struct(jen.Id(fieldNameNative).Uintptr())

	f.Line()
}

func (r *Record) generateNewFromNative(f *file) {
	// GEN: func SomeClassNewFromNative(native uintptr) *SomeClass {...}
	f.
		Func().
		Id(r.newFromNativeName).
		Params(jen.Id(fieldNameNative).Uintptr()).
		Params(jen.Op("*").Id(r.goName)).
		BlockFunc(r.generateNewFromNativeBody)
}

func (r *Record) generateNewFromNativeBody(g *jen.Group) {
	g.
		Return(jen.
			Op("&").
			Id(r.goName).
			Values(jen.Dict{
				jen.Id(fieldNameNative): jen.Id(fieldNameNative),
			}),
		)
}

func (r *Record) generateAncestorAccessors(f *file) {
	ancestort := r.parent

	for ancestort != nil {
		accessorName := ancestort.goName
		if ancestort.goName == "Object" && ancestort.namespace.Name != "GObject" {
			// avoid name clash
			accessorName += ancestort.namespace.Name
		}

		parentType := jen.Op("*").Id(ancestort.goName)
		if ancestort.namespace != r.namespace {
			parentType = jen.Op("*").Qual(ancestort.namespace.goFullPackageName, ancestort.goName)
		}

		// GEN: func (recv *SomeClass) AncestorName() *AncestorName {...}
		f.
			Func().
			Params(jen.Id(receiverName).Op("*").Id(r.goName)).
			Id(accessorName).
			Params().
			Params(parentType).
			BlockFunc(func(g *jen.Group) {
				r.generateAncestorAccessorBody(g, ancestort)
			})

		ancestort = ancestort.parent
	}
}

func (r *Record) generateAncestorAccessorBody(g *jen.Group, ancestor *Class) {
	// recv.native
	native := jen.Id(receiverName).Dot(fieldNameNative)

	// SomeClassNewFromNative
	//	OR
	// other.SomeClassNewFromNative
	ancestorNewFromNative := jen.Id(ancestor.newFromNativeName)
	if ancestor.namespace != r.namespace {
		ancestorNewFromNative = jen.Qual(ancestor.namespace.goFullPackageName, ancestor.newFromNativeName)
	}

	// GEN: return SomeClassNewFromNative(native)
	g.Return(ancestorNewFromNative.Call(native))
}

func (r *Record) createFromArgument(g *jen.Group, argName *jen.Statement, argValue *jen.Statement) {
	g.
		Add(argName).
		Op(":=").
		Id(r.newFromNativeName).
		Call(argValue)
}

func (r *Record) generateStructFinalizer(f *file) {
	if r.isClass {
		return
	}

	// func finalizeSomeStruct(obj *SomeStruct) {
	//   ...
	// }
	f.
		Func().
		Id("finalize" + r.goName).
		Params(jen.Id("obj").Op("*").Id(r.goName)).
		BlockFunc(func(g *jen.Group) {
			// SomeStuct.Free(obj.native)
			g.
				Id(r.giInfoGoName).Dot("Free").
				Call(jen.Id("obj").Dot(fieldNameNative))
		})
}

func (r *Record) generateStructConstructor(f *file) {
	if r.isClass {
		return
	}

	f.Line()

	funcName := r.goName + "Struct"

	f.Commentf("%s creates an uninitialised %s.", funcName, r.goName)
	f.
		Func().
		Id(funcName).
		Params().
		Params(jen.Op("*").Id(r.goName)).
		BlockFunc(func(g *jen.Group) {
			// GEN: err := someStruct_Set()
			g.
				Id("err").
				Op(":=").
				Id(r.giInfoSetFuncGoName).
				Call()

			// GEN:
			//	if err != nil {
			//		return nil
			//	}
			g.
				If(jen.Id("err").Op("!=").Nil()).
				Block(jen.Return().Nil())

			g.Line()

			// GEN: structGo := &SomeStruct{native: SomeStruct.Alloc()}
			struct_ := jen.
				Id(r.giInfoGoName).Dot("Alloc").
				Call()
			r.createFromArgument(g, jen.Id("structGo"), struct_)

			// GEN: runtime.SetFinalizer(structGo, finalizeSomeType)
			g.
				Qual("runtime", "SetFinalizer").
				Call(
					jen.Id("structGo"),
					jen.Id("finalize"+r.goName),
				)

			// GEN: return structGo
			g.Return(jen.Id("structGo"))
		})
}
