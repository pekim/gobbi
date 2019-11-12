package generate

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type repository struct {
	// xml mappings
	XMLName xml.Name `xml:"repository"`
	Version string   `xml:"Version,attr"`
	//Includes  []Include  `xml:"http://www.gtk.org/introspection/core/1.0 include"`
	Namespace *Namespace `xml:"namespace"`

	addenda *repository
}

func (r *repository) loadFromFile(filename string, required bool) {
	filepath := projectFilepath("gir-files", filename)
	source, err := ioutil.ReadFile(filepath)
	if err != nil {
		if os.IsNotExist(err) && !required {
			return
		}

		panic(err)
	}

	err = xml.Unmarshal(source, r)
	if err != nil {
		panic(fmt.Errorf("Failed to parse %s : %s", filepath, err))
	}
}

func repositoryFromFile(spec RepositorySpec) *repository {
	fullname := spec.Name + "-" + spec.Version

	r := &repository{
		addenda: &repository{},
	}
	r.loadFromFile(fullname+".gir", true)
	r.addenda.loadFromFile(fullname+"-addenda.gir", false)

	//fmt.Println(r.Namespace.Name, r.Namespace.Version)
	return r
}

///*
//	#cgo pkg-config: gobject-introspection-1.0
//	#include <gobject-introspection-1.0/girepository.h>
//*/
//import "C"
//
//import (
//	"fmt"
//	"github.com/dave/jennifer/jen"
//	"strings"
//)
//
//type repository struct {
//	name              string
//	Version           string
//	goPackageName     string
//	fullGoPackageName string
//	libDir            string
//	file              *jen.File
//
//	cName      *C.char
//	cVersion   *C.char
//	repository *C.GIRepository
//}
//
//func ForRepository(name string, Version string) {
//	goPackageName := strings.ToLower(name)
//
//	r := &repository{
//		name:              name,
//		Version:           Version,
//		goPackageName:     goPackageName,
//		fullGoPackageName: fmt.Sprintf("github.com/pekim/gobbi/lib/%s", goPackageName),
//		libDir:            projectFilepath("..", "lib", goPackageName),
//
//		cName:      C.CString(name),
//		cVersion:   C.CString(Version),
//		repository: C.g_irepository_get_default(),
//	}
//
//	r.require()
//	r.generate()
//}
//
//// require loads the repo.
//func (r *repository) require() {
//	var err *C.GError
//	C.g_irepository_require(r.repository, r.cName, r.cVersion, 0, &err)
//
//	if err != nil {
//		message := C.GoString(err.message)
//		panic(message)
//	}
//}
//
//func (r *repository) generate() {
//	r.generateLibDir()
//
//	r.generateInfoType(C.GI_INFO_TYPE_CONSTANT, "constant", r.generateConstant)
//	r.generateInfoType(C.GI_INFO_TYPE_ENUM, "enum", r.generateEnum)
//	r.generateInfoType(C.GI_INFO_TYPE_FLAGS, "flags", r.generateEnum)
//}
//
//// generateInfoType generates a file for an info type.
//func (r *repository) generateInfoType(infoType C.GIInfoType, filename string, fn func(*C.GIBaseInfo)) {
//	r.generateFile(filename, func(f *jen.File) {
//		r.foreachInfo(infoType, fn)
//	})
//}
//
//// foreachInfo calls a function for each info type of a specific type.
//func (r *repository) foreachInfo(infoType C.GIInfoType, fn func(*C.GIBaseInfo)) {
//	n := C.g_irepository_get_n_infos(r.repository, r.cName)
//	for i := C.int(0); i < n; i++ {
//		info := C.g_irepository_get_info(r.repository, r.cName, i)
//
//		if C.g_base_info_get_type(info) == infoType {
//			fn(info)
//		}
//	}
//}
