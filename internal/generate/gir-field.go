package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/pekim/gobbi/internal/gi"
)

type Field struct {
	Name     string `xml:"name,attr"`
	Writable int    `xml:"writable,attr"`
	Bits     int    `xml:"bits,attr"`
	Private  bool   `xml:"private,attr"`
	Doc      *Doc   `xml:"doc"`
	Type     *Type  `xml:"type"`

	goName    string
	namespace *Namespace
	record    *Record
}

func (f *Field) init(ns *Namespace, record *Record) {
	f.namespace = ns
	f.record = record
	f.goName = makeExportedGoName(f.Name)

	if f.Type != nil {
		f.Type.init(ns)
	}
}

func (f *Field) generate(fi *file) {
	if f.Private {
		return
	}

	f.generateGetter(fi)
	f.generateSetter(fi)
}

func (f *Field) generateGetter(fi *file) {
	goType, err := f.Type.jenGoType()
	if err != nil {
		fi.Comment(unsupportedf(f.Name, "for field getter : %s", err.Error()))
		fi.Line()
		return
	}

	funcName := "Field" + f.goName

	fi.Commentf("// %s returns the C field '%s'.", funcName, f.Name)
	// GEN: func (recv *SomeRecord) Field...() (someType) {...}
	fi.
		Func().
		Add(generateReceiverParams(f.record)).
		Id(funcName).
		Params().
		Add(goType).
		BlockFunc(f.generateGetterBody)
}

func (f *Field) generateGetterBody(g *jen.Group) {
	// GEN: argValue := gi.StructFieldGet(someStruct, recv.native, "field-name")
	g.
		Id("argValue").
		Op(":=").
		Qual(gi.PackageName, f.record.giInfotype+"FieldGet").
		Call(
			jen.Id(f.record.giInfoGoName),
			jen.Id(receiverName).Dot(fieldNameNative),
			jen.Lit(f.Name),
		)

	value := jen.Id("value")
	f.Type.generateOutArgValue(g, value, jen.Id("argValue"), f.transferOwnership())
	g.Return(value)
}

func (f *Field) generateSetter(fi *file) {
	goType, err := f.Type.jenGoType()
	if err != nil {
		fi.Comment(unsupportedf(f.Name, "for field setter : %s", err.Error()))
		fi.Line()
		return
	}

	funcName := "SetField" + f.goName

	fi.Commentf("// %s sets the value of the C field '%s'.", funcName, f.Name)
	// GEN: func (recv *SomeRecord) SetField...(value someType) {...}
	fi.
		Func().
		Add(generateReceiverParams(f.record)).
		Id(funcName).
		Params(jen.Id("value").Add(goType)).
		BlockFunc(f.generateSetterBody)
}

func (f *Field) generateSetterBody(g *jen.Group) {
	jenValue := jen.Id("value")

	if f.Type.isAlias() {
		typ := f.Type.resolvedType()

		jenValue = jen.
			Add(jenGoTypes[typ.Name]).
			Parens(jenValue)
	}

	if f.Type.isEnumeration() {
		jenValue = jen.
			Add(jenGoTypes["int"]).
			Parens(jenValue)
	}

	if f.Type.isRecord() {
		jenValue = jenValue.Dot(fieldNameNative)
	}

	value := f.Type.createFromInArgument(jenValue)

	// var argValue gi.Argument
	g.
		Var().
		Id("argValue").
		Qual(gi.PackageName, "Argument")

	// argValue.Set...(value)
	typ := f.Type.resolvedType()
	g.
		Id("argValue").
		Dot(typ.argumentValueSetFunctionName()).
		Call(value)

	g.
		Qual(gi.PackageName, f.record.giInfotype+"FieldSet").
		Call(
			jen.Id(f.record.giInfoGoName),
			jen.Id(receiverName).Dot(fieldNameNative),
			jen.Lit(f.Name),
			jen.Id("argValue"),
		)
}

func (f *Field) transferOwnership() *bool {
	if !f.Type.isString() {
		return nil
	}

	to := false
	return &to
}
