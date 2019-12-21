// Code generated - DO NOT EDIT.
// +build gobject_2.4

package gobject

import (
	c "github.com/pekim/gobbi/lib/internal/c"
	"unsafe"
)

// #include <glib-object.h>
import "C"

// bitfields
type ConnectFlags C.GConnectFlags
type ParamFlags C.GParamFlags
type SignalFlags C.GSignalFlags
type SignalMatchType C.GSignalMatchType
type TypeDebugFlags C.GTypeDebugFlags
type TypeFlags C.GTypeFlags
type TypeFundamentalFlags C.GTypeFundamentalFlags

// enumerations

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

// classes
type InitiallyUnowned C.GInitiallyUnowned
type Object C.GObject
type ParamSpec C.GParamSpec
type ParamSpecBoolean C.GParamSpecBoolean
type ParamSpecBoxed C.GParamSpecBoxed
type ParamSpecChar C.GParamSpecChar
type ParamSpecDouble C.GParamSpecDouble
type ParamSpecEnum C.GParamSpecEnum
type ParamSpecFlags C.GParamSpecFlags
type ParamSpecFloat C.GParamSpecFloat
type ParamSpecInt C.GParamSpecInt
type ParamSpecInt64 C.GParamSpecInt64
type ParamSpecLong C.GParamSpecLong
type ParamSpecObject C.GParamSpecObject
type ParamSpecOverride C.GParamSpecOverride
type ParamSpecParam C.GParamSpecParam
type ParamSpecPointer C.GParamSpecPointer
type ParamSpecString C.GParamSpecString
type ParamSpecUChar C.GParamSpecUChar
type ParamSpecUInt C.GParamSpecUInt
type ParamSpecUInt64 C.GParamSpecUInt64
type ParamSpecULong C.GParamSpecULong
type ParamSpecUnichar C.GParamSpecUnichar
type ParamSpecValueArray C.GParamSpecValueArray
type TypeModule C.GTypeModule

// interfaces
type TypePlugin C.GTypePlugin

func Fn_boxed_copy(boxedType c.UndefinedParamType, srcBoxed unsafe.Pointer) {}

func Fn_boxed_free(boxedType c.UndefinedParamType, boxed unsafe.Pointer) {}

func Fn_boxed_type_register_static(name string, boxedCopy c.UndefinedParamType, boxedFree c.UndefinedParamType) {
}

func Fn_cclosure_marshal_BOOLEAN__BOXED_BOXED(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_BOOLEAN__FLAGS(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_STRING__OBJECT_POINTER(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__BOOLEAN(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__BOXED(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__CHAR(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__DOUBLE(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__ENUM(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__FLAGS(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__FLOAT(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__INT(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__LONG(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__OBJECT(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__PARAM(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__POINTER(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__STRING(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__UCHAR(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__UINT(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__UINT_POINTER(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__ULONG(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_marshal_VOID__VOID(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues uint, paramValues c.UndefinedParamType, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
}

func Fn_cclosure_new(callbackFunc c.UndefinedParamType, userData unsafe.Pointer, destroyData c.UndefinedParamType) {
}

func Fn_cclosure_new_object(callbackFunc c.UndefinedParamType, object c.UndefinedParamType) {}

func Fn_cclosure_new_object_swap(callbackFunc c.UndefinedParamType, object c.UndefinedParamType) {}

func Fn_cclosure_new_swap(callbackFunc c.UndefinedParamType, userData unsafe.Pointer, destroyData c.UndefinedParamType) {
}

func Fn_enum_complete_type_info(gEnumType c.UndefinedParamType, constValues c.UndefinedParamType) {}

func Fn_enum_get_value(enumClass c.UndefinedParamType, value int) {}

func Fn_enum_get_value_by_name(enumClass c.UndefinedParamType, name string) {}

func Fn_enum_get_value_by_nick(enumClass c.UndefinedParamType, nick string) {}

func Fn_enum_register_static(name string, constStaticValues c.UndefinedParamType) {}

func Fn_flags_complete_type_info(gFlagsType c.UndefinedParamType, constValues c.UndefinedParamType) {}

func Fn_flags_get_first_value(flagsClass c.UndefinedParamType, value uint) {}

func Fn_flags_get_value_by_name(flagsClass c.UndefinedParamType, name string) {}

func Fn_flags_get_value_by_nick(flagsClass c.UndefinedParamType, nick string) {}

func Fn_flags_register_static(name string, constStaticValues c.UndefinedParamType) {}

func Fn_gtype_get_type() {}

func Fn_param_spec_boolean(name string, nick string, blurb string, defaultValue bool, flags c.UndefinedParamType) {
}

func Fn_param_spec_boxed(name string, nick string, blurb string, boxedType c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_char(name string, nick string, blurb string, minimum int8, maximum int8, defaultValue int8, flags c.UndefinedParamType) {
}

func Fn_param_spec_double(name string, nick string, blurb string, minimum float64, maximum float64, defaultValue float64, flags c.UndefinedParamType) {
}

func Fn_param_spec_enum(name string, nick string, blurb string, enumType c.UndefinedParamType, defaultValue int, flags c.UndefinedParamType) {
}

func Fn_param_spec_flags(name string, nick string, blurb string, flagsType c.UndefinedParamType, defaultValue uint, flags c.UndefinedParamType) {
}

func Fn_param_spec_float(name string, nick string, blurb string, minimum float32, maximum float32, defaultValue float32, flags c.UndefinedParamType) {
}

func Fn_param_spec_int(name string, nick string, blurb string, minimum int, maximum int, defaultValue int, flags c.UndefinedParamType) {
}

func Fn_param_spec_int64(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags c.UndefinedParamType) {
}

func Fn_param_spec_long(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags c.UndefinedParamType) {
}

func Fn_param_spec_object(name string, nick string, blurb string, objectType c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_override(name string, overridden c.UndefinedParamType) {}

func Fn_param_spec_param(name string, nick string, blurb string, paramType c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_pointer(name string, nick string, blurb string, flags c.UndefinedParamType) {}

func Fn_param_spec_pool_new(typePrefixing bool) {}

func Fn_param_spec_string(name string, nick string, blurb string, defaultValue string, flags c.UndefinedParamType) {
}

func Fn_param_spec_uchar(name string, nick string, blurb string, minimum uint8, maximum uint8, defaultValue uint8, flags c.UndefinedParamType) {
}

func Fn_param_spec_uint(name string, nick string, blurb string, minimum uint, maximum uint, defaultValue uint, flags c.UndefinedParamType) {
}

func Fn_param_spec_uint64(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags c.UndefinedParamType) {
}

func Fn_param_spec_ulong(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags c.UndefinedParamType) {
}

func Fn_param_spec_unichar(name string, nick string, blurb string, defaultValue rune, flags c.UndefinedParamType) {
}

func Fn_param_spec_value_array(name string, nick string, blurb string, elementSpec c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_type_register_static(name string, pspecInfo c.UndefinedParamType) {}

func Fn_param_value_convert(pspec c.UndefinedParamType, srcValue c.UndefinedParamType, destValue c.UndefinedParamType, strictValidation bool) {
}

func Fn_param_value_defaults(pspec c.UndefinedParamType, value c.UndefinedParamType) {}

func Fn_param_value_set_default(pspec c.UndefinedParamType, value c.UndefinedParamType) {}

func Fn_param_value_validate(pspec c.UndefinedParamType, value c.UndefinedParamType) {}

func Fn_param_values_cmp(pspec c.UndefinedParamType, value1 c.UndefinedParamType, value2 c.UndefinedParamType) {
}

func Fn_pointer_type_register_static(name string) {}

func Fn_signal_accumulator_true_handled(ihint c.UndefinedParamType, returnAccu c.UndefinedParamType, handlerReturn c.UndefinedParamType, dummy unsafe.Pointer) {
}

func Fn_signal_add_emission_hook(signalId uint, detail c.UndefinedParamType, hookFunc c.UndefinedParamType, hookData unsafe.Pointer, dataDestroy c.UndefinedParamType) {
}

func Fn_signal_chain_from_overridden(instanceAndParams c.UndefinedParamType, returnValue c.UndefinedParamType) {
}

// UNSUPPORTED : signal_chain_from_overridden_handler : has varargs

func Fn_signal_connect_closure(instance unsafe.Pointer, detailedSignal string, closure c.UndefinedParamType, after bool) {
}

func Fn_signal_connect_closure_by_id(instance unsafe.Pointer, signalId uint, detail c.UndefinedParamType, closure c.UndefinedParamType, after bool) {
}

func Fn_signal_connect_data(instance unsafe.Pointer, detailedSignal string, cHandler c.UndefinedParamType, data unsafe.Pointer, destroyData c.UndefinedParamType, connectFlags c.UndefinedParamType) {
}

func Fn_signal_connect_object(instance unsafe.Pointer, detailedSignal string, cHandler c.UndefinedParamType, gobject unsafe.Pointer, connectFlags c.UndefinedParamType) {
}

// UNSUPPORTED : signal_emit : has varargs

// UNSUPPORTED : signal_emit_by_name : has varargs

func Fn_signal_emit_valist(instance unsafe.Pointer, signalId uint, detail c.UndefinedParamType, varArgs c.UndefinedParamType) {
}

func Fn_signal_emitv(instanceAndParams c.UndefinedParamType, signalId uint, detail c.UndefinedParamType, returnValue c.UndefinedParamType) {
}

func Fn_signal_get_invocation_hint(instance unsafe.Pointer) {}

func Fn_signal_handler_block(instance unsafe.Pointer, handlerId uint64) {}

func Fn_signal_handler_disconnect(instance unsafe.Pointer, handlerId uint64) {}

func Fn_signal_handler_find(instance unsafe.Pointer, mask c.UndefinedParamType, signalId uint, detail c.UndefinedParamType, closure c.UndefinedParamType, func_ unsafe.Pointer, data unsafe.Pointer) {
}

func Fn_signal_handler_is_connected(instance unsafe.Pointer, handlerId uint64) {}

func Fn_signal_handler_unblock(instance unsafe.Pointer, handlerId uint64) {}

func Fn_signal_handlers_block_matched(instance unsafe.Pointer, mask c.UndefinedParamType, signalId uint, detail c.UndefinedParamType, closure c.UndefinedParamType, func_ unsafe.Pointer, data unsafe.Pointer) {
}

func Fn_signal_handlers_destroy(instance unsafe.Pointer) {}

func Fn_signal_handlers_disconnect_matched(instance unsafe.Pointer, mask c.UndefinedParamType, signalId uint, detail c.UndefinedParamType, closure c.UndefinedParamType, func_ unsafe.Pointer, data unsafe.Pointer) {
}

func Fn_signal_handlers_unblock_matched(instance unsafe.Pointer, mask c.UndefinedParamType, signalId uint, detail c.UndefinedParamType, closure c.UndefinedParamType, func_ unsafe.Pointer, data unsafe.Pointer) {
}

func Fn_signal_has_handler_pending(instance unsafe.Pointer, signalId uint, detail c.UndefinedParamType, mayBeBlocked bool) {
}

func Fn_signal_list_ids(itype c.UndefinedParamType) {}

func Fn_signal_lookup(name string, itype c.UndefinedParamType) {}

func Fn_signal_name(signalId uint) {}

// UNSUPPORTED : signal_new : has varargs

// UNSUPPORTED : signal_new_class_handler : has varargs

func Fn_signal_new_valist(signalName string, itype c.UndefinedParamType, signalFlags c.UndefinedParamType, classClosure c.UndefinedParamType, accumulator c.UndefinedParamType, accuData unsafe.Pointer, cMarshaller c.UndefinedParamType, returnType c.UndefinedParamType, nParams uint, args c.UndefinedParamType) {
}

func Fn_signal_newv(signalName string, itype c.UndefinedParamType, signalFlags c.UndefinedParamType, classClosure c.UndefinedParamType, accumulator c.UndefinedParamType, accuData unsafe.Pointer, cMarshaller c.UndefinedParamType, returnType c.UndefinedParamType, nParams uint, paramTypes c.UndefinedParamType) {
}

func Fn_signal_override_class_closure(signalId uint, instanceType c.UndefinedParamType, classClosure c.UndefinedParamType) {
}

func Fn_signal_parse_name(detailedSignal string, itype c.UndefinedParamType, forceDetailQuark bool) {}

func Fn_signal_query(signalId uint) {}

func Fn_signal_remove_emission_hook(signalId uint, hookId uint64) {}

func Fn_signal_stop_emission(instance unsafe.Pointer, signalId uint, detail c.UndefinedParamType) {}

func Fn_signal_stop_emission_by_name(instance unsafe.Pointer, detailedSignal string) {}

func Fn_signal_type_cclosure_new(itype c.UndefinedParamType, structOffset uint) {}

func Fn_source_set_closure(source c.UndefinedParamType, closure c.UndefinedParamType) {}

func Fn_source_set_dummy_callback(source c.UndefinedParamType) {}

func Fn_strdup_value_contents(value c.UndefinedParamType) {}

func Fn_type_add_class_cache_func(cacheData unsafe.Pointer, cacheFunc c.UndefinedParamType) {}

func Fn_type_add_instance_private(classType c.UndefinedParamType, privateSize uint64) {}

func Fn_type_add_interface_check(checkData unsafe.Pointer, checkFunc c.UndefinedParamType) {}

func Fn_type_add_interface_dynamic(instanceType c.UndefinedParamType, interfaceType c.UndefinedParamType, plugin c.UndefinedParamType) {
}

func Fn_type_add_interface_static(instanceType c.UndefinedParamType, interfaceType c.UndefinedParamType, info c.UndefinedParamType) {
}

func Fn_type_check_class_cast(gClass c.UndefinedParamType, isAType c.UndefinedParamType) {}

func Fn_type_check_class_is_a(gClass c.UndefinedParamType, isAType c.UndefinedParamType) {}

func Fn_type_check_instance(instance c.UndefinedParamType) {}

func Fn_type_check_instance_cast(instance c.UndefinedParamType, ifaceType c.UndefinedParamType) {}

func Fn_type_check_instance_is_a(instance c.UndefinedParamType, ifaceType c.UndefinedParamType) {}

func Fn_type_check_instance_is_fundamentally_a(instance c.UndefinedParamType, fundamentalType c.UndefinedParamType) {
}

func Fn_type_check_is_value_type(type_ c.UndefinedParamType) {}

func Fn_type_check_value(value c.UndefinedParamType) {}

func Fn_type_check_value_holds(value c.UndefinedParamType, type_ c.UndefinedParamType) {}

func Fn_type_children(type_ c.UndefinedParamType) {}

func Fn_type_class_adjust_private_offset(gClass unsafe.Pointer, privateSizeOrOffset *int) {}

func Fn_type_class_peek(type_ c.UndefinedParamType) {}

func Fn_type_class_peek_static(type_ c.UndefinedParamType) {}

func Fn_type_class_ref(type_ c.UndefinedParamType) {}

func Fn_type_create_instance(type_ c.UndefinedParamType) {}

func Fn_type_default_interface_peek(gType c.UndefinedParamType) {}

func Fn_type_default_interface_ref(gType c.UndefinedParamType) {}

func Fn_type_default_interface_unref(gIface unsafe.Pointer) {}

func Fn_type_depth(type_ c.UndefinedParamType) {}

func Fn_type_free_instance(instance c.UndefinedParamType) {}

func Fn_type_from_name(name string) {}

func Fn_type_fundamental(typeId c.UndefinedParamType) {}

func Fn_type_fundamental_next() {}

func Fn_type_get_plugin(type_ c.UndefinedParamType) {}

func Fn_type_get_qdata(type_ c.UndefinedParamType, quark c.UndefinedParamType) {}

func Fn_type_init() {}

func Fn_type_init_with_debug_flags(debugFlags c.UndefinedParamType) {}

func Fn_type_interface_add_prerequisite(interfaceType c.UndefinedParamType, prerequisiteType c.UndefinedParamType) {
}

func Fn_type_interface_get_plugin(instanceType c.UndefinedParamType, interfaceType c.UndefinedParamType) {
}

func Fn_type_interface_peek(instanceClass unsafe.Pointer, ifaceType c.UndefinedParamType) {}

func Fn_type_interface_prerequisites(interfaceType c.UndefinedParamType) {}

func Fn_type_interfaces(type_ c.UndefinedParamType) {}

func Fn_type_is_a(type_ c.UndefinedParamType, isAType c.UndefinedParamType) {}

func Fn_type_name(type_ c.UndefinedParamType) {}

func Fn_type_name_from_class(gClass c.UndefinedParamType) {}

func Fn_type_name_from_instance(instance c.UndefinedParamType) {}

func Fn_type_next_base(leafType c.UndefinedParamType, rootType c.UndefinedParamType) {}

func Fn_type_parent(type_ c.UndefinedParamType) {}

func Fn_type_qname(type_ c.UndefinedParamType) {}

func Fn_type_query(type_ c.UndefinedParamType) {}

func Fn_type_register_dynamic(parentType c.UndefinedParamType, typeName string, plugin c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_type_register_fundamental(typeId c.UndefinedParamType, typeName string, info c.UndefinedParamType, finfo c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_type_register_static(parentType c.UndefinedParamType, typeName string, info c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_type_remove_class_cache_func(cacheData unsafe.Pointer, cacheFunc c.UndefinedParamType) {}

func Fn_type_remove_interface_check(checkData unsafe.Pointer, checkFunc c.UndefinedParamType) {}

func Fn_type_set_qdata(type_ c.UndefinedParamType, quark c.UndefinedParamType, data unsafe.Pointer) {}

func Fn_type_test_flags(type_ c.UndefinedParamType, flags uint) {}

func Fn_type_value_table_peek(type_ c.UndefinedParamType) {}

func Fn_value_register_transform_func(srcType c.UndefinedParamType, destType c.UndefinedParamType, transformFunc c.UndefinedParamType) {
}

func Fn_value_type_compatible(srcType c.UndefinedParamType, destType c.UndefinedParamType) {}

func Fn_value_type_transformable(srcType c.UndefinedParamType, destType c.UndefinedParamType) {}

func Fn_bad() {}
