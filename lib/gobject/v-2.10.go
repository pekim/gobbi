// Code generated - DO NOT EDIT.
// +build gobject_2.10

package gobject

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/internal/c/gobject"
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

// BoxedCopy wraps the C function g_boxed_copy.
func BoxedCopy(boxedType uint64, srcBoxed unsafe.Pointer) unsafe.Pointer {
	sys_boxedType := boxedType
	sys_srcBoxed := srcBoxed
	retSys := gobject.Fn_g_boxed_copy(sys_boxedType, sys_srcBoxed)
	ret := retSys

	return ret
}

// BoxedFree wraps the C function g_boxed_free.
func BoxedFree(boxedType uint64, boxed unsafe.Pointer) {
	sys_boxedType := boxedType
	sys_boxed := boxed
	gobject.Fn_g_boxed_free(sys_boxedType, sys_boxed)
}

// UNSUPPORTED : g_boxed_type_register_static : parameter 'boxed_copy' is callback

// UNSUPPORTED : g_cclosure_new : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_object_swap : parameter 'callback_func' is callback

// UNSUPPORTED : g_cclosure_new_swap : parameter 'callback_func' is callback

// EnumCompleteTypeInfo wraps the C function g_enum_complete_type_info.
func EnumCompleteTypeInfo(gEnumType uint64, constValues *EnumValue) *TypeInfo {
	sys_gEnumType := gEnumType
	sys_constValues := constValues.ToC()
	gobject.Fn_g_enum_complete_type_info(sys_gEnumType, sys_constValues)
}

// EnumGetValue wraps the C function g_enum_get_value.
func EnumGetValue(enumClass *EnumClass, value int) *EnumValue {
	sys_enumClass := enumClass.ToC()
	sys_value := value
	retSys := gobject.Fn_g_enum_get_value(sys_enumClass, sys_value)
	ret := EnumValueNewFromC(retSys)

	return ret
}

// EnumGetValueByName wraps the C function g_enum_get_value_by_name.
func EnumGetValueByName(enumClass *EnumClass, name string) *EnumValue {
	sys_enumClass := enumClass.ToC()
	sys_name := name
	retSys := gobject.Fn_g_enum_get_value_by_name(sys_enumClass, sys_name)
	ret := EnumValueNewFromC(retSys)

	return ret
}

// EnumGetValueByNick wraps the C function g_enum_get_value_by_nick.
func EnumGetValueByNick(enumClass *EnumClass, nick string) *EnumValue {
	sys_enumClass := enumClass.ToC()
	sys_nick := nick
	retSys := gobject.Fn_g_enum_get_value_by_nick(sys_enumClass, sys_nick)
	ret := EnumValueNewFromC(retSys)

	return ret
}

// EnumRegisterStatic wraps the C function g_enum_register_static.
func EnumRegisterStatic(name string, constStaticValues *EnumValue) uint64 {
	sys_name := name
	sys_constStaticValues := constStaticValues.ToC()
	retSys := gobject.Fn_g_enum_register_static(sys_name, sys_constStaticValues)
	ret := retSys

	return ret
}

// FlagsCompleteTypeInfo wraps the C function g_flags_complete_type_info.
func FlagsCompleteTypeInfo(gFlagsType uint64, constValues *FlagsValue) *TypeInfo {
	sys_gFlagsType := gFlagsType
	sys_constValues := constValues.ToC()
	gobject.Fn_g_flags_complete_type_info(sys_gFlagsType, sys_constValues)
}

// FlagsGetFirstValue wraps the C function g_flags_get_first_value.
func FlagsGetFirstValue(flagsClass *FlagsClass, value uint) *FlagsValue {
	sys_flagsClass := flagsClass.ToC()
	sys_value := value
	retSys := gobject.Fn_g_flags_get_first_value(sys_flagsClass, sys_value)
	ret := FlagsValueNewFromC(retSys)

	return ret
}

// FlagsGetValueByName wraps the C function g_flags_get_value_by_name.
func FlagsGetValueByName(flagsClass *FlagsClass, name string) *FlagsValue {
	sys_flagsClass := flagsClass.ToC()
	sys_name := name
	retSys := gobject.Fn_g_flags_get_value_by_name(sys_flagsClass, sys_name)
	ret := FlagsValueNewFromC(retSys)

	return ret
}

// FlagsGetValueByNick wraps the C function g_flags_get_value_by_nick.
func FlagsGetValueByNick(flagsClass *FlagsClass, nick string) *FlagsValue {
	sys_flagsClass := flagsClass.ToC()
	sys_nick := nick
	retSys := gobject.Fn_g_flags_get_value_by_nick(sys_flagsClass, sys_nick)
	ret := FlagsValueNewFromC(retSys)

	return ret
}

// FlagsRegisterStatic wraps the C function g_flags_register_static.
func FlagsRegisterStatic(name string, constStaticValues *FlagsValue) uint64 {
	sys_name := name
	sys_constStaticValues := constStaticValues.ToC()
	retSys := gobject.Fn_g_flags_register_static(sys_name, sys_constStaticValues)
	ret := retSys

	return ret
}

// GtypeGetType wraps the C function g_gtype_get_type.
func GtypeGetType() uint64 {
	retSys := gobject.Fn_g_gtype_get_type()
	ret := retSys

	return ret
}

// ParamSpecBoolean_ wraps the C function g_param_spec_boolean.
func ParamSpecBoolean_(name string, nick string, blurb string, defaultValue bool, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_defaultValue := defaultValue
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_boolean(sys_name, sys_nick, sys_blurb, sys_defaultValue, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecBoxed_ wraps the C function g_param_spec_boxed.
func ParamSpecBoxed_(name string, nick string, blurb string, boxedType uint64, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_boxedType := boxedType
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_boxed(sys_name, sys_nick, sys_blurb, sys_boxedType, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecChar_ wraps the C function g_param_spec_char.
func ParamSpecChar_(name string, nick string, blurb string, minimum int8, maximum int8, defaultValue int8, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_char(sys_name, sys_nick, sys_blurb, sys_minimum, sys_maximum, sys_defaultValue, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecDouble_ wraps the C function g_param_spec_double.
func ParamSpecDouble_(name string, nick string, blurb string, minimum float64, maximum float64, defaultValue float64, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_double(sys_name, sys_nick, sys_blurb, sys_minimum, sys_maximum, sys_defaultValue, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecEnum_ wraps the C function g_param_spec_enum.
func ParamSpecEnum_(name string, nick string, blurb string, enumType uint64, defaultValue int, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_enumType := enumType
	sys_defaultValue := defaultValue
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_enum(sys_name, sys_nick, sys_blurb, sys_enumType, sys_defaultValue, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecFlags_ wraps the C function g_param_spec_flags.
func ParamSpecFlags_(name string, nick string, blurb string, flagsType uint64, defaultValue uint, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_flagsType := flagsType
	sys_defaultValue := defaultValue
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_flags(sys_name, sys_nick, sys_blurb, sys_flagsType, sys_defaultValue, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecFloat_ wraps the C function g_param_spec_float.
func ParamSpecFloat_(name string, nick string, blurb string, minimum float32, maximum float32, defaultValue float32, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_float(sys_name, sys_nick, sys_blurb, sys_minimum, sys_maximum, sys_defaultValue, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecGtype wraps the C function g_param_spec_gtype.
//
// since 2.10
func ParamSpecGtype(name string, nick string, blurb string, isAType uint64, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_isAType := isAType
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_gtype(sys_name, sys_nick, sys_blurb, sys_isAType, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecInt_ wraps the C function g_param_spec_int.
func ParamSpecInt_(name string, nick string, blurb string, minimum int, maximum int, defaultValue int, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_int(sys_name, sys_nick, sys_blurb, sys_minimum, sys_maximum, sys_defaultValue, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecInt64_ wraps the C function g_param_spec_int64.
func ParamSpecInt64_(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_int64(sys_name, sys_nick, sys_blurb, sys_minimum, sys_maximum, sys_defaultValue, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecLong_ wraps the C function g_param_spec_long.
func ParamSpecLong_(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_long(sys_name, sys_nick, sys_blurb, sys_minimum, sys_maximum, sys_defaultValue, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecObject_ wraps the C function g_param_spec_object.
func ParamSpecObject_(name string, nick string, blurb string, objectType uint64, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_objectType := objectType
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_object(sys_name, sys_nick, sys_blurb, sys_objectType, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecOverride_ wraps the C function g_param_spec_override.
//
// since 2.4
func ParamSpecOverride_(name string, overridden *ParamSpec) *ParamSpec {
	sys_name := name
	sys_overridden := overridden.ToC()
	retSys := gobject.Fn_g_param_spec_override(sys_name, sys_overridden)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecParam_ wraps the C function g_param_spec_param.
func ParamSpecParam_(name string, nick string, blurb string, paramType uint64, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_paramType := paramType
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_param(sys_name, sys_nick, sys_blurb, sys_paramType, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecPointer_ wraps the C function g_param_spec_pointer.
func ParamSpecPointer_(name string, nick string, blurb string, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_pointer(sys_name, sys_nick, sys_blurb, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecString_ wraps the C function g_param_spec_string.
func ParamSpecString_(name string, nick string, blurb string, defaultValue string, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_defaultValue := defaultValue
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_string(sys_name, sys_nick, sys_blurb, sys_defaultValue, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecUchar wraps the C function g_param_spec_uchar.
func ParamSpecUchar(name string, nick string, blurb string, minimum uint8, maximum uint8, defaultValue uint8, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_uchar(sys_name, sys_nick, sys_blurb, sys_minimum, sys_maximum, sys_defaultValue, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecUint wraps the C function g_param_spec_uint.
func ParamSpecUint(name string, nick string, blurb string, minimum uint, maximum uint, defaultValue uint, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_uint(sys_name, sys_nick, sys_blurb, sys_minimum, sys_maximum, sys_defaultValue, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecUint64 wraps the C function g_param_spec_uint64.
func ParamSpecUint64(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_uint64(sys_name, sys_nick, sys_blurb, sys_minimum, sys_maximum, sys_defaultValue, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecUlong wraps the C function g_param_spec_ulong.
func ParamSpecUlong(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_ulong(sys_name, sys_nick, sys_blurb, sys_minimum, sys_maximum, sys_defaultValue, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecUnichar_ wraps the C function g_param_spec_unichar.
func ParamSpecUnichar_(name string, nick string, blurb string, defaultValue rune, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_defaultValue := defaultValue
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_unichar(sys_name, sys_nick, sys_blurb, sys_defaultValue, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamSpecValueArray_ wraps the C function g_param_spec_value_array.
func ParamSpecValueArray_(name string, nick string, blurb string, elementSpec *ParamSpec, flags ParamFlags) *ParamSpec {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_elementSpec := elementSpec.ToC()
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_param_spec_value_array(sys_name, sys_nick, sys_blurb, sys_elementSpec, sys_flags)
	ret := ParamSpecNewFromC(retSys)

	return ret
}

// ParamTypeRegisterStatic wraps the C function g_param_type_register_static.
func ParamTypeRegisterStatic(name string, pspecInfo *ParamSpecTypeInfo) uint64 {
	sys_name := name
	sys_pspecInfo := pspecInfo.ToC()
	retSys := gobject.Fn_g_param_type_register_static(sys_name, sys_pspecInfo)
	ret := retSys

	return ret
}

// ParamValueConvert wraps the C function g_param_value_convert.
func ParamValueConvert(pspec *ParamSpec, srcValue *Value, destValue *Value, strictValidation bool) bool {
	sys_pspec := pspec.ToC()
	sys_srcValue := srcValue.ToC()
	sys_destValue := destValue.ToC()
	sys_strictValidation := strictValidation
	retSys := gobject.Fn_g_param_value_convert(sys_pspec, sys_srcValue, sys_destValue, sys_strictValidation)
	ret := retSys

	return ret
}

// ParamValueDefaults wraps the C function g_param_value_defaults.
func ParamValueDefaults(pspec *ParamSpec, value *Value) bool {
	sys_pspec := pspec.ToC()
	sys_value := value.ToC()
	retSys := gobject.Fn_g_param_value_defaults(sys_pspec, sys_value)
	ret := retSys

	return ret
}

// ParamValueSetDefault wraps the C function g_param_value_set_default.
func ParamValueSetDefault(pspec *ParamSpec, value *Value) {
	sys_pspec := pspec.ToC()
	sys_value := value.ToC()
	gobject.Fn_g_param_value_set_default(sys_pspec, sys_value)
}

// ParamValueValidate wraps the C function g_param_value_validate.
func ParamValueValidate(pspec *ParamSpec, value *Value) bool {
	sys_pspec := pspec.ToC()
	sys_value := value.ToC()
	retSys := gobject.Fn_g_param_value_validate(sys_pspec, sys_value)
	ret := retSys

	return ret
}

// ParamValuesCmp wraps the C function g_param_values_cmp.
func ParamValuesCmp(pspec *ParamSpec, value1 *Value, value2 *Value) int {
	sys_pspec := pspec.ToC()
	sys_value1 := value1.ToC()
	sys_value2 := value2.ToC()
	retSys := gobject.Fn_g_param_values_cmp(sys_pspec, sys_value1, sys_value2)
	ret := retSys

	return ret
}

// PointerTypeRegisterStatic wraps the C function g_pointer_type_register_static.
func PointerTypeRegisterStatic(name string) uint64 {
	sys_name := name
	retSys := gobject.Fn_g_pointer_type_register_static(sys_name)
	ret := retSys

	return ret
}

// SignalAccumulatorTrueHandled wraps the C function g_signal_accumulator_true_handled.
//
// since 2.4
func SignalAccumulatorTrueHandled(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value, dummy unsafe.Pointer) bool {
	sys_ihint := ihint.ToC()
	sys_returnAccu := returnAccu.ToC()
	sys_handlerReturn := handlerReturn.ToC()
	sys_dummy := dummy
	retSys := gobject.Fn_g_signal_accumulator_true_handled(sys_ihint, sys_returnAccu, sys_handlerReturn, sys_dummy)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_signal_add_emission_hook : parameter 'hook_func' is callback

// UNSUPPORTED : g_signal_chain_from_overridden : parameter 'instance_and_params' is array parameter without length parameter

// SignalConnectClosure wraps the C function g_signal_connect_closure.
func SignalConnectClosure(instance unsafe.Pointer, detailedSignal string, closure *Closure, after bool) uint64 {
	sys_instance := instance
	sys_detailedSignal := detailedSignal
	sys_closure := closure.ToC()
	sys_after := after
	retSys := gobject.Fn_g_signal_connect_closure(sys_instance, sys_detailedSignal, sys_closure, sys_after)
	ret := retSys

	return ret
}

// SignalConnectClosureById wraps the C function g_signal_connect_closure_by_id.
func SignalConnectClosureById(instance unsafe.Pointer, signalId uint, detail uint32, closure *Closure, after bool) uint64 {
	sys_instance := instance
	sys_signalId := signalId
	sys_detail := detail
	sys_closure := closure.ToC()
	sys_after := after
	retSys := gobject.Fn_g_signal_connect_closure_by_id(sys_instance, sys_signalId, sys_detail, sys_closure, sys_after)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_signal_connect_data : parameter 'c_handler' is callback

// UNSUPPORTED : g_signal_connect_object : parameter 'c_handler' is callback

// SignalEmit wraps the C function g_signal_emit.
func SignalEmit(instance unsafe.Pointer, signalId uint, detail uint32) {
	sys_instance := instance
	sys_signalId := signalId
	sys_detail := detail
	gobject.Fn_g_signal_emit(sys_instance, sys_signalId, sys_detail)
}

// SignalEmitByName wraps the C function g_signal_emit_by_name.
func SignalEmitByName(instance unsafe.Pointer, detailedSignal string) {
	sys_instance := instance
	sys_detailedSignal := detailedSignal
	gobject.Fn_g_signal_emit_by_name(sys_instance, sys_detailedSignal)
}

// SignalEmitValist wraps the C function g_signal_emit_valist.
func SignalEmitValist(instance unsafe.Pointer, signalId uint, detail uint32) {
	sys_instance := instance
	sys_signalId := signalId
	sys_detail := detail
	gobject.Fn_g_signal_emit_valist(sys_instance, sys_signalId, sys_detail)
}

// UNSUPPORTED : g_signal_emitv : parameter 'instance_and_params' is array parameter without length parameter

// SignalGetInvocationHint wraps the C function g_signal_get_invocation_hint.
func SignalGetInvocationHint(instance unsafe.Pointer) *SignalInvocationHint {
	sys_instance := instance
	retSys := gobject.Fn_g_signal_get_invocation_hint(sys_instance)
	ret := SignalInvocationHintNewFromC(retSys)

	return ret
}

// SignalHandlerBlock wraps the C function g_signal_handler_block.
func SignalHandlerBlock(instance unsafe.Pointer, handlerId uint64) {
	sys_instance := instance
	sys_handlerId := handlerId
	gobject.Fn_g_signal_handler_block(sys_instance, sys_handlerId)
}

// SignalHandlerDisconnect wraps the C function g_signal_handler_disconnect.
func SignalHandlerDisconnect(instance unsafe.Pointer, handlerId uint64) {
	sys_instance := instance
	sys_handlerId := handlerId
	gobject.Fn_g_signal_handler_disconnect(sys_instance, sys_handlerId)
}

// SignalHandlerFind wraps the C function g_signal_handler_find.
func SignalHandlerFind(instance unsafe.Pointer, mask SignalMatchType, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) uint64 {
	sys_instance := instance
	sys_mask := (int)(mask)
	sys_signalId := signalId
	sys_detail := detail
	sys_closure := closure.ToC()
	sys_func_ := func_
	sys_data := data
	retSys := gobject.Fn_g_signal_handler_find(sys_instance, sys_mask, sys_signalId, sys_detail, sys_closure, sys_func_, sys_data)
	ret := retSys

	return ret
}

// SignalHandlerIsConnected wraps the C function g_signal_handler_is_connected.
func SignalHandlerIsConnected(instance unsafe.Pointer, handlerId uint64) bool {
	sys_instance := instance
	sys_handlerId := handlerId
	retSys := gobject.Fn_g_signal_handler_is_connected(sys_instance, sys_handlerId)
	ret := retSys

	return ret
}

// SignalHandlerUnblock wraps the C function g_signal_handler_unblock.
func SignalHandlerUnblock(instance unsafe.Pointer, handlerId uint64) {
	sys_instance := instance
	sys_handlerId := handlerId
	gobject.Fn_g_signal_handler_unblock(sys_instance, sys_handlerId)
}

// SignalHandlersBlockMatched wraps the C function g_signal_handlers_block_matched.
func SignalHandlersBlockMatched(instance unsafe.Pointer, mask SignalMatchType, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) uint {
	sys_instance := instance
	sys_mask := (int)(mask)
	sys_signalId := signalId
	sys_detail := detail
	sys_closure := closure.ToC()
	sys_func_ := func_
	sys_data := data
	retSys := gobject.Fn_g_signal_handlers_block_matched(sys_instance, sys_mask, sys_signalId, sys_detail, sys_closure, sys_func_, sys_data)
	ret := retSys

	return ret
}

// SignalHandlersDestroy wraps the C function g_signal_handlers_destroy.
func SignalHandlersDestroy(instance unsafe.Pointer) {
	sys_instance := instance
	gobject.Fn_g_signal_handlers_destroy(sys_instance)
}

// SignalHandlersDisconnectMatched wraps the C function g_signal_handlers_disconnect_matched.
func SignalHandlersDisconnectMatched(instance unsafe.Pointer, mask SignalMatchType, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) uint {
	sys_instance := instance
	sys_mask := (int)(mask)
	sys_signalId := signalId
	sys_detail := detail
	sys_closure := closure.ToC()
	sys_func_ := func_
	sys_data := data
	retSys := gobject.Fn_g_signal_handlers_disconnect_matched(sys_instance, sys_mask, sys_signalId, sys_detail, sys_closure, sys_func_, sys_data)
	ret := retSys

	return ret
}

// SignalHandlersUnblockMatched wraps the C function g_signal_handlers_unblock_matched.
func SignalHandlersUnblockMatched(instance unsafe.Pointer, mask SignalMatchType, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) uint {
	sys_instance := instance
	sys_mask := (int)(mask)
	sys_signalId := signalId
	sys_detail := detail
	sys_closure := closure.ToC()
	sys_func_ := func_
	sys_data := data
	retSys := gobject.Fn_g_signal_handlers_unblock_matched(sys_instance, sys_mask, sys_signalId, sys_detail, sys_closure, sys_func_, sys_data)
	ret := retSys

	return ret
}

// SignalHasHandlerPending wraps the C function g_signal_has_handler_pending.
func SignalHasHandlerPending(instance unsafe.Pointer, signalId uint, detail uint32, mayBeBlocked bool) bool {
	sys_instance := instance
	sys_signalId := signalId
	sys_detail := detail
	sys_mayBeBlocked := mayBeBlocked
	retSys := gobject.Fn_g_signal_has_handler_pending(sys_instance, sys_signalId, sys_detail, sys_mayBeBlocked)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_signal_list_ids : has array return

// SignalLookup wraps the C function g_signal_lookup.
func SignalLookup(name string, itype uint64) uint {
	sys_name := name
	sys_itype := itype
	retSys := gobject.Fn_g_signal_lookup(sys_name, sys_itype)
	ret := retSys

	return ret
}

// SignalName wraps the C function g_signal_name.
func SignalName(signalId uint) string {
	sys_signalId := signalId
	retSys := gobject.Fn_g_signal_name(sys_signalId)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_signal_new : parameter 'accumulator' is callback

// UNSUPPORTED : g_signal_new_class_handler : parameter 'class_handler' is callback

// UNSUPPORTED : g_signal_new_valist : parameter 'accumulator' is callback

// UNSUPPORTED : g_signal_newv : parameter 'accumulator' is callback

// SignalOverrideClassClosure wraps the C function g_signal_override_class_closure.
func SignalOverrideClassClosure(signalId uint, instanceType uint64, classClosure *Closure) {
	sys_signalId := signalId
	sys_instanceType := instanceType
	sys_classClosure := classClosure.ToC()
	gobject.Fn_g_signal_override_class_closure(sys_signalId, sys_instanceType, sys_classClosure)
}

// UNSUPPORTED : g_signal_override_class_handler : parameter 'class_handler' is callback

// SignalParseName wraps the C function g_signal_parse_name.
func SignalParseName(detailedSignal string, itype uint64, forceDetailQuark bool) (bool, uint, uint32) {
	sys_detailedSignal := detailedSignal
	sys_itype := itype
	sys_forceDetailQuark := forceDetailQuark
	retSys := gobject.Fn_g_signal_parse_name(sys_detailedSignal, sys_itype, sys_forceDetailQuark)
	ret := retSys

	return ret
}

// SignalQuery_ wraps the C function g_signal_query.
func SignalQuery_(signalId uint) *SignalQuery {
	sys_signalId := signalId
	gobject.Fn_g_signal_query(sys_signalId)
}

// SignalRemoveEmissionHook wraps the C function g_signal_remove_emission_hook.
func SignalRemoveEmissionHook(signalId uint, hookId uint64) {
	sys_signalId := signalId
	sys_hookId := hookId
	gobject.Fn_g_signal_remove_emission_hook(sys_signalId, sys_hookId)
}

// UNSUPPORTED : g_signal_set_va_marshaller : blacklisted

// SignalStopEmission wraps the C function g_signal_stop_emission.
func SignalStopEmission(instance unsafe.Pointer, signalId uint, detail uint32) {
	sys_instance := instance
	sys_signalId := signalId
	sys_detail := detail
	gobject.Fn_g_signal_stop_emission(sys_instance, sys_signalId, sys_detail)
}

// SignalStopEmissionByName wraps the C function g_signal_stop_emission_by_name.
func SignalStopEmissionByName(instance unsafe.Pointer, detailedSignal string) {
	sys_instance := instance
	sys_detailedSignal := detailedSignal
	gobject.Fn_g_signal_stop_emission_by_name(sys_instance, sys_detailedSignal)
}

// SignalTypeCclosureNew wraps the C function g_signal_type_cclosure_new.
func SignalTypeCclosureNew(itype uint64, structOffset uint) *Closure {
	sys_itype := itype
	sys_structOffset := structOffset
	retSys := gobject.Fn_g_signal_type_cclosure_new(sys_itype, sys_structOffset)
	ret := ClosureNewFromC(retSys)

	return ret
}

// SourceSetClosure wraps the C function g_source_set_closure.
func SourceSetClosure(source *glib.Source, closure *Closure) {
	sys_source := source.ToC()
	sys_closure := closure.ToC()
	gobject.Fn_g_source_set_closure(sys_source, sys_closure)
}

// SourceSetDummyCallback wraps the C function g_source_set_dummy_callback.
func SourceSetDummyCallback(source *glib.Source) {
	sys_source := source.ToC()
	gobject.Fn_g_source_set_dummy_callback(sys_source)
}

// StrdupValueContents wraps the C function g_strdup_value_contents.
func StrdupValueContents(value *Value) string {
	sys_value := value.ToC()
	retSys := gobject.Fn_g_strdup_value_contents(sys_value)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_type_add_class_cache_func : parameter 'cache_func' is callback

// TypeAddInstancePrivate wraps the C function g_type_add_instance_private.
func TypeAddInstancePrivate(classType uint64, privateSize uint64) int {
	sys_classType := classType
	sys_privateSize := privateSize
	retSys := gobject.Fn_g_type_add_instance_private(sys_classType, sys_privateSize)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_type_add_interface_check : parameter 'check_func' is callback

// TypeAddInterfaceDynamic wraps the C function g_type_add_interface_dynamic.
func TypeAddInterfaceDynamic(instanceType uint64, interfaceType uint64, plugin *TypePlugin) {
	sys_instanceType := instanceType
	sys_interfaceType := interfaceType
	sys_plugin := plugin.ToC()
	gobject.Fn_g_type_add_interface_dynamic(sys_instanceType, sys_interfaceType, sys_plugin)
}

// TypeAddInterfaceStatic wraps the C function g_type_add_interface_static.
func TypeAddInterfaceStatic(instanceType uint64, interfaceType uint64, info *InterfaceInfo) {
	sys_instanceType := instanceType
	sys_interfaceType := interfaceType
	sys_info := info.ToC()
	gobject.Fn_g_type_add_interface_static(sys_instanceType, sys_interfaceType, sys_info)
}

// TypeCheckClassCast wraps the C function g_type_check_class_cast.
func TypeCheckClassCast(gClass *TypeClass, isAType uint64) *TypeClass {
	sys_gClass := gClass.ToC()
	sys_isAType := isAType
	retSys := gobject.Fn_g_type_check_class_cast(sys_gClass, sys_isAType)
	ret := TypeClassNewFromC(retSys)

	return ret
}

// TypeCheckClassIsA wraps the C function g_type_check_class_is_a.
func TypeCheckClassIsA(gClass *TypeClass, isAType uint64) bool {
	sys_gClass := gClass.ToC()
	sys_isAType := isAType
	retSys := gobject.Fn_g_type_check_class_is_a(sys_gClass, sys_isAType)
	ret := retSys

	return ret
}

// TypeCheckInstance wraps the C function g_type_check_instance.
func TypeCheckInstance(instance *TypeInstance) bool {
	sys_instance := instance.ToC()
	retSys := gobject.Fn_g_type_check_instance(sys_instance)
	ret := retSys

	return ret
}

// TypeCheckInstanceCast wraps the C function g_type_check_instance_cast.
func TypeCheckInstanceCast(instance *TypeInstance, ifaceType uint64) *TypeInstance {
	sys_instance := instance.ToC()
	sys_ifaceType := ifaceType
	retSys := gobject.Fn_g_type_check_instance_cast(sys_instance, sys_ifaceType)
	ret := TypeInstanceNewFromC(retSys)

	return ret
}

// TypeCheckInstanceIsA wraps the C function g_type_check_instance_is_a.
func TypeCheckInstanceIsA(instance *TypeInstance, ifaceType uint64) bool {
	sys_instance := instance.ToC()
	sys_ifaceType := ifaceType
	retSys := gobject.Fn_g_type_check_instance_is_a(sys_instance, sys_ifaceType)
	ret := retSys

	return ret
}

// TypeCheckInstanceIsFundamentallyA wraps the C function g_type_check_instance_is_fundamentally_a.
func TypeCheckInstanceIsFundamentallyA(instance *TypeInstance, fundamentalType uint64) bool {
	sys_instance := instance.ToC()
	sys_fundamentalType := fundamentalType
	retSys := gobject.Fn_g_type_check_instance_is_fundamentally_a(sys_instance, sys_fundamentalType)
	ret := retSys

	return ret
}

// TypeCheckIsValueType wraps the C function g_type_check_is_value_type.
func TypeCheckIsValueType(type_ uint64) bool {
	sys_type_ := type_
	retSys := gobject.Fn_g_type_check_is_value_type(sys_type_)
	ret := retSys

	return ret
}

// TypeCheckValue wraps the C function g_type_check_value.
func TypeCheckValue(value *Value) bool {
	sys_value := value.ToC()
	retSys := gobject.Fn_g_type_check_value(sys_value)
	ret := retSys

	return ret
}

// TypeCheckValueHolds wraps the C function g_type_check_value_holds.
func TypeCheckValueHolds(value *Value, type_ uint64) bool {
	sys_value := value.ToC()
	sys_type_ := type_
	retSys := gobject.Fn_g_type_check_value_holds(sys_value, sys_type_)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_type_children : has array return

// TypeCreateInstance wraps the C function g_type_create_instance.
func TypeCreateInstance(type_ uint64) *TypeInstance {
	sys_type_ := type_
	retSys := gobject.Fn_g_type_create_instance(sys_type_)
	ret := TypeInstanceNewFromC(retSys)

	return ret
}

// TypeDefaultInterfacePeek wraps the C function g_type_default_interface_peek.
//
// since 2.4
func TypeDefaultInterfacePeek(gType uint64) *TypeInterface {
	sys_gType := gType
	retSys := gobject.Fn_g_type_default_interface_peek(sys_gType)
	ret := TypeInterfaceNewFromC(retSys)

	return ret
}

// TypeDefaultInterfaceRef wraps the C function g_type_default_interface_ref.
//
// since 2.4
func TypeDefaultInterfaceRef(gType uint64) *TypeInterface {
	sys_gType := gType
	retSys := gobject.Fn_g_type_default_interface_ref(sys_gType)
	ret := TypeInterfaceNewFromC(retSys)

	return ret
}

// TypeDefaultInterfaceUnref wraps the C function g_type_default_interface_unref.
//
// since 2.4
func TypeDefaultInterfaceUnref(gIface unsafe.Pointer) {
	sys_gIface := gIface
	gobject.Fn_g_type_default_interface_unref(sys_gIface)
}

// TypeDepth wraps the C function g_type_depth.
func TypeDepth(type_ uint64) uint {
	sys_type_ := type_
	retSys := gobject.Fn_g_type_depth(sys_type_)
	ret := retSys

	return ret
}

// TypeFreeInstance wraps the C function g_type_free_instance.
func TypeFreeInstance(instance *TypeInstance) {
	sys_instance := instance.ToC()
	gobject.Fn_g_type_free_instance(sys_instance)
}

// TypeFromName wraps the C function g_type_from_name.
func TypeFromName(name string) uint64 {
	sys_name := name
	retSys := gobject.Fn_g_type_from_name(sys_name)
	ret := retSys

	return ret
}

// TypeFundamental wraps the C function g_type_fundamental.
func TypeFundamental(typeId uint64) uint64 {
	sys_typeId := typeId
	retSys := gobject.Fn_g_type_fundamental(sys_typeId)
	ret := retSys

	return ret
}

// TypeFundamentalNext wraps the C function g_type_fundamental_next.
func TypeFundamentalNext() uint64 {
	retSys := gobject.Fn_g_type_fundamental_next()
	ret := retSys

	return ret
}

// TypeGetPlugin wraps the C function g_type_get_plugin.
func TypeGetPlugin(type_ uint64) *TypePlugin {
	sys_type_ := type_
	retSys := gobject.Fn_g_type_get_plugin(sys_type_)
	ret := TypePluginNewFromC(retSys)

	return ret
}

// TypeGetQdata wraps the C function g_type_get_qdata.
func TypeGetQdata(type_ uint64, quark uint32) unsafe.Pointer {
	sys_type_ := type_
	sys_quark := quark
	retSys := gobject.Fn_g_type_get_qdata(sys_type_, sys_quark)
	ret := retSys

	return ret
}

// TypeInit wraps the C function g_type_init.
func TypeInit() {
	gobject.Fn_g_type_init()
}

// TypeInitWithDebugFlags wraps the C function g_type_init_with_debug_flags.
func TypeInitWithDebugFlags(debugFlags TypeDebugFlags) {
	sys_debugFlags := (int)(debugFlags)
	gobject.Fn_g_type_init_with_debug_flags(sys_debugFlags)
}

// UNSUPPORTED : g_type_interface_prerequisites : has array return

// UNSUPPORTED : g_type_interfaces : has array return

// TypeIsA wraps the C function g_type_is_a.
func TypeIsA(type_ uint64, isAType uint64) bool {
	sys_type_ := type_
	sys_isAType := isAType
	retSys := gobject.Fn_g_type_is_a(sys_type_, sys_isAType)
	ret := retSys

	return ret
}

// TypeName wraps the C function g_type_name.
func TypeName(type_ uint64) string {
	sys_type_ := type_
	retSys := gobject.Fn_g_type_name(sys_type_)
	ret := retSys

	return ret
}

// TypeNameFromClass wraps the C function g_type_name_from_class.
func TypeNameFromClass(gClass *TypeClass) string {
	sys_gClass := gClass.ToC()
	retSys := gobject.Fn_g_type_name_from_class(sys_gClass)
	ret := retSys

	return ret
}

// TypeNameFromInstance wraps the C function g_type_name_from_instance.
func TypeNameFromInstance(instance *TypeInstance) string {
	sys_instance := instance.ToC()
	retSys := gobject.Fn_g_type_name_from_instance(sys_instance)
	ret := retSys

	return ret
}

// TypeNextBase wraps the C function g_type_next_base.
func TypeNextBase(leafType uint64, rootType uint64) uint64 {
	sys_leafType := leafType
	sys_rootType := rootType
	retSys := gobject.Fn_g_type_next_base(sys_leafType, sys_rootType)
	ret := retSys

	return ret
}

// TypeParent wraps the C function g_type_parent.
func TypeParent(type_ uint64) uint64 {
	sys_type_ := type_
	retSys := gobject.Fn_g_type_parent(sys_type_)
	ret := retSys

	return ret
}

// TypeQname wraps the C function g_type_qname.
func TypeQname(type_ uint64) uint32 {
	sys_type_ := type_
	retSys := gobject.Fn_g_type_qname(sys_type_)
	ret := retSys

	return ret
}

// TypeQuery_ wraps the C function g_type_query.
func TypeQuery_(type_ uint64) *TypeQuery {
	sys_type_ := type_
	gobject.Fn_g_type_query(sys_type_)
}

// TypeRegisterDynamic wraps the C function g_type_register_dynamic.
func TypeRegisterDynamic(parentType uint64, typeName string, plugin *TypePlugin, flags TypeFlags) uint64 {
	sys_parentType := parentType
	sys_typeName := typeName
	sys_plugin := plugin.ToC()
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_type_register_dynamic(sys_parentType, sys_typeName, sys_plugin, sys_flags)
	ret := retSys

	return ret
}

// TypeRegisterFundamental wraps the C function g_type_register_fundamental.
func TypeRegisterFundamental(typeId uint64, typeName string, info *TypeInfo, finfo *TypeFundamentalInfo, flags TypeFlags) uint64 {
	sys_typeId := typeId
	sys_typeName := typeName
	sys_info := info.ToC()
	sys_finfo := finfo.ToC()
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_type_register_fundamental(sys_typeId, sys_typeName, sys_info, sys_finfo, sys_flags)
	ret := retSys

	return ret
}

// TypeRegisterStatic wraps the C function g_type_register_static.
func TypeRegisterStatic(parentType uint64, typeName string, info *TypeInfo, flags TypeFlags) uint64 {
	sys_parentType := parentType
	sys_typeName := typeName
	sys_info := info.ToC()
	sys_flags := (int)(flags)
	retSys := gobject.Fn_g_type_register_static(sys_parentType, sys_typeName, sys_info, sys_flags)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_type_register_static_simple : parameter 'class_init' is callback

// UNSUPPORTED : g_type_remove_class_cache_func : parameter 'cache_func' is callback

// UNSUPPORTED : g_type_remove_interface_check : parameter 'check_func' is callback

// TypeSetQdata wraps the C function g_type_set_qdata.
func TypeSetQdata(type_ uint64, quark uint32, data unsafe.Pointer) {
	sys_type_ := type_
	sys_quark := quark
	sys_data := data
	gobject.Fn_g_type_set_qdata(sys_type_, sys_quark, sys_data)
}

// TypeTestFlags wraps the C function g_type_test_flags.
func TypeTestFlags(type_ uint64, flags uint) bool {
	sys_type_ := type_
	sys_flags := flags
	retSys := gobject.Fn_g_type_test_flags(sys_type_, sys_flags)
	ret := retSys

	return ret
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

// CClosureNewFromC creates a new CClosure from a pointer to the C GCClosure that represents the CClosure.
func CClosureNewFromC(native unsafe.Pointer) *CClosure {
	return &CClosure{native: native}
}

// Closure is a representation of the C record GClosure.
type Closure struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GClosure that represents the Closure.
func (recv *Closure) ToC() unsafe.Pointer {
	return recv.native
}

// ClosureNewFromC creates a new Closure from a pointer to the C GClosure that represents the Closure.
func ClosureNewFromC(native unsafe.Pointer) *Closure {
	return &Closure{native: native}
}

// ClosureNotifyData is a representation of the C record GClosureNotifyData.
type ClosureNotifyData struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GClosureNotifyData that represents the ClosureNotifyData.
func (recv *ClosureNotifyData) ToC() unsafe.Pointer {
	return recv.native
}

// ClosureNotifyDataNewFromC creates a new ClosureNotifyData from a pointer to the C GClosureNotifyData that represents the ClosureNotifyData.
func ClosureNotifyDataNewFromC(native unsafe.Pointer) *ClosureNotifyData {
	return &ClosureNotifyData{native: native}
}

// EnumClass is a representation of the C record GEnumClass.
type EnumClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GEnumClass that represents the EnumClass.
func (recv *EnumClass) ToC() unsafe.Pointer {
	return recv.native
}

// EnumClassNewFromC creates a new EnumClass from a pointer to the C GEnumClass that represents the EnumClass.
func EnumClassNewFromC(native unsafe.Pointer) *EnumClass {
	return &EnumClass{native: native}
}

// EnumValue is a representation of the C record GEnumValue.
type EnumValue struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GEnumValue that represents the EnumValue.
func (recv *EnumValue) ToC() unsafe.Pointer {
	return recv.native
}

// EnumValueNewFromC creates a new EnumValue from a pointer to the C GEnumValue that represents the EnumValue.
func EnumValueNewFromC(native unsafe.Pointer) *EnumValue {
	return &EnumValue{native: native}
}

// FlagsClass is a representation of the C record GFlagsClass.
type FlagsClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFlagsClass that represents the FlagsClass.
func (recv *FlagsClass) ToC() unsafe.Pointer {
	return recv.native
}

// FlagsClassNewFromC creates a new FlagsClass from a pointer to the C GFlagsClass that represents the FlagsClass.
func FlagsClassNewFromC(native unsafe.Pointer) *FlagsClass {
	return &FlagsClass{native: native}
}

// FlagsValue is a representation of the C record GFlagsValue.
type FlagsValue struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFlagsValue that represents the FlagsValue.
func (recv *FlagsValue) ToC() unsafe.Pointer {
	return recv.native
}

// FlagsValueNewFromC creates a new FlagsValue from a pointer to the C GFlagsValue that represents the FlagsValue.
func FlagsValueNewFromC(native unsafe.Pointer) *FlagsValue {
	return &FlagsValue{native: native}
}

// InitiallyUnownedClass is a representation of the C record GInitiallyUnownedClass.
type InitiallyUnownedClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInitiallyUnownedClass that represents the InitiallyUnownedClass.
func (recv *InitiallyUnownedClass) ToC() unsafe.Pointer {
	return recv.native
}

// InitiallyUnownedClassNewFromC creates a new InitiallyUnownedClass from a pointer to the C GInitiallyUnownedClass that represents the InitiallyUnownedClass.
func InitiallyUnownedClassNewFromC(native unsafe.Pointer) *InitiallyUnownedClass {
	return &InitiallyUnownedClass{native: native}
}

// InterfaceInfo is a representation of the C record GInterfaceInfo.
type InterfaceInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInterfaceInfo that represents the InterfaceInfo.
func (recv *InterfaceInfo) ToC() unsafe.Pointer {
	return recv.native
}

// InterfaceInfoNewFromC creates a new InterfaceInfo from a pointer to the C GInterfaceInfo that represents the InterfaceInfo.
func InterfaceInfoNewFromC(native unsafe.Pointer) *InterfaceInfo {
	return &InterfaceInfo{native: native}
}

// ObjectClass is a representation of the C record GObjectClass.
type ObjectClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GObjectClass that represents the ObjectClass.
func (recv *ObjectClass) ToC() unsafe.Pointer {
	return recv.native
}

// ObjectClassNewFromC creates a new ObjectClass from a pointer to the C GObjectClass that represents the ObjectClass.
func ObjectClassNewFromC(native unsafe.Pointer) *ObjectClass {
	return &ObjectClass{native: native}
}

// ObjectConstructParam is a representation of the C record GObjectConstructParam.
type ObjectConstructParam struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GObjectConstructParam that represents the ObjectConstructParam.
func (recv *ObjectConstructParam) ToC() unsafe.Pointer {
	return recv.native
}

// ObjectConstructParamNewFromC creates a new ObjectConstructParam from a pointer to the C GObjectConstructParam that represents the ObjectConstructParam.
func ObjectConstructParamNewFromC(native unsafe.Pointer) *ObjectConstructParam {
	return &ObjectConstructParam{native: native}
}

// ParamSpecClass is a representation of the C record GParamSpecClass.
type ParamSpecClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecClass that represents the ParamSpecClass.
func (recv *ParamSpecClass) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecClassNewFromC creates a new ParamSpecClass from a pointer to the C GParamSpecClass that represents the ParamSpecClass.
func ParamSpecClassNewFromC(native unsafe.Pointer) *ParamSpecClass {
	return &ParamSpecClass{native: native}
}

// ParamSpecPool is a representation of the C record GParamSpecPool.
type ParamSpecPool struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecPool that represents the ParamSpecPool.
func (recv *ParamSpecPool) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecPoolNewFromC creates a new ParamSpecPool from a pointer to the C GParamSpecPool that represents the ParamSpecPool.
func ParamSpecPoolNewFromC(native unsafe.Pointer) *ParamSpecPool {
	return &ParamSpecPool{native: native}
}

// ParamSpecTypeInfo is a representation of the C record GParamSpecTypeInfo.
type ParamSpecTypeInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecTypeInfo that represents the ParamSpecTypeInfo.
func (recv *ParamSpecTypeInfo) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecTypeInfoNewFromC creates a new ParamSpecTypeInfo from a pointer to the C GParamSpecTypeInfo that represents the ParamSpecTypeInfo.
func ParamSpecTypeInfoNewFromC(native unsafe.Pointer) *ParamSpecTypeInfo {
	return &ParamSpecTypeInfo{native: native}
}

// Parameter is a representation of the C record GParameter.
type Parameter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParameter that represents the Parameter.
func (recv *Parameter) ToC() unsafe.Pointer {
	return recv.native
}

// ParameterNewFromC creates a new Parameter from a pointer to the C GParameter that represents the Parameter.
func ParameterNewFromC(native unsafe.Pointer) *Parameter {
	return &Parameter{native: native}
}

// SignalInvocationHint is a representation of the C record GSignalInvocationHint.
type SignalInvocationHint struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSignalInvocationHint that represents the SignalInvocationHint.
func (recv *SignalInvocationHint) ToC() unsafe.Pointer {
	return recv.native
}

// SignalInvocationHintNewFromC creates a new SignalInvocationHint from a pointer to the C GSignalInvocationHint that represents the SignalInvocationHint.
func SignalInvocationHintNewFromC(native unsafe.Pointer) *SignalInvocationHint {
	return &SignalInvocationHint{native: native}
}

// SignalQuery is a representation of the C record GSignalQuery.
type SignalQuery struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSignalQuery that represents the SignalQuery.
func (recv *SignalQuery) ToC() unsafe.Pointer {
	return recv.native
}

// SignalQueryNewFromC creates a new SignalQuery from a pointer to the C GSignalQuery that represents the SignalQuery.
func SignalQueryNewFromC(native unsafe.Pointer) *SignalQuery {
	return &SignalQuery{native: native}
}

// TypeClass is a representation of the C record GTypeClass.
type TypeClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeClass that represents the TypeClass.
func (recv *TypeClass) ToC() unsafe.Pointer {
	return recv.native
}

// TypeClassNewFromC creates a new TypeClass from a pointer to the C GTypeClass that represents the TypeClass.
func TypeClassNewFromC(native unsafe.Pointer) *TypeClass {
	return &TypeClass{native: native}
}

// TypeFundamentalInfo is a representation of the C record GTypeFundamentalInfo.
type TypeFundamentalInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeFundamentalInfo that represents the TypeFundamentalInfo.
func (recv *TypeFundamentalInfo) ToC() unsafe.Pointer {
	return recv.native
}

// TypeFundamentalInfoNewFromC creates a new TypeFundamentalInfo from a pointer to the C GTypeFundamentalInfo that represents the TypeFundamentalInfo.
func TypeFundamentalInfoNewFromC(native unsafe.Pointer) *TypeFundamentalInfo {
	return &TypeFundamentalInfo{native: native}
}

// TypeInfo is a representation of the C record GTypeInfo.
type TypeInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeInfo that represents the TypeInfo.
func (recv *TypeInfo) ToC() unsafe.Pointer {
	return recv.native
}

// TypeInfoNewFromC creates a new TypeInfo from a pointer to the C GTypeInfo that represents the TypeInfo.
func TypeInfoNewFromC(native unsafe.Pointer) *TypeInfo {
	return &TypeInfo{native: native}
}

// TypeInstance is a representation of the C record GTypeInstance.
type TypeInstance struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeInstance that represents the TypeInstance.
func (recv *TypeInstance) ToC() unsafe.Pointer {
	return recv.native
}

// TypeInstanceNewFromC creates a new TypeInstance from a pointer to the C GTypeInstance that represents the TypeInstance.
func TypeInstanceNewFromC(native unsafe.Pointer) *TypeInstance {
	return &TypeInstance{native: native}
}

// TypeInterface is a representation of the C record GTypeInterface.
type TypeInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeInterface that represents the TypeInterface.
func (recv *TypeInterface) ToC() unsafe.Pointer {
	return recv.native
}

// TypeInterfaceNewFromC creates a new TypeInterface from a pointer to the C GTypeInterface that represents the TypeInterface.
func TypeInterfaceNewFromC(native unsafe.Pointer) *TypeInterface {
	return &TypeInterface{native: native}
}

// TypeModuleClass is a representation of the C record GTypeModuleClass.
type TypeModuleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeModuleClass that represents the TypeModuleClass.
func (recv *TypeModuleClass) ToC() unsafe.Pointer {
	return recv.native
}

// TypeModuleClassNewFromC creates a new TypeModuleClass from a pointer to the C GTypeModuleClass that represents the TypeModuleClass.
func TypeModuleClassNewFromC(native unsafe.Pointer) *TypeModuleClass {
	return &TypeModuleClass{native: native}
}

// TypePluginClass is a representation of the C record GTypePluginClass.
type TypePluginClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypePluginClass that represents the TypePluginClass.
func (recv *TypePluginClass) ToC() unsafe.Pointer {
	return recv.native
}

// TypePluginClassNewFromC creates a new TypePluginClass from a pointer to the C GTypePluginClass that represents the TypePluginClass.
func TypePluginClassNewFromC(native unsafe.Pointer) *TypePluginClass {
	return &TypePluginClass{native: native}
}

// TypeQuery is a representation of the C record GTypeQuery.
type TypeQuery struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeQuery that represents the TypeQuery.
func (recv *TypeQuery) ToC() unsafe.Pointer {
	return recv.native
}

// TypeQueryNewFromC creates a new TypeQuery from a pointer to the C GTypeQuery that represents the TypeQuery.
func TypeQueryNewFromC(native unsafe.Pointer) *TypeQuery {
	return &TypeQuery{native: native}
}

// TypeValueTable is a representation of the C record GTypeValueTable.
type TypeValueTable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeValueTable that represents the TypeValueTable.
func (recv *TypeValueTable) ToC() unsafe.Pointer {
	return recv.native
}

// TypeValueTableNewFromC creates a new TypeValueTable from a pointer to the C GTypeValueTable that represents the TypeValueTable.
func TypeValueTableNewFromC(native unsafe.Pointer) *TypeValueTable {
	return &TypeValueTable{native: native}
}

// Value is a representation of the C record GValue.
type Value struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GValue that represents the Value.
func (recv *Value) ToC() unsafe.Pointer {
	return recv.native
}

// ValueNewFromC creates a new Value from a pointer to the C GValue that represents the Value.
func ValueNewFromC(native unsafe.Pointer) *Value {
	return &Value{native: native}
}

// ValueArray is a representation of the C record GValueArray.
type ValueArray struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GValueArray that represents the ValueArray.
func (recv *ValueArray) ToC() unsafe.Pointer {
	return recv.native
}

// ValueArrayNewFromC creates a new ValueArray from a pointer to the C GValueArray that represents the ValueArray.
func ValueArrayNewFromC(native unsafe.Pointer) *ValueArray {
	return &ValueArray{native: native}
}

// WeakRef is a representation of the C record GWeakRef.
type WeakRef struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GWeakRef that represents the WeakRef.
func (recv *WeakRef) ToC() unsafe.Pointer {
	return recv.native
}

// WeakRefNewFromC creates a new WeakRef from a pointer to the C GWeakRef that represents the WeakRef.
func WeakRefNewFromC(native unsafe.Pointer) *WeakRef {
	return &WeakRef{native: native}
}

// InitiallyUnowned is a representation of the C record GInitiallyUnowned.
type InitiallyUnowned struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInitiallyUnowned that represents the InitiallyUnowned.
func (recv *InitiallyUnowned) ToC() unsafe.Pointer {
	return recv.native
}

// InitiallyUnownedNewFromC creates a new InitiallyUnowned from a pointer to the C GInitiallyUnowned that represents the InitiallyUnowned.
func InitiallyUnownedNewFromC(native unsafe.Pointer) *InitiallyUnowned {
	return &InitiallyUnowned{native: native}
}

// Object is a representation of the C record GObject.
type Object struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GObject that represents the Object.
func (recv *Object) ToC() unsafe.Pointer {
	return recv.native
}

// ObjectNewFromC creates a new Object from a pointer to the C GObject that represents the Object.
func ObjectNewFromC(native unsafe.Pointer) *Object {
	return &Object{native: native}
}

// ParamSpec is a representation of the C record GParamSpec.
type ParamSpec struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpec that represents the ParamSpec.
func (recv *ParamSpec) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecNewFromC creates a new ParamSpec from a pointer to the C GParamSpec that represents the ParamSpec.
func ParamSpecNewFromC(native unsafe.Pointer) *ParamSpec {
	return &ParamSpec{native: native}
}

// ParamSpecBoolean is a representation of the C record GParamSpecBoolean.
type ParamSpecBoolean struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecBoolean that represents the ParamSpecBoolean.
func (recv *ParamSpecBoolean) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecBooleanNewFromC creates a new ParamSpecBoolean from a pointer to the C GParamSpecBoolean that represents the ParamSpecBoolean.
func ParamSpecBooleanNewFromC(native unsafe.Pointer) *ParamSpecBoolean {
	return &ParamSpecBoolean{native: native}
}

// ParamSpecBoxed is a representation of the C record GParamSpecBoxed.
type ParamSpecBoxed struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecBoxed that represents the ParamSpecBoxed.
func (recv *ParamSpecBoxed) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecBoxedNewFromC creates a new ParamSpecBoxed from a pointer to the C GParamSpecBoxed that represents the ParamSpecBoxed.
func ParamSpecBoxedNewFromC(native unsafe.Pointer) *ParamSpecBoxed {
	return &ParamSpecBoxed{native: native}
}

// ParamSpecChar is a representation of the C record GParamSpecChar.
type ParamSpecChar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecChar that represents the ParamSpecChar.
func (recv *ParamSpecChar) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecCharNewFromC creates a new ParamSpecChar from a pointer to the C GParamSpecChar that represents the ParamSpecChar.
func ParamSpecCharNewFromC(native unsafe.Pointer) *ParamSpecChar {
	return &ParamSpecChar{native: native}
}

// ParamSpecDouble is a representation of the C record GParamSpecDouble.
type ParamSpecDouble struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecDouble that represents the ParamSpecDouble.
func (recv *ParamSpecDouble) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecDoubleNewFromC creates a new ParamSpecDouble from a pointer to the C GParamSpecDouble that represents the ParamSpecDouble.
func ParamSpecDoubleNewFromC(native unsafe.Pointer) *ParamSpecDouble {
	return &ParamSpecDouble{native: native}
}

// ParamSpecEnum is a representation of the C record GParamSpecEnum.
type ParamSpecEnum struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecEnum that represents the ParamSpecEnum.
func (recv *ParamSpecEnum) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecEnumNewFromC creates a new ParamSpecEnum from a pointer to the C GParamSpecEnum that represents the ParamSpecEnum.
func ParamSpecEnumNewFromC(native unsafe.Pointer) *ParamSpecEnum {
	return &ParamSpecEnum{native: native}
}

// ParamSpecFlags is a representation of the C record GParamSpecFlags.
type ParamSpecFlags struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecFlags that represents the ParamSpecFlags.
func (recv *ParamSpecFlags) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecFlagsNewFromC creates a new ParamSpecFlags from a pointer to the C GParamSpecFlags that represents the ParamSpecFlags.
func ParamSpecFlagsNewFromC(native unsafe.Pointer) *ParamSpecFlags {
	return &ParamSpecFlags{native: native}
}

// ParamSpecFloat is a representation of the C record GParamSpecFloat.
type ParamSpecFloat struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecFloat that represents the ParamSpecFloat.
func (recv *ParamSpecFloat) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecFloatNewFromC creates a new ParamSpecFloat from a pointer to the C GParamSpecFloat that represents the ParamSpecFloat.
func ParamSpecFloatNewFromC(native unsafe.Pointer) *ParamSpecFloat {
	return &ParamSpecFloat{native: native}
}

// ParamSpecGType is a representation of the C record GParamSpecGType.
//
// since 2.10
type ParamSpecGType struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecGType that represents the ParamSpecGType.
func (recv *ParamSpecGType) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecGTypeNewFromC creates a new ParamSpecGType from a pointer to the C GParamSpecGType that represents the ParamSpecGType.
func ParamSpecGTypeNewFromC(native unsafe.Pointer) *ParamSpecGType {
	return &ParamSpecGType{native: native}
}

// ParamSpecInt is a representation of the C record GParamSpecInt.
type ParamSpecInt struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecInt that represents the ParamSpecInt.
func (recv *ParamSpecInt) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecIntNewFromC creates a new ParamSpecInt from a pointer to the C GParamSpecInt that represents the ParamSpecInt.
func ParamSpecIntNewFromC(native unsafe.Pointer) *ParamSpecInt {
	return &ParamSpecInt{native: native}
}

// ParamSpecInt64 is a representation of the C record GParamSpecInt64.
type ParamSpecInt64 struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecInt64 that represents the ParamSpecInt64.
func (recv *ParamSpecInt64) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecInt64NewFromC creates a new ParamSpecInt64 from a pointer to the C GParamSpecInt64 that represents the ParamSpecInt64.
func ParamSpecInt64NewFromC(native unsafe.Pointer) *ParamSpecInt64 {
	return &ParamSpecInt64{native: native}
}

// ParamSpecLong is a representation of the C record GParamSpecLong.
type ParamSpecLong struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecLong that represents the ParamSpecLong.
func (recv *ParamSpecLong) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecLongNewFromC creates a new ParamSpecLong from a pointer to the C GParamSpecLong that represents the ParamSpecLong.
func ParamSpecLongNewFromC(native unsafe.Pointer) *ParamSpecLong {
	return &ParamSpecLong{native: native}
}

// ParamSpecObject is a representation of the C record GParamSpecObject.
type ParamSpecObject struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecObject that represents the ParamSpecObject.
func (recv *ParamSpecObject) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecObjectNewFromC creates a new ParamSpecObject from a pointer to the C GParamSpecObject that represents the ParamSpecObject.
func ParamSpecObjectNewFromC(native unsafe.Pointer) *ParamSpecObject {
	return &ParamSpecObject{native: native}
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

// ParamSpecOverrideNewFromC creates a new ParamSpecOverride from a pointer to the C GParamSpecOverride that represents the ParamSpecOverride.
func ParamSpecOverrideNewFromC(native unsafe.Pointer) *ParamSpecOverride {
	return &ParamSpecOverride{native: native}
}

// ParamSpecParam is a representation of the C record GParamSpecParam.
type ParamSpecParam struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecParam that represents the ParamSpecParam.
func (recv *ParamSpecParam) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecParamNewFromC creates a new ParamSpecParam from a pointer to the C GParamSpecParam that represents the ParamSpecParam.
func ParamSpecParamNewFromC(native unsafe.Pointer) *ParamSpecParam {
	return &ParamSpecParam{native: native}
}

// ParamSpecPointer is a representation of the C record GParamSpecPointer.
type ParamSpecPointer struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecPointer that represents the ParamSpecPointer.
func (recv *ParamSpecPointer) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecPointerNewFromC creates a new ParamSpecPointer from a pointer to the C GParamSpecPointer that represents the ParamSpecPointer.
func ParamSpecPointerNewFromC(native unsafe.Pointer) *ParamSpecPointer {
	return &ParamSpecPointer{native: native}
}

// ParamSpecString is a representation of the C record GParamSpecString.
type ParamSpecString struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecString that represents the ParamSpecString.
func (recv *ParamSpecString) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecStringNewFromC creates a new ParamSpecString from a pointer to the C GParamSpecString that represents the ParamSpecString.
func ParamSpecStringNewFromC(native unsafe.Pointer) *ParamSpecString {
	return &ParamSpecString{native: native}
}

// ParamSpecUChar is a representation of the C record GParamSpecUChar.
type ParamSpecUChar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecUChar that represents the ParamSpecUChar.
func (recv *ParamSpecUChar) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecUCharNewFromC creates a new ParamSpecUChar from a pointer to the C GParamSpecUChar that represents the ParamSpecUChar.
func ParamSpecUCharNewFromC(native unsafe.Pointer) *ParamSpecUChar {
	return &ParamSpecUChar{native: native}
}

// ParamSpecUInt is a representation of the C record GParamSpecUInt.
type ParamSpecUInt struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecUInt that represents the ParamSpecUInt.
func (recv *ParamSpecUInt) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecUIntNewFromC creates a new ParamSpecUInt from a pointer to the C GParamSpecUInt that represents the ParamSpecUInt.
func ParamSpecUIntNewFromC(native unsafe.Pointer) *ParamSpecUInt {
	return &ParamSpecUInt{native: native}
}

// ParamSpecUInt64 is a representation of the C record GParamSpecUInt64.
type ParamSpecUInt64 struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecUInt64 that represents the ParamSpecUInt64.
func (recv *ParamSpecUInt64) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecUInt64NewFromC creates a new ParamSpecUInt64 from a pointer to the C GParamSpecUInt64 that represents the ParamSpecUInt64.
func ParamSpecUInt64NewFromC(native unsafe.Pointer) *ParamSpecUInt64 {
	return &ParamSpecUInt64{native: native}
}

// ParamSpecULong is a representation of the C record GParamSpecULong.
type ParamSpecULong struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecULong that represents the ParamSpecULong.
func (recv *ParamSpecULong) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecULongNewFromC creates a new ParamSpecULong from a pointer to the C GParamSpecULong that represents the ParamSpecULong.
func ParamSpecULongNewFromC(native unsafe.Pointer) *ParamSpecULong {
	return &ParamSpecULong{native: native}
}

// ParamSpecUnichar is a representation of the C record GParamSpecUnichar.
type ParamSpecUnichar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecUnichar that represents the ParamSpecUnichar.
func (recv *ParamSpecUnichar) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecUnicharNewFromC creates a new ParamSpecUnichar from a pointer to the C GParamSpecUnichar that represents the ParamSpecUnichar.
func ParamSpecUnicharNewFromC(native unsafe.Pointer) *ParamSpecUnichar {
	return &ParamSpecUnichar{native: native}
}

// ParamSpecValueArray is a representation of the C record GParamSpecValueArray.
type ParamSpecValueArray struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecValueArray that represents the ParamSpecValueArray.
func (recv *ParamSpecValueArray) ToC() unsafe.Pointer {
	return recv.native
}

// ParamSpecValueArrayNewFromC creates a new ParamSpecValueArray from a pointer to the C GParamSpecValueArray that represents the ParamSpecValueArray.
func ParamSpecValueArrayNewFromC(native unsafe.Pointer) *ParamSpecValueArray {
	return &ParamSpecValueArray{native: native}
}

// TypeModule is a representation of the C record GTypeModule.
type TypeModule struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeModule that represents the TypeModule.
func (recv *TypeModule) ToC() unsafe.Pointer {
	return recv.native
}

// TypeModuleNewFromC creates a new TypeModule from a pointer to the C GTypeModule that represents the TypeModule.
func TypeModuleNewFromC(native unsafe.Pointer) *TypeModule {
	return &TypeModule{native: native}
}

// TypePlugin is a representation of the C interface GTypePlugin.
type TypePlugin struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypePlugin that represents the TypePlugin.
func (recv *TypePlugin) ToC() unsafe.Pointer {
	return recv.native
}

// TypePluginNewFromC creates a new TypePlugin from a pointer to the C GTypePlugin that represents the TypePlugin.
func TypePluginNewFromC(native unsafe.Pointer) *TypePlugin {
	return &TypePlugin{native: native}
}

// TypeCValue is a representation of the C union GTypeCValue.
type TypeCValue struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeCValue that represents the TypeCValue.
func (recv *TypeCValue) ToC() unsafe.Pointer {
	return recv.native
}

// TypeCValueNewFromC creates a new TypeCValue from a pointer to the C GTypeCValue that represents the TypeCValue.
func TypeCValueNewFromC(native unsafe.Pointer) *TypeCValue {
	return &TypeCValue{native: native}
}

// UNSUPPORTED : _Value__data__union : blacklisted
