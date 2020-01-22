// Code generated - DO NOT EDIT.
// +build gobject_2.38

package gobject

import "unsafe"

// #include <glib-object.h>
// #include <stdlib.h>
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

func Fn_g_type_class_get_instance_private_offset(paramInstance unsafe.Pointer) int {
	cValueInstance := (C.gpointer)(paramInstance)

	ret := C.g_type_class_get_instance_private_offset(cValueInstance)

	return (int)(ret)
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

func Fn_g_binding_unbind(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GBinding)(unsafe.Pointer(paramInstance))

	C.g_binding_unbind(cValueInstance)
}

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

func Fn_g_param_spec_get_default_value(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))

	ret := C.g_param_spec_get_default_value(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_param_spec_set_qdata_full : parameter 'destroy' is callback
