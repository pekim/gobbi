package generate

import (
	"strings"
)

var goNameReservedWords = map[string]bool{
	"func":      true,
	"interface": true,
	"select":    true,
	"string":    true,
	"type":      true,
}

func makeExportedGoName(cName string) string {
	return makeGoName(cName, true)
}

func makeUnexportedGoName(cName string) string {
	goName := makeGoName(cName, false)

	if _, isReserved := goNameReservedWords[goName]; isReserved {
		goName += "_"
	}

	return goName
}

func makeGoName(cName string, uppercaseFirstChar bool) string {
	if len(cName) == 0 {
		return cName
	}

	cParts := strings.Split(cName, "_")
	if len(cParts) == 1 {
		cParts = strings.Split(cName, "-")
	}
	if cParts[0] == "" {
		cParts = cParts[1:]
	}

	goParts := []string{}

	for i, cPart := range cParts {
		goPart := cPart

		if i == 0 {
			if len(goPart) > 0 {
				if uppercaseFirstChar {
					goPart = strings.ToUpper(goPart[0:1]) + goPart[1:]
				} else {
					goPart = strings.ToLower(goPart[0:1]) + goPart[1:]
				}
			}
		} else {
			goPart = strings.Title(goPart)
		}

		goParts = append(goParts, goPart)
	}

	return strings.Join(goParts, "")
}
