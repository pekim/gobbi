package gobject

// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// BoxedCopy is a wrapper around the C function g_boxed_copy.
func BoxedCopy(boxedType int, srcBoxed int) {}

// BoxedFree is a wrapper around the C function g_boxed_free.
func BoxedFree(boxedType int, boxed int) {}

// BoxedTypeRegisterStatic is a wrapper around the C function g_boxed_type_register_static.
func BoxedTypeRegisterStatic(name int, boxedCopy int, boxedFree int) {}

// CclosureMarshalBooleanBoxedBoxed is a wrapper around the C function g_cclosure_marshal_BOOLEAN__BOXED_BOXED.
func CclosureMarshalBooleanBoxedBoxed(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalBooleanFlags is a wrapper around the C function g_cclosure_marshal_BOOLEAN__FLAGS.
func CclosureMarshalBooleanFlags(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalStringObjectPointer is a wrapper around the C function g_cclosure_marshal_STRING__OBJECT_POINTER.
func CclosureMarshalStringObjectPointer(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalVoidBoolean is a wrapper around the C function g_cclosure_marshal_VOID__BOOLEAN.
func CclosureMarshalVoidBoolean(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalVoidBoxed is a wrapper around the C function g_cclosure_marshal_VOID__BOXED.
func CclosureMarshalVoidBoxed(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalVoidChar is a wrapper around the C function g_cclosure_marshal_VOID__CHAR.
func CclosureMarshalVoidChar(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalVoidDouble is a wrapper around the C function g_cclosure_marshal_VOID__DOUBLE.
func CclosureMarshalVoidDouble(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalVoidEnum is a wrapper around the C function g_cclosure_marshal_VOID__ENUM.
func CclosureMarshalVoidEnum(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalVoidFlags is a wrapper around the C function g_cclosure_marshal_VOID__FLAGS.
func CclosureMarshalVoidFlags(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalVoidFloat is a wrapper around the C function g_cclosure_marshal_VOID__FLOAT.
func CclosureMarshalVoidFloat(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalVoidInt is a wrapper around the C function g_cclosure_marshal_VOID__INT.
func CclosureMarshalVoidInt(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalVoidLong is a wrapper around the C function g_cclosure_marshal_VOID__LONG.
func CclosureMarshalVoidLong(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalVoidObject is a wrapper around the C function g_cclosure_marshal_VOID__OBJECT.
func CclosureMarshalVoidObject(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalVoidParam is a wrapper around the C function g_cclosure_marshal_VOID__PARAM.
func CclosureMarshalVoidParam(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalVoidPointer is a wrapper around the C function g_cclosure_marshal_VOID__POINTER.
func CclosureMarshalVoidPointer(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalVoidString is a wrapper around the C function g_cclosure_marshal_VOID__STRING.
func CclosureMarshalVoidString(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalVoidUchar is a wrapper around the C function g_cclosure_marshal_VOID__UCHAR.
func CclosureMarshalVoidUchar(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalVoidUint is a wrapper around the C function g_cclosure_marshal_VOID__UINT.
func CclosureMarshalVoidUint(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalVoidUintPointer is a wrapper around the C function g_cclosure_marshal_VOID__UINT_POINTER.
func CclosureMarshalVoidUintPointer(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalVoidUlong is a wrapper around the C function g_cclosure_marshal_VOID__ULONG.
func CclosureMarshalVoidUlong(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalVoidVariant is a wrapper around the C function g_cclosure_marshal_VOID__VARIANT.
func CclosureMarshalVoidVariant(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// CclosureMarshalVoidVoid is a wrapper around the C function g_cclosure_marshal_VOID__VOID.
func CclosureMarshalVoidVoid(closure int, returnValue int, nParamValues int, paramValues int, invocationHint int, marshalData int) {
}

// Blacklisted function: g_cclosure_new

// CclosureNewObject is a wrapper around the C function g_cclosure_new_object.
func CclosureNewObject(callbackFunc int, object int) {}

// CclosureNewObjectSwap is a wrapper around the C function g_cclosure_new_object_swap.
func CclosureNewObjectSwap(callbackFunc int, object int) {}

// Blacklisted function: g_cclosure_new_swap

// EnumCompleteTypeInfo is a wrapper around the C function g_enum_complete_type_info.
func EnumCompleteTypeInfo(gEnumType int, info int, constValues int) {}

// EnumGetValue is a wrapper around the C function g_enum_get_value.
func EnumGetValue(enumClass int, value int) {}

// EnumGetValueByName is a wrapper around the C function g_enum_get_value_by_name.
func EnumGetValueByName(enumClass int, name int) {}

// EnumGetValueByNick is a wrapper around the C function g_enum_get_value_by_nick.
func EnumGetValueByNick(enumClass int, nick int) {}

// EnumRegisterStatic is a wrapper around the C function g_enum_register_static.
func EnumRegisterStatic(name int, constStaticValues int) {}

// FlagsCompleteTypeInfo is a wrapper around the C function g_flags_complete_type_info.
func FlagsCompleteTypeInfo(gFlagsType int, info int, constValues int) {}

// FlagsGetFirstValue is a wrapper around the C function g_flags_get_first_value.
func FlagsGetFirstValue(flagsClass int, value int) {}

// FlagsGetValueByName is a wrapper around the C function g_flags_get_value_by_name.
func FlagsGetValueByName(flagsClass int, name int) {}

// FlagsGetValueByNick is a wrapper around the C function g_flags_get_value_by_nick.
func FlagsGetValueByNick(flagsClass int, nick int) {}

// FlagsRegisterStatic is a wrapper around the C function g_flags_register_static.
func FlagsRegisterStatic(name int, constStaticValues int) {}

// GtypeGetType is a wrapper around the C function g_gtype_get_type.
func GtypeGetType() {}

// ParamSpecBoolean is a wrapper around the C function g_param_spec_boolean.
func ParamSpecBoolean(name int, nick int, blurb int, defaultValue int, flags int) {}

// ParamSpecBoxed is a wrapper around the C function g_param_spec_boxed.
func ParamSpecBoxed(name int, nick int, blurb int, boxedType int, flags int) {}

// ParamSpecChar is a wrapper around the C function g_param_spec_char.
func ParamSpecChar(name int, nick int, blurb int, minimum int, maximum int, defaultValue int, flags int) {}

// ParamSpecDouble is a wrapper around the C function g_param_spec_double.
func ParamSpecDouble(name int, nick int, blurb int, minimum int, maximum int, defaultValue int, flags int) {
}

// ParamSpecEnum is a wrapper around the C function g_param_spec_enum.
func ParamSpecEnum(name int, nick int, blurb int, enumType int, defaultValue int, flags int) {}

// ParamSpecFlags is a wrapper around the C function g_param_spec_flags.
func ParamSpecFlags(name int, nick int, blurb int, flagsType int, defaultValue int, flags int) {}

// ParamSpecFloat is a wrapper around the C function g_param_spec_float.
func ParamSpecFloat(name int, nick int, blurb int, minimum int, maximum int, defaultValue int, flags int) {}

// ParamSpecInt is a wrapper around the C function g_param_spec_int.
func ParamSpecInt(name int, nick int, blurb int, minimum int, maximum int, defaultValue int, flags int) {}

// ParamSpecInt64 is a wrapper around the C function g_param_spec_int64.
func ParamSpecInt64(name int, nick int, blurb int, minimum int, maximum int, defaultValue int, flags int) {}

// ParamSpecLong is a wrapper around the C function g_param_spec_long.
func ParamSpecLong(name int, nick int, blurb int, minimum int, maximum int, defaultValue int, flags int) {}

// ParamSpecObject is a wrapper around the C function g_param_spec_object.
func ParamSpecObject(name int, nick int, blurb int, objectType int, flags int) {}

// ParamSpecParam is a wrapper around the C function g_param_spec_param.
func ParamSpecParam(name int, nick int, blurb int, paramType int, flags int) {}

// ParamSpecPointer is a wrapper around the C function g_param_spec_pointer.
func ParamSpecPointer(name int, nick int, blurb int, flags int) {}

// ParamSpecPoolNew is a wrapper around the C function g_param_spec_pool_new.
func ParamSpecPoolNew(typePrefixing int) {}

// ParamSpecString is a wrapper around the C function g_param_spec_string.
func ParamSpecString(name int, nick int, blurb int, defaultValue int, flags int) {}

// ParamSpecUchar is a wrapper around the C function g_param_spec_uchar.
func ParamSpecUchar(name int, nick int, blurb int, minimum int, maximum int, defaultValue int, flags int) {}

// ParamSpecUint is a wrapper around the C function g_param_spec_uint.
func ParamSpecUint(name int, nick int, blurb int, minimum int, maximum int, defaultValue int, flags int) {}

// ParamSpecUint64 is a wrapper around the C function g_param_spec_uint64.
func ParamSpecUint64(name int, nick int, blurb int, minimum int, maximum int, defaultValue int, flags int) {
}

// ParamSpecUlong is a wrapper around the C function g_param_spec_ulong.
func ParamSpecUlong(name int, nick int, blurb int, minimum int, maximum int, defaultValue int, flags int) {}

// ParamSpecUnichar is a wrapper around the C function g_param_spec_unichar.
func ParamSpecUnichar(name int, nick int, blurb int, defaultValue int, flags int) {}

// ParamSpecValueArray is a wrapper around the C function g_param_spec_value_array.
func ParamSpecValueArray(name int, nick int, blurb int, elementSpec int, flags int) {}

// ParamTypeRegisterStatic is a wrapper around the C function g_param_type_register_static.
func ParamTypeRegisterStatic(name int, pspecInfo int) {}

// ParamValueConvert is a wrapper around the C function g_param_value_convert.
func ParamValueConvert(pspec int, srcValue int, destValue int, strictValidation int) {}

// ParamValueDefaults is a wrapper around the C function g_param_value_defaults.
func ParamValueDefaults(pspec int, value int) {}

// ParamValueSetDefault is a wrapper around the C function g_param_value_set_default.
func ParamValueSetDefault(pspec int, value int) {}

// ParamValueValidate is a wrapper around the C function g_param_value_validate.
func ParamValueValidate(pspec int, value int) {}

// ParamValuesCmp is a wrapper around the C function g_param_values_cmp.
func ParamValuesCmp(pspec int, value1 int, value2 int) {}

// PointerTypeRegisterStatic is a wrapper around the C function g_pointer_type_register_static.
func PointerTypeRegisterStatic(name int) {}

// SignalAddEmissionHook is a wrapper around the C function g_signal_add_emission_hook.
func SignalAddEmissionHook(signalId int, detail int, hookFunc int, hookData int, dataDestroy int) {}

// SignalChainFromOverridden is a wrapper around the C function g_signal_chain_from_overridden.
func SignalChainFromOverridden(instanceAndParams int, returnValue int) {}

// SignalConnectClosure is a wrapper around the C function g_signal_connect_closure.
func SignalConnectClosure(instance int, detailedSignal int, closure int, after int) {}

// SignalConnectClosureById is a wrapper around the C function g_signal_connect_closure_by_id.
func SignalConnectClosureById(instance int, signalId int, detail int, closure int, after int) {}

// SignalConnectData is a wrapper around the C function g_signal_connect_data.
func SignalConnectData(instance int, detailedSignal int, cHandler int, data int, destroyData int, connectFlags int) {
}

// SignalConnectObject is a wrapper around the C function g_signal_connect_object.
func SignalConnectObject(instance int, detailedSignal int, cHandler int, gobject int, connectFlags int) {}

// Unsupported function: g_signal_emit : unsupported parameter ... : varargs

// Unsupported function: g_signal_emit_by_name : unsupported parameter ... : varargs

// SignalEmitValist is a wrapper around the C function g_signal_emit_valist.
func SignalEmitValist(instance int, signalId int, detail int, varArgs int) {}

// SignalEmitv is a wrapper around the C function g_signal_emitv.
func SignalEmitv(instanceAndParams int, signalId int, detail int, returnValue int) {}

// SignalGetInvocationHint is a wrapper around the C function g_signal_get_invocation_hint.
func SignalGetInvocationHint(instance int) {}

// SignalHandlerBlock is a wrapper around the C function g_signal_handler_block.
func SignalHandlerBlock(instance int, handlerId int) {}

// SignalHandlerDisconnect is a wrapper around the C function g_signal_handler_disconnect.
func SignalHandlerDisconnect(instance int, handlerId int) {}

// SignalHandlerFind is a wrapper around the C function g_signal_handler_find.
func SignalHandlerFind(instance int, mask int, signalId int, detail int, closure int, func_ int, data int) {
}

// SignalHandlerIsConnected is a wrapper around the C function g_signal_handler_is_connected.
func SignalHandlerIsConnected(instance int, handlerId int) {}

// SignalHandlerUnblock is a wrapper around the C function g_signal_handler_unblock.
func SignalHandlerUnblock(instance int, handlerId int) {}

// SignalHandlersBlockMatched is a wrapper around the C function g_signal_handlers_block_matched.
func SignalHandlersBlockMatched(instance int, mask int, signalId int, detail int, closure int, func_ int, data int) {
}

// SignalHandlersDestroy is a wrapper around the C function g_signal_handlers_destroy.
func SignalHandlersDestroy(instance int) {}

// SignalHandlersDisconnectMatched is a wrapper around the C function g_signal_handlers_disconnect_matched.
func SignalHandlersDisconnectMatched(instance int, mask int, signalId int, detail int, closure int, func_ int, data int) {
}

// SignalHandlersUnblockMatched is a wrapper around the C function g_signal_handlers_unblock_matched.
func SignalHandlersUnblockMatched(instance int, mask int, signalId int, detail int, closure int, func_ int, data int) {
}

// SignalHasHandlerPending is a wrapper around the C function g_signal_has_handler_pending.
func SignalHasHandlerPending(instance int, signalId int, detail int, mayBeBlocked int) {}

// SignalListIds is a wrapper around the C function g_signal_list_ids.
func SignalListIds(itype int, nIds int) {}

// SignalLookup is a wrapper around the C function g_signal_lookup.
func SignalLookup(name int, itype int) {}

// SignalName is a wrapper around the C function g_signal_name.
func SignalName(signalId int) {}

// Unsupported function: g_signal_new : unsupported parameter ... : varargs

// SignalNewValist is a wrapper around the C function g_signal_new_valist.
func SignalNewValist(signalName int, itype int, signalFlags int, classClosure int, accumulator int, accuData int, cMarshaller int, returnType int, nParams int, args int) {
}

// SignalNewv is a wrapper around the C function g_signal_newv.
func SignalNewv(signalName int, itype int, signalFlags int, classClosure int, accumulator int, accuData int, cMarshaller int, returnType int, nParams int, paramTypes int) {
}

// SignalOverrideClassClosure is a wrapper around the C function g_signal_override_class_closure.
func SignalOverrideClassClosure(signalId int, instanceType int, classClosure int) {}

// SignalParseName is a wrapper around the C function g_signal_parse_name.
func SignalParseName(detailedSignal int, itype int, signalIdP int, detailP int, forceDetailQuark int) {}

// SignalQuery is a wrapper around the C function g_signal_query.
func SignalQuery(signalId int, query int) {}

// SignalRemoveEmissionHook is a wrapper around the C function g_signal_remove_emission_hook.
func SignalRemoveEmissionHook(signalId int, hookId int) {}

// SignalStopEmission is a wrapper around the C function g_signal_stop_emission.
func SignalStopEmission(instance int, signalId int, detail int) {}

// SignalStopEmissionByName is a wrapper around the C function g_signal_stop_emission_by_name.
func SignalStopEmissionByName(instance int, detailedSignal int) {}

// SignalTypeCclosureNew is a wrapper around the C function g_signal_type_cclosure_new.
func SignalTypeCclosureNew(itype int, structOffset int) {}

// SourceSetClosure is a wrapper around the C function g_source_set_closure.
func SourceSetClosure(source int, closure int) {}

// SourceSetDummyCallback is a wrapper around the C function g_source_set_dummy_callback.
func SourceSetDummyCallback(source int) {}

// StrdupValueContents is a wrapper around the C function g_strdup_value_contents.
func StrdupValueContents(value int) {}

// Blacklisted function: g_type_add_class_cache_func

// TypeAddInstancePrivate is a wrapper around the C function g_type_add_instance_private.
func TypeAddInstancePrivate(classType int, privateSize int) {}

// TypeAddInterfaceDynamic is a wrapper around the C function g_type_add_interface_dynamic.
func TypeAddInterfaceDynamic(instanceType int, interfaceType int, plugin int) {}

// TypeAddInterfaceStatic is a wrapper around the C function g_type_add_interface_static.
func TypeAddInterfaceStatic(instanceType int, interfaceType int, info int) {}

// TypeCheckClassCast is a wrapper around the C function g_type_check_class_cast.
func TypeCheckClassCast(gClass int, isAType int) {}

// TypeCheckClassIsA is a wrapper around the C function g_type_check_class_is_a.
func TypeCheckClassIsA(gClass int, isAType int) {}

// TypeCheckInstance is a wrapper around the C function g_type_check_instance.
func TypeCheckInstance(instance int) {}

// TypeCheckInstanceCast is a wrapper around the C function g_type_check_instance_cast.
func TypeCheckInstanceCast(instance int, ifaceType int) {}

// TypeCheckInstanceIsA is a wrapper around the C function g_type_check_instance_is_a.
func TypeCheckInstanceIsA(instance int, ifaceType int) {}

// TypeCheckInstanceIsFundamentallyA is a wrapper around the C function g_type_check_instance_is_fundamentally_a.
func TypeCheckInstanceIsFundamentallyA(instance int, fundamentalType int) {}

// TypeCheckIsValueType is a wrapper around the C function g_type_check_is_value_type.
func TypeCheckIsValueType(type_ int) {}

// TypeCheckValue is a wrapper around the C function g_type_check_value.
func TypeCheckValue(value int) {}

// TypeCheckValueHolds is a wrapper around the C function g_type_check_value_holds.
func TypeCheckValueHolds(value int, type_ int) {}

// TypeChildren is a wrapper around the C function g_type_children.
func TypeChildren(type_ int, nChildren int) {}

// TypeClassAdjustPrivateOffset is a wrapper around the C function g_type_class_adjust_private_offset.
func TypeClassAdjustPrivateOffset(gClass int, privateSizeOrOffset int) {}

// TypeClassPeek is a wrapper around the C function g_type_class_peek.
func TypeClassPeek(type_ int) {}

// TypeClassRef is a wrapper around the C function g_type_class_ref.
func TypeClassRef(type_ int) {}

// TypeCreateInstance is a wrapper around the C function g_type_create_instance.
func TypeCreateInstance(type_ int) {}

// TypeDepth is a wrapper around the C function g_type_depth.
func TypeDepth(type_ int) {}

// TypeFreeInstance is a wrapper around the C function g_type_free_instance.
func TypeFreeInstance(instance int) {}

// TypeFromName is a wrapper around the C function g_type_from_name.
func TypeFromName(name int) {}

// TypeFundamental is a wrapper around the C function g_type_fundamental.
func TypeFundamental(typeId int) {}

// TypeFundamentalNext is a wrapper around the C function g_type_fundamental_next.
func TypeFundamentalNext() {}

// TypeGetPlugin is a wrapper around the C function g_type_get_plugin.
func TypeGetPlugin(type_ int) {}

// TypeGetQdata is a wrapper around the C function g_type_get_qdata.
func TypeGetQdata(type_ int, quark int) {}

// TypeInit is a wrapper around the C function g_type_init.
func TypeInit() {}

// TypeInitWithDebugFlags is a wrapper around the C function g_type_init_with_debug_flags.
func TypeInitWithDebugFlags(debugFlags int) {}

// TypeInterfaceAddPrerequisite is a wrapper around the C function g_type_interface_add_prerequisite.
func TypeInterfaceAddPrerequisite(interfaceType int, prerequisiteType int) {}

// TypeInterfaceGetPlugin is a wrapper around the C function g_type_interface_get_plugin.
func TypeInterfaceGetPlugin(instanceType int, interfaceType int) {}

// TypeInterfacePeek is a wrapper around the C function g_type_interface_peek.
func TypeInterfacePeek(instanceClass int, ifaceType int) {}

// TypeInterfaces is a wrapper around the C function g_type_interfaces.
func TypeInterfaces(type_ int, nInterfaces int) {}

// TypeIsA is a wrapper around the C function g_type_is_a.
func TypeIsA(type_ int, isAType int) {}

// TypeName is a wrapper around the C function g_type_name.
func TypeName(type_ int) {}

// TypeNameFromClass is a wrapper around the C function g_type_name_from_class.
func TypeNameFromClass(gClass int) {}

// TypeNameFromInstance is a wrapper around the C function g_type_name_from_instance.
func TypeNameFromInstance(instance int) {}

// TypeNextBase is a wrapper around the C function g_type_next_base.
func TypeNextBase(leafType int, rootType int) {}

// TypeParent is a wrapper around the C function g_type_parent.
func TypeParent(type_ int) {}

// TypeQname is a wrapper around the C function g_type_qname.
func TypeQname(type_ int) {}

// TypeQuery is a wrapper around the C function g_type_query.
func TypeQuery(type_ int, query int) {}

// TypeRegisterDynamic is a wrapper around the C function g_type_register_dynamic.
func TypeRegisterDynamic(parentType int, typeName int, plugin int, flags int) {}

// TypeRegisterFundamental is a wrapper around the C function g_type_register_fundamental.
func TypeRegisterFundamental(typeId int, typeName int, info int, finfo int, flags int) {}

// TypeRegisterStatic is a wrapper around the C function g_type_register_static.
func TypeRegisterStatic(parentType int, typeName int, info int, flags int) {}

// Blacklisted function: g_type_remove_class_cache_func

// TypeSetQdata is a wrapper around the C function g_type_set_qdata.
func TypeSetQdata(type_ int, quark int, data int) {}

// TypeTestFlags is a wrapper around the C function g_type_test_flags.
func TypeTestFlags(type_ int, flags int) {}

// TypeValueTablePeek is a wrapper around the C function g_type_value_table_peek.
func TypeValueTablePeek(type_ int) {}

// ValueRegisterTransformFunc is a wrapper around the C function g_value_register_transform_func.
func ValueRegisterTransformFunc(srcType int, destType int, transformFunc int) {}

// ValueTypeCompatible is a wrapper around the C function g_value_type_compatible.
func ValueTypeCompatible(srcType int, destType int) {}

// ValueTypeTransformable is a wrapper around the C function g_value_type_transformable.
func ValueTypeTransformable(srcType int, destType int) {}
