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
	//f.generateSetter(g)
}

func (f *Field) generateGetter(fi *file) {
	goType, err := f.Type.jenGoType()
	if err != nil {
		fi.Comment(unsupportedf(f.Name, "for field getter : %s", err.Error()))
		fi.Line()
		return
	}

	fi.Commentf("// %s returns the C field '%s'.", f.goName, f.Name)
	fi.
		Func().
		Params(
			jen.
				Id(receiverName).
				Op("*").
				Id(f.record.goName),
		).
		Id(f.goName).
		Params().
		Add(goType).
		BlockFunc(f.generateGetterBody)
}

func (f *Field) generateGetterBody(g *jen.Group) {
	g.
		Id("argValue").
		Op(":=").
		Qual(gi.PackageName, "FieldGet").
		Call(
			jen.Id(f.record.structInfoGoName),
			jen.Id(receiverName).Dot("native"),
			jen.Lit(f.Name),
		)

	typ := f.Type.resolvedType()
	argValue := jen.
		Id("argValue").
		Dot(typ.argumentValueGetFunctionName()).
		CallFunc(func(g *jen.Group) {
			if typ.isString() {
				g.False()
			}
		})

	if f.Type.isAlias() {
		argValue = jen.
			Id(f.Type.Name).
			Parens(argValue)
	}

	createFromArgument := typ.createFromOutArgumentFunction()

	g.
		Id("value").
		Op(":=").
		Do(func(s *jen.Statement) {
			if createFromArgument != nil {
				createFromArgument(s, argValue)
			} else {
				s.Add(argValue)
			}
		})

	g.Return(jen.Id("value"))
}
