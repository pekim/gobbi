// This is a generated file - DO NOT EDIT
// +build glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Blacklisted : GVariant

// Unsupported : g_variant_builder_new : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_variant_builder_add : unsupported parameter ... : varargs

// Unsupported : g_variant_builder_add_value : unsupported parameter value : Blacklisted record : GVariant

// Releases all memory associated with a #GVariantBuilder without
// freeing the #GVariantBuilder structure itself.
//
// It typically only makes sense to do this on a stack-allocated
// #GVariantBuilder if you want to abort building the value part-way
// through.  This function need not be called if you call
// g_variant_builder_end() and it also doesn't need to be called on
// builders allocated with g_variant_builder_new() (see
// g_variant_builder_unref() for that).
//
// This function leaves the #GVariantBuilder structure set to all-zeros.
// It is valid to call this function on either an initialised
// #GVariantBuilder or one that is set to all-zeros but it is not valid
// to call this function on uninitialised memory.
/*

C function : g_variant_builder_clear
*/
func (recv *VariantBuilder) Clear() {
	C.g_variant_builder_clear((*C.GVariantBuilder)(recv.native))

	return
}

// Closes the subcontainer inside the given @builder that was opened by
// the most recent call to g_variant_builder_open().
//
// It is an error to call this function in any way that would create an
// inconsistent value to be constructed (ie: too few values added to the
// subcontainer).
/*

C function : g_variant_builder_close
*/
func (recv *VariantBuilder) Close() {
	C.g_variant_builder_close((*C.GVariantBuilder)(recv.native))

	return
}

// Unsupported : g_variant_builder_end : return type : Blacklisted record : GVariant

// Unsupported : g_variant_builder_init : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_variant_builder_open : unsupported parameter type : Blacklisted record : GVariantType

// Increases the reference count on @builder.
//
// Don't call this on stack-allocated #GVariantBuilder instances or bad
// things will happen.
/*

C function : g_variant_builder_ref
*/
func (recv *VariantBuilder) Ref() *VariantBuilder {
	retC := C.g_variant_builder_ref((*C.GVariantBuilder)(recv.native))
	retGo := VariantBuilderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Decreases the reference count on @builder.
//
// In the event that there are no more references, releases all memory
// associated with the #GVariantBuilder.
//
// Don't call this on stack-allocated #GVariantBuilder instances or bad
// things will happen.
/*

C function : g_variant_builder_unref
*/
func (recv *VariantBuilder) Unref() {
	C.g_variant_builder_unref((*C.GVariantBuilder)(recv.native))

	return
}

// Creates a new heap-allocated #GVariantIter to iterate over the
// container that was being iterated over by @iter.  Iteration begins on
// the new iterator from the current position of the old iterator but
// the two copies are independent past that point.
//
// Use g_variant_iter_free() to free the return value when you no longer
// need it.
//
// A reference is taken to the container that @iter is iterating over
// and will be releated only when g_variant_iter_free() is called.
/*

C function : g_variant_iter_copy
*/
func (recv *VariantIter) Copy() *VariantIter {
	retC := C.g_variant_iter_copy((*C.GVariantIter)(recv.native))
	retGo := VariantIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Frees a heap-allocated #GVariantIter.  Only call this function on
// iterators that were returned by g_variant_iter_new() or
// g_variant_iter_copy().
/*

C function : g_variant_iter_free
*/
func (recv *VariantIter) Free() {
	C.g_variant_iter_free((*C.GVariantIter)(recv.native))

	return
}

// Unsupported : g_variant_iter_init : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_variant_iter_loop : unsupported parameter ... : varargs

// Queries the number of child items in the container that we are
// iterating over.  This is the total number of items -- not the number
// of items remaining.
//
// This function might be useful for preallocation of arrays.
/*

C function : g_variant_iter_n_children
*/
func (recv *VariantIter) NChildren() uint64 {
	retC := C.g_variant_iter_n_children((*C.GVariantIter)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported : g_variant_iter_next : unsupported parameter ... : varargs

// Unsupported : g_variant_iter_next_value : return type : Blacklisted record : GVariant
