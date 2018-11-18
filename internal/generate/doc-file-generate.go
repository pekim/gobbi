package generate

import (
	"fmt"
	"os"
	"path"
)

type FrontmatterParam struct {
	name  string
	value string
}

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

func (df *DocFile) writeFrontmatter(params ...FrontmatterParam) {
	df.writeLine("+++")
	for _, param := range params {
		df.writeLinef(`%s = "%s"`, param.name, param.value)
	}
	df.writeLine("+++")
}

// writeDocTextConditionally writes a line if there is Doc.Text.
//
// More specifically nothing will be written if Doc is nil, or if
// Doc.Text is an empty string.
//
// The format arg should have a single string (%s) placeholder.
func (df *DocFile) writeDocTextLine(format string, doc *Doc) {
	if doc == nil || doc.Text == "" {
		return
	}

	df.writeLinef(format, doc.Text)
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
