// Code generated - DO NOT EDIT.
// +build gobject_2.4

package gobject

import "unsafe"

// #include <glib-object.h>
import "C"

type CClosure C.GCClosure
type Closure C.GClosure
type ClosureNotifyData C.GClosureNotifyData
type EnumClass C.GEnumClass
type EnumValue C.GEnumValue
type FlagsClass C.GFlagsClass
type FlagsValue C.GFlagsValue
type InitiallyUnownedClass C.GInitiallyUnownedClass
type InterfaceInfo C.GInterfaceInfo
type ObjectClass C.GObjectClass
type ObjectConstructParam C.GObjectConstructParam
type ParamSpecClass C.GParamSpecClass
type ParamSpecPool C.GParamSpecPool
type ParamSpecTypeInfo C.GParamSpecTypeInfo
type Parameter C.GParameter
type SignalInvocationHint C.GSignalInvocationHint
type SignalQuery C.GSignalQuery
type TypeClass C.GTypeClass
type TypeFundamentalInfo C.GTypeFundamentalInfo
type TypeInfo C.GTypeInfo
type TypeInstance C.GTypeInstance
type TypeInterface C.GTypeInterface
type TypeModuleClass C.GTypeModuleClass
type TypePluginClass C.GTypePluginClass
type TypeQuery C.GTypeQuery
type TypeValueTable C.GTypeValueTable
type Value C.GValue
type ValueArray C.GValueArray
type WeakRef C.GWeakRef

func Fn_g_boxed_copy(param0 uint64, param1 unsafe.Pointer) {}

func Fn_g_boxed_free(param0 uint64, param1 unsafe.Pointer) {}

// UNSUPPORTED : boxed_type_register_static : has callback

func Fn_g_cclosure_marshal_BOOLEAN__BOXED_BOXED(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_cclosure_marshal_BOOLEAN__FLAGS(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_cclosure_marshal_STRING__OBJECT_POINTER(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_cclosure_marshal_VOID__BOOLEAN(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_cclosure_marshal_VOID__BOXED(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_cclosure_marshal_VOID__CHAR(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_cclosure_marshal_VOID__DOUBLE(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_cclosure_marshal_VOID__ENUM(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_cclosure_marshal_VOID__FLAGS(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_cclosure_marshal_VOID__FLOAT(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_cclosure_marshal_VOID__INT(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_cclosure_marshal_VOID__LONG(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_cclosure_marshal_VOID__OBJECT(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_cclosure_marshal_VOID__PARAM(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_cclosure_marshal_VOID__POINTER(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_cclosure_marshal_VOID__STRING(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_cclosure_marshal_VOID__UCHAR(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_cclosure_marshal_VOID__UINT(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_cclosure_marshal_VOID__UINT_POINTER(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_cclosure_marshal_VOID__ULONG(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_cclosure_marshal_VOID__VOID(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

// UNSUPPORTED : cclosure_new : has callback

// UNSUPPORTED : cclosure_new_object : has callback

// UNSUPPORTED : cclosure_new_object_swap : has callback

// UNSUPPORTED : cclosure_new_swap : has callback

func Fn_g_enum_complete_type_info(param0 uint64, param1 unsafe.Pointer, param2 unsafe.Pointer) {}

func Fn_g_enum_get_value(param0 unsafe.Pointer, param1 int) {}

func Fn_g_enum_get_value_by_name(param0 unsafe.Pointer, param1 string) {}

func Fn_g_enum_get_value_by_nick(param0 unsafe.Pointer, param1 string) {}

func Fn_g_enum_register_static(param0 string, param1 unsafe.Pointer) {}

func Fn_g_flags_complete_type_info(param0 uint64, param1 unsafe.Pointer, param2 unsafe.Pointer) {}

func Fn_g_flags_get_first_value(param0 unsafe.Pointer, param1 uint) {}

func Fn_g_flags_get_value_by_name(param0 unsafe.Pointer, param1 string) {}

func Fn_g_flags_get_value_by_nick(param0 unsafe.Pointer, param1 string) {}

func Fn_g_flags_register_static(param0 string, param1 unsafe.Pointer) {}

func Fn_g_gtype_get_type() {
	C.g_gtype_get_type()
}

func Fn_g_param_spec_boolean(param0 string, param1 string, param2 string, param3 bool, param4 int) {}

func Fn_g_param_spec_boxed(param0 string, param1 string, param2 string, param3 uint64, param4 int) {}

func Fn_g_param_spec_char(param0 string, param1 string, param2 string, param3 int8, param4 int8, param5 int8, param6 int) {
}

func Fn_g_param_spec_double(param0 string, param1 string, param2 string, param3 float64, param4 float64, param5 float64, param6 int) {
}

func Fn_g_param_spec_enum(param0 string, param1 string, param2 string, param3 uint64, param4 int, param5 int) {
}

func Fn_g_param_spec_flags(param0 string, param1 string, param2 string, param3 uint64, param4 uint, param5 int) {
}

func Fn_g_param_spec_float(param0 string, param1 string, param2 string, param3 float32, param4 float32, param5 float32, param6 int) {
}

func Fn_g_param_spec_int(param0 string, param1 string, param2 string, param3 int, param4 int, param5 int, param6 int) {
}

func Fn_g_param_spec_int64(param0 string, param1 string, param2 string, param3 int64, param4 int64, param5 int64, param6 int) {
}

func Fn_g_param_spec_long(param0 string, param1 string, param2 string, param3 int64, param4 int64, param5 int64, param6 int) {
}

func Fn_g_param_spec_object(param0 string, param1 string, param2 string, param3 uint64, param4 int) {}

func Fn_g_param_spec_override(param0 string, param1 unsafe.Pointer) {}

func Fn_g_param_spec_param(param0 string, param1 string, param2 string, param3 uint64, param4 int) {}

func Fn_g_param_spec_pointer(param0 string, param1 string, param2 string, param3 int) {}

func Fn_g_param_spec_pool_new(param0 bool) {}

func Fn_g_param_spec_string(param0 string, param1 string, param2 string, param3 string, param4 int) {}

func Fn_g_param_spec_uchar(param0 string, param1 string, param2 string, param3 uint8, param4 uint8, param5 uint8, param6 int) {
}

func Fn_g_param_spec_uint(param0 string, param1 string, param2 string, param3 uint, param4 uint, param5 uint, param6 int) {
}

func Fn_g_param_spec_uint64(param0 string, param1 string, param2 string, param3 uint64, param4 uint64, param5 uint64, param6 int) {
}

func Fn_g_param_spec_ulong(param0 string, param1 string, param2 string, param3 uint64, param4 uint64, param5 uint64, param6 int) {
}

func Fn_g_param_spec_unichar(param0 string, param1 string, param2 string, param3 rune, param4 int) {}

func Fn_g_param_spec_value_array(param0 string, param1 string, param2 string, param3 unsafe.Pointer, param4 int) {
}

func Fn_g_param_type_register_static(param0 string, param1 unsafe.Pointer) {}

func Fn_g_param_value_convert(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 bool) {
}

func Fn_g_param_value_defaults(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_param_value_set_default(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_param_value_validate(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_param_values_cmp(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {}

func Fn_g_pointer_type_register_static(param0 string) {}

func Fn_g_signal_accumulator_true_handled(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer) {
}

// UNSUPPORTED : signal_add_emission_hook : has callback

func Fn_g_signal_chain_from_overridden(param0 []Value, param1 unsafe.Pointer) {}

// UNSUPPORTED : signal_chain_from_overridden_handler : has varargs

func Fn_g_signal_connect_closure(param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer, param3 bool) {
}

func Fn_g_signal_connect_closure_by_id(param0 unsafe.Pointer, param1 uint, param2 uint32, param3 unsafe.Pointer, param4 bool) {
}

// UNSUPPORTED : signal_connect_data : has callback

// UNSUPPORTED : signal_connect_object : has callback

// UNSUPPORTED : signal_emit : has varargs

// UNSUPPORTED : signal_emit_by_name : has varargs

// UNSUPPORTED : signal_emit_valist : has va_list

func Fn_g_signal_emitv(param0 []Value, param1 uint, param2 uint32, param3 unsafe.Pointer) {}

func Fn_g_signal_get_invocation_hint(param0 unsafe.Pointer) {}

func Fn_g_signal_handler_block(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_signal_handler_disconnect(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_signal_handler_find(param0 unsafe.Pointer, param1 int, param2 uint, param3 uint32, param4 unsafe.Pointer, param5 unsafe.Pointer, param6 unsafe.Pointer) {
}

func Fn_g_signal_handler_is_connected(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_signal_handler_unblock(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_signal_handlers_block_matched(param0 unsafe.Pointer, param1 int, param2 uint, param3 uint32, param4 unsafe.Pointer, param5 unsafe.Pointer, param6 unsafe.Pointer) {
}

func Fn_g_signal_handlers_destroy(param0 unsafe.Pointer) {}

func Fn_g_signal_handlers_disconnect_matched(param0 unsafe.Pointer, param1 int, param2 uint, param3 uint32, param4 unsafe.Pointer, param5 unsafe.Pointer, param6 unsafe.Pointer) {
}

func Fn_g_signal_handlers_unblock_matched(param0 unsafe.Pointer, param1 int, param2 uint, param3 uint32, param4 unsafe.Pointer, param5 unsafe.Pointer, param6 unsafe.Pointer) {
}

func Fn_g_signal_has_handler_pending(param0 unsafe.Pointer, param1 uint, param2 uint32, param3 bool) {
}

func Fn_g_signal_list_ids(param0 uint64, param1 *uint) {}

func Fn_g_signal_lookup(param0 string, param1 uint64) {}

func Fn_g_signal_name(param0 uint) {}

// UNSUPPORTED : signal_new : has varargs

// UNSUPPORTED : signal_new_class_handler : has varargs

// UNSUPPORTED : signal_new_valist : has va_list

// UNSUPPORTED : signal_newv : has callback

func Fn_g_signal_override_class_closure(param0 uint, param1 uint64, param2 unsafe.Pointer) {}

// UNSUPPORTED : signal_override_class_handler : has callback

func Fn_g_signal_parse_name(param0 string, param1 uint64, param2 *uint, param3 uint32, param4 bool) {}

func Fn_g_signal_query(param0 uint, param1 unsafe.Pointer) {}

func Fn_g_signal_remove_emission_hook(param0 uint, param1 uint64) {}

func Fn_g_signal_stop_emission(param0 unsafe.Pointer, param1 uint, param2 uint32) {}

func Fn_g_signal_stop_emission_by_name(param0 unsafe.Pointer, param1 string) {}

func Fn_g_signal_type_cclosure_new(param0 uint64, param1 uint) {}

func Fn_g_source_set_closure(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_source_set_dummy_callback(param0 unsafe.Pointer) {}

func Fn_g_strdup_value_contents(param0 unsafe.Pointer) {}

// UNSUPPORTED : type_add_class_cache_func : has callback

func Fn_g_type_add_instance_private(param0 uint64, param1 uint64) {}

// UNSUPPORTED : type_add_interface_check : has callback

func Fn_g_type_add_interface_dynamic(param0 uint64, param1 uint64, param2 unsafe.Pointer) {}

func Fn_g_type_add_interface_static(param0 uint64, param1 uint64, param2 unsafe.Pointer) {}

func Fn_g_type_check_class_cast(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_type_check_class_is_a(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_type_check_instance(param0 unsafe.Pointer) {}

func Fn_g_type_check_instance_cast(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_type_check_instance_is_a(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_type_check_instance_is_fundamentally_a(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_type_check_is_value_type(param0 uint64) {}

func Fn_g_type_check_value(param0 unsafe.Pointer) {}

func Fn_g_type_check_value_holds(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_type_children(param0 uint64, param1 *uint) {}

func Fn_g_type_class_adjust_private_offset(param0 unsafe.Pointer, param1 *int) {}

func Fn_g_type_class_peek(param0 uint64) {}

func Fn_g_type_class_peek_static(param0 uint64) {}

func Fn_g_type_class_ref(param0 uint64) {}

func Fn_g_type_create_instance(param0 uint64) {}

func Fn_g_type_default_interface_peek(param0 uint64) {}

func Fn_g_type_default_interface_ref(param0 uint64) {}

func Fn_g_type_default_interface_unref(param0 unsafe.Pointer) {}

func Fn_g_type_depth(param0 uint64) {}

func Fn_g_type_free_instance(param0 unsafe.Pointer) {}

func Fn_g_type_from_name(param0 string) {}

func Fn_g_type_fundamental(param0 uint64) {}

func Fn_g_type_fundamental_next() {
	C.g_type_fundamental_next()
}

func Fn_g_type_get_plugin(param0 uint64) {}

func Fn_g_type_get_qdata(param0 uint64, param1 uint32) {}

func Fn_g_type_init() {
	C.g_type_init()
}

func Fn_g_type_init_with_debug_flags(param0 int) {}

func Fn_g_type_interface_add_prerequisite(param0 uint64, param1 uint64) {}

func Fn_g_type_interface_get_plugin(param0 uint64, param1 uint64) {}

func Fn_g_type_interface_peek(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_type_interface_prerequisites(param0 uint64, param1 *uint) {}

func Fn_g_type_interfaces(param0 uint64, param1 *uint) {}

func Fn_g_type_is_a(param0 uint64, param1 uint64) {}

func Fn_g_type_name(param0 uint64) {}

func Fn_g_type_name_from_class(param0 unsafe.Pointer) {}

func Fn_g_type_name_from_instance(param0 unsafe.Pointer) {}

func Fn_g_type_next_base(param0 uint64, param1 uint64) {}

func Fn_g_type_parent(param0 uint64) {}

func Fn_g_type_qname(param0 uint64) {}

func Fn_g_type_query(param0 uint64, param1 unsafe.Pointer) {}

func Fn_g_type_register_dynamic(param0 uint64, param1 string, param2 unsafe.Pointer, param3 int) {}

func Fn_g_type_register_fundamental(param0 uint64, param1 string, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int) {
}

func Fn_g_type_register_static(param0 uint64, param1 string, param2 unsafe.Pointer, param3 int) {}

// UNSUPPORTED : type_register_static_simple : has callback

// UNSUPPORTED : type_remove_class_cache_func : has callback

// UNSUPPORTED : type_remove_interface_check : has callback

func Fn_g_type_set_qdata(param0 uint64, param1 uint32, param2 unsafe.Pointer) {}

func Fn_g_type_test_flags(param0 uint64, param1 uint) {}

func Fn_g_type_value_table_peek(param0 uint64) {}

// UNSUPPORTED : value_register_transform_func : has callback

func Fn_g_value_type_compatible(param0 uint64, param1 uint64) {}

func Fn_g_value_type_transformable(param0 uint64, param1 uint64) {}

// UNSUPPORTED : new : has varargs

// UNSUPPORTED : new_valist : has va_list

func Fn_g_object_newv(param0 uint64, param1 uint, param2 []Parameter) {}
