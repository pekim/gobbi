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

// Unsupported : g_cclosure_new : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_object : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_object_swap : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_swap : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// EnumCompleteTypeInfo is a wrapper around the C function g_enum_complete_type_info.
func EnumCompleteTypeInfo(gEnumType Type, constValues *EnumValue) *TypeInfo {
	c_g_enum_type := (C.GType)(gEnumType)

	var c_info C.GTypeInfo

	c_const_values := (*C.GEnumValue)(constValues.ToC())

	C.g_enum_complete_type_info(c_g_enum_type, &c_info, c_const_values)

	info := TypeInfoNewFromC(unsafe.Pointer(&c_info))

	return info
}

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

// EnumRegisterStatic is a wrapper around the C function g_enum_register_static.
func EnumRegisterStatic(name string, constStaticValues *EnumValue) Type {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_const_static_values := (*C.GEnumValue)(constStaticValues.ToC())

	retC := C.g_enum_register_static(c_name, c_const_static_values)
	retGo := (Type)(retC)

	return retGo
}

// FlagsCompleteTypeInfo is a wrapper around the C function g_flags_complete_type_info.
func FlagsCompleteTypeInfo(gFlagsType Type, constValues *FlagsValue) *TypeInfo {
	c_g_flags_type := (C.GType)(gFlagsType)

	var c_info C.GTypeInfo

	c_const_values := (*C.GFlagsValue)(constValues.ToC())

	C.g_flags_complete_type_info(c_g_flags_type, &c_info, c_const_values)

	info := TypeInfoNewFromC(unsafe.Pointer(&c_info))

	return info
}

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

// FlagsRegisterStatic is a wrapper around the C function g_flags_register_static.
func FlagsRegisterStatic(name string, constStaticValues *FlagsValue) Type {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_const_static_values := (*C.GFlagsValue)(constStaticValues.ToC())

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

// Unsupported : g_param_spec_boolean : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_boxed : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_char : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_double : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_enum : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_flags : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_float : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_int : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_int64 : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_long : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_object : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_param : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_pointer : return type : Blacklisted record : GParamSpec

// ParamSpecPoolNew is a wrapper around the C function g_param_spec_pool_new.
func ParamSpecPoolNew(typePrefixing bool) *ParamSpecPool {
	c_type_prefixing :=
		boolToGboolean(typePrefixing)

	retC := C.g_param_spec_pool_new(c_type_prefixing)
	retGo := ParamSpecPoolNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_param_spec_string : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_uchar : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_uint : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_uint64 : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_ulong : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_unichar : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_value_array : unsupported parameter element_spec : Blacklisted record : GParamSpec

// ParamTypeRegisterStatic is a wrapper around the C function g_param_type_register_static.
func ParamTypeRegisterStatic(name string, pspecInfo *ParamSpecTypeInfo) Type {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_pspec_info := (*C.GParamSpecTypeInfo)(pspecInfo.ToC())

	retC := C.g_param_type_register_static(c_name, c_pspec_info)
	retGo := (Type)(retC)

	return retGo
}

// Unsupported : g_param_value_convert : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : g_param_value_defaults : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : g_param_value_set_default : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : g_param_value_validate : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : g_param_values_cmp : unsupported parameter pspec : Blacklisted record : GParamSpec

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

// Unsupported : g_signal_list_ids : no return type

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

	c_class_closure := (*C.GClosure)(classClosure.ToC())

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

	c_info := (*C.GInterfaceInfo)(info.ToC())

	C.g_type_add_interface_static(c_instance_type, c_interface_type, c_info)

	return
}

// TypeCheckClassCast is a wrapper around the C function g_type_check_class_cast.
func TypeCheckClassCast(gClass *TypeClass, isAType Type) *TypeClass {
	c_g_class := (*C.GTypeClass)(gClass.ToC())

	c_is_a_type := (C.GType)(isAType)

	retC := C.g_type_check_class_cast(c_g_class, c_is_a_type)
	retGo := TypeClassNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeCheckClassIsA is a wrapper around the C function g_type_check_class_is_a.
func TypeCheckClassIsA(gClass *TypeClass, isAType Type) bool {
	c_g_class := (*C.GTypeClass)(gClass.ToC())

	c_is_a_type := (C.GType)(isAType)

	retC := C.g_type_check_class_is_a(c_g_class, c_is_a_type)
	retGo := retC == C.TRUE

	return retGo
}

// TypeCheckInstance is a wrapper around the C function g_type_check_instance.
func TypeCheckInstance(instance *TypeInstance) bool {
	c_instance := (*C.GTypeInstance)(instance.ToC())

	retC := C.g_type_check_instance(c_instance)
	retGo := retC == C.TRUE

	return retGo
}

// TypeCheckInstanceCast is a wrapper around the C function g_type_check_instance_cast.
func TypeCheckInstanceCast(instance *TypeInstance, ifaceType Type) *TypeInstance {
	c_instance := (*C.GTypeInstance)(instance.ToC())

	c_iface_type := (C.GType)(ifaceType)

	retC := C.g_type_check_instance_cast(c_instance, c_iface_type)
	retGo := TypeInstanceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeCheckInstanceIsA is a wrapper around the C function g_type_check_instance_is_a.
func TypeCheckInstanceIsA(instance *TypeInstance, ifaceType Type) bool {
	c_instance := (*C.GTypeInstance)(instance.ToC())

	c_iface_type := (C.GType)(ifaceType)

	retC := C.g_type_check_instance_is_a(c_instance, c_iface_type)
	retGo := retC == C.TRUE

	return retGo
}

// TypeCheckInstanceIsFundamentallyA is a wrapper around the C function g_type_check_instance_is_fundamentally_a.
func TypeCheckInstanceIsFundamentallyA(instance *TypeInstance, fundamentalType Type) bool {
	c_instance := (*C.GTypeInstance)(instance.ToC())

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
	c_value := (*C.GValue)(value.ToC())

	retC := C.g_type_check_value(c_value)
	retGo := retC == C.TRUE

	return retGo
}

// TypeCheckValueHolds is a wrapper around the C function g_type_check_value_holds.
func TypeCheckValueHolds(value *Value, type_ Type) bool {
	c_value := (*C.GValue)(value.ToC())

	c_type := (C.GType)(type_)

	retC := C.g_type_check_value_holds(c_value, c_type)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_type_children : no return type

// TypeClassAdjustPrivateOffset is a wrapper around the C function g_type_class_adjust_private_offset.
func TypeClassAdjustPrivateOffset(gClass uintptr, privateSizeOrOffset int32) {
	c_g_class := (C.gpointer)(gClass)

	c_private_size_or_offset := (C.gint)(privateSizeOrOffset)

	C.g_type_class_adjust_private_offset(c_g_class, &c_private_size_or_offset)

	return
}

// TypeClassPeek is a wrapper around the C function g_type_class_peek.
func TypeClassPeek(type_ Type) uintptr {
	c_type := (C.GType)(type_)

	retC := C.g_type_class_peek(c_type)
	retGo := (uintptr)(retC)

	return retGo
}

// TypeClassRef is a wrapper around the C function g_type_class_ref.
func TypeClassRef(type_ Type) uintptr {
	c_type := (C.GType)(type_)

	retC := C.g_type_class_ref(c_type)
	retGo := (uintptr)(retC)

	return retGo
}

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
	c_instance := (*C.GTypeInstance)(instance.ToC())

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

// TypeInterfaceAddPrerequisite is a wrapper around the C function g_type_interface_add_prerequisite.
func TypeInterfaceAddPrerequisite(interfaceType Type, prerequisiteType Type) {
	c_interface_type := (C.GType)(interfaceType)

	c_prerequisite_type := (C.GType)(prerequisiteType)

	C.g_type_interface_add_prerequisite(c_interface_type, c_prerequisite_type)

	return
}

// TypeInterfaceGetPlugin is a wrapper around the C function g_type_interface_get_plugin.
func TypeInterfaceGetPlugin(instanceType Type, interfaceType Type) *TypePlugin {
	c_instance_type := (C.GType)(instanceType)

	c_interface_type := (C.GType)(interfaceType)

	retC := C.g_type_interface_get_plugin(c_instance_type, c_interface_type)
	retGo := TypePluginNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeInterfacePeek is a wrapper around the C function g_type_interface_peek.
func TypeInterfacePeek(instanceClass uintptr, ifaceType Type) uintptr {
	c_instance_class := (C.gpointer)(instanceClass)

	c_iface_type := (C.GType)(ifaceType)

	retC := C.g_type_interface_peek(c_instance_class, c_iface_type)
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_type_interfaces : no return type

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

	c_info := (*C.GTypeInfo)(info.ToC())

	c_finfo := (*C.GTypeFundamentalInfo)(finfo.ToC())

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

	c_info := (*C.GTypeInfo)(info.ToC())

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

// TypeValueTablePeek is a wrapper around the C function g_type_value_table_peek.
func TypeValueTablePeek(type_ Type) *TypeValueTable {
	c_type := (C.GType)(type_)

	retC := C.g_type_value_table_peek(c_type)
	retGo := TypeValueTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_value_register_transform_func : unsupported parameter transform_func : no type generator for ValueTransform (GValueTransform) for param transform_func

// ValueTypeCompatible is a wrapper around the C function g_value_type_compatible.
func ValueTypeCompatible(srcType Type, destType Type) bool {
	c_src_type := (C.GType)(srcType)

	c_dest_type := (C.GType)(destType)

	retC := C.g_value_type_compatible(c_src_type, c_dest_type)
	retGo := retC == C.TRUE

	return retGo
}

// ValueTypeTransformable is a wrapper around the C function g_value_type_transformable.
func ValueTypeTransformable(srcType Type, destType Type) bool {
	c_src_type := (C.GType)(srcType)

	c_dest_type := (C.GType)(destType)

	retC := C.g_value_type_transformable(c_src_type, c_dest_type)
	retGo := retC == C.TRUE

	return retGo
}
