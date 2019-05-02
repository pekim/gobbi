// This is a generated file - DO NOT EDIT
// +build glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import (
	"fmt"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
/*

	static GVariant* _g_variant_new(const gchar* format_string) {
		return g_variant_new(format_string);
    }
*/
/*

	static GVariant* _g_variant_new_parsed(const gchar* format) {
		return g_variant_new_parsed(format);
    }
*/
import "C"

// Unsupported : g_malloc0_n : no return generator

// Unsupported : g_malloc_n : no return generator

// Unsupported : g_realloc_n : unsupported parameter mem : no type generator for gpointer (gpointer) for param mem

// Unsupported : g_try_malloc0_n : no return generator

// Unsupported : g_try_malloc_n : no return generator

// Unsupported : g_try_realloc_n : unsupported parameter mem : no type generator for gpointer (gpointer) for param mem

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

// VariantNew is a wrapper around the C function g_variant_new.
func VariantNew(formatString string, args ...interface{}) *Variant {
	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	retC := C._g_variant_new(c_format_string)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_new_array : unsupported parameter children :

// VariantNewBoolean is a wrapper around the C function g_variant_new_boolean.
func VariantNewBoolean(value bool) *Variant {
	c_value :=
		boolToGboolean(value)

	retC := C.g_variant_new_boolean(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewByte is a wrapper around the C function g_variant_new_byte.
func VariantNewByte(value uint8) *Variant {
	c_value := (C.guchar)(value)

	retC := C.g_variant_new_byte(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewDictEntry is a wrapper around the C function g_variant_new_dict_entry.
func VariantNewDictEntry(key *Variant, value *Variant) *Variant {
	c_key := (*C.GVariant)(C.NULL)
	if key != nil {
		c_key = (*C.GVariant)(key.ToC())
	}

	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	retC := C.g_variant_new_dict_entry(c_key, c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewDouble is a wrapper around the C function g_variant_new_double.
func VariantNewDouble(value float64) *Variant {
	c_value := (C.gdouble)(value)

	retC := C.g_variant_new_double(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_new_from_data : unsupported parameter notify : no type generator for DestroyNotify (GDestroyNotify) for param notify

// VariantNewHandle is a wrapper around the C function g_variant_new_handle.
func VariantNewHandle(value int32) *Variant {
	c_value := (C.gint32)(value)

	retC := C.g_variant_new_handle(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewInt16 is a wrapper around the C function g_variant_new_int16.
func VariantNewInt16(value int16) *Variant {
	c_value := (C.gint16)(value)

	retC := C.g_variant_new_int16(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewInt32 is a wrapper around the C function g_variant_new_int32.
func VariantNewInt32(value int32) *Variant {
	c_value := (C.gint32)(value)

	retC := C.g_variant_new_int32(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewInt64 is a wrapper around the C function g_variant_new_int64.
func VariantNewInt64(value int64) *Variant {
	c_value := (C.gint64)(value)

	retC := C.g_variant_new_int64(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewMaybe is a wrapper around the C function g_variant_new_maybe.
func VariantNewMaybe(childType *VariantType, child *Variant) *Variant {
	c_child_type := (*C.GVariantType)(C.NULL)
	if childType != nil {
		c_child_type = (*C.GVariantType)(childType.ToC())
	}

	c_child := (*C.GVariant)(C.NULL)
	if child != nil {
		c_child = (*C.GVariant)(child.ToC())
	}

	retC := C.g_variant_new_maybe(c_child_type, c_child)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewObjectPath is a wrapper around the C function g_variant_new_object_path.
func VariantNewObjectPath(objectPath string) *Variant {
	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	retC := C.g_variant_new_object_path(c_object_path)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewParsed is a wrapper around the C function g_variant_new_parsed.
func VariantNewParsed(format string, args ...interface{}) *Variant {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	retC := C._g_variant_new_parsed(c_format)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_new_parsed_va : unsupported parameter app : no type generator for va_list (va_list*) for param app

// VariantNewSignature is a wrapper around the C function g_variant_new_signature.
func VariantNewSignature(signature string) *Variant {
	c_signature := C.CString(signature)
	defer C.free(unsafe.Pointer(c_signature))

	retC := C.g_variant_new_signature(c_signature)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewString is a wrapper around the C function g_variant_new_string.
func VariantNewString(string_ string) *Variant {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_variant_new_string(c_string)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewStrv is a wrapper around the C function g_variant_new_strv.
func VariantNewStrv(strv []string) *Variant {
	c_strv_array := make([]*C.gchar, len(strv)+1, len(strv)+1)
	for i, item := range strv {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_strv_array[i] = c
	}
	c_strv_array[len(strv)] = nil
	c_strv_arrayPtr := &c_strv_array[0]
	c_strv := (**C.gchar)(unsafe.Pointer(c_strv_arrayPtr))

	c_length := (C.gssize)(len(strv))

	retC := C.g_variant_new_strv(c_strv, c_length)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_new_tuple : unsupported parameter children :

// VariantNewUint16 is a wrapper around the C function g_variant_new_uint16.
func VariantNewUint16(value uint16) *Variant {
	c_value := (C.guint16)(value)

	retC := C.g_variant_new_uint16(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewUint32 is a wrapper around the C function g_variant_new_uint32.
func VariantNewUint32(value uint32) *Variant {
	c_value := (C.guint32)(value)

	retC := C.g_variant_new_uint32(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewUint64 is a wrapper around the C function g_variant_new_uint64.
func VariantNewUint64(value uint64) *Variant {
	c_value := (C.guint64)(value)

	retC := C.g_variant_new_uint64(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_new_va : unsupported parameter endptr : in string with indirection level of 2

// VariantNewVariant is a wrapper around the C function g_variant_new_variant.
func VariantNewVariant(value *Variant) *Variant {
	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	retC := C.g_variant_new_variant(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantBuilderNew is a wrapper around the C function g_variant_builder_new.
func VariantBuilderNew(type_ *VariantType) *VariantBuilder {
	c_type := (*C.GVariantType)(C.NULL)
	if type_ != nil {
		c_type = (*C.GVariantType)(type_.ToC())
	}

	retC := C.g_variant_builder_new(c_type)
	retGo := VariantBuilderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantTypeNew is a wrapper around the C function g_variant_type_new.
func VariantTypeNew(typeString string) *VariantType {
	c_type_string := C.CString(typeString)
	defer C.free(unsafe.Pointer(c_type_string))

	retC := C.g_variant_type_new(c_type_string)
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}
