package generate

func (ns *Namespace) generateDoc() {
	generateDocFile(ns.writeDocForPackage,
		"content", "api", ns.goPackageName, "_index.md")

	generateDocFile(ns.Constants.generateDocs, "content", "api", ns.goPackageName, "constants.md")
}

func (ns *Namespace) writeDocForPackage(file *DocFile) {
	file.writeFrontmatter(
		FrontmatterParam{"title", ns.goPackageName},
		FrontmatterParam{"goPackageName", ns.goPackageName},
		FrontmatterParam{"libraryName", ns.Name},
		FrontmatterParam{"cDocPath", ns.CDocPath},
	)
}
