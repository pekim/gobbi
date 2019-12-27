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
