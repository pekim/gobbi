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

var goTypeValueSetterFunctions = map[string]string{
	"gint":     "SetInt",
	"guint":    "SetUint",
	"gint64":   "SetInt64",
	"guint64":  "SetUint64",
	"gdouble":  "SetDouble",
	"utf8":     "SetString",
	"gboolean": "SetBoolean",
}

func (a Argument) generateValueFromObject(s *jen.Statement, objectVarName string) {
	typ := a.Type
	resolvedType := typ.resolvedType()

	if resolvedType.isClass() || resolvedType.isRecord() || resolvedType.isInterface() {
		newFromNative := resolvedType.jenNewFromNative()

		if resolvedType.isClass() || resolvedType.isInterface() {

			// GEN: WidgetNewFromNative(value0.GetObject().Native())
			s.
				Add(newFromNative).
				Call(jen.
					Id(objectVarName).Dot("GetObject").Call().
					Dot("Native").Call())
			return
		} else {
			// GEN: WidgetNewFromNative(value0.GetBoxed().Native())
			s.
				Add(newFromNative).
				Call(jen.
					Id(objectVarName).Dot("GetBoxed").Call())
			return
		}
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

	panic(fmt.Sprintf("unable to unmarshal %s.%s from Object", resolvedType.namespace.Name, resolvedType.Name))
}

func (a Argument) generateObjectFromValue(g *jen.Group, objectVarName string, valueVarName string) {
	typ := a.Type
	resolvedType := typ.resolvedType()

	if funcName, found := goTypeValueSetterFunctions[resolvedType.Name]; found {
		g.Id(objectVarName).Dot(funcName).Call(jen.Id(valueVarName))
		return
	}

	if resolvedType.isClass() {
		g.Id(objectVarName).Dot("SetObject").CallFunc(func(g *jen.Group) {
			if a.Type.Name == "Object" || a.Type.foreignName == "Object" {
				g.Id(valueVarName)
			} else {
				g.Id(valueVarName).Dot("Object").Call()
			}
		})

		return
	}

	if resolvedType.isRecord() {
		g.Id(objectVarName).Dot("SetBoxed").Call(
			jen.Id(valueVarName).Dot("Native").Call())
		return
	}

	if resolvedType.isEnumeration() {
		g.Id(objectVarName).Dot("SetInt").Call(jen.Int32().Params(jen.Id(valueVarName)))
		return
	}

	panic(fmt.Sprintf("unable to unmarshal value to %s.%s Object", resolvedType.namespace.Name, resolvedType.Name))
}
