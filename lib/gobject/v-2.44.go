// Code generated - DO NOT EDIT.
// +build gobject_2.44 gobject_2.46 gobject_2.54 gobject_2.62

package gobject

import gobject "github.com/pekim/gobbi/lib/internal/c/gobject"

// UNSUPPORTED : SignalCMarshaller : blacklisted

// UNSUPPORTED : SignalCVaMarshaller : blacklisted

// UNSUPPORTED : g_boxed_type_register_static : parameter 'boxed_copy' is callback

// UNSUPPORTED : g_cclosure_new : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object_swap : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_swap : parameter 'callback_func' is callback

// UNSUPPORTED : g_enum_complete_type_info : has [in]out param, info

// UNSUPPORTED : g_flags_complete_type_info : has [in]out param, info

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

// TypeGetInstanceCount wraps the C function g_type_get_instance_count.
//
// since 2.44
func TypeGetInstanceCount(type_ uint64) int {
	sys_type_ := type_
	retSys := gobject.Fn_g_type_get_instance_count(sys_type_)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_type_interface_prerequisites : has [in]out param, n_prerequisites

// UNSUPPORTED : g_type_interfaces : has [in]out param, n_interfaces

// UNSUPPORTED : g_type_query : has [in]out param, query

// UNSUPPORTED : g_type_register_static_simple : parameter 'class_init' is callback

// UNSUPPORTED : g_type_remove_class_cache_func : parameter 'cache_func' is callback

// UNSUPPORTED : g_type_remove_interface_check : parameter 'check_func' is callback

// UNSUPPORTED : g_value_register_transform_func : parameter 'transform_func' is callback
