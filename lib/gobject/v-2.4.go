// Code generated - DO NOT EDIT.
// +build gobject_2.4

package gobject

import (
	gobject "github.com/pekim/gobbi/lib/internal/c/gobject"
	"unsafe"
)

// UNSUPPORTED : SignalCMarshaller : blacklisted

// UNSUPPORTED : SignalCVaMarshaller : blacklisted

// UNSUPPORTED : g_boxed_type_register_static : parameter 'boxed_copy' is callback

// UNSUPPORTED : g_cclosure_new : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object_swap : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_swap : parameter 'callback_func' is callback

// UNSUPPORTED : g_enum_complete_type_info : has [in]out param, info

// UNSUPPORTED : g_flags_complete_type_info : has [in]out param, info

// ParamSpecOverride_ wraps the C function g_param_spec_override.
//
// since 2.4
func ParamSpecOverride_(name string, overridden *ParamSpec) *ParamSpec {
	sys_name := name
	sys_overridden := overridden.ToC()
	retSys := gobject.Fn_g_param_spec_override(sys_name, sys_overridden)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// SignalAccumulatorTrueHandled wraps the C function g_signal_accumulator_true_handled.
//
// since 2.4
func SignalAccumulatorTrueHandled(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value, dummy unsafe.Pointer) bool {
	sys_ihint := ihint.ToC()
	sys_returnAccu := returnAccu.ToC()
	sys_handlerReturn := handlerReturn.ToC()
	sys_dummy := dummy
	retSys := gobject.Fn_g_signal_accumulator_true_handled(sys_ihint, sys_returnAccu, sys_handlerReturn, sys_dummy)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_signal_add_emission_hook : parameter 'hook_func' is callback

// UNSUPPORTED : g_signal_chain_from_overridden : parameter 'instance_and_params' is array parameter without length parameter

// UNSUPPORTED : g_signal_connect_data : parameter 'c_handler' is callback

// UNSUPPORTED : g_signal_connect_object : parameter 'c_handler' is callback

// UNSUPPORTED : g_signal_emitv : parameter 'instance_and_params' is array parameter without length parameter

// UNSUPPORTED : g_signal_list_ids : has [in]out param, n_ids

// UNSUPPORTED : g_signal_new : parameter 'accumulator' is callback

// UNSUPPORTED : g_signal_new_class_handler : parameter 'class_handler' is callback

// UNSUPPORTED : g_signal_new_valist : parameter 'accumulator' is callback

// UNSUPPORTED : g_signal_newv : parameter 'accumulator' is callback

// UNSUPPORTED : g_signal_override_class_handler : parameter 'class_handler' is callback

// UNSUPPORTED : g_signal_parse_name : has [in]out param, signal_id_p

// UNSUPPORTED : g_signal_query : has [in]out param, query

// UNSUPPORTED : g_signal_set_va_marshaller : blacklisted

// UNSUPPORTED : g_type_add_class_cache_func : parameter 'cache_func' is callback

// UNSUPPORTED : g_type_add_interface_check : parameter 'check_func' is callback

// UNSUPPORTED : g_type_children : has [in]out param, n_children

// TypeDefaultInterfacePeek wraps the C function g_type_default_interface_peek.
//
// since 2.4
func TypeDefaultInterfacePeek(gType uint64) *TypeInterface {
	sys_gType := gType
	retSys := gobject.Fn_g_type_default_interface_peek(sys_gType)
	ret := TypeInterfaceNewFromC(retSys)

	return ret
}

// TypeDefaultInterfaceRef wraps the C function g_type_default_interface_ref.
//
// since 2.4
func TypeDefaultInterfaceRef(gType uint64) *TypeInterface {
	sys_gType := gType
	retSys := gobject.Fn_g_type_default_interface_ref(sys_gType)
	ret := TypeInterfaceNewFromC(retSys)

	return ret
}

// TypeDefaultInterfaceUnref wraps the C function g_type_default_interface_unref.
//
// since 2.4
func TypeDefaultInterfaceUnref(gIface unsafe.Pointer) {
	sys_gIface := gIface
	gobject.Fn_g_type_default_interface_unref(sys_gIface)
}

// UNSUPPORTED : g_type_interface_prerequisites : has [in]out param, n_prerequisites

// UNSUPPORTED : g_type_interfaces : has [in]out param, n_interfaces

// UNSUPPORTED : g_type_query : has [in]out param, query

// UNSUPPORTED : g_type_register_static_simple : parameter 'class_init' is callback

// UNSUPPORTED : g_type_remove_class_cache_func : parameter 'cache_func' is callback

// UNSUPPORTED : g_type_remove_interface_check : parameter 'check_func' is callback

// UNSUPPORTED : g_value_register_transform_func : parameter 'transform_func' is callback

// ParamSpecOverride is a representation of the C record GParamSpecOverride.
//
// since 2.4
type ParamSpecOverride struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecOverride that represents the ParamSpecOverride.
func (recv *ParamSpecOverride) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecOverrideNewFromC creates a new ParamSpecOverride from a pointer to the C GParamSpecOverride that represents the ParamSpecOverride.
func ParamSpecOverrideNewFromC(native unsafe.Pointer) *ParamSpecOverride {
	return &ParamSpecOverride{native: native}
}

// TypeCValue is a representation of the C union GTypeCValue.
type TypeCValue struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeCValue that represents the TypeCValue.
func (recv *TypeCValue) ToC() unsafe.Pointer {
	return recv.native
}

// TypeCValueNewFromC creates a new TypeCValue from a pointer to the C GTypeCValue that represents the TypeCValue.
func TypeCValueNewFromC(native unsafe.Pointer) *TypeCValue {
	return &TypeCValue{native: native}
}

// UNSUPPORTED : _Value__data__union : blacklisted
