package generate

import "C"
import (
	"fmt"
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
	"strconv"
	"strings"
)

type Function struct {
	Name              string `xml:"name,attr"`
	Version           string `xml:"version,attr"`
	MovedTo           string `xml:"moved-to,attr"`
	ShadowedBy        string `xml:"shadowed-by,attr"`
	CIdentifier       string `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	Deprecated        int    `xml:"deprecated,attr"`
	DeprecatedVersion string `xml:"deprecated-version,attr"`
	//Doc               *Doc   `xml:"doc"`
	InstanceParameter *Parameter   `xml:"parameters>instance-parameter"`
	Parameters        Parameters   `xml:"parameters>parameter"`
	ReturnValue       *ReturnValue `xml:"return-value"`
	Throws            bool         `xml:"throws,attr"`
	Introspectable    string       `xml:"introspectable,attr"`

	namespace *Namespace
	blacklist bool
	version   semver.Version
	sysName   string
}

func (f *Function) init(ns *Namespace, record *Record, receiver bool) {
	f.namespace = ns
	f.applyAddenda()
	f.version = versionNew(f.Version)
	f.sysName = "Fn_" + f.CIdentifier
	if f.InstanceParameter != nil {
		f.InstanceParameter.init(ns)
	}
	f.Parameters.init(ns)
	f.ReturnValue.init(ns)
}

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

func (f *Function) isSupported() (bool, string) {
	if f.blacklist {
		return false, "blacklisted"
	}

	if supported, reason := f.Parameters.allSupported(); !supported {
		return false, reason
	}

	if supported, reason := f.ReturnValue.isSupported(); !supported {
		return false, reason
	}

	for _, param := range f.Parameters {
		if param.Array != nil && !param.Array.Type.isString() {
			return false, fmt.Sprintf("has non-string array param %s", param.Name)
		}
	}

	return true, ""
}

func (f *Function) generateSysParamsDeclaration(g *jen.Group) {
	if f.InstanceParameter != nil {
		goType := f.InstanceParameter.sysParamGoType()
		g.Id("paramInstance").Add(goType)

	}

	for i, param := range f.Parameters {
		if param.isVarargsOrValist() {
			continue
		}

		paramName := "param" + strconv.Itoa(i)
		goType := param.sysParamGoType()

		g.Id(paramName).Add(goType)
	}

	if f.Throws {
		g.Id("error").Add(jenUnsafePointer())
	}
}

func (f *Function) generateSysReturnTypeDeclaration(s *jen.Statement) {
	for _, param := range f.Parameters {
		if param.Array != nil && !param.Array.Type.isString() {
			return
		}
	}

	if supported, _ := f.Parameters.allSupported(); !supported {
		return
	}

	if supported, _ := f.ReturnValue.isSupported(); !supported {
		return
	}

	if f.ReturnValue.isVoid() {
		return
	}

	s.Add(f.ReturnValue.Type.sysParamGoType())
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
		f.InstanceParameter.generateSysCArg(g, "paramInstance", "cValueInstance")
		g.Line()
	}

	for i, param := range f.Parameters {
		if param.isVarargsOrValist() {
			continue
		}

		paramName := "param" + strconv.Itoa(i)
		cVarName := "cValue" + strconv.Itoa(i)

		param.generateSysCArg(g, paramName, cVarName)
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
	for i, param := range f.Parameters {
		paramName := "param" + strconv.Itoa(i)
		cVarName := "cValue" + strconv.Itoa(i)

		param.generateSysCArgOut(g, paramName, cVarName)
	}
}

func (f *Function) generateSysCallParams(g *jen.Group) {
	if f.InstanceParameter != nil {
		g.Id("cValueInstance")
	}

	for i, param := range f.Parameters {
		if param.isVarargsOrValist() {
			continue
		}

		g.Id("cValue" + strconv.Itoa(i))
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
	g.Return().Add(f.ReturnValue.generateSysGoValue("ret"))
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
