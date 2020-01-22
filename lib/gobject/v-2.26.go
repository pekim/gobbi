// Code generated - DO NOT EDIT.
// +build gobject_2.26

package gobject

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/internal/c/gobject"
	"unsafe"
)

// UNSUPPORTED : SignalCMarshaller : blacklisted

// UNSUPPORTED : SignalCVaMarshaller : blacklisted

// BindingFlags is a representation of the C bitfield GBindingFlags.
type BindingFlags int

// BindingFlags_default is a representation of the C bitfield member G_BINDING_DEFAULT.
const BindingFlags_default = BindingFlags(0)

// BindingFlags_bidirectional is a representation of the C bitfield member G_BINDING_BIDIRECTIONAL.
const BindingFlags_bidirectional = BindingFlags(1)

// BindingFlags_sync_create is a representation of the C bitfield member G_BINDING_SYNC_CREATE.
const BindingFlags_sync_create = BindingFlags(2)

// BindingFlags_invert_boolean is a representation of the C bitfield member G_BINDING_INVERT_BOOLEAN.
const BindingFlags_invert_boolean = BindingFlags(4)

// UNSUPPORTED : g_boxed_type_register_static : parameter 'boxed_copy' is callback

// UNSUPPORTED : g_cclosure_new : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object_swap : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_swap : parameter 'callback_func' is callback

// UNSUPPORTED : g_enum_complete_type_info : has [in]out param, info

// UNSUPPORTED : g_flags_complete_type_info : has [in]out param, info

// ParamSpecVariant_ wraps the C function g_param_spec_variant.
//
// since 2.26
func ParamSpecVariant_(name string, nick string, blurb string, type_ *glib.VariantType, defaultValue *glib.Variant, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_type_ := type_.ToC()
	sys_defaultValue := defaultValue.ToC()
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_variant(sys_name, sys_nick, sys_blurb, sys_type_, sys_defaultValue, sys_flags)
	ret := ParamSpecNewFromC(retSys)

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

// UNSUPPORTED : g_type_interface_prerequisites : has [in]out param, n_prerequisites

// UNSUPPORTED : g_type_interfaces : has [in]out param, n_interfaces

// UNSUPPORTED : g_type_query : has [in]out param, query

// UNSUPPORTED : g_type_register_static_simple : parameter 'class_init' is callback

// UNSUPPORTED : g_type_remove_class_cache_func : parameter 'cache_func' is callback

// UNSUPPORTED : g_type_remove_interface_check : parameter 'check_func' is callback

// UNSUPPORTED : g_value_register_transform_func : parameter 'transform_func' is callback

// Binding is a representation of the C record GBinding.
//
// since 2.26
type Binding struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBinding that represents the Binding.
func (recv *Binding) ToC() unsafe.Pointer {
	return recv.native
}

// BindingNewFromC creates a new Binding from a pointer to the C GBinding that represents the Binding.
func BindingNewFromC(native unsafe.Pointer) *Binding {
	return &Binding{native: native}
}

// ParamSpecVariant is a representation of the C record GParamSpecVariant.
//
// since 2.26
type ParamSpecVariant struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecVariant that represents the ParamSpecVariant.
func (recv *ParamSpecVariant) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecVariantNewFromC creates a new ParamSpecVariant from a pointer to the C GParamSpecVariant that represents the ParamSpecVariant.
func ParamSpecVariantNewFromC(native unsafe.Pointer) *ParamSpecVariant {
	return &ParamSpecVariant{native: native}
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
