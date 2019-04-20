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
	cf.sortFunctionNames()
	fmt.Println("g_printf", cf.functionNameIndexes["g_printf"])

	filepath := projectFilepath("lib", "internal", "call", "call.c")
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	cf.file = file
	defer file.Close()

	cf.writeIncludeStatements()
	cf.writeLibraryGlobals()
	cf.writeFunctionGlobals()
	cf.writeOpenFunction()
	cf.writeCloseFunction()
	cf.writeGetFunctionFunction()
	cf.writeCallFunctionFunction()
}

func (cf *CallFile) writeIncludeStatements() {
	cf.write(`
#include <avcall.h>
#include <dlfcn.h>
#include <stdio.h>
#include <stdlib.h>
	`)
}

func (cf *CallFile) writeLibraryGlobals() {
	cf.write(`
void *library_handles[%d];
	`, len(cf.libraryNames))

	cf.write(`
char *library_names[] = {%s};
	`, "\n\t\""+strings.Join(cf.libraryNames, "\",\n\t\"")+"\"\n")
}

func (cf *CallFile) writeFunctionGlobals() {
	cf.write(`
void *functions[%d];
	`, len(cf.functionNames))

	cf.write(`
char *function_names[] = {%s};
	`, "\n\t\""+strings.Join(cf.functionNames, "\",\n\t\"")+"\"\n")
}

func (cf *CallFile) writeOpenFunction() {
	cf.write(`
char* open() {
    char *error;

	int i;
	for ( i = 0; i < %d; i++ ) {
		library_handles[i] = dlopen(library_names[i], RTLD_LAZY);
		if (!library_handles[i])
		{
			return dlerror();
		}
	}

    return NULL;
}
	`, len(cf.libraryNames))
}

func (cf *CallFile) writeCloseFunction() {
	cf.write(`
void close() {
	int i;
	for ( i = 0; i < %d; i++ ) {
	    dlclose(library_handles[i]);
	}
}
	`, len(cf.libraryNames))
}

func (cf *CallFile) writeGetFunctionFunction() {
	cf.write(`
void *get_function(int function_index) {
	char *function_name = function_names[function_index];
	void *fn = functions[function_index];
	if (fn) {
		return fn; 
	}

	int i;
	for ( i = 0; i < %d; i++ ) {
		fn = dlsym(library_handles[i], function_name);
		if (fn)
		{
			functions[function_index] = fn;
			return fn;
		}
	}

	return NULL;
}
	`, len(cf.libraryNames))
}

func (cf *CallFile) writeCallFunctionFunction() {
	cf.write(`
void call_function(int function_index) {
	void* fn = get_function(function_index);
	if (!fn) {
		printf("failed to load function %%d, %%s\n", function_index, function_names[function_index]);
		exit(1);
	}

	int ret;
	av_alist alist;
	//av_start_void (alist, fn);
	av_start_int (alist, fn, &ret);
	av_ptr(alist, char*, "qaz %%d\n");
	av_int(alist, 42);
	av_call(alist);
}
	`)
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
