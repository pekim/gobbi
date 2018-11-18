package generate

func (ns *Namespace) generateDoc() {
	generateDocFile(ns.writeDocForPackage,
		"content", "api", ns.goPackageName, "_index.md")

	generateDocFile(ns.Constants.generateDocs, "content", "api", ns.goPackageName, "constants.md")
	generateDocFile(func(df *DocFile) { ns.Bitfields.generateDocs(df, "bitfields") },
		"content", "api", ns.goPackageName, "bitfields.md")
	generateDocFile(func(df *DocFile) { ns.Enumerations.generateDocs(df, "enums") },
		"content", "api", ns.goPackageName, "enums.md")
}

func (ns *Namespace) writeDocForPackage(file *DocFile) {
	file.writeFrontmatter(
		FrontmatterParam{"title", ns.goPackageName},
		FrontmatterParam{"goPackageName", ns.goPackageName},
		FrontmatterParam{"libraryName", ns.Name},
		FrontmatterParam{"cDocPath", ns.CDocPath},
	)
}
