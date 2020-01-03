// Code generated - DO NOT EDIT.
// +build gobject_2.30

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

// BoxedCopy is analogous to the C function g_boxed_copy.
func BoxedCopy(boxedType uint64, srcBoxed unsafe.Pointer) unsafe.Pointer {
	sys_boxedType := boxedType
	sys_srcBoxed := srcBoxed
	ret := gobject.Fn_g_boxed_copy(sys_boxedType, sys_srcBoxed)

	return ret
}

// BoxedFree is analogous to the C function g_boxed_free.
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

// ClearObject is analogous to the C function g_clear_object.
func ClearObject(objectPtr **Object) {
	sys_objectPtr := objectPtr.ToC()
	gobject.Fn_g_clear_object(sys_objectPtr)
}

// UNSUPPORTED : g_enum_complete_type_info : has array [in]out, info

// EnumGetValue is analogous to the C function g_enum_get_value.
func EnumGetValue(enumClass *EnumClass, value int) unsafe.Pointer {
	sys_enumClass := enumClass.ToC()
	sys_value := value
	ret := gobject.Fn_g_enum_get_value(sys_enumClass, sys_value)

	return ret
}

// EnumGetValueByName is analogous to the C function g_enum_get_value_by_name.
func EnumGetValueByName(enumClass *EnumClass, name string) unsafe.Pointer {
	sys_enumClass := enumClass.ToC()
	sys_name := name
	ret := gobject.Fn_g_enum_get_value_by_name(sys_enumClass, sys_name)

	return ret
}

// EnumGetValueByNick is analogous to the C function g_enum_get_value_by_nick.
func EnumGetValueByNick(enumClass *EnumClass, nick string) unsafe.Pointer {
	sys_enumClass := enumClass.ToC()
	sys_nick := nick
	ret := gobject.Fn_g_enum_get_value_by_nick(sys_enumClass, sys_nick)

	return ret
}

// EnumRegisterStatic is analogous to the C function g_enum_register_static.
func EnumRegisterStatic(name string, constStaticValues *EnumValue) uint64 {
	sys_name := name
	sys_constStaticValues := constStaticValues.ToC()
	ret := gobject.Fn_g_enum_register_static(sys_name, sys_constStaticValues)

	return ret
}

// UNSUPPORTED : g_flags_complete_type_info : has array [in]out, info

// FlagsGetFirstValue is analogous to the C function g_flags_get_first_value.
func FlagsGetFirstValue(flagsClass *FlagsClass, value uint) unsafe.Pointer {
	sys_flagsClass := flagsClass.ToC()
	sys_value := value
	ret := gobject.Fn_g_flags_get_first_value(sys_flagsClass, sys_value)

	return ret
}

// FlagsGetValueByName is analogous to the C function g_flags_get_value_by_name.
func FlagsGetValueByName(flagsClass *FlagsClass, name string) unsafe.Pointer {
	sys_flagsClass := flagsClass.ToC()
	sys_name := name
	ret := gobject.Fn_g_flags_get_value_by_name(sys_flagsClass, sys_name)

	return ret
}

// FlagsGetValueByNick is analogous to the C function g_flags_get_value_by_nick.
func FlagsGetValueByNick(flagsClass *FlagsClass, nick string) unsafe.Pointer {
	sys_flagsClass := flagsClass.ToC()
	sys_nick := nick
	ret := gobject.Fn_g_flags_get_value_by_nick(sys_flagsClass, sys_nick)

	return ret
}

// FlagsRegisterStatic is analogous to the C function g_flags_register_static.
func FlagsRegisterStatic(name string, constStaticValues *FlagsValue) uint64 {
	sys_name := name
	sys_constStaticValues := constStaticValues.ToC()
	ret := gobject.Fn_g_flags_register_static(sys_name, sys_constStaticValues)

	return ret
}

// GtypeGetType is analogous to the C function g_gtype_get_type.
func GtypeGetType() uint64 {
	ret := gobject.Fn_g_gtype_get_type()

	return ret
}

// ParamSpecBoolean_ is analogous to the C function g_param_spec_boolean.
func ParamSpecBoolean_(name string, nick string, blurb string, defaultValue bool, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_defaultValue := defaultValue
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_boolean(sys_name, sys_nick, sys_blurb, sys_defaultValue, sys_flags)

	return ret
}

// ParamSpecBoxed_ is analogous to the C function g_param_spec_boxed.
func ParamSpecBoxed_(name string, nick string, blurb string, boxedType uint64, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_boxedType := boxedType
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_boxed(sys_name, sys_nick, sys_blurb, sys_boxedType, sys_flags)

	return ret
}

// ParamSpecChar_ is analogous to the C function g_param_spec_char.
func ParamSpecChar_(name string, nick string, blurb string, minimum int8, maximum int8, defaultValue int8, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_char(sys_name, sys_nick, sys_blurb, sys_minimum, sys_maximum, sys_defaultValue, sys_flags)

	return ret
}

// ParamSpecDouble_ is analogous to the C function g_param_spec_double.
func ParamSpecDouble_(name string, nick string, blurb string, minimum float64, maximum float64, defaultValue float64, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_double(sys_name, sys_nick, sys_blurb, sys_minimum, sys_maximum, sys_defaultValue, sys_flags)

	return ret
}

// ParamSpecEnum_ is analogous to the C function g_param_spec_enum.
func ParamSpecEnum_(name string, nick string, blurb string, enumType uint64, defaultValue int, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_enumType := enumType
	sys_defaultValue := defaultValue
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_enum(sys_name, sys_nick, sys_blurb, sys_enumType, sys_defaultValue, sys_flags)

	return ret
}

// ParamSpecFlags_ is analogous to the C function g_param_spec_flags.
func ParamSpecFlags_(name string, nick string, blurb string, flagsType uint64, defaultValue uint, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_flagsType := flagsType
	sys_defaultValue := defaultValue
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_flags(sys_name, sys_nick, sys_blurb, sys_flagsType, sys_defaultValue, sys_flags)

	return ret
}

// ParamSpecFloat_ is analogous to the C function g_param_spec_float.
func ParamSpecFloat_(name string, nick string, blurb string, minimum float32, maximum float32, defaultValue float32, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_float(sys_name, sys_nick, sys_blurb, sys_minimum, sys_maximum, sys_defaultValue, sys_flags)

	return ret
}

// ParamSpecGtype is analogous to the C function g_param_spec_gtype.
func ParamSpecGtype(name string, nick string, blurb string, isAType uint64, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_isAType := isAType
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_gtype(sys_name, sys_nick, sys_blurb, sys_isAType, sys_flags)

	return ret
}

// ParamSpecInt_ is analogous to the C function g_param_spec_int.
func ParamSpecInt_(name string, nick string, blurb string, minimum int, maximum int, defaultValue int, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_int(sys_name, sys_nick, sys_blurb, sys_minimum, sys_maximum, sys_defaultValue, sys_flags)

	return ret
}

// ParamSpecInt64_ is analogous to the C function g_param_spec_int64.
func ParamSpecInt64_(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_int64(sys_name, sys_nick, sys_blurb, sys_minimum, sys_maximum, sys_defaultValue, sys_flags)

	return ret
}

// ParamSpecLong_ is analogous to the C function g_param_spec_long.
func ParamSpecLong_(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_long(sys_name, sys_nick, sys_blurb, sys_minimum, sys_maximum, sys_defaultValue, sys_flags)

	return ret
}

// ParamSpecObject_ is analogous to the C function g_param_spec_object.
func ParamSpecObject_(name string, nick string, blurb string, objectType uint64, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_objectType := objectType
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_object(sys_name, sys_nick, sys_blurb, sys_objectType, sys_flags)

	return ret
}

// ParamSpecOverride_ is analogous to the C function g_param_spec_override.
func ParamSpecOverride_(name string, overridden *ParamSpec) unsafe.Pointer {
	sys_name := name
	sys_overridden := overridden.ToC()
	ret := gobject.Fn_g_param_spec_override(sys_name, sys_overridden)

	return ret
}

// ParamSpecParam_ is analogous to the C function g_param_spec_param.
func ParamSpecParam_(name string, nick string, blurb string, paramType uint64, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_paramType := paramType
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_param(sys_name, sys_nick, sys_blurb, sys_paramType, sys_flags)

	return ret
}

// ParamSpecPointer_ is analogous to the C function g_param_spec_pointer.
func ParamSpecPointer_(name string, nick string, blurb string, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_pointer(sys_name, sys_nick, sys_blurb, sys_flags)

	return ret
}

// ParamSpecString_ is analogous to the C function g_param_spec_string.
func ParamSpecString_(name string, nick string, blurb string, defaultValue string, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_defaultValue := defaultValue
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_string(sys_name, sys_nick, sys_blurb, sys_defaultValue, sys_flags)

	return ret
}

// ParamSpecUchar is analogous to the C function g_param_spec_uchar.
func ParamSpecUchar(name string, nick string, blurb string, minimum uint8, maximum uint8, defaultValue uint8, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_uchar(sys_name, sys_nick, sys_blurb, sys_minimum, sys_maximum, sys_defaultValue, sys_flags)

	return ret
}

// ParamSpecUint is analogous to the C function g_param_spec_uint.
func ParamSpecUint(name string, nick string, blurb string, minimum uint, maximum uint, defaultValue uint, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_uint(sys_name, sys_nick, sys_blurb, sys_minimum, sys_maximum, sys_defaultValue, sys_flags)

	return ret
}

// ParamSpecUint64 is analogous to the C function g_param_spec_uint64.
func ParamSpecUint64(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_uint64(sys_name, sys_nick, sys_blurb, sys_minimum, sys_maximum, sys_defaultValue, sys_flags)

	return ret
}

// ParamSpecUlong is analogous to the C function g_param_spec_ulong.
func ParamSpecUlong(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_minimum := minimum
	sys_maximum := maximum
	sys_defaultValue := defaultValue
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_ulong(sys_name, sys_nick, sys_blurb, sys_minimum, sys_maximum, sys_defaultValue, sys_flags)

	return ret
}

// ParamSpecUnichar_ is analogous to the C function g_param_spec_unichar.
func ParamSpecUnichar_(name string, nick string, blurb string, defaultValue rune, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_defaultValue := defaultValue
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_unichar(sys_name, sys_nick, sys_blurb, sys_defaultValue, sys_flags)

	return ret
}

// ParamSpecValueArray_ is analogous to the C function g_param_spec_value_array.
func ParamSpecValueArray_(name string, nick string, blurb string, elementSpec *ParamSpec, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_elementSpec := elementSpec.ToC()
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_value_array(sys_name, sys_nick, sys_blurb, sys_elementSpec, sys_flags)

	return ret
}

// ParamSpecVariant_ is analogous to the C function g_param_spec_variant.
func ParamSpecVariant_(name string, nick string, blurb string, type_ *glib.VariantType, defaultValue *glib.Variant, flags int) unsafe.Pointer {
	sys_name := name
	sys_nick := nick
	sys_blurb := blurb
	sys_type_ := type_.ToC()
	sys_defaultValue := defaultValue.ToC()
	sys_flags := flags
	ret := gobject.Fn_g_param_spec_variant(sys_name, sys_nick, sys_blurb, sys_type_, sys_defaultValue, sys_flags)

	return ret
}

// ParamTypeRegisterStatic is analogous to the C function g_param_type_register_static.
func ParamTypeRegisterStatic(name string, pspecInfo *ParamSpecTypeInfo) uint64 {
	sys_name := name
	sys_pspecInfo := pspecInfo.ToC()
	ret := gobject.Fn_g_param_type_register_static(sys_name, sys_pspecInfo)

	return ret
}

// ParamValueConvert is analogous to the C function g_param_value_convert.
func ParamValueConvert(pspec *ParamSpec, srcValue *Value, destValue *Value, strictValidation bool) bool {
	sys_pspec := pspec.ToC()
	sys_srcValue := srcValue.ToC()
	sys_destValue := destValue.ToC()
	sys_strictValidation := strictValidation
	ret := gobject.Fn_g_param_value_convert(sys_pspec, sys_srcValue, sys_destValue, sys_strictValidation)

	return ret
}

// ParamValueDefaults is analogous to the C function g_param_value_defaults.
func ParamValueDefaults(pspec *ParamSpec, value *Value) bool {
	sys_pspec := pspec.ToC()
	sys_value := value.ToC()
	ret := gobject.Fn_g_param_value_defaults(sys_pspec, sys_value)

	return ret
}

// ParamValueSetDefault is analogous to the C function g_param_value_set_default.
func ParamValueSetDefault(pspec *ParamSpec, value *Value) {
	sys_pspec := pspec.ToC()
	sys_value := value.ToC()
	gobject.Fn_g_param_value_set_default(sys_pspec, sys_value)
}

// ParamValueValidate is analogous to the C function g_param_value_validate.
func ParamValueValidate(pspec *ParamSpec, value *Value) bool {
	sys_pspec := pspec.ToC()
	sys_value := value.ToC()
	ret := gobject.Fn_g_param_value_validate(sys_pspec, sys_value)

	return ret
}

// ParamValuesCmp is analogous to the C function g_param_values_cmp.
func ParamValuesCmp(pspec *ParamSpec, value1 *Value, value2 *Value) int {
	sys_pspec := pspec.ToC()
	sys_value1 := value1.ToC()
	sys_value2 := value2.ToC()
	ret := gobject.Fn_g_param_values_cmp(sys_pspec, sys_value1, sys_value2)

	return ret
}

// PointerTypeRegisterStatic is analogous to the C function g_pointer_type_register_static.
func PointerTypeRegisterStatic(name string) uint64 {
	sys_name := name
	ret := gobject.Fn_g_pointer_type_register_static(sys_name)

	return ret
}

// SignalAccumulatorFirstWins is analogous to the C function g_signal_accumulator_first_wins.
func SignalAccumulatorFirstWins(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value, dummy unsafe.Pointer) bool {
	sys_ihint := ihint.ToC()
	sys_returnAccu := returnAccu.ToC()
	sys_handlerReturn := handlerReturn.ToC()
	sys_dummy := dummy
	ret := gobject.Fn_g_signal_accumulator_first_wins(sys_ihint, sys_returnAccu, sys_handlerReturn, sys_dummy)

	return ret
}

// SignalAccumulatorTrueHandled is analogous to the C function g_signal_accumulator_true_handled.
func SignalAccumulatorTrueHandled(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value, dummy unsafe.Pointer) bool {
	sys_ihint := ihint.ToC()
	sys_returnAccu := returnAccu.ToC()
	sys_handlerReturn := handlerReturn.ToC()
	sys_dummy := dummy
	ret := gobject.Fn_g_signal_accumulator_true_handled(sys_ihint, sys_returnAccu, sys_handlerReturn, sys_dummy)

	return ret
}

// UNSUPPORTED : g_signal_add_emission_hook : parameter 'hook_func' is callback

// UNSUPPORTED : g_signal_chain_from_overridden : parameter 'instance_and_params' is array parameter without length parameter

// SignalChainFromOverriddenHandler is analogous to the C function g_signal_chain_from_overridden_handler.
func SignalChainFromOverriddenHandler(instance unsafe.Pointer) {
	sys_instance := instance
	gobject.Fn_g_signal_chain_from_overridden_handler(sys_instance)
}

// SignalConnectClosure is analogous to the C function g_signal_connect_closure.
func SignalConnectClosure(instance unsafe.Pointer, detailedSignal string, closure *Closure, after bool) uint64 {
	sys_instance := instance
	sys_detailedSignal := detailedSignal
	sys_closure := closure.ToC()
	sys_after := after
	ret := gobject.Fn_g_signal_connect_closure(sys_instance, sys_detailedSignal, sys_closure, sys_after)

	return ret
}

// SignalConnectClosureById is analogous to the C function g_signal_connect_closure_by_id.
func SignalConnectClosureById(instance unsafe.Pointer, signalId uint, detail uint32, closure *Closure, after bool) uint64 {
	sys_instance := instance
	sys_signalId := signalId
	sys_detail := detail
	sys_closure := closure.ToC()
	sys_after := after
	ret := gobject.Fn_g_signal_connect_closure_by_id(sys_instance, sys_signalId, sys_detail, sys_closure, sys_after)

	return ret
}

// UNSUPPORTED : g_signal_connect_data : parameter 'c_handler' is callback

// UNSUPPORTED : g_signal_connect_object : parameter 'c_handler' is callback

// SignalEmit is analogous to the C function g_signal_emit.
func SignalEmit(instance unsafe.Pointer, signalId uint, detail uint32) {
	sys_instance := instance
	sys_signalId := signalId
	sys_detail := detail
	gobject.Fn_g_signal_emit(sys_instance, sys_signalId, sys_detail)
}

// SignalEmitByName is analogous to the C function g_signal_emit_by_name.
func SignalEmitByName(instance unsafe.Pointer, detailedSignal string) {
	sys_instance := instance
	sys_detailedSignal := detailedSignal
	gobject.Fn_g_signal_emit_by_name(sys_instance, sys_detailedSignal)
}

// SignalEmitValist is analogous to the C function g_signal_emit_valist.
func SignalEmitValist(instance unsafe.Pointer, signalId uint, detail uint32) {
	sys_instance := instance
	sys_signalId := signalId
	sys_detail := detail
	gobject.Fn_g_signal_emit_valist(sys_instance, sys_signalId, sys_detail)
}

// UNSUPPORTED : g_signal_emitv : parameter 'instance_and_params' is array parameter without length parameter

// SignalGetInvocationHint is analogous to the C function g_signal_get_invocation_hint.
func SignalGetInvocationHint(instance unsafe.Pointer) unsafe.Pointer {
	sys_instance := instance
	ret := gobject.Fn_g_signal_get_invocation_hint(sys_instance)

	return ret
}

// SignalHandlerBlock is analogous to the C function g_signal_handler_block.
func SignalHandlerBlock(instance unsafe.Pointer, handlerId uint64) {
	sys_instance := instance
	sys_handlerId := handlerId
	gobject.Fn_g_signal_handler_block(sys_instance, sys_handlerId)
}

// SignalHandlerDisconnect is analogous to the C function g_signal_handler_disconnect.
func SignalHandlerDisconnect(instance unsafe.Pointer, handlerId uint64) {
	sys_instance := instance
	sys_handlerId := handlerId
	gobject.Fn_g_signal_handler_disconnect(sys_instance, sys_handlerId)
}

// SignalHandlerFind is analogous to the C function g_signal_handler_find.
func SignalHandlerFind(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) uint64 {
	sys_instance := instance
	sys_mask := mask
	sys_signalId := signalId
	sys_detail := detail
	sys_closure := closure.ToC()
	sys_func_ := func_
	sys_data := data
	ret := gobject.Fn_g_signal_handler_find(sys_instance, sys_mask, sys_signalId, sys_detail, sys_closure, sys_func_, sys_data)

	return ret
}

// SignalHandlerIsConnected is analogous to the C function g_signal_handler_is_connected.
func SignalHandlerIsConnected(instance unsafe.Pointer, handlerId uint64) bool {
	sys_instance := instance
	sys_handlerId := handlerId
	ret := gobject.Fn_g_signal_handler_is_connected(sys_instance, sys_handlerId)

	return ret
}

// SignalHandlerUnblock is analogous to the C function g_signal_handler_unblock.
func SignalHandlerUnblock(instance unsafe.Pointer, handlerId uint64) {
	sys_instance := instance
	sys_handlerId := handlerId
	gobject.Fn_g_signal_handler_unblock(sys_instance, sys_handlerId)
}

// SignalHandlersBlockMatched is analogous to the C function g_signal_handlers_block_matched.
func SignalHandlersBlockMatched(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) uint {
	sys_instance := instance
	sys_mask := mask
	sys_signalId := signalId
	sys_detail := detail
	sys_closure := closure.ToC()
	sys_func_ := func_
	sys_data := data
	ret := gobject.Fn_g_signal_handlers_block_matched(sys_instance, sys_mask, sys_signalId, sys_detail, sys_closure, sys_func_, sys_data)

	return ret
}

// SignalHandlersDestroy is analogous to the C function g_signal_handlers_destroy.
func SignalHandlersDestroy(instance unsafe.Pointer) {
	sys_instance := instance
	gobject.Fn_g_signal_handlers_destroy(sys_instance)
}

// SignalHandlersDisconnectMatched is analogous to the C function g_signal_handlers_disconnect_matched.
func SignalHandlersDisconnectMatched(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) uint {
	sys_instance := instance
	sys_mask := mask
	sys_signalId := signalId
	sys_detail := detail
	sys_closure := closure.ToC()
	sys_func_ := func_
	sys_data := data
	ret := gobject.Fn_g_signal_handlers_disconnect_matched(sys_instance, sys_mask, sys_signalId, sys_detail, sys_closure, sys_func_, sys_data)

	return ret
}

// SignalHandlersUnblockMatched is analogous to the C function g_signal_handlers_unblock_matched.
func SignalHandlersUnblockMatched(instance unsafe.Pointer, mask int, signalId uint, detail uint32, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) uint {
	sys_instance := instance
	sys_mask := mask
	sys_signalId := signalId
	sys_detail := detail
	sys_closure := closure.ToC()
	sys_func_ := func_
	sys_data := data
	ret := gobject.Fn_g_signal_handlers_unblock_matched(sys_instance, sys_mask, sys_signalId, sys_detail, sys_closure, sys_func_, sys_data)

	return ret
}

// SignalHasHandlerPending is analogous to the C function g_signal_has_handler_pending.
func SignalHasHandlerPending(instance unsafe.Pointer, signalId uint, detail uint32, mayBeBlocked bool) bool {
	sys_instance := instance
	sys_signalId := signalId
	sys_detail := detail
	sys_mayBeBlocked := mayBeBlocked
	ret := gobject.Fn_g_signal_has_handler_pending(sys_instance, sys_signalId, sys_detail, sys_mayBeBlocked)

	return ret
}

// UNSUPPORTED : g_signal_list_ids : has array [in]out, n_ids

// SignalLookup is analogous to the C function g_signal_lookup.
func SignalLookup(name string, itype uint64) uint {
	sys_name := name
	sys_itype := itype
	ret := gobject.Fn_g_signal_lookup(sys_name, sys_itype)

	return ret
}

// SignalName is analogous to the C function g_signal_name.
func SignalName(signalId uint) string {
	sys_signalId := signalId
	ret := gobject.Fn_g_signal_name(sys_signalId)

	return ret
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
	gobject.Fn_g_signal_override_class_closure(sys_signalId, sys_instanceType, sys_classClosure)
}

// UNSUPPORTED : g_signal_override_class_handler : parameter 'class_handler' is callback

// UNSUPPORTED : g_signal_parse_name : has array [in]out, signal_id_p

// UNSUPPORTED : g_signal_query : has array [in]out, query

// SignalRemoveEmissionHook is analogous to the C function g_signal_remove_emission_hook.
func SignalRemoveEmissionHook(signalId uint, hookId uint64) {
	sys_signalId := signalId
	sys_hookId := hookId
	gobject.Fn_g_signal_remove_emission_hook(sys_signalId, sys_hookId)
}

// UNSUPPORTED : g_signal_set_va_marshaller : blacklisted

// SignalStopEmission is analogous to the C function g_signal_stop_emission.
func SignalStopEmission(instance unsafe.Pointer, signalId uint, detail uint32) {
	sys_instance := instance
	sys_signalId := signalId
	sys_detail := detail
	gobject.Fn_g_signal_stop_emission(sys_instance, sys_signalId, sys_detail)
}

// SignalStopEmissionByName is analogous to the C function g_signal_stop_emission_by_name.
func SignalStopEmissionByName(instance unsafe.Pointer, detailedSignal string) {
	sys_instance := instance
	sys_detailedSignal := detailedSignal
	gobject.Fn_g_signal_stop_emission_by_name(sys_instance, sys_detailedSignal)
}

// SignalTypeCclosureNew is analogous to the C function g_signal_type_cclosure_new.
func SignalTypeCclosureNew(itype uint64, structOffset uint) unsafe.Pointer {
	sys_itype := itype
	sys_structOffset := structOffset
	ret := gobject.Fn_g_signal_type_cclosure_new(sys_itype, sys_structOffset)

	return ret
}

// SourceSetClosure is analogous to the C function g_source_set_closure.
func SourceSetClosure(source *glib.Source, closure *Closure) {
	sys_source := source.ToC()
	sys_closure := closure.ToC()
	gobject.Fn_g_source_set_closure(sys_source, sys_closure)
}

// SourceSetDummyCallback is analogous to the C function g_source_set_dummy_callback.
func SourceSetDummyCallback(source *glib.Source) {
	sys_source := source.ToC()
	gobject.Fn_g_source_set_dummy_callback(sys_source)
}

// StrdupValueContents is analogous to the C function g_strdup_value_contents.
func StrdupValueContents(value *Value) string {
	sys_value := value.ToC()
	ret := gobject.Fn_g_strdup_value_contents(sys_value)

	return ret
}

// UNSUPPORTED : g_type_add_class_cache_func : parameter 'cache_func' is callback

// TypeAddClassPrivate is analogous to the C function g_type_add_class_private.
func TypeAddClassPrivate(classType uint64, privateSize uint64) {
	sys_classType := classType
	sys_privateSize := privateSize
	gobject.Fn_g_type_add_class_private(sys_classType, sys_privateSize)
}

// TypeAddInstancePrivate is analogous to the C function g_type_add_instance_private.
func TypeAddInstancePrivate(classType uint64, privateSize uint64) int {
	sys_classType := classType
	sys_privateSize := privateSize
	ret := gobject.Fn_g_type_add_instance_private(sys_classType, sys_privateSize)

	return ret
}

// UNSUPPORTED : g_type_add_interface_check : parameter 'check_func' is callback

// TypeAddInterfaceDynamic is analogous to the C function g_type_add_interface_dynamic.
func TypeAddInterfaceDynamic(instanceType uint64, interfaceType uint64, plugin *TypePlugin) {
	sys_instanceType := instanceType
	sys_interfaceType := interfaceType
	sys_plugin := plugin.ToC()
	gobject.Fn_g_type_add_interface_dynamic(sys_instanceType, sys_interfaceType, sys_plugin)
}

// TypeAddInterfaceStatic is analogous to the C function g_type_add_interface_static.
func TypeAddInterfaceStatic(instanceType uint64, interfaceType uint64, info *InterfaceInfo) {
	sys_instanceType := instanceType
	sys_interfaceType := interfaceType
	sys_info := info.ToC()
	gobject.Fn_g_type_add_interface_static(sys_instanceType, sys_interfaceType, sys_info)
}

// TypeCheckClassCast is analogous to the C function g_type_check_class_cast.
func TypeCheckClassCast(gClass *TypeClass, isAType uint64) unsafe.Pointer {
	sys_gClass := gClass.ToC()
	sys_isAType := isAType
	ret := gobject.Fn_g_type_check_class_cast(sys_gClass, sys_isAType)

	return ret
}

// TypeCheckClassIsA is analogous to the C function g_type_check_class_is_a.
func TypeCheckClassIsA(gClass *TypeClass, isAType uint64) bool {
	sys_gClass := gClass.ToC()
	sys_isAType := isAType
	ret := gobject.Fn_g_type_check_class_is_a(sys_gClass, sys_isAType)

	return ret
}

// TypeCheckInstance is analogous to the C function g_type_check_instance.
func TypeCheckInstance(instance *TypeInstance) bool {
	sys_instance := instance.ToC()
	ret := gobject.Fn_g_type_check_instance(sys_instance)

	return ret
}

// TypeCheckInstanceCast is analogous to the C function g_type_check_instance_cast.
func TypeCheckInstanceCast(instance *TypeInstance, ifaceType uint64) unsafe.Pointer {
	sys_instance := instance.ToC()
	sys_ifaceType := ifaceType
	ret := gobject.Fn_g_type_check_instance_cast(sys_instance, sys_ifaceType)

	return ret
}

// TypeCheckInstanceIsA is analogous to the C function g_type_check_instance_is_a.
func TypeCheckInstanceIsA(instance *TypeInstance, ifaceType uint64) bool {
	sys_instance := instance.ToC()
	sys_ifaceType := ifaceType
	ret := gobject.Fn_g_type_check_instance_is_a(sys_instance, sys_ifaceType)

	return ret
}

// TypeCheckInstanceIsFundamentallyA is analogous to the C function g_type_check_instance_is_fundamentally_a.
func TypeCheckInstanceIsFundamentallyA(instance *TypeInstance, fundamentalType uint64) bool {
	sys_instance := instance.ToC()
	sys_fundamentalType := fundamentalType
	ret := gobject.Fn_g_type_check_instance_is_fundamentally_a(sys_instance, sys_fundamentalType)

	return ret
}

// TypeCheckIsValueType is analogous to the C function g_type_check_is_value_type.
func TypeCheckIsValueType(type_ uint64) bool {
	sys_type_ := type_
	ret := gobject.Fn_g_type_check_is_value_type(sys_type_)

	return ret
}

// TypeCheckValue is analogous to the C function g_type_check_value.
func TypeCheckValue(value *Value) bool {
	sys_value := value.ToC()
	ret := gobject.Fn_g_type_check_value(sys_value)

	return ret
}

// TypeCheckValueHolds is analogous to the C function g_type_check_value_holds.
func TypeCheckValueHolds(value *Value, type_ uint64) bool {
	sys_value := value.ToC()
	sys_type_ := type_
	ret := gobject.Fn_g_type_check_value_holds(sys_value, sys_type_)

	return ret
}

// UNSUPPORTED : g_type_children : has array [in]out, n_children

// TypeCreateInstance is analogous to the C function g_type_create_instance.
func TypeCreateInstance(type_ uint64) unsafe.Pointer {
	sys_type_ := type_
	ret := gobject.Fn_g_type_create_instance(sys_type_)

	return ret
}

// TypeDefaultInterfacePeek is analogous to the C function g_type_default_interface_peek.
func TypeDefaultInterfacePeek(gType uint64) unsafe.Pointer {
	sys_gType := gType
	ret := gobject.Fn_g_type_default_interface_peek(sys_gType)

	return ret
}

// TypeDefaultInterfaceRef is analogous to the C function g_type_default_interface_ref.
func TypeDefaultInterfaceRef(gType uint64) unsafe.Pointer {
	sys_gType := gType
	ret := gobject.Fn_g_type_default_interface_ref(sys_gType)

	return ret
}

// TypeDefaultInterfaceUnref is analogous to the C function g_type_default_interface_unref.
func TypeDefaultInterfaceUnref(gIface unsafe.Pointer) {
	sys_gIface := gIface
	gobject.Fn_g_type_default_interface_unref(sys_gIface)
}

// TypeDepth is analogous to the C function g_type_depth.
func TypeDepth(type_ uint64) uint {
	sys_type_ := type_
	ret := gobject.Fn_g_type_depth(sys_type_)

	return ret
}

// TypeFreeInstance is analogous to the C function g_type_free_instance.
func TypeFreeInstance(instance *TypeInstance) {
	sys_instance := instance.ToC()
	gobject.Fn_g_type_free_instance(sys_instance)
}

// TypeFromName is analogous to the C function g_type_from_name.
func TypeFromName(name string) uint64 {
	sys_name := name
	ret := gobject.Fn_g_type_from_name(sys_name)

	return ret
}

// TypeFundamental is analogous to the C function g_type_fundamental.
func TypeFundamental(typeId uint64) uint64 {
	sys_typeId := typeId
	ret := gobject.Fn_g_type_fundamental(sys_typeId)

	return ret
}

// TypeFundamentalNext is analogous to the C function g_type_fundamental_next.
func TypeFundamentalNext() uint64 {
	ret := gobject.Fn_g_type_fundamental_next()

	return ret
}

// TypeGetPlugin is analogous to the C function g_type_get_plugin.
func TypeGetPlugin(type_ uint64) unsafe.Pointer {
	sys_type_ := type_
	ret := gobject.Fn_g_type_get_plugin(sys_type_)

	return ret
}

// TypeGetQdata is analogous to the C function g_type_get_qdata.
func TypeGetQdata(type_ uint64, quark uint32) unsafe.Pointer {
	sys_type_ := type_
	sys_quark := quark
	ret := gobject.Fn_g_type_get_qdata(sys_type_, sys_quark)

	return ret
}

// TypeInit is analogous to the C function g_type_init.
func TypeInit() {
	gobject.Fn_g_type_init()
}

// TypeInitWithDebugFlags is analogous to the C function g_type_init_with_debug_flags.
func TypeInitWithDebugFlags(debugFlags int) {
	sys_debugFlags := debugFlags
	gobject.Fn_g_type_init_with_debug_flags(sys_debugFlags)
}

// UNSUPPORTED : g_type_interface_prerequisites : has array [in]out, n_prerequisites

// UNSUPPORTED : g_type_interfaces : has array [in]out, n_interfaces

// TypeIsA is analogous to the C function g_type_is_a.
func TypeIsA(type_ uint64, isAType uint64) bool {
	sys_type_ := type_
	sys_isAType := isAType
	ret := gobject.Fn_g_type_is_a(sys_type_, sys_isAType)

	return ret
}

// TypeName is analogous to the C function g_type_name.
func TypeName(type_ uint64) string {
	sys_type_ := type_
	ret := gobject.Fn_g_type_name(sys_type_)

	return ret
}

// TypeNameFromClass is analogous to the C function g_type_name_from_class.
func TypeNameFromClass(gClass *TypeClass) string {
	sys_gClass := gClass.ToC()
	ret := gobject.Fn_g_type_name_from_class(sys_gClass)

	return ret
}

// TypeNameFromInstance is analogous to the C function g_type_name_from_instance.
func TypeNameFromInstance(instance *TypeInstance) string {
	sys_instance := instance.ToC()
	ret := gobject.Fn_g_type_name_from_instance(sys_instance)

	return ret
}

// TypeNextBase is analogous to the C function g_type_next_base.
func TypeNextBase(leafType uint64, rootType uint64) uint64 {
	sys_leafType := leafType
	sys_rootType := rootType
	ret := gobject.Fn_g_type_next_base(sys_leafType, sys_rootType)

	return ret
}

// TypeParent is analogous to the C function g_type_parent.
func TypeParent(type_ uint64) uint64 {
	sys_type_ := type_
	ret := gobject.Fn_g_type_parent(sys_type_)

	return ret
}

// TypeQname is analogous to the C function g_type_qname.
func TypeQname(type_ uint64) uint32 {
	sys_type_ := type_
	ret := gobject.Fn_g_type_qname(sys_type_)

	return ret
}

// UNSUPPORTED : g_type_query : has array [in]out, query

// TypeRegisterDynamic is analogous to the C function g_type_register_dynamic.
func TypeRegisterDynamic(parentType uint64, typeName string, plugin *TypePlugin, flags int) uint64 {
	sys_parentType := parentType
	sys_typeName := typeName
	sys_plugin := plugin.ToC()
	sys_flags := flags
	ret := gobject.Fn_g_type_register_dynamic(sys_parentType, sys_typeName, sys_plugin, sys_flags)

	return ret
}

// TypeRegisterFundamental is analogous to the C function g_type_register_fundamental.
func TypeRegisterFundamental(typeId uint64, typeName string, info *TypeInfo, finfo *TypeFundamentalInfo, flags int) uint64 {
	sys_typeId := typeId
	sys_typeName := typeName
	sys_info := info.ToC()
	sys_finfo := finfo.ToC()
	sys_flags := flags
	ret := gobject.Fn_g_type_register_fundamental(sys_typeId, sys_typeName, sys_info, sys_finfo, sys_flags)

	return ret
}

// TypeRegisterStatic is analogous to the C function g_type_register_static.
func TypeRegisterStatic(parentType uint64, typeName string, info *TypeInfo, flags int) uint64 {
	sys_parentType := parentType
	sys_typeName := typeName
	sys_info := info.ToC()
	sys_flags := flags
	ret := gobject.Fn_g_type_register_static(sys_parentType, sys_typeName, sys_info, sys_flags)

	return ret
}

// UNSUPPORTED : g_type_register_static_simple : parameter 'class_init' is callback

// UNSUPPORTED : g_type_remove_class_cache_func : parameter 'cache_func' is callback

// UNSUPPORTED : g_type_remove_interface_check : parameter 'check_func' is callback

// TypeSetQdata is analogous to the C function g_type_set_qdata.
func TypeSetQdata(type_ uint64, quark uint32, data unsafe.Pointer) {
	sys_type_ := type_
	sys_quark := quark
	sys_data := data
	gobject.Fn_g_type_set_qdata(sys_type_, sys_quark, sys_data)
}

// TypeTestFlags is analogous to the C function g_type_test_flags.
func TypeTestFlags(type_ uint64, flags uint) bool {
	sys_type_ := type_
	sys_flags := flags
	ret := gobject.Fn_g_type_test_flags(sys_type_, sys_flags)

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

// Binding is a representation of the C record GBinding.
//
// since 2.26
type Binding struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBinding that represents the Binding.
func (recv *Binding) ToC() unsafe.Pointer {
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

// ParamSpecVariant is a representation of the C record GParamSpecVariant.
//
// since 2.26
type ParamSpecVariant struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GParamSpecVariant that represents the ParamSpecVariant.
func (recv *ParamSpecVariant) ToC() unsafe.Pointer {
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
type TypeCValue struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTypeCValue that represents the TypeCValue.
func (recv *TypeCValue) ToC() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : _Value__data__union : blacklisted
