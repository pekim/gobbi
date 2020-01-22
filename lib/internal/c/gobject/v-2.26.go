// Code generated - DO NOT EDIT.
// +build gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54 gobject_2.62

package gobject

import "unsafe"

// #include <glib-object.h>
// #include <stdlib.h>
import "C"

func Fn_g_cclosure_marshal_VOID__VARIANT(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_VOID__VARIANT(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
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

func Fn_g_object_class_install_properties(paramInstance unsafe.Pointer, param0 uint, param1 []unsafe.Pointer) {
	cValueInstance := (*C.GObjectClass)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (**C.GParamSpec)(unsafe.Pointer(&param1[0]))

	C.g_object_class_install_properties(cValueInstance, cValue0, cValue1)
}

func Fn_g_value_dup_variant(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GValue)(unsafe.Pointer(paramInstance))

	ret := C.g_value_dup_variant(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_value_get_variant(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GValue)(unsafe.Pointer(paramInstance))

	ret := C.g_value_get_variant(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_value_set_variant(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GValue)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	C.g_value_set_variant(cValueInstance, cValue0)
}

func Fn_g_value_take_variant(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GValue)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	C.g_value_take_variant(cValueInstance, cValue0)
}

// UNSUPPORTED : g_value_register_transform_func : parameter 'transform_func' is callback

// UNSUPPORTED : g_value_array_sort : parameter 'compare_func' is callback

// UNSUPPORTED : g_value_array_sort_with_data : parameter 'compare_func' is callback

// UNSUPPORTED : g_boxed_type_register_static : parameter 'boxed_copy' is callback

// UNSUPPORTED : g_cclosure_new : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object_swap : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_swap : parameter 'callback_func' is callback

func Fn_g_param_spec_variant(param0 string, param1 string, param2 string, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.GVariantType)(unsafe.Pointer(param3))

	cValue4 := (*C.GVariant)(unsafe.Pointer(param4))

	cValue5 := (C.GParamFlags)(param5)

	ret := C.g_param_spec_variant(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return unsafe.Pointer(ret)
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

// UNSUPPORTED : g_type_register_static_simple : parameter 'class_init' is callback

// UNSUPPORTED : g_type_remove_class_cache_func : parameter 'cache_func' is callback

// UNSUPPORTED : g_type_remove_interface_check : parameter 'check_func' is callback

// UNSUPPORTED : g_value_register_transform_func : parameter 'transform_func' is callback

func Fn_g_binding_get_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GBinding)(unsafe.Pointer(paramInstance))

	ret := C.g_binding_get_flags(cValueInstance)

	return (int)(ret)
}

func Fn_g_binding_get_source(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GBinding)(unsafe.Pointer(paramInstance))

	ret := C.g_binding_get_source(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_binding_get_source_property(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GBinding)(unsafe.Pointer(paramInstance))

	ret := C.g_binding_get_source_property(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_binding_get_target(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GBinding)(unsafe.Pointer(paramInstance))

	ret := C.g_binding_get_target(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_binding_get_target_property(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GBinding)(unsafe.Pointer(paramInstance))

	ret := C.g_binding_get_target_property(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : g_object_add_toggle_ref : parameter 'notify' is callback

func Fn_g_object_bind_property(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 string, param3 int) unsafe.Pointer {
	cValueInstance := (C.gpointer)(paramInstance)

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gpointer)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.GBindingFlags)(param3)

	ret := C.g_object_bind_property(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_object_bind_property_full : parameter 'transform_to' is callback

func Fn_g_object_bind_property_with_closures(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 string, param3 int, param4 unsafe.Pointer, param5 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (C.gpointer)(paramInstance)

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gpointer)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.GBindingFlags)(param3)

	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))

	cValue5 := (*C.GClosure)(unsafe.Pointer(param5))

	ret := C.g_object_bind_property_with_closures(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_object_dup_data : parameter 'dup_func' is callback

// UNSUPPORTED : g_object_dup_qdata : parameter 'dup_func' is callback

func Fn_g_object_notify_by_pspec(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	C.g_object_notify_by_pspec(cValueInstance, cValue0)
}

// UNSUPPORTED : g_object_remove_toggle_ref : parameter 'notify' is callback

// UNSUPPORTED : g_object_replace_data : parameter 'destroy' is callback

// UNSUPPORTED : g_object_replace_qdata : parameter 'destroy' is callback

// UNSUPPORTED : g_object_set_data_full : parameter 'destroy' is callback

// UNSUPPORTED : g_object_set_qdata_full : parameter 'destroy' is callback

// UNSUPPORTED : g_object_weak_ref : parameter 'notify' is callback

// UNSUPPORTED : g_object_weak_unref : parameter 'notify' is callback

// UNSUPPORTED : g_param_spec_set_qdata_full : parameter 'destroy' is callback
