// This is a generated file - DO NOT EDIT
// +build glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Close is a wrapper around the C function g_close.
func Close(fd int32) (bool, error) {
	c_fd := (C.gint)(fd)

	var cThrowableError *C.GError

	retC := C.g_close(c_fd, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetNumProcessors is a wrapper around the C function g_get_num_processors.
func GetNumProcessors() uint32 {
	retC := C.g_get_num_processors()
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_unix_fd_add : unsupported parameter function : no type generator for UnixFDSourceFunc (GUnixFDSourceFunc) for param function

// Unsupported : g_unix_fd_add_full : unsupported parameter function : no type generator for UnixFDSourceFunc (GUnixFDSourceFunc) for param function

// UnixFdSourceNew is a wrapper around the C function g_unix_fd_source_new.
func UnixFdSourceNew(fd int32, condition IOCondition) *Source {
	c_fd := (C.gint)(fd)

	c_condition := (C.GIOCondition)(condition)

	retC := C.g_unix_fd_source_new(c_fd, c_condition)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}
