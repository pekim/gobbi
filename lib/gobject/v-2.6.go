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

func Fn_g_boxed_copy(boxedType uint64, srcBoxed unsafe.Pointer) {}

func Fn_g_boxed_free(boxedType uint64, boxed unsafe.Pointer) {}

// UNSUPPORTED : g_boxed_type_register_static : parameter 'boxed_copy' is callback

// UNSUPPORTED : g_cclosure_new : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object_swap : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_swap : parameter 'callback_func' is callback

func Fn_g_enum_complete_type_info(gEnumType uint64, constValues *EnumValue) {}

func Fn_g_enum_get_value(enumClass *EnumClass, value int) {}

func Fn_g_enum_get_value_by_name(enumClass *EnumClass, name string) {}

func Fn_g_enum_get_value_by_nick(enumClass *EnumClass, nick string) {}

func Fn_g_enum_register_static(name string, constStaticValues *EnumValue) {}

func Fn_g_flags_complete_type_info(gFlagsType uint64, constValues *FlagsValue) {}

func Fn_g_flags_get_first_value(flagsClass *FlagsClass, value uint) {}

func Fn_g_flags_get_value_by_name(flagsClass *FlagsClass, name string) {}

func Fn_g_flags_get_value_by_nick(flagsClass *FlagsClass, nick string) {}

func Fn_g_flags_register_static(name string, constStaticValues *FlagsValue) {}

func Fn_g_gtype_get_type() {}

func Fn_g_param_spec_boolean(name string, nick string, blurb string, defaultValue bool, flags int) {}

func Fn_g_param_spec_boxed(name string, nick string, blurb string, boxedType uint64, flags int) {}

func Fn_g_param_spec_char(name string, nick string, blurb string, minimum int8, maximum int8, defaultValue int8, flags int) {
}

func Fn_g_param_spec_double(name string, nick string, blurb string, minimum float64, maximum float64, defaultValue float64, flags int) {
}

func Fn_g_param_spec_enum(name string, nick string, blurb string, enumType uint64, defaultValue int, flags int) {
}

func Fn_g_param_spec_flags(name string, nick string, blurb string, flagsType uint64, defaultValue uint, flags int) {
}

func Fn_g_param_spec_float(name string, nick string, blurb string, minimum float32, maximum float32, defaultValue float32, flags int) {
}

func Fn_g_param_spec_int(name string, nick string, blurb string, minimum int, maximum int, defaultValue int, flags int) {
}

func Fn_g_param_spec_int64(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags int) {
}

func Fn_g_param_spec_long(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags int) {
}

func Fn_g_param_spec_object(name string, nick string, blurb string, objectType uint64, flags int) {}

func Fn_g_param_spec_override(name string, overridden *ParamSpec) {}

func Fn_g_param_spec_param(name string, nick string, blurb string, paramType uint64, flags int) {}

func Fn_g_param_spec_pointer(name string, nick string, blurb string, flags int) {}

func Fn_g_param_spec_string(name string, nick string, blurb string, defaultValue string, flags int) {
}

func Fn_g_param_spec_uchar(name string, nick string, blurb string, minimum uint8, maximum uint8, defaultValue uint8, flags int) {
}

func Fn_g_param_spec_uint(name string, nick string, blurb string, minimum uint, maximum uint, defaultValue uint, flags int) {
}

func Fn_g_param_spec_uint64(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags int) {
}

func Fn_g_param_spec_ulong(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags int) {
}

func Fn_g_param_spec_unichar(name string, nick string, blurb string, defaultValue rune, flags int) {}

func Fn_g_param_spec_value_array(name string, nick string, blurb string, elementSpec *ParamSpec, flags int) {
}

func Fn_g_param_type_register_static(name string, pspecInfo *ParamSpecTypeInfo) {}

func Fn_g_param_value_convert(pspec *ParamSpec, srcValue *Value, destValue *Value, strictValidation bool) {
}

func Fn_g_param_value_defaults(pspec *ParamSpec, value *Value) {}

func Fn_g_param_value_set_default(pspec *ParamSpec, value *Value) {}

func Fn_g_param_value_validate(pspec *ParamSpec, value *Value) {}

func Fn_g_param_values_cmp(pspec *ParamSpec, value1 *Value, value2 *Value) {}

func Fn_g_pointer_type_register_static(name string) {}

func Fn_g_signal_accumulator_true_handled(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value, dummy unsafe.Pointer) {
}

// UNSUPPORTED : g_signal_add_emission_hook : parameter 'hook_func' is callback

// UNSUPPORTED : g_signal_chain_from_overridden : parameter 'instance_and_params' is array parameter without length parameter

func Fn_g_signal_connect_closure(instance unsafe.Pointer, detailedSignal string, closure *Closure, after bool) {
}

func Fn_g_signal_connect_closure_by_id(instance unsafe.Pointer, signalId uint, detail uint32, closure *Closure, after bool) {
}

// UNSUPPORTED : g_signal_connect_data : parameter 'c_handler' is callback

// UNSUPPORTED : g_signal_connect_object : parameter 'c_handler' is callback

func Fn_g_signal_emit(instance unsafe.Pointer, signalId uint, detail uint32) {}

func Fn_g_signal_emit_by_name(instance unsafe.Pointer, detailedSignal string) {}

func Fn_g_signal_emit_valist(instance unsafe.Pointer, signalId uint, detail uint32) {}

// UNSUPPORTED : g_signal_emitv : parameter 'instance_and_params' is array parameter without length parameter

func Fn_g_signal_get_invocation_hint(instance unsafe.Pointer) {}

func Fn_g_signal_handler_block(instance unsafe.Pointer, handlerId uint64) {}

func Fn_g_signal_handler_disconnect(instance unsafe.Pointer, handlerId uint64) {}

func Fn_g_signal_handler_find(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) {
}

func Fn_g_signal_handler_is_connected(instance unsafe.Pointer, handlerId uint64) {}

func Fn_g_signal_handler_unblock(instance unsafe.Pointer, handlerId uint64) {}

func Fn_g_signal_handlers_block_matched(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) {
}

func Fn_g_signal_handlers_destroy(instance unsafe.Pointer) {}

func Fn_g_signal_handlers_disconnect_matched(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) {
}

func Fn_g_signal_handlers_unblock_matched(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) {
}

func Fn_g_signal_has_handler_pending(instance unsafe.Pointer, signalId uint, detail uint32, mayBeBlocked bool) {
}

func Fn_g_signal_list_ids(itype uint64) {}

func Fn_g_signal_lookup(name string, itype uint64) {}

func Fn_g_signal_name(signalId uint) {}

// UNSUPPORTED : g_signal_new : parameter 'accumulator' is callback

// UNSUPPORTED : g_signal_new_class_handler : parameter 'class_handler' is callback

// UNSUPPORTED : g_signal_new_valist : parameter 'accumulator' is callback

// UNSUPPORTED : g_signal_newv : parameter 'accumulator' is callback

func Fn_g_signal_override_class_closure(signalId uint, instanceType uint64, classClosure *Closure) {
}

// UNSUPPORTED : g_signal_override_class_handler : parameter 'class_handler' is callback

func Fn_g_signal_parse_name(detailedSignal string, itype uint64, forceDetailQuark bool) {}

func Fn_g_signal_query(signalId uint) {}

func Fn_g_signal_remove_emission_hook(signalId uint, hookId uint64) {}

// UNSUPPORTED : g_signal_set_va_marshaller : blacklisted

func Fn_g_signal_stop_emission(instance unsafe.Pointer, signalId uint, detail uint32) {}

func Fn_g_signal_stop_emission_by_name(instance unsafe.Pointer, detailedSignal string) {}

func Fn_g_signal_type_cclosure_new(itype uint64, structOffset uint) {}

func Fn_g_source_set_closure(source *glib.Source, closure *Closure) {}

func Fn_g_source_set_dummy_callback(source *glib.Source) {}

func Fn_g_strdup_value_contents(value *Value) {}

// UNSUPPORTED : g_type_add_class_cache_func : parameter 'cache_func' is callback

func Fn_g_type_add_instance_private(classType uint64, privateSize uint64) {}

// UNSUPPORTED : g_type_add_interface_check : parameter 'check_func' is callback

func Fn_g_type_add_interface_dynamic(instanceType uint64, interfaceType uint64, plugin *TypePlugin) {
}

func Fn_g_type_add_interface_static(instanceType uint64, interfaceType uint64, info *InterfaceInfo) {
}

func Fn_g_type_check_class_cast(gClass *TypeClass, isAType uint64) {}

func Fn_g_type_check_class_is_a(gClass *TypeClass, isAType uint64) {}

func Fn_g_type_check_instance(instance *TypeInstance) {}

func Fn_g_type_check_instance_cast(instance *TypeInstance, ifaceType uint64) {}

func Fn_g_type_check_instance_is_a(instance *TypeInstance, ifaceType uint64) {}

func Fn_g_type_check_instance_is_fundamentally_a(instance *TypeInstance, fundamentalType uint64) {}

func Fn_g_type_check_is_value_type(type_ uint64) {}

func Fn_g_type_check_value(value *Value) {}

func Fn_g_type_check_value_holds(value *Value, type_ uint64) {}

func Fn_g_type_children(type_ uint64) {}

func Fn_g_type_create_instance(type_ uint64) {}

func Fn_g_type_default_interface_peek(gType uint64) {}

func Fn_g_type_default_interface_ref(gType uint64) {}

func Fn_g_type_default_interface_unref(gIface unsafe.Pointer) {}

func Fn_g_type_depth(type_ uint64) {}

func Fn_g_type_free_instance(instance *TypeInstance) {}

func Fn_g_type_from_name(name string) {}

func Fn_g_type_fundamental(typeId uint64) {}

func Fn_g_type_fundamental_next() {}

func Fn_g_type_get_plugin(type_ uint64) {}

func Fn_g_type_get_qdata(type_ uint64, quark uint32) {}

func Fn_g_type_init() {}

func Fn_g_type_init_with_debug_flags(debugFlags int) {}

func Fn_g_type_interfaces(type_ uint64) {}

func Fn_g_type_is_a(type_ uint64, isAType uint64) {}

func Fn_g_type_name(type_ uint64) {}

func Fn_g_type_name_from_class(gClass *TypeClass) {}

func Fn_g_type_name_from_instance(instance *TypeInstance) {}

func Fn_g_type_next_base(leafType uint64, rootType uint64) {}

func Fn_g_type_parent(type_ uint64) {}

func Fn_g_type_qname(type_ uint64) {}

func Fn_g_type_query(type_ uint64) {}

func Fn_g_type_register_dynamic(parentType uint64, typeName string, plugin *TypePlugin, flags int) {
}

func Fn_g_type_register_fundamental(typeId uint64, typeName string, info *TypeInfo, finfo *TypeFundamentalInfo, flags int) {
}

func Fn_g_type_register_static(parentType uint64, typeName string, info *TypeInfo, flags int) {}

// UNSUPPORTED : g_type_register_static_simple : parameter 'class_init' is callback

// UNSUPPORTED : g_type_remove_class_cache_func : parameter 'cache_func' is callback

// UNSUPPORTED : g_type_remove_interface_check : parameter 'check_func' is callback

func Fn_g_type_set_qdata(type_ uint64, quark uint32, data unsafe.Pointer) {}

func Fn_g_type_test_flags(type_ uint64, flags uint) {}

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

// TypeModule is a representation of the C record GTypeModule.
type TypeModule struct {
	native unsafe.Pointer
}

// TypePlugin is a representation of the C interface GTypePlugin.
type TypePlugin struct {
	native unsafe.Pointer
}
