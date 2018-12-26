// This is a generated file - DO NOT EDIT
// +build gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_object_class_install_properties : unsupported parameter pspecs :

// DupVariant is a wrapper around the C function g_value_dup_variant.
func (recv *Value) DupVariant() *glib.Variant {
	retC := C.g_value_dup_variant((*C.GValue)(recv.native))
	var retGo (*glib.Variant)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.VariantNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetVariant is a wrapper around the C function g_value_get_variant.
func (recv *Value) GetVariant() *glib.Variant {
	retC := C.g_value_get_variant((*C.GValue)(recv.native))
	var retGo (*glib.Variant)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.VariantNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SetVariant is a wrapper around the C function g_value_set_variant.
func (recv *Value) SetVariant(variant *glib.Variant) {
	c_variant := (*C.GVariant)(C.NULL)
	if variant != nil {
		c_variant = (*C.GVariant)(variant.ToC())
	}

	C.g_value_set_variant((*C.GValue)(recv.native), c_variant)

	return
}

// TakeVariant is a wrapper around the C function g_value_take_variant.
func (recv *Value) TakeVariant(variant *glib.Variant) {
	c_variant := (*C.GVariant)(C.NULL)
	if variant != nil {
		c_variant = (*C.GVariant)(variant.ToC())
	}

	C.g_value_take_variant((*C.GValue)(recv.native), c_variant)

	return
}
