package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

type Callback struct {
	*Function
	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	varNameId        string
	varNameMap       string
	varNameLock      string
	goNameHandler    string
	callbackTypeName string
}

func (c *Callback) init(ns *Namespace) {
	c.Function.init(ns, nil, "")
	c.initNames()
}

func (c *Callback) initNames() {
	callbackGoName := makeExportedGoName(c.Name)

	c.varNameId = fmt.Sprintf("callback%sId", callbackGoName)
	c.varNameMap = fmt.Sprintf("callback%sMap", callbackGoName)
	c.varNameLock = fmt.Sprintf("callback%sLock", callbackGoName)

	c.goNameHandler = fmt.Sprintf("callback_%sHandler",
		makeGoNameInternal(c.Name, false))

	c.callbackTypeName = fmt.Sprintf("%sCallback",
		callbackGoName)
}

func (c *Callback) supported() (supported bool, reason string) {
	for _, p := range c.Parameters {
		if p.Name == "data" || p.Name == "user_data" {
			continue
		}

		if supported, reason := p.isSupported(); !supported {
			return supported, fmt.Sprintf("callback %s : unsupported parameter %s : %s", c.Name, p.Name, reason)
		}

		if p.Type != nil && p.Type.Name == "ignore" {
			return supported, fmt.Sprintf("callback %s : unsupported ignore parameter %s", c.Name, p.Name)
		}
	}

	if supported, reason := c.ReturnValue.isSupported(); !supported {
		return false, fmt.Sprintf("callback %s : %s", c.Name, reason)
	}

	if c.ReturnValue.Array != nil {
		return supported, fmt.Sprintf("callback %s : unsupported return array", c.Name)
	}

	return true, ""
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

	c.generateCgoPreamble()
	c.generateVariables(g)
	c.generateCallbackType(g)
	//c.generateConnectFunction(g)
	//c.generateDisconnectFunction(g)
	//c.generateHandlerFunction(g)
}

func (c *Callback) generateCgoPreamble() {
	handlerReturn := c.cTypeDeclaration(c.ReturnValue.Type)

	params := []string{"GObject *"}
	for _, param := range c.Parameters {
		if param.Array != nil {
			params = append(params, param.Array.CType)
		} else {
			params = append(params, c.cTypeDeclaration(param.Type))
		}
	}
	params = append(params, "gpointer")
	handlerParams := strings.Join(params, ", ")

	c.Namespace.jenFile.CgoPreamble(
		fmt.Sprintf(`
	%s %s(%s);

`,
			handlerReturn,
			c.goNameHandler,
			handlerParams,
		))
}

func (c *Callback) generateVariables(g *jen.Group) {
	g.Var().Id(c.varNameId).Int()
	g.Var().Id(c.varNameMap).Op("=").Make(jen.Map(jen.Int()).Id(c.callbackTypeName))
	g.Var().Id(c.varNameLock).Qual("sync", "RWMutex")

	g.Line()
}

func (c *Callback) cTypeDeclaration(typ *Type) string {
	cDeclaration := typ.CType
	qname := QNameNew(c.Namespace, typ.Name)

	if cDeclaration == "" {
		record, found := qname.namespace.recordOrClassRecordForName(qname.name)
		if found {
			return record.CType + " *"
		}

		iface, found := qname.namespace.interfaceForName(qname.name)
		if found {
			return iface.CType + " *"
		}

		bitfield, found := qname.namespace.bitfieldForName(qname.name)
		if found {
			return bitfield.CType
		}

		enum, found := qname.namespace.enumForName(qname.name)
		if found {
			return enum.CType
		}

		panic(fmt.Sprintf("Not found bitfield, class, enum interface, or record %s, for callback %s",
			qname.name, c.Name))
	}

	return cDeclaration
}

func (c *Callback) generateCallbackType(g *jen.Group) {
	params := Parameters{}
	for _, p := range c.Parameters {
		if p.Name == "data" || p.Name == "user_data" {
			continue
		}

		params = append(params, p)
	}

	g.Commentf("%s is a callback function for a '%s' callback.",
		c.callbackTypeName,
		c.Name)

	g.
		Type().
		Id(c.callbackTypeName).
		Func().
		ParamsFunc(params.generateFunctionDeclaration).
		ParamsFunc(c.generateCallbackReturnDeclaration)

	g.Line()
}

func (c *Callback) generateCallbackReturnDeclaration(g *jen.Group) {
	c.ReturnValue.generateFunctionDeclaration(g)
}
