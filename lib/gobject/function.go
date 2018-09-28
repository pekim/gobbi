// This is a generated file - DO NOT EDIT

package gobject

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_boxed_copy : unsupported parameter boxed_type : no type generator for GType, GType

// Unsupported : g_boxed_free : unsupported parameter boxed_type : no type generator for GType, GType

// Unsupported : g_boxed_type_register_static : unsupported parameter boxed_copy : no type generator for BoxedCopyFunc, GBoxedCopyFunc

// CclosureMarshalBooleanBoxedBoxed is a wrapper around the C function g_cclosure_marshal_BOOLEAN__BOXED_BOXED.
func CclosureMarshalBooleanBoxedBoxed(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_BOOLEAN__BOXED_BOXED(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalBooleanFlags is a wrapper around the C function g_cclosure_marshal_BOOLEAN__FLAGS.
func CclosureMarshalBooleanFlags(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_BOOLEAN__FLAGS(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalStringObjectPointer is a wrapper around the C function g_cclosure_marshal_STRING__OBJECT_POINTER.
func CclosureMarshalStringObjectPointer(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_STRING__OBJECT_POINTER(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalVoidBoolean is a wrapper around the C function g_cclosure_marshal_VOID__BOOLEAN.
func CclosureMarshalVoidBoolean(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__BOOLEAN(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalVoidBoxed is a wrapper around the C function g_cclosure_marshal_VOID__BOXED.
func CclosureMarshalVoidBoxed(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__BOXED(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalVoidChar is a wrapper around the C function g_cclosure_marshal_VOID__CHAR.
func CclosureMarshalVoidChar(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__CHAR(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalVoidDouble is a wrapper around the C function g_cclosure_marshal_VOID__DOUBLE.
func CclosureMarshalVoidDouble(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__DOUBLE(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalVoidEnum is a wrapper around the C function g_cclosure_marshal_VOID__ENUM.
func CclosureMarshalVoidEnum(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__ENUM(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalVoidFlags is a wrapper around the C function g_cclosure_marshal_VOID__FLAGS.
func CclosureMarshalVoidFlags(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__FLAGS(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalVoidFloat is a wrapper around the C function g_cclosure_marshal_VOID__FLOAT.
func CclosureMarshalVoidFloat(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__FLOAT(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalVoidInt is a wrapper around the C function g_cclosure_marshal_VOID__INT.
func CclosureMarshalVoidInt(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__INT(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalVoidLong is a wrapper around the C function g_cclosure_marshal_VOID__LONG.
func CclosureMarshalVoidLong(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__LONG(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalVoidObject is a wrapper around the C function g_cclosure_marshal_VOID__OBJECT.
func CclosureMarshalVoidObject(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__OBJECT(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalVoidParam is a wrapper around the C function g_cclosure_marshal_VOID__PARAM.
func CclosureMarshalVoidParam(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__PARAM(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalVoidPointer is a wrapper around the C function g_cclosure_marshal_VOID__POINTER.
func CclosureMarshalVoidPointer(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__POINTER(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalVoidString is a wrapper around the C function g_cclosure_marshal_VOID__STRING.
func CclosureMarshalVoidString(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__STRING(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalVoidUchar is a wrapper around the C function g_cclosure_marshal_VOID__UCHAR.
func CclosureMarshalVoidUchar(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__UCHAR(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalVoidUint is a wrapper around the C function g_cclosure_marshal_VOID__UINT.
func CclosureMarshalVoidUint(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__UINT(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalVoidUintPointer is a wrapper around the C function g_cclosure_marshal_VOID__UINT_POINTER.
func CclosureMarshalVoidUintPointer(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__UINT_POINTER(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalVoidUlong is a wrapper around the C function g_cclosure_marshal_VOID__ULONG.
func CclosureMarshalVoidUlong(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__ULONG(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalVoidVariant is a wrapper around the C function g_cclosure_marshal_VOID__VARIANT.
func CclosureMarshalVoidVariant(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__VARIANT(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CclosureMarshalVoidVoid is a wrapper around the C function g_cclosure_marshal_VOID__VOID.
func CclosureMarshalVoidVoid(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_value := (*C.GValue)(returnValue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__VOID(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// Unsupported : g_cclosure_new : unsupported parameter callback_func : no type generator for Callback, GCallback

// Unsupported : g_cclosure_new_object : unsupported parameter callback_func : no type generator for Callback, GCallback

// Unsupported : g_cclosure_new_object_swap : unsupported parameter callback_func : no type generator for Callback, GCallback

// Unsupported : g_cclosure_new_swap : unsupported parameter callback_func : no type generator for Callback, GCallback

// Unsupported : g_enum_complete_type_info : unsupported parameter g_enum_type : no type generator for GType, GType

// EnumGetValue is a wrapper around the C function g_enum_get_value.
func EnumGetValue(enumClass *EnumClass, value int32) *EnumValue {
	c_enum_class := (*C.GEnumClass)(enumClass.ToC())

	c_value := (C.gint)(value)

	retC := C.g_enum_get_value(c_enum_class, c_value)
	retGo := EnumValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EnumGetValueByName is a wrapper around the C function g_enum_get_value_by_name.
func EnumGetValueByName(enumClass *EnumClass, name string) *EnumValue {
	c_enum_class := (*C.GEnumClass)(enumClass.ToC())

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_enum_get_value_by_name(c_enum_class, c_name)
	retGo := EnumValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EnumGetValueByNick is a wrapper around the C function g_enum_get_value_by_nick.
func EnumGetValueByNick(enumClass *EnumClass, nick string) *EnumValue {
	c_enum_class := (*C.GEnumClass)(enumClass.ToC())

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	retC := C.g_enum_get_value_by_nick(c_enum_class, c_nick)
	retGo := EnumValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_enum_register_static : no return generator

// Unsupported : g_flags_complete_type_info : unsupported parameter g_flags_type : no type generator for GType, GType

// FlagsGetFirstValue is a wrapper around the C function g_flags_get_first_value.
func FlagsGetFirstValue(flagsClass *FlagsClass, value uint32) *FlagsValue {
	c_flags_class := (*C.GFlagsClass)(flagsClass.ToC())

	c_value := (C.guint)(value)

	retC := C.g_flags_get_first_value(c_flags_class, c_value)
	retGo := FlagsValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FlagsGetValueByName is a wrapper around the C function g_flags_get_value_by_name.
func FlagsGetValueByName(flagsClass *FlagsClass, name string) *FlagsValue {
	c_flags_class := (*C.GFlagsClass)(flagsClass.ToC())

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_flags_get_value_by_name(c_flags_class, c_name)
	retGo := FlagsValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FlagsGetValueByNick is a wrapper around the C function g_flags_get_value_by_nick.
func FlagsGetValueByNick(flagsClass *FlagsClass, nick string) *FlagsValue {
	c_flags_class := (*C.GFlagsClass)(flagsClass.ToC())

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	retC := C.g_flags_get_value_by_nick(c_flags_class, c_nick)
	retGo := FlagsValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_flags_register_static : no return generator

// Unsupported : g_gtype_get_type : no return generator

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

// Unsupported : g_param_spec_boxed : unsupported parameter boxed_type : no type generator for GType, GType

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

// Unsupported : g_param_spec_enum : unsupported parameter enum_type : no type generator for GType, GType

// Unsupported : g_param_spec_flags : unsupported parameter flags_type : no type generator for GType, GType

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

// Unsupported : g_param_spec_object : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_param_spec_param : unsupported parameter param_type : no type generator for GType, GType

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

// ParamSpecPoolNew is a wrapper around the C function g_param_spec_pool_new.
func ParamSpecPoolNew(typePrefixing bool) *ParamSpecPool {
	c_type_prefixing :=
		boolToGboolean(typePrefixing)

	retC := C.g_param_spec_pool_new(c_type_prefixing)
	retGo := ParamSpecPoolNewFromC(unsafe.Pointer(retC))

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

	c_element_spec := (*C.GParamSpec)(elementSpec.ToC())

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_value_array(c_name, c_nick, c_blurb, c_element_spec, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_param_type_register_static : no return generator

// ParamValueConvert is a wrapper around the C function g_param_value_convert.
func ParamValueConvert(pspec *ParamSpec, srcValue *Value, destValue *Value, strictValidation bool) bool {
	c_pspec := (*C.GParamSpec)(pspec.ToC())

	c_src_value := (*C.GValue)(srcValue.ToC())

	c_dest_value := (*C.GValue)(destValue.ToC())

	c_strict_validation :=
		boolToGboolean(strictValidation)

	retC := C.g_param_value_convert(c_pspec, c_src_value, c_dest_value, c_strict_validation)
	retGo := retC == C.TRUE

	return retGo
}

// ParamValueDefaults is a wrapper around the C function g_param_value_defaults.
func ParamValueDefaults(pspec *ParamSpec, value *Value) bool {
	c_pspec := (*C.GParamSpec)(pspec.ToC())

	c_value := (*C.GValue)(value.ToC())

	retC := C.g_param_value_defaults(c_pspec, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// ParamValueSetDefault is a wrapper around the C function g_param_value_set_default.
func ParamValueSetDefault(pspec *ParamSpec, value *Value) {
	c_pspec := (*C.GParamSpec)(pspec.ToC())

	c_value := (*C.GValue)(value.ToC())

	C.g_param_value_set_default(c_pspec, c_value)

	return
}

// ParamValueValidate is a wrapper around the C function g_param_value_validate.
func ParamValueValidate(pspec *ParamSpec, value *Value) bool {
	c_pspec := (*C.GParamSpec)(pspec.ToC())

	c_value := (*C.GValue)(value.ToC())

	retC := C.g_param_value_validate(c_pspec, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// ParamValuesCmp is a wrapper around the C function g_param_values_cmp.
func ParamValuesCmp(pspec *ParamSpec, value1 *Value, value2 *Value) int32 {
	c_pspec := (*C.GParamSpec)(pspec.ToC())

	c_value1 := (*C.GValue)(value1.ToC())

	c_value2 := (*C.GValue)(value2.ToC())

	retC := C.g_param_values_cmp(c_pspec, c_value1, c_value2)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_pointer_type_register_static : no return generator

// Unsupported : g_signal_add_emission_hook : unsupported parameter hook_func : no type generator for SignalEmissionHook, GSignalEmissionHook

// Unsupported : g_signal_chain_from_overridden : unsupported parameter instance_and_params : no param type

// SignalConnectClosure is a wrapper around the C function g_signal_connect_closure.
func SignalConnectClosure(instance uintptr, detailedSignal string, closure *Closure, after bool) uint64 {
	c_instance := (C.gpointer)(instance)

	c_detailed_signal := C.CString(detailedSignal)
	defer C.free(unsafe.Pointer(c_detailed_signal))

	c_closure := (*C.GClosure)(closure.ToC())

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

	c_closure := (*C.GClosure)(closure.ToC())

	c_after :=
		boolToGboolean(after)

	retC := C.g_signal_connect_closure_by_id(c_instance, c_signal_id, c_detail, c_closure, c_after)
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported : g_signal_connect_data : unsupported parameter c_handler : no type generator for Callback, GCallback

// Unsupported : g_signal_connect_object : unsupported parameter c_handler : no type generator for Callback, GCallback

// Unsupported : g_signal_emit : unsupported parameter ... : varargs

// Unsupported : g_signal_emit_by_name : unsupported parameter ... : varargs

// Unsupported : g_signal_emit_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : g_signal_emitv : unsupported parameter instance_and_params : no param type

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

	c_closure := (*C.GClosure)(closure.ToC())

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

	c_closure := (*C.GClosure)(closure.ToC())

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

	c_closure := (*C.GClosure)(closure.ToC())

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

	c_closure := (*C.GClosure)(closure.ToC())

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

// Unsupported : g_signal_list_ids : unsupported parameter itype : no type generator for GType, GType

// Unsupported : g_signal_lookup : unsupported parameter itype : no type generator for GType, GType

// SignalName is a wrapper around the C function g_signal_name.
func SignalName(signalId uint32) string {
	c_signal_id := (C.guint)(signalId)

	retC := C.g_signal_name(c_signal_id)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_signal_new : unsupported parameter itype : no type generator for GType, GType

// Unsupported : g_signal_new_valist : unsupported parameter itype : no type generator for GType, GType

// Unsupported : g_signal_newv : unsupported parameter itype : no type generator for GType, GType

// Unsupported : g_signal_override_class_closure : unsupported parameter instance_type : no type generator for GType, GType

// Unsupported : g_signal_parse_name : unsupported parameter itype : no type generator for GType, GType

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

// Unsupported : g_signal_type_cclosure_new : unsupported parameter itype : no type generator for GType, GType

// SourceSetClosure is a wrapper around the C function g_source_set_closure.
func SourceSetClosure(source *glib.Source, closure *Closure) {
	c_source := (*C.GSource)(source.ToC())

	c_closure := (*C.GClosure)(closure.ToC())

	C.g_source_set_closure(c_source, c_closure)

	return
}

// SourceSetDummyCallback is a wrapper around the C function g_source_set_dummy_callback.
func SourceSetDummyCallback(source *glib.Source) {
	c_source := (*C.GSource)(source.ToC())

	C.g_source_set_dummy_callback(c_source)

	return
}

// StrdupValueContents is a wrapper around the C function g_strdup_value_contents.
func StrdupValueContents(value *Value) string {
	c_value := (*C.GValue)(value.ToC())

	retC := C.g_strdup_value_contents(c_value)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_type_add_class_cache_func : unsupported parameter cache_func : no type generator for TypeClassCacheFunc, GTypeClassCacheFunc

// Unsupported : g_type_add_instance_private : unsupported parameter class_type : no type generator for GType, GType

// Unsupported : g_type_add_interface_dynamic : unsupported parameter instance_type : no type generator for GType, GType

// Unsupported : g_type_add_interface_static : unsupported parameter instance_type : no type generator for GType, GType

// Unsupported : g_type_check_class_cast : unsupported parameter is_a_type : no type generator for GType, GType

// Unsupported : g_type_check_class_is_a : unsupported parameter is_a_type : no type generator for GType, GType

// TypeCheckInstance is a wrapper around the C function g_type_check_instance.
func TypeCheckInstance(instance *TypeInstance) bool {
	c_instance := (*C.GTypeInstance)(instance.ToC())

	retC := C.g_type_check_instance(c_instance)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_type_check_instance_cast : unsupported parameter iface_type : no type generator for GType, GType

// Unsupported : g_type_check_instance_is_a : unsupported parameter iface_type : no type generator for GType, GType

// Unsupported : g_type_check_instance_is_fundamentally_a : unsupported parameter fundamental_type : no type generator for GType, GType

// Unsupported : g_type_check_is_value_type : unsupported parameter type : no type generator for GType, GType

// TypeCheckValue is a wrapper around the C function g_type_check_value.
func TypeCheckValue(value *Value) bool {
	c_value := (*C.GValue)(value.ToC())

	retC := C.g_type_check_value(c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_type_check_value_holds : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_children : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_class_adjust_private_offset : unsupported parameter private_size_or_offset : no type generator for gint, gint*

// Unsupported : g_type_class_peek : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_class_ref : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_create_instance : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_depth : unsupported parameter type : no type generator for GType, GType

// TypeFreeInstance is a wrapper around the C function g_type_free_instance.
func TypeFreeInstance(instance *TypeInstance) {
	c_instance := (*C.GTypeInstance)(instance.ToC())

	C.g_type_free_instance(c_instance)

	return
}

// Unsupported : g_type_from_name : no return generator

// Unsupported : g_type_fundamental : unsupported parameter type_id : no type generator for GType, GType

// Unsupported : g_type_fundamental_next : no return generator

// Unsupported : g_type_get_plugin : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_get_qdata : unsupported parameter type : no type generator for GType, GType

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

// Unsupported : g_type_interface_add_prerequisite : unsupported parameter interface_type : no type generator for GType, GType

// Unsupported : g_type_interface_get_plugin : unsupported parameter instance_type : no type generator for GType, GType

// Unsupported : g_type_interface_peek : unsupported parameter iface_type : no type generator for GType, GType

// Unsupported : g_type_interfaces : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_is_a : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_name : unsupported parameter type : no type generator for GType, GType

// TypeNameFromClass is a wrapper around the C function g_type_name_from_class.
func TypeNameFromClass(gClass *TypeClass) string {
	c_g_class := (*C.GTypeClass)(gClass.ToC())

	retC := C.g_type_name_from_class(c_g_class)
	retGo := C.GoString(retC)

	return retGo
}

// TypeNameFromInstance is a wrapper around the C function g_type_name_from_instance.
func TypeNameFromInstance(instance *TypeInstance) string {
	c_instance := (*C.GTypeInstance)(instance.ToC())

	retC := C.g_type_name_from_instance(c_instance)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_type_next_base : unsupported parameter leaf_type : no type generator for GType, GType

// Unsupported : g_type_parent : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_qname : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_query : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_register_dynamic : unsupported parameter parent_type : no type generator for GType, GType

// Unsupported : g_type_register_fundamental : unsupported parameter type_id : no type generator for GType, GType

// Unsupported : g_type_register_static : unsupported parameter parent_type : no type generator for GType, GType

// Unsupported : g_type_remove_class_cache_func : unsupported parameter cache_func : no type generator for TypeClassCacheFunc, GTypeClassCacheFunc

// Unsupported : g_type_set_qdata : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_test_flags : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_value_table_peek : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_value_register_transform_func : unsupported parameter src_type : no type generator for GType, GType

// Unsupported : g_value_type_compatible : unsupported parameter src_type : no type generator for GType, GType

// Unsupported : g_value_type_transformable : unsupported parameter src_type : no type generator for GType, GType
