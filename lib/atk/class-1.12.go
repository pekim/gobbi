// This is a generated file - DO NOT EDIT
// +build atk_1.12 atk_1.13 atk_1.20 atk_1.30 atk_2.8 atk_2.12

package atk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// GetAttributes is a wrapper around the C function atk_object_get_attributes.
func (recv *Object) GetAttributes() *AttributeSet {
	retC := C.atk_object_get_attributes((*C.AtkObject)(recv.native))
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}
