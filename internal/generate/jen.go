package generate

import "github.com/dave/jennifer/jen"

// jenUnsafePointer returns a statement representing "unsafe.Pointer".
//
// Reduces the risk of typos.
func jenUnsafePointer() *jen.Statement {
	return jen.Qual("unsafe", "Pointer")
}
