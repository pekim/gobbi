package generate

import "C"
import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
)

func (f *Function) generateLib(fi *jen.File, version semver.Version) {
	if supported, reason := f.isSupported(); !supported {
		fi.Commentf("UNSUPPORTED : %s : %s", f.CIdentifier, reason)
		fi.Line()
		return
	}

	// TODO
	if f.Throws {
		fi.Commentf("UNSUPPORTED : %s : throws", f.CIdentifier)
		fi.Line()
		return
	}

	// TODO
	for _, p := range f.Parameters {
		if p.Array != nil {
			fi.Commentf("UNSUPPORTED : %s : has array param, %s", f.CIdentifier, p.Name)
			fi.Line()
			return
		}

		//if p.isOut() {
		//	fi.Commentf("UNSUPPORTED : %s : has [in]out param, %s", f.CIdentifier, p.Name)
		//	fi.Line()
		//	return
		//}
	}
	if f.ReturnValue.Array != nil {
		fi.Commentf("UNSUPPORTED : %s : has array return", f.CIdentifier)
		fi.Line()
		return
	}

	if f.MovedTo != "" || f.ShadowedBy != "" {
		return
	}

	if f.version.GT(version) {
		return
	}

	fi.Commentf("%s wraps the C function %s.", f.goName, f.CIdentifier)
	docVersion(fi, f.Version)

	// func Fn_some_function(...) [return type] {...}
	fi.
		Func().
		Do(f.generateLibReceiverDeclaration).
		Id(f.goName).
		ParamsFunc(f.generateLibParamsDeclaration).
		ParamsFunc(f.generateLibReturnTypesDeclaration).
		BlockFunc(f.generateLibBody)

	fi.Line()
}

func (f *Function) generateLibReceiverDeclaration(s *jen.Statement) {
	if f.InstanceParameter != nil {
		s.Id("recv").Add(jen.Id(f.InstanceParameter.Type.Name))

	}
}

func (f *Function) generateLibParamsDeclaration(g *jen.Group) {
	for _, param := range f.Parameters {
		if param.isVarargsOrValist() {
			continue
		}

		if !param.isIn() {
			continue
		}

		goType := param.libParamGoType()

		g.Id(param.goVarName).Add(goType)
	}

	//if f.Throws {
	//	g.Id("error").Add(jenUnsafePointer())
	//}
}

func (f *Function) generateLibReturnTypesDeclaration(g *jen.Group) {
	if !f.ReturnValue.isVoid() {
		// return value
		typ := f.ReturnValue.Type
		if typ.isStruct() {
			// TODO below belongs in Type.sysParamGoType ???
			g.Op("*").Add(typ.idOrQual())
		} else {
			g.Add(f.ReturnValue.sysParamGoType())
		}
	}

	// out params
	for _, param := range f.Parameters {
		if !param.isOut() {
			continue
		}

		if param.Type.isStruct() {
			// TODO below belongs in Type.sysParamGoType ???
			g.Op("*").Add(param.Type.idOrQual())
		} else {
			g.Add(param.sysReturnGoType())
		}
	}
}

func (f *Function) generateLibBody(g *jen.Group) {
	// create args for sys function
	for _, p := range f.Parameters {
		if p.isVarargsOrValist() {
			continue
		}

		argVarName := "sys_" + p.goVarName
		p.generateLibArg(g, argVarName)
	}

	// call sys function
	g.
		Do(func(s *jen.Statement) {
			if !f.ReturnValue.isVoid() {
				s.Id("retSys").Op(":=")
			}
		}).
		Qual(f.namespace.goFullSysPackageName, f.sysName).CallFunc(f.generateLibCallParams)

	f.generateReturn(g)
}

func (f *Function) generateLibCallParams(g *jen.Group) {
	//if f.InstanceParameter != nil {
	//	g.Id("cValueInstance")
	//}

	for _, param := range f.Parameters {
		if param.isVarargsOrValist() {
			continue
		}

		if param.isOut() {
			g.Op("&").Id("sys_" + param.goVarName)
		} else {
			g.Id("sys_" + param.goVarName)
		}
	}

	//if f.Throws {
	//	g.Id("cError")
	//}
}

func (f *Function) generateReturn(g *jen.Group) {
	if f.ReturnValue.isVoid() && f.Parameters.outCount() == 0 {
		return
	}

	g.Line()

	g.ReturnFunc(func(g *jen.Group) {
		if !f.ReturnValue.isVoid() {
			f.generateMarshalReturnValue(g, "retSys", f.ReturnValue.Type)
		}

		for _, param := range f.Parameters {
			if !param.isOut() {
				continue
			}

			argVarName := "sys_" + param.goVarName
			f.generateMarshalReturnValue(g, argVarName, param.Type)
		}
	})
}

func (f *Function) generateMarshalReturnValue(g *jen.Group, varName string, typ *Type) {
	if typ.isStruct() {
		if typ.isQualifiedName() {
			ctorName := typ.foreignName + "NewFromC"
			g.Qual(typ.foreignNamespace.goFullPackageName, ctorName).Call(jen.Id(varName))
			return
		}

		ctorName := typ.Name + "NewFromC"
		g.Id(ctorName).Call(jen.Id(varName))
		return
	}

	g.Id(varName)
}

//func (f *Function) generateLibVarArgsCFunction(fi *jen.File) {
//	params := []string{}
//	args := []string{}
//
//	if f.InstanceParameter != nil {
//		if f.InstanceParameter.Type != nil {
//			params = append(params, f.InstanceParameter.Type.CType+" "+f.InstanceParameter.Name)
//			args = append(args, f.InstanceParameter.Name)
//		} else if f.InstanceParameter.Array != nil {
//			params = append(params, f.InstanceParameter.Array.CType+" "+f.InstanceParameter.Name)
//			args = append(args, f.InstanceParameter.Name)
//		} else {
//			panic(fmt.Sprintf("Do not know how to generate instance parameter %s for %s", f.InstanceParameter.Name, f.CIdentifier))
//		}
//	}
//
//	for _, param := range f.Parameters {
//		if param.isVarargsOrValist() {
//			args = append(args, "NULL")
//		} else if param.Type != nil {
//			params = append(params, param.Type.CType+" "+param.Name)
//			args = append(args, param.Name)
//		} else if param.Array != nil {
//			params = append(params, param.Array.CType+" "+param.Name)
//			args = append(args, param.Name)
//		} else {
//			panic(fmt.Sprintf("Do not know how to generate parameter %s for %s", param.Name, f.CIdentifier))
//		}
//	}
//
//	fnDecl := fmt.Sprintf(`
//static %s c_%s(%s) {
//    return %s(%s);
//}
//`,
//		f.ReturnValue.Type.CType,
//		f.CIdentifier,
//		strings.Join(params, ", "),
//		f.CIdentifier,
//		strings.Join(args, ", "),
//	)
//
//	fi.CgoPreamble(fnDecl)
//}
