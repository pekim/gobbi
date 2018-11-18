package generate

type Enumerations []*Enumeration

func (ee Enumerations) init(ns *Namespace) {
	for _, enum := range ee {
		enum.init(ns)
	}
}

func (ee Enumerations) versionList() Versions {
	var versions Versions

	for _, e := range ee {
		if e.Version != "" {
			versions = append(versions, VersionNew(e.Version))
		}
	}

	return versions
}

func (ee Enumerations) entities() []Generatable {
	var generatables []Generatable

	for _, enum := range ee {
		generatables = append(generatables, enum)
	}

	return generatables
}

func (ee Enumerations) forName(name string) *Enumeration {
	for _, enum := range ee {
		if enum.Name == name {
			return enum
		}
	}

	return nil
}

func (ee Enumerations) mergeAddenda(addenda Enumerations) {
	for _, addendaEnum := range addenda {
		if enum := ee.forName(addendaEnum.Name); enum != nil {
			enum.mergeAddenda(addendaEnum)
		}
	}
}

func (ee Enumerations) generateDocs(df *DocFile, title string) {
	df.writeFrontmatter(
		FrontmatterParam{"title", title},
		//FrontmatterParam{"layout", "constants"},
		//FrontmatterParam{"type", "api"},
	)

	for _, e := range ee {
		if e.Blacklist {
			continue
		}

		df.writeLinef(`<p class="api-heading">%s</p>`, e.Name)
		df.writeDocTextLine(`<p class="api-doc">%s</p>`, e.Doc)
		df.writeLine(`<div class="api-notes">`)
		df.writeLinef(`  <p class="api-ctype">%s</p>`, e.CType)
		if e.Version != "" {
			df.writeLinef(`  <p class="api-since">since %s</p>`, e.Version)
		}

		df.writeLine("<table>")
		for _, m := range e.Members {
			df.writeLine("<tr>")
			df.writeLinef(`<td class="name">%s</td>`, m.goName)
			df.writeLinef(`<td class="value">%d</td>`, m.Value)
			df.writeDocTextLine(`<td class="doc">%s</td>`, m.Doc)
			df.writeLine("</tr>")
		}
		df.writeLine("</table>")

		df.writeLine(`</div>`)
	}
}
