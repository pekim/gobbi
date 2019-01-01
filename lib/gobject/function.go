// This is a generated file - DO NOT EDIT

package gobject

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// BoxedCopy is a wrapper around the C function g_boxed_copy.
func BoxedCopy(boxedType Type, srcBoxed uintptr) uintptr {
	c_boxed_type := (C.GType)(boxedType)

	c_src_boxed := (C.gconstpointer)(srcBoxed)

	retC := C.g_boxed_copy(c_boxed_type, c_src_boxed)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// BoxedFree is a wrapper around the C function g_boxed_free.
func BoxedFree(boxedType Type, boxed uintptr) {
	c_boxed_type := (C.GType)(boxedType)

	c_boxed := (C.gpointer)(boxed)

	C.g_boxed_free(c_boxed_type, c_boxed)

	return
}

// Unsupported : g_boxed_type_register_static : unsupported parameter boxed_copy : no type generator for BoxedCopyFunc (GBoxedCopyFunc) for param boxed_copy

// Unsupported : g_cclosure_new : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_object : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_object_swap : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_swap : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// EnumCompleteTypeInfo is a wrapper around the C function g_enum_complete_type_info.
func EnumCompleteTypeInfo(gEnumType Type, constValues *EnumValue) *TypeInfo {
	c_g_enum_type := (C.GType)(gEnumType)

	var c_info C.GTypeInfo

	c_const_values := (*C.GEnumValue)(C.NULL)
	if constValues != nil {
		c_const_values = (*C.GEnumValue)(constValues.ToC())
	}

	C.g_enum_complete_type_info(c_g_enum_type, &c_info, c_const_values)

	info := TypeInfoNewFromC(unsafe.Pointer(&c_info))

	return info
}

// EnumGetValue is a wrapper around the C function g_enum_get_value.
func EnumGetValue(enumClass *EnumClass, value int32) *EnumValue {
	c_enum_class := (*C.GEnumClass)(C.NULL)
	if enumClass != nil {
		c_enum_class = (*C.GEnumClass)(enumClass.ToC())
	}

	c_value := (C.gint)(value)

	retC := C.g_enum_get_value(c_enum_class, c_value)
	retGo := EnumValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EnumGetValueByName is a wrapper around the C function g_enum_get_value_by_name.
func EnumGetValueByName(enumClass *EnumClass, name string) *EnumValue {
	c_enum_class := (*C.GEnumClass)(C.NULL)
	if enumClass != nil {
		c_enum_class = (*C.GEnumClass)(enumClass.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_enum_get_value_by_name(c_enum_class, c_name)
	retGo := EnumValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EnumGetValueByNick is a wrapper around the C function g_enum_get_value_by_nick.
func EnumGetValueByNick(enumClass *EnumClass, nick string) *EnumValue {
	c_enum_class := (*C.GEnumClass)(C.NULL)
	if enumClass != nil {
		c_enum_class = (*C.GEnumClass)(enumClass.ToC())
	}

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	retC := C.g_enum_get_value_by_nick(c_enum_class, c_nick)
	retGo := EnumValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EnumRegisterStatic is a wrapper around the C function g_enum_register_static.
func EnumRegisterStatic(name string, constStaticValues *EnumValue) Type {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_const_static_values := (*C.GEnumValue)(C.NULL)
	if constStaticValues != nil {
		c_const_static_values = (*C.GEnumValue)(constStaticValues.ToC())
	}

	retC := C.g_enum_register_static(c_name, c_const_static_values)
	retGo := (Type)(retC)

	return retGo
}

// FlagsCompleteTypeInfo is a wrapper around the C function g_flags_complete_type_info.
func FlagsCompleteTypeInfo(gFlagsType Type, constValues *FlagsValue) *TypeInfo {
	c_g_flags_type := (C.GType)(gFlagsType)

	var c_info C.GTypeInfo

	c_const_values := (*C.GFlagsValue)(C.NULL)
	if constValues != nil {
		c_const_values = (*C.GFlagsValue)(constValues.ToC())
	}

	C.g_flags_complete_type_info(c_g_flags_type, &c_info, c_const_values)

	info := TypeInfoNewFromC(unsafe.Pointer(&c_info))

	return info
}

// FlagsGetFirstValue is a wrapper around the C function g_flags_get_first_value.
func FlagsGetFirstValue(flagsClass *FlagsClass, value uint32) *FlagsValue {
	c_flags_class := (*C.GFlagsClass)(C.NULL)
	if flagsClass != nil {
		c_flags_class = (*C.GFlagsClass)(flagsClass.ToC())
	}

	c_value := (C.guint)(value)

	retC := C.g_flags_get_first_value(c_flags_class, c_value)
	retGo := FlagsValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FlagsGetValueByName is a wrapper around the C function g_flags_get_value_by_name.
func FlagsGetValueByName(flagsClass *FlagsClass, name string) *FlagsValue {
	c_flags_class := (*C.GFlagsClass)(C.NULL)
	if flagsClass != nil {
		c_flags_class = (*C.GFlagsClass)(flagsClass.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_flags_get_value_by_name(c_flags_class, c_name)
	retGo := FlagsValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FlagsGetValueByNick is a wrapper around the C function g_flags_get_value_by_nick.
func FlagsGetValueByNick(flagsClass *FlagsClass, nick string) *FlagsValue {
	c_flags_class := (*C.GFlagsClass)(C.NULL)
	if flagsClass != nil {
		c_flags_class = (*C.GFlagsClass)(flagsClass.ToC())
	}

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	retC := C.g_flags_get_value_by_nick(c_flags_class, c_nick)
	retGo := FlagsValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FlagsRegisterStatic is a wrapper around the C function g_flags_register_static.
func FlagsRegisterStatic(name string, constStaticValues *FlagsValue) Type {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_const_static_values := (*C.GFlagsValue)(C.NULL)
	if constStaticValues != nil {
		c_const_static_values = (*C.GFlagsValue)(constStaticValues.ToC())
	}

	retC := C.g_flags_register_static(c_name, c_const_static_values)
	retGo := (Type)(retC)

	return retGo
}

// GtypeGetType is a wrapper around the C function g_gtype_get_type.
func GtypeGetType() Type {
	retC := C.g_gtype_get_type()
	retGo := (Type)(retC)

	return retGo
}

// ParamSpecBoolean_ is a wrapper around the C function g_param_spec_boolean.
func ParamSpecBoolean_(name string, nick string, blurb string, defaultValue bool, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_default_value :=
		boolToGboolean(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_boolean(c_name, c_nick, c_blurb, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecBoxed_ is a wrapper around the C function g_param_spec_boxed.
func ParamSpecBoxed_(name string, nick string, blurb string, boxedType Type, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_boxed_type := (C.GType)(boxedType)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_boxed(c_name, c_nick, c_blurb, c_boxed_type, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecChar_ is a wrapper around the C function g_param_spec_char.
func ParamSpecChar_(name string, nick string, blurb string, minimum int8, maximum int8, defaultValue int8, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_minimum := (C.gint8)(minimum)

	c_maximum := (C.gint8)(maximum)

	c_default_value := (C.gint8)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_char(c_name, c_nick, c_blurb, c_minimum, c_maximum, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecDouble_ is a wrapper around the C function g_param_spec_double.
func ParamSpecDouble_(name string, nick string, blurb string, minimum float64, maximum float64, defaultValue float64, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_minimum := (C.gdouble)(minimum)

	c_maximum := (C.gdouble)(maximum)

	c_default_value := (C.gdouble)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_double(c_name, c_nick, c_blurb, c_minimum, c_maximum, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecEnum_ is a wrapper around the C function g_param_spec_enum.
func ParamSpecEnum_(name string, nick string, blurb string, enumType Type, defaultValue int32, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_enum_type := (C.GType)(enumType)

	c_default_value := (C.gint)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_enum(c_name, c_nick, c_blurb, c_enum_type, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecFlags_ is a wrapper around the C function g_param_spec_flags.
func ParamSpecFlags_(name string, nick string, blurb string, flagsType Type, defaultValue uint32, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_flags_type := (C.GType)(flagsType)

	c_default_value := (C.guint)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_flags(c_name, c_nick, c_blurb, c_flags_type, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecFloat_ is a wrapper around the C function g_param_spec_float.
func ParamSpecFloat_(name string, nick string, blurb string, minimum float32, maximum float32, defaultValue float32, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_minimum := (C.gfloat)(minimum)

	c_maximum := (C.gfloat)(maximum)

	c_default_value := (C.gfloat)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_float(c_name, c_nick, c_blurb, c_minimum, c_maximum, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecInt_ is a wrapper around the C function g_param_spec_int.
func ParamSpecInt_(name string, nick string, blurb string, minimum int32, maximum int32, defaultValue int32, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_minimum := (C.gint)(minimum)

	c_maximum := (C.gint)(maximum)

	c_default_value := (C.gint)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_int(c_name, c_nick, c_blurb, c_minimum, c_maximum, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecInt64_ is a wrapper around the C function g_param_spec_int64.
func ParamSpecInt64_(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_minimum := (C.gint64)(minimum)

	c_maximum := (C.gint64)(maximum)

	c_default_value := (C.gint64)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_int64(c_name, c_nick, c_blurb, c_minimum, c_maximum, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamLong_ is a wrapper around the C function g_param_spec_long.
func ParamLong_(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_minimum := (C.glong)(minimum)

	c_maximum := (C.glong)(maximum)

	c_default_value := (C.glong)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_long(c_name, c_nick, c_blurb, c_minimum, c_maximum, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecObject_ is a wrapper around the C function g_param_spec_object.
func ParamSpecObject_(name string, nick string, blurb string, objectType Type, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_object_type := (C.GType)(objectType)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_object(c_name, c_nick, c_blurb, c_object_type, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecParam_ is a wrapper around the C function g_param_spec_param.
func ParamSpecParam_(name string, nick string, blurb string, paramType Type, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_param_type := (C.GType)(paramType)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_param(c_name, c_nick, c_blurb, c_param_type, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecPointer_ is a wrapper around the C function g_param_spec_pointer.
func ParamSpecPointer_(name string, nick string, blurb string, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_pointer(c_name, c_nick, c_blurb, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecString_ is a wrapper around the C function g_param_spec_string.
func ParamSpecString_(name string, nick string, blurb string, defaultValue string, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_default_value := C.CString(defaultValue)
	defer C.free(unsafe.Pointer(c_default_value))

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_string(c_name, c_nick, c_blurb, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecUchar is a wrapper around the C function g_param_spec_uchar.
func ParamSpecUchar(name string, nick string, blurb string, minimum uint8, maximum uint8, defaultValue uint8, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_minimum := (C.guint8)(minimum)

	c_maximum := (C.guint8)(maximum)

	c_default_value := (C.guint8)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_uchar(c_name, c_nick, c_blurb, c_minimum, c_maximum, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecUint is a wrapper around the C function g_param_spec_uint.
func ParamSpecUint(name string, nick string, blurb string, minimum uint32, maximum uint32, defaultValue uint32, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_minimum := (C.guint)(minimum)

	c_maximum := (C.guint)(maximum)

	c_default_value := (C.guint)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_uint(c_name, c_nick, c_blurb, c_minimum, c_maximum, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecUint64 is a wrapper around the C function g_param_spec_uint64.
func ParamSpecUint64(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_minimum := (C.guint64)(minimum)

	c_maximum := (C.guint64)(maximum)

	c_default_value := (C.guint64)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_uint64(c_name, c_nick, c_blurb, c_minimum, c_maximum, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecUlong is a wrapper around the C function g_param_spec_ulong.
func ParamSpecUlong(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_minimum := (C.gulong)(minimum)

	c_maximum := (C.gulong)(maximum)

	c_default_value := (C.gulong)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_ulong(c_name, c_nick, c_blurb, c_minimum, c_maximum, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecUnichar_ is a wrapper around the C function g_param_spec_unichar.
func ParamSpecUnichar_(name string, nick string, blurb string, defaultValue rune, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_default_value := (C.gunichar)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_unichar(c_name, c_nick, c_blurb, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecValueArray_ is a wrapper around the C function g_param_spec_value_array.
func ParamSpecValueArray_(name string, nick string, blurb string, elementSpec *ParamSpec, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_element_spec := (*C.GParamSpec)(C.NULL)
	if elementSpec != nil {
		c_element_spec = (*C.GParamSpec)(elementSpec.ToC())
	}

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_value_array(c_name, c_nick, c_blurb, c_element_spec, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamTypeRegisterStatic is a wrapper around the C function g_param_type_register_static.
func ParamTypeRegisterStatic(name string, pspecInfo *ParamSpecTypeInfo) Type {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_pspec_info := (*C.GParamSpecTypeInfo)(C.NULL)
	if pspecInfo != nil {
		c_pspec_info = (*C.GParamSpecTypeInfo)(pspecInfo.ToC())
	}

	retC := C.g_param_type_register_static(c_name, c_pspec_info)
	retGo := (Type)(retC)

	return retGo
}

// ParamValueConvert is a wrapper around the C function g_param_value_convert.
func ParamValueConvert(pspec *ParamSpec, srcValue *Value, destValue *Value, strictValidation bool) bool {
	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	c_src_value := (*C.GValue)(C.NULL)
	if srcValue != nil {
		c_src_value = (*C.GValue)(srcValue.ToC())
	}

	c_dest_value := (*C.GValue)(C.NULL)
	if destValue != nil {
		c_dest_value = (*C.GValue)(destValue.ToC())
	}

	c_strict_validation :=
		boolToGboolean(strictValidation)

	retC := C.g_param_value_convert(c_pspec, c_src_value, c_dest_value, c_strict_validation)
	retGo := retC == C.TRUE

	return retGo
}

// ParamValueDefaults is a wrapper around the C function g_param_value_defaults.
func ParamValueDefaults(pspec *ParamSpec, value *Value) bool {
	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.g_param_value_defaults(c_pspec, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// ParamValueSetDefault is a wrapper around the C function g_param_value_set_default.
func ParamValueSetDefault(pspec *ParamSpec, value *Value) {
	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	C.g_param_value_set_default(c_pspec, c_value)

	return
}

// ParamValueValidate is a wrapper around the C function g_param_value_validate.
func ParamValueValidate(pspec *ParamSpec, value *Value) bool {
	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.g_param_value_validate(c_pspec, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// ParamValuesCmp is a wrapper around the C function g_param_values_cmp.
func ParamValuesCmp(pspec *ParamSpec, value1 *Value, value2 *Value) int32 {
	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	c_value1 := (*C.GValue)(C.NULL)
	if value1 != nil {
		c_value1 = (*C.GValue)(value1.ToC())
	}

	c_value2 := (*C.GValue)(C.NULL)
	if value2 != nil {
		c_value2 = (*C.GValue)(value2.ToC())
	}

	retC := C.g_param_values_cmp(c_pspec, c_value1, c_value2)
	retGo := (int32)(retC)

	return retGo
}

// PointerTypeRegisterStatic is a wrapper around the C function g_pointer_type_register_static.
func PointerTypeRegisterStatic(name string) Type {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_pointer_type_register_static(c_name)
	retGo := (Type)(retC)

	return retGo
}

// Unsupported : g_signal_add_emission_hook : unsupported parameter hook_func : no type generator for SignalEmissionHook (GSignalEmissionHook) for param hook_func

// Unsupported : g_signal_chain_from_overridden : unsupported parameter instance_and_params :

// SignalConnectClosure is a wrapper around the C function g_signal_connect_closure.
func SignalConnectClosure(instance uintptr, detailedSignal string, closure *Closure, after bool) uint64 {
	c_instance := (C.gpointer)(instance)

	c_detailed_signal := C.CString(detailedSignal)
	defer C.free(unsafe.Pointer(c_detailed_signal))

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_after :=
		boolToGboolean(after)

	retC := C.g_signal_connect_closure(c_instance, c_detailed_signal, c_closure, c_after)
	retGo := (uint64)(retC)

	return retGo
}

// SignalConnectClosureById is a wrapper around the C function g_signal_connect_closure_by_id.
func SignalConnectClosureById(instance uintptr, signalId uint32, detail glib.Quark, closure *Closure, after bool) uint64 {
	c_instance := (C.gpointer)(instance)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_after :=
		boolToGboolean(after)

	retC := C.g_signal_connect_closure_by_id(c_instance, c_signal_id, c_detail, c_closure, c_after)
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported : g_signal_connect_data : unsupported parameter c_handler : no type generator for Callback (GCallback) for param c_handler

// Unsupported : g_signal_connect_object : unsupported parameter c_handler : no type generator for Callback (GCallback) for param c_handler

// Unsupported : g_signal_emit : unsupported parameter ... : varargs

// Unsupported : g_signal_emit_by_name : unsupported parameter ... : varargs

// Unsupported : g_signal_emit_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Unsupported : g_signal_emitv : unsupported parameter instance_and_params :

// SignalGetInvocationHint is a wrapper around the C function g_signal_get_invocation_hint.
func SignalGetInvocationHint(instance uintptr) *SignalInvocationHint {
	c_instance := (C.gpointer)(instance)

	retC := C.g_signal_get_invocation_hint(c_instance)
	retGo := SignalInvocationHintNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SignalHandlerBlock is a wrapper around the C function g_signal_handler_block.
func SignalHandlerBlock(instance uintptr, handlerId uint64) {
	c_instance := (C.gpointer)(instance)

	c_handler_id := (C.gulong)(handlerId)

	C.g_signal_handler_block(c_instance, c_handler_id)

	return
}

// SignalHandlerDisconnect is a wrapper around the C function g_signal_handler_disconnect.
func SignalHandlerDisconnect(instance uintptr, handlerId uint64) {
	c_instance := (C.gpointer)(instance)

	c_handler_id := (C.gulong)(handlerId)

	C.g_signal_handler_disconnect(c_instance, c_handler_id)

	return
}

// SignalHandlerFind is a wrapper around the C function g_signal_handler_find.
func SignalHandlerFind(instance uintptr, mask SignalMatchType, signalId uint32, detail glib.Quark, closure *Closure, func_ uintptr, data uintptr) uint64 {
	c_instance := (C.gpointer)(instance)

	c_mask := (C.GSignalMatchType)(mask)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_func := (C.gpointer)(func_)

	c_data := (C.gpointer)(data)

	retC := C.g_signal_handler_find(c_instance, c_mask, c_signal_id, c_detail, c_closure, c_func, c_data)
	retGo := (uint64)(retC)

	return retGo
}

// SignalHandlerIsConnected is a wrapper around the C function g_signal_handler_is_connected.
func SignalHandlerIsConnected(instance uintptr, handlerId uint64) bool {
	c_instance := (C.gpointer)(instance)

	c_handler_id := (C.gulong)(handlerId)

	retC := C.g_signal_handler_is_connected(c_instance, c_handler_id)
	retGo := retC == C.TRUE

	return retGo
}

// SignalHandlerUnblock is a wrapper around the C function g_signal_handler_unblock.
func SignalHandlerUnblock(instance uintptr, handlerId uint64) {
	c_instance := (C.gpointer)(instance)

	c_handler_id := (C.gulong)(handlerId)

	C.g_signal_handler_unblock(c_instance, c_handler_id)

	return
}

// SignalHandlersBlockMatched is a wrapper around the C function g_signal_handlers_block_matched.
func SignalHandlersBlockMatched(instance uintptr, mask SignalMatchType, signalId uint32, detail glib.Quark, closure *Closure, func_ uintptr, data uintptr) uint32 {
	c_instance := (C.gpointer)(instance)

	c_mask := (C.GSignalMatchType)(mask)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_func := (C.gpointer)(func_)

	c_data := (C.gpointer)(data)

	retC := C.g_signal_handlers_block_matched(c_instance, c_mask, c_signal_id, c_detail, c_closure, c_func, c_data)
	retGo := (uint32)(retC)

	return retGo
}

// SignalHandlersDestroy is a wrapper around the C function g_signal_handlers_destroy.
func SignalHandlersDestroy(instance uintptr) {
	c_instance := (C.gpointer)(instance)

	C.g_signal_handlers_destroy(c_instance)

	return
}

// SignalHandlersDisconnectMatched is a wrapper around the C function g_signal_handlers_disconnect_matched.
func SignalHandlersDisconnectMatched(instance uintptr, mask SignalMatchType, signalId uint32, detail glib.Quark, closure *Closure, func_ uintptr, data uintptr) uint32 {
	c_instance := (C.gpointer)(instance)

	c_mask := (C.GSignalMatchType)(mask)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_func := (C.gpointer)(func_)

	c_data := (C.gpointer)(data)

	retC := C.g_signal_handlers_disconnect_matched(c_instance, c_mask, c_signal_id, c_detail, c_closure, c_func, c_data)
	retGo := (uint32)(retC)

	return retGo
}

// SignalHandlersUnblockMatched is a wrapper around the C function g_signal_handlers_unblock_matched.
func SignalHandlersUnblockMatched(instance uintptr, mask SignalMatchType, signalId uint32, detail glib.Quark, closure *Closure, func_ uintptr, data uintptr) uint32 {
	c_instance := (C.gpointer)(instance)

	c_mask := (C.GSignalMatchType)(mask)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_func := (C.gpointer)(func_)

	c_data := (C.gpointer)(data)

	retC := C.g_signal_handlers_unblock_matched(c_instance, c_mask, c_signal_id, c_detail, c_closure, c_func, c_data)
	retGo := (uint32)(retC)

	return retGo
}

// SignalHasHandlerPending is a wrapper around the C function g_signal_has_handler_pending.
func SignalHasHandlerPending(instance uintptr, signalId uint32, detail glib.Quark, mayBeBlocked bool) bool {
	c_instance := (C.gpointer)(instance)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_may_be_blocked :=
		boolToGboolean(mayBeBlocked)

	retC := C.g_signal_has_handler_pending(c_instance, c_signal_id, c_detail, c_may_be_blocked)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_signal_list_ids : array return type :

// SignalLookup is a wrapper around the C function g_signal_lookup.
func SignalLookup(name string, itype Type) uint32 {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_itype := (C.GType)(itype)

	retC := C.g_signal_lookup(c_name, c_itype)
	retGo := (uint32)(retC)

	return retGo
}

// SignalName is a wrapper around the C function g_signal_name.
func SignalName(signalId uint32) string {
	c_signal_id := (C.guint)(signalId)

	retC := C.g_signal_name(c_signal_id)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_signal_new : unsupported parameter accumulator : no type generator for SignalAccumulator (GSignalAccumulator) for param accumulator

// Unsupported : g_signal_new_valist : unsupported parameter accumulator : no type generator for SignalAccumulator (GSignalAccumulator) for param accumulator

// Unsupported : g_signal_newv : unsupported parameter accumulator : no type generator for SignalAccumulator (GSignalAccumulator) for param accumulator

// SignalOverrideClassClosure is a wrapper around the C function g_signal_override_class_closure.
func SignalOverrideClassClosure(signalId uint32, instanceType Type, classClosure *Closure) {
	c_signal_id := (C.guint)(signalId)

	c_instance_type := (C.GType)(instanceType)

	c_class_closure := (*C.GClosure)(C.NULL)
	if classClosure != nil {
		c_class_closure = (*C.GClosure)(classClosure.ToC())
	}

	C.g_signal_override_class_closure(c_signal_id, c_instance_type, c_class_closure)

	return
}

// SignalParseName is a wrapper around the C function g_signal_parse_name.
func SignalParseName(detailedSignal string, itype Type, forceDetailQuark bool) (bool, uint32, glib.Quark) {
	c_detailed_signal := C.CString(detailedSignal)
	defer C.free(unsafe.Pointer(c_detailed_signal))

	c_itype := (C.GType)(itype)

	var c_signal_id_p C.guint

	var c_detail_p C.GQuark

	c_force_detail_quark :=
		boolToGboolean(forceDetailQuark)

	retC := C.g_signal_parse_name(c_detailed_signal, c_itype, &c_signal_id_p, &c_detail_p, c_force_detail_quark)
	retGo := retC == C.TRUE

	signalIdP := (uint32)(c_signal_id_p)

	detailP := (glib.Quark)(c_detail_p)

	return retGo, signalIdP, detailP
}

// SignalQuery_ is a wrapper around the C function g_signal_query.
func SignalQuery_(signalId uint32) *SignalQuery {
	c_signal_id := (C.guint)(signalId)

	var c_query C.GSignalQuery

	C.g_signal_query(c_signal_id, &c_query)

	query := SignalQueryNewFromC(unsafe.Pointer(&c_query))

	return query
}

// SignalRemoveEmissionHook is a wrapper around the C function g_signal_remove_emission_hook.
func SignalRemoveEmissionHook(signalId uint32, hookId uint64) {
	c_signal_id := (C.guint)(signalId)

	c_hook_id := (C.gulong)(hookId)

	C.g_signal_remove_emission_hook(c_signal_id, c_hook_id)

	return
}

// SignalStopEmission is a wrapper around the C function g_signal_stop_emission.
func SignalStopEmission(instance uintptr, signalId uint32, detail glib.Quark) {
	c_instance := (C.gpointer)(instance)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	C.g_signal_stop_emission(c_instance, c_signal_id, c_detail)

	return
}

// SignalStopEmissionByName is a wrapper around the C function g_signal_stop_emission_by_name.
func SignalStopEmissionByName(instance uintptr, detailedSignal string) {
	c_instance := (C.gpointer)(instance)

	c_detailed_signal := C.CString(detailedSignal)
	defer C.free(unsafe.Pointer(c_detailed_signal))

	C.g_signal_stop_emission_by_name(c_instance, c_detailed_signal)

	return
}

// SignalTypeCclosureNew is a wrapper around the C function g_signal_type_cclosure_new.
func SignalTypeCclosureNew(itype Type, structOffset uint32) *Closure {
	c_itype := (C.GType)(itype)

	c_struct_offset := (C.guint)(structOffset)

	retC := C.g_signal_type_cclosure_new(c_itype, c_struct_offset)
	retGo := ClosureNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SourceSetClosure is a wrapper around the C function g_source_set_closure.
func SourceSetClosure(source *glib.Source, closure *Closure) {
	c_source := (*C.GSource)(C.NULL)
	if source != nil {
		c_source = (*C.GSource)(source.ToC())
	}

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	C.g_source_set_closure(c_source, c_closure)

	return
}

// SourceSetDummyCallback is a wrapper around the C function g_source_set_dummy_callback.
func SourceSetDummyCallback(source *glib.Source) {
	c_source := (*C.GSource)(C.NULL)
	if source != nil {
		c_source = (*C.GSource)(source.ToC())
	}

	C.g_source_set_dummy_callback(c_source)

	return
}

// StrdupValueContents is a wrapper around the C function g_strdup_value_contents.
func StrdupValueContents(value *Value) string {
	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.g_strdup_value_contents(c_value)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_type_add_class_cache_func : unsupported parameter cache_func : no type generator for TypeClassCacheFunc (GTypeClassCacheFunc) for param cache_func

// TypeAddInstancePrivate is a wrapper around the C function g_type_add_instance_private.
func TypeAddInstancePrivate(classType Type, privateSize uint64) int32 {
	c_class_type := (C.GType)(classType)

	c_private_size := (C.gsize)(privateSize)

	retC := C.g_type_add_instance_private(c_class_type, c_private_size)
	retGo := (int32)(retC)

	return retGo
}

// TypeAddInterfaceDynamic is a wrapper around the C function g_type_add_interface_dynamic.
func TypeAddInterfaceDynamic(instanceType Type, interfaceType Type, plugin *TypePlugin) {
	c_instance_type := (C.GType)(instanceType)

	c_interface_type := (C.GType)(interfaceType)

	c_plugin := (*C.GTypePlugin)(plugin.ToC())

	C.g_type_add_interface_dynamic(c_instance_type, c_interface_type, c_plugin)

	return
}

// TypeAddInterfaceStatic is a wrapper around the C function g_type_add_interface_static.
func TypeAddInterfaceStatic(instanceType Type, interfaceType Type, info *InterfaceInfo) {
	c_instance_type := (C.GType)(instanceType)

	c_interface_type := (C.GType)(interfaceType)

	c_info := (*C.GInterfaceInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GInterfaceInfo)(info.ToC())
	}

	C.g_type_add_interface_static(c_instance_type, c_interface_type, c_info)

	return
}

// TypeCheckClassCast is a wrapper around the C function g_type_check_class_cast.
func TypeCheckClassCast(gClass *TypeClass, isAType Type) *TypeClass {
	c_g_class := (*C.GTypeClass)(C.NULL)
	if gClass != nil {
		c_g_class = (*C.GTypeClass)(gClass.ToC())
	}

	c_is_a_type := (C.GType)(isAType)

	retC := C.g_type_check_class_cast(c_g_class, c_is_a_type)
	retGo := TypeClassNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeCheckClassIsA is a wrapper around the C function g_type_check_class_is_a.
func TypeCheckClassIsA(gClass *TypeClass, isAType Type) bool {
	c_g_class := (*C.GTypeClass)(C.NULL)
	if gClass != nil {
		c_g_class = (*C.GTypeClass)(gClass.ToC())
	}

	c_is_a_type := (C.GType)(isAType)

	retC := C.g_type_check_class_is_a(c_g_class, c_is_a_type)
	retGo := retC == C.TRUE

	return retGo
}

// TypeCheckInstance is a wrapper around the C function g_type_check_instance.
func TypeCheckInstance(instance *TypeInstance) bool {
	c_instance := (*C.GTypeInstance)(C.NULL)
	if instance != nil {
		c_instance = (*C.GTypeInstance)(instance.ToC())
	}

	retC := C.g_type_check_instance(c_instance)
	retGo := retC == C.TRUE

	return retGo
}

// TypeCheckInstanceCast is a wrapper around the C function g_type_check_instance_cast.
func TypeCheckInstanceCast(instance *TypeInstance, ifaceType Type) *TypeInstance {
	c_instance := (*C.GTypeInstance)(C.NULL)
	if instance != nil {
		c_instance = (*C.GTypeInstance)(instance.ToC())
	}

	c_iface_type := (C.GType)(ifaceType)

	retC := C.g_type_check_instance_cast(c_instance, c_iface_type)
	retGo := TypeInstanceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeCheckInstanceIsA is a wrapper around the C function g_type_check_instance_is_a.
func TypeCheckInstanceIsA(instance *TypeInstance, ifaceType Type) bool {
	c_instance := (*C.GTypeInstance)(C.NULL)
	if instance != nil {
		c_instance = (*C.GTypeInstance)(instance.ToC())
	}

	c_iface_type := (C.GType)(ifaceType)

	retC := C.g_type_check_instance_is_a(c_instance, c_iface_type)
	retGo := retC == C.TRUE

	return retGo
}

// TypeCheckInstanceIsFundamentallyA is a wrapper around the C function g_type_check_instance_is_fundamentally_a.
func TypeCheckInstanceIsFundamentallyA(instance *TypeInstance, fundamentalType Type) bool {
	c_instance := (*C.GTypeInstance)(C.NULL)
	if instance != nil {
		c_instance = (*C.GTypeInstance)(instance.ToC())
	}

	c_fundamental_type := (C.GType)(fundamentalType)

	retC := C.g_type_check_instance_is_fundamentally_a(c_instance, c_fundamental_type)
	retGo := retC == C.TRUE

	return retGo
}

// TypeCheckIsValueType is a wrapper around the C function g_type_check_is_value_type.
func TypeCheckIsValueType(type_ Type) bool {
	c_type := (C.GType)(type_)

	retC := C.g_type_check_is_value_type(c_type)
	retGo := retC == C.TRUE

	return retGo
}

// TypeCheckValue is a wrapper around the C function g_type_check_value.
func TypeCheckValue(value *Value) bool {
	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.g_type_check_value(c_value)
	retGo := retC == C.TRUE

	return retGo
}

// TypeCheckValueHolds is a wrapper around the C function g_type_check_value_holds.
func TypeCheckValueHolds(value *Value, type_ Type) bool {
	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	c_type := (C.GType)(type_)

	retC := C.g_type_check_value_holds(c_value, c_type)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_type_children : array return type :

// TypeCreateInstance is a wrapper around the C function g_type_create_instance.
func TypeCreateInstance(type_ Type) *TypeInstance {
	c_type := (C.GType)(type_)

	retC := C.g_type_create_instance(c_type)
	retGo := TypeInstanceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeDepth is a wrapper around the C function g_type_depth.
func TypeDepth(type_ Type) uint32 {
	c_type := (C.GType)(type_)

	retC := C.g_type_depth(c_type)
	retGo := (uint32)(retC)

	return retGo
}

// TypeFreeInstance is a wrapper around the C function g_type_free_instance.
func TypeFreeInstance(instance *TypeInstance) {
	c_instance := (*C.GTypeInstance)(C.NULL)
	if instance != nil {
		c_instance = (*C.GTypeInstance)(instance.ToC())
	}

	C.g_type_free_instance(c_instance)

	return
}

// TypeFromName is a wrapper around the C function g_type_from_name.
func TypeFromName(name string) Type {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_type_from_name(c_name)
	retGo := (Type)(retC)

	return retGo
}

// TypeFundamental is a wrapper around the C function g_type_fundamental.
func TypeFundamental(typeId Type) Type {
	c_type_id := (C.GType)(typeId)

	retC := C.g_type_fundamental(c_type_id)
	retGo := (Type)(retC)

	return retGo
}

// TypeFundamentalNext is a wrapper around the C function g_type_fundamental_next.
func TypeFundamentalNext() Type {
	retC := C.g_type_fundamental_next()
	retGo := (Type)(retC)

	return retGo
}

// TypeGetPlugin is a wrapper around the C function g_type_get_plugin.
func TypeGetPlugin(type_ Type) *TypePlugin {
	c_type := (C.GType)(type_)

	retC := C.g_type_get_plugin(c_type)
	retGo := TypePluginNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeGetQdata is a wrapper around the C function g_type_get_qdata.
func TypeGetQdata(type_ Type, quark glib.Quark) uintptr {
	c_type := (C.GType)(type_)

	c_quark := (C.GQuark)(quark)

	retC := C.g_type_get_qdata(c_type, c_quark)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// TypeInit is a wrapper around the C function g_type_init.
func TypeInit() {
	C.g_type_init()

	return
}

// TypeInitWithDebugFlags is a wrapper around the C function g_type_init_with_debug_flags.
func TypeInitWithDebugFlags(debugFlags TypeDebugFlags) {
	c_debug_flags := (C.GTypeDebugFlags)(debugFlags)

	C.g_type_init_with_debug_flags(c_debug_flags)

	return
}

// Unsupported : g_type_interfaces : array return type :

// TypeIsA is a wrapper around the C function g_type_is_a.
func TypeIsA(type_ Type, isAType Type) bool {
	c_type := (C.GType)(type_)

	c_is_a_type := (C.GType)(isAType)

	retC := C.g_type_is_a(c_type, c_is_a_type)
	retGo := retC == C.TRUE

	return retGo
}

// TypeName is a wrapper around the C function g_type_name.
func TypeName(type_ Type) string {
	c_type := (C.GType)(type_)

	retC := C.g_type_name(c_type)
	retGo := C.GoString(retC)

	return retGo
}

// TypeNameFromClass is a wrapper around the C function g_type_name_from_class.
func TypeNameFromClass(gClass *TypeClass) string {
	c_g_class := (*C.GTypeClass)(C.NULL)
	if gClass != nil {
		c_g_class = (*C.GTypeClass)(gClass.ToC())
	}

	retC := C.g_type_name_from_class(c_g_class)
	retGo := C.GoString(retC)

	return retGo
}

// TypeNameFromInstance is a wrapper around the C function g_type_name_from_instance.
func TypeNameFromInstance(instance *TypeInstance) string {
	c_instance := (*C.GTypeInstance)(C.NULL)
	if instance != nil {
		c_instance = (*C.GTypeInstance)(instance.ToC())
	}

	retC := C.g_type_name_from_instance(c_instance)
	retGo := C.GoString(retC)

	return retGo
}

// TypeNextBase is a wrapper around the C function g_type_next_base.
func TypeNextBase(leafType Type, rootType Type) Type {
	c_leaf_type := (C.GType)(leafType)

	c_root_type := (C.GType)(rootType)

	retC := C.g_type_next_base(c_leaf_type, c_root_type)
	retGo := (Type)(retC)

	return retGo
}

// TypeParent is a wrapper around the C function g_type_parent.
func TypeParent(type_ Type) Type {
	c_type := (C.GType)(type_)

	retC := C.g_type_parent(c_type)
	retGo := (Type)(retC)

	return retGo
}

// TypeQname is a wrapper around the C function g_type_qname.
func TypeQname(type_ Type) glib.Quark {
	c_type := (C.GType)(type_)

	retC := C.g_type_qname(c_type)
	retGo := (glib.Quark)(retC)

	return retGo
}

// TypeQuery_ is a wrapper around the C function g_type_query.
func TypeQuery_(type_ Type) *TypeQuery {
	c_type := (C.GType)(type_)

	var c_query C.GTypeQuery

	C.g_type_query(c_type, &c_query)

	query := TypeQueryNewFromC(unsafe.Pointer(&c_query))

	return query
}

// TypeRegisterDynamic is a wrapper around the C function g_type_register_dynamic.
func TypeRegisterDynamic(parentType Type, typeName string, plugin *TypePlugin, flags TypeFlags) Type {
	c_parent_type := (C.GType)(parentType)

	c_type_name := C.CString(typeName)
	defer C.free(unsafe.Pointer(c_type_name))

	c_plugin := (*C.GTypePlugin)(plugin.ToC())

	c_flags := (C.GTypeFlags)(flags)

	retC := C.g_type_register_dynamic(c_parent_type, c_type_name, c_plugin, c_flags)
	retGo := (Type)(retC)

	return retGo
}

// TypeRegisterFundamental is a wrapper around the C function g_type_register_fundamental.
func TypeRegisterFundamental(typeId Type, typeName string, info *TypeInfo, finfo *TypeFundamentalInfo, flags TypeFlags) Type {
	c_type_id := (C.GType)(typeId)

	c_type_name := C.CString(typeName)
	defer C.free(unsafe.Pointer(c_type_name))

	c_info := (*C.GTypeInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GTypeInfo)(info.ToC())
	}

	c_finfo := (*C.GTypeFundamentalInfo)(C.NULL)
	if finfo != nil {
		c_finfo = (*C.GTypeFundamentalInfo)(finfo.ToC())
	}

	c_flags := (C.GTypeFlags)(flags)

	retC := C.g_type_register_fundamental(c_type_id, c_type_name, c_info, c_finfo, c_flags)
	retGo := (Type)(retC)

	return retGo
}

// TypeRegisterStatic is a wrapper around the C function g_type_register_static.
func TypeRegisterStatic(parentType Type, typeName string, info *TypeInfo, flags TypeFlags) Type {
	c_parent_type := (C.GType)(parentType)

	c_type_name := C.CString(typeName)
	defer C.free(unsafe.Pointer(c_type_name))

	c_info := (*C.GTypeInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GTypeInfo)(info.ToC())
	}

	c_flags := (C.GTypeFlags)(flags)

	retC := C.g_type_register_static(c_parent_type, c_type_name, c_info, c_flags)
	retGo := (Type)(retC)

	return retGo
}

// Unsupported : g_type_remove_class_cache_func : unsupported parameter cache_func : no type generator for TypeClassCacheFunc (GTypeClassCacheFunc) for param cache_func

// TypeSetQdata is a wrapper around the C function g_type_set_qdata.
func TypeSetQdata(type_ Type, quark glib.Quark, data uintptr) {
	c_type := (C.GType)(type_)

	c_quark := (C.GQuark)(quark)

	c_data := (C.gpointer)(data)

	C.g_type_set_qdata(c_type, c_quark, c_data)

	return
}

// TypeTestFlags is a wrapper around the C function g_type_test_flags.
func TypeTestFlags(type_ Type, flags uint32) bool {
	c_type := (C.GType)(type_)

	c_flags := (C.guint)(flags)

	retC := C.g_type_test_flags(c_type, c_flags)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_value_register_transform_func : unsupported parameter transform_func : no type generator for ValueTransform (GValueTransform) for param transform_func
