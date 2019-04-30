// This is a generated file - DO NOT EDIT
// +build gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// ParamSpecGType is a wrapper around the C record GParamSpecGType.
type ParamSpecGType struct {
	native *C.GParamSpecGType
	// parent_instance : record
	IsAType Type
}

func ParamSpecGTypeNewFromC(u unsafe.Pointer) *ParamSpecGType {
	c := (*C.GParamSpecGType)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecGType{
		IsAType: (Type)(c.is_a_type),
		native:  c,
	}

	return g
}

func (recv *ParamSpecGType) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
