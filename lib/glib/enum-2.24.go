// This is a generated file - DO NOT EDIT
// +build glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// The range of possible top-level types of #GVariant instances.
type VariantClass C.GVariantClass

const (
	// The #GVariant is a boolean.
	VARIANT_CLASS_BOOLEAN VariantClass = 98
	// The #GVariant is a byte.
	VARIANT_CLASS_BYTE VariantClass = 121
	// The #GVariant is a signed 16 bit integer.
	VARIANT_CLASS_INT16 VariantClass = 110
	// The #GVariant is an unsigned 16 bit integer.
	VARIANT_CLASS_UINT16 VariantClass = 113
	// The #GVariant is a signed 32 bit integer.
	VARIANT_CLASS_INT32 VariantClass = 105
	// The #GVariant is an unsigned 32 bit integer.
	VARIANT_CLASS_UINT32 VariantClass = 117
	// The #GVariant is a signed 64 bit integer.
	VARIANT_CLASS_INT64 VariantClass = 120
	// The #GVariant is an unsigned 64 bit integer.
	VARIANT_CLASS_UINT64 VariantClass = 116
	// The #GVariant is a file handle index.
	VARIANT_CLASS_HANDLE VariantClass = 104
	/*
	   The #GVariant is a double precision floating
	                            point value.
	*/
	VARIANT_CLASS_DOUBLE VariantClass = 100
	// The #GVariant is a normal string.
	VARIANT_CLASS_STRING VariantClass = 115
	/*
	   The #GVariant is a D-Bus object path
	                                 string.
	*/
	VARIANT_CLASS_OBJECT_PATH VariantClass = 111
	// The #GVariant is a D-Bus signature string.
	VARIANT_CLASS_SIGNATURE VariantClass = 103
	// The #GVariant is a variant.
	VARIANT_CLASS_VARIANT VariantClass = 118
	// The #GVariant is a maybe-typed value.
	VARIANT_CLASS_MAYBE VariantClass = 109
	// The #GVariant is an array.
	VARIANT_CLASS_ARRAY VariantClass = 97
	// The #GVariant is a tuple.
	VARIANT_CLASS_TUPLE VariantClass = 40
	// The #GVariant is a dictionary entry.
	VARIANT_CLASS_DICT_ENTRY VariantClass = 123
)
