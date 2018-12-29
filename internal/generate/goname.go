package generate

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

var cNameReservedWords = map[string]bool{
	"func":      true,
	"interface": true,
	"select":    true,
	"type":      true,
}

var goNameReservedWords = map[string]bool{
	"func":      true,
	"interface": true,
	"select":    true,
	"string":    true,
	"type":      true,
}

func makeSafeCName(cName string) string {
	if _, isReserved := cNameReservedWords[cName]; isReserved {
		return "_" + cName
	}

	return cName
}

func makeGoName(cName string) string {
	name := makeGoNameInternal(cName, false)

	if _, isReserved := goNameReservedWords[name]; isReserved {
		name += "_"
	}

	return name
}

func makeExportedGoName(cName string) string {
	return makeGoNameInternal(cName, true)
}

func makeGoNameInternal(cName string, uppercaseFirstChar bool) string {
	cParts := strings.Split(cName, "_")
	if len(cParts) == 1 {
		cParts = strings.Split(cName, "-")
	}

	goParts := []string{}

	for i, cPart := range cParts {
		goPart := strings.ToLower(cPart)
		if uppercaseFirstChar || i > 0 {
			goPart = strings.Title(goPart)
		}

		goParts = append(goParts, goPart)
	}

	return strings.Join(goParts, "")
}

func lowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}
