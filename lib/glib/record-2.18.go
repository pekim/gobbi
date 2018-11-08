// This is a generated file - DO NOT EDIT
// +build glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Reset is a wrapper around the C function g_checksum_reset.
func (recv *Checksum) Reset() {
	C.g_checksum_reset((*C.GChecksum)(recv.native))

	return
}

// Unsupported : g_error_new : unsupported parameter ... : varargs

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data_dnotify : no type generator for DestroyNotify (GDestroyNotify) for param user_data_dnotify

// GetUserData is a wrapper around the C function g_markup_parse_context_get_user_data.
func (recv *MarkupParseContext) GetUserData() uintptr {
	retC := C.g_markup_parse_context_get_user_data((*C.GMarkupParseContext)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Pop is a wrapper around the C function g_markup_parse_context_pop.
func (recv *MarkupParseContext) Pop() uintptr {
	retC := C.g_markup_parse_context_pop((*C.GMarkupParseContext)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Push is a wrapper around the C function g_markup_parse_context_push.
func (recv *MarkupParseContext) Push(parser *MarkupParser, userData uintptr) {
	c_parser := (*C.GMarkupParser)(parser.ToC())

	c_user_data := (C.gpointer)(userData)

	C.g_markup_parse_context_push((*C.GMarkupParseContext)(recv.native), c_parser, c_user_data)

	return
}

// Unsupported : g_option_group_new : unsupported parameter destroy : no type generator for DestroyNotify (GDestroyNotify) for param destroy

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Unsupported : g_variant_builder_new : unsupported parameter type : Blacklisted record : GVariantType
