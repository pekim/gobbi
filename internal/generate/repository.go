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
	"unsafe"
)

type repository struct {
	name       string
	version    string
	format     bool
	repository *C.GIRepository
}

func ForRepository(name string, version string, format bool) {
	repo:= &repository{
		name:       name,
		version:    version,
		format:     format,
		repository: C.g_irepository_get_default(),
	}

	repo.require()
}

func (r*repository) require() {
	cName:=C.CString(r.name)
	defer C.free(unsafe.Pointer(cName))

	cVersion:=C.CString(r.version)
	defer C.free(unsafe.Pointer(cVersion))

	var error *C.GError
	C.g_irepository_require (r.repository, cName, cVersion, 0, &error);

	if error!=nil {
		message:=C.GoString(error.message)
		fmt.Printf("\nFailed to require %s %s, %s\n",r.name, r.version,message)
		os.Exit(1)
	}
}
