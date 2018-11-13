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

// Writes the contents of @key_file to @filename using
// g_file_set_contents().
//
// This function can fail for any of the reasons that
// g_file_set_contents() may fail.
/*

C function : g_key_file_save_to_file
*/
func (recv *KeyFile) SaveToFile(filename string) (bool, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.g_key_file_save_to_file((*C.GKeyFile)(recv.native), c_filename, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_option_context_parse_strv : unsupported parameter arguments :

// #GVariantDict is a mutable interface to #GVariant dictionaries.
//
// It can be used for doing a sequence of dictionary lookups in an
// efficient way on an existing #GVariant dictionary or it can be used
// to construct new dictionaries with a hashtable-like interface.  It
// can also be used for taking existing dictionaries and modifying them
// in order to create new ones.
//
// #GVariantDict can only be used with %G_VARIANT_TYPE_VARDICT
// dictionaries.
//
// It is possible to use #GVariantDict allocated on the stack or on the
// heap.  When using a stack-allocated #GVariantDict, you begin with a
// call to g_variant_dict_init() and free the resources with a call to
// g_variant_dict_clear().
//
// Heap-allocated #GVariantDict follows normal refcounting rules: you
// allocate it with g_variant_dict_new() and use g_variant_dict_ref()
// and g_variant_dict_unref().
//
// g_variant_dict_end() is used to convert the #GVariantDict back into a
// dictionary-type #GVariant.  When used with stack-allocated instances,
// this also implicitly frees all associated memory, but for
// heap-allocated instances, you must still call g_variant_dict_unref()
// afterwards.
//
// You will typically want to use a heap-allocated #GVariantDict when
// you expose it as part of an API.  For most other uses, the
// stack-allocated form will be more convenient.
//
// Consider the following two examples that do the same thing in each
// style: take an existing dictionary and look up the "count" uint32
// key, adding 1 to it if it is found, or returning an error if the
// key is not found.  Each returns the new dictionary as a floating
// #GVariant.
//
// ## Using a stack-allocated GVariantDict
//
// |[<!-- language="C" -->
// GVariant *
// add_to_count (GVariant  *orig,
// GError   **error)
// {
// GVariantDict dict;
// guint32 count;
//
// g_variant_dict_init (&dict, orig);
// if (!g_variant_dict_lookup (&dict, "count", "u", &count))
// {
// g_set_error (...);
// g_variant_dict_clear (&dict);
// return NULL;
// }
//
// g_variant_dict_insert (&dict, "count", "u", count + 1);
//
// return g_variant_dict_end (&dict);
// }
// ]|
//
// ## Using heap-allocated GVariantDict
//
// |[<!-- language="C" -->
// GVariant *
// add_to_count (GVariant  *orig,
// GError   **error)
// {
// GVariantDict *dict;
// GVariant *result;
// guint32 count;
//
// dict = g_variant_dict_new (orig);
//
// if (g_variant_dict_lookup (dict, "count", "u", &count))
// {
// g_variant_dict_insert (dict, "count", "u", count + 1);
// result = g_variant_dict_end (dict);
// }
// else
// {
// g_set_error (...);
// result = NULL;
// }
//
// g_variant_dict_unref (dict);
//
// return result;
// }
// ]|
/*

C record/class : GVariantDict
*/
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

// Unsupported : g_variant_dict_new : unsupported parameter from_asv : Blacklisted record : GVariant

// Releases all memory associated with a #GVariantDict without freeing
// the #GVariantDict structure itself.
//
// It typically only makes sense to do this on a stack-allocated
// #GVariantDict if you want to abort building the value part-way
// through.  This function need not be called if you call
// g_variant_dict_end() and it also doesn't need to be called on dicts
// allocated with g_variant_dict_new (see g_variant_dict_unref() for
// that).
//
// It is valid to call this function on either an initialised
// #GVariantDict or one that was previously cleared by an earlier call
// to g_variant_dict_clear() but it is not valid to call this function
// on uninitialised memory.
/*

C function : g_variant_dict_clear
*/
func (recv *VariantDict) Clear() {
	C.g_variant_dict_clear((*C.GVariantDict)(recv.native))

	return
}

// Checks if @key exists in @dict.
/*

C function : g_variant_dict_contains
*/
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

// Increases the reference count on @dict.
//
// Don't call this on stack-allocated #GVariantDict instances or bad
// things will happen.
/*

C function : g_variant_dict_ref
*/
func (recv *VariantDict) Ref() *VariantDict {
	retC := C.g_variant_dict_ref((*C.GVariantDict)(recv.native))
	retGo := VariantDictNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Removes a key and its associated value from a #GVariantDict.
/*

C function : g_variant_dict_remove
*/
func (recv *VariantDict) Remove(key string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_variant_dict_remove((*C.GVariantDict)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// Decreases the reference count on @dict.
//
// In the event that there are no more references, releases all memory
// associated with the #GVariantDict.
//
// Don't call this on stack-allocated #GVariantDict instances or bad
// things will happen.
/*

C function : g_variant_dict_unref
*/
func (recv *VariantDict) Unref() {
	C.g_variant_dict_unref((*C.GVariantDict)(recv.native))

	return
}
