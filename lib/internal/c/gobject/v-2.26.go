// Code generated - DO NOT EDIT.
// +build gobject_2.26

package gobject

import c "github.com/pekim/gobbi/lib/internal/c"

// #include <glib-object.h>
import "C"

// bitfields
type BindingFlags C.GBindingFlags
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
type Binding C.GBinding
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
type ParamSpecGType C.GParamSpecGType
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
type ParamSpecVariant C.GParamSpecVariant
type TypeModule C.GTypeModule

// interfaces
type TypePlugin C.GTypePlugin

func Fn_boxed_copy(boxedType c.UndefinedParamType, srcBoxed c.UndefinedParamType) {}

func Fn_boxed_free(boxedType c.UndefinedParamType, boxed c.UndefinedParamType) {}

func Fn_boxed_type_register_static(name c.UndefinedParamType, boxedCopy c.UndefinedParamType, boxedFree c.UndefinedParamType) {
}

func Fn_cclosure_marshal_BOOLEAN__BOXED_BOXED(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_BOOLEAN__FLAGS(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_STRING__OBJECT_POINTER(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_VOID__BOOLEAN(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_VOID__BOXED(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_VOID__CHAR(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_VOID__DOUBLE(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_VOID__ENUM(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_VOID__FLAGS(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_VOID__FLOAT(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_VOID__INT(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_VOID__LONG(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_VOID__OBJECT(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_VOID__PARAM(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_VOID__POINTER(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_VOID__STRING(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_VOID__UCHAR(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_VOID__UINT(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_VOID__UINT_POINTER(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_VOID__ULONG(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_VOID__VARIANT(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_marshal_VOID__VOID(closure c.UndefinedParamType, returnValue c.UndefinedParamType, nParamValues c.UndefinedParamType, paramValues c.UndefinedParamType, invocationHint c.UndefinedParamType, marshalData c.UndefinedParamType) {
}

func Fn_cclosure_new(callbackFunc c.UndefinedParamType, userData c.UndefinedParamType, destroyData c.UndefinedParamType) {
}

func Fn_cclosure_new_object(callbackFunc c.UndefinedParamType, object c.UndefinedParamType) {}

func Fn_cclosure_new_object_swap(callbackFunc c.UndefinedParamType, object c.UndefinedParamType) {}

func Fn_cclosure_new_swap(callbackFunc c.UndefinedParamType, userData c.UndefinedParamType, destroyData c.UndefinedParamType) {
}

func Fn_enum_complete_type_info(gEnumType c.UndefinedParamType, info c.UndefinedParamType, constValues c.UndefinedParamType) {
}

func Fn_enum_get_value(enumClass c.UndefinedParamType, value c.UndefinedParamType) {}

func Fn_enum_get_value_by_name(enumClass c.UndefinedParamType, name c.UndefinedParamType) {}

func Fn_enum_get_value_by_nick(enumClass c.UndefinedParamType, nick c.UndefinedParamType) {}

func Fn_enum_register_static(name c.UndefinedParamType, constStaticValues c.UndefinedParamType) {}

func Fn_flags_complete_type_info(gFlagsType c.UndefinedParamType, info c.UndefinedParamType, constValues c.UndefinedParamType) {
}

func Fn_flags_get_first_value(flagsClass c.UndefinedParamType, value c.UndefinedParamType) {}

func Fn_flags_get_value_by_name(flagsClass c.UndefinedParamType, name c.UndefinedParamType) {}

func Fn_flags_get_value_by_nick(flagsClass c.UndefinedParamType, nick c.UndefinedParamType) {}

func Fn_flags_register_static(name c.UndefinedParamType, constStaticValues c.UndefinedParamType) {}

func Fn_gtype_get_type() {}

func Fn_param_spec_boolean(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, defaultValue c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_boxed(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, boxedType c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_char(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, minimum c.UndefinedParamType, maximum c.UndefinedParamType, defaultValue c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_double(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, minimum c.UndefinedParamType, maximum c.UndefinedParamType, defaultValue c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_enum(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, enumType c.UndefinedParamType, defaultValue c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_flags(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, flagsType c.UndefinedParamType, defaultValue c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_float(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, minimum c.UndefinedParamType, maximum c.UndefinedParamType, defaultValue c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_gtype(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, isAType c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_int(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, minimum c.UndefinedParamType, maximum c.UndefinedParamType, defaultValue c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_int64(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, minimum c.UndefinedParamType, maximum c.UndefinedParamType, defaultValue c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_long(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, minimum c.UndefinedParamType, maximum c.UndefinedParamType, defaultValue c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_object(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, objectType c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_override(name c.UndefinedParamType, overridden c.UndefinedParamType) {}

func Fn_param_spec_param(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, paramType c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_pointer(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_pool_new(typePrefixing c.UndefinedParamType) {}

func Fn_param_spec_string(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, defaultValue c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_uchar(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, minimum c.UndefinedParamType, maximum c.UndefinedParamType, defaultValue c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_uint(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, minimum c.UndefinedParamType, maximum c.UndefinedParamType, defaultValue c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_uint64(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, minimum c.UndefinedParamType, maximum c.UndefinedParamType, defaultValue c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_ulong(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, minimum c.UndefinedParamType, maximum c.UndefinedParamType, defaultValue c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_unichar(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, defaultValue c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_value_array(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, elementSpec c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_spec_variant(name c.UndefinedParamType, nick c.UndefinedParamType, blurb c.UndefinedParamType, type_ c.UndefinedParamType, defaultValue c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_param_type_register_static(name c.UndefinedParamType, pspecInfo c.UndefinedParamType) {}

func Fn_param_value_convert(pspec c.UndefinedParamType, srcValue c.UndefinedParamType, destValue c.UndefinedParamType, strictValidation c.UndefinedParamType) {
}

func Fn_param_value_defaults(pspec c.UndefinedParamType, value c.UndefinedParamType) {}

func Fn_param_value_set_default(pspec c.UndefinedParamType, value c.UndefinedParamType) {}

func Fn_param_value_validate(pspec c.UndefinedParamType, value c.UndefinedParamType) {}

func Fn_param_values_cmp(pspec c.UndefinedParamType, value1 c.UndefinedParamType, value2 c.UndefinedParamType) {
}

func Fn_pointer_type_register_static(name c.UndefinedParamType) {}

func Fn_signal_accumulator_true_handled(ihint c.UndefinedParamType, returnAccu c.UndefinedParamType, handlerReturn c.UndefinedParamType, dummy c.UndefinedParamType) {
}

func Fn_signal_add_emission_hook(signalId c.UndefinedParamType, detail c.UndefinedParamType, hookFunc c.UndefinedParamType, hookData c.UndefinedParamType, dataDestroy c.UndefinedParamType) {
}

func Fn_signal_chain_from_overridden(instanceAndParams c.UndefinedParamType, returnValue c.UndefinedParamType) {
}

// UNSUPPORTED : signal_chain_from_overridden_handler : has varargs

func Fn_signal_connect_closure(instance c.UndefinedParamType, detailedSignal c.UndefinedParamType, closure c.UndefinedParamType, after c.UndefinedParamType) {
}

func Fn_signal_connect_closure_by_id(instance c.UndefinedParamType, signalId c.UndefinedParamType, detail c.UndefinedParamType, closure c.UndefinedParamType, after c.UndefinedParamType) {
}

func Fn_signal_connect_data(instance c.UndefinedParamType, detailedSignal c.UndefinedParamType, cHandler c.UndefinedParamType, data c.UndefinedParamType, destroyData c.UndefinedParamType, connectFlags c.UndefinedParamType) {
}

func Fn_signal_connect_object(instance c.UndefinedParamType, detailedSignal c.UndefinedParamType, cHandler c.UndefinedParamType, gobject c.UndefinedParamType, connectFlags c.UndefinedParamType) {
}

// UNSUPPORTED : signal_emit : has varargs

// UNSUPPORTED : signal_emit_by_name : has varargs

func Fn_signal_emit_valist(instance c.UndefinedParamType, signalId c.UndefinedParamType, detail c.UndefinedParamType, varArgs c.UndefinedParamType) {
}

func Fn_signal_emitv(instanceAndParams c.UndefinedParamType, signalId c.UndefinedParamType, detail c.UndefinedParamType, returnValue c.UndefinedParamType) {
}

func Fn_signal_get_invocation_hint(instance c.UndefinedParamType) {}

func Fn_signal_handler_block(instance c.UndefinedParamType, handlerId c.UndefinedParamType) {}

func Fn_signal_handler_disconnect(instance c.UndefinedParamType, handlerId c.UndefinedParamType) {}

func Fn_signal_handler_find(instance c.UndefinedParamType, mask c.UndefinedParamType, signalId c.UndefinedParamType, detail c.UndefinedParamType, closure c.UndefinedParamType, func_ c.UndefinedParamType, data c.UndefinedParamType) {
}

func Fn_signal_handler_is_connected(instance c.UndefinedParamType, handlerId c.UndefinedParamType) {}

func Fn_signal_handler_unblock(instance c.UndefinedParamType, handlerId c.UndefinedParamType) {}

func Fn_signal_handlers_block_matched(instance c.UndefinedParamType, mask c.UndefinedParamType, signalId c.UndefinedParamType, detail c.UndefinedParamType, closure c.UndefinedParamType, func_ c.UndefinedParamType, data c.UndefinedParamType) {
}

func Fn_signal_handlers_destroy(instance c.UndefinedParamType) {}

func Fn_signal_handlers_disconnect_matched(instance c.UndefinedParamType, mask c.UndefinedParamType, signalId c.UndefinedParamType, detail c.UndefinedParamType, closure c.UndefinedParamType, func_ c.UndefinedParamType, data c.UndefinedParamType) {
}

func Fn_signal_handlers_unblock_matched(instance c.UndefinedParamType, mask c.UndefinedParamType, signalId c.UndefinedParamType, detail c.UndefinedParamType, closure c.UndefinedParamType, func_ c.UndefinedParamType, data c.UndefinedParamType) {
}

func Fn_signal_has_handler_pending(instance c.UndefinedParamType, signalId c.UndefinedParamType, detail c.UndefinedParamType, mayBeBlocked c.UndefinedParamType) {
}

func Fn_signal_list_ids(itype c.UndefinedParamType, nIds c.UndefinedParamType) {}

func Fn_signal_lookup(name c.UndefinedParamType, itype c.UndefinedParamType) {}

func Fn_signal_name(signalId c.UndefinedParamType) {}

// UNSUPPORTED : signal_new : has varargs

// UNSUPPORTED : signal_new_class_handler : has varargs

func Fn_signal_new_valist(signalName c.UndefinedParamType, itype c.UndefinedParamType, signalFlags c.UndefinedParamType, classClosure c.UndefinedParamType, accumulator c.UndefinedParamType, accuData c.UndefinedParamType, cMarshaller c.UndefinedParamType, returnType c.UndefinedParamType, nParams c.UndefinedParamType, args c.UndefinedParamType) {
}

func Fn_signal_newv(signalName c.UndefinedParamType, itype c.UndefinedParamType, signalFlags c.UndefinedParamType, classClosure c.UndefinedParamType, accumulator c.UndefinedParamType, accuData c.UndefinedParamType, cMarshaller c.UndefinedParamType, returnType c.UndefinedParamType, nParams c.UndefinedParamType, paramTypes c.UndefinedParamType) {
}

func Fn_signal_override_class_closure(signalId c.UndefinedParamType, instanceType c.UndefinedParamType, classClosure c.UndefinedParamType) {
}

func Fn_signal_override_class_handler(signalName c.UndefinedParamType, instanceType c.UndefinedParamType, classHandler c.UndefinedParamType) {
}

func Fn_signal_parse_name(detailedSignal c.UndefinedParamType, itype c.UndefinedParamType, signalIdP c.UndefinedParamType, detailP c.UndefinedParamType, forceDetailQuark c.UndefinedParamType) {
}

func Fn_signal_query(signalId c.UndefinedParamType, query c.UndefinedParamType) {}

func Fn_signal_remove_emission_hook(signalId c.UndefinedParamType, hookId c.UndefinedParamType) {}

func Fn_signal_stop_emission(instance c.UndefinedParamType, signalId c.UndefinedParamType, detail c.UndefinedParamType) {
}

func Fn_signal_stop_emission_by_name(instance c.UndefinedParamType, detailedSignal c.UndefinedParamType) {
}

func Fn_signal_type_cclosure_new(itype c.UndefinedParamType, structOffset c.UndefinedParamType) {}

func Fn_source_set_closure(source c.UndefinedParamType, closure c.UndefinedParamType) {}

func Fn_source_set_dummy_callback(source c.UndefinedParamType) {}

func Fn_strdup_value_contents(value c.UndefinedParamType) {}

func Fn_type_add_class_cache_func(cacheData c.UndefinedParamType, cacheFunc c.UndefinedParamType) {}

func Fn_type_add_class_private(classType c.UndefinedParamType, privateSize c.UndefinedParamType) {}

func Fn_type_add_instance_private(classType c.UndefinedParamType, privateSize c.UndefinedParamType) {}

func Fn_type_add_interface_check(checkData c.UndefinedParamType, checkFunc c.UndefinedParamType) {}

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

func Fn_type_children(type_ c.UndefinedParamType, nChildren c.UndefinedParamType) {}

func Fn_type_class_adjust_private_offset(gClass c.UndefinedParamType, privateSizeOrOffset c.UndefinedParamType) {
}

func Fn_type_class_peek(type_ c.UndefinedParamType) {}

func Fn_type_class_peek_static(type_ c.UndefinedParamType) {}

func Fn_type_class_ref(type_ c.UndefinedParamType) {}

func Fn_type_create_instance(type_ c.UndefinedParamType) {}

func Fn_type_default_interface_peek(gType c.UndefinedParamType) {}

func Fn_type_default_interface_ref(gType c.UndefinedParamType) {}

func Fn_type_default_interface_unref(gIface c.UndefinedParamType) {}

func Fn_type_depth(type_ c.UndefinedParamType) {}

func Fn_type_free_instance(instance c.UndefinedParamType) {}

func Fn_type_from_name(name c.UndefinedParamType) {}

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

func Fn_type_interface_peek(instanceClass c.UndefinedParamType, ifaceType c.UndefinedParamType) {}

func Fn_type_interface_prerequisites(interfaceType c.UndefinedParamType, nPrerequisites c.UndefinedParamType) {
}

func Fn_type_interfaces(type_ c.UndefinedParamType, nInterfaces c.UndefinedParamType) {}

func Fn_type_is_a(type_ c.UndefinedParamType, isAType c.UndefinedParamType) {}

func Fn_type_name(type_ c.UndefinedParamType) {}

func Fn_type_name_from_class(gClass c.UndefinedParamType) {}

func Fn_type_name_from_instance(instance c.UndefinedParamType) {}

func Fn_type_next_base(leafType c.UndefinedParamType, rootType c.UndefinedParamType) {}

func Fn_type_parent(type_ c.UndefinedParamType) {}

func Fn_type_qname(type_ c.UndefinedParamType) {}

func Fn_type_query(type_ c.UndefinedParamType, query c.UndefinedParamType) {}

func Fn_type_register_dynamic(parentType c.UndefinedParamType, typeName c.UndefinedParamType, plugin c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_type_register_fundamental(typeId c.UndefinedParamType, typeName c.UndefinedParamType, info c.UndefinedParamType, finfo c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_type_register_static(parentType c.UndefinedParamType, typeName c.UndefinedParamType, info c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_type_register_static_simple(parentType c.UndefinedParamType, typeName c.UndefinedParamType, classSize c.UndefinedParamType, classInit c.UndefinedParamType, instanceSize c.UndefinedParamType, instanceInit c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_type_remove_class_cache_func(cacheData c.UndefinedParamType, cacheFunc c.UndefinedParamType) {}

func Fn_type_remove_interface_check(checkData c.UndefinedParamType, checkFunc c.UndefinedParamType) {}

func Fn_type_set_qdata(type_ c.UndefinedParamType, quark c.UndefinedParamType, data c.UndefinedParamType) {
}

func Fn_type_test_flags(type_ c.UndefinedParamType, flags c.UndefinedParamType) {}

func Fn_type_value_table_peek(type_ c.UndefinedParamType) {}

func Fn_value_register_transform_func(srcType c.UndefinedParamType, destType c.UndefinedParamType, transformFunc c.UndefinedParamType) {
}

func Fn_value_type_compatible(srcType c.UndefinedParamType, destType c.UndefinedParamType) {}

func Fn_value_type_transformable(srcType c.UndefinedParamType, destType c.UndefinedParamType) {}

func Fn_bad() {}
