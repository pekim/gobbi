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

// ParamSpecVariant is a wrapper around the C function g_param_spec_variant.
func ParamSpecVariant(name string, nick string, blurb string, type_ *glib.VariantType, defaultValue *glib.Variant, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_type := (*C.GVariantType)(C.NULL)
	if type_ != nil {
		c_type = (*C.GVariantType)(type_.ToC())
	}

	c_default_value := (*C.GVariant)(C.NULL)
	if defaultValue != nil {
		c_default_value = (*C.GVariant)(defaultValue.ToC())
	}

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_variant(c_name, c_nick, c_blurb, c_type, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}
