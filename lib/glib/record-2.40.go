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

// g_hash_table_get_keys_as_array : no type generator for gpointer (gpointer) for array return
// Blacklisted : g_key_file_save_to_file

// Blacklisted : g_option_context_parse_strv

// Blacklisted : g_variant_parse_error_print_context

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

// Equals compares this VariantDict with another VariantDict, and returns true if they represent the same GObject.
func (recv *VariantDict) Equals(other *VariantDict) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_variant_dict_new

// Blacklisted : g_variant_dict_clear

// Blacklisted : g_variant_dict_contains

// Blacklisted : g_variant_dict_end

// Init is a wrapper around the C function g_variant_dict_init.
func (recv *VariantDict) Init(fromAsv *Variant) {
	c_from_asv := (*C.GVariant)(C.NULL)
	if fromAsv != nil {
		c_from_asv = (*C.GVariant)(fromAsv.ToC())
	}

	C.g_variant_dict_init((*C.GVariantDict)(recv.native), c_from_asv)

	return
}

// Blacklisted : g_variant_dict_insert

// Blacklisted : g_variant_dict_insert_value

// Blacklisted : g_variant_dict_lookup

// Blacklisted : g_variant_dict_lookup_value

// Blacklisted : g_variant_dict_ref

// Blacklisted : g_variant_dict_remove

// Blacklisted : g_variant_dict_unref
