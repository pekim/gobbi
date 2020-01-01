package generate

import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
)

type Interface struct {
	*Record
	//VirtualMethods Methods `xml:"virtual-method"`
}

func (i *Interface) generateSys(f *jen.File, version semver.Version) {
	i.Methods.generateSys(f, version)
}

func (i *Interface) generateLib(f *jen.File, version semver.Version) {
	if i.version.GT(version) {
		return
	}

	i.generateLibType(f, version)
	//i.Methods.generateLib(f, version)
}

func (i *Interface) generateLibType(f *jen.File, version semver.Version) {
	f.Commentf("%s is a representation of the C interface %s.", i.goName, i.CType)
	docVersion(f, i.Version)

	f.Type().Id(i.goName).Struct(
		jen.Id("native").Add(jenUnsafePointer()),
	)

	f.Line()
}
