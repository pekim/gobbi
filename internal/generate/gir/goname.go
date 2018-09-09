package gir

import "strings"

var reservedWords = map[string]bool{
	"func": true,
	"type": true,
}

func makeGoName(cName string) string {
	name := makeGoNameInternal(cName, false)

	if _, isReserved := reservedWords[name]; isReserved {
		name += "_"
	}

	return name
}

func makeExportedGoName(cName string) string {
	return makeGoNameInternal(cName, true)
}

func makeGoNameInternal(cName string, uppercaseFirstChar bool) string {
	cParts := strings.Split(cName, "_")
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
