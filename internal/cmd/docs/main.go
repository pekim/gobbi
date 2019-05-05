package main

import (
	"github.com/gomarkdown/markdown"
	"io/ioutil"
	"log"
	"path"
	"strings"
)

const srcDir = "./docs-src/content"
const destDir = "./docs"

type page struct {
	file  string
	title string
}

func main() {
	pages := []page{
		page{file: "index", title: "Introduction"},
		page{file: "getting-started", title: "Getting started"},
		page{file: "application-lifecycle", title: "Application lifecycle"},
		page{file: "build-tags", title: "Build tags"},
		page{file: "goroutines", title: "Goroutines"},
		page{file: "gvalue", title: "gobject.Value"},
		page{file: "casting", title: "Casting"},
		page{file: "signal-handling", title: "Signal handling"},
		page{file: "variadic-functions", title: "Variadic functions"},
		page{file: "reference-counting", title: "Reference counting"},
		page{file: "api", title: "API docs"},
	}

	for _, page := range pages {
		generateHtmlFile(page)
	}
}

func generateHtmlFile(page page) {
	template, err := ioutil.ReadFile(path.Join(srcDir, "template.html"))
	errorIsFatal(err)

	md, err := ioutil.ReadFile(path.Join(srcDir, page.file+".md"))
	errorIsFatal(err)

	content := string(markdown.ToHTML(md, nil, nil))

	html := string(template)
	html = strings.ReplaceAll(html, "{title}", page.title)
	html = strings.Replace(html, "{content}", content, 1)

	ioutil.WriteFile(path.Join(destDir, page.file+".html"), []byte(html), 0644)
}

func errorIsFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
