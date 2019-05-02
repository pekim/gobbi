// This is a generated file - DO NOT EDIT
// +build glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import (
	"C"
	"unsafe"
)

// Unsupported : g_clear_pointer : unsupported parameter pp : no type generator for gpointer (gpointer*) for param pp

// ComputeChecksumForBytes is a wrapper around the C function g_compute_checksum_for_bytes.
func ComputeChecksumForBytes(checksumType ChecksumType, data *Bytes) string {
	c_checksum_type := (C.GChecksumType)(checksumType)

	c_data := (*C.GBytes)(C.NULL)
	if data != nil {
		c_data = (*C.GBytes)(data.ToC())
	}

	retC := C.g_compute_checksum_for_bytes(c_checksum_type, c_data)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_datalist_id_dup_data : unsupported parameter dup_func : no type generator for DuplicateFunc (GDuplicateFunc) for param dup_func

// Unsupported : g_datalist_id_replace_data : unsupported parameter oldval : no type generator for gpointer (gpointer) for param oldval

// SpawnCheckExitStatus is a wrapper around the C function g_spawn_check_exit_status.
func SpawnCheckExitStatus(exitStatus int32) (bool, error) {
	c_exit_status := (C.gint)(exitStatus)

	var cThrowableError *C.GError

	retC := C.g_spawn_check_exit_status(c_exit_status, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_test_add_data_func_full : unsupported parameter test_data : no type generator for gpointer (gpointer) for param test_data

// TestExpectMessage is a wrapper around the C function g_test_expect_message.
func TestExpectMessage(logDomain string, logLevel LogLevelFlags, pattern string) {
	c_log_domain := C.CString(logDomain)
	defer C.free(unsafe.Pointer(c_log_domain))

	c_log_level := (C.GLogLevelFlags)(logLevel)

	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	C.g_test_expect_message(c_log_domain, c_log_level, c_pattern)

	return
}
