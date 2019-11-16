package generate

import "strings"

var goNameReservedWords = map[string]bool{
	"func":      true,
	"interface": true,
	"select":    true,
	"string":    true,
	"type":      true,
}

func makeExportedGoName(cName string) string {
	return makeUnexportedGoName(cName, true)
}

func makeUnexportedGoName(cName string, uppercaseFirstChar bool) string {
	goName := makeGoName(cName, uppercaseFirstChar)

	if _, isReserved := goNameReservedWords[goName]; isReserved {
		goName += "_"
	}

	return goName
}

func makeGoName(cName string, uppercaseFirstChar bool) string {
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
