package main

import (
	"fmt"
	"github.com/boyter/scc/processor"
	"github.com/pekim/gobbi/internal/generate"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	generateLibraries()
	outputScc()
}

func generateLibraries() {
	start := time.Now()
	generate.FromRoot("Gtk", "3.0")
	fmt.Printf("\ngeneration %.0fms\n\n", time.Now().Sub(start).Seconds()*1000)
}

func outputScc() {
	processor.ProcessConstants()

	var fileCount int
	var allFiles processor.FileJob

	filepath.Walk("./lib", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		content, _ := ioutil.ReadFile(path)
		filejob := processor.FileJob{
			Language: "Go",
			Content:  content,
		}
		processor.CountStats(&filejob)

		allFiles.Blank += filejob.Blank
		allFiles.Code += filejob.Code
		allFiles.Comment += filejob.Comment
		allFiles.Lines += filejob.Lines
		fileCount++

		return nil
	})

	fmt.Printf("%9s %9s %9s %9s %9s\n", "files", "lines", "code", "comment", "blank")
	fmt.Printf("%9d %9d %9d %9d %9d\n", fileCount, allFiles.Lines, allFiles.Code, allFiles.Comment, allFiles.Blank)
}
