package generate

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (ns *Namespace) generateDocDir() {
	err := os.MkdirAll(ns.docDir, 0775)
	if err != nil {
		panic(err)
	}
}

func (ns *Namespace) generateDocForPackage() {
	content := fmt.Sprintf(`
+++
title="%s"
type="page"
+++

The %s package provides Go bindings for the Gnome %s library.
`, ns.goPackageName, "`"+ns.goPackageName+"`", ns.Name)

	filename := filepath.Join(ns.docDir, "_index.md")
	ioutil.WriteFile(filename, []byte(content), 0664)
}
