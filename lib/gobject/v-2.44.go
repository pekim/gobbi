// Code generated - DO NOT EDIT.
// +build gobject_2.44

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

func BoxedCopy(boxedType uint64, srcBoxed unsafe.Pointer) {
	sys_boxedType := uint64(boxedType)
	sys_srcBoxed := unsafe.Pointer(srcBoxed)
}

func BoxedFree(boxedType uint64, boxed unsafe.Pointer) {
	sys_boxedType := uint64(boxedType)
	sys_boxed := unsafe.Pointer(boxed)
}

// UNSUPPORTED : g_boxed_type_register_static : parameter 'boxed_copy' is callback

// UNSUPPORTED : g_cclosure_new : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object_swap : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_swap : parameter 'callback_func' is callback

func ClearObject(objectPtr **Object) {
	sys_objectPtr := *unsafe.Pointer(objectPtr)
}

func EnumCompleteTypeInfo(gEnumType uint64, constValues *EnumValue) {
	sys_gEnumType := uint64(gEnumType)
	sys_constValues := unsafe.Pointer(constValues)
}

func EnumGetValue(enumClass *EnumClass, value int) {
	sys_enumClass := unsafe.Pointer(enumClass)
	sys_value := int(value)
}

func EnumGetValueByName(enumClass *EnumClass, name string) {
	sys_enumClass := unsafe.Pointer(enumClass)
	sys_name := string(name)
}

func EnumGetValueByNick(enumClass *EnumClass, nick string) {
	sys_enumClass := unsafe.Pointer(enumClass)
	sys_nick := string(nick)
}

func EnumRegisterStatic(name string, constStaticValues *EnumValue) {
	sys_name := string(name)
	sys_constStaticValues := unsafe.Pointer(constStaticValues)
}

func FlagsCompleteTypeInfo(gFlagsType uint64, constValues *FlagsValue) {
	sys_gFlagsType := uint64(gFlagsType)
	sys_constValues := unsafe.Pointer(constValues)
}

func FlagsGetFirstValue(flagsClass *FlagsClass, value uint) {
	sys_flagsClass := unsafe.Pointer(flagsClass)
	sys_value := uint(value)
}

func FlagsGetValueByName(flagsClass *FlagsClass, name string) {
	sys_flagsClass := unsafe.Pointer(flagsClass)
	sys_name := string(name)
}

func FlagsGetValueByNick(flagsClass *FlagsClass, nick string) {
	sys_flagsClass := unsafe.Pointer(flagsClass)
	sys_nick := string(nick)
}

func FlagsRegisterStatic(name string, constStaticValues *FlagsValue) {
	sys_name := string(name)
	sys_constStaticValues := unsafe.Pointer(constStaticValues)
}

func GtypeGetType() {}

func ParamSpecBoolean(name string, nick string, blurb string, defaultValue bool, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_defaultValue := bool(defaultValue)
	sys_flags := int(flags)
}

func ParamSpecBoxed(name string, nick string, blurb string, boxedType uint64, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_boxedType := uint64(boxedType)
	sys_flags := int(flags)
}

func ParamSpecChar(name string, nick string, blurb string, minimum int8, maximum int8, defaultValue int8, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_minimum := int8(minimum)
	sys_maximum := int8(maximum)
	sys_defaultValue := int8(defaultValue)
	sys_flags := int(flags)
}

func ParamSpecDouble(name string, nick string, blurb string, minimum float64, maximum float64, defaultValue float64, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_minimum := float64(minimum)
	sys_maximum := float64(maximum)
	sys_defaultValue := float64(defaultValue)
	sys_flags := int(flags)
}

func ParamSpecEnum(name string, nick string, blurb string, enumType uint64, defaultValue int, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_enumType := uint64(enumType)
	sys_defaultValue := int(defaultValue)
	sys_flags := int(flags)
}

func ParamSpecFlags(name string, nick string, blurb string, flagsType uint64, defaultValue uint, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_flagsType := uint64(flagsType)
	sys_defaultValue := uint(defaultValue)
	sys_flags := int(flags)
}

func ParamSpecFloat(name string, nick string, blurb string, minimum float32, maximum float32, defaultValue float32, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_minimum := float32(minimum)
	sys_maximum := float32(maximum)
	sys_defaultValue := float32(defaultValue)
	sys_flags := int(flags)
}

func ParamSpecGtype(name string, nick string, blurb string, isAType uint64, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_isAType := uint64(isAType)
	sys_flags := int(flags)
}

func ParamSpecInt(name string, nick string, blurb string, minimum int, maximum int, defaultValue int, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_minimum := int(minimum)
	sys_maximum := int(maximum)
	sys_defaultValue := int(defaultValue)
	sys_flags := int(flags)
}

func ParamSpecInt64(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_minimum := int64(minimum)
	sys_maximum := int64(maximum)
	sys_defaultValue := int64(defaultValue)
	sys_flags := int(flags)
}

func ParamSpecLong(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_minimum := int64(minimum)
	sys_maximum := int64(maximum)
	sys_defaultValue := int64(defaultValue)
	sys_flags := int(flags)
}

func ParamSpecObject(name string, nick string, blurb string, objectType uint64, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_objectType := uint64(objectType)
	sys_flags := int(flags)
}

func ParamSpecOverride(name string, overridden *ParamSpec) {
	sys_name := string(name)
	sys_overridden := unsafe.Pointer(overridden)
}

func ParamSpecParam(name string, nick string, blurb string, paramType uint64, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_paramType := uint64(paramType)
	sys_flags := int(flags)
}

func ParamSpecPointer(name string, nick string, blurb string, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_flags := int(flags)
}

func ParamSpecString(name string, nick string, blurb string, defaultValue string, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_defaultValue := string(defaultValue)
	sys_flags := int(flags)
}

func ParamSpecUchar(name string, nick string, blurb string, minimum uint8, maximum uint8, defaultValue uint8, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_minimum := uint8(minimum)
	sys_maximum := uint8(maximum)
	sys_defaultValue := uint8(defaultValue)
	sys_flags := int(flags)
}

func ParamSpecUint(name string, nick string, blurb string, minimum uint, maximum uint, defaultValue uint, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_minimum := uint(minimum)
	sys_maximum := uint(maximum)
	sys_defaultValue := uint(defaultValue)
	sys_flags := int(flags)
}

func ParamSpecUint64(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_minimum := uint64(minimum)
	sys_maximum := uint64(maximum)
	sys_defaultValue := uint64(defaultValue)
	sys_flags := int(flags)
}

func ParamSpecUlong(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_minimum := uint64(minimum)
	sys_maximum := uint64(maximum)
	sys_defaultValue := uint64(defaultValue)
	sys_flags := int(flags)
}

func ParamSpecUnichar(name string, nick string, blurb string, defaultValue rune, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_defaultValue := rune(defaultValue)
	sys_flags := int(flags)
}

func ParamSpecValueArray(name string, nick string, blurb string, elementSpec *ParamSpec, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_elementSpec := unsafe.Pointer(elementSpec)
	sys_flags := int(flags)
}

func ParamSpecVariant(name string, nick string, blurb string, type_ *glib.VariantType, defaultValue *glib.Variant, flags int) {
	sys_name := string(name)
	sys_nick := string(nick)
	sys_blurb := string(blurb)
	sys_type_ := unsafe.Pointer(type_)
	sys_defaultValue := unsafe.Pointer(defaultValue)
	sys_flags := int(flags)
}

func ParamTypeRegisterStatic(name string, pspecInfo *ParamSpecTypeInfo) {
	sys_name := string(name)
	sys_pspecInfo := unsafe.Pointer(pspecInfo)
}

func ParamValueConvert(pspec *ParamSpec, srcValue *Value, destValue *Value, strictValidation bool) {
	sys_pspec := unsafe.Pointer(pspec)
	sys_srcValue := unsafe.Pointer(srcValue)
	sys_destValue := unsafe.Pointer(destValue)
	sys_strictValidation := bool(strictValidation)
}

func ParamValueDefaults(pspec *ParamSpec, value *Value) {
	sys_pspec := unsafe.Pointer(pspec)
	sys_value := unsafe.Pointer(value)
}

func ParamValueSetDefault(pspec *ParamSpec, value *Value) {
	sys_pspec := unsafe.Pointer(pspec)
	sys_value := unsafe.Pointer(value)
}

func ParamValueValidate(pspec *ParamSpec, value *Value) {
	sys_pspec := unsafe.Pointer(pspec)
	sys_value := unsafe.Pointer(value)
}

func ParamValuesCmp(pspec *ParamSpec, value1 *Value, value2 *Value) {
	sys_pspec := unsafe.Pointer(pspec)
	sys_value1 := unsafe.Pointer(value1)
	sys_value2 := unsafe.Pointer(value2)
}

func PointerTypeRegisterStatic(name string) {
	sys_name := string(name)
}

func SignalAccumulatorFirstWins(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value, dummy unsafe.Pointer) {
	sys_ihint := unsafe.Pointer(ihint)
	sys_returnAccu := unsafe.Pointer(returnAccu)
	sys_handlerReturn := unsafe.Pointer(handlerReturn)
	sys_dummy := unsafe.Pointer(dummy)
}

func SignalAccumulatorTrueHandled(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value, dummy unsafe.Pointer) {
	sys_ihint := unsafe.Pointer(ihint)
	sys_returnAccu := unsafe.Pointer(returnAccu)
	sys_handlerReturn := unsafe.Pointer(handlerReturn)
	sys_dummy := unsafe.Pointer(dummy)
}

// UNSUPPORTED : g_signal_add_emission_hook : parameter 'hook_func' is callback

// UNSUPPORTED : g_signal_chain_from_overridden : parameter 'instance_and_params' is array parameter without length parameter

func SignalChainFromOverriddenHandler(instance unsafe.Pointer) {
	sys_instance := unsafe.Pointer(instance)
}

func SignalConnectClosure(instance unsafe.Pointer, detailedSignal string, closure *Closure, after bool) {
	sys_instance := unsafe.Pointer(instance)
	sys_detailedSignal := string(detailedSignal)
	sys_closure := unsafe.Pointer(closure)
	sys_after := bool(after)
}

func SignalConnectClosureById(instance unsafe.Pointer, signalId uint, detail uint32, closure *Closure, after bool) {
	sys_instance := unsafe.Pointer(instance)
	sys_signalId := uint(signalId)
	sys_detail := uint32(detail)
	sys_closure := unsafe.Pointer(closure)
	sys_after := bool(after)
}

// UNSUPPORTED : g_signal_connect_data : parameter 'c_handler' is callback

// UNSUPPORTED : g_signal_connect_object : parameter 'c_handler' is callback

func SignalEmit(instance unsafe.Pointer, signalId uint, detail uint32) {
	sys_instance := unsafe.Pointer(instance)
	sys_signalId := uint(signalId)
	sys_detail := uint32(detail)
}

func SignalEmitByName(instance unsafe.Pointer, detailedSignal string) {
	sys_instance := unsafe.Pointer(instance)
	sys_detailedSignal := string(detailedSignal)
}

func SignalEmitValist(instance unsafe.Pointer, signalId uint, detail uint32) {
	sys_instance := unsafe.Pointer(instance)
	sys_signalId := uint(signalId)
	sys_detail := uint32(detail)
}

// UNSUPPORTED : g_signal_emitv : parameter 'instance_and_params' is array parameter without length parameter

func SignalGetInvocationHint(instance unsafe.Pointer) {
	sys_instance := unsafe.Pointer(instance)
}

func SignalHandlerBlock(instance unsafe.Pointer, handlerId uint64) {
	sys_instance := unsafe.Pointer(instance)
	sys_handlerId := uint64(handlerId)
}

func SignalHandlerDisconnect(instance unsafe.Pointer, handlerId uint64) {
	sys_instance := unsafe.Pointer(instance)
	sys_handlerId := uint64(handlerId)
}

func SignalHandlerFind(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) {
	sys_instance := unsafe.Pointer(instance)
	sys_mask := int(mask)
	sys_signalId := uint(signalId)
	sys_detail := uint32(detail)
	sys_closure := unsafe.Pointer(closure)
	sys_func_ := unsafe.Pointer(func_)
	sys_data := unsafe.Pointer(data)
}

func SignalHandlerIsConnected(instance unsafe.Pointer, handlerId uint64) {
	sys_instance := unsafe.Pointer(instance)
	sys_handlerId := uint64(handlerId)
}

func SignalHandlerUnblock(instance unsafe.Pointer, handlerId uint64) {
	sys_instance := unsafe.Pointer(instance)
	sys_handlerId := uint64(handlerId)
}

func SignalHandlersBlockMatched(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) {
	sys_instance := unsafe.Pointer(instance)
	sys_mask := int(mask)
	sys_signalId := uint(signalId)
	sys_detail := uint32(detail)
	sys_closure := unsafe.Pointer(closure)
	sys_func_ := unsafe.Pointer(func_)
	sys_data := unsafe.Pointer(data)
}

func SignalHandlersDestroy(instance unsafe.Pointer) {
	sys_instance := unsafe.Pointer(instance)
}

func SignalHandlersDisconnectMatched(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) {
	sys_instance := unsafe.Pointer(instance)
	sys_mask := int(mask)
	sys_signalId := uint(signalId)
	sys_detail := uint32(detail)
	sys_closure := unsafe.Pointer(closure)
	sys_func_ := unsafe.Pointer(func_)
	sys_data := unsafe.Pointer(data)
}

func SignalHandlersUnblockMatched(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) {
	sys_instance := unsafe.Pointer(instance)
	sys_mask := int(mask)
	sys_signalId := uint(signalId)
	sys_detail := uint32(detail)
	sys_closure := unsafe.Pointer(closure)
	sys_func_ := unsafe.Pointer(func_)
	sys_data := unsafe.Pointer(data)
}

func SignalHasHandlerPending(instance unsafe.Pointer, signalId uint, detail uint32, mayBeBlocked bool) {
	sys_instance := unsafe.Pointer(instance)
	sys_signalId := uint(signalId)
	sys_detail := uint32(detail)
	sys_mayBeBlocked := bool(mayBeBlocked)
}

func SignalListIds(itype uint64) {
	sys_itype := uint64(itype)
}

func SignalLookup(name string, itype uint64) {
	sys_name := string(name)
	sys_itype := uint64(itype)
}

func SignalName(signalId uint) {
	sys_signalId := uint(signalId)
}

// UNSUPPORTED : g_signal_new : parameter 'accumulator' is callback

// UNSUPPORTED : g_signal_new_class_handler : parameter 'class_handler' is callback

// UNSUPPORTED : g_signal_new_valist : parameter 'accumulator' is callback

// UNSUPPORTED : g_signal_newv : parameter 'accumulator' is callback

func SignalOverrideClassClosure(signalId uint, instanceType uint64, classClosure *Closure) {
	sys_signalId := uint(signalId)
	sys_instanceType := uint64(instanceType)
	sys_classClosure := unsafe.Pointer(classClosure)
}

// UNSUPPORTED : g_signal_override_class_handler : parameter 'class_handler' is callback

func SignalParseName(detailedSignal string, itype uint64, forceDetailQuark bool) {
	sys_detailedSignal := string(detailedSignal)
	sys_itype := uint64(itype)
	sys_forceDetailQuark := bool(forceDetailQuark)
}

func SignalQuery(signalId uint) {
	sys_signalId := uint(signalId)
}

func SignalRemoveEmissionHook(signalId uint, hookId uint64) {
	sys_signalId := uint(signalId)
	sys_hookId := uint64(hookId)
}

// UNSUPPORTED : g_signal_set_va_marshaller : blacklisted

func SignalStopEmission(instance unsafe.Pointer, signalId uint, detail uint32) {
	sys_instance := unsafe.Pointer(instance)
	sys_signalId := uint(signalId)
	sys_detail := uint32(detail)
}

func SignalStopEmissionByName(instance unsafe.Pointer, detailedSignal string) {
	sys_instance := unsafe.Pointer(instance)
	sys_detailedSignal := string(detailedSignal)
}

func SignalTypeCclosureNew(itype uint64, structOffset uint) {
	sys_itype := uint64(itype)
	sys_structOffset := uint(structOffset)
}

func SourceSetClosure(source *glib.Source, closure *Closure) {
	sys_source := unsafe.Pointer(source)
	sys_closure := unsafe.Pointer(closure)
}

func SourceSetDummyCallback(source *glib.Source) {
	sys_source := unsafe.Pointer(source)
}

func StrdupValueContents(value *Value) {
	sys_value := unsafe.Pointer(value)
}

// UNSUPPORTED : g_type_add_class_cache_func : parameter 'cache_func' is callback

func TypeAddClassPrivate(classType uint64, privateSize uint64) {
	sys_classType := uint64(classType)
	sys_privateSize := uint64(privateSize)
}

func TypeAddInstancePrivate(classType uint64, privateSize uint64) {
	sys_classType := uint64(classType)
	sys_privateSize := uint64(privateSize)
}

// UNSUPPORTED : g_type_add_interface_check : parameter 'check_func' is callback

func TypeAddInterfaceDynamic(instanceType uint64, interfaceType uint64, plugin *TypePlugin) {
	sys_instanceType := uint64(instanceType)
	sys_interfaceType := uint64(interfaceType)
	sys_plugin := unsafe.Pointer(plugin)
}

func TypeAddInterfaceStatic(instanceType uint64, interfaceType uint64, info *InterfaceInfo) {
	sys_instanceType := uint64(instanceType)
	sys_interfaceType := uint64(interfaceType)
	sys_info := unsafe.Pointer(info)
}

func TypeCheckClassCast(gClass *TypeClass, isAType uint64) {
	sys_gClass := unsafe.Pointer(gClass)
	sys_isAType := uint64(isAType)
}

func TypeCheckClassIsA(gClass *TypeClass, isAType uint64) {
	sys_gClass := unsafe.Pointer(gClass)
	sys_isAType := uint64(isAType)
}

func TypeCheckInstance(instance *TypeInstance) {
	sys_instance := unsafe.Pointer(instance)
}

func TypeCheckInstanceCast(instance *TypeInstance, ifaceType uint64) {
	sys_instance := unsafe.Pointer(instance)
	sys_ifaceType := uint64(ifaceType)
}

func TypeCheckInstanceIsA(instance *TypeInstance, ifaceType uint64) {
	sys_instance := unsafe.Pointer(instance)
	sys_ifaceType := uint64(ifaceType)
}

func TypeCheckInstanceIsFundamentallyA(instance *TypeInstance, fundamentalType uint64) {
	sys_instance := unsafe.Pointer(instance)
	sys_fundamentalType := uint64(fundamentalType)
}

func TypeCheckIsValueType(type_ uint64) {
	sys_type_ := uint64(type_)
}

func TypeCheckValue(value *Value) {
	sys_value := unsafe.Pointer(value)
}

func TypeCheckValueHolds(value *Value, type_ uint64) {
	sys_value := unsafe.Pointer(value)
	sys_type_ := uint64(type_)
}

func TypeChildren(type_ uint64) {
	sys_type_ := uint64(type_)
}

func TypeCreateInstance(type_ uint64) {
	sys_type_ := uint64(type_)
}

func TypeDefaultInterfacePeek(gType uint64) {
	sys_gType := uint64(gType)
}

func TypeDefaultInterfaceRef(gType uint64) {
	sys_gType := uint64(gType)
}

func TypeDefaultInterfaceUnref(gIface unsafe.Pointer) {
	sys_gIface := unsafe.Pointer(gIface)
}

func TypeDepth(type_ uint64) {
	sys_type_ := uint64(type_)
}

func TypeEnsure(type_ uint64) {
	sys_type_ := uint64(type_)
}

func TypeFreeInstance(instance *TypeInstance) {
	sys_instance := unsafe.Pointer(instance)
}

func TypeFromName(name string) {
	sys_name := string(name)
}

func TypeFundamental(typeId uint64) {
	sys_typeId := uint64(typeId)
}

func TypeFundamentalNext() {}

func TypeGetInstanceCount(type_ uint64) {
	sys_type_ := uint64(type_)
}

func TypeGetPlugin(type_ uint64) {
	sys_type_ := uint64(type_)
}

func TypeGetQdata(type_ uint64, quark uint32) {
	sys_type_ := uint64(type_)
	sys_quark := uint32(quark)
}

func TypeGetTypeRegistrationSerial() {}

func TypeInit() {}

func TypeInitWithDebugFlags(debugFlags int) {
	sys_debugFlags := int(debugFlags)
}

func TypeInterfaces(type_ uint64) {
	sys_type_ := uint64(type_)
}

func TypeIsA(type_ uint64, isAType uint64) {
	sys_type_ := uint64(type_)
	sys_isAType := uint64(isAType)
}

func TypeName(type_ uint64) {
	sys_type_ := uint64(type_)
}

func TypeNameFromClass(gClass *TypeClass) {
	sys_gClass := unsafe.Pointer(gClass)
}

func TypeNameFromInstance(instance *TypeInstance) {
	sys_instance := unsafe.Pointer(instance)
}

func TypeNextBase(leafType uint64, rootType uint64) {
	sys_leafType := uint64(leafType)
	sys_rootType := uint64(rootType)
}

func TypeParent(type_ uint64) {
	sys_type_ := uint64(type_)
}

func TypeQname(type_ uint64) {
	sys_type_ := uint64(type_)
}

func TypeQuery(type_ uint64) {
	sys_type_ := uint64(type_)
}

func TypeRegisterDynamic(parentType uint64, typeName string, plugin *TypePlugin, flags int) {
	sys_parentType := uint64(parentType)
	sys_typeName := string(typeName)
	sys_plugin := unsafe.Pointer(plugin)
	sys_flags := int(flags)
}

func TypeRegisterFundamental(typeId uint64, typeName string, info *TypeInfo, finfo *TypeFundamentalInfo, flags int) {
	sys_typeId := uint64(typeId)
	sys_typeName := string(typeName)
	sys_info := unsafe.Pointer(info)
	sys_finfo := unsafe.Pointer(finfo)
	sys_flags := int(flags)
}

func TypeRegisterStatic(parentType uint64, typeName string, info *TypeInfo, flags int) {
	sys_parentType := uint64(parentType)
	sys_typeName := string(typeName)
	sys_info := unsafe.Pointer(info)
	sys_flags := int(flags)
}

// UNSUPPORTED : g_type_register_static_simple : parameter 'class_init' is callback

// UNSUPPORTED : g_type_remove_class_cache_func : parameter 'cache_func' is callback

// UNSUPPORTED : g_type_remove_interface_check : parameter 'check_func' is callback

func TypeSetQdata(type_ uint64, quark uint32, data unsafe.Pointer) {
	sys_type_ := uint64(type_)
	sys_quark := uint32(quark)
	sys_data := unsafe.Pointer(data)
}

func TypeTestFlags(type_ uint64, flags uint) {
	sys_type_ := uint64(type_)
	sys_flags := uint(flags)
}

// UNSUPPORTED : g_value_register_transform_func : parameter 'transform_func' is callback

// CClosure is a representation of the C record GCClosure.
type CClosure struct {
	native unsafe.Pointer
}

// Closure is a representation of the C record GClosure.
type Closure struct {
	native unsafe.Pointer
}

// ClosureNotifyData is a representation of the C record GClosureNotifyData.
type ClosureNotifyData struct {
	native unsafe.Pointer
}

// EnumClass is a representation of the C record GEnumClass.
type EnumClass struct {
	native unsafe.Pointer
}

// EnumValue is a representation of the C record GEnumValue.
type EnumValue struct {
	native unsafe.Pointer
}

// FlagsClass is a representation of the C record GFlagsClass.
type FlagsClass struct {
	native unsafe.Pointer
}

// FlagsValue is a representation of the C record GFlagsValue.
type FlagsValue struct {
	native unsafe.Pointer
}

// InitiallyUnownedClass is a representation of the C record GInitiallyUnownedClass.
type InitiallyUnownedClass struct {
	native unsafe.Pointer
}

// InterfaceInfo is a representation of the C record GInterfaceInfo.
type InterfaceInfo struct {
	native unsafe.Pointer
}

// ObjectClass is a representation of the C record GObjectClass.
type ObjectClass struct {
	native unsafe.Pointer
}

// ObjectConstructParam is a representation of the C record GObjectConstructParam.
type ObjectConstructParam struct {
	native unsafe.Pointer
}

// ParamSpecClass is a representation of the C record GParamSpecClass.
type ParamSpecClass struct {
	native unsafe.Pointer
}

// ParamSpecPool is a representation of the C record GParamSpecPool.
type ParamSpecPool struct {
	native unsafe.Pointer
}

// ParamSpecTypeInfo is a representation of the C record GParamSpecTypeInfo.
type ParamSpecTypeInfo struct {
	native unsafe.Pointer
}

// Parameter is a representation of the C record GParameter.
type Parameter struct {
	native unsafe.Pointer
}

// SignalInvocationHint is a representation of the C record GSignalInvocationHint.
type SignalInvocationHint struct {
	native unsafe.Pointer
}

// SignalQuery is a representation of the C record GSignalQuery.
type SignalQuery struct {
	native unsafe.Pointer
}

// TypeClass is a representation of the C record GTypeClass.
type TypeClass struct {
	native unsafe.Pointer
}

// TypeFundamentalInfo is a representation of the C record GTypeFundamentalInfo.
type TypeFundamentalInfo struct {
	native unsafe.Pointer
}

// TypeInfo is a representation of the C record GTypeInfo.
type TypeInfo struct {
	native unsafe.Pointer
}

// TypeInstance is a representation of the C record GTypeInstance.
type TypeInstance struct {
	native unsafe.Pointer
}

// TypeInterface is a representation of the C record GTypeInterface.
type TypeInterface struct {
	native unsafe.Pointer
}

// TypeModuleClass is a representation of the C record GTypeModuleClass.
type TypeModuleClass struct {
	native unsafe.Pointer
}

// TypePluginClass is a representation of the C record GTypePluginClass.
type TypePluginClass struct {
	native unsafe.Pointer
}

// TypeQuery is a representation of the C record GTypeQuery.
type TypeQuery struct {
	native unsafe.Pointer
}

// TypeValueTable is a representation of the C record GTypeValueTable.
type TypeValueTable struct {
	native unsafe.Pointer
}

// Value is a representation of the C record GValue.
type Value struct {
	native unsafe.Pointer
}

// ValueArray is a representation of the C record GValueArray.
type ValueArray struct {
	native unsafe.Pointer
}

// WeakRef is a representation of the C record GWeakRef.
type WeakRef struct {
	native unsafe.Pointer
}

// Binding is a representation of the C record GBinding.
//
// since 2.26
type Binding struct {
	native unsafe.Pointer
}

// InitiallyUnowned is a representation of the C record GInitiallyUnowned.
type InitiallyUnowned struct {
	native unsafe.Pointer
}

// Object is a representation of the C record GObject.
type Object struct {
	native unsafe.Pointer
}

// ParamSpec is a representation of the C record GParamSpec.
type ParamSpec struct {
	native unsafe.Pointer
}

// ParamSpecBoolean is a representation of the C record GParamSpecBoolean.
type ParamSpecBoolean struct {
	native unsafe.Pointer
}

// ParamSpecBoxed is a representation of the C record GParamSpecBoxed.
type ParamSpecBoxed struct {
	native unsafe.Pointer
}

// ParamSpecChar is a representation of the C record GParamSpecChar.
type ParamSpecChar struct {
	native unsafe.Pointer
}

// ParamSpecDouble is a representation of the C record GParamSpecDouble.
type ParamSpecDouble struct {
	native unsafe.Pointer
}

// ParamSpecEnum is a representation of the C record GParamSpecEnum.
type ParamSpecEnum struct {
	native unsafe.Pointer
}

// ParamSpecFlags is a representation of the C record GParamSpecFlags.
type ParamSpecFlags struct {
	native unsafe.Pointer
}

// ParamSpecFloat is a representation of the C record GParamSpecFloat.
type ParamSpecFloat struct {
	native unsafe.Pointer
}

// ParamSpecGType is a representation of the C record GParamSpecGType.
//
// since 2.10
type ParamSpecGType struct {
	native unsafe.Pointer
}

// ParamSpecInt is a representation of the C record GParamSpecInt.
type ParamSpecInt struct {
	native unsafe.Pointer
}

// ParamSpecInt64 is a representation of the C record GParamSpecInt64.
type ParamSpecInt64 struct {
	native unsafe.Pointer
}

// ParamSpecLong is a representation of the C record GParamSpecLong.
type ParamSpecLong struct {
	native unsafe.Pointer
}

// ParamSpecObject is a representation of the C record GParamSpecObject.
type ParamSpecObject struct {
	native unsafe.Pointer
}

// ParamSpecOverride is a representation of the C record GParamSpecOverride.
//
// since 2.4
type ParamSpecOverride struct {
	native unsafe.Pointer
}

// ParamSpecParam is a representation of the C record GParamSpecParam.
type ParamSpecParam struct {
	native unsafe.Pointer
}

// ParamSpecPointer is a representation of the C record GParamSpecPointer.
type ParamSpecPointer struct {
	native unsafe.Pointer
}

// ParamSpecString is a representation of the C record GParamSpecString.
type ParamSpecString struct {
	native unsafe.Pointer
}

// ParamSpecUChar is a representation of the C record GParamSpecUChar.
type ParamSpecUChar struct {
	native unsafe.Pointer
}

// ParamSpecUInt is a representation of the C record GParamSpecUInt.
type ParamSpecUInt struct {
	native unsafe.Pointer
}

// ParamSpecUInt64 is a representation of the C record GParamSpecUInt64.
type ParamSpecUInt64 struct {
	native unsafe.Pointer
}

// ParamSpecULong is a representation of the C record GParamSpecULong.
type ParamSpecULong struct {
	native unsafe.Pointer
}

// ParamSpecUnichar is a representation of the C record GParamSpecUnichar.
type ParamSpecUnichar struct {
	native unsafe.Pointer
}

// ParamSpecValueArray is a representation of the C record GParamSpecValueArray.
type ParamSpecValueArray struct {
	native unsafe.Pointer
}

// ParamSpecVariant is a representation of the C record GParamSpecVariant.
//
// since 2.26
type ParamSpecVariant struct {
	native unsafe.Pointer
}

// TypeModule is a representation of the C record GTypeModule.
type TypeModule struct {
	native unsafe.Pointer
}

// TypePlugin is a representation of the C interface GTypePlugin.
type TypePlugin struct {
	native unsafe.Pointer
}

// TypeCValue is a representation of the C union GTypeCValue.
type TypeCValue struct{}

// _Value__data__union is a representation of the C union .
type _Value__data__union struct{}
