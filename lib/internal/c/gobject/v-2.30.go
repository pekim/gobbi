// Code generated - DO NOT EDIT.
// +build gobject_2.30

package gobject

import "unsafe"

// #include <glib-object.h>
// #include <stdlib.h>
/*

static void c_g_cclosure_marshal_generic_va(GClosure* closure, GValue* return_value, gpointer instance, gpointer marshal_data, int n_params, GType* param_types) {
    return g_cclosure_marshal_generic_va(closure, return_value, instance, NULL, marshal_data, n_params, param_types);
}
*/
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}
func toGoBool(b C.gboolean) bool {
	return b == C.TRUE
}
func Fn_g_cclosure_marshal_generic(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_generic(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_generic_va(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param4 unsafe.Pointer, param5 int, param6 []uint64) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.gpointer)(param2)

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.int)(param5)

	cValue6 := (*C.GType)(unsafe.Pointer(&param6[0]))

	C.c_g_cclosure_marshal_generic_va(cValue0, cValue1, cValue2, cValue4, cValue5, cValue6)
}

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
