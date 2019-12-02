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

	goName           string
	namespace        *Namespace
	newFromCFuncName string
	parentNamespace  *Namespace
	parent           *Class

	structInfoGoName        string
	structInfoOnceGoName    string
	structInfoSetFuncGoName string
}

func (r *Record) init(ns *Namespace) {
	r.namespace = ns
	r.goName = r.Name
	if r.namespace.Name == "GObject" && r.goName == "SignalQuery" {
		// avoid name clash with a function of the same name
		r.goName = "SignalQuery_"
	}
	r.newFromCFuncName = fmt.Sprintf("%sNewFromC", r.Name)

	r.structInfoGoName = fmt.Sprintf("%sStruct", makeUnexportedGoName(r.Name))
	r.structInfoOnceGoName = fmt.Sprintf("%s_Once", r.structInfoGoName)
	r.structInfoSetFuncGoName = fmt.Sprintf("%s_Set", r.structInfoGoName)

	r.Constructors.init(ns, r)
	r.Functions.init(ns /*r.GoName*/)
	r.Methods.init(ns, r)
	r.Fields.init(ns, r)
	//r.Signals.init(ns, r)
}

func (r *Record) generate(f *file) {
	r.generateStructInfo(f)
	r.generateType(f)
	r.Fields.generate(f)
	r.Constructors.generate(f)
	r.Methods.generate(f)

	if len(r.Constructors) == 0 {
		r.generateStructConstructor(f)
		r.generateStructFinalizer(f)
	}

}

func (r *Record) generateStructInfo(f *file) {
	// var colorStruct *gi.Struct
	f.
		Var().
		Id(r.structInfoGoName).
		Op("*").
		Qual(gi.PackageName, "Struct")

	// var colorStructOnce sync.Once
	f.
		Var().
		Id(r.structInfoOnceGoName).
		Qual("sync", "Once")

	// func colorStructSet() error {
	//   ...
	// }
	f.
		Func().
		Id(r.structInfoSetFuncGoName).
		Params().
		List(jen.Id("error")).
		BlockFunc(func(g *jen.Group) {
			// var err error
			g.Var().Id("err").Id("error")

			//   colorStructOnce.Do(func() {
			//     colorStruct, err = gi.StructNew("Gdk", "Color")
			//   })
			g.
				Id(r.structInfoOnceGoName).
				Dot("Do").
				Call(jen.
					Func().
					Params().
					Block(jen.
						Id(r.structInfoGoName).
						Op(",").
						Id("err").
						Op("=").
						Qual(gi.PackageName, "StructNew").
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
		StructFunc(r.generateStructContent)

	f.Line()
}

func (r *Record) generateStructContent(g *jen.Group) {
	if r.parent != nil {
		if r.parentNamespace != nil {
			g.Qual(r.parentNamespace.goFullPackageName, r.parent.Name)
		} else {
			g.Id(r.parent.Name)
		}
	} else {
		g.Id(fieldNameNative).Uintptr()
	}
}

func (r *Record) supportedAsOutParameter() bool {
	return true
}

func (r *Record) createFromArgument(g *jen.Group, argName *jen.Statement, argValue *jen.Statement) {
	g.
		Add(argName).
		Op(":=").
		Op("&").
		Id(r.goName).
		Values()

	g.
		Add(argName).Dot(fieldNameNative).
		Op("=").
		Add(argValue)
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
				Id(r.structInfoGoName).Dot("Free").
				Call(jen.Id("obj").Dot(fieldNameNative))
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
				Id(r.structInfoSetFuncGoName).
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
				Id(r.structInfoGoName).Dot("Alloc").
				Call()
			r.createFromArgument(g, jen.Id("structGo"), struct_)
			//g.
			//	Id("structGo").
			//	Op(":=").
			//	Do(func(s *jen.Statement) {
			//		// SomeStruct.Alloc()
			//		struct_ := jen.
			//			Id(r.structInfoGoName).Dot("Alloc").
			//			Call()
			//
			//		s.Add(r.createFromArgument(struct_))
			//	})

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
