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
	if a.Type.isClass() || a.Type.isRecord() || a.Type.isInterface() {
		newFromNative := a.Type.jenNewFromNative()

		if a.Type.isClass() || a.Type.isInterface() {

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

	if a.Type.isBitfield() {
		bitfield, _ := a.Type.namespace.Bitfields.byName(a.Type.Name)
		s.Params(jen.Id(bitfield.goTypeName)).
			Params(jen.Id(objectVarName).Dot("GetInt").Call())

		return
	}

	if a.Type.isEnumeration() {
		enum, _ := a.Type.namespace.Enumerations.byName(a.Type.Name)
		s.Params(jen.Id(enum.goTypeName)).
			Params(jen.Id(objectVarName).Dot("GetInt").Call())

		return
	}

	if funcName, found := goTypeValueGetterFunctions[a.Type.Name]; found {
		s.Id(objectVarName).Dot(funcName).Call()
		return
	}

	if a.Type.Name == "gpointer" {
		s.Id(objectVarName).Dot("GetPointer()")
		return
	}

	panic(fmt.Sprintf("unable to unmarshal %s.%s from Object", a.Type.namespace.Name, a.Type.Name))
}

func (a Argument) generateObjectFromValue(g *jen.Group, objectVarName string, valueVarName string) {
	if funcName, found := goTypeValueSetterFunctions[a.Type.Name]; found {
		g.Id(objectVarName).Dot(funcName).Call(jen.Id(valueVarName))
		return
	}

	if a.Type.isClass() {
		g.Id(objectVarName).Dot("SetObject").CallFunc(func(g *jen.Group) {
			if a.Type.Name == "Object" || a.Type.foreignName == "Object" {
				g.Id(valueVarName)
			} else {
				g.Id(valueVarName).Dot("Object").Call()
			}
		})

		return
	}

	if a.Type.isRecord() {
		g.Id(objectVarName).Dot("SetBoxed").Call(
			jen.Id(valueVarName).Dot("Native").Call())
		return
	}

	if a.Type.isEnumeration() {
		g.Id(objectVarName).Dot("SetInt").Call(jen.Int32().Params(jen.Id(valueVarName)))
		return
	}

	panic(fmt.Sprintf("unable to unmarshal value to %s.%s Object", a.Type.namespace.Name, a.Type.Name))
}
