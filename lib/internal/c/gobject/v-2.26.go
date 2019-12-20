// Code generated - DO NOT EDIT.
// +build gobject_2.26

package gobject

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

func Fn_boxed_copy(boxedType string, srcBoxed string) {}

func Fn_boxed_free(boxedType string, boxed string) {}

func Fn_boxed_type_register_static(name string, boxedCopy string, boxedFree string) {}

func Fn_cclosure_marshal_BOOLEAN__BOXED_BOXED(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_BOOLEAN__FLAGS(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_STRING__OBJECT_POINTER(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_VOID__BOOLEAN(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_VOID__BOXED(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_VOID__CHAR(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_VOID__DOUBLE(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_VOID__ENUM(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_VOID__FLAGS(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_VOID__FLOAT(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_VOID__INT(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_VOID__LONG(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_VOID__OBJECT(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_VOID__PARAM(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_VOID__POINTER(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_VOID__STRING(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_VOID__UCHAR(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_VOID__UINT(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_VOID__UINT_POINTER(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_VOID__ULONG(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_VOID__VARIANT(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_marshal_VOID__VOID(closure string, returnValue string, nParamValues string, paramValues string, invocationHint string, marshalData string) {
}

func Fn_cclosure_new(callbackFunc string, userData string, destroyData string) {}

func Fn_cclosure_new_object(callbackFunc string, object string) {}

func Fn_cclosure_new_object_swap(callbackFunc string, object string) {}

func Fn_cclosure_new_swap(callbackFunc string, userData string, destroyData string) {}

func Fn_enum_complete_type_info(gEnumType string, info string, constValues string) {}

func Fn_enum_get_value(enumClass string, value string) {}

func Fn_enum_get_value_by_name(enumClass string, name string) {}

func Fn_enum_get_value_by_nick(enumClass string, nick string) {}

func Fn_enum_register_static(name string, constStaticValues string) {}

func Fn_flags_complete_type_info(gFlagsType string, info string, constValues string) {}

func Fn_flags_get_first_value(flagsClass string, value string) {}

func Fn_flags_get_value_by_name(flagsClass string, name string) {}

func Fn_flags_get_value_by_nick(flagsClass string, nick string) {}

func Fn_flags_register_static(name string, constStaticValues string) {}

func Fn_gtype_get_type() {}

func Fn_param_spec_boolean(name string, nick string, blurb string, defaultValue string, flags string) {}

func Fn_param_spec_boxed(name string, nick string, blurb string, boxedType string, flags string) {}

func Fn_param_spec_char(name string, nick string, blurb string, minimum string, maximum string, defaultValue string, flags string) {
}

func Fn_param_spec_double(name string, nick string, blurb string, minimum string, maximum string, defaultValue string, flags string) {
}

func Fn_param_spec_enum(name string, nick string, blurb string, enumType string, defaultValue string, flags string) {
}

func Fn_param_spec_flags(name string, nick string, blurb string, flagsType string, defaultValue string, flags string) {
}

func Fn_param_spec_float(name string, nick string, blurb string, minimum string, maximum string, defaultValue string, flags string) {
}

func Fn_param_spec_gtype(name string, nick string, blurb string, isAType string, flags string) {}

func Fn_param_spec_int(name string, nick string, blurb string, minimum string, maximum string, defaultValue string, flags string) {
}

func Fn_param_spec_int64(name string, nick string, blurb string, minimum string, maximum string, defaultValue string, flags string) {
}

func Fn_param_spec_long(name string, nick string, blurb string, minimum string, maximum string, defaultValue string, flags string) {
}

func Fn_param_spec_object(name string, nick string, blurb string, objectType string, flags string) {}

func Fn_param_spec_override(name string, overridden string) {}

func Fn_param_spec_param(name string, nick string, blurb string, paramType string, flags string) {}

func Fn_param_spec_pointer(name string, nick string, blurb string, flags string) {}

func Fn_param_spec_pool_new(typePrefixing string) {}

func Fn_param_spec_string(name string, nick string, blurb string, defaultValue string, flags string) {}

func Fn_param_spec_uchar(name string, nick string, blurb string, minimum string, maximum string, defaultValue string, flags string) {
}

func Fn_param_spec_uint(name string, nick string, blurb string, minimum string, maximum string, defaultValue string, flags string) {
}

func Fn_param_spec_uint64(name string, nick string, blurb string, minimum string, maximum string, defaultValue string, flags string) {
}

func Fn_param_spec_ulong(name string, nick string, blurb string, minimum string, maximum string, defaultValue string, flags string) {
}

func Fn_param_spec_unichar(name string, nick string, blurb string, defaultValue string, flags string) {}

func Fn_param_spec_value_array(name string, nick string, blurb string, elementSpec string, flags string) {
}

func Fn_param_spec_variant(name string, nick string, blurb string, type_ string, defaultValue string, flags string) {
}

func Fn_param_type_register_static(name string, pspecInfo string) {}

func Fn_param_value_convert(pspec string, srcValue string, destValue string, strictValidation string) {}

func Fn_param_value_defaults(pspec string, value string) {}

func Fn_param_value_set_default(pspec string, value string) {}

func Fn_param_value_validate(pspec string, value string) {}

func Fn_param_values_cmp(pspec string, value1 string, value2 string) {}

func Fn_pointer_type_register_static(name string) {}

func Fn_signal_accumulator_true_handled(ihint string, returnAccu string, handlerReturn string, dummy string) {
}

func Fn_signal_add_emission_hook(signalId string, detail string, hookFunc string, hookData string, dataDestroy string) {
}

func Fn_signal_chain_from_overridden(instanceAndParams string, returnValue string) {}

// UNSUPPORTED : signal_chain_from_overridden_handler : has varargs

func Fn_signal_connect_closure(instance string, detailedSignal string, closure string, after string) {}

func Fn_signal_connect_closure_by_id(instance string, signalId string, detail string, closure string, after string) {
}

func Fn_signal_connect_data(instance string, detailedSignal string, cHandler string, data string, destroyData string, connectFlags string) {
}

func Fn_signal_connect_object(instance string, detailedSignal string, cHandler string, gobject string, connectFlags string) {
}

// UNSUPPORTED : signal_emit : has varargs

// UNSUPPORTED : signal_emit_by_name : has varargs

func Fn_signal_emit_valist(instance string, signalId string, detail string, varArgs string) {}

func Fn_signal_emitv(instanceAndParams string, signalId string, detail string, returnValue string) {}

func Fn_signal_get_invocation_hint(instance string) {}

func Fn_signal_handler_block(instance string, handlerId string) {}

func Fn_signal_handler_disconnect(instance string, handlerId string) {}

func Fn_signal_handler_find(instance string, mask string, signalId string, detail string, closure string, func_ string, data string) {
}

func Fn_signal_handler_is_connected(instance string, handlerId string) {}

func Fn_signal_handler_unblock(instance string, handlerId string) {}

func Fn_signal_handlers_block_matched(instance string, mask string, signalId string, detail string, closure string, func_ string, data string) {
}

func Fn_signal_handlers_destroy(instance string) {}

func Fn_signal_handlers_disconnect_matched(instance string, mask string, signalId string, detail string, closure string, func_ string, data string) {
}

func Fn_signal_handlers_unblock_matched(instance string, mask string, signalId string, detail string, closure string, func_ string, data string) {
}

func Fn_signal_has_handler_pending(instance string, signalId string, detail string, mayBeBlocked string) {
}

func Fn_signal_list_ids(itype string, nIds string) {}

func Fn_signal_lookup(name string, itype string) {}

func Fn_signal_name(signalId string) {}

// UNSUPPORTED : signal_new : has varargs

// UNSUPPORTED : signal_new_class_handler : has varargs

func Fn_signal_new_valist(signalName string, itype string, signalFlags string, classClosure string, accumulator string, accuData string, cMarshaller string, returnType string, nParams string, args string) {
}

func Fn_signal_newv(signalName string, itype string, signalFlags string, classClosure string, accumulator string, accuData string, cMarshaller string, returnType string, nParams string, paramTypes string) {
}

func Fn_signal_override_class_closure(signalId string, instanceType string, classClosure string) {}

func Fn_signal_override_class_handler(signalName string, instanceType string, classHandler string) {}

func Fn_signal_parse_name(detailedSignal string, itype string, signalIdP string, detailP string, forceDetailQuark string) {
}

func Fn_signal_query(signalId string, query string) {}

func Fn_signal_remove_emission_hook(signalId string, hookId string) {}

func Fn_signal_stop_emission(instance string, signalId string, detail string) {}

func Fn_signal_stop_emission_by_name(instance string, detailedSignal string) {}

func Fn_signal_type_cclosure_new(itype string, structOffset string) {}

func Fn_source_set_closure(source string, closure string) {}

func Fn_source_set_dummy_callback(source string) {}

func Fn_strdup_value_contents(value string) {}

func Fn_type_add_class_cache_func(cacheData string, cacheFunc string) {}

func Fn_type_add_class_private(classType string, privateSize string) {}

func Fn_type_add_instance_private(classType string, privateSize string) {}

func Fn_type_add_interface_check(checkData string, checkFunc string) {}

func Fn_type_add_interface_dynamic(instanceType string, interfaceType string, plugin string) {}

func Fn_type_add_interface_static(instanceType string, interfaceType string, info string) {}

func Fn_type_check_class_cast(gClass string, isAType string) {}

func Fn_type_check_class_is_a(gClass string, isAType string) {}

func Fn_type_check_instance(instance string) {}

func Fn_type_check_instance_cast(instance string, ifaceType string) {}

func Fn_type_check_instance_is_a(instance string, ifaceType string) {}

func Fn_type_check_instance_is_fundamentally_a(instance string, fundamentalType string) {}

func Fn_type_check_is_value_type(type_ string) {}

func Fn_type_check_value(value string) {}

func Fn_type_check_value_holds(value string, type_ string) {}

func Fn_type_children(type_ string, nChildren string) {}

func Fn_type_class_adjust_private_offset(gClass string, privateSizeOrOffset string) {}

func Fn_type_class_peek(type_ string) {}

func Fn_type_class_peek_static(type_ string) {}

func Fn_type_class_ref(type_ string) {}

func Fn_type_create_instance(type_ string) {}

func Fn_type_default_interface_peek(gType string) {}

func Fn_type_default_interface_ref(gType string) {}

func Fn_type_default_interface_unref(gIface string) {}

func Fn_type_depth(type_ string) {}

func Fn_type_free_instance(instance string) {}

func Fn_type_from_name(name string) {}

func Fn_type_fundamental(typeId string) {}

func Fn_type_fundamental_next() {}

func Fn_type_get_plugin(type_ string) {}

func Fn_type_get_qdata(type_ string, quark string) {}

func Fn_type_init() {}

func Fn_type_init_with_debug_flags(debugFlags string) {}

func Fn_type_interface_add_prerequisite(interfaceType string, prerequisiteType string) {}

func Fn_type_interface_get_plugin(instanceType string, interfaceType string) {}

func Fn_type_interface_peek(instanceClass string, ifaceType string) {}

func Fn_type_interface_prerequisites(interfaceType string, nPrerequisites string) {}

func Fn_type_interfaces(type_ string, nInterfaces string) {}

func Fn_type_is_a(type_ string, isAType string) {}

func Fn_type_name(type_ string) {}

func Fn_type_name_from_class(gClass string) {}

func Fn_type_name_from_instance(instance string) {}

func Fn_type_next_base(leafType string, rootType string) {}

func Fn_type_parent(type_ string) {}

func Fn_type_qname(type_ string) {}

func Fn_type_query(type_ string, query string) {}

func Fn_type_register_dynamic(parentType string, typeName string, plugin string, flags string) {}

func Fn_type_register_fundamental(typeId string, typeName string, info string, finfo string, flags string) {
}

func Fn_type_register_static(parentType string, typeName string, info string, flags string) {}

func Fn_type_register_static_simple(parentType string, typeName string, classSize string, classInit string, instanceSize string, instanceInit string, flags string) {
}

func Fn_type_remove_class_cache_func(cacheData string, cacheFunc string) {}

func Fn_type_remove_interface_check(checkData string, checkFunc string) {}

func Fn_type_set_qdata(type_ string, quark string, data string) {}

func Fn_type_test_flags(type_ string, flags string) {}

func Fn_type_value_table_peek(type_ string) {}

func Fn_value_register_transform_func(srcType string, destType string, transformFunc string) {}

func Fn_value_type_compatible(srcType string, destType string) {}

func Fn_value_type_transformable(srcType string, destType string) {}

func Fn_bad() {}
