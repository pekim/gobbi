package generate

/*
	#cgo pkg-config: gobject-introspection-1.0
	#include <gobject-introspection-1.0/girepository.h>
	#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"os"
)

type repository struct {
	name    string
	version string
	format  bool

	cName      *C.char
	cVersion   *C.char
	repository *C.GIRepository
}

func ForRepository(name string, version string, format bool) {
	r := &repository{
		name:    name,
		version: version,
		format:  format,

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
		fmt.Printf("\nFailed to require %s %s, %s\n", r.name, r.version, message)
		os.Exit(1)
	}
}

func (r *repository) generate() {
	r.foreachInfo(C.GI_INFO_TYPE_CONSTANT, r.generateConstant)
}

func (r *repository) foreachInfo(infoType C.GIInfoType, fn func(*C.GIBaseInfo)) {
	n := C.g_irepository_get_n_infos(r.repository, r.cName)
	for i := C.int(0); i < n; i++ {
		info := C.g_irepository_get_info(r.repository, r.cName, i)

		if C.g_base_info_get_type(info) == infoType {
			fn(info)
		}
	}
}

func (r *repository) generateConstant(info *C.GIBaseInfo) {
	cName := C.g_base_info_get_name(info)
	name := C.GoString(cName)

	constant := (*C.GIConstantInfo)(info)
	var value C.GIArgument
	size := C.g_constant_info_get_value(constant, &value)

	fmt.Println(name, size)
}
