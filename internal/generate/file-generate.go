package generate

import "C"
import (
	"os"
)

func (r *repository) generateLibDir() {
	err := os.MkdirAll(r.libDir, 0775)
	if err != nil {
		panic(err)
	}
}

//func (r *repository) generateFile(name string, generateContent func(f *jen.File)) {
//	r.file = jen.NewFile(r.goPackageName)
//
//	// Use a standard generated file comment format.
//	// https://github.com/golang/go/issues/13560#issuecomment-288457920
//	r.file.HeaderComment(("Code generated - DO NOT EDIT."))
//	r.file.Line()
//
//	generateContent(r.file)
//
//	filepath := path.Join(r.libDir, fmt.Sprintf("%s.go", name))
//	err := r.file.Save(filepath)
//	if err != nil {
//		panic(err)
//	}
//}
