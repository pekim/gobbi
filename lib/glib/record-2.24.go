// This is a generated file - DO NOT EDIT
// +build glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

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

// Variant is a wrapper around the C record GVariant.
type Variant struct {
	native *C.GVariant
}

func VariantNewFromC(u unsafe.Pointer) *Variant {
	c := (*C.GVariant)(u)
	if c == nil {
		return nil
	}

	g := &Variant{native: c}

	return g
}

func (recv *Variant) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Variant with another Variant, and returns true if they represent the same GObject.
func (recv *Variant) Equals(other *Variant) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_variant_new

// Unsupported : g_variant_new_array : unsupported parameter children :

// Blacklisted : g_variant_new_boolean

// Blacklisted : g_variant_new_byte

// Blacklisted : g_variant_new_dict_entry

// Blacklisted : g_variant_new_double

// Unsupported : g_variant_new_from_data : unsupported parameter notify : no type generator for DestroyNotify (GDestroyNotify) for param notify

// Blacklisted : g_variant_new_handle

// Blacklisted : g_variant_new_int16

// Blacklisted : g_variant_new_int32

// Blacklisted : g_variant_new_int64

// Blacklisted : g_variant_new_maybe

// Blacklisted : g_variant_new_object_path

// Blacklisted : g_variant_new_parsed

// Unsupported : g_variant_new_parsed_va : unsupported parameter app : no type generator for va_list (va_list*) for param app

// Blacklisted : g_variant_new_signature

// Blacklisted : g_variant_new_string

// Blacklisted : g_variant_new_strv

// Unsupported : g_variant_new_tuple : unsupported parameter children :

// Blacklisted : g_variant_new_uint16

// Blacklisted : g_variant_new_uint32

// Blacklisted : g_variant_new_uint64

// Unsupported : g_variant_new_va : unsupported parameter endptr : in string with indirection level of 2

// Blacklisted : g_variant_new_variant

// Blacklisted : g_variant_is_object_path

// Blacklisted : g_variant_is_signature

// Blacklisted : g_variant_byteswap

// Blacklisted : g_variant_classify

// Blacklisted : g_variant_dup_string

// Blacklisted : g_variant_dup_strv

// Blacklisted : g_variant_equal

// Blacklisted : g_variant_get

// Blacklisted : g_variant_get_boolean

// Blacklisted : g_variant_get_byte

// Blacklisted : g_variant_get_child

// Blacklisted : g_variant_get_child_value

// Unsupported : g_variant_get_data : no return generator

// Blacklisted : g_variant_get_double

// Unsupported : g_variant_get_fixed_array : no type generator for gpointer (gconstpointer) for array return

// Blacklisted : g_variant_get_handle

// Blacklisted : g_variant_get_int16

// Blacklisted : g_variant_get_int32

// Blacklisted : g_variant_get_int64

// Blacklisted : g_variant_get_maybe

// Blacklisted : g_variant_get_normal_form

// Blacklisted : g_variant_get_size

// Blacklisted : g_variant_get_string

// Blacklisted : g_variant_get_strv

// Blacklisted : g_variant_get_type

// Blacklisted : g_variant_get_type_string

// Blacklisted : g_variant_get_uint16

// Blacklisted : g_variant_get_uint32

// Blacklisted : g_variant_get_uint64

// Unsupported : g_variant_get_va : unsupported parameter endptr : in string with indirection level of 2

// Blacklisted : g_variant_get_variant

// Blacklisted : g_variant_hash

// Blacklisted : g_variant_is_container

// Blacklisted : g_variant_is_normal_form

// Blacklisted : g_variant_is_of_type

// Blacklisted : g_variant_iter_new

// Blacklisted : g_variant_n_children

// Blacklisted : g_variant_print

// Blacklisted : g_variant_print_string

// Blacklisted : g_variant_ref

// Blacklisted : g_variant_ref_sink

// Unsupported : g_variant_store : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Blacklisted : g_variant_take_ref

// Blacklisted : g_variant_unref

// Blacklisted : g_variant_builder_new

// Blacklisted : g_variant_builder_add

// Blacklisted : g_variant_builder_add_value

// Blacklisted : g_variant_builder_clear

// Blacklisted : g_variant_builder_close

// Blacklisted : g_variant_builder_end

// Init is a wrapper around the C function g_variant_builder_init.
func (recv *VariantBuilder) Init(type_ *VariantType) {
	c_type := (*C.GVariantType)(C.NULL)
	if type_ != nil {
		c_type = (*C.GVariantType)(type_.ToC())
	}

	C.g_variant_builder_init((*C.GVariantBuilder)(recv.native), c_type)

	return
}

// Blacklisted : g_variant_builder_open

// Blacklisted : g_variant_builder_ref

// Blacklisted : g_variant_builder_unref

// Blacklisted : g_variant_iter_copy

// Blacklisted : g_variant_iter_free

// Init is a wrapper around the C function g_variant_iter_init.
func (recv *VariantIter) Init(value *Variant) uint64 {
	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	retC := C.g_variant_iter_init((*C.GVariantIter)(recv.native), c_value)
	retGo := (uint64)(retC)

	return retGo
}

// Blacklisted : g_variant_iter_loop

// Blacklisted : g_variant_iter_n_children

// Blacklisted : g_variant_iter_next

// Blacklisted : g_variant_iter_next_value

// Blacklisted : g_variant_type_new

// Blacklisted : g_variant_type_string_scan
