// This is a generated file - DO NOT EDIT
// +build glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import call "github.com/pekim/gobbi/lib/internal/call"

// Unsupported : g_double_equal : unsupported parameter v1 : no type generator for gpointer (gconstpointer) for param v1

// Unsupported : g_double_hash : unsupported parameter v : no type generator for gpointer (gconstpointer) for param v

// Unsupported : g_hostname_is_ascii_encoded : return type :

// Unsupported : g_hostname_is_ip_address : return type :

// Unsupported : g_hostname_is_non_ascii : return type :

// Unsupported : g_hostname_to_ascii : return type :

// Unsupported : g_hostname_to_unicode : return type :

// Unsupported : g_int64_equal : unsupported parameter v1 : no type generator for gpointer (gconstpointer) for param v1

// Unsupported : g_int64_hash : unsupported parameter v : no type generator for gpointer (gconstpointer) for param v

// Unsupported : g_main_context_get_thread_default : return type :

// ReloadUserSpecialDirsCache is a wrapper around the C function g_reload_user_special_dirs_cache.
func ReloadUserSpecialDirsCache() {
	data := call.Data{Return: call.Value{Type: call.TYPE_VOID}}
	call.Function(2611, &data)
	return
}

// Unsupported : g_test_log_set_fatal_handler : unsupported parameter log_func : no type generator for TestLogFatalFunc (GTestLogFatalFunc) for param log_func
