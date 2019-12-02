package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/pekim/gobbi/internal/gi"
)

type Parameters []*Parameter

func (pp Parameters) init(ns *Namespace) {
	//pp.fixupArgcArgv()
	//pp.fixupStringLengthParams()
	//pp.fixupFormatArgs()

	for _, param := range pp {
		param.init(ns)

		//if param.Array != nil {
		//	if param.Array.Type != nil {
		//		param.Array.Type.init(ns)
		//		if param.Array.Type.generator == nil && param.Name == "files" {
		//			fmt.Println(param.Name, param.Array.Type.CType)
		//			panic("xxx")
		//		}
		//	}
		//
		//	if param.Array.Length != nil {
		//		// Provide an array length param with a reference to its array param.
		//		paramIndex := *param.Array.Length
		//		pp[paramIndex].arrayLengthFor = param
		//
		//		// And the reverse.
		//		param.Array.lengthParam = pp[paramIndex]
		//	}
		//}
	}
}

func (pp Parameters) supported() (bool, string) {
	for _, param := range pp {
		if supported, reason := param.supported(); !supported {
			return supported, reason
		}
	}

	return true, ""
}

func (pp Parameters) generateInDeclarations(g *jen.Group) {
	for _, param := range pp {
		if param.isIn() {
			param.generateInDeclaration(g)
		}
	}
}

func (pp Parameters) generateReturnDeclarations(g *jen.Group) {
	for _, param := range pp {
		if param.isOut() {
			param.generateReturnDeclaration(g)
		}
	}
}

func (pp Parameters) inCount(receiver bool) int {
	count := 0
	if receiver {
		count++
	}
	for _, param := range pp {
		if param.isIn() {
			count++
		}
	}

	return count
}

func (pp Parameters) outCount() int {
	count := 0
	for _, param := range pp {
		if param.isOut() {
			count++
		}
	}

	return count
}

func (pp Parameters) generateInArgs(g *jen.Group, receiver bool) {
	count := pp.inCount(receiver)
	if count == 0 {
		return
	}

	// "var inArgs [n]gi.Arguments"
	g.
		Var().
		Id("inArgs").
		Index(jen.Lit(count)).
		Qual(gi.PackageName, "Argument")

	// set values in inArgs
	n := 0
	if receiver {
		g.
			Id("inArgs").
			Index(jen.Lit(n)).
			Dot("SetPointer").
			Call(jen.Id(receiverName).Dot(fieldNameNative))

		n++
	}
	for _, param := range pp {
		if param.isIn() {
			param.generateInArg(g, n)
			n++
		}
	}
}

func (pp Parameters) generateOutValues(g *jen.Group, varNamePrefix string) {
	n := 0
	for _, param := range pp {
		if param.isOut() {
			param.generateOutValue(g, varNamePrefix, n)
			n++
		}
	}
}

func (pp Parameters) generateOutArgs(g *jen.Group) {
	count := pp.outCount()
	if count == 0 {
		return
	}

	// "var outArgs [n]gi.Arguments"
	g.
		Var().
		Id("outArgs").
		Index(jen.Lit(count)).
		Qual(gi.PackageName, "Argument")
}

func (pp Parameters) generateCallParams(g *jen.Group, receiver bool) {
	// in args
	if pp.inCount(receiver) > 0 {
		g.
			Id("inArgs").
			Index(jen.Op(":"))
	} else {
		g.Nil()
	}

	// out args
	if pp.outCount() > 0 {
		g.
			Id("outArgs").
			Index(jen.Op(":"))
	} else {
		g.Nil()
	}
}
