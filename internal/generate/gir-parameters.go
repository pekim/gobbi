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

func (pp Parameters) generateInDeclarations(g *group) {
	for _, param := range pp {
		if param.isIn() {
			param.generateInDeclaration(g)
		}
	}
}

func (pp Parameters) generateReturnDeclarations(g *group) {
	for _, param := range pp {
		if param.isOut() {
			param.generateReturnDeclaration(g)
		}
	}
}

func (pp Parameters) inCount() int {
	count := 0
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

func (pp Parameters) generateInArgs(g *group) {
	count := pp.inCount()
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
	for _, param := range pp {
		if param.isIn() {
			param.generateInArg(g, n)
			n++
		}
	}

	g.Line()
}

func (pp Parameters) generateOutValues(g *jen.Group) {
	n := 0
	for _, param := range pp {
		if param.isOut() {
			param.generateOutValue(g, n)
			n++
		}
	}
}

func (pp Parameters) generateOutArgs(g *group) {
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
	g.Line()
}

func (pp Parameters) generateCallParams(g *jen.Group) {
	// in args
	if pp.inCount() > 0 {
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
