// Code generated - DO NOT EDIT.
// +build gobject_2.6

package gobject

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// UNSUPPORTED : SignalCMarshaller : blacklisted

// UNSUPPORTED : SignalCVaMarshaller : blacklisted

// Type is a representation of the C alias GType.
type Type uint64

// PARAM_MASK is a representation of the C constant G_PARAM_MASK.
const PARAM_MASK = 255

// PARAM_STATIC_STRINGS is a representation of the C constant G_PARAM_STATIC_STRINGS.
const PARAM_STATIC_STRINGS = 224

// PARAM_USER_SHIFT is a representation of the C constant G_PARAM_USER_SHIFT.
const PARAM_USER_SHIFT = 8

// SIGNAL_FLAGS_MASK is a representation of the C constant G_SIGNAL_FLAGS_MASK.
const SIGNAL_FLAGS_MASK = 511

// SIGNAL_MATCH_MASK is a representation of the C constant G_SIGNAL_MATCH_MASK.
const SIGNAL_MATCH_MASK = 63

// UNSUPPORTED the alias TYPE_FLAG_RESERVED_ID_BIT is qualified, 'GLib.Type'

// TYPE_FUNDAMENTAL_MAX is a representation of the C constant G_TYPE_FUNDAMENTAL_MAX.
const TYPE_FUNDAMENTAL_MAX = 255

// TYPE_FUNDAMENTAL_SHIFT is a representation of the C constant G_TYPE_FUNDAMENTAL_SHIFT.
const TYPE_FUNDAMENTAL_SHIFT = 2

// TYPE_RESERVED_BSE_FIRST is a representation of the C constant G_TYPE_RESERVED_BSE_FIRST.
const TYPE_RESERVED_BSE_FIRST = 32

// TYPE_RESERVED_BSE_LAST is a representation of the C constant G_TYPE_RESERVED_BSE_LAST.
const TYPE_RESERVED_BSE_LAST = 48

// TYPE_RESERVED_GLIB_FIRST is a representation of the C constant G_TYPE_RESERVED_GLIB_FIRST.
const TYPE_RESERVED_GLIB_FIRST = 22

// TYPE_RESERVED_GLIB_LAST is a representation of the C constant G_TYPE_RESERVED_GLIB_LAST.
const TYPE_RESERVED_GLIB_LAST = 31

// TYPE_RESERVED_USER_FIRST is a representation of the C constant G_TYPE_RESERVED_USER_FIRST.
const TYPE_RESERVED_USER_FIRST = 49

// VALUE_NOCOPY_CONTENTS is a representation of the C constant G_VALUE_NOCOPY_CONTENTS.
const VALUE_NOCOPY_CONTENTS = 134217728

// ConnectFlags is a representation of the C bitfield GConnectFlags.
type ConnectFlags int

// ConnectFlags_after is a representation of the C bitfield member G_CONNECT_AFTER.
const ConnectFlags_after = ConnectFlags(1)

// ConnectFlags_swapped is a representation of the C bitfield member G_CONNECT_SWAPPED.
const ConnectFlags_swapped = ConnectFlags(2)

// ParamFlags is a representation of the C bitfield GParamFlags.
type ParamFlags int

// ParamFlags_readable is a representation of the C bitfield member G_PARAM_READABLE.
const ParamFlags_readable = ParamFlags(1)

// ParamFlags_writable is a representation of the C bitfield member G_PARAM_WRITABLE.
const ParamFlags_writable = ParamFlags(2)

// ParamFlags_readwrite is a representation of the C bitfield member G_PARAM_READWRITE.
const ParamFlags_readwrite = ParamFlags(3)

// ParamFlags_construct is a representation of the C bitfield member G_PARAM_CONSTRUCT.
const ParamFlags_construct = ParamFlags(4)

// ParamFlags_construct_only is a representation of the C bitfield member G_PARAM_CONSTRUCT_ONLY.
const ParamFlags_construct_only = ParamFlags(8)

// ParamFlags_lax_validation is a representation of the C bitfield member G_PARAM_LAX_VALIDATION.
const ParamFlags_lax_validation = ParamFlags(16)

// ParamFlags_static_name is a representation of the C bitfield member G_PARAM_STATIC_NAME.
const ParamFlags_static_name = ParamFlags(32)

// ParamFlags_private is a representation of the C bitfield member G_PARAM_PRIVATE.
const ParamFlags_private = ParamFlags(32)

// ParamFlags_static_nick is a representation of the C bitfield member G_PARAM_STATIC_NICK.
const ParamFlags_static_nick = ParamFlags(64)

// ParamFlags_static_blurb is a representation of the C bitfield member G_PARAM_STATIC_BLURB.
const ParamFlags_static_blurb = ParamFlags(128)

// ParamFlags_explicit_notify is a representation of the C bitfield member G_PARAM_EXPLICIT_NOTIFY.
const ParamFlags_explicit_notify = ParamFlags(1073741824)

// ParamFlags_deprecated is a representation of the C bitfield member G_PARAM_DEPRECATED.
const ParamFlags_deprecated = ParamFlags(2147483648)

// SignalFlags is a representation of the C bitfield GSignalFlags.
type SignalFlags int

// SignalFlags_run_first is a representation of the C bitfield member G_SIGNAL_RUN_FIRST.
const SignalFlags_run_first = SignalFlags(1)

// SignalFlags_run_last is a representation of the C bitfield member G_SIGNAL_RUN_LAST.
const SignalFlags_run_last = SignalFlags(2)

// SignalFlags_run_cleanup is a representation of the C bitfield member G_SIGNAL_RUN_CLEANUP.
const SignalFlags_run_cleanup = SignalFlags(4)

// SignalFlags_no_recurse is a representation of the C bitfield member G_SIGNAL_NO_RECURSE.
const SignalFlags_no_recurse = SignalFlags(8)

// SignalFlags_detailed is a representation of the C bitfield member G_SIGNAL_DETAILED.
const SignalFlags_detailed = SignalFlags(16)

// SignalFlags_action is a representation of the C bitfield member G_SIGNAL_ACTION.
const SignalFlags_action = SignalFlags(32)

// SignalFlags_no_hooks is a representation of the C bitfield member G_SIGNAL_NO_HOOKS.
const SignalFlags_no_hooks = SignalFlags(64)

// SignalFlags_must_collect is a representation of the C bitfield member G_SIGNAL_MUST_COLLECT.
const SignalFlags_must_collect = SignalFlags(128)

// SignalFlags_deprecated is a representation of the C bitfield member G_SIGNAL_DEPRECATED.
const SignalFlags_deprecated = SignalFlags(256)

// SignalMatchType is a representation of the C bitfield GSignalMatchType.
type SignalMatchType int

// SignalMatchType_id is a representation of the C bitfield member G_SIGNAL_MATCH_ID.
const SignalMatchType_id = SignalMatchType(1)

// SignalMatchType_detail is a representation of the C bitfield member G_SIGNAL_MATCH_DETAIL.
const SignalMatchType_detail = SignalMatchType(2)

// SignalMatchType_closure is a representation of the C bitfield member G_SIGNAL_MATCH_CLOSURE.
const SignalMatchType_closure = SignalMatchType(4)

// SignalMatchType_func is a representation of the C bitfield member G_SIGNAL_MATCH_FUNC.
const SignalMatchType_func = SignalMatchType(8)

// SignalMatchType_data is a representation of the C bitfield member G_SIGNAL_MATCH_DATA.
const SignalMatchType_data = SignalMatchType(16)

// SignalMatchType_unblocked is a representation of the C bitfield member G_SIGNAL_MATCH_UNBLOCKED.
const SignalMatchType_unblocked = SignalMatchType(32)

// TypeDebugFlags is a representation of the C bitfield GTypeDebugFlags.
type TypeDebugFlags int

// TypeDebugFlags_none is a representation of the C bitfield member G_TYPE_DEBUG_NONE.
const TypeDebugFlags_none = TypeDebugFlags(0)

// TypeDebugFlags_objects is a representation of the C bitfield member G_TYPE_DEBUG_OBJECTS.
const TypeDebugFlags_objects = TypeDebugFlags(1)

// TypeDebugFlags_signals is a representation of the C bitfield member G_TYPE_DEBUG_SIGNALS.
const TypeDebugFlags_signals = TypeDebugFlags(2)

// TypeDebugFlags_instance_count is a representation of the C bitfield member G_TYPE_DEBUG_INSTANCE_COUNT.
const TypeDebugFlags_instance_count = TypeDebugFlags(4)

// TypeDebugFlags_mask is a representation of the C bitfield member G_TYPE_DEBUG_MASK.
const TypeDebugFlags_mask = TypeDebugFlags(7)

// TypeFlags is a representation of the C bitfield GTypeFlags.
type TypeFlags int

// TypeFlags_abstract is a representation of the C bitfield member G_TYPE_FLAG_ABSTRACT.
const TypeFlags_abstract = TypeFlags(16)

// TypeFlags_value_abstract is a representation of the C bitfield member G_TYPE_FLAG_VALUE_ABSTRACT.
const TypeFlags_value_abstract = TypeFlags(32)

// TypeFundamentalFlags is a representation of the C bitfield GTypeFundamentalFlags.
type TypeFundamentalFlags int

// TypeFundamentalFlags_classed is a representation of the C bitfield member G_TYPE_FLAG_CLASSED.
const TypeFundamentalFlags_classed = TypeFundamentalFlags(1)

// TypeFundamentalFlags_instantiatable is a representation of the C bitfield member G_TYPE_FLAG_INSTANTIATABLE.
const TypeFundamentalFlags_instantiatable = TypeFundamentalFlags(2)

// TypeFundamentalFlags_derivable is a representation of the C bitfield member G_TYPE_FLAG_DERIVABLE.
const TypeFundamentalFlags_derivable = TypeFundamentalFlags(4)

// TypeFundamentalFlags_deep_derivable is a representation of the C bitfield member G_TYPE_FLAG_DEEP_DERIVABLE.
const TypeFundamentalFlags_deep_derivable = TypeFundamentalFlags(8)

// BoxedCopy is analogous to the C function g_boxed_copy.
func BoxedCopy(boxedType uint64, srcBoxed unsafe.Pointer) {
	sys_boxedType := boxedType
	sys_srcBoxed := srcBoxed
}

// BoxedFree is analogous to the C function g_boxed_free.
func BoxedFree(boxedType uint64, boxed unsafe.Pointer) {
	sys_boxedType := boxedType
	sys_boxed := boxed
}

// UNSUPPORTED : g_boxed_type_register_static : parameter 'boxed_copy' is callback

// UNSUPPORTED : g_cclosure_new : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object_swap : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_swap : parameter 'callback_func' is callback

// EnumCompleteTypeInfo is analogous to the C function g_enum_complete_type_info.
func EnumCompleteTypeInfo(gEnumType uint64, constValues *EnumValue) {
	sys_gEnumType := gEnumType
	sys_constValues := constValues.ToC()
}

// EnumGetValue is analogous to the C function g_enum_get_value.
func EnumGetValue(enumClass *EnumClass, value int) {
	sys_enumClass := enumClass.ToC()
	sys_value := value
}

// EnumGetValueByName is analogous to the C function g_enum_get_value_by_name.
func EnumGetValueByName(enumClass *EnumClass, name string) {
	sys_enumClass := enumClass.ToC()
	sys_name := name
}

// EnumGetValueByNick is analogous to the C function g_enum_get_value_by_nick.
func EnumGetValueByNick(enumClass *EnumClass, nick string) {
	sys_enumClass := enumClass.ToC()
	sys_nick := nick
}

// EnumRegisterStatic is analogous to the C function g_enum_register_static.
func EnumRegisterStatic(name string, constStaticValues *EnumValue) {
	sys_name := name
	sys_constStaticValues := constStaticValues.ToC()
}

// FlagsCompleteTypeInfo is analogous to the C function g_flags_complete_type_info.
func FlagsCompleteTypeInfo(gFlagsType uint64, constValues *FlagsValue) {
	sys_gFlagsType := gFlagsType
	sys_constValues := constValues.ToC()
}

// FlagsGetFirstValue is analogous to the C function g_flags_get_first_value.
func FlagsGetFirstValue(flagsClass *FlagsClass, value uint) {
	sys_flagsClass := flagsClass.ToC()
	sys_value := value
}

// FlagsGetValueByName is analogous to the C function g_flags_get_value_by_name.
func FlagsGetValueByName(flagsClass *FlagsClass, name string) {
	sys_flagsClass := flagsClass.ToC()
	sys_name := name
}

// FlagsGetValueByNick is analogous to the C function g_flags_get_value_by_nick.
func FlagsGetValueByNick(flagsClass *FlagsClass, nick string) {
	sys_flagsClass := flagsClass.ToC()
	sys_nick := nick
}

// FlagsRegisterStatic is analogous to the C function g_flags_register_static.
func FlagsRegisterStatic(name string, constStaticValues *FlagsValue) {
	sys_name := name
	sys_constStaticValues := constStaticValues.ToC()
}

// GtypeGetType is analogous to the C function g_gtype_get_type.
func GtypeGetType() {}

// ParamSpecBoolean_ is analogous to the C function g_param_spec_boolean.
func ParamSpecBoolean_(name string, nick string, blurb string, defaultValue bool, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_defaultValue := defaultValue
	sys_flags := flags
}

// ParamSpecBoxed_ is analogous to the C function g_param_spec_boxed.
func ParamSpecBoxed_(name string, nick string, blurb string, boxedType uint64, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_boxedType := boxedType
	sys_flags := flags
}

// ParamSpecChar_ is analogous to the C function g_param_spec_char.
func ParamSpecChar_(name string, nick string, blurb string, minimum int8, maximum int8, defaultValue int8, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
}

// ParamSpecDouble_ is analogous to the C function g_param_spec_double.
func ParamSpecDouble_(name string, nick string, blurb string, minimum float64, maximum float64, defaultValue float64, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
}

// ParamSpecEnum_ is analogous to the C function g_param_spec_enum.
func ParamSpecEnum_(name string, nick string, blurb string, enumType uint64, defaultValue int, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_enumType := enumType
	sys_defaultValue := defaultValue
	sys_flags := flags
}

// ParamSpecFlags_ is analogous to the C function g_param_spec_flags.
func ParamSpecFlags_(name string, nick string, blurb string, flagsType uint64, defaultValue uint, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_flagsType := flagsType
	sys_defaultValue := defaultValue
	sys_flags := flags
}

// ParamSpecFloat_ is analogous to the C function g_param_spec_float.
func ParamSpecFloat_(name string, nick string, blurb string, minimum float32, maximum float32, defaultValue float32, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
}

// ParamSpecInt_ is analogous to the C function g_param_spec_int.
func ParamSpecInt_(name string, nick string, blurb string, minimum int, maximum int, defaultValue int, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
}

// ParamSpecInt64_ is analogous to the C function g_param_spec_int64.
func ParamSpecInt64_(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
}

// ParamSpecLong_ is analogous to the C function g_param_spec_long.
func ParamSpecLong_(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
}

// ParamSpecObject_ is analogous to the C function g_param_spec_object.
func ParamSpecObject_(name string, nick string, blurb string, objectType uint64, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_objectType := objectType
	sys_flags := flags
}

// ParamSpecOverride_ is analogous to the C function g_param_spec_override.
func ParamSpecOverride_(name string, overridden *ParamSpec) {
	sys_name := name
	sys_overridden := overridden.ToC()
}

// ParamSpecParam_ is analogous to the C function g_param_spec_param.
func ParamSpecParam_(name string, nick string, blurb string, paramType uint64, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_paramType := paramType
	sys_flags := flags
}

// ParamSpecPointer_ is analogous to the C function g_param_spec_pointer.
func ParamSpecPointer_(name string, nick string, blurb string, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_flags := flags
}

// ParamSpecString_ is analogous to the C function g_param_spec_string.
func ParamSpecString_(name string, nick string, blurb string, defaultValue string, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_defaultValue := defaultValue
	sys_flags := flags
}

// ParamSpecUchar is analogous to the C function g_param_spec_uchar.
func ParamSpecUchar(name string, nick string, blurb string, minimum uint8, maximum uint8, defaultValue uint8, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
}

// ParamSpecUint is analogous to the C function g_param_spec_uint.
func ParamSpecUint(name string, nick string, blurb string, minimum uint, maximum uint, defaultValue uint, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
}

// ParamSpecUint64 is analogous to the C function g_param_spec_uint64.
func ParamSpecUint64(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
}

// ParamSpecUlong is analogous to the C function g_param_spec_ulong.
func ParamSpecUlong(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
}

// ParamSpecUnichar_ is analogous to the C function g_param_spec_unichar.
func ParamSpecUnichar_(name string, nick string, blurb string, defaultValue rune, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_defaultValue := defaultValue
	sys_flags := flags
}

// ParamSpecValueArray_ is analogous to the C function g_param_spec_value_array.
func ParamSpecValueArray_(name string, nick string, blurb string, elementSpec *ParamSpec, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_elementSpec := elementSpec.ToC()
	sys_flags := flags
}

// ParamTypeRegisterStatic is analogous to the C function g_param_type_register_static.
func ParamTypeRegisterStatic(name string, pspecInfo *ParamSpecTypeInfo) {
	sys_name := name
	sys_pspecInfo := pspecInfo.ToC()
}

// ParamValueConvert is analogous to the C function g_param_value_convert.
func ParamValueConvert(pspec *ParamSpec, srcValue *Value, destValue *Value, strictValidation bool) {
	sys_pspec := pspec.ToC()
	sys_srcValue := srcValue.ToC()
	sys_destValue := destValue.ToC()
	sys_strictValidation := strictValidation
}

// ParamValueDefaults is analogous to the C function g_param_value_defaults.
func ParamValueDefaults(pspec *ParamSpec, value *Value) {
	sys_pspec := pspec.ToC()
	sys_value := value.ToC()
}

// ParamValueSetDefault is analogous to the C function g_param_value_set_default.
func ParamValueSetDefault(pspec *ParamSpec, value *Value) {
	sys_pspec := pspec.ToC()
	sys_value := value.ToC()
}

// ParamValueValidate is analogous to the C function g_param_value_validate.
func ParamValueValidate(pspec *ParamSpec, value *Value) {
	sys_pspec := pspec.ToC()
	sys_value := value.ToC()
}

// ParamValuesCmp is analogous to the C function g_param_values_cmp.
func ParamValuesCmp(pspec *ParamSpec, value1 *Value, value2 *Value) {
	sys_pspec := pspec.ToC()
	sys_value1 := value1.ToC()
	sys_value2 := value2.ToC()
}

// PointerTypeRegisterStatic is analogous to the C function g_pointer_type_register_static.
func PointerTypeRegisterStatic(name string) {
	sys_name := name
}

// SignalAccumulatorTrueHandled is analogous to the C function g_signal_accumulator_true_handled.
func SignalAccumulatorTrueHandled(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value, dummy unsafe.Pointer) {
	sys_ihint := ihint.ToC()
	sys_returnAccu := returnAccu.ToC()
	sys_handlerReturn := handlerReturn.ToC()
	sys_dummy := dummy
}

// UNSUPPORTED : g_signal_add_emission_hook : parameter 'hook_func' is callback

// UNSUPPORTED : g_signal_chain_from_overridden : parameter 'instance_and_params' is array parameter without length parameter

// SignalConnectClosure is analogous to the C function g_signal_connect_closure.
func SignalConnectClosure(instance unsafe.Pointer, detailedSignal string, closure *Closure, after bool) {
	sys_instance := instance
	sys_detailedSignal := detailedSignal
	sys_closure := closure.ToC()
	sys_after := after
}

// SignalConnectClosureById is analogous to the C function g_signal_connect_closure_by_id.
func SignalConnectClosureById(instance unsafe.Pointer, signalId uint, detail uint32, closure *Closure, after bool) {
	sys_instance := instance
	sys_signalId := signalId
	sys_detail := detail
	sys_closure := closure.ToC()
	sys_after := after
}

// UNSUPPORTED : g_signal_connect_data : parameter 'c_handler' is callback

// UNSUPPORTED : g_signal_connect_object : parameter 'c_handler' is callback

// SignalEmit is analogous to the C function g_signal_emit.
func SignalEmit(instance unsafe.Pointer, signalId uint, detail uint32) {
	sys_instance := instance
	sys_signalId := signalId
	sys_detail := detail
}

// SignalEmitByName is analogous to the C function g_signal_emit_by_name.
func SignalEmitByName(instance unsafe.Pointer, detailedSignal string) {
	sys_instance := instance
	sys_detailedSignal := detailedSignal
}

// SignalEmitValist is analogous to the C function g_signal_emit_valist.
func SignalEmitValist(instance unsafe.Pointer, signalId uint, detail uint32) {
	sys_instance := instance
	sys_signalId := signalId
	sys_detail := detail
}

// UNSUPPORTED : g_signal_emitv : parameter 'instance_and_params' is array parameter without length parameter

// SignalGetInvocationHint is analogous to the C function g_signal_get_invocation_hint.
func SignalGetInvocationHint(instance unsafe.Pointer) {
	sys_instance := instance
}

// SignalHandlerBlock is analogous to the C function g_signal_handler_block.
func SignalHandlerBlock(instance unsafe.Pointer, handlerId uint64) {
	sys_instance := instance
	sys_handlerId := handlerId
}

// SignalHandlerDisconnect is analogous to the C function g_signal_handler_disconnect.
func SignalHandlerDisconnect(instance unsafe.Pointer, handlerId uint64) {
	sys_instance := instance
	sys_handlerId := handlerId
}

// SignalHandlerFind is analogous to the C function g_signal_handler_find.
func SignalHandlerFind(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) {
	sys_instance := instance
	sys_mask := mask
	sys_signalId := signalId
	sys_detail := detail
	sys_closure := closure.ToC()
	sys_func_ := func_
	sys_data := data
}

// SignalHandlerIsConnected is analogous to the C function g_signal_handler_is_connected.
func SignalHandlerIsConnected(instance unsafe.Pointer, handlerId uint64) {
	sys_instance := instance
	sys_handlerId := handlerId
}

// SignalHandlerUnblock is analogous to the C function g_signal_handler_unblock.
func SignalHandlerUnblock(instance unsafe.Pointer, handlerId uint64) {
	sys_instance := instance
	sys_handlerId := handlerId
}

// SignalHandlersBlockMatched is analogous to the C function g_signal_handlers_block_matched.
func SignalHandlersBlockMatched(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) {
	sys_instance := instance
	sys_mask := mask
	sys_signalId := signalId
	sys_detail := detail
	sys_closure := closure.ToC()
	sys_func_ := func_
	sys_data := data
}

// SignalHandlersDestroy is analogous to the C function g_signal_handlers_destroy.
func SignalHandlersDestroy(instance unsafe.Pointer) {
	sys_instance := instance
}

// SignalHandlersDisconnectMatched is analogous to the C function g_signal_handlers_disconnect_matched.
func SignalHandlersDisconnectMatched(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) {
	sys_instance := instance
	sys_mask := mask
	sys_signalId := signalId
	sys_detail := detail
	sys_closure := closure.ToC()
	sys_func_ := func_
	sys_data := data
}

// SignalHandlersUnblockMatched is analogous to the C function g_signal_handlers_unblock_matched.
func SignalHandlersUnblockMatched(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) {
	sys_instance := instance
	sys_mask := mask
	sys_signalId := signalId
	sys_detail := detail
	sys_closure := closure.ToC()
	sys_func_ := func_
	sys_data := data
}

// SignalHasHandlerPending is analogous to the C function g_signal_has_handler_pending.
func SignalHasHandlerPending(instance unsafe.Pointer, signalId uint, detail uint32, mayBeBlocked bool) {
	sys_instance := instance
	sys_signalId := signalId
	sys_detail := detail
	sys_mayBeBlocked := mayBeBlocked
}

// SignalListIds is analogous to the C function g_signal_list_ids.
func SignalListIds(itype uint64) {
	sys_itype := itype
}

// SignalLookup is analogous to the C function g_signal_lookup.
func SignalLookup(name string, itype uint64) {
	sys_name := name
	sys_itype := itype
}

// SignalName is analogous to the C function g_signal_name.
func SignalName(signalId uint) {
	sys_signalId := signalId
}

// UNSUPPORTED : g_signal_new : parameter 'accumulator' is callback

// UNSUPPORTED : g_signal_new_class_handler : parameter 'class_handler' is callback

// UNSUPPORTED : g_signal_new_valist : parameter 'accumulator' is callback

// UNSUPPORTED : g_signal_newv : parameter 'accumulator' is callback

// SignalOverrideClassClosure is analogous to the C function g_signal_override_class_closure.
func SignalOverrideClassClosure(signalId uint, instanceType uint64, classClosure *Closure) {
	sys_signalId := signalId
	sys_instanceType := instanceType
	sys_classClosure := classClosure.ToC()
}

// UNSUPPORTED : g_signal_override_class_handler : parameter 'class_handler' is callback

// SignalParseName is analogous to the C function g_signal_parse_name.
func SignalParseName(detailedSignal string, itype uint64, forceDetailQuark bool) {
	sys_detailedSignal := detailedSignal
	sys_itype := itype
	sys_forceDetailQuark := forceDetailQuark
}

// SignalQuery_ is analogous to the C function g_signal_query.
func SignalQuery_(signalId uint) {
	sys_signalId := signalId
}

// SignalRemoveEmissionHook is analogous to the C function g_signal_remove_emission_hook.
func SignalRemoveEmissionHook(signalId uint, hookId uint64) {
	sys_signalId := signalId
	sys_hookId := hookId
}

// UNSUPPORTED : g_signal_set_va_marshaller : blacklisted

// SignalStopEmission is analogous to the C function g_signal_stop_emission.
func SignalStopEmission(instance unsafe.Pointer, signalId uint, detail uint32) {
	sys_instance := instance
	sys_signalId := signalId
	sys_detail := detail
}

// SignalStopEmissionByName is analogous to the C function g_signal_stop_emission_by_name.
func SignalStopEmissionByName(instance unsafe.Pointer, detailedSignal string) {
	sys_instance := instance
	sys_detailedSignal := detailedSignal
}

// SignalTypeCclosureNew is analogous to the C function g_signal_type_cclosure_new.
func SignalTypeCclosureNew(itype uint64, structOffset uint) {
	sys_itype := itype
	sys_structOffset := structOffset
}

// SourceSetClosure is analogous to the C function g_source_set_closure.
func SourceSetClosure(source *glib.Source, closure *Closure) {
	sys_source := source.ToC()
	sys_closure := closure.ToC()
}

// SourceSetDummyCallback is analogous to the C function g_source_set_dummy_callback.
func SourceSetDummyCallback(source *glib.Source) {
	sys_source := source.ToC()
}

// StrdupValueContents is analogous to the C function g_strdup_value_contents.
func StrdupValueContents(value *Value) {
	sys_value := value.ToC()
}

// UNSUPPORTED : g_type_add_class_cache_func : parameter 'cache_func' is callback

// TypeAddInstancePrivate is analogous to the C function g_type_add_instance_private.
func TypeAddInstancePrivate(classType uint64, privateSize uint64) {
	sys_classType := classType
	sys_privateSize := privateSize
}

// UNSUPPORTED : g_type_add_interface_check : parameter 'check_func' is callback

// TypeAddInterfaceDynamic is analogous to the C function g_type_add_interface_dynamic.
func TypeAddInterfaceDynamic(instanceType uint64, interfaceType uint64, plugin *TypePlugin) {
	sys_instanceType := instanceType
	sys_interfaceType := interfaceType
	sys_plugin := plugin.ToC()
}

// TypeAddInterfaceStatic is analogous to the C function g_type_add_interface_static.
func TypeAddInterfaceStatic(instanceType uint64, interfaceType uint64, info *InterfaceInfo) {
	sys_instanceType := instanceType
	sys_interfaceType := interfaceType
	sys_info := info.ToC()
}

// TypeCheckClassCast is analogous to the C function g_type_check_class_cast.
func TypeCheckClassCast(gClass *TypeClass, isAType uint64) {
	sys_gClass := gClass.ToC()
	sys_isAType := isAType
}

// TypeCheckClassIsA is analogous to the C function g_type_check_class_is_a.
func TypeCheckClassIsA(gClass *TypeClass, isAType uint64) {
	sys_gClass := gClass.ToC()
	sys_isAType := isAType
}

// TypeCheckInstance is analogous to the C function g_type_check_instance.
func TypeCheckInstance(instance *TypeInstance) {
	sys_instance := instance.ToC()
}

// TypeCheckInstanceCast is analogous to the C function g_type_check_instance_cast.
func TypeCheckInstanceCast(instance *TypeInstance, ifaceType uint64) {
	sys_instance := instance.ToC()
	sys_ifaceType := ifaceType
}

// TypeCheckInstanceIsA is analogous to the C function g_type_check_instance_is_a.
func TypeCheckInstanceIsA(instance *TypeInstance, ifaceType uint64) {
	sys_instance := instance.ToC()
	sys_ifaceType := ifaceType
}

// TypeCheckInstanceIsFundamentallyA is analogous to the C function g_type_check_instance_is_fundamentally_a.
func TypeCheckInstanceIsFundamentallyA(instance *TypeInstance, fundamentalType uint64) {
	sys_instance := instance.ToC()
	sys_fundamentalType := fundamentalType
}

// TypeCheckIsValueType is analogous to the C function g_type_check_is_value_type.
func TypeCheckIsValueType(type_ uint64) {
	sys_type_ := type_
}

// TypeCheckValue is analogous to the C function g_type_check_value.
func TypeCheckValue(value *Value) {
	sys_value := value.ToC()
}

// TypeCheckValueHolds is analogous to the C function g_type_check_value_holds.
func TypeCheckValueHolds(value *Value, type_ uint64) {
	sys_value := value.ToC()
	sys_type_ := type_
}

// TypeChildren is analogous to the C function g_type_children.
func TypeChildren(type_ uint64) {
	sys_type_ := type_
}

// TypeCreateInstance is analogous to the C function g_type_create_instance.
func TypeCreateInstance(type_ uint64) {
	sys_type_ := type_
}

// TypeDefaultInterfacePeek is analogous to the C function g_type_default_interface_peek.
func TypeDefaultInterfacePeek(gType uint64) {
	sys_gType := gType
}

// TypeDefaultInterfaceRef is analogous to the C function g_type_default_interface_ref.
func TypeDefaultInterfaceRef(gType uint64) {
	sys_gType := gType
}

// TypeDefaultInterfaceUnref is analogous to the C function g_type_default_interface_unref.
func TypeDefaultInterfaceUnref(gIface unsafe.Pointer) {
	sys_gIface := gIface
}

// TypeDepth is analogous to the C function g_type_depth.
func TypeDepth(type_ uint64) {
	sys_type_ := type_
}

// TypeFreeInstance is analogous to the C function g_type_free_instance.
func TypeFreeInstance(instance *TypeInstance) {
	sys_instance := instance.ToC()
}

// TypeFromName is analogous to the C function g_type_from_name.
func TypeFromName(name string) {
	sys_name := name
}

// TypeFundamental is analogous to the C function g_type_fundamental.
func TypeFundamental(typeId uint64) {
	sys_typeId := typeId
}

// TypeFundamentalNext is analogous to the C function g_type_fundamental_next.
func TypeFundamentalNext() {}

// TypeGetPlugin is analogous to the C function g_type_get_plugin.
func TypeGetPlugin(type_ uint64) {
	sys_type_ := type_
}

// TypeGetQdata is analogous to the C function g_type_get_qdata.
func TypeGetQdata(type_ uint64, quark uint32) {
	sys_type_ := type_
	sys_quark := quark
}

// TypeInit is analogous to the C function g_type_init.
func TypeInit() {}

// TypeInitWithDebugFlags is analogous to the C function g_type_init_with_debug_flags.
func TypeInitWithDebugFlags(debugFlags int) {
	sys_debugFlags := debugFlags
}

// TypeInterfaces is analogous to the C function g_type_interfaces.
func TypeInterfaces(type_ uint64) {
	sys_type_ := type_
}

// TypeIsA is analogous to the C function g_type_is_a.
func TypeIsA(type_ uint64, isAType uint64) {
	sys_type_ := type_
	sys_isAType := isAType
}

// TypeName is analogous to the C function g_type_name.
func TypeName(type_ uint64) {
	sys_type_ := type_
}

// TypeNameFromClass is analogous to the C function g_type_name_from_class.
func TypeNameFromClass(gClass *TypeClass) {
	sys_gClass := gClass.ToC()
}

// TypeNameFromInstance is analogous to the C function g_type_name_from_instance.
func TypeNameFromInstance(instance *TypeInstance) {
	sys_instance := instance.ToC()
}

// TypeNextBase is analogous to the C function g_type_next_base.
func TypeNextBase(leafType uint64, rootType uint64) {
	sys_leafType := leafType
	sys_rootType := rootType
}

// TypeParent is analogous to the C function g_type_parent.
func TypeParent(type_ uint64) {
	sys_type_ := type_
}

// TypeQname is analogous to the C function g_type_qname.
func TypeQname(type_ uint64) {
	sys_type_ := type_
}

// TypeQuery_ is analogous to the C function g_type_query.
func TypeQuery_(type_ uint64) {
	sys_type_ := type_
}

// TypeRegisterDynamic is analogous to the C function g_type_register_dynamic.
func TypeRegisterDynamic(parentType uint64, typeName string, plugin *TypePlugin, flags int) {
	sys_parentType := parentType
	sys_typeName := typeName
	sys_plugin := plugin.ToC()
	sys_flags := flags
}

// TypeRegisterFundamental is analogous to the C function g_type_register_fundamental.
func TypeRegisterFundamental(typeId uint64, typeName string, info *TypeInfo, finfo *TypeFundamentalInfo, flags int) {
	sys_typeId := typeId
	sys_typeName := typeName
	sys_info := info.ToC()
	sys_finfo := finfo.ToC()
	sys_flags := flags
}

// TypeRegisterStatic is analogous to the C function g_type_register_static.
func TypeRegisterStatic(parentType uint64, typeName string, info *TypeInfo, flags int) {
	sys_parentType := parentType
	sys_typeName := typeName
	sys_info := info.ToC()
	sys_flags := flags
}

// UNSUPPORTED : g_type_register_static_simple : parameter 'class_init' is callback

// UNSUPPORTED : g_type_remove_class_cache_func : parameter 'cache_func' is callback

// UNSUPPORTED : g_type_remove_interface_check : parameter 'check_func' is callback

// TypeSetQdata is analogous to the C function g_type_set_qdata.
func TypeSetQdata(type_ uint64, quark uint32, data unsafe.Pointer) {
	sys_type_ := type_
	sys_quark := quark
	sys_data := data
}

// TypeTestFlags is analogous to the C function g_type_test_flags.
func TypeTestFlags(type_ uint64, flags uint) {
	sys_type_ := type_
	sys_flags := flags
}

// UNSUPPORTED : g_value_register_transform_func : parameter 'transform_func' is callback

// CClosure is a representation of the C record GCClosure.
type CClosure struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GCClosure that represents the CClosure.
func (recv *CClosure) ToC() unsafe.Pointer {
	return recv.native
}

// Closure is a representation of the C record GClosure.
type Closure struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GClosure that represents the Closure.
func (recv *Closure) ToC() unsafe.Pointer {
	return recv.native
}

// ClosureNotifyData is a representation of the C record GClosureNotifyData.
type ClosureNotifyData struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GClosureNotifyData that represents the ClosureNotifyData.
func (recv *ClosureNotifyData) ToC() unsafe.Pointer {
	return recv.native
}

// EnumClass is a representation of the C record GEnumClass.
type EnumClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GEnumClass that represents the EnumClass.
func (recv *EnumClass) ToC() unsafe.Pointer {
	return recv.native
}

// EnumValue is a representation of the C record GEnumValue.
type EnumValue struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GEnumValue that represents the EnumValue.
func (recv *EnumValue) ToC() unsafe.Pointer {
	return recv.native
}

// FlagsClass is a representation of the C record GFlagsClass.
type FlagsClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFlagsClass that represents the FlagsClass.
func (recv *FlagsClass) ToC() unsafe.Pointer {
	return recv.native
}

// FlagsValue is a representation of the C record GFlagsValue.
type FlagsValue struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFlagsValue that represents the FlagsValue.
func (recv *FlagsValue) ToC() unsafe.Pointer {
	return recv.native
}

// InitiallyUnownedClass is a representation of the C record GInitiallyUnownedClass.
type InitiallyUnownedClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInitiallyUnownedClass that represents the InitiallyUnownedClass.
func (recv *InitiallyUnownedClass) ToC() unsafe.Pointer {
	return recv.native
}

// InterfaceInfo is a representation of the C record GInterfaceInfo.
type InterfaceInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInterfaceInfo that represents the InterfaceInfo.
func (recv *InterfaceInfo) ToC() unsafe.Pointer {
	return recv.native
}

// ObjectClass is a representation of the C record GObjectClass.
type ObjectClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GObjectClass that represents the ObjectClass.
func (recv *ObjectClass) ToC() unsafe.Pointer {
	return recv.native
}

// ObjectConstructParam is a representation of the C record GObjectConstructParam.
type ObjectConstructParam struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GObjectConstructParam that represents the ObjectConstructParam.
func (recv *ObjectConstructParam) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecClass is a representation of the C record GParamSpecClass.
type ParamSpecClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecClass that represents the ParamSpecClass.
func (recv *ParamSpecClass) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecPool is a representation of the C record GParamSpecPool.
type ParamSpecPool struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecPool that represents the ParamSpecPool.
func (recv *ParamSpecPool) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecTypeInfo is a representation of the C record GParamSpecTypeInfo.
type ParamSpecTypeInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecTypeInfo that represents the ParamSpecTypeInfo.
func (recv *ParamSpecTypeInfo) ToC() unsafe.Pointer {
	return recv.native
}

// Parameter is a representation of the C record GParameter.
type Parameter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParameter that represents the Parameter.
func (recv *Parameter) ToC() unsafe.Pointer {
	return recv.native
}

// SignalInvocationHint is a representation of the C record GSignalInvocationHint.
type SignalInvocationHint struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSignalInvocationHint that represents the SignalInvocationHint.
func (recv *SignalInvocationHint) ToC() unsafe.Pointer {
	return recv.native
}

// SignalQuery is a representation of the C record GSignalQuery.
type SignalQuery struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSignalQuery that represents the SignalQuery.
func (recv *SignalQuery) ToC() unsafe.Pointer {
	return recv.native
}

// TypeClass is a representation of the C record GTypeClass.
type TypeClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeClass that represents the TypeClass.
func (recv *TypeClass) ToC() unsafe.Pointer {
	return recv.native
}

// TypeFundamentalInfo is a representation of the C record GTypeFundamentalInfo.
type TypeFundamentalInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeFundamentalInfo that represents the TypeFundamentalInfo.
func (recv *TypeFundamentalInfo) ToC() unsafe.Pointer {
	return recv.native
}

// TypeInfo is a representation of the C record GTypeInfo.
type TypeInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeInfo that represents the TypeInfo.
func (recv *TypeInfo) ToC() unsafe.Pointer {
	return recv.native
}

// TypeInstance is a representation of the C record GTypeInstance.
type TypeInstance struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeInstance that represents the TypeInstance.
func (recv *TypeInstance) ToC() unsafe.Pointer {
	return recv.native
}

// TypeInterface is a representation of the C record GTypeInterface.
type TypeInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeInterface that represents the TypeInterface.
func (recv *TypeInterface) ToC() unsafe.Pointer {
	return recv.native
}

// TypeModuleClass is a representation of the C record GTypeModuleClass.
type TypeModuleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeModuleClass that represents the TypeModuleClass.
func (recv *TypeModuleClass) ToC() unsafe.Pointer {
	return recv.native
}

// TypePluginClass is a representation of the C record GTypePluginClass.
type TypePluginClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypePluginClass that represents the TypePluginClass.
func (recv *TypePluginClass) ToC() unsafe.Pointer {
	return recv.native
}

// TypeQuery is a representation of the C record GTypeQuery.
type TypeQuery struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeQuery that represents the TypeQuery.
func (recv *TypeQuery) ToC() unsafe.Pointer {
	return recv.native
}

// TypeValueTable is a representation of the C record GTypeValueTable.
type TypeValueTable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeValueTable that represents the TypeValueTable.
func (recv *TypeValueTable) ToC() unsafe.Pointer {
	return recv.native
}

// Value is a representation of the C record GValue.
type Value struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GValue that represents the Value.
func (recv *Value) ToC() unsafe.Pointer {
	return recv.native
}

// ValueArray is a representation of the C record GValueArray.
type ValueArray struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GValueArray that represents the ValueArray.
func (recv *ValueArray) ToC() unsafe.Pointer {
	return recv.native
}

// WeakRef is a representation of the C record GWeakRef.
type WeakRef struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GWeakRef that represents the WeakRef.
func (recv *WeakRef) ToC() unsafe.Pointer {
	return recv.native
}

// InitiallyUnowned is a representation of the C record GInitiallyUnowned.
type InitiallyUnowned struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInitiallyUnowned that represents the InitiallyUnowned.
func (recv *InitiallyUnowned) ToC() unsafe.Pointer {
	return recv.native
}

// Object is a representation of the C record GObject.
type Object struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GObject that represents the Object.
func (recv *Object) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpec is a representation of the C record GParamSpec.
type ParamSpec struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpec that represents the ParamSpec.
func (recv *ParamSpec) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecBoolean is a representation of the C record GParamSpecBoolean.
type ParamSpecBoolean struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecBoolean that represents the ParamSpecBoolean.
func (recv *ParamSpecBoolean) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecBoxed is a representation of the C record GParamSpecBoxed.
type ParamSpecBoxed struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecBoxed that represents the ParamSpecBoxed.
func (recv *ParamSpecBoxed) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecChar is a representation of the C record GParamSpecChar.
type ParamSpecChar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecChar that represents the ParamSpecChar.
func (recv *ParamSpecChar) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecDouble is a representation of the C record GParamSpecDouble.
type ParamSpecDouble struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecDouble that represents the ParamSpecDouble.
func (recv *ParamSpecDouble) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecEnum is a representation of the C record GParamSpecEnum.
type ParamSpecEnum struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecEnum that represents the ParamSpecEnum.
func (recv *ParamSpecEnum) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecFlags is a representation of the C record GParamSpecFlags.
type ParamSpecFlags struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecFlags that represents the ParamSpecFlags.
func (recv *ParamSpecFlags) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecFloat is a representation of the C record GParamSpecFloat.
type ParamSpecFloat struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecFloat that represents the ParamSpecFloat.
func (recv *ParamSpecFloat) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecInt is a representation of the C record GParamSpecInt.
type ParamSpecInt struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecInt that represents the ParamSpecInt.
func (recv *ParamSpecInt) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecInt64 is a representation of the C record GParamSpecInt64.
type ParamSpecInt64 struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecInt64 that represents the ParamSpecInt64.
func (recv *ParamSpecInt64) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecLong is a representation of the C record GParamSpecLong.
type ParamSpecLong struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecLong that represents the ParamSpecLong.
func (recv *ParamSpecLong) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecObject is a representation of the C record GParamSpecObject.
type ParamSpecObject struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecObject that represents the ParamSpecObject.
func (recv *ParamSpecObject) ToC() unsafe.Pointer {
	return recv.native
}

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

// ParamSpecParam is a representation of the C record GParamSpecParam.
type ParamSpecParam struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecParam that represents the ParamSpecParam.
func (recv *ParamSpecParam) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecPointer is a representation of the C record GParamSpecPointer.
type ParamSpecPointer struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecPointer that represents the ParamSpecPointer.
func (recv *ParamSpecPointer) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecString is a representation of the C record GParamSpecString.
type ParamSpecString struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecString that represents the ParamSpecString.
func (recv *ParamSpecString) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecUChar is a representation of the C record GParamSpecUChar.
type ParamSpecUChar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecUChar that represents the ParamSpecUChar.
func (recv *ParamSpecUChar) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecUInt is a representation of the C record GParamSpecUInt.
type ParamSpecUInt struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecUInt that represents the ParamSpecUInt.
func (recv *ParamSpecUInt) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecUInt64 is a representation of the C record GParamSpecUInt64.
type ParamSpecUInt64 struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecUInt64 that represents the ParamSpecUInt64.
func (recv *ParamSpecUInt64) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecULong is a representation of the C record GParamSpecULong.
type ParamSpecULong struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecULong that represents the ParamSpecULong.
func (recv *ParamSpecULong) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecUnichar is a representation of the C record GParamSpecUnichar.
type ParamSpecUnichar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecUnichar that represents the ParamSpecUnichar.
func (recv *ParamSpecUnichar) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecValueArray is a representation of the C record GParamSpecValueArray.
type ParamSpecValueArray struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecValueArray that represents the ParamSpecValueArray.
func (recv *ParamSpecValueArray) ToC() unsafe.Pointer {
	return recv.native
}

// TypeModule is a representation of the C record GTypeModule.
type TypeModule struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeModule that represents the TypeModule.
func (recv *TypeModule) ToC() unsafe.Pointer {
	return recv.native
}

// TypePlugin is a representation of the C interface GTypePlugin.
type TypePlugin struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypePlugin that represents the TypePlugin.
func (recv *TypePlugin) ToC() unsafe.Pointer {
	return recv.native
}

// TypeCValue is a representation of the C union GTypeCValue.
type TypeCValue struct{}

// _Value__data__union is a representation of the C union .
type _Value__data__union struct{}
