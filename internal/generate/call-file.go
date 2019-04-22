package generate

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

type CallFile struct {
	file                *os.File
	packageNames        []string
	libraryNames        []string
	functionNames       []string
	functionNameIndexes map[string]int
}

func (cf *CallFile) generate() {
	cf.setLibraryNames()

	filepath := projectFilepath("lib", "internal", "call", "call_data.c")
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	cf.file = file

	cf.writeLibraryGlobals()
	cf.writeFunctionGlobals()

	err = file.Close()
	if err != nil {
		panic(err)
	}
}

func (cf *CallFile) writeLibraryGlobals() {
	cf.write(`
int library_count = %d;
`, len(cf.libraryNames))

	cf.write(`
void *library_handles[%d];
`, len(cf.libraryNames))

	cf.write(`
char *library_names[] = {%s};
	`, "\n\t\""+strings.Join(cf.libraryNames, "\",\n\t\"")+"\"\n")
}

func (cf *CallFile) writeFunctionGlobals() {
	cf.write(`
int function_count = %d;
	`, len(cf.functionNames))

	cf.write(`
void *functions[%d];
	`, len(cf.functionNames))

	cf.write(`
char *function_names[] = {%s};
	`, "\n\t\""+strings.Join(cf.functionNames, "\",\n\t\"")+"\"\n")
}

func (cf *CallFile) write(format string, args ...interface{}) {
	format = strings.TrimSuffix(format, "\t")

	_, err := cf.file.WriteString(fmt.Sprintf(format, args...))
	if err != nil {
		panic(err)
	}
}

func (cf *CallFile) setLibraryNames() {
	pkgConfigArgs := append([]string{"--libs"}, cf.packageNames...)
	pkgconfigOutput, err := exec.Command("pkg-config", pkgConfigArgs...).Output()
	if err != nil {
		fmt.Println("pkg-config", pkgConfigArgs)
		panic(err)
	}

	libraryNames := []string{}
	for _, outItem := range strings.Split(string(pkgconfigOutput), " ") {
		outItem = strings.TrimSpace(outItem)
		if !strings.HasPrefix(outItem, "-l") {
			continue
		}

		libraryName := fmt.Sprintf("lib%s.so", strings.TrimPrefix(outItem, "-l"))
		libraryNames = append(libraryNames, libraryName)
	}

	sort.Strings(libraryNames)
	cf.libraryNames = stringsUniq(libraryNames)
}

func (cf *CallFile) addFunctionName(name string) {
	cf.functionNames = append(cf.functionNames, name)
}

func (cf *CallFile) sortFunctionNames() {
	sort.Strings(cf.functionNames)

	cf.functionNameIndexes = make(map[string]int)
	for i, name := range cf.functionNames {
		cf.functionNameIndexes[name] = i
	}
}

func (cf *CallFile) getFunctionIndex(name string) int {
	return cf.functionNameIndexes[name]
}

func stringsUniq(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}
