// This is a generated file - DO NOT EDIT
// +build gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// This function is intended for #GObject implementations to re-enforce
// a [floating][floating-ref] object reference. Doing this is seldom
// required: all #GInitiallyUnowneds are created with a floating reference
// which usually just needs to be sunken by calling g_object_ref_sink().
/*

C function : g_object_force_floating
*/
func (recv *Object) ForceFloating() {
	C.g_object_force_floating((*C.GObject)(recv.native))

	return
}

// Checks whether @object has a [floating][floating-ref] reference.
/*

C function : g_object_is_floating
*/
func (recv *Object) IsFloating() bool {
	retC := C.g_object_is_floating((C.gpointer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Increase the reference count of @object, and possibly remove the
// [floating][floating-ref] reference, if @object has a floating reference.
//
// In other words, if the object is floating, then this call "assumes
// ownership" of the floating reference, converting it to a normal
// reference by clearing the floating flag while leaving the reference
// count unchanged.  If the object is not floating, then this call
// adds a new normal reference increasing the reference count by one.
//
// Since GLib 2.56, the type of @object will be propagated to the return type
// under the same conditions as for g_object_ref().
/*

C function : g_object_ref_sink
*/
func (recv *Object) RefSink() uintptr {
	retC := C.g_object_ref_sink((C.gpointer)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// A #GParamSpec derived structure that contains the meta data for #GType properties.
/*

C record/class : GParamSpecGType
*/
type ParamSpecGType struct {
	native *C.GParamSpecGType
	// parent_instance : record
	IsAType Type
}

func ParamSpecGTypeNewFromC(u unsafe.Pointer) *ParamSpecGType {
	c := (*C.GParamSpecGType)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecGType{
		IsAType: (Type)(c.is_a_type),
		native:  c,
	}

	return g
}

func (recv *ParamSpecGType) ToC() unsafe.Pointer {
	recv.native.is_a_type =
		(C.GType)(recv.IsAType)

	return (unsafe.Pointer)(recv.native)
}
