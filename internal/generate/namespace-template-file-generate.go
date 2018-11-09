package generate

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

func (ns *Namespace) generateTemplatedFiles() {
	templateDir := projectFilepath("internal", "template")
	fileInfos, err := ioutil.ReadDir(templateDir)

	if err != nil {
		panic(err)
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			panic(fmt.Sprintf("Unexpected template dir, %", fileInfo.Name()))
		}

		// read template file
		filepath := path.Join(templateDir, fileInfo.Name())
		templateBytes, err := ioutil.ReadFile(filepath)
		if err != nil {
			panic(err)
		}

		// replace package with the library's package name
		packageLine := fmt.Sprintf("package %s\n", ns.goPackageName)
		template := strings.Replace(string(templateBytes), "package template\n", packageLine, 1)

		libraryFilepath := path.Join(ns.libDir, fileInfo.Name())
		err = ioutil.WriteFile(libraryFilepath, []byte(template), 0644)
		if err != nil {
			panic(err)
		}
	}
}
