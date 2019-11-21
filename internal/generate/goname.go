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
	if len(goName) > 0 {
		goName = strings.ToLower(goName[0:1]) + goName[1:]
	}

	if _, isReserved := goNameReservedWords[goName]; isReserved {
		goName += "_"
	}

	return goName
}

func makeGoName(cName string, uppercaseFirstChar bool) string {
	if len(cName) == 0 {
		return cName
	}

	//if uppercaseFirstChar{
	//	cName = strings.ToLower(goName[0:1]) + goName[1:]
	//}else {
	//
	//}

	cParts := strings.Split(cName, "_")
	if len(cParts) == 1 {
		cParts = strings.Split(cName, "-")
	}

	goParts := []string{}

	for i, cPart := range cParts {
		goPart := cPart
		if uppercaseFirstChar || i > 0 {
			goPart = strings.Title(goPart)
		}

		goParts = append(goParts, goPart)
	}

	return strings.Join(goParts, "")
}
