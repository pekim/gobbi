// This is a generated file - DO NOT EDIT
// +build glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import call "github.com/pekim/gobbi/lib/internal/call"

// Unsupported : g_get_environ : array return type :

// Unsupported : g_get_locale_variants : array return type :

// GetMonotonicTime is a wrapper around the C function g_get_monotonic_time.
func GetMonotonicTime() int64 {
	data := call.Data{Return: call.Return{Type: call.RT_LONG}}
	call.Function(1610, &data)
	ret := data.Return.Int64()

	return ret
}

// GetRealTime is a wrapper around the C function g_get_real_time.
func GetRealTime() int64 {
	data := call.Data{Return: call.Return{Type: call.RT_LONG}}
	call.Function(1614, &data)
	ret := data.Return.Int64()

	return ret
}

// Unsupported : g_get_user_runtime_dir : return type :

// Unsupported : g_test_add_data_func : unsupported parameter test_data : no type generator for gpointer (gconstpointer) for param test_data

// Unsupported : g_test_add_vtable : unsupported parameter test_data : no type generator for gpointer (gconstpointer) for param test_data

// Unsupported : g_test_create_case : unsupported parameter test_data : no type generator for gpointer (gconstpointer) for param test_data
