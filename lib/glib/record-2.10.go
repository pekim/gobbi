// This is a generated file - DO NOT EDIT
// +build glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_async_queue_push_sorted : unsupported parameter func : no type generator for CompareDataFunc (GCompareDataFunc) for param func

// Unsupported : g_async_queue_push_sorted_unlocked : unsupported parameter func : no type generator for CompareDataFunc (GCompareDataFunc) for param func

// Unsupported : g_async_queue_sort : unsupported parameter func : no type generator for CompareDataFunc (GCompareDataFunc) for param func

// Unsupported : g_async_queue_sort_unlocked : unsupported parameter func : no type generator for CompareDataFunc (GCompareDataFunc) for param func

// Unsupported : g_date_set_time_t : unsupported parameter timet : no type generator for glong (time_t) for param timet

// SetTimeVal is a wrapper around the C function g_date_set_time_val.
func (recv *Date) SetTimeVal(timeval *TimeVal) {
	c_timeval := (*C.GTimeVal)(timeval.ToC())

	C.g_date_set_time_val((*C.GDate)(recv.native), c_timeval)

	return
}

// Unsupported : g_error_new : unsupported parameter ... : varargs

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// IsOwner is a wrapper around the C function g_main_context_is_owner.
func (recv *MainContext) IsOwner() bool {
	retC := C.g_main_context_is_owner((*C.GMainContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data_dnotify : no type generator for DestroyNotify (GDestroyNotify) for param user_data_dnotify

// Unsupported : g_option_group_new : unsupported parameter destroy : no type generator for DestroyNotify (GDestroyNotify) for param destroy

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Unsupported : g_thread_pool_set_sort_function : unsupported parameter func : no type generator for CompareDataFunc (GCompareDataFunc) for param func

// Unsupported : g_variant_builder_new : unsupported parameter type : Blacklisted record : GVariantType
