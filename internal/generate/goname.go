package generate

import (
	"fmt"
	"strings"
	"unicode"
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

	if len(cName) > 0 && unicode.IsUpper(rune(goName[0])) {
		fmt.Println(cName, goName)
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
	//if cParts[0] == "" {
	//	cParts = cParts[1:]
	//}

	goParts := []string{}

	for _, cPart := range cParts {
		goPart := cPart

		//if i == 0 {
		//	if len(goPart) > 0 {
		//		if uppercaseFirstChar {
		//			goPart = strings.ToUpper(goPart[0:1]) + goPart[1:]
		//		} else {
		//			goPart = strings.ToLower(goPart[0:1]) + goPart[1:]
		//		}
		//	}
		//} else {
		goPart = strings.Title(goPart)
		//}

		goParts = append(goParts, goPart)
	}

	goName := strings.Join(goParts, "")
	if uppercaseFirstChar {
		goName = strings.ToUpper(goName[0:1]) + goName[1:]
	} else {
		//fmt.Println("1                                  ", goName)
		goName = strings.ToLower(goName[0:1]) + goName[1:]
		//fmt.Println("2  ", goName)
		if unicode.IsUpper(rune(goName[0])) {
			fmt.Println(cName, goName)
		}
	}

	return goName
}
