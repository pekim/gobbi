package generate

import (
	"fmt"
	"path"
)

func (ns *Namespace) generateTemplatedFiles() {

	filepath := path.Join(ns.libDir, "test.go")
	fmt.Println(filepath)
}
