// This is a generated file - DO NOT EDIT
// +build atk_1.6 atk_1.20 atk_2.8

package atk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// GetFocusObject is a wrapper around the C function atk_get_focus_object.
func GetFocusObject() *Object {
	retC := C.atk_get_focus_object()
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}
