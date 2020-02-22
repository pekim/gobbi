package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"regexp"
	"strings"
)

var cTypeRegex = regexp.MustCompile(" *(const |volatile )* *([a-zA-Z0-9_ ]+) *(.*)?")

type cType struct {
	typ              string
	stars            string
	indirectionCount int
}

func parseCtype(raw string) cType {
	parts := cTypeRegex.FindStringSubmatch(raw)
	if parts == nil {
		panic(fmt.Sprintf("Failed to parse type ; '%s'", raw))
	}

	ct := cType{}
	ct.typ = parts[2]
	ct.indirectionCount = strings.Count(parts[3], "*")
	ct.stars = strings.Repeat("*", ct.indirectionCount)
	if ct.typ == "" {
		panic(fmt.Sprintf("Failed to parse type ; '%s'", raw))
	}

	return ct
}

func (ct cType) jenCgoType() *jen.Statement {
	return jen.Parens(jen.Op(ct.stars).Qual("C", ct.typ))
}
