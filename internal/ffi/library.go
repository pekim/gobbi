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
	name     string
	filename string
	handle   unsafe.Pointer
}

func OpenLibrary(name string, filename string) *Library {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	handle := C.dlopen(cFilename, C.RTLD_LAZY)
	if handle == nil {
		dlError := C.GoString(C.dlerror())
		panic(fmt.Sprintf("Failed to load library : %s", dlError))
	}

	return &Library{
		name:     name,
		filename: filename,
		handle:   handle,
	}
}

func (l *Library) function(name string, argTypes []Type, returnType Type) *Function {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.dlerror()
	fn := C.dlsym(l.handle, cName)
	if fn == nil {
		dlError := C.GoString(C.dlerror())
		err := fmt.Errorf("Failed to find function : %s", dlError)
		handleError(err)
		return nil
	}

	f := &Function{
		library:    l,
		name:       name,
		argTypes:   argTypes,
		returnType: returnType,
		fn:         fn,
	}
	f.prepare()

	return f
}
