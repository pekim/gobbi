package generate

import (
	"github.com/blang/semver"
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
	//goName    string
}

func (a *Alias) init(ns *Namespace) {
	a.namespace = ns
	a.applyAddenda()
	a.version = versionNew(a.Version)
	//a.goName = makeExportedGoName(a.Name)
	a.Type.init(ns)
}

//func (a *Alias) generateSys(f *jen.File, version semver.Version) {
//	if a.blacklist {
//		f.Commentf("UNSUPPORTED : %s : blacklisted", a.Name)
//		return
//	}
//
//	if a.version.GT(version) {
//		return
//	}
//
//	// GEN: type SomeAlias SomeCType
//	f.Type().Id(a.Name).Qual("C", a.CType)
//}
