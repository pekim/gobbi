package generate

import "strings"

func makeExportedGoName(cName string) string {
	return makeUnexportedGoName(cName, true)
}

func makeUnexportedGoName(cName string, uppercaseFirstChar bool) string {
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
