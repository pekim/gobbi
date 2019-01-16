package generate

import (
	"github.com/dave/jennifer/jen"
)

type Callback struct {
	*Function
	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
}

func (c *Callback) generate(g *jen.Group, version *Version) {
	if !supportedByVersion(c, version) {
		return
	}

	supported, reason := c.supported()
	if !supported {
		g.Comment(reason)
		return
	}

	g.Commentf("potentially support callback : %s", c.Name)

	//if blacklisted, detail := f.blacklisted(); blacklisted {
	//	g.Commentf("Blacklisted : %s", detail)
	//	g.Line()
	//	return
	//}
	//
	//if f.Parameters.hasFormatArgs() {
	//	f.generateFormatArgsCWrapper()
	//}
	//
	//g.Commentf("%s is a wrapper around the C function %s.", f.GoName, f.CIdentifier)
	//
	//g.
	//	Func().
	//	Do(f.generateReceiverDeclaration).
	//	Id(f.GoName).                                         // name
	//	ParamsFunc(f.Parameters.generateFunctionDeclaration). // params
	//	ParamsFunc(f.generateReturnDeclaration).              // returns
	//	BlockFunc(f.generateBody).                            // body
	//	Line()
}
