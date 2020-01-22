package generate

import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
)

func (r Record) generateSysType(f *jen.File, version semver.Version) {
	if r.blacklist {
		f.Commentf("UNSUPPORTED : %s : blacklisted", r.Name)
		return
	}

	if !generateEntityForVersion(version, r.version) {
		return
	}

	// GEN: type SomeRecord SomeCType
	f.Type().Id(r.Name).Qual("C", r.CType)
}

func (r *Record) generateSys(f *jen.File, version semver.Version) {
	r.Constructors.generateSys(f, version)
	r.Methods.generateSys(f, version)
	r.Functions.generateSys(f, version)
}
