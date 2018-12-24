// This is a generated file - DO NOT EDIT
// +build glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

type VariantClass C.GVariantClass

const (
	VARIANT_CLASS_BOOLEAN     VariantClass = 98
	VARIANT_CLASS_BYTE        VariantClass = 121
	VARIANT_CLASS_INT16       VariantClass = 110
	VARIANT_CLASS_UINT16      VariantClass = 113
	VARIANT_CLASS_INT32       VariantClass = 105
	VARIANT_CLASS_UINT32      VariantClass = 117
	VARIANT_CLASS_INT64       VariantClass = 120
	VARIANT_CLASS_UINT64      VariantClass = 116
	VARIANT_CLASS_HANDLE      VariantClass = 104
	VARIANT_CLASS_DOUBLE      VariantClass = 100
	VARIANT_CLASS_STRING      VariantClass = 115
	VARIANT_CLASS_OBJECT_PATH VariantClass = 111
	VARIANT_CLASS_SIGNATURE   VariantClass = 103
	VARIANT_CLASS_VARIANT     VariantClass = 118
	VARIANT_CLASS_MAYBE       VariantClass = 109
	VARIANT_CLASS_ARRAY       VariantClass = 97
	VARIANT_CLASS_TUPLE       VariantClass = 40
	VARIANT_CLASS_DICT_ENTRY  VariantClass = 123
)
