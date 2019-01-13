// This is a generated file - DO NOT EDIT
// +build glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

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

	static void _g_variant_dict_insert(GVariantDict* dict, const gchar* key, const gchar* format_string) {
		return g_variant_dict_insert(dict, key, format_string);
    }
*/
/*

	static gboolean _g_variant_dict_lookup(GVariantDict* dict, const gchar* key, const gchar* format_string) {
		return g_variant_dict_lookup(dict, key, format_string);
    }
*/
import "C"

// g_hash_table_get_keys_as_array : no type generator for gpointer (gpointer) for array return
// SaveToFile is a wrapper around the C function g_key_file_save_to_file.
func (recv *KeyFile) SaveToFile(filename string) (bool, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.g_key_file_save_to_file((*C.GKeyFile)(recv.native), c_filename, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Blacklisted : g_option_context_parse_strv

// VariantParseErrorPrintContext is a wrapper around the C function g_variant_parse_error_print_context.
func VariantParseErrorPrintContext(error *Error, sourceStr string) string {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	c_source_str := C.CString(sourceStr)
	defer C.free(unsafe.Pointer(c_source_str))

	retC := C.g_variant_parse_error_print_context(c_error, c_source_str)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

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

// Clear is a wrapper around the C function g_variant_dict_clear.
func (recv *VariantDict) Clear() {
	C.g_variant_dict_clear((*C.GVariantDict)(recv.native))

	return
}

// Contains is a wrapper around the C function g_variant_dict_contains.
func (recv *VariantDict) Contains(key string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_variant_dict_contains((*C.GVariantDict)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// End is a wrapper around the C function g_variant_dict_end.
func (recv *VariantDict) End() *Variant {
	retC := C.g_variant_dict_end((*C.GVariantDict)(recv.native))
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Init is a wrapper around the C function g_variant_dict_init.
func (recv *VariantDict) Init(fromAsv *Variant) {
	c_from_asv := (*C.GVariant)(C.NULL)
	if fromAsv != nil {
		c_from_asv = (*C.GVariant)(fromAsv.ToC())
	}

	C.g_variant_dict_init((*C.GVariantDict)(recv.native), c_from_asv)

	return
}

// Insert is a wrapper around the C function g_variant_dict_insert.
func (recv *VariantDict) Insert(key string, formatString string, args ...interface{}) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	C._g_variant_dict_insert((*C.GVariantDict)(recv.native), c_key, c_format_string)

	return
}

// InsertValue is a wrapper around the C function g_variant_dict_insert_value.
func (recv *VariantDict) InsertValue(key string, value *Variant) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	C.g_variant_dict_insert_value((*C.GVariantDict)(recv.native), c_key, c_value)

	return
}

// Lookup is a wrapper around the C function g_variant_dict_lookup.
func (recv *VariantDict) Lookup(key string, formatString string, args ...interface{}) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	retC := C._g_variant_dict_lookup((*C.GVariantDict)(recv.native), c_key, c_format_string)
	retGo := retC == C.TRUE

	return retGo
}

// LookupValue is a wrapper around the C function g_variant_dict_lookup_value.
func (recv *VariantDict) LookupValue(key string, expectedType *VariantType) *Variant {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_expected_type := (*C.GVariantType)(C.NULL)
	if expectedType != nil {
		c_expected_type = (*C.GVariantType)(expectedType.ToC())
	}

	retC := C.g_variant_dict_lookup_value((*C.GVariantDict)(recv.native), c_key, c_expected_type)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ref is a wrapper around the C function g_variant_dict_ref.
func (recv *VariantDict) Ref() *VariantDict {
	retC := C.g_variant_dict_ref((*C.GVariantDict)(recv.native))
	retGo := VariantDictNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Remove is a wrapper around the C function g_variant_dict_remove.
func (recv *VariantDict) Remove(key string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_variant_dict_remove((*C.GVariantDict)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// Unref is a wrapper around the C function g_variant_dict_unref.
func (recv *VariantDict) Unref() {
	C.g_variant_dict_unref((*C.GVariantDict)(recv.native))

	return
}
