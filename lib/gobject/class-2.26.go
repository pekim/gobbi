// This is a generated file - DO NOT EDIT
// +build gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import (
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Binding is a wrapper around the C record GBinding.
type Binding struct {
	native *C.GBinding
}

func BindingNewFromC(u unsafe.Pointer) *Binding {
	c := (*C.GBinding)(u)
	if c == nil {
		return nil
	}

	g := &Binding{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Binding) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Binding) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Binding with another Binding, and returns true if they represent the same GObject.
func (recv *Binding) Equals(other *Binding) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Binding) Object() *Object {
	return ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Binding.
// Exercise care, as this is a potentially dangerous function if the Object is not a Binding.
func CastToBinding(object *Object) *Binding {
	return BindingNewFromC(object.ToC())
}

// Blacklisted : g_binding_get_flags

// Blacklisted : g_binding_get_source

// Blacklisted : g_binding_get_source_property

// Blacklisted : g_binding_get_target

// Blacklisted : g_binding_get_target_property

// BindProperty is a wrapper around the C function g_object_bind_property.
func (recv *Object) BindProperty(sourceProperty string, target *Object, targetProperty string, flags BindingFlags) *Binding {
	c_source_property := C.CString(sourceProperty)
	defer C.free(unsafe.Pointer(c_source_property))

	c_target := (C.gpointer)(C.NULL)
	if target != nil {
		c_target = (C.gpointer)(target.ToC())
	}

	c_target_property := C.CString(targetProperty)
	defer C.free(unsafe.Pointer(c_target_property))

	c_flags := (C.GBindingFlags)(flags)

	retC := C.g_object_bind_property((C.gpointer)(recv.native), c_source_property, c_target, c_target_property, c_flags)
	retGo := BindingNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_object_bind_property_full : unsupported parameter transform_to : no type generator for BindingTransformFunc (GBindingTransformFunc) for param transform_to

// Blacklisted : g_object_bind_property_with_closures

// Blacklisted : g_object_notify_by_pspec

// Blacklisted : GParamSpecVariant
