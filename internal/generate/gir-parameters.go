package generate

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
