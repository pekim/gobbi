// Code generated - DO NOT EDIT.
// +build gobject_2.4

package gobject

import "unsafe"

// #include <glib-object.h>
import "C"

// records
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

func Fn_boxed_copy(boxedType uint64, srcBoxed unsafe.Pointer) {}

func Fn_boxed_free(boxedType uint64, boxed unsafe.Pointer) {}

// UNSUPPORTED : boxed_type_register_static : has callback

func Fn_cclosure_marshal_BOOLEAN__BOXED_BOXED(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_BOOLEAN__FLAGS(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_STRING__OBJECT_POINTER(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__BOOLEAN(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__BOXED(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__CHAR(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__DOUBLE(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__ENUM(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__FLAGS(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__FLOAT(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__INT(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__LONG(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__OBJECT(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__PARAM(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__POINTER(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__STRING(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__UCHAR(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__UINT(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__UINT_POINTER(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__ULONG(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__VOID(closure unsafe.Pointer, returnValue unsafe.Pointer, nParamValues uint, paramValues unsafe.Pointer, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

// UNSUPPORTED : cclosure_new : has callback

// UNSUPPORTED : cclosure_new_object : has callback

// UNSUPPORTED : cclosure_new_object_swap : has callback

// UNSUPPORTED : cclosure_new_swap : has callback

func Fn_enum_complete_type_info(gEnumType uint64, constValues unsafe.Pointer) {}

func Fn_enum_get_value(enumClass unsafe.Pointer, value int) {}

func Fn_enum_get_value_by_name(enumClass unsafe.Pointer, name string) {}

func Fn_enum_get_value_by_nick(enumClass unsafe.Pointer, nick string) {}

func Fn_enum_register_static(name string, constStaticValues unsafe.Pointer) {}

func Fn_flags_complete_type_info(gFlagsType uint64, constValues unsafe.Pointer) {}

func Fn_flags_get_first_value(flagsClass unsafe.Pointer, value uint) {}

func Fn_flags_get_value_by_name(flagsClass unsafe.Pointer, name string) {}

func Fn_flags_get_value_by_nick(flagsClass unsafe.Pointer, nick string) {}

func Fn_flags_register_static(name string, constStaticValues unsafe.Pointer) {}

func Fn_gtype_get_type() {
	C.g_gtype_get_type()
}

func Fn_param_spec_boolean(name string, nick string, blurb string, defaultValue bool, flags int) {}

func Fn_param_spec_boxed(name string, nick string, blurb string, boxedType uint64, flags int) {}

func Fn_param_spec_char(name string, nick string, blurb string, minimum int8, maximum int8, defaultValue int8, flags int) {
}

func Fn_param_spec_double(name string, nick string, blurb string, minimum float64, maximum float64, defaultValue float64, flags int) {
}

func Fn_param_spec_enum(name string, nick string, blurb string, enumType uint64, defaultValue int, flags int) {
}

func Fn_param_spec_flags(name string, nick string, blurb string, flagsType uint64, defaultValue uint, flags int) {
}

func Fn_param_spec_float(name string, nick string, blurb string, minimum float32, maximum float32, defaultValue float32, flags int) {
}

func Fn_param_spec_int(name string, nick string, blurb string, minimum int, maximum int, defaultValue int, flags int) {
}

func Fn_param_spec_int64(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags int) {
}

func Fn_param_spec_long(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags int) {
}

func Fn_param_spec_object(name string, nick string, blurb string, objectType uint64, flags int) {}

func Fn_param_spec_override(name string, overridden unsafe.Pointer) {}

func Fn_param_spec_param(name string, nick string, blurb string, paramType uint64, flags int) {}

func Fn_param_spec_pointer(name string, nick string, blurb string, flags int) {}

func Fn_param_spec_pool_new(typePrefixing bool) {}

func Fn_param_spec_string(name string, nick string, blurb string, defaultValue string, flags int) {}

func Fn_param_spec_uchar(name string, nick string, blurb string, minimum uint8, maximum uint8, defaultValue uint8, flags int) {
}

func Fn_param_spec_uint(name string, nick string, blurb string, minimum uint, maximum uint, defaultValue uint, flags int) {
}

func Fn_param_spec_uint64(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags int) {
}

func Fn_param_spec_ulong(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags int) {
}

func Fn_param_spec_unichar(name string, nick string, blurb string, defaultValue rune, flags int) {}

func Fn_param_spec_value_array(name string, nick string, blurb string, elementSpec unsafe.Pointer, flags int) {
}

func Fn_param_type_register_static(name string, pspecInfo unsafe.Pointer) {}

func Fn_param_value_convert(pspec unsafe.Pointer, srcValue unsafe.Pointer, destValue unsafe.Pointer, strictValidation bool) {
}

func Fn_param_value_defaults(pspec unsafe.Pointer, value unsafe.Pointer) {}

func Fn_param_value_set_default(pspec unsafe.Pointer, value unsafe.Pointer) {}

func Fn_param_value_validate(pspec unsafe.Pointer, value unsafe.Pointer) {}

func Fn_param_values_cmp(pspec unsafe.Pointer, value1 unsafe.Pointer, value2 unsafe.Pointer) {}

func Fn_pointer_type_register_static(name string) {}

func Fn_signal_accumulator_true_handled(ihint unsafe.Pointer, returnAccu unsafe.Pointer, handlerReturn unsafe.Pointer, dummy unsafe.Pointer) {
}

// UNSUPPORTED : signal_add_emission_hook : has callback

func Fn_signal_chain_from_overridden(instanceAndParams *Value, returnValue unsafe.Pointer) {}

// UNSUPPORTED : signal_chain_from_overridden_handler : has varargs

func Fn_signal_connect_closure(instance unsafe.Pointer, detailedSignal string, closure unsafe.Pointer, after bool) {
}

func Fn_signal_connect_closure_by_id(instance unsafe.Pointer, signalId uint, detail uint32, closure unsafe.Pointer, after bool) {
}

// UNSUPPORTED : signal_connect_data : has callback

// UNSUPPORTED : signal_connect_object : has callback

// UNSUPPORTED : signal_emit : has varargs

// UNSUPPORTED : signal_emit_by_name : has varargs

// UNSUPPORTED : signal_emit_valist : has va_list

func Fn_signal_emitv(instanceAndParams *Value, signalId uint, detail uint32, returnValue unsafe.Pointer) {
}

func Fn_signal_get_invocation_hint(instance unsafe.Pointer) {}

func Fn_signal_handler_block(instance unsafe.Pointer, handlerId uint64) {}

func Fn_signal_handler_disconnect(instance unsafe.Pointer, handlerId uint64) {}

func Fn_signal_handler_find(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure unsafe.Pointer, func_ unsafe.Pointer, data unsafe.Pointer) {
}

func Fn_signal_handler_is_connected(instance unsafe.Pointer, handlerId uint64) {}

func Fn_signal_handler_unblock(instance unsafe.Pointer, handlerId uint64) {}

func Fn_signal_handlers_block_matched(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure unsafe.Pointer, func_ unsafe.Pointer, data unsafe.Pointer) {
}

func Fn_signal_handlers_destroy(instance unsafe.Pointer) {}

func Fn_signal_handlers_disconnect_matched(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure unsafe.Pointer, func_ unsafe.Pointer, data unsafe.Pointer) {
}

func Fn_signal_handlers_unblock_matched(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure unsafe.Pointer, func_ unsafe.Pointer, data unsafe.Pointer) {
}

func Fn_signal_has_handler_pending(instance unsafe.Pointer, signalId uint, detail uint32, mayBeBlocked bool) {
}

func Fn_signal_list_ids(itype uint64) {}

func Fn_signal_lookup(name string, itype uint64) {}

func Fn_signal_name(signalId uint) {}

// UNSUPPORTED : signal_new : has varargs

// UNSUPPORTED : signal_new_class_handler : has varargs

// UNSUPPORTED : signal_new_valist : has va_list

// UNSUPPORTED : signal_newv : has callback

func Fn_signal_override_class_closure(signalId uint, instanceType uint64, classClosure unsafe.Pointer) {
}

// UNSUPPORTED : signal_override_class_handler : has callback

func Fn_signal_parse_name(detailedSignal string, itype uint64, forceDetailQuark bool) {}

func Fn_signal_query(signalId uint) {}

func Fn_signal_remove_emission_hook(signalId uint, hookId uint64) {}

func Fn_signal_stop_emission(instance unsafe.Pointer, signalId uint, detail uint32) {}

func Fn_signal_stop_emission_by_name(instance unsafe.Pointer, detailedSignal string) {}

func Fn_signal_type_cclosure_new(itype uint64, structOffset uint) {}

func Fn_source_set_closure(source unsafe.Pointer, closure unsafe.Pointer) {}

func Fn_source_set_dummy_callback(source unsafe.Pointer) {}

func Fn_strdup_value_contents(value unsafe.Pointer) {}

// UNSUPPORTED : type_add_class_cache_func : has callback

func Fn_type_add_instance_private(classType uint64, privateSize uint64) {}

// UNSUPPORTED : type_add_interface_check : has callback

func Fn_type_add_interface_dynamic(instanceType uint64, interfaceType uint64, plugin unsafe.Pointer) {
}

func Fn_type_add_interface_static(instanceType uint64, interfaceType uint64, info unsafe.Pointer) {}

func Fn_type_check_class_cast(gClass unsafe.Pointer, isAType uint64) {}

func Fn_type_check_class_is_a(gClass unsafe.Pointer, isAType uint64) {}

func Fn_type_check_instance(instance unsafe.Pointer) {}

func Fn_type_check_instance_cast(instance unsafe.Pointer, ifaceType uint64) {}

func Fn_type_check_instance_is_a(instance unsafe.Pointer, ifaceType uint64) {}

func Fn_type_check_instance_is_fundamentally_a(instance unsafe.Pointer, fundamentalType uint64) {}

func Fn_type_check_is_value_type(type_ uint64) {}

func Fn_type_check_value(value unsafe.Pointer) {}

func Fn_type_check_value_holds(value unsafe.Pointer, type_ uint64) {}

func Fn_type_children(type_ uint64) {}

func Fn_type_class_adjust_private_offset(gClass unsafe.Pointer, privateSizeOrOffset *int) {}

func Fn_type_class_peek(type_ uint64) {}

func Fn_type_class_peek_static(type_ uint64) {}

func Fn_type_class_ref(type_ uint64) {}

func Fn_type_create_instance(type_ uint64) {}

func Fn_type_default_interface_peek(gType uint64) {}

func Fn_type_default_interface_ref(gType uint64) {}

func Fn_type_default_interface_unref(gIface unsafe.Pointer) {}

func Fn_type_depth(type_ uint64) {}

func Fn_type_free_instance(instance unsafe.Pointer) {}

func Fn_type_from_name(name string) {}

func Fn_type_fundamental(typeId uint64) {}

func Fn_type_fundamental_next() {
	C.g_type_fundamental_next()
}

func Fn_type_get_plugin(type_ uint64) {}

func Fn_type_get_qdata(type_ uint64, quark uint32) {}

func Fn_type_init() {
	C.g_type_init()
}

func Fn_type_init_with_debug_flags(debugFlags int) {}

func Fn_type_interface_add_prerequisite(interfaceType uint64, prerequisiteType uint64) {}

func Fn_type_interface_get_plugin(instanceType uint64, interfaceType uint64) {}

func Fn_type_interface_peek(instanceClass unsafe.Pointer, ifaceType uint64) {}

func Fn_type_interface_prerequisites(interfaceType uint64) {}

func Fn_type_interfaces(type_ uint64) {}

func Fn_type_is_a(type_ uint64, isAType uint64) {}

func Fn_type_name(type_ uint64) {}

func Fn_type_name_from_class(gClass unsafe.Pointer) {}

func Fn_type_name_from_instance(instance unsafe.Pointer) {}

func Fn_type_next_base(leafType uint64, rootType uint64) {}

func Fn_type_parent(type_ uint64) {}

func Fn_type_qname(type_ uint64) {}

func Fn_type_query(type_ uint64) {}

func Fn_type_register_dynamic(parentType uint64, typeName string, plugin unsafe.Pointer, flags int) {}

func Fn_type_register_fundamental(typeId uint64, typeName string, info unsafe.Pointer, finfo unsafe.Pointer, flags int) {
}

func Fn_type_register_static(parentType uint64, typeName string, info unsafe.Pointer, flags int) {}

// UNSUPPORTED : type_register_static_simple : has callback

// UNSUPPORTED : type_remove_class_cache_func : has callback

// UNSUPPORTED : type_remove_interface_check : has callback

func Fn_type_set_qdata(type_ uint64, quark uint32, data unsafe.Pointer) {}

func Fn_type_test_flags(type_ uint64, flags uint) {}

func Fn_type_value_table_peek(type_ uint64) {}

// UNSUPPORTED : value_register_transform_func : has callback

func Fn_value_type_compatible(srcType uint64, destType uint64) {}

func Fn_value_type_transformable(srcType uint64, destType uint64) {}
