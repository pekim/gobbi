package ffi

/*
#include <stdlib.h>

#include <dlfcn.h>
#cgo LDFLAGS: -ldl

*/
import "C"

import (
	"fmt"
	"unsafe"
)

type Library struct {
	name   string
	handle unsafe.Pointer
}

func OpenLibrary(name string) *Library {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	handle := C.dlopen(cName, C.RTLD_LAZY)
	if handle == nil {
		dlError := C.GoString(C.dlerror())
		panic(fmt.Sprintf("Failed to load library : %s", dlError))
	}

	return &Library{
		name:   name,
		handle: handle,
	}
}

func (l *Library) function(name string) (*Function, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.dlerror()
	fn := C.dlsym(l.handle, cName)
	if fn == nil {
		dlError := C.GoString(C.dlerror())
		return nil, fmt.Errorf("Failed to find function : %s", dlError)
	}

	return &Function{
		library: l,
		name:    name,
		fn:      fn,
	}, nil
}
