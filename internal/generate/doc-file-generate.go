package generate

import (
	"fmt"
	"os"
	"path"
)

type DocFile struct {
	file *os.File
}

func (df *DocFile) write(text string) {
	_, err := df.file.WriteString(text)
	if err != nil {
		panic(err)
	}
}

func (df *DocFile) writef(format string, args ...interface{}) {
	df.write(fmt.Sprintf(format, args...))
}

func (df *DocFile) writeLinef(format string, args ...interface{}) {
	df.writef(format+"\n", args...)
}

func (df *DocFile) writeLine(line string) {
	df.writef(line + "\n")
}

func generateDocDir(filepath string) {
	err := os.MkdirAll(path.Dir(filepath), 0775)
	if err != nil {
		panic(err)
	}
}

func generateDocFile(fn func(*DocFile), paths ...string) {
	filepath := projectFilepath(append([]string{"docs-src"}, paths...)...)
	generateDocDir(filepath)

	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}

	df := &DocFile{
		file: file,
	}

	fn(df)

	df.file.Close()
	if err != nil {
		panic(err)
	}
}
