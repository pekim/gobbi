package main

import (
	"fmt"
	"github.com/boyter/scc/processor"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func outputScc() {
	fmt.Printf("%4s %9s %9s %9s %9s %9s\n", "type", "files", "lines", "code", "comment", "blank")
	outputSccForType("go")
	outputSccForType("c")
}

func outputSccForType(typ string) {
	processor.ProcessConstants()

	var fileCount int
	var allFiles processor.FileJob

	err := filepath.Walk("./lib", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path, "."+typ) {
			c(path, &allFiles)
			fileCount++
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%4s %9d %9d %9d %9d %9d\n",
		strings.Title(typ), fileCount, allFiles.Lines, allFiles.Code, allFiles.Comment, allFiles.Blank)
}

func c(path string, allFiles *processor.FileJob) {
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
}
