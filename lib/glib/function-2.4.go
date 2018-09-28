// This is a generated file - DO NOT EDIT
// +build glib_2.4 glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_atomic_int_add : unsupported parameter atomic : no type generator for gint, volatile gint*

// Unsupported : g_atomic_int_compare_and_exchange : unsupported parameter atomic : no type generator for gint, volatile gint*

// Unsupported : g_atomic_int_dec_and_test : unsupported parameter atomic : no type generator for gint, volatile gint*

// Unsupported : g_atomic_int_exchange_and_add : unsupported parameter atomic : no type generator for gint, volatile gint*

// Unsupported : g_atomic_int_get : unsupported parameter atomic : no type generator for gint, volatile const gint*

// Unsupported : g_atomic_int_inc : unsupported parameter atomic : no type generator for gint, volatile gint*

// Unsupported : g_atomic_int_set : unsupported parameter atomic : no type generator for gint, volatile gint*

// Unsupported : g_atomic_pointer_compare_and_exchange : unsupported parameter atomic : no type generator for gpointer, void*

// Unsupported : g_atomic_pointer_get : unsupported parameter atomic : no type generator for gpointer, void*

// Unsupported : g_atomic_pointer_set : unsupported parameter atomic : no type generator for gpointer, void*

// Unsupported : g_child_watch_add : unsupported parameter function : no type generator for ChildWatchFunc, GChildWatchFunc

// Unsupported : g_child_watch_add_full : unsupported parameter function : no type generator for ChildWatchFunc, GChildWatchFunc

// ChildWatchSourceNew is a wrapper around the C function g_child_watch_source_new.
func ChildWatchSourceNew(pid Pid) *Source {
	c_pid := (C.GPid)(pid)

	retC := C.g_child_watch_source_new(c_pid)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FileReadLink is a wrapper around the C function g_file_read_link.
func FileReadLink(filename string) (string, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.g_file_read_link(c_filename, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_markup_printf_escaped : unsupported parameter ... : varargs

// Unsupported : g_markup_vprintf_escaped : unsupported parameter args : no type generator for va_list, va_list

// Setenv is a wrapper around the C function g_setenv.
func Setenv(variable string, value string, overwrite bool) bool {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	c_overwrite :=
		boolToGboolean(overwrite)

	retC := C.g_setenv(c_variable, c_value, c_overwrite)
	retGo := retC == C.TRUE

	return retGo
}

// StripContext is a wrapper around the C function g_strip_context.
func StripContext(msgid string, msgval string) string {
	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	c_msgval := C.CString(msgval)
	defer C.free(unsafe.Pointer(c_msgval))

	retC := C.g_strip_context(c_msgid, c_msgval)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_strsplit_set : no return type

// Unsupported : g_unichar_get_mirror_char : unsupported parameter mirrored_ch : no type generator for gunichar, gunichar*

// Unsetenv is a wrapper around the C function g_unsetenv.
func Unsetenv(variable string) {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	C.g_unsetenv(c_variable)

	return
}

// Unsupported : g_vasprintf : unsupported parameter string : in string with indirection level of 2
