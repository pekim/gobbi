package generate

import (
	"fmt"
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

	g.Commentf("potentially supported callback : %s", c.Name)
	g.Line()
}

func (c *Callback) supported() (supported bool, reason string) {
	for _, p := range c.Parameters {
		if p.Name == "data" || p.Name == "user_data" {
			continue
		}

		if supported, reason := p.isSupported(); !supported {
			return supported, fmt.Sprintf("callback %s : unsupported parameter %s : %s", c.Name, p.Name, reason)
		}
	}

	if supported, reason := c.ReturnValue.isSupported(); !supported {
		return false, fmt.Sprintf("%s : %s", c.Name, reason)
	}

	return true, ""
}
