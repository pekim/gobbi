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

func (pp Parameters) generateDeclarations(g *group) {
	for _, param := range pp {
		param.generateDeclaration(g)
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

func (pp Parameters) generateInArgs(g *group) {
	count := pp.inCount()
	if count == 0 {
		return
	}

	g.
		Id("inArgs").
		Op(":=").
		Make(
			jen.Index().Qual(gi.PackageName, "Argument"),
			jen.Lit(count),
		)

	for n, param := range pp {
		if param.isIn() {
			param.generateInArg(g, n)
		}
	}

	g.Line()
}

func (pp Parameters) generateCallParams(g *jen.Group) {
	if pp.inCount() > 0 {
		g.Id("inArgs")
	} else {
		g.Nil()
	}
}
