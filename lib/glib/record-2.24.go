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

// Clear is a wrapper around the C function g_variant_builder_clear.
func (recv *VariantBuilder) Clear() {
	C.g_variant_builder_clear((*C.GVariantBuilder)(recv.native))

	return
}

// Close is a wrapper around the C function g_variant_builder_close.
func (recv *VariantBuilder) Close() {
	C.g_variant_builder_close((*C.GVariantBuilder)(recv.native))

	return
}

// Unsupported : g_variant_builder_end : return type : Blacklisted record : GVariant

// Unsupported : g_variant_builder_init : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_variant_builder_open : unsupported parameter type : Blacklisted record : GVariantType

// Ref is a wrapper around the C function g_variant_builder_ref.
func (recv *VariantBuilder) Ref() *VariantBuilder {
	retC := C.g_variant_builder_ref((*C.GVariantBuilder)(recv.native))
	retGo := VariantBuilderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_variant_builder_unref.
func (recv *VariantBuilder) Unref() {
	C.g_variant_builder_unref((*C.GVariantBuilder)(recv.native))

	return
}

// Copy is a wrapper around the C function g_variant_iter_copy.
func (recv *VariantIter) Copy() *VariantIter {
	retC := C.g_variant_iter_copy((*C.GVariantIter)(recv.native))
	retGo := VariantIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function g_variant_iter_free.
func (recv *VariantIter) Free() {
	C.g_variant_iter_free((*C.GVariantIter)(recv.native))

	return
}

// Unsupported : g_variant_iter_init : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_variant_iter_loop : unsupported parameter ... : varargs

// NChildren is a wrapper around the C function g_variant_iter_n_children.
func (recv *VariantIter) NChildren() uint64 {
	retC := C.g_variant_iter_n_children((*C.GVariantIter)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported : g_variant_iter_next : unsupported parameter ... : varargs

// Unsupported : g_variant_iter_next_value : return type : Blacklisted record : GVariant
