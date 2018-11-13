// This is a generated file - DO NOT EDIT
// +build gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
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

	return g
}

func (recv *Binding) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Binding) Object() *Object {
	return ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Binding.
// Exercise care, as this is a potentially dangerous function if the Object is not a Binding.
func CastToBinding(object *Object) *Binding {
	return BindingNewFromC(object.ToC())
}

// Retrieves the flags passed when constructing the #GBinding.
/*

C function : g_binding_get_flags
*/
func (recv *Binding) GetFlags() BindingFlags {
	retC := C.g_binding_get_flags((*C.GBinding)(recv.native))
	retGo := (BindingFlags)(retC)

	return retGo
}

// Retrieves the #GObject instance used as the source of the binding.
/*

C function : g_binding_get_source
*/
func (recv *Binding) GetSource() *Object {
	retC := C.g_binding_get_source((*C.GBinding)(recv.native))
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the name of the property of #GBinding:source used as the source
// of the binding.
/*

C function : g_binding_get_source_property
*/
func (recv *Binding) GetSourceProperty() string {
	retC := C.g_binding_get_source_property((*C.GBinding)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Retrieves the #GObject instance used as the target of the binding.
/*

C function : g_binding_get_target
*/
func (recv *Binding) GetTarget() *Object {
	retC := C.g_binding_get_target((*C.GBinding)(recv.native))
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the name of the property of #GBinding:target used as the target
// of the binding.
/*

C function : g_binding_get_target_property
*/
func (recv *Binding) GetTargetProperty() string {
	retC := C.g_binding_get_target_property((*C.GBinding)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Creates a binding between @source_property on @source and @target_property
// on @target. Whenever the @source_property is changed the @target_property is
// updated using the same value. For instance:
//
// |[
// g_object_bind_property (action, "active", widget, "sensitive", 0);
// ]|
//
// Will result in the "sensitive" property of the widget #GObject instance to be
// updated with the same value of the "active" property of the action #GObject
// instance.
//
// If @flags contains %G_BINDING_BIDIRECTIONAL then the binding will be mutual:
// if @target_property on @target changes then the @source_property on @source
// will be updated as well.
//
// The binding will automatically be removed when either the @source or the
// @target instances are finalized. To remove the binding without affecting the
// @source and the @target you can just call g_object_unref() on the returned
// #GBinding instance.
//
// A #GObject can have multiple bindings.
/*

C function : g_object_bind_property
*/
func (recv *Object) BindProperty(sourceProperty string, target uintptr, targetProperty string, flags BindingFlags) *Binding {
	c_source_property := C.CString(sourceProperty)
	defer C.free(unsafe.Pointer(c_source_property))

	c_target := (C.gpointer)(target)

	c_target_property := C.CString(targetProperty)
	defer C.free(unsafe.Pointer(c_target_property))

	c_flags := (C.GBindingFlags)(flags)

	retC := C.g_object_bind_property((C.gpointer)(recv.native), c_source_property, c_target, c_target_property, c_flags)
	retGo := BindingNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_object_bind_property_full : unsupported parameter transform_to : no type generator for BindingTransformFunc (GBindingTransformFunc) for param transform_to

// Creates a binding between @source_property on @source and @target_property
// on @target, allowing you to set the transformation functions to be used by
// the binding.
//
// This function is the language bindings friendly version of
// g_object_bind_property_full(), using #GClosures instead of
// function pointers.
/*

C function : g_object_bind_property_with_closures
*/
func (recv *Object) BindPropertyWithClosures(sourceProperty string, target uintptr, targetProperty string, flags BindingFlags, transformTo *Closure, transformFrom *Closure) *Binding {
	c_source_property := C.CString(sourceProperty)
	defer C.free(unsafe.Pointer(c_source_property))

	c_target := (C.gpointer)(target)

	c_target_property := C.CString(targetProperty)
	defer C.free(unsafe.Pointer(c_target_property))

	c_flags := (C.GBindingFlags)(flags)

	c_transform_to := (*C.GClosure)(C.NULL)
	if transformTo != nil {
		c_transform_to = (*C.GClosure)(transformTo.ToC())
	}

	c_transform_from := (*C.GClosure)(C.NULL)
	if transformFrom != nil {
		c_transform_from = (*C.GClosure)(transformFrom.ToC())
	}

	retC := C.g_object_bind_property_with_closures((C.gpointer)(recv.native), c_source_property, c_target, c_target_property, c_flags, c_transform_to, c_transform_from)
	retGo := BindingNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_object_notify_by_pspec : unsupported parameter pspec : Blacklisted record : GParamSpec

// ParamSpecVariant is a wrapper around the C record GParamSpecVariant.
type ParamSpecVariant struct {
	native *C.GParamSpecVariant
	// parent_instance : record
	// _type : record
	// default_value : record
	// Private : padding
}

func ParamSpecVariantNewFromC(u unsafe.Pointer) *ParamSpecVariant {
	c := (*C.GParamSpecVariant)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecVariant{native: c}

	return g
}

func (recv *ParamSpecVariant) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
