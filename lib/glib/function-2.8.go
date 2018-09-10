// +build glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_access : unsupported parameter filename : type filename, const gchar*

// Unsupported : g_build_filenamev : unsupported parameter args : no type

// Unsupported : g_build_pathv : unsupported parameter args : no type

// Unsupported : g_chdir : unsupported parameter path : type filename, const gchar*

// Unsupported : g_datalist_get_flags : unsupported parameter datalist : type Data, GData**

// Unsupported : g_datalist_set_flags : unsupported parameter datalist : type Data, GData**

// Unsupported : g_datalist_unset_flags : unsupported parameter datalist : type Data, GData**

// Unsupported : g_file_set_contents : unsupported parameter filename : type filename, const gchar*

// GetHostName is a wrapper around the C function g_get_host_name.
func GetHostName() {}

// Listenv is a wrapper around the C function g_listenv.
func Listenv() {}

// Unsupported : g_mkdir_with_parents : unsupported parameter pathname : type filename, const gchar*

// TryMalloc0 is a wrapper around the C function g_try_malloc0.
func TryMalloc0(nBytes uint64) {}

// Utf8CollateKeyForFilename is a wrapper around the C function g_utf8_collate_key_for_filename.
func Utf8CollateKeyForFilename(str string, len int64) {}
