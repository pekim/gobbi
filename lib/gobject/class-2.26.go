// This is a generated file - DO NOT EDIT
// +build gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.44 gobject_2.54

package gobject

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Binding is a wrapper around the C record GBinding.
type Binding struct {
	native *C.GBinding
}

func bindingNewFromC(c *C.GBinding) *Binding {
	if c == nil {
		return nil
	}

	g := &Binding{native: c}

	return g
}

// Unsupported : g_binding_get_flags : no return generator

// Unsupported : g_binding_get_source : no return generator

// GetSourceProperty is a wrapper around the C function g_binding_get_source_property.
func (recv *Binding) GetSourceProperty() string {
	retC := C.g_binding_get_source_property(recv.native)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_binding_get_target : no return generator

// GetTargetProperty is a wrapper around the C function g_binding_get_target_property.
func (recv *Binding) GetTargetProperty() string {
	retC := C.g_binding_get_target_property(recv.native)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_binding_unbind : no return generator

// ParamSpecVariant is a wrapper around the C record GParamSpecVariant.
type ParamSpecVariant struct {
	native *C.GParamSpecVariant
	// parent_instance : no type generator for ParamSpec, GParamSpec
	// _type : no type generator for GLib.VariantType, GVariantType*
	// default_value : no type generator for GLib.Variant, GVariant*
	// no type for padding
}

func paramSpecVariantNewFromC(c *C.GParamSpecVariant) *ParamSpecVariant {
	if c == nil {
		return nil
	}

	g := &ParamSpecVariant{native: c}

	return g
}
