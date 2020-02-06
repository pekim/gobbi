package ffi

// #include <dlfcn.h>
// #include <stdlib.h>
// #cgo LDFLAGS: -ldl
import "C"

import (
	"fmt"
	"unsafe"
)

type Library struct {
	name   string
	handle unsafe.Pointer
}

func NewLibrary(name string) *Library {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	handle := C.dlopen(cName, C.RTLD_NOW)
	if handle == nil {
		error := C.GoString(C.dlerror())
		panic(fmt.Sprintf("Failed to load library : %s", error))
	}

	return &Library{
		name:   name,
		handle: handle,
	}
}
