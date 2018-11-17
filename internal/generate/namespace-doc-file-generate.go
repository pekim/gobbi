package generate

import (
	"fmt"
)

func (ns *Namespace) generateDocForPackage() {
	content := fmt.Sprintf(`
+++
title="%s"
type="page"
+++

The %s package provides Go bindings for the Gnome %s library.
`, ns.goPackageName, "`"+ns.goPackageName+"`", ns.Name)

	generateDocFile(func(file *DocFile) {
		file.write(content)
	}, "content", "api", ns.goPackageName, "_index.md")
}
