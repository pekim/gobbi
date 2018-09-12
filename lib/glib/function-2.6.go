// +build glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// CheckVersion is a wrapper around the C function glib_check_version.
func CheckVersion(requiredMajor uint32, requiredMinor uint32, requiredMicro uint32) string {
	c_required_major := (C.guint)(requiredMajor)

	c_required_minor := (C.guint)(requiredMinor)

	c_required_micro := (C.guint)(requiredMicro)

	retC := C.glib_check_version(c_required_major, c_required_minor, c_required_micro)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Unsupported : g_filename_display_basename : unsupported parameter filename : no param type for filename, const gchar*

// Unsupported : g_filename_display_name : unsupported parameter filename : no param type for filename, const gchar*

// Unsupported : g_get_filename_charsets : unsupported parameter charsets : in for string with indirection level of 3

// Unsupported : g_get_language_names : no return type

// Unsupported : g_get_system_config_dirs : no return type

// Unsupported : g_get_system_data_dirs : no return type

// Unsupported : g_get_user_cache_dir : no return type

// Unsupported : g_get_user_config_dir : no return type

// Unsupported : g_get_user_data_dir : no return type

// Unsupported : g_log_set_default_handler : unsupported parameter log_func : no param type for LogFunc, GLogFunc

// Unsupported : g_rmdir : unsupported parameter filename : no param type for filename, const gchar*

// Unsupported : g_strv_length : unsupported parameter str_array : in for string with indirection level of 2

// Unsupported : g_unlink : unsupported parameter filename : no param type for filename, const gchar*

// Unsupported : g_uri_list_extract_uris : no return type
