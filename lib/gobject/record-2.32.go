// This is a generated file - DO NOT EDIT
// +build gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Get the contents of a %G_TYPE_CHAR #GValue.
/*

C function

g_value_get_schar
*/
func (recv *Value) GetSchar() int8 {
	retC := C.g_value_get_schar((*C.GValue)(recv.native))
	retGo := (int8)(retC)

	return retGo
}

// Set the contents of a %G_TYPE_CHAR #GValue to @v_char.
/*

C function

g_value_set_schar
*/
func (recv *Value) SetSchar(vChar int8) {
	c_v_char := (C.gint8)(vChar)

	C.g_value_set_schar((*C.GValue)(recv.native), c_v_char)

	return
}

// Frees resources associated with a non-statically-allocated #GWeakRef.
// After this call, the #GWeakRef is left in an undefined state.
//
// You should only call this on a #GWeakRef that previously had
// g_weak_ref_init() called on it.
/*

C function

g_weak_ref_clear
*/
func (recv *WeakRef) Clear() {
	C.g_weak_ref_clear((*C.GWeakRef)(recv.native))

	return
}

// If @weak_ref is not empty, atomically acquire a strong
// reference to the object it points to, and return that reference.
//
// This function is needed because of the potential race between taking
// the pointer value and g_object_ref() on it, if the object was losing
// its last reference at the same time in a different thread.
//
// The caller should release the resulting reference in the usual way,
// by using g_object_unref().
/*

C function

g_weak_ref_get
*/
func (recv *WeakRef) Get() uintptr {
	retC := C.g_weak_ref_get((*C.GWeakRef)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// Initialise a non-statically-allocated #GWeakRef.
//
// This function also calls g_weak_ref_set() with @object on the
// freshly-initialised weak reference.
//
// This function should always be matched with a call to
// g_weak_ref_clear().  It is not necessary to use this function for a
// #GWeakRef in static storage because it will already be
// properly initialised.  Just use g_weak_ref_set() directly.
/*

C function

g_weak_ref_init
*/
func (recv *WeakRef) Init(object uintptr) {
	c_object := (C.gpointer)(object)

	C.g_weak_ref_init((*C.GWeakRef)(recv.native), c_object)

	return
}

// Change the object to which @weak_ref points, or set it to
// %NULL.
//
// You must own a strong reference on @object while calling this
// function.
/*

C function

g_weak_ref_set
*/
func (recv *WeakRef) Set(object uintptr) {
	c_object := (C.gpointer)(object)

	C.g_weak_ref_set((*C.GWeakRef)(recv.native), c_object)

	return
}
