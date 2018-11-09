// This is a generated file - DO NOT EDIT
// +build glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// GetMaxLookbehind is a wrapper around the C function g_regex_get_max_lookbehind.
func (recv *Regex) GetMaxLookbehind() int32 {
	retC := C.g_regex_get_max_lookbehind((*C.GRegex)(recv.native))
	retGo := (int32)(retC)

	return retGo
}
