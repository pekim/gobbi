// This is a generated file - DO NOT EDIT
// +build glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// ArrayGetElementSize is a wrapper around the C function g_array_get_element_size.
func ArrayGetElementSize(array []uintptr) uint32 {
	c_array := &array[0]

	retC := C.g_array_get_element_size((*C.GArray)(unsafe.Pointer(c_array)))
	retGo := (uint32)(retC)

	return retGo
}

// g_array_ref : array return type :
// ArrayUnref is a wrapper around the C function g_array_unref.
func ArrayUnref(array []uintptr) {
	c_array := &array[0]

	C.g_array_unref((*C.GArray)(unsafe.Pointer(c_array)))

	return
}

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// MainContextGetThreadDefault is a wrapper around the C function g_main_context_get_thread_default.
func MainContextGetThreadDefault() *MainContext {
	retC := C.g_main_context_get_thread_default()
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PopThreadDefault is a wrapper around the C function g_main_context_pop_thread_default.
func (recv *MainContext) PopThreadDefault() {
	C.g_main_context_pop_thread_default((*C.GMainContext)(recv.native))

	return
}

// PushThreadDefault is a wrapper around the C function g_main_context_push_thread_default.
func (recv *MainContext) PushThreadDefault() {
	C.g_main_context_push_thread_default((*C.GMainContext)(recv.native))

	return
}

// Ref is a wrapper around the C function g_mapped_file_ref.
func (recv *MappedFile) Ref() *MappedFile {
	retC := C.g_mapped_file_ref((*C.GMappedFile)(recv.native))
	retGo := MappedFileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ref is a wrapper around the C function g_tree_ref.
func (recv *Tree) Ref() *Tree {
	retC := C.g_tree_ref((*C.GTree)(recv.native))
	retGo := TreeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_tree_unref.
func (recv *Tree) Unref() {
	C.g_tree_unref((*C.GTree)(recv.native))

	return
}
