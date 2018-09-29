// This is a generated file - DO NOT EDIT
// +build atk_2.12

package atk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
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

// Copy is a wrapper around the C function atk_range_copy.
func (recv *Range) Copy() *Range {
	retC := C.atk_range_copy((*C.AtkRange)(recv.native))
	retGo := RangeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function atk_range_free.
func (recv *Range) Free() {
	C.atk_range_free((*C.AtkRange)(recv.native))

	return
}

// GetDescription is a wrapper around the C function atk_range_get_description.
func (recv *Range) GetDescription() string {
	retC := C.atk_range_get_description((*C.AtkRange)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetLowerLimit is a wrapper around the C function atk_range_get_lower_limit.
func (recv *Range) GetLowerLimit() float64 {
	retC := C.atk_range_get_lower_limit((*C.AtkRange)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetUpperLimit is a wrapper around the C function atk_range_get_upper_limit.
func (recv *Range) GetUpperLimit() float64 {
	retC := C.atk_range_get_upper_limit((*C.AtkRange)(recv.native))
	retGo := (float64)(retC)

	return retGo
}
