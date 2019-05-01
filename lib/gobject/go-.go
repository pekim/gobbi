// This is a generated file - DO NOT EDIT

package gobject

import call "github.com/pekim/gobbi/lib/internal/call"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Unsupported : alias has no type generator for ClosureMarshal, GClosureMarshal

// Unsupported : alias has no type generator for VaClosureMarshal, GVaClosureMarshal

// Type is a representation of the C alias GType.
type Type uint64

type ConnectFlags int

const (
	CONNECT_AFTER   ConnectFlags = 1
	CONNECT_SWAPPED ConnectFlags = 2
)

type ParamFlags int

const (
	PARAM_READABLE        ParamFlags = 1
	PARAM_WRITABLE        ParamFlags = 2
	PARAM_READWRITE       ParamFlags = 3
	PARAM_CONSTRUCT       ParamFlags = 4
	PARAM_CONSTRUCT_ONLY  ParamFlags = 8
	PARAM_LAX_VALIDATION  ParamFlags = 16
	PARAM_STATIC_NAME     ParamFlags = 32
	PARAM_PRIVATE         ParamFlags = 32
	PARAM_STATIC_NICK     ParamFlags = 64
	PARAM_STATIC_BLURB    ParamFlags = 128
	PARAM_EXPLICIT_NOTIFY ParamFlags = 1073741824

// Blacklisted member : G_PARAM_DEPRECATED
)

type SignalFlags int

const (
	SIGNAL_RUN_FIRST    SignalFlags = 1
	SIGNAL_RUN_LAST     SignalFlags = 2
	SIGNAL_RUN_CLEANUP  SignalFlags = 4
	SIGNAL_NO_RECURSE   SignalFlags = 8
	SIGNAL_DETAILED     SignalFlags = 16
	SIGNAL_ACTION       SignalFlags = 32
	SIGNAL_NO_HOOKS     SignalFlags = 64
	SIGNAL_MUST_COLLECT SignalFlags = 128
	SIGNAL_DEPRECATED   SignalFlags = 256
)

type SignalMatchType int

const (
	SIGNAL_MATCH_ID        SignalMatchType = 1
	SIGNAL_MATCH_DETAIL    SignalMatchType = 2
	SIGNAL_MATCH_CLOSURE   SignalMatchType = 4
	SIGNAL_MATCH_FUNC      SignalMatchType = 8
	SIGNAL_MATCH_DATA      SignalMatchType = 16
	SIGNAL_MATCH_UNBLOCKED SignalMatchType = 32
)

type TypeDebugFlags int

const (
	TYPE_DEBUG_NONE           TypeDebugFlags = 0
	TYPE_DEBUG_OBJECTS        TypeDebugFlags = 1
	TYPE_DEBUG_SIGNALS        TypeDebugFlags = 2
	TYPE_DEBUG_INSTANCE_COUNT TypeDebugFlags = 4
	TYPE_DEBUG_MASK           TypeDebugFlags = 7
)

type TypeFlags int

const (
	TYPE_FLAG_ABSTRACT       TypeFlags = 16
	TYPE_FLAG_VALUE_ABSTRACT TypeFlags = 32
)

type TypeFundamentalFlags int

const (
	TYPE_FLAG_CLASSED        TypeFundamentalFlags = 1
	TYPE_FLAG_INSTANTIATABLE TypeFundamentalFlags = 2
	TYPE_FLAG_DERIVABLE      TypeFundamentalFlags = 4
	TYPE_FLAG_DEEP_DERIVABLE TypeFundamentalFlags = 8
)

// Unsupported : g_object_new : unsupported parameter ... : varargs

// Unsupported : g_object_new_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Unsupported : g_object_newv : unsupported parameter parameters :

// Blacklisted : GTypeModule

const PARAM_MASK int32 = 255
const PARAM_STATIC_STRINGS int32 = 224
const PARAM_USER_SHIFT int32 = 8
const SIGNAL_FLAGS_MASK int32 = 511
const SIGNAL_MATCH_MASK int32 = 63

// Unsupported : type GLib.Type for TYPE_FLAG_RESERVED_ID_BIT

const TYPE_FUNDAMENTAL_MAX int32 = 255
const TYPE_FUNDAMENTAL_SHIFT int32 = 2
const TYPE_RESERVED_BSE_FIRST int32 = 32
const TYPE_RESERVED_BSE_LAST int32 = 48
const TYPE_RESERVED_GLIB_FIRST int32 = 22
const TYPE_RESERVED_GLIB_LAST int32 = 31
const TYPE_RESERVED_USER_FIRST int32 = 49

// Blacklisted : VALUE_COLLECT_FORMAT_MAX_LENGTH

const VALUE_NOCOPY_CONTENTS int32 = 134217728

// Unsupported : g_boxed_copy : unsupported parameter src_boxed : no type generator for gpointer (gconstpointer) for param src_boxed

// Unsupported : g_boxed_free : unsupported parameter boxed : no type generator for gpointer (gpointer) for param boxed

// Unsupported : g_boxed_type_register_static : unsupported parameter boxed_copy : no type generator for BoxedCopyFunc (GBoxedCopyFunc) for param boxed_copy

// Unsupported : g_cclosure_marshal_BOOLEAN__BOXED_BOXED : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_BOOLEAN__FLAGS : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_STRING__OBJECT_POINTER : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__BOOLEAN : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__BOXED : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__CHAR : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__DOUBLE : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__ENUM : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__FLAGS : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__FLOAT : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__INT : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__LONG : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__OBJECT : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__PARAM : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__POINTER : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__STRING : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__UCHAR : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__UINT : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__UINT_POINTER : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__ULONG : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__VARIANT : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__VOID : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_new : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_object : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_object_swap : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_swap : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_enum_get_value : return type :

// Unsupported : g_enum_get_value_by_name : return type :

// Unsupported : g_enum_get_value_by_nick : return type :

// Unsupported : g_enum_register_static : return type :

// Unsupported : g_flags_get_first_value : return type :

// Unsupported : g_flags_get_value_by_name : return type :

// Unsupported : g_flags_get_value_by_nick : return type :

// Unsupported : g_flags_register_static : return type :

// Unsupported : g_gtype_get_type : return type :

// Unsupported : g_param_spec_boolean : return type :

// Unsupported : g_param_spec_boxed : return type :

// Unsupported : g_param_spec_char : return type :

// Unsupported : g_param_spec_double : return type :

// Unsupported : g_param_spec_enum : return type :

// Unsupported : g_param_spec_flags : return type :

// Unsupported : g_param_spec_float : return type :

// Unsupported : g_param_spec_int : return type :

// Unsupported : g_param_spec_int64 : return type :

// Unsupported : g_param_spec_long : return type :

// Unsupported : g_param_spec_object : return type :

// Unsupported : g_param_spec_param : return type :

// Unsupported : g_param_spec_pointer : return type :

// Unsupported : g_param_spec_pool_new : return type :

// Unsupported : g_param_spec_string : return type :

// Unsupported : g_param_spec_uchar : return type :

// Unsupported : g_param_spec_uint : return type :

// Unsupported : g_param_spec_uint64 : return type :

// Unsupported : g_param_spec_ulong : return type :

// Unsupported : g_param_spec_unichar : return type :

// Unsupported : g_param_spec_value_array : return type :

// Unsupported : g_param_type_register_static : return type :

// Unsupported : g_param_value_convert : return type :

// Unsupported : g_param_value_defaults : return type :

// Unsupported : g_param_value_validate : return type :

// Unsupported : g_pointer_type_register_static : return type :

// Unsupported : g_signal_add_emission_hook : unsupported parameter hook_func : no type generator for SignalEmissionHook (GSignalEmissionHook) for param hook_func

// Unsupported : g_signal_chain_from_overridden : unsupported parameter instance_and_params :

// Unsupported : g_signal_connect_data : unsupported parameter c_handler : no type generator for Callback (GCallback) for param c_handler

// Unsupported : g_signal_connect_object : unsupported parameter c_handler : no type generator for Callback (GCallback) for param c_handler

// Unsupported : g_signal_emit : unsupported parameter ... : varargs

// Unsupported : g_signal_emit_by_name : unsupported parameter ... : varargs

// Unsupported : g_signal_emit_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Unsupported : g_signal_emitv : unsupported parameter instance_and_params :

// Unsupported : g_signal_get_invocation_hint : return type :

// Unsupported : g_signal_handler_find : unsupported parameter func : no type generator for gpointer (gpointer) for param func

// Unsupported : g_signal_handler_is_connected : return type :

// Unsupported : g_signal_handlers_block_matched : unsupported parameter func : no type generator for gpointer (gpointer) for param func

// Unsupported : g_signal_handlers_disconnect_matched : unsupported parameter func : no type generator for gpointer (gpointer) for param func

// Unsupported : g_signal_handlers_unblock_matched : unsupported parameter func : no type generator for gpointer (gpointer) for param func

// Unsupported : g_signal_has_handler_pending : return type :

// Unsupported : g_signal_list_ids : array return type :

// Unsupported : g_signal_name : return type :

// Unsupported : g_signal_new : unsupported parameter accumulator : no type generator for SignalAccumulator (GSignalAccumulator) for param accumulator

// Unsupported : g_signal_new_valist : unsupported parameter accumulator : no type generator for SignalAccumulator (GSignalAccumulator) for param accumulator

// Unsupported : g_signal_newv : unsupported parameter accumulator : no type generator for SignalAccumulator (GSignalAccumulator) for param accumulator

// Unsupported : g_signal_parse_name : return type :

// Unsupported : g_signal_type_cclosure_new : return type :

// Unsupported : g_strdup_value_contents : return type :

// Unsupported : g_type_add_class_cache_func : unsupported parameter cache_data : no type generator for gpointer (gpointer) for param cache_data

// Unsupported : g_type_check_class_cast : return type :

// Unsupported : g_type_check_class_is_a : return type :

// Unsupported : g_type_check_instance : return type :

// Unsupported : g_type_check_instance_cast : return type :

// Unsupported : g_type_check_instance_is_a : return type :

// Unsupported : g_type_check_instance_is_fundamentally_a : return type :

// Unsupported : g_type_check_is_value_type : return type :

// Unsupported : g_type_check_value : return type :

// Unsupported : g_type_check_value_holds : return type :

// Unsupported : g_type_children : array return type :

// Unsupported : g_type_class_adjust_private_offset : unsupported parameter g_class : no type generator for gpointer (gpointer) for param g_class

// Unsupported : g_type_class_peek : return type :

// Unsupported : g_type_class_ref : return type :

// Unsupported : g_type_create_instance : return type :

// Unsupported : g_type_from_name : return type :

// Unsupported : g_type_fundamental : return type :

// Unsupported : g_type_fundamental_next : return type :

// Unsupported : g_type_get_plugin : return type :

// Unsupported : g_type_get_qdata : no return generator

// TypeInit is a wrapper around the C function g_type_init.
func TypeInit() {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_VOID},
	}
	call.Function(3573, &data)
	return
}

// Unsupported : g_type_interface_get_plugin : return type :

// Unsupported : g_type_interface_peek : return type :

// Unsupported : g_type_interfaces : array return type :

// Unsupported : g_type_is_a : return type :

// Unsupported : g_type_name : return type :

// Unsupported : g_type_name_from_class : return type :

// Unsupported : g_type_name_from_instance : return type :

// Unsupported : g_type_next_base : return type :

// Unsupported : g_type_parent : return type :

// Unsupported : g_type_qname : return type :

// Unsupported : g_type_register_dynamic : return type :

// Unsupported : g_type_register_fundamental : return type :

// Unsupported : g_type_register_static : return type :

// Unsupported : g_type_remove_class_cache_func : unsupported parameter cache_data : no type generator for gpointer (gpointer) for param cache_data

// Unsupported : g_type_set_qdata : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_type_test_flags : return type :

// Unsupported : g_type_value_table_peek : return type :

// Unsupported : g_value_register_transform_func : unsupported parameter transform_func : no type generator for ValueTransform (GValueTransform) for param transform_func

// Unsupported : g_value_type_compatible : return type :

// Unsupported : g_value_type_transformable : return type :

// Unsupported : g_closure_new_object : return type :

// Unsupported : g_closure_new_simple : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_value_array_new : return type :
