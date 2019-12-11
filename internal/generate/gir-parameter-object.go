package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

var goTypeValueGetterFunctions = map[string]string{
	"gint":     "GetInt",
	"guint":    "GetUint",
	"gint64":   "GetInt64",
	"guint64":  "GetUint64",
	"gdouble":  "GetDouble",
	"utf8":     "GetString",
	"gboolean": "GetBoolean",
}

func (p Parameter) generateValueFromObject(s *jen.Statement, objectVarName string) {
	typ := p.Type
	resolvedType := typ.resolvedType()

	if resolvedType.isClass() {
		var newFromNative *jen.Statement
		if resolvedType.isQualifiedName() {
			class, _ := resolvedType.foreignNamespace.Classes.byName(resolvedType.foreignName)
			newFromNative = jen.Qual(resolvedType.foreignNamespace.goFullPackageName, class.newFromNativeName)
		} else {
			class, _ := resolvedType.namespace.Classes.byName(resolvedType.Name)
			newFromNative = jen.Id(class.newFromNativeName)
		}

		// GEN: WidgetNewFromNative(value0.GetObject().Native())
		s.
			Add(newFromNative).
			Call(jen.
				Id(objectVarName).Dot("GetObject").Call().
				Dot("Native").Call())
		return
	}

	if resolvedType.isRecord() {
		var newFromNative *jen.Statement
		if resolvedType.isQualifiedName() {
			record, _ := resolvedType.foreignNamespace.Records.byName(resolvedType.foreignName)
			newFromNative = jen.Qual(resolvedType.foreignNamespace.goFullPackageName, record.newFromNativeName)
		} else {
			record, _ := resolvedType.namespace.Records.byName(resolvedType.Name)
			newFromNative = jen.Id(record.newFromNativeName)
		}

		// GEN: WidgetNewFromNative(value0.GetObject().Native())
		s.
			Add(newFromNative).
			Call(jen.
				Id(objectVarName).Dot("GetObject").Call().
				Dot("Native").Call())
		return
	}

	if resolvedType.isBitfield() {
		bitfield, _ := resolvedType.namespace.Bitfields.byName(resolvedType.Name)
		s.Params(jen.Id(bitfield.goTypeName)).
			Params(jen.Id(objectVarName).Dot("GetInt").Call())

		return
	}

	if resolvedType.isEnumeration() {
		enum, _ := resolvedType.namespace.Enumerations.byName(resolvedType.Name)
		s.Params(jen.Id(enum.goTypeName)).
			Params(jen.Id(objectVarName).Dot("GetInt").Call())

		return
	}

	if funcName, found := goTypeValueGetterFunctions[resolvedType.Name]; found {
		s.Id(objectVarName).Dot(funcName).Call()
		return
	}

	if resolvedType.Name == "gpointer" {
		s.Id(objectVarName).Dot("GetPointer()")
		return
	}

	fmt.Println(resolvedType.Name)
	s.Lit(42)
}
