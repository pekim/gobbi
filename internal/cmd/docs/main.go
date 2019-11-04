package main

import (
	"github.com/gomarkdown/markdown"
	"io/ioutil"
	"log"
	"os"
	"path"
	"text/template"
)

const srcDir = "./docs-src"
const destDir = "../docs"

type Page struct {
	File  string
	Title string
}

var pages = []Page{
	Page{File: "index", Title: "Introduction"},
	Page{File: "getting-started", Title: "Getting started"},
	Page{File: "application-lifecycle", Title: "Application lifecycle"},
	Page{File: "build-tags", Title: "Build tags"},
	Page{File: "goroutines", Title: "Goroutines"},
	Page{File: "casting", Title: "Casting"},
	Page{File: "signal-handling", Title: "Signal handling"},
	Page{File: "variadic-functions", Title: "Variadic functions"},
	Page{File: "gvalue", Title: "gobject.Value"},
	Page{File: "reference-counting", Title: "Reference counting"},
	Page{File: "subclassing", Title: "Subclassing"},
	Page{File: "api", Title: "API docs"},
}

func main() {
	for _, page := range pages {
		generateHtmlFile(page)
	}
}

func generateHtmlFile(page Page) {
	template, err := template.ParseFiles(
		path.Join(srcDir, "_template.html"),
		path.Join(srcDir, "_navigation.html"),
		path.Join(srcDir, "_header.html"),
	)
	errorIsFatal(err)

	md, err := ioutil.ReadFile(path.Join(srcDir, page.File+".md"))
	errorIsFatal(err)

	content := string(markdown.ToHTML(md, nil, nil))

	out, err := os.OpenFile(path.Join(destDir, page.File+".html"), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	errorIsFatal(err)
	defer out.Close()

	type model struct {
		Title   string
		Pages   []Page
		Content string
	}
	err = template.ExecuteTemplate(out, "_template.html", model{
		Title:   page.Title,
		Pages:   pages,
		Content: content,
	})
	errorIsFatal(err)
}

func errorIsFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
