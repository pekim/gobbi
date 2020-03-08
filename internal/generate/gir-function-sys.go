package generate

import "C"
import (
	"fmt"
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
	"strings"
)

func (f *Function) generateSys(fi *jen.File, version semver.Version) {
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

	if f.Parameters.hasVarargs() || f.Parameters.hasVaList() {
		f.generateSysVarArgsCFunction(fi)
	}

	// func Fn_some_function(...) [return type] {...}
	fi.
		Func().
		Id(f.sysName).
		ParamsFunc(f.generateSysParamsDeclaration).
		Do(f.generateSysReturnTypeDeclaration).
		BlockFunc(f.generateSysBody)

	fi.Line()
}

func (f *Function) generateSysParamsDeclaration(g *jen.Group) {
	if f.InstanceParameter != nil {
		goType := f.InstanceParameter.sysParamGoType()
		g.Id(f.InstanceParameter.goVarName).Add(goType)

	}

	for _, param := range f.Parameters {
		if param.isVarargsOrValist() {
			continue
		}

		goType := param.sysParamGoType()

		g.Id(param.goVarName).Add(goType)
	}

	if f.Throws {
		g.Id("error").Add(jenUnsafePointer())
	}
}

func (f *Function) generateSysReturnTypeDeclaration(s *jen.Statement) {
	if f.ReturnValue.isVoid() {
		return
	}

	s.Add(f.ReturnValue.sysParamGoType())
}

func (f *Function) generateSysBody(g *jen.Group) {
	f.generateSysCArgs(g)

	cIdentifier := f.CIdentifier
	if f.Parameters.hasVaList() || f.Parameters.hasVarargs() {
		cIdentifier = "c_" + cIdentifier
	}

	// [ret :=] C.somefunction(...)
	g.
		Do(func(s *jen.Statement) {
			if !f.ReturnValue.isVoid() {
				s.Id("ret").Op(":=")
			}
		}).
		Qual("C", cIdentifier).
		CallFunc(f.generateSysCallParams)

	f.generateSysCArgsOut(g)
	f.generateSysReturn(g)
}

func (f *Function) generateSysCArgs(g *jen.Group) {
	if f.InstanceParameter != nil {
		f.InstanceParameter.generateSysCArg(g)
		g.Line()
	}

	for _, param := range f.Parameters {
		if param.isVarargsOrValist() {
			continue
		}

		param.generateSysCArg(g)
		g.Line()
	}

	if f.Throws {
		g.Id("cError").Op(":=").
			Parens(jen.Op("**").Qual("C", "GError")).
			Parens(jen.Id("error"))
		g.Line()
	}
}

func (f *Function) generateSysCArgsOut(g *jen.Group) {
	for _, param := range f.Parameters {
		param.generateSysCArgOut(g, param.goVarName, param.cVarName)
	}
}

func (f *Function) generateSysCallParams(g *jen.Group) {
	if f.InstanceParameter != nil {
		g.Id(f.InstanceParameter.cVarName)
	}

	for _, param := range f.Parameters {
		if param.isVarargsOrValist() {
			continue
		}

		g.Id(param.cVarName)
	}

	if f.Throws {
		g.Id("cError")
	}
}

func (f *Function) generateSysReturn(g *jen.Group) {
	if f.ReturnValue.isVoid() {
		return
	}

	g.Line()
	retValue := f.ReturnValue.generateSysGoValue(g, "ret")
	g.Return().Add(retValue)
}

func (f *Function) generateSysVarArgsCFunction(fi *jen.File) {
	params := []string{}
	args := []string{}

	if f.InstanceParameter != nil {
		if f.InstanceParameter.Type != nil {
			params = append(params, f.InstanceParameter.Type.CType+" "+f.InstanceParameter.Name)
			args = append(args, f.InstanceParameter.Name)
		} else if f.InstanceParameter.Array != nil {
			params = append(params, f.InstanceParameter.Array.CType+" "+f.InstanceParameter.Name)
			args = append(args, f.InstanceParameter.Name)
		} else {
			panic(fmt.Sprintf("Do not know how to generate instance parameter %s for %s", f.InstanceParameter.Name, f.CIdentifier))
		}
	}

	for _, param := range f.Parameters {
		if param.isVarargsOrValist() {
			args = append(args, "NULL")
		} else if param.Type != nil {
			params = append(params, param.Type.CType+" "+param.Name)
			args = append(args, param.Name)
		} else if param.Array != nil {
			params = append(params, param.Array.CType+" "+param.Name)
			args = append(args, param.Name)
		} else {
			panic(fmt.Sprintf("Do not know how to generate parameter %s for %s", param.Name, f.CIdentifier))
		}
	}

	fnDecl := fmt.Sprintf(`
static %s c_%s(%s) {
    return %s(%s);
}
`,
		f.ReturnValue.Type.CType,
		f.CIdentifier,
		strings.Join(params, ", "),
		f.CIdentifier,
		strings.Join(args, ", "),
	)

	fi.CgoPreamble(fnDecl)
}
