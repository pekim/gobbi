package generate

import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
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
	Parameters Parameters `xml:"parameters>parameter"`
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
	f.sysName = "Fn_" + f.Name
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
	for _, param := range f.Parameters {
		if !param.isIn() {
			continue
		}

		goType := param.sysParamGoType()
		g.Id(param.goVarName).Add(goType)
	}
}

func (f *Function) generateSysBody(g *jen.Group) {
	if len(f.Parameters) > 0 {
		return
	}

	g.
		Qual("C", f.CIdentifier).
		Call()
}
