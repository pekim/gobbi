// +build glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported function: g_clear_pointer : unsupported parameter pp : type gpointer, gpointer*

// Unsupported function: g_compute_checksum_for_bytes : unsupported parameter checksum_type : type ChecksumType, GChecksumType

// Unsupported function: g_datalist_id_dup_data : unsupported parameter datalist : type Data, GData**

// Unsupported function: g_datalist_id_replace_data : unsupported parameter datalist : type Data, GData**

// SpawnCheckExitStatus is a wrapper around the C function g_spawn_check_exit_status.
func SpawnCheckExitStatus(exitStatus int32) {}

// Unsupported function: g_test_add_data_func_full : unsupported parameter test_func : type TestDataFunc, GTestDataFunc

// Unsupported function: g_test_expect_message : unsupported parameter log_level : type LogLevelFlags, GLogLevelFlags
