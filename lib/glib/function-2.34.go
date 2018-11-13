// This is a generated file - DO NOT EDIT
// +build glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_clear_pointer : unsupported parameter destroy : no type generator for DestroyNotify (GDestroyNotify) for param destroy

// Computes the checksum for a binary @data. This is a
// convenience wrapper for g_checksum_new(), g_checksum_get_string()
// and g_checksum_free().
//
// The hexadecimal string returned will be in lower case.
/*

C function

g_compute_checksum_for_bytes
*/
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

// Unsupported : g_datalist_id_replace_data : unsupported parameter destroy : no type generator for DestroyNotify (GDestroyNotify) for param destroy

// Set @error if @exit_status indicates the child exited abnormally
// (e.g. with a nonzero exit code, or via a fatal signal).
//
// The g_spawn_sync() and g_child_watch_add() family of APIs return an
// exit status for subprocesses encoded in a platform-specific way.
// On Unix, this is guaranteed to be in the same format waitpid() returns,
// and on Windows it is guaranteed to be the result of GetExitCodeProcess().
//
// Prior to the introduction of this function in GLib 2.34, interpreting
// @exit_status required use of platform-specific APIs, which is problematic
// for software using GLib as a cross-platform layer.
//
// Additionally, many programs simply want to determine whether or not
// the child exited successfully, and either propagate a #GError or
// print a message to standard error. In that common case, this function
// can be used. Note that the error message in @error will contain
// human-readable information about the exit status.
//
// The @domain and @code of @error have special semantics in the case
// where the process has an "exit code", as opposed to being killed by
// a signal. On Unix, this happens if WIFEXITED() would be true of
// @exit_status. On Windows, it is always the case.
//
// The special semantics are that the actual exit code will be the
// code set in @error, and the domain will be %G_SPAWN_EXIT_ERROR.
// This allows you to differentiate between different exit codes.
//
// If the process was terminated by some means other than an exit
// status, the domain will be %G_SPAWN_ERROR, and the code will be
// %G_SPAWN_ERROR_FAILED.
//
// This function just offers convenience; you can of course also check
// the available platform via a macro such as %G_OS_UNIX, and use
// WIFEXITED() and WEXITSTATUS() on @exit_status directly. Do not attempt
// to scan or parse the error message string; it may be translated and/or
// change in future versions of GLib.
/*

C function

g_spawn_check_exit_status
*/
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

// Unsupported : g_test_add_data_func_full : unsupported parameter test_func : no type generator for TestDataFunc (GTestDataFunc) for param test_func

// Indicates that a message with the given @log_domain and @log_level,
// with text matching @pattern, is expected to be logged. When this
// message is logged, it will not be printed, and the test case will
// not abort.
//
// This API may only be used with the old logging API (g_log() without
// %G_LOG_USE_STRUCTURED defined). It will not work with the structured logging
// API. See [Testing for Messages][testing-for-messages].
//
// Use g_test_assert_expected_messages() to assert that all
// previously-expected messages have been seen and suppressed.
//
// You can call this multiple times in a row, if multiple messages are
// expected as a result of a single call. (The messages must appear in
// the same order as the calls to g_test_expect_message().)
//
// For example:
//
// |[<!-- language="C" -->
// g_main_context_push_thread_default() should fail if the
// context is already owned by another thread.
// g_test_expect_message (G_LOG_DOMAIN,
// G_LOG_LEVEL_CRITICAL,
// "assertion*acquired_context*failed");
// g_main_context_push_thread_default (bad_context);
// g_test_assert_expected_messages ();
// ]|
//
// Note that you cannot use this to test g_error() messages, since
// g_error() intentionally never returns even if the program doesn't
// abort; use g_test_trap_subprocess() in this case.
//
// If messages at %G_LOG_LEVEL_DEBUG are emitted, but not explicitly
// expected via g_test_expect_message() then they will be ignored.
/*

C function

g_test_expect_message
*/
func TestExpectMessage(logDomain string, logLevel LogLevelFlags, pattern string) {
	c_log_domain := C.CString(logDomain)
	defer C.free(unsafe.Pointer(c_log_domain))

	c_log_level := (C.GLogLevelFlags)(logLevel)

	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	C.g_test_expect_message(c_log_domain, c_log_level, c_pattern)

	return
}
