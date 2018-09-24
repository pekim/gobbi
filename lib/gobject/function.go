// This is a generated file - DO NOT EDIT

package gobject

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_boxed_copy : unsupported parameter boxed_type : no type generator for GType, GType

// Unsupported : g_boxed_free : unsupported parameter boxed_type : no type generator for GType, GType

// Unsupported : g_boxed_type_register_static : unsupported parameter boxed_copy : no type generator for BoxedCopyFunc, GBoxedCopyFunc

// Unsupported : g_cclosure_marshal_BOOLEAN__BOXED_BOXED : no return generator

// Unsupported : g_cclosure_marshal_BOOLEAN__FLAGS : no return generator

// Unsupported : g_cclosure_marshal_STRING__OBJECT_POINTER : no return generator

// Unsupported : g_cclosure_marshal_VOID__BOOLEAN : no return generator

// Unsupported : g_cclosure_marshal_VOID__BOXED : no return generator

// Unsupported : g_cclosure_marshal_VOID__CHAR : no return generator

// Unsupported : g_cclosure_marshal_VOID__DOUBLE : no return generator

// Unsupported : g_cclosure_marshal_VOID__ENUM : no return generator

// Unsupported : g_cclosure_marshal_VOID__FLAGS : no return generator

// Unsupported : g_cclosure_marshal_VOID__FLOAT : no return generator

// Unsupported : g_cclosure_marshal_VOID__INT : no return generator

// Unsupported : g_cclosure_marshal_VOID__LONG : no return generator

// Unsupported : g_cclosure_marshal_VOID__OBJECT : no return generator

// Unsupported : g_cclosure_marshal_VOID__PARAM : no return generator

// Unsupported : g_cclosure_marshal_VOID__POINTER : no return generator

// Unsupported : g_cclosure_marshal_VOID__STRING : no return generator

// Unsupported : g_cclosure_marshal_VOID__UCHAR : no return generator

// Unsupported : g_cclosure_marshal_VOID__UINT : no return generator

// Unsupported : g_cclosure_marshal_VOID__UINT_POINTER : no return generator

// Unsupported : g_cclosure_marshal_VOID__ULONG : no return generator

// Unsupported : g_cclosure_marshal_VOID__VARIANT : no return generator

// Unsupported : g_cclosure_marshal_VOID__VOID : no return generator

// Unsupported : g_cclosure_new : unsupported parameter callback_func : no type generator for Callback, GCallback

// Unsupported : g_cclosure_new_object : unsupported parameter callback_func : no type generator for Callback, GCallback

// Unsupported : g_cclosure_new_object_swap : unsupported parameter callback_func : no type generator for Callback, GCallback

// Unsupported : g_cclosure_new_swap : unsupported parameter callback_func : no type generator for Callback, GCallback

// Unsupported : g_enum_complete_type_info : unsupported parameter g_enum_type : no type generator for GType, GType

// EnumGetValue is a wrapper around the C function g_enum_get_value.
func EnumGetValue(enumClass *EnumClass, value int32) *EnumValue {
	c_enum_class := enumClass.toC()

	c_value := (C.gint)(value)

	retC := C.g_enum_get_value(c_enum_class, c_value)
	retGo := EnumValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EnumGetValueByName is a wrapper around the C function g_enum_get_value_by_name.
func EnumGetValueByName(enumClass *EnumClass, name string) *EnumValue {
	c_enum_class := enumClass.toC()

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_enum_get_value_by_name(c_enum_class, c_name)
	retGo := EnumValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EnumGetValueByNick is a wrapper around the C function g_enum_get_value_by_nick.
func EnumGetValueByNick(enumClass *EnumClass, nick string) *EnumValue {
	c_enum_class := enumClass.toC()

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	retC := C.g_enum_get_value_by_nick(c_enum_class, c_nick)
	retGo := EnumValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_enum_register_static : no return generator

// Unsupported : g_flags_complete_type_info : unsupported parameter g_flags_type : no type generator for GType, GType

// FlagsGetFirstValue is a wrapper around the C function g_flags_get_first_value.
func FlagsGetFirstValue(flagsClass *FlagsClass, value uint32) *FlagsValue {
	c_flags_class := flagsClass.toC()

	c_value := (C.guint)(value)

	retC := C.g_flags_get_first_value(c_flags_class, c_value)
	retGo := FlagsValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FlagsGetValueByName is a wrapper around the C function g_flags_get_value_by_name.
func FlagsGetValueByName(flagsClass *FlagsClass, name string) *FlagsValue {
	c_flags_class := flagsClass.toC()

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_flags_get_value_by_name(c_flags_class, c_name)
	retGo := FlagsValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FlagsGetValueByNick is a wrapper around the C function g_flags_get_value_by_nick.
func FlagsGetValueByNick(flagsClass *FlagsClass, nick string) *FlagsValue {
	c_flags_class := flagsClass.toC()

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	retC := C.g_flags_get_value_by_nick(c_flags_class, c_nick)
	retGo := FlagsValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_flags_register_static : no return generator

// Unsupported : g_gtype_get_type : no return generator

// Unsupported : g_param_spec_boolean : unsupported parameter default_value : no type generator for gboolean, gboolean

// Unsupported : g_param_spec_boxed : unsupported parameter boxed_type : no type generator for GType, GType

// Unsupported : g_param_spec_char : no return generator

// Unsupported : g_param_spec_double : no return generator

// Unsupported : g_param_spec_enum : unsupported parameter enum_type : no type generator for GType, GType

// Unsupported : g_param_spec_flags : unsupported parameter flags_type : no type generator for GType, GType

// Unsupported : g_param_spec_float : no return generator

// Unsupported : g_param_spec_int : no return generator

// Unsupported : g_param_spec_int64 : no return generator

// Unsupported : g_param_spec_long : no return generator

// Unsupported : g_param_spec_object : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_param_spec_param : unsupported parameter param_type : no type generator for GType, GType

// Unsupported : g_param_spec_pointer : no return generator

// Unsupported : g_param_spec_pool_new : unsupported parameter type_prefixing : no type generator for gboolean, gboolean

// Unsupported : g_param_spec_string : no return generator

// Unsupported : g_param_spec_uchar : no return generator

// Unsupported : g_param_spec_uint : no return generator

// Unsupported : g_param_spec_uint64 : no return generator

// Unsupported : g_param_spec_ulong : no return generator

// Unsupported : g_param_spec_unichar : no return generator

// Unsupported : g_param_spec_value_array : unsupported parameter element_spec : no type generator for ParamSpec, GParamSpec*

// Unsupported : g_param_type_register_static : no return generator

// Unsupported : g_param_value_convert : unsupported parameter pspec : no type generator for ParamSpec, GParamSpec*

// Unsupported : g_param_value_defaults : unsupported parameter pspec : no type generator for ParamSpec, GParamSpec*

// Unsupported : g_param_value_set_default : unsupported parameter pspec : no type generator for ParamSpec, GParamSpec*

// Unsupported : g_param_value_validate : unsupported parameter pspec : no type generator for ParamSpec, GParamSpec*

// Unsupported : g_param_values_cmp : unsupported parameter pspec : no type generator for ParamSpec, GParamSpec*

// Unsupported : g_pointer_type_register_static : no return generator

// Unsupported : g_signal_add_emission_hook : unsupported parameter hook_func : no type generator for SignalEmissionHook, GSignalEmissionHook

// Unsupported : g_signal_chain_from_overridden : unsupported parameter instance_and_params : no param type

// Unsupported : g_signal_connect_closure : unsupported parameter after : no type generator for gboolean, gboolean

// Unsupported : g_signal_connect_closure_by_id : unsupported parameter after : no type generator for gboolean, gboolean

// Unsupported : g_signal_connect_data : unsupported parameter c_handler : no type generator for Callback, GCallback

// Unsupported : g_signal_connect_object : unsupported parameter c_handler : no type generator for Callback, GCallback

// Unsupported : g_signal_emit : unsupported parameter ... : varargs

// Unsupported : g_signal_emit_by_name : unsupported parameter ... : varargs

// Unsupported : g_signal_emit_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : g_signal_emitv : unsupported parameter instance_and_params : no param type

// SignalGetInvocationHint is a wrapper around the C function g_signal_get_invocation_hint.
func SignalGetInvocationHint(instance uintptr) *SignalInvocationHint {
	c_instance := (C.gpointer)(instance)

	retC := C.g_signal_get_invocation_hint(c_instance)
	retGo := SignalInvocationHintNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_signal_handler_block : no return generator

// Unsupported : g_signal_handler_disconnect : no return generator

// SignalHandlerFind is a wrapper around the C function g_signal_handler_find.
func SignalHandlerFind(instance uintptr, mask SignalMatchType, signalId uint32, detail glib.Quark, closure *Closure, func_ uintptr, data uintptr) uint64 {
	c_instance := (C.gpointer)(instance)

	c_mask := (C.GSignalMatchType)(mask)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_closure := closure.toC()

	c_func := (C.gpointer)(func_)

	c_data := (C.gpointer)(data)

	retC := C.g_signal_handler_find(c_instance, c_mask, c_signal_id, c_detail, c_closure, c_func, c_data)
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported : g_signal_handler_is_connected : no return generator

// Unsupported : g_signal_handler_unblock : no return generator

// SignalHandlersBlockMatched is a wrapper around the C function g_signal_handlers_block_matched.
func SignalHandlersBlockMatched(instance uintptr, mask SignalMatchType, signalId uint32, detail glib.Quark, closure *Closure, func_ uintptr, data uintptr) uint32 {
	c_instance := (C.gpointer)(instance)

	c_mask := (C.GSignalMatchType)(mask)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_closure := closure.toC()

	c_func := (C.gpointer)(func_)

	c_data := (C.gpointer)(data)

	retC := C.g_signal_handlers_block_matched(c_instance, c_mask, c_signal_id, c_detail, c_closure, c_func, c_data)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_signal_handlers_destroy : no return generator

// SignalHandlersDisconnectMatched is a wrapper around the C function g_signal_handlers_disconnect_matched.
func SignalHandlersDisconnectMatched(instance uintptr, mask SignalMatchType, signalId uint32, detail glib.Quark, closure *Closure, func_ uintptr, data uintptr) uint32 {
	c_instance := (C.gpointer)(instance)

	c_mask := (C.GSignalMatchType)(mask)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_closure := closure.toC()

	c_func := (C.gpointer)(func_)

	c_data := (C.gpointer)(data)

	retC := C.g_signal_handlers_disconnect_matched(c_instance, c_mask, c_signal_id, c_detail, c_closure, c_func, c_data)
	retGo := (uint32)(retC)

	return retGo
}

// SignalHandlersUnblockMatched is a wrapper around the C function g_signal_handlers_unblock_matched.
func SignalHandlersUnblockMatched(instance uintptr, mask SignalMatchType, signalId uint32, detail glib.Quark, closure *Closure, func_ uintptr, data uintptr) uint32 {
	c_instance := (C.gpointer)(instance)

	c_mask := (C.GSignalMatchType)(mask)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_closure := closure.toC()

	c_func := (C.gpointer)(func_)

	c_data := (C.gpointer)(data)

	retC := C.g_signal_handlers_unblock_matched(c_instance, c_mask, c_signal_id, c_detail, c_closure, c_func, c_data)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_signal_has_handler_pending : unsupported parameter may_be_blocked : no type generator for gboolean, gboolean

// Unsupported : g_signal_list_ids : unsupported parameter itype : no type generator for GType, GType

// Unsupported : g_signal_lookup : unsupported parameter itype : no type generator for GType, GType

// SignalName is a wrapper around the C function g_signal_name.
func SignalName(signalId uint32) string {
	c_signal_id := (C.guint)(signalId)

	retC := C.g_signal_name(c_signal_id)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_signal_new : unsupported parameter itype : no type generator for GType, GType

// Unsupported : g_signal_new_valist : unsupported parameter itype : no type generator for GType, GType

// Unsupported : g_signal_newv : unsupported parameter itype : no type generator for GType, GType

// Unsupported : g_signal_override_class_closure : unsupported parameter instance_type : no type generator for GType, GType

// Unsupported : g_signal_parse_name : unsupported parameter itype : no type generator for GType, GType

// Unsupported : g_signal_query : no return generator

// Unsupported : g_signal_remove_emission_hook : no return generator

// Unsupported : g_signal_stop_emission : no return generator

// Unsupported : g_signal_stop_emission_by_name : no return generator

// Unsupported : g_signal_type_cclosure_new : unsupported parameter itype : no type generator for GType, GType

// Unsupported : g_source_set_closure : no return generator

// Unsupported : g_source_set_dummy_callback : no return generator

// StrdupValueContents is a wrapper around the C function g_strdup_value_contents.
func StrdupValueContents(value *Value) string {
	c_value := value.toC()

	retC := C.g_strdup_value_contents(c_value)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_type_add_class_cache_func : unsupported parameter cache_func : no type generator for TypeClassCacheFunc, GTypeClassCacheFunc

// Unsupported : g_type_add_instance_private : unsupported parameter class_type : no type generator for GType, GType

// Unsupported : g_type_add_interface_dynamic : unsupported parameter instance_type : no type generator for GType, GType

// Unsupported : g_type_add_interface_static : unsupported parameter instance_type : no type generator for GType, GType

// Unsupported : g_type_check_class_cast : unsupported parameter is_a_type : no type generator for GType, GType

// Unsupported : g_type_check_class_is_a : unsupported parameter is_a_type : no type generator for GType, GType

// Unsupported : g_type_check_instance : no return generator

// Unsupported : g_type_check_instance_cast : unsupported parameter iface_type : no type generator for GType, GType

// Unsupported : g_type_check_instance_is_a : unsupported parameter iface_type : no type generator for GType, GType

// Unsupported : g_type_check_instance_is_fundamentally_a : unsupported parameter fundamental_type : no type generator for GType, GType

// Unsupported : g_type_check_is_value_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_check_value : no return generator

// Unsupported : g_type_check_value_holds : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_children : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_class_adjust_private_offset : unsupported parameter private_size_or_offset : no type generator for gint, gint*

// Unsupported : g_type_class_peek : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_class_ref : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_create_instance : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_depth : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_free_instance : no return generator

// Unsupported : g_type_from_name : no return generator

// Unsupported : g_type_fundamental : unsupported parameter type_id : no type generator for GType, GType

// Unsupported : g_type_fundamental_next : no return generator

// Unsupported : g_type_get_plugin : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_get_qdata : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_init : no return generator

// Unsupported : g_type_init_with_debug_flags : no return generator

// Unsupported : g_type_interface_add_prerequisite : unsupported parameter interface_type : no type generator for GType, GType

// Unsupported : g_type_interface_get_plugin : unsupported parameter instance_type : no type generator for GType, GType

// Unsupported : g_type_interface_peek : unsupported parameter iface_type : no type generator for GType, GType

// Unsupported : g_type_interfaces : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_is_a : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_name : unsupported parameter type : no type generator for GType, GType

// TypeNameFromClass is a wrapper around the C function g_type_name_from_class.
func TypeNameFromClass(gClass *TypeClass) string {
	c_g_class := gClass.toC()

	retC := C.g_type_name_from_class(c_g_class)
	retGo := C.GoString(retC)

	return retGo
}

// TypeNameFromInstance is a wrapper around the C function g_type_name_from_instance.
func TypeNameFromInstance(instance *TypeInstance) string {
	c_instance := instance.toC()

	retC := C.g_type_name_from_instance(c_instance)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_type_next_base : unsupported parameter leaf_type : no type generator for GType, GType

// Unsupported : g_type_parent : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_qname : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_query : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_register_dynamic : unsupported parameter parent_type : no type generator for GType, GType

// Unsupported : g_type_register_fundamental : unsupported parameter type_id : no type generator for GType, GType

// Unsupported : g_type_register_static : unsupported parameter parent_type : no type generator for GType, GType

// Unsupported : g_type_remove_class_cache_func : unsupported parameter cache_func : no type generator for TypeClassCacheFunc, GTypeClassCacheFunc

// Unsupported : g_type_set_qdata : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_test_flags : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_value_table_peek : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_value_register_transform_func : unsupported parameter src_type : no type generator for GType, GType

// Unsupported : g_value_type_compatible : unsupported parameter src_type : no type generator for GType, GType

// Unsupported : g_value_type_transformable : unsupported parameter src_type : no type generator for GType, GType
