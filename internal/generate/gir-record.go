package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/pekim/gobbi/internal/cgo/callback"
	"github.com/pekim/gobbi/internal/cgo/gi"
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
	Signals        Signals      `xml:"http://www.gtk.org/introspection/glib/1.0 signal"`

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

func (r *Record) init(ns *Namespace, giInfoType string) {
	r.namespace = ns

	r.goName = r.Name
	if r.namespace.Name == "GObject" && r.goName == "SignalQuery" {
		// avoid name clash with a function of the same name
		r.goName = "SignalQuery_"
	}

	r.newFromNativeName = fmt.Sprintf("%sNewFromNative", r.goName)

	r.giInfotype = giInfoType
	r.giInfoGoName = fmt.Sprintf("%s%s", makeUnexportedGoName(r.Name), r.giInfotype)
	r.giInfoOnceGoName = fmt.Sprintf("%s_Once", r.giInfoGoName)
	r.giInfoSetFuncGoName = fmt.Sprintf("%s_Set", r.giInfoGoName)

	r.Constructors.init(ns, r)
	r.Functions.init(ns /*r.GoName*/)
	r.Methods.init(ns, r)
	r.Fields.init(ns, r)
	r.Signals.init(ns, r)
}

func (r *Record) generate(f *file) {
	r.generateGiInfo(f)
	r.generateType(f)
	r.generateNewFromNative(f)
	r.generateAncestorAccessors(f)
	r.generateDownCast(f)
	r.generateEquals(f)
	r.generateNativeAccessor(f)
	r.Fields.generate(f)
	r.Constructors.generate(f)
	r.Methods.generate(f)
	r.Signals.generate(f)
	r.generateSignalDisconnect(f)

	if r.giInfotype == "Struct" && len(r.Constructors) == 0 {
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
						Qual(gi.PackageName, r.giInfotype+"New").
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
		Struct(jen.Id(fieldNameNative).Qual("unsafe", "Pointer"))

	f.Line()
}

func (r *Record) generateNewFromNative(f *file) {
	// GEN: func SomeClassNewFromNative(native unsafe.Pointer) *SomeClass {...}
	f.
		Func().
		Id(r.newFromNativeName).
		Params(jen.Id(fieldNameNative).Qual("unsafe", "Pointer")).
		Params(jen.Op("*").Id(r.goName)).
		BlockFunc(r.generateNewFromNativeBody)

	f.Line()
}

func (r *Record) generateNewFromNativeBody(g *jen.Group) {
	instance := jen.Id("instance")
	object := jen.Id("object")

	// GEN: instance := &SomeRecord{native: native}
	g.
		Add(instance).
		Op(":=").
		Op("&").
		Id(r.goName).
		Values(jen.Dict{
			jen.Id(fieldNameNative): jen.Id(fieldNameNative),
		})
	g.Line()

	// Increment ref count as appropriate.
	if r.derivedFromObject() {
		g.
			Add(object).
			Op(":=").
			Add(instance).Dot("Object").
			Call()

		g.
			If(jen.Add(object).Dot("IsFloating").Call()).
			Block(jen.Add(object).Dot("RefSink").Call()).
			Else().
			Block(jen.Add(object).Dot("Ref").Call())

		g.Line()
	}

	r.setFinalizer(g, object)

	// GEN: return instance
	g.Return(instance)
}

func (r *Record) setFinalizer(g *jen.Group, instance *jen.Statement) {
	if !r.derivedFromObject() {
		return
	}

	object := jen.Id("o")
	gobjectNs, _ := r.namespace.namespaces.byName("GObject")

	g.
		// GEN: runtime.SetFinalizer(...)
		Qual("runtime", "SetFinalizer").
		Call(
			// GEN: instance
			jen.Add(instance),

			// GEN: func (instance *SomeRecord) { o.Object().Unref() }
			jen.
				Func().
				Params(jen.Add(object).Op("*").Qual(gobjectNs.goFullPackageName, "Object")).
				Block(jen.Add(object).
					Dot("Unref").Call(),
				),
		)

	g.Line()
}

func (r *Record) derivedFromInitiallyUnowned() bool {
	return r.derivedFrom("InitiallyUnowned")
}

func (r *Record) derivedFromObject() bool {
	return r.derivedFrom("Object")
}

func (r *Record) derivedFrom(name string) bool {
	if r.namespace.Name == "GObject" {
		return false
	}

	ancestor := r.parent
	for ancestor != nil {
		if ancestor.namespace.Name == "GObject" && ancestor.Name == name {
			return true
		}

		ancestor = ancestor.parent
	}

	return false
}

func (r *Record) generateAncestorAccessors(f *file) {
	ancestor := r.parent

	for ancestor != nil {
		accessorName := ancestor.goName
		if ancestor.goName == "Object" && ancestor.namespace.Name != "GObject" {
			// avoid name clash
			accessorName += ancestor.namespace.Name
		}

		parentType := jen.Op("*").Id(ancestor.goName)
		if ancestor.namespace != r.namespace {
			parentType = jen.Op("*").Qual(ancestor.namespace.goFullPackageName, ancestor.goName)
		}

		f.Commentf("%s upcasts to *%s", accessorName, accessorName)

		// GEN: func (recv *SomeClass) AncestorName() *AncestorName {...}
		r.methodPrelude(f, accessorName).
			Params().
			Params(parentType).
			BlockFunc(func(g *jen.Group) {
				r.generateAncestorAccessorBody(g, ancestor)
			})

		ancestor = ancestor.parent
	}
}

func (r *Record) generateAncestorAccessorBody(g *jen.Group, ancestor *Class) {
	// recv.native
	native := jen.Id(receiverName).Dot(nativeAccessorName).Call()

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

func (r *Record) generateDownCast(f *file) {
	if r.namespace.Name == "GObject" || r.namespace.Name == "GLib" {
		return
	}

	funcName := "CastTo" + r.goName
	gobjectNs, _ := r.namespace.namespaces.byName("GObject")

	f.Commentf(`%s down casts any arbitrary Object to %s.
Exercise care, as this is a potentially dangerous function
if the Object is not a %s.`,
		funcName, r.goName, r.goName)

	f.
		Func().
		Id(funcName).
		Params(jen.
			Id("object").
			Op("*").Qual(gobjectNs.goFullPackageName, "Object")).
		Params(jen.Op("*").Id(r.goName)).
		BlockFunc(func(g *jen.Group) {
			native := jen.Id("object").Dot(nativeAccessorName).Call()
			g.Return(jen.Id(r.newFromNativeName).Call(native))
		})

	f.Line()
}

func (r *Record) generateNativeAccessor(f *file) {
	r.methodPrelude(f, nativeAccessorName).
		Params().
		Params(jen.Qual("unsafe", "Pointer")).
		Block(jen.
			Return().
			Id(receiverName).Dot(fieldNameNative))

	f.Line()
}

func (r *Record) createFromArgument(g *jen.Group, foreignNamespace *Namespace, argName *jen.Statement, argValue *jen.Statement) {
	newFromNativeFunc := jen.Id(r.newFromNativeName)
	if foreignNamespace != nil {
		newFromNativeFunc = jen.Qual(foreignNamespace.goFullPackageName, r.newFromNativeName)
	}

	g.
		Add(argName).
		Op(":=").
		Add(newFromNativeFunc).
		Call(argValue)
}

func (r *Record) generateStructFinalizer(f *file) {
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
				Call(jen.Id("obj").Dot(nativeAccessorName).Call())
		})
}

func (r *Record) generateStructConstructor(f *file) {
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
			r.createFromArgument(g, nil, jen.Id("structGo"), struct_)

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

func (r *Record) methodPrelude(f *file, methodName string) *jen.Statement {
	return f.
		Func().
		Params(jen.Id(receiverName).Op("*").Id(r.goName)).
		Id(methodName)
}

func (r *Record) generateEquals(f *file) {
	f.Commentf("Equals compares this %s with another %s, and returns true if they represent the same Object.",
		r.goName, r.goName)

	// GEN: func (recv *Cursor) Equals(other *Cursor) bool {...}
	r.methodPrelude(f, "Equals").
		Params(jen.Id("other").Op("*").Id(r.goName)).
		Params(jen.Bool()).
		Block(
			// GEN: return other.Native() == recv.Native()
			jen.Return(jen.
				Id("other").Dot(nativeAccessorName).Call().
				Op("==").
				Id(receiverName).Dot(nativeAccessorName).Call()))

	f.Line()
}

func (r *Record) generateSignalDisconnect(f *file) {
	if len(r.Signals) == 0 {
		// If there are no signals there is no point in generating
		// a Disconnect method.
		return
	}

	id := "connectionID"

	f.Commentf(`Disconnect disconnects a callback previously registered with a Connect...() method.

The %s should be a value returned from a call to a Connect...() method.`,
		id)

	// GEN: func (recv *Device) DisconnectSignal(connectionId int) {...}
	r.methodPrelude(f, "DisconnectSignal").
		Params(jen.Id(id).Int()).
		Block(jen.
			// GEN: callback.DisconnectSignal(connectionId)
			Qual(callback.PackageName, "DisconnectSignal").Call(jen.Id(id)))

}
