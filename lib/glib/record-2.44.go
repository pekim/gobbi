// This is a generated file - DO NOT EDIT
// +build glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_bytes_new_with_free_func : unsupported parameter free_func : no type generator for DestroyNotify (GDestroyNotify) for param free_func

// Unsupported : g_error_new : unsupported parameter ... : varargs

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data_dnotify : no type generator for DestroyNotify (GDestroyNotify) for param user_data_dnotify

// GetStrictPosix is a wrapper around the C function g_option_context_get_strict_posix.
func (recv *OptionContext) GetStrictPosix() bool {
	retC := C.g_option_context_get_strict_posix((*C.GOptionContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetStrictPosix is a wrapper around the C function g_option_context_set_strict_posix.
func (recv *OptionContext) SetStrictPosix(strictPosix bool) {
	c_strict_posix :=
		boolToGboolean(strictPosix)

	C.g_option_context_set_strict_posix((*C.GOptionContext)(recv.native), c_strict_posix)

	return
}

// Unsupported : g_option_group_new : unsupported parameter destroy : no type generator for DestroyNotify (GDestroyNotify) for param destroy

// Ref is a wrapper around the C function g_option_group_ref.
func (recv *OptionGroup) Ref() *OptionGroup {
	retC := C.g_option_group_ref((*C.GOptionGroup)(recv.native))
	retGo := OptionGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_option_group_unref.
func (recv *OptionGroup) Unref() {
	C.g_option_group_unref((*C.GOptionGroup)(recv.native))

	return
}

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Unsupported : g_variant_builder_new : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_variant_dict_new : unsupported parameter from_asv : Blacklisted record : GVariant
