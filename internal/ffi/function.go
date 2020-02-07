package ffi

// #include <ffi.h>
// #cgo pkg-config: libffi
import "C"

import (
	"fmt"
	"unsafe"
)

type Function struct {
	library *Library
	name    string
	fn      unsafe.Pointer
	cif     C.ffi_cif
}

func (f *Function) prepare() {
	status := C.ffi_prep_cif(&f.cif, C.FFI_DEFAULT_ABI, 0, &C.ffi_type_void, nil)
	if status != C.FFI_OK {
		handleError(fmt.Errorf("Failed to prepare cif for %s.%s, status %d",
			f.library.name, f.name, status))
	}
}

func (f *Function) Call() {
}
