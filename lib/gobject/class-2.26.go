// This is a generated file - DO NOT EDIT
// +build gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// #GBinding is the representation of a binding between a property on a
// #GObject instance (or source) and another property on another #GObject
// instance (or target). Whenever the source property changes, the same
// value is applied to the target property; for instance, the following
// binding:
//
// |[<!-- language="C" -->
// g_object_bind_property (object1, "property-a",
// object2, "property-b",
// G_BINDING_DEFAULT);
// ]|
//
// will cause the property named "property-b" of @object2 to be updated
// every time g_object_set() or the specific accessor changes the value of
// the property "property-a" of @object1.
//
// It is possible to create a bidirectional binding between two properties
// of two #GObject instances, so that if either property changes, the
// other is updated as well, for instance:
//
// |[<!-- language="C" -->
// g_object_bind_property (object1, "property-a",
// object2, "property-b",
// G_BINDING_BIDIRECTIONAL);
// ]|
//
// will keep the two properties in sync.
//
// It is also possible to set a custom transformation function (in both
// directions, in case of a bidirectional binding) to apply a custom
// transformation from the source value to the target value before
// applying it; for instance, the following binding:
//
// |[<!-- language="C" -->
// g_object_bind_property_full (adjustment1, "value",
// adjustment2, "value",
// G_BINDING_BIDIRECTIONAL,
// celsius_to_fahrenheit,
// fahrenheit_to_celsius,
// NULL, NULL);
// ]|
//
// will keep the "value" property of the two adjustments in sync; the
// @celsius_to_fahrenheit function will be called whenever the "value"
// property of @adjustment1 changes and will transform the current value
// of the property before applying it to the "value" property of @adjustment2.
//
// Vice versa, the @fahrenheit_to_celsius function will be called whenever
// the "value" property of @adjustment2 changes, and will transform the
// current value of the property before applying it to the "value" property
// of @adjustment1.
//
// Note that #GBinding does not resolve cycles by itself; a cycle like
//
// |[
// object1:propertyA -> object2:propertyB
// object2:propertyB -> object3:propertyC
// object3:propertyC -> object1:propertyA
// ]|
//
// might lead to an infinite loop. The loop, in this particular case,
// can be avoided if the objects emit the #GObject::notify signal only
// if the value has effectively been changed. A binding is implemented
// using the #GObject::notify signal, so it is susceptible to all the
// various ways of blocking a signal emission, like g_signal_stop_emission()
// or g_signal_handler_block().
//
// A binding will be severed, and the resources it allocates freed, whenever
// either one of the #GObject instances it refers to are finalized, or when
// the #GBinding instance loses its last reference.
//
// Bindings for languages with garbage collection can use
// g_binding_unbind() to explicitly release a binding between the source
// and target properties, instead of relying on the last reference on the
// binding, source, and target instances to drop.
//
// #GBinding is available since GObject 2.26
/*

C type

GBinding
*/
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

C function

g_binding_get_flags
*/
func (recv *Binding) GetFlags() BindingFlags {
	retC := C.g_binding_get_flags((*C.GBinding)(recv.native))
	retGo := (BindingFlags)(retC)

	return retGo
}

// Retrieves the #GObject instance used as the source of the binding.
/*

C function

g_binding_get_source
*/
func (recv *Binding) GetSource() *Object {
	retC := C.g_binding_get_source((*C.GBinding)(recv.native))
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the name of the property of #GBinding:source used as the source
// of the binding.
/*

C function

g_binding_get_source_property
*/
func (recv *Binding) GetSourceProperty() string {
	retC := C.g_binding_get_source_property((*C.GBinding)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Retrieves the #GObject instance used as the target of the binding.
/*

C function

g_binding_get_target
*/
func (recv *Binding) GetTarget() *Object {
	retC := C.g_binding_get_target((*C.GBinding)(recv.native))
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the name of the property of #GBinding:target used as the target
// of the binding.
/*

C function

g_binding_get_target_property
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

C function

g_object_bind_property
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

C function

g_object_bind_property_with_closures
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

// A #GParamSpec derived structure that contains the meta data for #GVariant properties.
/*

C type

GParamSpecVariant
*/
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
