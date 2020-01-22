// Code generated - DO NOT EDIT.
// +build gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54 gobject_2.62

package gobject

import "unsafe"

// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// UNSUPPORTED : g_cclosure_new : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object_swap : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_swap : parameter 'callback_func' is callback

// UNSUPPORTED : g_closure_add_finalize_notifier : parameter 'notify_func' is callback

// UNSUPPORTED : g_closure_add_invalidate_notifier : parameter 'notify_func' is callback

// UNSUPPORTED : g_closure_add_marshal_guards : parameter 'pre_marshal_notify' is callback

// UNSUPPORTED : g_closure_remove_finalize_notifier : parameter 'notify_func' is callback

// UNSUPPORTED : g_closure_remove_invalidate_notifier : parameter 'notify_func' is callback

// UNSUPPORTED : g_closure_set_marshal : parameter 'marshal' is callback

// UNSUPPORTED : g_closure_set_meta_marshal : parameter 'meta_marshal' is callback

func Fn_g_value_get_gtype(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GValue)(unsafe.Pointer(paramInstance))

	ret := C.g_value_get_gtype(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_value_set_gtype(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (*C.GValue)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	C.g_value_set_gtype(cValueInstance, cValue0)
}

// UNSUPPORTED : g_value_register_transform_func : parameter 'transform_func' is callback

// UNSUPPORTED : g_value_array_sort : parameter 'compare_func' is callback

// UNSUPPORTED : g_value_array_sort_with_data : parameter 'compare_func' is callback

// UNSUPPORTED : g_boxed_type_register_static : parameter 'boxed_copy' is callback

// UNSUPPORTED : g_cclosure_new : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object_swap : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_swap : parameter 'callback_func' is callback

// UNSUPPORTED : g_signal_add_emission_hook : parameter 'hook_func' is callback

// UNSUPPORTED : g_signal_chain_from_overridden : parameter 'instance_and_params' is array parameter without length parameter

// UNSUPPORTED : g_signal_connect_data : parameter 'c_handler' is callback

// UNSUPPORTED : g_signal_connect_object : parameter 'c_handler' is callback

// UNSUPPORTED : g_signal_emitv : parameter 'instance_and_params' is array parameter without length parameter

// UNSUPPORTED : g_signal_new : parameter 'accumulator' is callback

// UNSUPPORTED : g_signal_new_class_handler : parameter 'class_handler' is callback

// UNSUPPORTED : g_signal_new_valist : parameter 'accumulator' is callback

// UNSUPPORTED : g_signal_newv : parameter 'accumulator' is callback

// UNSUPPORTED : g_signal_override_class_handler : parameter 'class_handler' is callback

// UNSUPPORTED : g_signal_set_va_marshaller : blacklisted

// UNSUPPORTED : g_type_add_class_cache_func : parameter 'cache_func' is callback

// UNSUPPORTED : g_type_add_interface_check : parameter 'check_func' is callback

// UNSUPPORTED : g_type_register_static_simple : parameter 'class_init' is callback

// UNSUPPORTED : g_type_remove_class_cache_func : parameter 'cache_func' is callback

// UNSUPPORTED : g_type_remove_interface_check : parameter 'check_func' is callback

// UNSUPPORTED : g_value_register_transform_func : parameter 'transform_func' is callback

// UNSUPPORTED : g_object_add_toggle_ref : parameter 'notify' is callback

// UNSUPPORTED : g_object_bind_property_full : parameter 'transform_to' is callback

// UNSUPPORTED : g_object_dup_data : parameter 'dup_func' is callback

// UNSUPPORTED : g_object_dup_qdata : parameter 'dup_func' is callback

// UNSUPPORTED : g_object_remove_toggle_ref : parameter 'notify' is callback

// UNSUPPORTED : g_object_replace_data : parameter 'destroy' is callback

// UNSUPPORTED : g_object_replace_qdata : parameter 'destroy' is callback

// UNSUPPORTED : g_object_set_data_full : parameter 'destroy' is callback

// UNSUPPORTED : g_object_set_qdata_full : parameter 'destroy' is callback

// UNSUPPORTED : g_object_weak_ref : parameter 'notify' is callback

// UNSUPPORTED : g_object_weak_unref : parameter 'notify' is callback

// UNSUPPORTED : g_param_spec_set_qdata_full : parameter 'destroy' is callback
