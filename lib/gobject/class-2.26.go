// This is a generated file - DO NOT EDIT
// +build gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import (
	"C"
	"unsafe"
)

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

// GetFlags is a wrapper around the C function g_binding_get_flags.
func (recv *Binding) GetFlags() BindingFlags {
	retC := C.g_binding_get_flags((*C.GBinding)(recv.native))
	retGo := (BindingFlags)(retC)

	return retGo
}

// GetSource is a wrapper around the C function g_binding_get_source.
func (recv *Binding) GetSource() *Object {
	retC := C.g_binding_get_source((*C.GBinding)(recv.native))
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSourceProperty is a wrapper around the C function g_binding_get_source_property.
func (recv *Binding) GetSourceProperty() string {
	retC := C.g_binding_get_source_property((*C.GBinding)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetTarget is a wrapper around the C function g_binding_get_target.
func (recv *Binding) GetTarget() *Object {
	retC := C.g_binding_get_target((*C.GBinding)(recv.native))
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetTargetProperty is a wrapper around the C function g_binding_get_target_property.
func (recv *Binding) GetTargetProperty() string {
	retC := C.g_binding_get_target_property((*C.GBinding)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

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

// BindPropertyWithClosures is a wrapper around the C function g_object_bind_property_with_closures.
func (recv *Object) BindPropertyWithClosures(sourceProperty string, target *Object, targetProperty string, flags BindingFlags, transformTo *Closure, transformFrom *Closure) *Binding {
	c_source_property := C.CString(sourceProperty)
	defer C.free(unsafe.Pointer(c_source_property))

	c_target := (C.gpointer)(C.NULL)
	if target != nil {
		c_target = (C.gpointer)(target.ToC())
	}

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

// NotifyByPspec is a wrapper around the C function g_object_notify_by_pspec.
func (recv *Object) NotifyByPspec(pspec *ParamSpec) {
	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	C.g_object_notify_by_pspec((*C.GObject)(recv.native), c_pspec)

	return
}

// Equals compares this ParamSpecVariant with another ParamSpecVariant, and returns true if they represent the same GObject.
func (recv *ParamSpecVariant) Equals(other *ParamSpecVariant) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecVariant) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}
