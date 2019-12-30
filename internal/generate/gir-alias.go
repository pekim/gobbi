package generate

import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
)

type Alias struct {
	Name  string `xml:"name,attr"`
	CType string `xml:"type,attr"`
	Type  *Type  `xml:"type"`
	//Doc   *Doc   `xml:"doc"`
	Version string

	blacklist bool
	namespace *Namespace
	version   semver.Version
	goName    string
}

func (a *Alias) init(ns *Namespace) {
	a.namespace = ns
	a.applyAddenda()
	a.version = versionNew(a.Version)
	a.goName = makeExportedGoName(a.Name)
	a.Type.init(ns)
}

func (a *Alias) generateLib(f *jen.File, version semver.Version) {
	if a.blacklist {
		f.Commentf("UNSUPPORTED : %s : blacklisted", a.Name)
		f.Line()
		return
	}

	if a.version.GT(version) {
		return
	}

	goType := a.Type.libParamGoType(false)
	//f.docForC(a.goName, a.Name)
	f.
		Type().
		Id(a.goName).
		Add(goType)

	f.Line()
}
