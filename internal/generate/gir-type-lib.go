package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

func (t *Type) libParamGoType(decrementIndirectionCount bool) *jen.Statement {
	if t.isString() {
		stars := ""
		if t.cType.indirectionCount > 0 {
			stars = strings.Repeat("*", t.cType.indirectionCount-1)
		}
		return jen.Op(stars).String()
	}

	if t.cType.typ == "void" && t.cType.indirectionCount == 1 {
		return jenUnsafePointer()
	}

	// pango specific
	if t.cType.typ == "FILE" && t.cType.indirectionCount == 1 {
		return jenUnsafePointer()
	}

	stars := t.cType.stars
	if decrementIndirectionCount && t.cType.indirectionCount > 0 {
		stars = strings.Repeat("*", t.cType.indirectionCount-1)
	}

	if t.isBitfield() {
		if t.isQualifiedName() {
			bitfield, _ := t.foreignNamespace.Bitfields.byName(t.foreignName)
			return jen.Op(stars).Qual(t.foreignNamespace.goFullPackageName, bitfield.goTypeName)
		}

		bitfield, _ := t.namespace.Bitfields.byName(t.Name)
		return jen.Op(stars).Id(bitfield.goTypeName)
	}

	if t.isEnumeration() {
		if t.isQualifiedName() {
			enum, _ := t.foreignNamespace.Enumerations.byName(t.foreignName)
			return jen.Op(stars).Qual(t.foreignNamespace.goFullPackageName, enum.goTypeName)
		}

		enum, _ := t.namespace.Enumerations.byName(t.Name)
		return jen.Op(stars).Id(enum.goTypeName)
	}

	if simpleGoType, ok := simpleSysParamGoTypes[t.cType.typ]; ok {
		return jen.Op(stars).Add(simpleGoType)
	}
	if simpleGoType, ok := simpleSysParamGoTypes[t.Name]; ok {
		return jen.Op(stars).Add(simpleGoType)
	}

	if t.isAlias() {
		if t.isQualifiedName() {
			alias, _ := t.foreignNamespace.Aliases.byName(t.foreignName)
			return jen.Op(stars).Add(alias.Type.sysParamGoType(false))
		}

		alias, _ := t.namespace.Aliases.byName(t.Name)
		return jen.Op(stars).Add(alias.Type.sysParamGoType(false))
	}

	if t.isRecord() && t.cType.indirectionCount == 0 {
		if t.isQualifiedName() {
			return jen.Qual(t.foreignNamespace.goFullPackageName, t.foreignName)
		}
		return jen.Op(stars).Id(t.Name)
	}

	if t.isStruct() {
		if t.cType.indirectionCount == 0 {
			return jenUnsafePointer()
		}

		if t.isQualifiedName() {
			return jen.Op(t.cType.stars).Qual(t.foreignNamespace.goFullPackageName, t.foreignName)
		}
		return jen.Op(t.cType.stars).Id(t.Name)
	}

	panic(fmt.Sprintf("Unsupported type : %s %s (%s)", t.namespace.Name, t.Name, t.CType))
}
