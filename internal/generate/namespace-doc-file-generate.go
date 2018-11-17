package generate

func (ns *Namespace) generateDocForPackage() {
	generateDocFile(ns.writeDocForPackage,
		"content", "api", ns.goPackageName, "_index.md")
}

func (ns *Namespace) writeDocForPackage(file *DocFile) {
	file.writeFrontmatter(
		FrontmatterParam{"title", ns.goPackageName},
		FrontmatterParam{"goPackageName", ns.goPackageName},
		FrontmatterParam{"libraryName", ns.Name},
		FrontmatterParam{"cDocPath", ns.CDocPath},
	)
}
