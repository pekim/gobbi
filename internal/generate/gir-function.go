package generate

import "C"
import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
	"strconv"
)

type Function struct {
	Name              string `xml:"name,attr"`
	Version           string `xml:"version,attr"`
	MovedTo           string `xml:"moved-to,attr"`
	CIdentifier       string `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	Deprecated        int    `xml:"deprecated,attr"`
	DeprecatedVersion string `xml:"deprecated-version,attr"`
	//Doc               *Doc   `xml:"doc"`
	//InstanceParameter *Parameter   `xml:"parameters>instance-parameter"`
	InstanceParameter *Parameter `xml:"parameters>instance-parameter"`
	Parameters        Parameters `xml:"parameters>parameter"`
	//ReturnValue    *ReturnValue `xml:"return-value"`
	Throws         int    `xml:"throws,attr"`
	Introspectable string `xml:"introspectable,attr"`

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
}

func (f *Function) generateSys(fi *jen.File, version semver.Version) {
	if f.blacklist {
		fi.Commentf("UNSUPPORTED : %s : blacklisted", f.Name)
		return
	}

	if supported, reason := f.Parameters.allSupported(); !supported {
		fi.Commentf("UNSUPPORTED : %s : %s", f.Name, reason)
		fi.Line()
		return
	}

	if f.version.GT(version) {
		return
	}

	fi.
		Func().
		Id(f.sysName).
		ParamsFunc(f.generateSysParamDeclaration).
		BlockFunc(f.generateSysBody)

	fi.Line()
}

func (f *Function) generateSysParamDeclaration(g *jen.Group) {
	if f.InstanceParameter != nil {
		goType := f.InstanceParameter.sysParamGoType()
		g.Id("paramInstance").Add(goType)

	}

	for i, param := range f.Parameters {
		paramName := "param" + strconv.Itoa(i)
		goType := param.sysParamGoType()

		g.Id(paramName).Add(goType)
	}
}

func (f *Function) generateSysBody(g *jen.Group) {
	for _, param := range f.Parameters {
		if param.Array != nil {
			g.Comment("has array param")
			return
		}
	}

	f.generateSysCArgs(g)

	g.
		Qual("C", f.CIdentifier).
		CallFunc(f.generateSysCallParams)
}

func (f *Function) generateSysCArgs(g *jen.Group) {
	if f.InstanceParameter != nil {
		f.generateSysCArg(g, f.InstanceParameter, "paramInstance", "cValueInstance")
	}

	for i, param := range f.Parameters {
		paramName := "param" + strconv.Itoa(i)
		cVarName := "cValue" + strconv.Itoa(i)

		f.generateSysCArg(g, param, paramName, cVarName)
	}

	g.Line()
}

func (f *Function) generateSysCArg(g *jen.Group, param *Parameter, paramName string, cVarName string) {
	cValue := param.generateSysCValue(paramName)
	g.Id(cVarName).Op(":=").Add(cValue)
}

func (f *Function) generateSysCallParams(g *jen.Group) {
	if f.InstanceParameter != nil {
		g.Id("cValueInstance")
	}

	for i, _ := range f.Parameters {
		g.Id("cValue" + strconv.Itoa(i))
	}
}
