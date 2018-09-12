// +build glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_access : unsupported parameter filename : no param type for filename, const gchar*

// Unsupported : g_build_filenamev : unsupported parameter args : no param type

// Unsupported : g_build_pathv : unsupported parameter args : no param type

// Unsupported : g_chdir : unsupported parameter path : no param type for filename, const gchar*

// Unsupported : g_datalist_get_flags : unsupported parameter datalist : no param type for Data, GData**

// Unsupported : g_datalist_set_flags : unsupported parameter datalist : no param type for Data, GData**

// Unsupported : g_datalist_unset_flags : unsupported parameter datalist : no param type for Data, GData**

// Unsupported : g_file_set_contents : unsupported parameter filename : no param type for filename, const gchar*

// Unsupported : g_get_host_name : no return type

// Unsupported : g_listenv : no return type

// Unsupported : g_mkdir_with_parents : unsupported parameter pathname : no param type for filename, const gchar*

// TryMalloc0 is a wrapper around the C function g_try_malloc0.
func TryMalloc0(nBytes uint64) uintptr {
	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_try_malloc0(c_n_bytes)
	retGo :=
		(uintptr)(retC)

	return retGo
}

// Unsupported : g_utf8_collate_key_for_filename : no return type
