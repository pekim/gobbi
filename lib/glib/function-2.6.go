// +build glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #include <glib.h>
// #include <stdlib.h>
import "C"

// CheckVersion is a wrapper around the C function glib_check_version.
func CheckVersion(requiredMajor uint32, requiredMinor uint32, requiredMicro uint32) {
	c_required_major := (C.guint)(requiredMajor)

	c_required_minor := (C.guint)(requiredMinor)

	c_required_micro := (C.guint)(requiredMicro)

	C.glib_check_version(c_required_major, c_required_minor, c_required_micro)
}

// Unsupported : g_filename_display_basename : unsupported parameter filename : type filename, const gchar*

// Unsupported : g_filename_display_name : unsupported parameter filename : type filename, const gchar*

// GetFilenameCharsets is a wrapper around the C function g_get_filename_charsets.
func GetFilenameCharsets(charsets string) {
	c_charsets := C.CString(charsets)
	defer C.free(unsafe.Pointer(c_charsets))

	C.g_get_filename_charsets(c_charsets)
}

// GetLanguageNames is a wrapper around the C function g_get_language_names.
func GetLanguageNames() {
	C.g_get_language_names()
}

// GetSystemConfigDirs is a wrapper around the C function g_get_system_config_dirs.
func GetSystemConfigDirs() {
	C.g_get_system_config_dirs()
}

// GetSystemDataDirs is a wrapper around the C function g_get_system_data_dirs.
func GetSystemDataDirs() {
	C.g_get_system_data_dirs()
}

// GetUserCacheDir is a wrapper around the C function g_get_user_cache_dir.
func GetUserCacheDir() {
	C.g_get_user_cache_dir()
}

// GetUserConfigDir is a wrapper around the C function g_get_user_config_dir.
func GetUserConfigDir() {
	C.g_get_user_config_dir()
}

// GetUserDataDir is a wrapper around the C function g_get_user_data_dir.
func GetUserDataDir() {
	C.g_get_user_data_dir()
}

// Unsupported : g_log_set_default_handler : unsupported parameter log_func : type LogFunc, GLogFunc

// Unsupported : g_rmdir : unsupported parameter filename : type filename, const gchar*

// StrvLength is a wrapper around the C function g_strv_length.
func StrvLength(strArray string) {
	c_str_array := C.CString(strArray)
	defer C.free(unsafe.Pointer(c_str_array))

	C.g_strv_length(c_str_array)
}

// Unsupported : g_unlink : unsupported parameter filename : type filename, const gchar*

// UriListExtractUris is a wrapper around the C function g_uri_list_extract_uris.
func UriListExtractUris(uriList string) {
	c_uri_list := C.CString(uriList)
	defer C.free(unsafe.Pointer(c_uri_list))

	C.g_uri_list_extract_uris(c_uri_list)
}
