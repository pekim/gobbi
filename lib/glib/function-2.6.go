// +build glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// CheckVersion is a wrapper around the C function glib_check_version.
func CheckVersion(requiredMajor uint32, requiredMinor uint32, requiredMicro uint32) {}

// Unsupported : g_filename_display_basename : unsupported parameter filename : type filename, const gchar*

// Unsupported : g_filename_display_name : unsupported parameter filename : type filename, const gchar*

// GetFilenameCharsets is a wrapper around the C function g_get_filename_charsets.
func GetFilenameCharsets(charsets string) {}

// GetLanguageNames is a wrapper around the C function g_get_language_names.
func GetLanguageNames() {}

// GetSystemConfigDirs is a wrapper around the C function g_get_system_config_dirs.
func GetSystemConfigDirs() {}

// GetSystemDataDirs is a wrapper around the C function g_get_system_data_dirs.
func GetSystemDataDirs() {}

// GetUserCacheDir is a wrapper around the C function g_get_user_cache_dir.
func GetUserCacheDir() {}

// GetUserConfigDir is a wrapper around the C function g_get_user_config_dir.
func GetUserConfigDir() {}

// GetUserDataDir is a wrapper around the C function g_get_user_data_dir.
func GetUserDataDir() {}

// Unsupported : g_log_set_default_handler : unsupported parameter log_func : type LogFunc, GLogFunc

// Unsupported : g_rmdir : unsupported parameter filename : type filename, const gchar*

// StrvLength is a wrapper around the C function g_strv_length.
func StrvLength(strArray string) {}

// Unsupported : g_unlink : unsupported parameter filename : type filename, const gchar*

// UriListExtractUris is a wrapper around the C function g_uri_list_extract_uris.
func UriListExtractUris(uriList string) {}
