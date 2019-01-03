// This is a generated file - DO NOT EDIT
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

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

// GetUserSpecialDir is a wrapper around the C function g_get_user_special_dir.
func GetUserSpecialDir(directory UserDirectory) string {
	c_directory := (C.GUserDirectory)(directory)

	retC := C.g_get_user_special_dir(c_directory)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_once_init_enter : unsupported parameter location : no type generator for gpointer (void*) for param location

// Unsupported : g_once_init_leave : unsupported parameter location : no type generator for gpointer (void*) for param location

// Unsupported : g_regex_escape_string : unsupported parameter string :

// Unsupported : g_sequence_get : no return generator

// Unsupported : g_sequence_insert_before : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_set : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_slice_copy : unsupported parameter mem_block : no type generator for gpointer (gconstpointer) for param mem_block

// Unsupported : g_timeout_add_seconds : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_timeout_add_seconds_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// TimeoutSourceNewSeconds is a wrapper around the C function g_timeout_source_new_seconds.
func TimeoutSourceNewSeconds(interval uint32) *Source {
	c_interval := (C.guint)(interval)

	retC := C.g_timeout_source_new_seconds(c_interval)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// UnicharCombiningClass is a wrapper around the C function g_unichar_combining_class.
func UnicharCombiningClass(uc rune) int32 {
	c_uc := (C.gunichar)(uc)

	retC := C.g_unichar_combining_class(c_uc)
	retGo := (int32)(retC)

	return retGo
}

// UnicharGetScript is a wrapper around the C function g_unichar_get_script.
func UnicharGetScript(ch rune) UnicodeScript {
	c_ch := (C.gunichar)(ch)

	retC := C.g_unichar_get_script(c_ch)
	retGo := (UnicodeScript)(retC)

	return retGo
}

// UnicharIsmark is a wrapper around the C function g_unichar_ismark.
func UnicharIsmark(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_ismark(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIszerowidth is a wrapper around the C function g_unichar_iszerowidth.
func UnicharIszerowidth(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_iszerowidth(c_c)
	retGo := retC == C.TRUE

	return retGo
}
