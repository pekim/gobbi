package generate

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const bookDir = "gitbook"

func initialiseDocsSummaryFile() {
	filename := projectFilepath(bookDir, "SUMMARY.md")
	baseFilename := projectFilepath(bookDir, "SUMMARY-base.md")

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	basefile, err := os.Open(baseFilename)
	if err != nil {
		panic(err)
	}
	defer basefile.Close()

	_, err = io.Copy(file, basefile)
	if err != nil {
		panic(err)
	}
}

func appendDocsSummaryFile(level int, name string, docFilepath string) {
	relativePath, err := filepath.Rel(
		projectFilepath(bookDir),
		docFilepath,
	)
	if err != nil {
		panic(err)
	}

	summaryFilename := projectFilepath(bookDir, "SUMMARY.md")
	f, err := os.OpenFile(summaryFilename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	indent := strings.Repeat("  ", level)
	line := fmt.Sprintf("%s- [%s](./%s)\n",
		indent, name, relativePath,
	)

	_, err = f.WriteString(line)
	if err != nil {
		panic(err)
	}

	f.Close()
}
