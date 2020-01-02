// Code generated - DO NOT EDIT.
// +build gobject_2.32

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
	sys_boxedType := boxedType
	sys_srcBoxed := srcBoxed
}

func BoxedFree(boxedType uint64, boxed unsafe.Pointer) {
	sys_boxedType := boxedType
	sys_boxed := boxed
}

// UNSUPPORTED : g_boxed_type_register_static : parameter 'boxed_copy' is callback

// UNSUPPORTED : g_cclosure_new : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object_swap : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_swap : parameter 'callback_func' is callback

func ClearObject(objectPtr **Object) {
	sys_objectPtr := objectPtr.ToC()
}

func EnumCompleteTypeInfo(gEnumType uint64, constValues *EnumValue) {
	sys_gEnumType := gEnumType
	sys_constValues := constValues.ToC()
}

func EnumGetValue(enumClass *EnumClass, value int) {
	sys_enumClass := enumClass.ToC()
	sys_value := value
}

func EnumGetValueByName(enumClass *EnumClass, name string) {
	sys_enumClass := enumClass.ToC()
	sys_name := name
}

func EnumGetValueByNick(enumClass *EnumClass, nick string) {
	sys_enumClass := enumClass.ToC()
	sys_nick := nick
}

func EnumRegisterStatic(name string, constStaticValues *EnumValue) {
	sys_name := name
	sys_constStaticValues := constStaticValues.ToC()
}

func FlagsCompleteTypeInfo(gFlagsType uint64, constValues *FlagsValue) {
	sys_gFlagsType := gFlagsType
	sys_constValues := constValues.ToC()
}

func FlagsGetFirstValue(flagsClass *FlagsClass, value uint) {
	sys_flagsClass := flagsClass.ToC()
	sys_value := value
}

func FlagsGetValueByName(flagsClass *FlagsClass, name string) {
	sys_flagsClass := flagsClass.ToC()
	sys_name := name
}

func FlagsGetValueByNick(flagsClass *FlagsClass, nick string) {
	sys_flagsClass := flagsClass.ToC()
	sys_nick := nick
}

func FlagsRegisterStatic(name string, constStaticValues *FlagsValue) {
	sys_name := name
	sys_constStaticValues := constStaticValues.ToC()
}

func GtypeGetType() {}

func ParamSpecBoolean(name string, nick string, blurb string, defaultValue bool, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_defaultValue := defaultValue
	sys_flags := flags
}

func ParamSpecBoxed(name string, nick string, blurb string, boxedType uint64, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_boxedType := boxedType
	sys_flags := flags
}

func ParamSpecChar(name string, nick string, blurb string, minimum int8, maximum int8, defaultValue int8, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
}

func ParamSpecDouble(name string, nick string, blurb string, minimum float64, maximum float64, defaultValue float64, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
}

func ParamSpecEnum(name string, nick string, blurb string, enumType uint64, defaultValue int, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_enumType := enumType
	sys_defaultValue := defaultValue
	sys_flags := flags
}

func ParamSpecFlags(name string, nick string, blurb string, flagsType uint64, defaultValue uint, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_flagsType := flagsType
	sys_defaultValue := defaultValue
	sys_flags := flags
}

func ParamSpecFloat(name string, nick string, blurb string, minimum float32, maximum float32, defaultValue float32, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
}

func ParamSpecGtype(name string, nick string, blurb string, isAType uint64, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_isAType := isAType
	sys_flags := flags
}

func ParamSpecInt(name string, nick string, blurb string, minimum int, maximum int, defaultValue int, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
}

func ParamSpecInt64(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
}

func ParamSpecLong(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
}

func ParamSpecObject(name string, nick string, blurb string, objectType uint64, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_objectType := objectType
	sys_flags := flags
}

func ParamSpecOverride(name string, overridden *ParamSpec) {
	sys_name := name
	sys_overridden := overridden.ToC()
}

func ParamSpecParam(name string, nick string, blurb string, paramType uint64, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_paramType := paramType
	sys_flags := flags
}

func ParamSpecPointer(name string, nick string, blurb string, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_flags := flags
}

func ParamSpecString(name string, nick string, blurb string, defaultValue string, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_defaultValue := defaultValue
	sys_flags := flags
}

func ParamSpecUchar(name string, nick string, blurb string, minimum uint8, maximum uint8, defaultValue uint8, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
}

func ParamSpecUint(name string, nick string, blurb string, minimum uint, maximum uint, defaultValue uint, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
}

func ParamSpecUint64(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
}

func ParamSpecUlong(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
}

func ParamSpecUnichar(name string, nick string, blurb string, defaultValue rune, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_defaultValue := defaultValue
	sys_flags := flags
}

func ParamSpecValueArray(name string, nick string, blurb string, elementSpec *ParamSpec, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_elementSpec := elementSpec.ToC()
	sys_flags := flags
}

func ParamSpecVariant(name string, nick string, blurb string, type_ *glib.VariantType, defaultValue *glib.Variant, flags int) {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_type_ := type_.ToC()
	sys_defaultValue := defaultValue.ToC()
	sys_flags := flags
}

func ParamTypeRegisterStatic(name string, pspecInfo *ParamSpecTypeInfo) {
	sys_name := name
	sys_pspecInfo := pspecInfo.ToC()
}

func ParamValueConvert(pspec *ParamSpec, srcValue *Value, destValue *Value, strictValidation bool) {
	sys_pspec := pspec.ToC()
	sys_srcValue := srcValue.ToC()
	sys_destValue := destValue.ToC()
	sys_strictValidation := strictValidation
}

func ParamValueDefaults(pspec *ParamSpec, value *Value) {
	sys_pspec := pspec.ToC()
	sys_value := value.ToC()
}

func ParamValueSetDefault(pspec *ParamSpec, value *Value) {
	sys_pspec := pspec.ToC()
	sys_value := value.ToC()
}

func ParamValueValidate(pspec *ParamSpec, value *Value) {
	sys_pspec := pspec.ToC()
	sys_value := value.ToC()
}

func ParamValuesCmp(pspec *ParamSpec, value1 *Value, value2 *Value) {
	sys_pspec := pspec.ToC()
	sys_value1 := value1.ToC()
	sys_value2 := value2.ToC()
}

func PointerTypeRegisterStatic(name string) {
	sys_name := name
}

func SignalAccumulatorFirstWins(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value, dummy unsafe.Pointer) {
	sys_ihint := ihint.ToC()
	sys_returnAccu := returnAccu.ToC()
	sys_handlerReturn := handlerReturn.ToC()
	sys_dummy := dummy
}

func SignalAccumulatorTrueHandled(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value, dummy unsafe.Pointer) {
	sys_ihint := ihint.ToC()
	sys_returnAccu := returnAccu.ToC()
	sys_handlerReturn := handlerReturn.ToC()
	sys_dummy := dummy
}

// UNSUPPORTED : g_signal_add_emission_hook : parameter 'hook_func' is callback

// UNSUPPORTED : g_signal_chain_from_overridden : parameter 'instance_and_params' is array parameter without length parameter

func SignalChainFromOverriddenHandler(instance unsafe.Pointer) {
	sys_instance := instance.ToC()
}

func SignalConnectClosure(instance unsafe.Pointer, detailedSignal string, closure *Closure, after bool) {
	sys_instance := instance.ToC()
	sys_detailedSignal := detailedSignal
	sys_closure := closure.ToC()
	sys_after := after
}

func SignalConnectClosureById(instance unsafe.Pointer, signalId uint, detail uint32, closure *Closure, after bool) {
	sys_instance := instance.ToC()
	sys_signalId := signalId
	sys_detail := detail
	sys_closure := closure.ToC()
	sys_after := after
}

// UNSUPPORTED : g_signal_connect_data : parameter 'c_handler' is callback

// UNSUPPORTED : g_signal_connect_object : parameter 'c_handler' is callback

func SignalEmit(instance unsafe.Pointer, signalId uint, detail uint32) {
	sys_instance := instance.ToC()
	sys_signalId := signalId
	sys_detail := detail
}

func SignalEmitByName(instance unsafe.Pointer, detailedSignal string) {
	sys_instance := instance.ToC()
	sys_detailedSignal := detailedSignal
}

func SignalEmitValist(instance unsafe.Pointer, signalId uint, detail uint32) {
	sys_instance := instance.ToC()
	sys_signalId := signalId
	sys_detail := detail
}

// UNSUPPORTED : g_signal_emitv : parameter 'instance_and_params' is array parameter without length parameter

func SignalGetInvocationHint(instance unsafe.Pointer) {
	sys_instance := instance.ToC()
}

func SignalHandlerBlock(instance unsafe.Pointer, handlerId uint64) {
	sys_instance := instance.ToC()
	sys_handlerId := handlerId
}

func SignalHandlerDisconnect(instance unsafe.Pointer, handlerId uint64) {
	sys_instance := instance.ToC()
	sys_handlerId := handlerId
}

func SignalHandlerFind(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) {
	sys_instance := instance.ToC()
	sys_mask := mask
	sys_signalId := signalId
	sys_detail := detail
	sys_closure := closure.ToC()
	sys_func_ := func_
	sys_data := data
}

func SignalHandlerIsConnected(instance unsafe.Pointer, handlerId uint64) {
	sys_instance := instance.ToC()
	sys_handlerId := handlerId
}

func SignalHandlerUnblock(instance unsafe.Pointer, handlerId uint64) {
	sys_instance := instance.ToC()
	sys_handlerId := handlerId
}

func SignalHandlersBlockMatched(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) {
	sys_instance := instance.ToC()
	sys_mask := mask
	sys_signalId := signalId
	sys_detail := detail
	sys_closure := closure.ToC()
	sys_func_ := func_
	sys_data := data
}

func SignalHandlersDestroy(instance unsafe.Pointer) {
	sys_instance := instance.ToC()
}

func SignalHandlersDisconnectMatched(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) {
	sys_instance := instance.ToC()
	sys_mask := mask
	sys_signalId := signalId
	sys_detail := detail
	sys_closure := closure.ToC()
	sys_func_ := func_
	sys_data := data
}

func SignalHandlersUnblockMatched(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) {
	sys_instance := instance.ToC()
	sys_mask := mask
	sys_signalId := signalId
	sys_detail := detail
	sys_closure := closure.ToC()
	sys_func_ := func_
	sys_data := data
}

func SignalHasHandlerPending(instance unsafe.Pointer, signalId uint, detail uint32, mayBeBlocked bool) {
	sys_instance := instance.ToC()
	sys_signalId := signalId
	sys_detail := detail
	sys_mayBeBlocked := mayBeBlocked
}

func SignalListIds(itype uint64) {
	sys_itype := itype
}

func SignalLookup(name string, itype uint64) {
	sys_name := name
	sys_itype := itype
}

func SignalName(signalId uint) {
	sys_signalId := signalId
}

// UNSUPPORTED : g_signal_new : parameter 'accumulator' is callback

// UNSUPPORTED : g_signal_new_class_handler : parameter 'class_handler' is callback

// UNSUPPORTED : g_signal_new_valist : parameter 'accumulator' is callback

// UNSUPPORTED : g_signal_newv : parameter 'accumulator' is callback

func SignalOverrideClassClosure(signalId uint, instanceType uint64, classClosure *Closure) {
	sys_signalId := signalId
	sys_instanceType := instanceType
	sys_classClosure := classClosure.ToC()
}

// UNSUPPORTED : g_signal_override_class_handler : parameter 'class_handler' is callback

func SignalParseName(detailedSignal string, itype uint64, forceDetailQuark bool) {
	sys_detailedSignal := detailedSignal
	sys_itype := itype
	sys_forceDetailQuark := forceDetailQuark
}

func SignalQuery(signalId uint) {
	sys_signalId := signalId
}

func SignalRemoveEmissionHook(signalId uint, hookId uint64) {
	sys_signalId := signalId
	sys_hookId := hookId
}

// UNSUPPORTED : g_signal_set_va_marshaller : blacklisted

func SignalStopEmission(instance unsafe.Pointer, signalId uint, detail uint32) {
	sys_instance := instance.ToC()
	sys_signalId := signalId
	sys_detail := detail
}

func SignalStopEmissionByName(instance unsafe.Pointer, detailedSignal string) {
	sys_instance := instance.ToC()
	sys_detailedSignal := detailedSignal
}

func SignalTypeCclosureNew(itype uint64, structOffset uint) {
	sys_itype := itype
	sys_structOffset := structOffset
}

func SourceSetClosure(source *glib.Source, closure *Closure) {
	sys_source := source.ToC()
	sys_closure := closure.ToC()
}

func SourceSetDummyCallback(source *glib.Source) {
	sys_source := source.ToC()
}

func StrdupValueContents(value *Value) {
	sys_value := value.ToC()
}

// UNSUPPORTED : g_type_add_class_cache_func : parameter 'cache_func' is callback

func TypeAddClassPrivate(classType uint64, privateSize uint64) {
	sys_classType := classType
	sys_privateSize := privateSize
}

func TypeAddInstancePrivate(classType uint64, privateSize uint64) {
	sys_classType := classType
	sys_privateSize := privateSize
}

// UNSUPPORTED : g_type_add_interface_check : parameter 'check_func' is callback

func TypeAddInterfaceDynamic(instanceType uint64, interfaceType uint64, plugin *TypePlugin) {
	sys_instanceType := instanceType
	sys_interfaceType := interfaceType
	sys_plugin := plugin.ToC()
}

func TypeAddInterfaceStatic(instanceType uint64, interfaceType uint64, info *InterfaceInfo) {
	sys_instanceType := instanceType
	sys_interfaceType := interfaceType
	sys_info := info.ToC()
}

func TypeCheckClassCast(gClass *TypeClass, isAType uint64) {
	sys_gClass := gClass.ToC()
	sys_isAType := isAType
}

func TypeCheckClassIsA(gClass *TypeClass, isAType uint64) {
	sys_gClass := gClass.ToC()
	sys_isAType := isAType
}

func TypeCheckInstance(instance *TypeInstance) {
	sys_instance := instance.ToC()
}

func TypeCheckInstanceCast(instance *TypeInstance, ifaceType uint64) {
	sys_instance := instance.ToC()
	sys_ifaceType := ifaceType
}

func TypeCheckInstanceIsA(instance *TypeInstance, ifaceType uint64) {
	sys_instance := instance.ToC()
	sys_ifaceType := ifaceType
}

func TypeCheckInstanceIsFundamentallyA(instance *TypeInstance, fundamentalType uint64) {
	sys_instance := instance.ToC()
	sys_fundamentalType := fundamentalType
}

func TypeCheckIsValueType(type_ uint64) {
	sys_type_ := type_
}

func TypeCheckValue(value *Value) {
	sys_value := value.ToC()
}

func TypeCheckValueHolds(value *Value, type_ uint64) {
	sys_value := value.ToC()
	sys_type_ := type_
}

func TypeChildren(type_ uint64) {
	sys_type_ := type_
}

func TypeCreateInstance(type_ uint64) {
	sys_type_ := type_
}

func TypeDefaultInterfacePeek(gType uint64) {
	sys_gType := gType
}

func TypeDefaultInterfaceRef(gType uint64) {
	sys_gType := gType
}

func TypeDefaultInterfaceUnref(gIface unsafe.Pointer) {
	sys_gIface := gIface.ToC()
}

func TypeDepth(type_ uint64) {
	sys_type_ := type_
}

func TypeFreeInstance(instance *TypeInstance) {
	sys_instance := instance.ToC()
}

func TypeFromName(name string) {
	sys_name := name
}

func TypeFundamental(typeId uint64) {
	sys_typeId := typeId
}

func TypeFundamentalNext() {}

func TypeGetPlugin(type_ uint64) {
	sys_type_ := type_
}

func TypeGetQdata(type_ uint64, quark uint32) {
	sys_type_ := type_
	sys_quark := quark
}

func TypeInit() {}

func TypeInitWithDebugFlags(debugFlags int) {
	sys_debugFlags := debugFlags
}

func TypeInterfaces(type_ uint64) {
	sys_type_ := type_
}

func TypeIsA(type_ uint64, isAType uint64) {
	sys_type_ := type_
	sys_isAType := isAType
}

func TypeName(type_ uint64) {
	sys_type_ := type_
}

func TypeNameFromClass(gClass *TypeClass) {
	sys_gClass := gClass.ToC()
}

func TypeNameFromInstance(instance *TypeInstance) {
	sys_instance := instance.ToC()
}

func TypeNextBase(leafType uint64, rootType uint64) {
	sys_leafType := leafType
	sys_rootType := rootType
}

func TypeParent(type_ uint64) {
	sys_type_ := type_
}

func TypeQname(type_ uint64) {
	sys_type_ := type_
}

func TypeQuery(type_ uint64) {
	sys_type_ := type_
}

func TypeRegisterDynamic(parentType uint64, typeName string, plugin *TypePlugin, flags int) {
	sys_parentType := parentType
	sys_typeName := typeName
	sys_plugin := plugin.ToC()
	sys_flags := flags
}

func TypeRegisterFundamental(typeId uint64, typeName string, info *TypeInfo, finfo *TypeFundamentalInfo, flags int) {
	sys_typeId := typeId
	sys_typeName := typeName
	sys_info := info.ToC()
	sys_finfo := finfo.ToC()
	sys_flags := flags
}

func TypeRegisterStatic(parentType uint64, typeName string, info *TypeInfo, flags int) {
	sys_parentType := parentType
	sys_typeName := typeName
	sys_info := info.ToC()
	sys_flags := flags
}

// UNSUPPORTED : g_type_register_static_simple : parameter 'class_init' is callback

// UNSUPPORTED : g_type_remove_class_cache_func : parameter 'cache_func' is callback

// UNSUPPORTED : g_type_remove_interface_check : parameter 'check_func' is callback

func TypeSetQdata(type_ uint64, quark uint32, data unsafe.Pointer) {
	sys_type_ := type_
	sys_quark := quark
	sys_data := data
}

func TypeTestFlags(type_ uint64, flags uint) {
	sys_type_ := type_
	sys_flags := flags
}

// UNSUPPORTED : g_value_register_transform_func : parameter 'transform_func' is callback

// CClosure is a representation of the C record GCClosure.
type CClosure struct {
	native unsafe.Pointer
}

func (recv *CClosure) ToC() unsafe.Pointer {
	return recv.native
}

// Closure is a representation of the C record GClosure.
type Closure struct {
	native unsafe.Pointer
}

func (recv *Closure) ToC() unsafe.Pointer {
	return recv.native
}

// ClosureNotifyData is a representation of the C record GClosureNotifyData.
type ClosureNotifyData struct {
	native unsafe.Pointer
}

func (recv *ClosureNotifyData) ToC() unsafe.Pointer {
	return recv.native
}

// EnumClass is a representation of the C record GEnumClass.
type EnumClass struct {
	native unsafe.Pointer
}

func (recv *EnumClass) ToC() unsafe.Pointer {
	return recv.native
}

// EnumValue is a representation of the C record GEnumValue.
type EnumValue struct {
	native unsafe.Pointer
}

func (recv *EnumValue) ToC() unsafe.Pointer {
	return recv.native
}

// FlagsClass is a representation of the C record GFlagsClass.
type FlagsClass struct {
	native unsafe.Pointer
}

func (recv *FlagsClass) ToC() unsafe.Pointer {
	return recv.native
}

// FlagsValue is a representation of the C record GFlagsValue.
type FlagsValue struct {
	native unsafe.Pointer
}

func (recv *FlagsValue) ToC() unsafe.Pointer {
	return recv.native
}

// InitiallyUnownedClass is a representation of the C record GInitiallyUnownedClass.
type InitiallyUnownedClass struct {
	native unsafe.Pointer
}

func (recv *InitiallyUnownedClass) ToC() unsafe.Pointer {
	return recv.native
}

// InterfaceInfo is a representation of the C record GInterfaceInfo.
type InterfaceInfo struct {
	native unsafe.Pointer
}

func (recv *InterfaceInfo) ToC() unsafe.Pointer {
	return recv.native
}

// ObjectClass is a representation of the C record GObjectClass.
type ObjectClass struct {
	native unsafe.Pointer
}

func (recv *ObjectClass) ToC() unsafe.Pointer {
	return recv.native
}

// ObjectConstructParam is a representation of the C record GObjectConstructParam.
type ObjectConstructParam struct {
	native unsafe.Pointer
}

func (recv *ObjectConstructParam) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecClass is a representation of the C record GParamSpecClass.
type ParamSpecClass struct {
	native unsafe.Pointer
}

func (recv *ParamSpecClass) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecPool is a representation of the C record GParamSpecPool.
type ParamSpecPool struct {
	native unsafe.Pointer
}

func (recv *ParamSpecPool) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecTypeInfo is a representation of the C record GParamSpecTypeInfo.
type ParamSpecTypeInfo struct {
	native unsafe.Pointer
}

func (recv *ParamSpecTypeInfo) ToC() unsafe.Pointer {
	return recv.native
}

// Parameter is a representation of the C record GParameter.
type Parameter struct {
	native unsafe.Pointer
}

func (recv *Parameter) ToC() unsafe.Pointer {
	return recv.native
}

// SignalInvocationHint is a representation of the C record GSignalInvocationHint.
type SignalInvocationHint struct {
	native unsafe.Pointer
}

func (recv *SignalInvocationHint) ToC() unsafe.Pointer {
	return recv.native
}

// SignalQuery is a representation of the C record GSignalQuery.
type SignalQuery struct {
	native unsafe.Pointer
}

func (recv *SignalQuery) ToC() unsafe.Pointer {
	return recv.native
}

// TypeClass is a representation of the C record GTypeClass.
type TypeClass struct {
	native unsafe.Pointer
}

func (recv *TypeClass) ToC() unsafe.Pointer {
	return recv.native
}

// TypeFundamentalInfo is a representation of the C record GTypeFundamentalInfo.
type TypeFundamentalInfo struct {
	native unsafe.Pointer
}

func (recv *TypeFundamentalInfo) ToC() unsafe.Pointer {
	return recv.native
}

// TypeInfo is a representation of the C record GTypeInfo.
type TypeInfo struct {
	native unsafe.Pointer
}

func (recv *TypeInfo) ToC() unsafe.Pointer {
	return recv.native
}

// TypeInstance is a representation of the C record GTypeInstance.
type TypeInstance struct {
	native unsafe.Pointer
}

func (recv *TypeInstance) ToC() unsafe.Pointer {
	return recv.native
}

// TypeInterface is a representation of the C record GTypeInterface.
type TypeInterface struct {
	native unsafe.Pointer
}

func (recv *TypeInterface) ToC() unsafe.Pointer {
	return recv.native
}

// TypeModuleClass is a representation of the C record GTypeModuleClass.
type TypeModuleClass struct {
	native unsafe.Pointer
}

func (recv *TypeModuleClass) ToC() unsafe.Pointer {
	return recv.native
}

// TypePluginClass is a representation of the C record GTypePluginClass.
type TypePluginClass struct {
	native unsafe.Pointer
}

func (recv *TypePluginClass) ToC() unsafe.Pointer {
	return recv.native
}

// TypeQuery is a representation of the C record GTypeQuery.
type TypeQuery struct {
	native unsafe.Pointer
}

func (recv *TypeQuery) ToC() unsafe.Pointer {
	return recv.native
}

// TypeValueTable is a representation of the C record GTypeValueTable.
type TypeValueTable struct {
	native unsafe.Pointer
}

func (recv *TypeValueTable) ToC() unsafe.Pointer {
	return recv.native
}

// Value is a representation of the C record GValue.
type Value struct {
	native unsafe.Pointer
}

func (recv *Value) ToC() unsafe.Pointer {
	return recv.native
}

// ValueArray is a representation of the C record GValueArray.
type ValueArray struct {
	native unsafe.Pointer
}

func (recv *ValueArray) ToC() unsafe.Pointer {
	return recv.native
}

// WeakRef is a representation of the C record GWeakRef.
type WeakRef struct {
	native unsafe.Pointer
}

func (recv *WeakRef) ToC() unsafe.Pointer {
	return recv.native
}

// Binding is a representation of the C record GBinding.
//
// since 2.26
type Binding struct {
	native unsafe.Pointer
}

func (recv *Binding) ToC() unsafe.Pointer {
	return recv.native
}

// InitiallyUnowned is a representation of the C record GInitiallyUnowned.
type InitiallyUnowned struct {
	native unsafe.Pointer
}

func (recv *InitiallyUnowned) ToC() unsafe.Pointer {
	return recv.native
}

// Object is a representation of the C record GObject.
type Object struct {
	native unsafe.Pointer
}

func (recv *Object) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpec is a representation of the C record GParamSpec.
type ParamSpec struct {
	native unsafe.Pointer
}

func (recv *ParamSpec) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecBoolean is a representation of the C record GParamSpecBoolean.
type ParamSpecBoolean struct {
	native unsafe.Pointer
}

func (recv *ParamSpecBoolean) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecBoxed is a representation of the C record GParamSpecBoxed.
type ParamSpecBoxed struct {
	native unsafe.Pointer
}

func (recv *ParamSpecBoxed) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecChar is a representation of the C record GParamSpecChar.
type ParamSpecChar struct {
	native unsafe.Pointer
}

func (recv *ParamSpecChar) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecDouble is a representation of the C record GParamSpecDouble.
type ParamSpecDouble struct {
	native unsafe.Pointer
}

func (recv *ParamSpecDouble) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecEnum is a representation of the C record GParamSpecEnum.
type ParamSpecEnum struct {
	native unsafe.Pointer
}

func (recv *ParamSpecEnum) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecFlags is a representation of the C record GParamSpecFlags.
type ParamSpecFlags struct {
	native unsafe.Pointer
}

func (recv *ParamSpecFlags) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecFloat is a representation of the C record GParamSpecFloat.
type ParamSpecFloat struct {
	native unsafe.Pointer
}

func (recv *ParamSpecFloat) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecGType is a representation of the C record GParamSpecGType.
//
// since 2.10
type ParamSpecGType struct {
	native unsafe.Pointer
}

func (recv *ParamSpecGType) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecInt is a representation of the C record GParamSpecInt.
type ParamSpecInt struct {
	native unsafe.Pointer
}

func (recv *ParamSpecInt) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecInt64 is a representation of the C record GParamSpecInt64.
type ParamSpecInt64 struct {
	native unsafe.Pointer
}

func (recv *ParamSpecInt64) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecLong is a representation of the C record GParamSpecLong.
type ParamSpecLong struct {
	native unsafe.Pointer
}

func (recv *ParamSpecLong) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecObject is a representation of the C record GParamSpecObject.
type ParamSpecObject struct {
	native unsafe.Pointer
}

func (recv *ParamSpecObject) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecOverride is a representation of the C record GParamSpecOverride.
//
// since 2.4
type ParamSpecOverride struct {
	native unsafe.Pointer
}

func (recv *ParamSpecOverride) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecParam is a representation of the C record GParamSpecParam.
type ParamSpecParam struct {
	native unsafe.Pointer
}

func (recv *ParamSpecParam) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecPointer is a representation of the C record GParamSpecPointer.
type ParamSpecPointer struct {
	native unsafe.Pointer
}

func (recv *ParamSpecPointer) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecString is a representation of the C record GParamSpecString.
type ParamSpecString struct {
	native unsafe.Pointer
}

func (recv *ParamSpecString) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecUChar is a representation of the C record GParamSpecUChar.
type ParamSpecUChar struct {
	native unsafe.Pointer
}

func (recv *ParamSpecUChar) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecUInt is a representation of the C record GParamSpecUInt.
type ParamSpecUInt struct {
	native unsafe.Pointer
}

func (recv *ParamSpecUInt) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecUInt64 is a representation of the C record GParamSpecUInt64.
type ParamSpecUInt64 struct {
	native unsafe.Pointer
}

func (recv *ParamSpecUInt64) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecULong is a representation of the C record GParamSpecULong.
type ParamSpecULong struct {
	native unsafe.Pointer
}

func (recv *ParamSpecULong) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecUnichar is a representation of the C record GParamSpecUnichar.
type ParamSpecUnichar struct {
	native unsafe.Pointer
}

func (recv *ParamSpecUnichar) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecValueArray is a representation of the C record GParamSpecValueArray.
type ParamSpecValueArray struct {
	native unsafe.Pointer
}

func (recv *ParamSpecValueArray) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecVariant is a representation of the C record GParamSpecVariant.
//
// since 2.26
type ParamSpecVariant struct {
	native unsafe.Pointer
}

func (recv *ParamSpecVariant) ToC() unsafe.Pointer {
	return recv.native
}

// TypeModule is a representation of the C record GTypeModule.
type TypeModule struct {
	native unsafe.Pointer
}

func (recv *TypeModule) ToC() unsafe.Pointer {
	return recv.native
}

// TypePlugin is a representation of the C interface GTypePlugin.
type TypePlugin struct {
	native unsafe.Pointer
}

// TypeCValue is a representation of the C union GTypeCValue.
type TypeCValue struct{}

// _Value__data__union is a representation of the C union .
type _Value__data__union struct{}
