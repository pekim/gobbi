package generate

import (
	"fmt"
	"os"
	"path"
)

type DocFile struct {
	file *os.File
	path string
}

func (f *DocFile) write(text string) {
	_, err := f.file.WriteString(text)
	if err != nil {
		panic(err)
	}
}

func (f *DocFile) writeLine(line string) {
	f.write(line + "\n")
}

func (f *DocFile) writeLinef(format string, args ...interface{}) {
	f.writef(format+"\n", args...)
}

func (f *DocFile) writeLines(lines ...string) {
	for _, line := range lines {
		f.writeLine(line)
	}
}

func (f *DocFile) writef(format string, args ...interface{}) {
	f.write(fmt.Sprintf(format, args...))
}

func (ns Namespace) generateDocFile(name string, callback func(docfile *DocFile)) {
	filename := path.Join(ns.docsDir, name)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	callback(&DocFile{
		file: file,
		path: filename,
	})

	err = file.Close()
	if err != nil {
		panic(err)
	}
}
