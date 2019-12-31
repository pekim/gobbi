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

	if f.MovedTo != "" || f.ShadowedBy != "" {
		return
	}

	if f.version.GT(version) {
		return
	}

	// func Fn_some_function(...) [return type] {...}
	fi.
		Func().
		Id(f.sysName).
		ParamsFunc(f.generateLibParamsDeclaration).
		Do(f.generateLibReturnTypeDeclaration).
		BlockFunc(f.generateLibBody)

	fi.Line()
}

func (f *Function) generateLibParamsDeclaration(g *jen.Group) {
	//if f.InstanceParameter != nil {
	//	goType := f.InstanceParameter.sysParamGoType()
	//	g.Id("paramInstance").Add(goType)
	//
	//}
	//
	//for i, param := range f.Parameters {
	//	if param.isVarargsOrValist() {
	//		continue
	//	}
	//
	//	paramName := "param" + strconv.Itoa(i)
	//	goType := param.sysParamGoType()
	//
	//	g.Id(paramName).Add(goType)
	//}
	//
	//if f.Throws {
	//	g.Id("error").Add(jenUnsafePointer())
	//}
}

func (f *Function) generateLibReturnTypeDeclaration(s *jen.Statement) {
	if f.ReturnValue.isVoid() {
		return
	}

	//s.Add(f.ReturnValue.sysParamGoType())
}

func (f *Function) generateLibBody(g *jen.Group) {
	//f.generateSysCArgs(g)
	//
	//cIdentifier := f.CIdentifier
	//if f.Parameters.hasVaList() || f.Parameters.hasVarargs() {
	//	cIdentifier = "c_" + cIdentifier
	//}
	//
	//// [ret :=] C.somefunction(...)
	//g.
	//	Do(func(s *jen.Statement) {
	//		if !f.ReturnValue.isVoid() {
	//			s.Id("ret").Op(":=")
	//		}
	//	}).
	//	Qual("C", cIdentifier).
	//	CallFunc(f.generateSysCallParams)
	//
	//f.generateSysCArgsOut(g)
	//f.generateSysReturn(g)
}

//func (f *Function) generateLibCArgs(g *jen.Group) {
//	if f.InstanceParameter != nil {
//		f.InstanceParameter.generateLibCArg(g, "paramInstance", "cValueInstance")
//		g.Line()
//	}
//
//	for i, param := range f.Parameters {
//		if param.isVarargsOrValist() {
//			continue
//		}
//
//		paramName := "param" + strconv.Itoa(i)
//		cVarName := "cValue" + strconv.Itoa(i)
//
//		param.generateLibCArg(g, paramName, cVarName)
//		g.Line()
//	}
//
//	if f.Throws {
//		g.Id("cError").Op(":=").
//			Parens(jen.Op("**").Qual("C", "GError")).
//			Parens(jen.Id("error"))
//		g.Line()
//	}
//}
//
//func (f *Function) generateLibCArgsOut(g *jen.Group) {
//	for i, param := range f.Parameters {
//		paramName := "param" + strconv.Itoa(i)
//		cVarName := "cValue" + strconv.Itoa(i)
//
//		param.generateLibCArgOut(g, paramName, cVarName)
//	}
//}
//
//func (f *Function) generateLibCallParams(g *jen.Group) {
//	if f.InstanceParameter != nil {
//		g.Id("cValueInstance")
//	}
//
//	for i, param := range f.Parameters {
//		if param.isVarargsOrValist() {
//			continue
//		}
//
//		g.Id("cValue" + strconv.Itoa(i))
//	}
//
//	if f.Throws {
//		g.Id("cError")
//	}
//}
//
//func (f *Function) generateLibReturn(g *jen.Group) {
//	if f.ReturnValue.isVoid() {
//		return
//	}
//
//	g.Line()
//	retValue := f.ReturnValue.generateLibGoValue(g, "ret")
//	g.Return().Add(retValue)
//}
//
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
