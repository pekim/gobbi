package main

import (
	"fmt"
	"github.com/boyter/scc/processor"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func outputFileStats() {
	processor.ProcessConstants()

	var fileCount int
	var allFiles processor.FileJob

	err := filepath.Walk("../lib", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		content, _ := ioutil.ReadFile(path)
		if !strings.Contains(string(content), "// Code generated - DO NOT EDIT") {
			return nil
		}

		c(content, &allFiles)
		fileCount++

		return nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("%9s %9s %9s %9s %9s\n", "files", "lines", "code", "comment", "blank")
	fmt.Printf("%9d %9d %9d %9d %9d\n", fileCount, allFiles.Lines, allFiles.Code, allFiles.Comment, allFiles.Blank)
	fmt.Println()
}

func c(content []byte, allFiles *processor.FileJob) {
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
