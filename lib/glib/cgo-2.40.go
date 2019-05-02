// This is a generated file - DO NOT EDIT
// +build glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

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

// Unsupported : g_str_tokenize_and_fold : unsupported parameter ascii_alternates : output array param ascii_alternates

// VariantDict is a wrapper around the C record GVariantDict.
type VariantDict struct {
	native *C.GVariantDict
}

func VariantDictNewFromC(u unsafe.Pointer) *VariantDict {
	c := (*C.GVariantDict)(u)
	if c == nil {
		return nil
	}

	g := &VariantDict{native: c}

	return g
}

func (recv *VariantDict) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VariantDictNew is a wrapper around the C function g_variant_dict_new.
func VariantDictNew(fromAsv *Variant) *VariantDict {
	c_from_asv := (*C.GVariant)(C.NULL)
	if fromAsv != nil {
		c_from_asv = (*C.GVariant)(fromAsv.ToC())
	}

	retC := C.g_variant_dict_new(c_from_asv)
	retGo := VariantDictNewFromC(unsafe.Pointer(retC))

	return retGo
}
