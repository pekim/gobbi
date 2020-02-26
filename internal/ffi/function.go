package ffi

import (
	"fmt"
	"unsafe"
)

// #include <ffi.h>
// #include <stdlib.h>
// #include <stdio.h>
// #cgo pkg-config: libffi
/*
static void cc() {
  ffi_cif cif;
  ffi_type *args[1];
  void *values[1];
  char *s;
  ffi_arg rc;

	args[0] = &ffi_type_pointer;
	values[0] = &s;

	if (ffi_prep_cif(&cif, FFI_DEFAULT_ABI, 1,
		&ffi_type_sint, args) == FFI_OK)
	{
		s = "";
		ffi_call(&cif, printf, &rc, values);
	}
}
*/
import "C"

type Function struct {
	library    *Library
	name       string
	argTypes   []Type
	returnType Type
	fn         unsafe.Pointer
	cif        C.ffi_cif
}

func (f *Function) prepare() {
	argCount := len(f.argTypes)

	var cArgTypes **C.ffi_type
	if argCount > 0 {
		cBlock := C.malloc(C.ulong(argCount) * (C.ulong)(unsafe.Sizeof(uintptr(0))))
		defer C.free(unsafe.Pointer(cBlock))
		argTypesSlice := (*[1 << 30]*C.ffi_type)(unsafe.Pointer(cBlock))[:argCount:argCount]

		for i, _ := range f.argTypes {
			argTypesSlice[i] = (*C.ffi_type)(&f.argTypes[i])
		}
		cArgTypes = &argTypesSlice[0]
	}

	status := C.ffi_prep_cif(&f.cif, C.FFI_DEFAULT_ABI, C.uint(argCount), &C.ffi_type_void, cArgTypes)
	if status != C.FFI_OK {
		handleError(fmt.Errorf("Failed to prepare cif for %s.%s, status %d",
			f.library.name, f.name, status))
	}
}

func (f *Function) Call() {
	C.cc()

	//	var rc C.ffi_arg
	//
	//	argCount := len(f.argTypes)
	//	cBlock := C.malloc(C.ulong(argCount) * (C.ulong)(unsafe.Sizeof(uintptr(0))))
	//	defer C.free(unsafe.Pointer(cBlock))
	//	valuesSlice := (*[1 << 30]*C.void)(unsafe.Pointer(cBlock))[:argCount:argCount]
	//
	//	cString := C.CString("qwerty")
	//	valuesSlice[0] = (*C.void)(unsafe.Pointer(&cString))
	//	v1 := 6
	//	valuesSlice[1] = (*C.void)(unsafe.Pointer(&v1))
	//
	//	C.ffi_call(&f.cif, (*[0]byte)(f.fn), unsafe.Pointer(&rc), (*unsafe.Pointer)(unsafe.Pointer(&valuesSlice[0])))
}
