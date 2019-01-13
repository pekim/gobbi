// This is a generated file - DO NOT EDIT
// +build glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

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

// AsciiStrtoll is a wrapper around the C function g_ascii_strtoll.
func AsciiStrtoll(nptr string, base uint32) (int64, string) {
	c_nptr := C.CString(nptr)
	defer C.free(unsafe.Pointer(c_nptr))

	var c_endptr *C.gchar

	c_base := (C.guint)(base)

	retC := C.g_ascii_strtoll(c_nptr, &c_endptr, c_base)
	retGo := (int64)(retC)

	endptr := C.GoString(c_endptr)

	return retGo, endptr
}

// Unsupported : g_base64_decode : array return type :

// Unsupported : g_base64_decode_step : unsupported parameter out : output array param out

// Base64Encode is a wrapper around the C function g_base64_encode.
func Base64Encode(data []uint8) string {
	c_data_array := make([]C.guint8, len(data), len(data))
	c_data := &c_data_array[0]

	c_len := (C.gsize)(len(data))

	retC := C.g_base64_encode(c_data, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_base64_encode_close : unsupported parameter out : output array param out

// Unsupported : g_base64_encode_step : unsupported parameter out : output array param out

// MainCurrentSource is a wrapper around the C function g_main_current_source.
func MainCurrentSource() *Source {
	retC := C.g_main_current_source()
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// UnicharIswideCjk is a wrapper around the C function g_unichar_iswide_cjk.
func UnicharIswideCjk(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_iswide_cjk(c_c)
	retGo := retC == C.TRUE

	return retGo
}
