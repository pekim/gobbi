package generate

/*
	#cgo pkg-config: gobject-introspection-1.0
	#include <gobject-introspection-1.0/girepository.h>
*/
import "C"

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

type repository struct {
	name              string
	version           string
	goPackageName     string
	fullGoPackageName string
	libDir            string
	file              *jen.File

	cName      *C.char
	cVersion   *C.char
	repository *C.GIRepository
}

func ForRepository(name string, version string) {
	goPackageName := strings.ToLower(name)

	r := &repository{
		name:              name,
		version:           version,
		goPackageName:     goPackageName,
		fullGoPackageName: fmt.Sprintf("github.com/pekim/gobbi/lib/%s", goPackageName),
		libDir:            projectFilepath("..", "lib", goPackageName),

		cName:      C.CString(name),
		cVersion:   C.CString(version),
		repository: C.g_irepository_get_default(),
	}

	r.require()
	r.generate()
}

// require loads the repo.
func (r *repository) require() {
	var err *C.GError
	C.g_irepository_require(r.repository, r.cName, r.cVersion, 0, &err)

	if err != nil {
		message := C.GoString(err.message)
		panic(message)
	}
}

func (r *repository) generate() {
	r.generateLibDir()

	r.generateInfoType(C.GI_INFO_TYPE_CONSTANT, "constant", r.generateConstant)
	r.generateInfoType(C.GI_INFO_TYPE_ENUM, "enum", r.generateEnum)
}

// generateInfoType generates a file for an info type.
func (r *repository) generateInfoType(infoType C.GIInfoType, filename string, fn func(*C.GIBaseInfo)) {
	r.generateFile(filename, func(f *jen.File) {
		r.foreachInfo(infoType, fn)
	})
}

// foreachInfo calls a function for each info type of a specific type.
func (r *repository) foreachInfo(infoType C.GIInfoType, fn func(*C.GIBaseInfo)) {
	n := C.g_irepository_get_n_infos(r.repository, r.cName)
	for i := C.int(0); i < n; i++ {
		info := C.g_irepository_get_info(r.repository, r.cName, i)

		if C.g_base_info_get_type(info) == infoType {
			fn(info)
		}
	}
}
