// This is a generated file - DO NOT EDIT
// +build gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_closure_add_finalize_notifier : unsupported parameter notify_func : no type generator for ClosureNotify, GClosureNotify

// Unsupported : g_closure_add_invalidate_notifier : unsupported parameter notify_func : no type generator for ClosureNotify, GClosureNotify

// Unsupported : g_closure_add_marshal_guards : unsupported parameter pre_marshal_notify : no type generator for ClosureNotify, GClosureNotify

// Unsupported : g_closure_invoke : unsupported parameter param_values : no param type

// Unsupported : g_closure_remove_finalize_notifier : unsupported parameter notify_func : no type generator for ClosureNotify, GClosureNotify

// Unsupported : g_closure_remove_invalidate_notifier : unsupported parameter notify_func : no type generator for ClosureNotify, GClosureNotify

// Unsupported : g_closure_set_marshal : unsupported parameter marshal : no type generator for ClosureMarshal, GClosureMarshal

// Unsupported : g_closure_set_meta_marshal : unsupported parameter meta_marshal : no type generator for ClosureMarshal, GClosureMarshal

// Unsupported : g_object_class_install_properties : unsupported parameter pspecs : no param type

// Unsupported : g_object_class_list_properties : unsupported parameter n_properties : no type generator for guint, guint*

// Unsupported : g_param_spec_pool_insert : unsupported parameter owner_type : no type generator for GType, GType

// Unsupported : g_param_spec_pool_list : unsupported parameter owner_type : no type generator for GType, GType

// Unsupported : g_param_spec_pool_list_owned : unsupported parameter owner_type : no type generator for GType, GType

// Unsupported : g_param_spec_pool_lookup : unsupported parameter owner_type : no type generator for GType, GType

// Unsupported : g_type_class_get_private : unsupported parameter private_type : no type generator for GType, GType

// Unsupported : g_type_instance_get_private : unsupported parameter private_type : no type generator for GType, GType

// Unsupported : g_value_dup_variant : return type : Blacklisted record : GVariant

// Unsupported : g_value_get_gtype : no return generator

// GetSchar is a wrapper around the C function g_value_get_schar.
func (recv *Value) GetSchar() int8 {
	retC := C.g_value_get_schar((*C.GValue)(recv.native))
	retGo := (int8)(retC)

	return retGo
}

// Unsupported : g_value_get_variant : return type : Blacklisted record : GVariant

// Unsupported : g_value_init : unsupported parameter g_type : no type generator for GType, GType

// Unsupported : g_value_set_gtype : unsupported parameter v_gtype : no type generator for GType, GType

// SetSchar is a wrapper around the C function g_value_set_schar.
func (recv *Value) SetSchar(vChar int8) {
	c_v_char := (C.gint8)(vChar)

	C.g_value_set_schar((*C.GValue)(recv.native), c_v_char)

	return
}

// Unsupported : g_value_set_variant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : g_value_take_variant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : g_value_array_sort : unsupported parameter compare_func : no type generator for GLib.CompareFunc, GCompareFunc

// Unsupported : g_value_array_sort_with_data : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc, GCompareDataFunc

// Clear is a wrapper around the C function g_weak_ref_clear.
func (recv *WeakRef) Clear() {
	C.g_weak_ref_clear((*C.GWeakRef)(recv.native))

	return
}

// Get is a wrapper around the C function g_weak_ref_get.
func (recv *WeakRef) Get() uintptr {
	retC := C.g_weak_ref_get((*C.GWeakRef)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// Init is a wrapper around the C function g_weak_ref_init.
func (recv *WeakRef) Init(object uintptr) {
	c_object := (C.gpointer)(object)

	C.g_weak_ref_init((*C.GWeakRef)(recv.native), c_object)

	return
}

// Set is a wrapper around the C function g_weak_ref_set.
func (recv *WeakRef) Set(object uintptr) {
	c_object := (C.gpointer)(object)

	C.g_weak_ref_set((*C.GWeakRef)(recv.native), c_object)

	return
}
