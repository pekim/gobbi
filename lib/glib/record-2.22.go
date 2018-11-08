// This is a generated file - DO NOT EDIT
// +build glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_error_new : unsupported parameter ... : varargs

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

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

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data_dnotify : no type generator for DestroyNotify (GDestroyNotify) for param user_data_dnotify

// Unsupported : g_option_group_new : unsupported parameter destroy : no type generator for DestroyNotify (GDestroyNotify) for param destroy

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

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

// Unsupported : g_variant_builder_new : unsupported parameter type : Blacklisted record : GVariantType
