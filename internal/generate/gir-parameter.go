package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"regexp"
)

var cTypeRegex = regexp.MustCompile(" *(const)? *([a-zA-Z0-9 ]+) *(\\**)? *")

type Parameter struct {
	Name      string `xml:"name,attr"`
	Direction string `xml:"direction,attr"`
	Argument
	AllowNone bool `xml:"allow-none,attr"`
	//Doc       *Doc      `xml:"doc"`
	Varargs *struct{} `xml:"varargs"`

	goVarName string
	namespace *Namespace

	lengthParam    *Parameter
	lengthForParam *Parameter
	argcParam      *Parameter
	argvParam      *Parameter
}

func (p *Parameter) init(ns *Namespace) {
	p.namespace = ns
	p.Array.init(ns)
	p.goVarName = makeUnexportedGoName(p.Name)
}

func (p *Parameter) sysParamGoType() *jen.Statement {
	if p.Type != nil {
		parts := cTypeRegex.FindStringSubmatch(p.Type.CType)
		//isConst := parts[1]
		typ := parts[2]
		stars := parts[3]
		//indirectionCount := len(stars)
		if typ == "" {
			panic(fmt.Sprintf("Failed to parse type ; '%s'", p.Type.CType))
		}

		if simpleGoType, ok := simpleSysParamGoTypes[typ]; ok {
			return jen.Op(stars).Add(simpleGoType)
		}
		//fmt.Println("type", p.Type.CType)
	}

	return jen.Qual("github.com/pekim/gobbi/lib/internal/c", "UndefinedParamType")
}

var simpleSysParamGoTypes = map[string]*jen.Statement{
	"gunichar":    jen.Rune(),
	"int":         jen.Int(),
	"gint":        jen.Int(),
	"gint8":       jen.Int8(),
	"gint16":      jen.Int16(),
	"gint32":      jen.Int32(),
	"gint64":      jen.Int64(),
	"guint":       jen.Uint(),
	"guint8":      jen.Uint8(),
	"guint16":     jen.Uint16(),
	"guint32":     jen.Uint32(),
	"guint64":     jen.Uint64(),
	"glong":       jen.Int64(),
	"gulong":      jen.Uint64(),
	"gloat":       jen.Float32(),
	"gfloat":      jen.Float32(),
	"double":      jen.Float64(),
	"long double": jen.Float64(), // not ideal, but Go has nothing better
	"gdouble":     jen.Float64(),
	"gsize":       jen.Uint64(),
	"gssize":      jen.Uint64(),
	"gboolean":    jen.Bool(),
	"gpointer":    jen.Qual("unsafe", "Pointer"),
}
