package generate

import (
	"github.com/dave/jennifer/jen"
	"regexp"
	"strings"
)

var reSeeSomethingElse = regexp.MustCompile("## \\{#(.*)\\}")

type Doc struct {
	Text string `xml:",chardata"`
}

func generateDoc(d *Doc, g *jen.Group) {
	if d != nil {
		d.generate(g)
	}
}

func (d *Doc) generate(g *jen.Group) {
	if d.Text == "" {
		return
	}

	lines := strings.Split(d.Text, "\n")
	for l, line := range lines {
		line = strings.TrimSpace(line)

		line = strings.TrimPrefix(line, "## ")
		line = reSeeSomethingElse.ReplaceAllString(line, "\n\n(see $1)")

		lines[l] = line
	}

	for _, line := range lines {
		g.Comment(line)
	}
}
