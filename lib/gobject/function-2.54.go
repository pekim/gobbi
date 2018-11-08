// This is a generated file - DO NOT EDIT
// +build gobject_2.54

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_boxed_type_register_static : unsupported parameter boxed_copy : no type generator for BoxedCopyFunc (GBoxedCopyFunc) for param boxed_copy

// Unsupported : g_cclosure_new : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_object : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_object_swap : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_swap : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_clear_object : unsupported parameter object_ptr : record with indirection level of 2

// EnumToString is a wrapper around the C function g_enum_to_string.
func EnumToString(gEnumType Type, value int32) string {
	c_g_enum_type := (C.GType)(gEnumType)

	c_value := (C.gint)(value)

	retC := C.g_enum_to_string(c_g_enum_type, c_value)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// FlagsToString is a wrapper around the C function g_flags_to_string.
func FlagsToString(flagsType Type, value uint32) string {
	c_flags_type := (C.GType)(flagsType)

	c_value := (C.guint)(value)

	retC := C.g_flags_to_string(c_flags_type, c_value)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_param_spec_boolean : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_boxed : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_char : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_double : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_enum : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_flags : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_float : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_gtype : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_int : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_int64 : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_long : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_object : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_override : unsupported parameter overridden : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_param : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_pointer : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_string : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_uchar : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_uint : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_uint64 : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_ulong : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_unichar : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_value_array : unsupported parameter element_spec : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_variant : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_param_value_convert : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : g_param_value_defaults : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : g_param_value_set_default : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : g_param_value_validate : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : g_param_values_cmp : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : g_signal_add_emission_hook : unsupported parameter hook_func : no type generator for SignalEmissionHook (GSignalEmissionHook) for param hook_func

// Unsupported : g_signal_chain_from_overridden : unsupported parameter instance_and_params :

// Unsupported : g_signal_chain_from_overridden_handler : unsupported parameter ... : varargs

// Unsupported : g_signal_connect_data : unsupported parameter c_handler : no type generator for Callback (GCallback) for param c_handler

// Unsupported : g_signal_connect_object : unsupported parameter c_handler : no type generator for Callback (GCallback) for param c_handler

// Unsupported : g_signal_emit : unsupported parameter ... : varargs

// Unsupported : g_signal_emit_by_name : unsupported parameter ... : varargs

// Unsupported : g_signal_emit_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Unsupported : g_signal_emitv : unsupported parameter instance_and_params :

// Unsupported : g_signal_list_ids : no return type

// Unsupported : g_signal_new : unsupported parameter accumulator : no type generator for SignalAccumulator (GSignalAccumulator) for param accumulator

// Unsupported : g_signal_new_class_handler : unsupported parameter class_handler : no type generator for Callback (GCallback) for param class_handler

// Unsupported : g_signal_new_valist : unsupported parameter accumulator : no type generator for SignalAccumulator (GSignalAccumulator) for param accumulator

// Unsupported : g_signal_newv : unsupported parameter accumulator : no type generator for SignalAccumulator (GSignalAccumulator) for param accumulator

// Unsupported : g_signal_override_class_handler : unsupported parameter class_handler : no type generator for Callback (GCallback) for param class_handler

// Unsupported : g_signal_set_va_marshaller : unsupported parameter va_marshaller : no type generator for SignalCVaMarshaller (GSignalCVaMarshaller) for param va_marshaller

// Unsupported : g_type_add_class_cache_func : unsupported parameter cache_func : no type generator for TypeClassCacheFunc (GTypeClassCacheFunc) for param cache_func

// Unsupported : g_type_add_interface_check : unsupported parameter check_func : no type generator for TypeInterfaceCheckFunc (GTypeInterfaceCheckFunc) for param check_func

// Unsupported : g_type_add_interface_dynamic : unsupported parameter plugin : no type generator for TypePlugin (GTypePlugin*) for param plugin

// Unsupported : g_type_children : no return type

// Unsupported : g_type_get_plugin : no return generator

// Unsupported : g_type_interface_get_plugin : no return generator

// Unsupported : g_type_interface_prerequisites : no return type

// Unsupported : g_type_interfaces : no return type

// Unsupported : g_type_register_dynamic : unsupported parameter plugin : no type generator for TypePlugin (GTypePlugin*) for param plugin

// Unsupported : g_type_register_static_simple : unsupported parameter class_init : no type generator for ClassInitFunc (GClassInitFunc) for param class_init

// Unsupported : g_type_remove_class_cache_func : unsupported parameter cache_func : no type generator for TypeClassCacheFunc (GTypeClassCacheFunc) for param cache_func

// Unsupported : g_type_remove_interface_check : unsupported parameter check_func : no type generator for TypeInterfaceCheckFunc (GTypeInterfaceCheckFunc) for param check_func

// Unsupported : g_value_register_transform_func : unsupported parameter transform_func : no type generator for ValueTransform (GValueTransform) for param transform_func
