// This is a generated file - DO NOT EDIT
// +build glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_bytes_new : unsupported parameter data : no type generator for guint8 () for array param data

// Unsupported : g_bytes_new_static : unsupported parameter data : no type generator for guint8 () for array param data

// Unsupported : g_bytes_new_take : unsupported parameter data : no type generator for guint8 () for array param data

// Unsupported : g_bytes_new_with_free_func : unsupported parameter data : no type generator for guint8 () for array param data

// Unsupported : g_error_new : unsupported parameter ... : varargs

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data_dnotify : no type generator for DestroyNotify (GDestroyNotify) for param user_data_dnotify

// Unsupported : g_option_group_new : unsupported parameter destroy : no type generator for DestroyNotify (GDestroyNotify) for param destroy

// GetMaxLookbehind is a wrapper around the C function g_regex_get_max_lookbehind.
func (recv *Regex) GetMaxLookbehind() int32 {
	retC := C.g_regex_get_max_lookbehind((*C.GRegex)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Unsupported : g_variant_builder_new : unsupported parameter type : Blacklisted record : GVariantType
