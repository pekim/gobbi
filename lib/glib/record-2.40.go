// This is a generated file - DO NOT EDIT
// +build glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// g_hash_table_get_keys_as_array : no return type
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

// Unsupported : g_option_context_parse_strv : unsupported parameter arguments :

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

// Unsupported : g_variant_dict_new : unsupported parameter from_asv : Blacklisted record : GVariant

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

// Unsupported : g_variant_dict_end : return type : Blacklisted record : GVariant

// Unsupported : g_variant_dict_init : unsupported parameter from_asv : Blacklisted record : GVariant

// Unsupported : g_variant_dict_insert : unsupported parameter ... : varargs

// Unsupported : g_variant_dict_insert_value : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_variant_dict_lookup : unsupported parameter ... : varargs

// Unsupported : g_variant_dict_lookup_value : unsupported parameter expected_type : Blacklisted record : GVariantType

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
