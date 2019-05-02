// This is a generated file - DO NOT EDIT
// +build atk_2.12

package atk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// RangeNew is a wrapper around the C function atk_range_new.
func RangeNew(lowerLimit float64, upperLimit float64, description string) *Range {
	c_lower_limit := (C.gdouble)(lowerLimit)

	c_upper_limit := (C.gdouble)(upperLimit)

	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	retC := C.atk_range_new(c_lower_limit, c_upper_limit, c_description)
	retGo := RangeNewFromC(unsafe.Pointer(retC))

	return retGo
}
