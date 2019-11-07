package generate

import (
	"fmt"
	"github.com/pekim/jennifer/jen"
)

type Field struct {
	Namespace *Namespace

	Name     string `xml:"name,attr"`
	Writable int    `xml:"writable,attr"`
	Bits     int    `xml:"bits,attr"`
	Private  bool   `xml:"private,attr"`
	Doc      *Doc   `xml:"doc"`
	Type     *Type  `xml:"type"`

	goVarName string
}

func (f *Field) init(ns *Namespace) {
	f.Namespace = ns
	f.goVarName = MakeExportedGoName(f.Name)
	f.Name = makeSafeCName(f.Name)

	if f.Type != nil {
		f.Type.init(ns)
	}
}

func (f Field) supported() (bool, string) {
	if f.Private {
		return false, fmt.Sprintf("Private : %s", f.Name)
	}

	if f.Bits > 0 {
		return false, fmt.Sprintf("Bitfield not supported : %2d %s", f.Bits, f.Name)
	}

	if f.Type == nil {
		return false, fmt.Sprintf("no type for %s", f.Name)
	}

	if f.Type.generator == nil {
		return false, fmt.Sprintf("%s : no type generator for %s, %s", f.Name, f.Type.Name, f.Type.CType)
	}

	if supportedAsField, reason := f.Type.generator.isSupportedAsField(); !supportedAsField {
		return false, fmt.Sprintf("%s : %s", f.Name, reason)
	}

	return true, ""
}

func (f Field) generate(g *jen.Group) {
	supported, reason := f.supported()
	if !supported {
		g.Comment(reason)
		return
	}

	f.Type.generator.generateDeclaration(g, f.goVarName)
}

func (f *Field) mergeAddenda(addenda *Field) {
	f.Private = addenda.Private
}
