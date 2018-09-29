// This is a generated file - DO NOT EDIT
// +build gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_object_new : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_new_valist : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_new_with_properties : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_newv : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_add_toggle_ref : unsupported parameter notify : no type generator for ToggleNotify, GToggleNotify

// Unsupported : g_object_add_weak_pointer : unsupported parameter weak_pointer_location : no type generator for gpointer, gpointer*

// Unsupported : g_object_bind_property_full : unsupported parameter transform_to : no type generator for BindingTransformFunc, GBindingTransformFunc

// Unsupported : g_object_connect : unsupported parameter ... : varargs

// Unsupported : g_object_disconnect : unsupported parameter ... : varargs

// Unsupported : g_object_dup_data : unsupported parameter dup_func : no type generator for GLib.DuplicateFunc, GDuplicateFunc

// Unsupported : g_object_dup_qdata : unsupported parameter dup_func : no type generator for GLib.DuplicateFunc, GDuplicateFunc

// ForceFloating is a wrapper around the C function g_object_force_floating.
func (recv *Object) ForceFloating() {
	C.g_object_force_floating((*C.GObject)(recv.native))

	return
}

// Unsupported : g_object_get : unsupported parameter ... : varargs

// Unsupported : g_object_get_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : g_object_getv : unsupported parameter names : no param type

// IsFloating is a wrapper around the C function g_object_is_floating.
func (recv *Object) IsFloating() bool {
	retC := C.g_object_is_floating((C.gpointer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// RefSink is a wrapper around the C function g_object_ref_sink.
func (recv *Object) RefSink() uintptr {
	retC := C.g_object_ref_sink((C.gpointer)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_object_remove_toggle_ref : unsupported parameter notify : no type generator for ToggleNotify, GToggleNotify

// Unsupported : g_object_remove_weak_pointer : unsupported parameter weak_pointer_location : no type generator for gpointer, gpointer*

// Unsupported : g_object_replace_data : unsupported parameter destroy : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_object_replace_qdata : unsupported parameter destroy : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_object_set : unsupported parameter ... : varargs

// Unsupported : g_object_set_data_full : unsupported parameter destroy : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_object_set_qdata_full : unsupported parameter destroy : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_object_set_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : g_object_setv : unsupported parameter names : no param type

// Unsupported : g_object_weak_ref : unsupported parameter notify : no type generator for WeakNotify, GWeakNotify

// Unsupported : g_object_weak_unref : unsupported parameter notify : no type generator for WeakNotify, GWeakNotify

// RefSink is a wrapper around the C function g_param_spec_ref_sink.
func (recv *ParamSpec) RefSink() *ParamSpec {
	retC := C.g_param_spec_ref_sink((*C.GParamSpec)(recv.native))
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_param_spec_set_qdata_full : unsupported parameter destroy : no type generator for GLib.DestroyNotify, GDestroyNotify

// ParamSpecGType is a wrapper around the C record GParamSpecGType.
type ParamSpecGType struct {
	native *C.GParamSpecGType
	// parent_instance : record
	// is_a_type : no type generator for GType, GType
}

func ParamSpecGTypeNewFromC(u unsafe.Pointer) *ParamSpecGType {
	c := (*C.GParamSpecGType)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecGType{native: c}

	return g
}

func (recv *ParamSpecGType) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_type_module_add_interface : unsupported parameter instance_type : no type generator for GType, GType

// Unsupported : g_type_module_register_enum : no return generator

// Unsupported : g_type_module_register_flags : no return generator

// Unsupported : g_type_module_register_type : unsupported parameter parent_type : no type generator for GType, GType
