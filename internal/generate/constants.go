package generate

type Constants []*Constant

func (cc Constants) init(ns *Namespace) {
	for _, constant := range cc {
		constant.init(ns)
	}
}

func (cc Constants) versionList() Versions {
	var versions Versions

	for _, c := range cc {
		if c.Version != "" {
			versions = append(versions, VersionNew(c.Version))
		}
	}

	return versions
}

func (cc Constants) entities() []Generatable {
	var generatables []Generatable

	for _, constant := range cc {
		generatables = append(generatables, constant)
	}

	return generatables
}

func (cc Constants) forName(name string) *Constant {
	for _, constant := range cc {
		if constant.Name == name {
			return constant
		}
	}

	return nil
}

func (cc Constants) mergeAddenda(addenda Constants) {
	for _, addendaContant := range addenda {
		if constant := cc.forName(addendaContant.Name); constant != nil {
			constant.mergeAddenda(addendaContant)
		}
	}
}

func (cc Constants) generateDocs(df *DocFile) {
	df.writeFrontmatter(
		FrontmatterParam{"title", "constants"},
		//FrontmatterParam{"layout", "constants"},
		//FrontmatterParam{"type", "api"},
	)

	for _, c := range cc {
		if c.Blacklist {
			continue
		}

		df.writeLinef(`<p class="api-heading">%s</p>`, c.Name)

		if c.Doc != nil && c.Doc.Text != "" {
			df.writeLinef(`<p class="api-doc">%s</p>`, c.Doc.Text)
		}

		df.writeLine(`<div class="api-notes">`)
		df.writeLinef(`  <p class="api-ctype">%s</p>`, c.CType)
		if c.Version != "" {
			df.writeLinef(`  <p class="api-since">since %s</p>`, c.Version)
		}
		df.writeLine(`</div>`)
	}
}
