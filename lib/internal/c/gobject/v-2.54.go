// Code generated - DO NOT EDIT.
// +build gobject_2.54 gobject_2.62

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

// UNSUPPORTED : g_value_register_transform_func : parameter 'transform_func' is callback

// UNSUPPORTED : g_value_array_sort : parameter 'compare_func' is callback

// UNSUPPORTED : g_value_array_sort_with_data : parameter 'compare_func' is callback

// UNSUPPORTED : g_boxed_type_register_static : parameter 'boxed_copy' is callback

// UNSUPPORTED : g_cclosure_new : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object_swap : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_swap : parameter 'callback_func' is callback

func Fn_g_enum_to_string(param0 uint64, param1 int) string {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.g_enum_to_string(cValue0, cValue1)

	return C.GoString(ret)
}

func Fn_g_flags_to_string(param0 uint64, param1 uint) string {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.guint)(param1)

	ret := C.g_flags_to_string(cValue0, cValue1)

	return C.GoString(ret)
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

func Fn_g_object_new_with_properties(param0 uint64, param1 uint, param2 []string, param3 []Value) unsafe.Pointer {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.guint)(param1)

	param2Len := len(param2)
	cValue2Array := C.malloc((C.ulong)(param2Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue2Array))
	param2Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue2Array))[:param2Len:param2Len]
	for param2i, param2String := range param2 {
		param2Slice[param2i] = (*C.gchar)(C.CString(param2String))
		defer C.free(unsafe.Pointer(param2Slice[param2i]))
	}
	cValue2 := &param2Slice[0]

	cValue3 := (*C.GValue)(unsafe.Pointer(&param3[0]))

	ret := C.g_object_new_with_properties(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_object_add_toggle_ref : parameter 'notify' is callback

// UNSUPPORTED : g_object_bind_property_full : parameter 'transform_to' is callback

// UNSUPPORTED : g_object_dup_data : parameter 'dup_func' is callback

// UNSUPPORTED : g_object_dup_qdata : parameter 'dup_func' is callback

func Fn_g_object_getv(paramInstance unsafe.Pointer, param0 uint, param1 []string, param2 []Value) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	param1Len := len(param1)
	cValue1Array := C.malloc((C.ulong)(param1Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue1Array))
	param1Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue1Array))[:param1Len:param1Len]
	for param1i, param1String := range param1 {
		param1Slice[param1i] = (*C.gchar)(C.CString(param1String))
		defer C.free(unsafe.Pointer(param1Slice[param1i]))
	}
	cValue1 := &param1Slice[0]

	cValue2 := (*C.GValue)(unsafe.Pointer(&param2[0]))

	C.g_object_getv(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : g_object_remove_toggle_ref : parameter 'notify' is callback

// UNSUPPORTED : g_object_replace_data : parameter 'destroy' is callback

// UNSUPPORTED : g_object_replace_qdata : parameter 'destroy' is callback

// UNSUPPORTED : g_object_set_data_full : parameter 'destroy' is callback

// UNSUPPORTED : g_object_set_qdata_full : parameter 'destroy' is callback

func Fn_g_object_setv(paramInstance unsafe.Pointer, param0 uint, param1 []string, param2 []Value) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	param1Len := len(param1)
	cValue1Array := C.malloc((C.ulong)(param1Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue1Array))
	param1Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue1Array))[:param1Len:param1Len]
	for param1i, param1String := range param1 {
		param1Slice[param1i] = (*C.gchar)(C.CString(param1String))
		defer C.free(unsafe.Pointer(param1Slice[param1i]))
	}
	cValue1 := &param1Slice[0]

	cValue2 := (*C.GValue)(unsafe.Pointer(&param2[0]))

	C.g_object_setv(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : g_object_weak_ref : parameter 'notify' is callback

// UNSUPPORTED : g_object_weak_unref : parameter 'notify' is callback

// UNSUPPORTED : g_param_spec_set_qdata_full : parameter 'destroy' is callback
