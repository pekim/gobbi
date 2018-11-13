// This is a generated file - DO NOT EDIT
// +build glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// This wraps the close() call; in case of error, %errno will be
// preserved, but the error will also be stored as a #GError in @error.
//
// Besides using #GError, there is another major reason to prefer this
// function over the call provided by the system; on Unix, it will
// attempt to correctly handle %EINTR, which has platform-specific
// semantics.
/*

C function : g_close
*/
func Close(fd int32) (bool, error) {
	c_fd := (C.gint)(fd)

	var cThrowableError *C.GError

	retC := C.g_close(c_fd, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Determine the approximate number of threads that the system will
// schedule simultaneously for this process.  This is intended to be
// used as a parameter to g_thread_pool_new() for CPU bound tasks and
// similar cases.
/*

C function : g_get_num_processors
*/
func GetNumProcessors() uint32 {
	retC := C.g_get_num_processors()
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_unix_fd_add : unsupported parameter function : no type generator for UnixFDSourceFunc (GUnixFDSourceFunc) for param function

// Unsupported : g_unix_fd_add_full : unsupported parameter function : no type generator for UnixFDSourceFunc (GUnixFDSourceFunc) for param function

// Creates a #GSource to watch for a particular IO condition on a file
// descriptor.
//
// The source will never close the fd -- you must do it yourself.
/*

C function : g_unix_fd_source_new
*/
func UnixFdSourceNew(fd int32, condition IOCondition) *Source {
	c_fd := (C.gint)(fd)

	c_condition := (C.GIOCondition)(condition)

	retC := C.g_unix_fd_source_new(c_fd, c_condition)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}
