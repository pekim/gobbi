// This is a generated file - DO NOT EDIT
// +build atk_1.6 atk_1.9 atk_1.12 atk_1.13 atk_1.20 atk_1.30 atk_2.8 atk_2.12

package atk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// Gets the currently focused object.
/*

C function

atk_get_focus_object
*/
func GetFocusObject() *Object {
	retC := C.atk_get_focus_object()
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}
