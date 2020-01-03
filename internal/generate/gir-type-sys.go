package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

func (t *Type) sysParamGoPlainType() *jen.Statement {
	if t.isString() {
		return jen.String()
	}

	//if t.cType.typ == "void" && t.cType.indirectionCount == 1 {
	//	return jenUnsafePointer()
	//}
	if t.cType.typ == "void" {
		return nil
	}

	// pango specific
	if t.cType.typ == "FILE" && t.cType.indirectionCount == 1 {
		return jenUnsafePointer()
	}

	if simpleGoType, ok := simpleSysParamGoTypes[t.cType.typ]; ok {
		return jen.Add(simpleGoType)
	}

	if t.isAlias() {
		if t.isQualifiedName() {
			alias, _ := t.foreignNamespace.Aliases.byName(t.foreignName)
			return jen.Add(alias.Type.sysParamGoType(false))
		}

		alias, _ := t.namespace.Aliases.byName(t.Name)
		return jen.Add(alias.Type.sysParamGoType(false))
	}

	if t.isBitfield() || t.isEnumeration() {
		return jen.Int()
	}

	if t.isRecord() && t.cType.indirectionCount == 0 {
		if t.isQualifiedName() {
			return jen.Qual(t.foreignNamespace.goFullSysPackageName, t.foreignName)
		}
		return jen.Id(t.Name)
	}

	if t.isClass() ||
		t.isRecord() ||
		t.isInterface() ||
		t.isUnion() {

		return jenUnsafePointer()
	}

	panic(fmt.Sprintf("Unsupported type : %s %s (%s)", t.namespace.Name, t.Name, t.CType))
}

func (t *Type) sysParamGoType(decrementIndirectionCount bool) *jen.Statement {
	if t.CType == "GdkAtom" {
		return jenUnsafePointer()
	}

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

	if t.isBitfield() || t.isEnumeration() {
		return jen.Op(stars).Int()
	}

	if t.isRecord() && t.cType.indirectionCount == 0 {
		if t.isQualifiedName() {
			return jen.Qual(t.foreignNamespace.goFullSysPackageName, t.foreignName)
		}
		return jen.Op(stars).Id(t.Name)
	}

	if t.isClass() ||
		t.isRecord() ||
		t.isInterface() ||
		t.isUnion() {

		if t.cType.indirectionCount == 0 {
			return jenUnsafePointer()
		}

		stars := strings.Repeat("*", t.cType.indirectionCount-1)
		return jen.Op(stars).Add(jenUnsafePointer())
	}

	panic(fmt.Sprintf("Unsupported type : %s %s (%s)", t.namespace.Name, t.Name, t.CType))
}
