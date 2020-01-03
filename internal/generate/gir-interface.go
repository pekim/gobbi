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

	generateLibStructType(f, "interface", i.Name, i.CType, i.Version)
	generateLibToC(f, i.Name, i.CType)
	//i.Methods.generateLib(f, version)
}
