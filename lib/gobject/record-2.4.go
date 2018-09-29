// This is a generated file - DO NOT EDIT
// +build gobject_2.4 gobject_2.6 gobject_2.8 gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import "unsafe"

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

// OverrideProperty is a wrapper around the C function g_object_class_override_property.
func (recv *ObjectClass) OverrideProperty(propertyId uint32, name string) {
	c_property_id := (C.guint)(propertyId)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.g_object_class_override_property((*C.GObjectClass)(recv.native), c_property_id, c_name)

	return
}

// Unsupported : g_param_spec_pool_insert : unsupported parameter owner_type : no type generator for GType, GType

// Unsupported : g_param_spec_pool_list : unsupported parameter owner_type : no type generator for GType, GType

// Unsupported : g_param_spec_pool_list_owned : unsupported parameter owner_type : no type generator for GType, GType

// Unsupported : g_param_spec_pool_lookup : unsupported parameter owner_type : no type generator for GType, GType

// AddPrivate is a wrapper around the C function g_type_class_add_private.
func (recv *TypeClass) AddPrivate(privateSize uint64) {
	c_private_size := (C.gsize)(privateSize)

	C.g_type_class_add_private((C.gpointer)(recv.native), c_private_size)

	return
}

// Unsupported : g_type_class_get_private : unsupported parameter private_type : no type generator for GType, GType

// Unsupported : g_type_instance_get_private : unsupported parameter private_type : no type generator for GType, GType

// Unsupported : g_value_dup_variant : return type : Blacklisted record : GVariant

// Unsupported : g_value_get_gtype : no return generator

// Unsupported : g_value_get_variant : return type : Blacklisted record : GVariant

// Unsupported : g_value_init : unsupported parameter g_type : no type generator for GType, GType

// Unsupported : g_value_set_gtype : unsupported parameter v_gtype : no type generator for GType, GType

// Unsupported : g_value_set_variant : unsupported parameter variant : Blacklisted record : GVariant

// TakeBoxed is a wrapper around the C function g_value_take_boxed.
func (recv *Value) TakeBoxed(vBoxed uintptr) {
	c_v_boxed := (C.gconstpointer)(vBoxed)

	C.g_value_take_boxed((*C.GValue)(recv.native), c_v_boxed)

	return
}

// TakeObject is a wrapper around the C function g_value_take_object.
func (recv *Value) TakeObject(vObject uintptr) {
	c_v_object := (C.gpointer)(vObject)

	C.g_value_take_object((*C.GValue)(recv.native), c_v_object)

	return
}

// TakeParam is a wrapper around the C function g_value_take_param.
func (recv *Value) TakeParam(param *ParamSpec) {
	c_param := (*C.GParamSpec)(param.ToC())

	C.g_value_take_param((*C.GValue)(recv.native), c_param)

	return
}

// TakeString is a wrapper around the C function g_value_take_string.
func (recv *Value) TakeString(vString string) {
	c_v_string := C.CString(vString)
	defer C.free(unsafe.Pointer(c_v_string))

	C.g_value_take_string((*C.GValue)(recv.native), c_v_string)

	return
}

// Unsupported : g_value_take_variant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : g_value_array_sort : unsupported parameter compare_func : no type generator for GLib.CompareFunc, GCompareFunc

// Unsupported : g_value_array_sort_with_data : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc, GCompareDataFunc
