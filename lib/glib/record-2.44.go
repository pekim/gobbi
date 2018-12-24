// This is a generated file - DO NOT EDIT
// +build glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

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
