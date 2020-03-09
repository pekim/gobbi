package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

func (p *Parameter) sysParamGoType() *jen.Statement {
	if p.isType() {
		// Atoms are really pointers underneath.
		if p.Type.CType == "GdkAtom" {
			return jenUnsafePointer()
		}

		star := ""
		if p.Nullable && !p.isOut() && !p.Type.isStruct() && !p.Type.isPointer() {
			// nullable simple type, so make it a pointer
			star = "*"
		}

		return jen.
			Op(star).
			Add(p.Type.sysParamGoType())
	}

	if p.isArray() {
		star := ""
		if p.isOut() {
			star = "*"
		}

		return jen.
			Op(star).
			Add(p.Array.sysParamGoType())
	}

	panic(fmt.Sprintf("Parameter is not a type or an array: %s", p.context))
}
