// This is a generated file - DO NOT EDIT
// +build glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_clear_pointer : unsupported parameter pp : no type generator for gpointer, gpointer*

// ComputeChecksumForBytes is a wrapper around the C function g_compute_checksum_for_bytes.
func ComputeChecksumForBytes(checksumType ChecksumType, data *Bytes) string {
	c_checksum_type := (C.GChecksumType)(checksumType)

	c_data := (*C.GBytes)(data.ToC())

	retC := C.g_compute_checksum_for_bytes(c_checksum_type, c_data)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_datalist_id_dup_data : unsupported parameter datalist : in string with indirection level of 2

// Unsupported : g_datalist_id_replace_data : unsupported parameter datalist : in string with indirection level of 2

// SpawnCheckExitStatus is a wrapper around the C function g_spawn_check_exit_status.
func SpawnCheckExitStatus(exitStatus int32) (bool, error) {
	c_exit_status := (C.gint)(exitStatus)

	var cThrowableError *C.GError

	retC := C.g_spawn_check_exit_status(c_exit_status, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_test_add_data_func_full : unsupported parameter test_func : no type generator for TestDataFunc, GTestDataFunc

// Unsupported : g_test_expect_message : no return generator
