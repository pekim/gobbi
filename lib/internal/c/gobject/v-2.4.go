// Code generated - DO NOT EDIT.
// +build gobject_2.4

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

func Fn_g_object_class_override_property(paramInstance unsafe.Pointer, param0 uint, param1 string) {
	cValueInstance := (*C.GObjectClass)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_object_class_override_property(cValueInstance, cValue0, cValue1)
}

func Fn_g_type_class_add_private(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (C.gpointer)(paramInstance)

	cValue0 := (C.gsize)(param0)

	C.g_type_class_add_private(cValueInstance, cValue0)
}

func Fn_g_type_class_peek_static(param0 uint64) unsafe.Pointer {
	cValue0 := (C.GType)(param0)

	ret := C.g_type_class_peek_static(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_value_take_boxed(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GValue)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gconstpointer)(param0)

	C.g_value_take_boxed(cValueInstance, cValue0)
}

func Fn_g_value_take_object(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GValue)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gpointer)(param0)

	C.g_value_take_object(cValueInstance, cValue0)
}

func Fn_g_value_take_param(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GValue)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	C.g_value_take_param(cValueInstance, cValue0)
}

func Fn_g_value_take_string(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GValue)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_value_take_string(cValueInstance, cValue0)
}

// UNSUPPORTED : g_value_register_transform_func : parameter 'transform_func' is callback

// UNSUPPORTED : g_value_array_sort : parameter 'compare_func' is callback

// UNSUPPORTED : g_value_array_sort_with_data : parameter 'compare_func' is callback

// UNSUPPORTED : g_boxed_type_register_static : parameter 'boxed_copy' is callback

// UNSUPPORTED : g_cclosure_new : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object_swap : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_swap : parameter 'callback_func' is callback

func Fn_g_param_spec_override(param0 string, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GParamSpec)(unsafe.Pointer(param1))

	ret := C.g_param_spec_override(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_signal_accumulator_true_handled(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer) bool {
	cValue0 := (*C.GSignalInvocationHint)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	cValue3 := (C.gpointer)(param3)

	ret := C.g_signal_accumulator_true_handled(cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

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

func Fn_g_type_default_interface_peek(param0 uint64) unsafe.Pointer {
	cValue0 := (C.GType)(param0)

	ret := C.g_type_default_interface_peek(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_type_default_interface_ref(param0 uint64) unsafe.Pointer {
	cValue0 := (C.GType)(param0)

	ret := C.g_type_default_interface_ref(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_type_default_interface_unref(param0 unsafe.Pointer) {
	cValue0 := (C.gpointer)(param0)

	C.g_type_default_interface_unref(cValue0)
}

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

func Fn_g_object_interface_find_property(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_object_interface_find_property(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_object_interface_install_property(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (*C.GParamSpec)(unsafe.Pointer(param1))

	C.g_object_interface_install_property(cValue0, cValue1)
}

func Fn_g_object_interface_list_properties(param0 unsafe.Pointer, param1 *uint) []unsafe.Pointer {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (*C.guint)(unsafe.Pointer(param1))

	ret := C.g_object_interface_list_properties(cValue0, cValue1)

	retLen := int(*cValue1)
	retGo := make([]unsafe.Pointer, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

func Fn_g_param_spec_get_redirect_target(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))

	ret := C.g_param_spec_get_redirect_target(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_param_spec_set_qdata_full : parameter 'destroy' is callback
