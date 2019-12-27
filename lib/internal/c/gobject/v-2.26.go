// Code generated - DO NOT EDIT.
// +build gobject_2.26

package gobject

import "unsafe"

// #include <glib-object.h>
// #include <stdlib.h>
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}
func toGoBool(b C.gboolean) bool {
	return b == C.TRUE
}

type CClosure C.GCClosure
type Closure C.GClosure
type ClosureNotifyData C.GClosureNotifyData
type EnumClass C.GEnumClass
type EnumValue C.GEnumValue
type FlagsClass C.GFlagsClass
type FlagsValue C.GFlagsValue
type InitiallyUnownedClass C.GInitiallyUnownedClass
type InterfaceInfo C.GInterfaceInfo
type ObjectClass C.GObjectClass
type ObjectConstructParam C.GObjectConstructParam
type ParamSpecClass C.GParamSpecClass
type ParamSpecPool C.GParamSpecPool
type ParamSpecTypeInfo C.GParamSpecTypeInfo
type Parameter C.GParameter
type SignalInvocationHint C.GSignalInvocationHint
type SignalQuery C.GSignalQuery
type TypeClass C.GTypeClass
type TypeFundamentalInfo C.GTypeFundamentalInfo
type TypeInfo C.GTypeInfo
type TypeInstance C.GTypeInstance
type TypeInterface C.GTypeInterface
type TypeModuleClass C.GTypeModuleClass
type TypePluginClass C.GTypePluginClass
type TypeQuery C.GTypeQuery
type TypeValueTable C.GTypeValueTable
type Value C.GValue
type ValueArray C.GValueArray
type WeakRef C.GWeakRef

func Fn_g_boxed_copy(param0 uint64, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.gconstpointer)(param1)

	ret := C.g_boxed_copy(cValue0, cValue1)

	return (unsafe.Pointer)(ret)
}

func Fn_g_boxed_free(param0 uint64, param1 unsafe.Pointer) {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.gpointer)(param1)

	C.g_boxed_free(cValue0, cValue1)
}

// UNSUPPORTED : boxed_type_register_static : has callback

func Fn_g_cclosure_marshal_BOOLEAN__BOXED_BOXED(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_BOOLEAN__BOXED_BOXED(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_BOOLEAN__FLAGS(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_BOOLEAN__FLAGS(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_STRING__OBJECT_POINTER(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_STRING__OBJECT_POINTER(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_VOID__BOOLEAN(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_VOID__BOOLEAN(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_VOID__BOXED(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_VOID__BOXED(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_VOID__CHAR(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_VOID__CHAR(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_VOID__DOUBLE(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_VOID__DOUBLE(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_VOID__ENUM(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_VOID__ENUM(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_VOID__FLAGS(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_VOID__FLAGS(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_VOID__FLOAT(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_VOID__FLOAT(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_VOID__INT(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_VOID__INT(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_VOID__LONG(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_VOID__LONG(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_VOID__OBJECT(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_VOID__OBJECT(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_VOID__PARAM(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_VOID__PARAM(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_VOID__POINTER(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_VOID__POINTER(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_VOID__STRING(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_VOID__STRING(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_VOID__UCHAR(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_VOID__UCHAR(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_VOID__UINT(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_VOID__UINT(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_VOID__UINT_POINTER(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_VOID__UINT_POINTER(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_VOID__ULONG(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_VOID__ULONG(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_VOID__VARIANT(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_VOID__VARIANT(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_cclosure_marshal_VOID__VOID(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	cValue5 := (C.gpointer)(param5)

	C.g_cclosure_marshal_VOID__VOID(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

// UNSUPPORTED : cclosure_new : has callback

// UNSUPPORTED : cclosure_new_object : has callback

// UNSUPPORTED : cclosure_new_object_swap : has callback

// UNSUPPORTED : cclosure_new_swap : has callback

func Fn_g_enum_complete_type_info(param0 uint64, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValue0 := (C.GType)(param0)

	cValue1 := (*C.GTypeInfo)(unsafe.Pointer(param1))

	cValue2 := (*C.GEnumValue)(unsafe.Pointer(param2))

	C.g_enum_complete_type_info(cValue0, cValue1, cValue2)
}

func Fn_g_enum_get_value(param0 unsafe.Pointer, param1 int) unsafe.Pointer {
	cValue0 := (*C.GEnumClass)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	ret := C.g_enum_get_value(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_enum_get_value_by_name(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GEnumClass)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_enum_get_value_by_name(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_enum_get_value_by_nick(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GEnumClass)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_enum_get_value_by_nick(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_enum_register_static(param0 string, param1 unsafe.Pointer) uint64 {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GEnumValue)(unsafe.Pointer(param1))

	ret := C.g_enum_register_static(cValue0, cValue1)

	return (uint64)(ret)
}

func Fn_g_flags_complete_type_info(param0 uint64, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValue0 := (C.GType)(param0)

	cValue1 := (*C.GTypeInfo)(unsafe.Pointer(param1))

	cValue2 := (*C.GFlagsValue)(unsafe.Pointer(param2))

	C.g_flags_complete_type_info(cValue0, cValue1, cValue2)
}

func Fn_g_flags_get_first_value(param0 unsafe.Pointer, param1 uint) unsafe.Pointer {
	cValue0 := (*C.GFlagsClass)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	ret := C.g_flags_get_first_value(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_flags_get_value_by_name(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GFlagsClass)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_flags_get_value_by_name(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_flags_get_value_by_nick(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GFlagsClass)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_flags_get_value_by_nick(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_flags_register_static(param0 string, param1 unsafe.Pointer) uint64 {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GFlagsValue)(unsafe.Pointer(param1))

	ret := C.g_flags_register_static(cValue0, cValue1)

	return (uint64)(ret)
}

func Fn_g_gtype_get_type() uint64 {
	ret := C.g_gtype_get_type()

	return (uint64)(ret)
}

func Fn_g_param_spec_boolean(param0 string, param1 string, param2 string, param3 bool, param4 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := toCBool(param3)

	cValue4 := (C.GParamFlags)(param4)

	ret := C.g_param_spec_boolean(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_boxed(param0 string, param1 string, param2 string, param3 uint64, param4 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.GType)(param3)

	cValue4 := (C.GParamFlags)(param4)

	ret := C.g_param_spec_boxed(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_char(param0 string, param1 string, param2 string, param3 int8, param4 int8, param5 int8, param6 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.gint8)(param3)

	cValue4 := (C.gint8)(param4)

	cValue5 := (C.gint8)(param5)

	cValue6 := (C.GParamFlags)(param6)

	ret := C.g_param_spec_char(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_double(param0 string, param1 string, param2 string, param3 float64, param4 float64, param5 float64, param6 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	cValue6 := (C.GParamFlags)(param6)

	ret := C.g_param_spec_double(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_enum(param0 string, param1 string, param2 string, param3 uint64, param4 int, param5 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.GType)(param3)

	cValue4 := (C.gint)(param4)

	cValue5 := (C.GParamFlags)(param5)

	ret := C.g_param_spec_enum(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_flags(param0 string, param1 string, param2 string, param3 uint64, param4 uint, param5 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.GType)(param3)

	cValue4 := (C.guint)(param4)

	cValue5 := (C.GParamFlags)(param5)

	ret := C.g_param_spec_flags(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_float(param0 string, param1 string, param2 string, param3 float32, param4 float32, param5 float32, param6 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.gfloat)(param3)

	cValue4 := (C.gfloat)(param4)

	cValue5 := (C.gfloat)(param5)

	cValue6 := (C.GParamFlags)(param6)

	ret := C.g_param_spec_float(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_gtype(param0 string, param1 string, param2 string, param3 uint64, param4 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.GType)(param3)

	cValue4 := (C.GParamFlags)(param4)

	ret := C.g_param_spec_gtype(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_int(param0 string, param1 string, param2 string, param3 int, param4 int, param5 int, param6 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.gint)(param3)

	cValue4 := (C.gint)(param4)

	cValue5 := (C.gint)(param5)

	cValue6 := (C.GParamFlags)(param6)

	ret := C.g_param_spec_int(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_int64(param0 string, param1 string, param2 string, param3 int64, param4 int64, param5 int64, param6 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.gint64)(param3)

	cValue4 := (C.gint64)(param4)

	cValue5 := (C.gint64)(param5)

	cValue6 := (C.GParamFlags)(param6)

	ret := C.g_param_spec_int64(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_long(param0 string, param1 string, param2 string, param3 int64, param4 int64, param5 int64, param6 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.glong)(param3)

	cValue4 := (C.glong)(param4)

	cValue5 := (C.glong)(param5)

	cValue6 := (C.GParamFlags)(param6)

	ret := C.g_param_spec_long(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_object(param0 string, param1 string, param2 string, param3 uint64, param4 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.GType)(param3)

	cValue4 := (C.GParamFlags)(param4)

	ret := C.g_param_spec_object(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_override(param0 string, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GParamSpec)(unsafe.Pointer(param1))

	ret := C.g_param_spec_override(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_param(param0 string, param1 string, param2 string, param3 uint64, param4 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.GType)(param3)

	cValue4 := (C.GParamFlags)(param4)

	ret := C.g_param_spec_param(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_pointer(param0 string, param1 string, param2 string, param3 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.GParamFlags)(param3)

	ret := C.g_param_spec_pointer(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_pool_new(param0 bool) unsafe.Pointer {
	cValue0 := toCBool(param0)

	ret := C.g_param_spec_pool_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_string(param0 string, param1 string, param2 string, param3 string, param4 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	cValue4 := (C.GParamFlags)(param4)

	ret := C.g_param_spec_string(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_uchar(param0 string, param1 string, param2 string, param3 uint8, param4 uint8, param5 uint8, param6 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.guint8)(param3)

	cValue4 := (C.guint8)(param4)

	cValue5 := (C.guint8)(param5)

	cValue6 := (C.GParamFlags)(param6)

	ret := C.g_param_spec_uchar(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_uint(param0 string, param1 string, param2 string, param3 uint, param4 uint, param5 uint, param6 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.guint)(param3)

	cValue4 := (C.guint)(param4)

	cValue5 := (C.guint)(param5)

	cValue6 := (C.GParamFlags)(param6)

	ret := C.g_param_spec_uint(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_uint64(param0 string, param1 string, param2 string, param3 uint64, param4 uint64, param5 uint64, param6 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.guint64)(param3)

	cValue4 := (C.guint64)(param4)

	cValue5 := (C.guint64)(param5)

	cValue6 := (C.GParamFlags)(param6)

	ret := C.g_param_spec_uint64(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_ulong(param0 string, param1 string, param2 string, param3 uint64, param4 uint64, param5 uint64, param6 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.gulong)(param3)

	cValue4 := (C.gulong)(param4)

	cValue5 := (C.gulong)(param5)

	cValue6 := (C.GParamFlags)(param6)

	ret := C.g_param_spec_ulong(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_unichar(param0 string, param1 string, param2 string, param3 rune, param4 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.gunichar)(param3)

	cValue4 := (C.GParamFlags)(param4)

	ret := C.g_param_spec_unichar(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_value_array(param0 string, param1 string, param2 string, param3 unsafe.Pointer, param4 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.GParamSpec)(unsafe.Pointer(param3))

	cValue4 := (C.GParamFlags)(param4)

	ret := C.g_param_spec_value_array(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_variant(param0 string, param1 string, param2 string, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.GVariantType)(unsafe.Pointer(param3))

	cValue4 := (*C.GVariant)(unsafe.Pointer(param4))

	cValue5 := (C.GParamFlags)(param5)

	ret := C.g_param_spec_variant(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return unsafe.Pointer(ret)
}

func Fn_g_param_type_register_static(param0 string, param1 unsafe.Pointer) uint64 {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GParamSpecTypeInfo)(unsafe.Pointer(param1))

	ret := C.g_param_type_register_static(cValue0, cValue1)

	return (uint64)(ret)
}

func Fn_g_param_value_convert(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 bool) bool {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	cValue3 := toCBool(param3)

	ret := C.g_param_value_convert(cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

func Fn_g_param_value_defaults(param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	ret := C.g_param_value_defaults(cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_param_value_set_default(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	C.g_param_value_set_default(cValue0, cValue1)
}

func Fn_g_param_value_validate(param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	ret := C.g_param_value_validate(cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_param_values_cmp(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) int {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	ret := C.g_param_values_cmp(cValue0, cValue1, cValue2)

	return (int)(ret)
}

func Fn_g_pointer_type_register_static(param0 string) uint64 {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_pointer_type_register_static(cValue0)

	return (uint64)(ret)
}

func Fn_g_signal_accumulator_true_handled(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer) bool {
	cValue0 := (*C.GSignalInvocationHint)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	cValue3 := (C.gpointer)(param3)

	ret := C.g_signal_accumulator_true_handled(cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

// UNSUPPORTED : signal_add_emission_hook : has callback

// UNSUPPORTED : signal_chain_from_overridden : has non-string array param instance_and_params

// UNSUPPORTED : signal_chain_from_overridden_handler : has varargs

func Fn_g_signal_connect_closure(param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer, param3 bool) uint64 {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GClosure)(unsafe.Pointer(param2))

	cValue3 := toCBool(param3)

	ret := C.g_signal_connect_closure(cValue0, cValue1, cValue2, cValue3)

	return (uint64)(ret)
}

func Fn_g_signal_connect_closure_by_id(param0 unsafe.Pointer, param1 uint, param2 uint32, param3 unsafe.Pointer, param4 bool) uint64 {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GQuark)(param2)

	cValue3 := (*C.GClosure)(unsafe.Pointer(param3))

	cValue4 := toCBool(param4)

	ret := C.g_signal_connect_closure_by_id(cValue0, cValue1, cValue2, cValue3, cValue4)

	return (uint64)(ret)
}

// UNSUPPORTED : signal_connect_data : has callback

// UNSUPPORTED : signal_connect_object : has callback

// UNSUPPORTED : signal_emit : has varargs

// UNSUPPORTED : signal_emit_by_name : has varargs

// UNSUPPORTED : signal_emit_valist : has va_list

// UNSUPPORTED : signal_emitv : has non-string array param instance_and_params

func Fn_g_signal_get_invocation_hint(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (C.gpointer)(param0)

	ret := C.g_signal_get_invocation_hint(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_signal_handler_block(param0 unsafe.Pointer, param1 uint64) {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (C.gulong)(param1)

	C.g_signal_handler_block(cValue0, cValue1)
}

func Fn_g_signal_handler_disconnect(param0 unsafe.Pointer, param1 uint64) {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (C.gulong)(param1)

	C.g_signal_handler_disconnect(cValue0, cValue1)
}

func Fn_g_signal_handler_find(param0 unsafe.Pointer, param1 int, param2 uint, param3 uint32, param4 unsafe.Pointer, param5 unsafe.Pointer, param6 unsafe.Pointer) uint64 {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (C.GSignalMatchType)(param1)

	cValue2 := (C.guint)(param2)

	cValue3 := (C.GQuark)(param3)

	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))

	cValue5 := (C.gpointer)(param5)

	cValue6 := (C.gpointer)(param6)

	ret := C.g_signal_handler_find(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)

	return (uint64)(ret)
}

func Fn_g_signal_handler_is_connected(param0 unsafe.Pointer, param1 uint64) bool {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (C.gulong)(param1)

	ret := C.g_signal_handler_is_connected(cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_signal_handler_unblock(param0 unsafe.Pointer, param1 uint64) {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (C.gulong)(param1)

	C.g_signal_handler_unblock(cValue0, cValue1)
}

func Fn_g_signal_handlers_block_matched(param0 unsafe.Pointer, param1 int, param2 uint, param3 uint32, param4 unsafe.Pointer, param5 unsafe.Pointer, param6 unsafe.Pointer) uint {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (C.GSignalMatchType)(param1)

	cValue2 := (C.guint)(param2)

	cValue3 := (C.GQuark)(param3)

	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))

	cValue5 := (C.gpointer)(param5)

	cValue6 := (C.gpointer)(param6)

	ret := C.g_signal_handlers_block_matched(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)

	return (uint)(ret)
}

func Fn_g_signal_handlers_destroy(param0 unsafe.Pointer) {
	cValue0 := (C.gpointer)(param0)

	C.g_signal_handlers_destroy(cValue0)
}

func Fn_g_signal_handlers_disconnect_matched(param0 unsafe.Pointer, param1 int, param2 uint, param3 uint32, param4 unsafe.Pointer, param5 unsafe.Pointer, param6 unsafe.Pointer) uint {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (C.GSignalMatchType)(param1)

	cValue2 := (C.guint)(param2)

	cValue3 := (C.GQuark)(param3)

	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))

	cValue5 := (C.gpointer)(param5)

	cValue6 := (C.gpointer)(param6)

	ret := C.g_signal_handlers_disconnect_matched(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)

	return (uint)(ret)
}

func Fn_g_signal_handlers_unblock_matched(param0 unsafe.Pointer, param1 int, param2 uint, param3 uint32, param4 unsafe.Pointer, param5 unsafe.Pointer, param6 unsafe.Pointer) uint {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (C.GSignalMatchType)(param1)

	cValue2 := (C.guint)(param2)

	cValue3 := (C.GQuark)(param3)

	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))

	cValue5 := (C.gpointer)(param5)

	cValue6 := (C.gpointer)(param6)

	ret := C.g_signal_handlers_unblock_matched(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)

	return (uint)(ret)
}

func Fn_g_signal_has_handler_pending(param0 unsafe.Pointer, param1 uint, param2 uint32, param3 bool) bool {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GQuark)(param2)

	cValue3 := toCBool(param3)

	ret := C.g_signal_has_handler_pending(cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

// UNSUPPORTED : signal_list_ids : has array return

func Fn_g_signal_lookup(param0 string, param1 uint64) uint {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GType)(param1)

	ret := C.g_signal_lookup(cValue0, cValue1)

	return (uint)(ret)
}

func Fn_g_signal_name(param0 uint) string {
	cValue0 := (C.guint)(param0)

	ret := C.g_signal_name(cValue0)

	return C.GoString(ret)
}

// UNSUPPORTED : signal_new : has varargs

// UNSUPPORTED : signal_new_class_handler : has varargs

// UNSUPPORTED : signal_new_valist : has va_list

// UNSUPPORTED : signal_newv : has callback

func Fn_g_signal_override_class_closure(param0 uint, param1 uint64, param2 unsafe.Pointer) {
	cValue0 := (C.guint)(param0)

	cValue1 := (C.GType)(param1)

	cValue2 := (*C.GClosure)(unsafe.Pointer(param2))

	C.g_signal_override_class_closure(cValue0, cValue1, cValue2)
}

// UNSUPPORTED : signal_override_class_handler : has callback

func Fn_g_signal_parse_name(param0 string, param1 uint64, param2 *uint, param3 *uint32, param4 bool) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GType)(param1)

	cValue2 := (*C.guint)(unsafe.Pointer(param2))

	cValue3 := (*C.GQuark)(unsafe.Pointer(param3))

	cValue4 := toCBool(param4)

	ret := C.g_signal_parse_name(cValue0, cValue1, cValue2, cValue3, cValue4)

	return toGoBool(ret)
}

func Fn_g_signal_query(param0 uint, param1 unsafe.Pointer) {
	cValue0 := (C.guint)(param0)

	cValue1 := (*C.GSignalQuery)(unsafe.Pointer(param1))

	C.g_signal_query(cValue0, cValue1)
}

func Fn_g_signal_remove_emission_hook(param0 uint, param1 uint64) {
	cValue0 := (C.guint)(param0)

	cValue1 := (C.gulong)(param1)

	C.g_signal_remove_emission_hook(cValue0, cValue1)
}

func Fn_g_signal_stop_emission(param0 unsafe.Pointer, param1 uint, param2 uint32) {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GQuark)(param2)

	C.g_signal_stop_emission(cValue0, cValue1, cValue2)
}

func Fn_g_signal_stop_emission_by_name(param0 unsafe.Pointer, param1 string) {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_signal_stop_emission_by_name(cValue0, cValue1)
}

func Fn_g_signal_type_cclosure_new(param0 uint64, param1 uint) unsafe.Pointer {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.guint)(param1)

	ret := C.g_signal_type_cclosure_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_source_set_closure(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GSource)(unsafe.Pointer(param0))

	cValue1 := (*C.GClosure)(unsafe.Pointer(param1))

	C.g_source_set_closure(cValue0, cValue1)
}

func Fn_g_source_set_dummy_callback(param0 unsafe.Pointer) {
	cValue0 := (*C.GSource)(unsafe.Pointer(param0))

	C.g_source_set_dummy_callback(cValue0)
}

func Fn_g_strdup_value_contents(param0 unsafe.Pointer) string {
	cValue0 := (*C.GValue)(unsafe.Pointer(param0))

	ret := C.g_strdup_value_contents(cValue0)

	return C.GoString(ret)
}

// UNSUPPORTED : type_add_class_cache_func : has callback

func Fn_g_type_add_class_private(param0 uint64, param1 uint64) {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.gsize)(param1)

	C.g_type_add_class_private(cValue0, cValue1)
}

func Fn_g_type_add_instance_private(param0 uint64, param1 uint64) int {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.gsize)(param1)

	ret := C.g_type_add_instance_private(cValue0, cValue1)

	return (int)(ret)
}

// UNSUPPORTED : type_add_interface_check : has callback

func Fn_g_type_add_interface_dynamic(param0 uint64, param1 uint64, param2 unsafe.Pointer) {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.GType)(param1)

	cValue2 := (*C.GTypePlugin)(unsafe.Pointer(param2))

	C.g_type_add_interface_dynamic(cValue0, cValue1, cValue2)
}

func Fn_g_type_add_interface_static(param0 uint64, param1 uint64, param2 unsafe.Pointer) {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.GType)(param1)

	cValue2 := (*C.GInterfaceInfo)(unsafe.Pointer(param2))

	C.g_type_add_interface_static(cValue0, cValue1, cValue2)
}

func Fn_g_type_check_class_cast(param0 unsafe.Pointer, param1 uint64) unsafe.Pointer {
	cValue0 := (*C.GTypeClass)(unsafe.Pointer(param0))

	cValue1 := (C.GType)(param1)

	ret := C.g_type_check_class_cast(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_type_check_class_is_a(param0 unsafe.Pointer, param1 uint64) bool {
	cValue0 := (*C.GTypeClass)(unsafe.Pointer(param0))

	cValue1 := (C.GType)(param1)

	ret := C.g_type_check_class_is_a(cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_type_check_instance(param0 unsafe.Pointer) bool {
	cValue0 := (*C.GTypeInstance)(unsafe.Pointer(param0))

	ret := C.g_type_check_instance(cValue0)

	return toGoBool(ret)
}

func Fn_g_type_check_instance_cast(param0 unsafe.Pointer, param1 uint64) unsafe.Pointer {
	cValue0 := (*C.GTypeInstance)(unsafe.Pointer(param0))

	cValue1 := (C.GType)(param1)

	ret := C.g_type_check_instance_cast(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_type_check_instance_is_a(param0 unsafe.Pointer, param1 uint64) bool {
	cValue0 := (*C.GTypeInstance)(unsafe.Pointer(param0))

	cValue1 := (C.GType)(param1)

	ret := C.g_type_check_instance_is_a(cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_type_check_instance_is_fundamentally_a(param0 unsafe.Pointer, param1 uint64) bool {
	cValue0 := (*C.GTypeInstance)(unsafe.Pointer(param0))

	cValue1 := (C.GType)(param1)

	ret := C.g_type_check_instance_is_fundamentally_a(cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_type_check_is_value_type(param0 uint64) bool {
	cValue0 := (C.GType)(param0)

	ret := C.g_type_check_is_value_type(cValue0)

	return toGoBool(ret)
}

func Fn_g_type_check_value(param0 unsafe.Pointer) bool {
	cValue0 := (*C.GValue)(unsafe.Pointer(param0))

	ret := C.g_type_check_value(cValue0)

	return toGoBool(ret)
}

func Fn_g_type_check_value_holds(param0 unsafe.Pointer, param1 uint64) bool {
	cValue0 := (*C.GValue)(unsafe.Pointer(param0))

	cValue1 := (C.GType)(param1)

	ret := C.g_type_check_value_holds(cValue0, cValue1)

	return toGoBool(ret)
}

// UNSUPPORTED : type_children : has array return

func Fn_g_type_class_adjust_private_offset(param0 unsafe.Pointer, param1 *int) {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.g_type_class_adjust_private_offset(cValue0, cValue1)
}

func Fn_g_type_class_peek(param0 uint64) unsafe.Pointer {
	cValue0 := (C.GType)(param0)

	ret := C.g_type_class_peek(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_type_class_peek_static(param0 uint64) unsafe.Pointer {
	cValue0 := (C.GType)(param0)

	ret := C.g_type_class_peek_static(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_type_class_ref(param0 uint64) unsafe.Pointer {
	cValue0 := (C.GType)(param0)

	ret := C.g_type_class_ref(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_type_create_instance(param0 uint64) unsafe.Pointer {
	cValue0 := (C.GType)(param0)

	ret := C.g_type_create_instance(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_type_default_interface_peek(param0 uint64) unsafe.Pointer {
	cValue0 := (C.GType)(param0)

	ret := C.g_type_default_interface_peek(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_type_default_interface_ref(param0 uint64) unsafe.Pointer {
	cValue0 := (C.GType)(param0)

	ret := C.g_type_default_interface_ref(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_type_default_interface_unref(param0 unsafe.Pointer) {
	cValue0 := (C.gpointer)(param0)

	C.g_type_default_interface_unref(cValue0)
}

func Fn_g_type_depth(param0 uint64) uint {
	cValue0 := (C.GType)(param0)

	ret := C.g_type_depth(cValue0)

	return (uint)(ret)
}

func Fn_g_type_free_instance(param0 unsafe.Pointer) {
	cValue0 := (*C.GTypeInstance)(unsafe.Pointer(param0))

	C.g_type_free_instance(cValue0)
}

func Fn_g_type_from_name(param0 string) uint64 {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_type_from_name(cValue0)

	return (uint64)(ret)
}

func Fn_g_type_fundamental(param0 uint64) uint64 {
	cValue0 := (C.GType)(param0)

	ret := C.g_type_fundamental(cValue0)

	return (uint64)(ret)
}

func Fn_g_type_fundamental_next() uint64 {
	ret := C.g_type_fundamental_next()

	return (uint64)(ret)
}

func Fn_g_type_get_plugin(param0 uint64) unsafe.Pointer {
	cValue0 := (C.GType)(param0)

	ret := C.g_type_get_plugin(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_type_get_qdata(param0 uint64, param1 uint32) unsafe.Pointer {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.GQuark)(param1)

	ret := C.g_type_get_qdata(cValue0, cValue1)

	return (unsafe.Pointer)(ret)
}

func Fn_g_type_init() {
	C.g_type_init()
}

func Fn_g_type_init_with_debug_flags(param0 int) {
	cValue0 := (C.GTypeDebugFlags)(param0)

	C.g_type_init_with_debug_flags(cValue0)
}

func Fn_g_type_interface_add_prerequisite(param0 uint64, param1 uint64) {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.GType)(param1)

	C.g_type_interface_add_prerequisite(cValue0, cValue1)
}

func Fn_g_type_interface_get_plugin(param0 uint64, param1 uint64) unsafe.Pointer {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.GType)(param1)

	ret := C.g_type_interface_get_plugin(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_type_interface_peek(param0 unsafe.Pointer, param1 uint64) unsafe.Pointer {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (C.GType)(param1)

	ret := C.g_type_interface_peek(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : type_interface_prerequisites : has array return

// UNSUPPORTED : type_interfaces : has array return

func Fn_g_type_is_a(param0 uint64, param1 uint64) bool {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.GType)(param1)

	ret := C.g_type_is_a(cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_type_name(param0 uint64) string {
	cValue0 := (C.GType)(param0)

	ret := C.g_type_name(cValue0)

	return C.GoString(ret)
}

func Fn_g_type_name_from_class(param0 unsafe.Pointer) string {
	cValue0 := (*C.GTypeClass)(unsafe.Pointer(param0))

	ret := C.g_type_name_from_class(cValue0)

	return C.GoString(ret)
}

func Fn_g_type_name_from_instance(param0 unsafe.Pointer) string {
	cValue0 := (*C.GTypeInstance)(unsafe.Pointer(param0))

	ret := C.g_type_name_from_instance(cValue0)

	return C.GoString(ret)
}

func Fn_g_type_next_base(param0 uint64, param1 uint64) uint64 {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.GType)(param1)

	ret := C.g_type_next_base(cValue0, cValue1)

	return (uint64)(ret)
}

func Fn_g_type_parent(param0 uint64) uint64 {
	cValue0 := (C.GType)(param0)

	ret := C.g_type_parent(cValue0)

	return (uint64)(ret)
}

func Fn_g_type_qname(param0 uint64) uint32 {
	cValue0 := (C.GType)(param0)

	ret := C.g_type_qname(cValue0)

	return (uint32)(ret)
}

func Fn_g_type_query(param0 uint64, param1 unsafe.Pointer) {
	cValue0 := (C.GType)(param0)

	cValue1 := (*C.GTypeQuery)(unsafe.Pointer(param1))

	C.g_type_query(cValue0, cValue1)
}

func Fn_g_type_register_dynamic(param0 uint64, param1 string, param2 unsafe.Pointer, param3 int) uint64 {
	cValue0 := (C.GType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GTypePlugin)(unsafe.Pointer(param2))

	cValue3 := (C.GTypeFlags)(param3)

	ret := C.g_type_register_dynamic(cValue0, cValue1, cValue2, cValue3)

	return (uint64)(ret)
}

func Fn_g_type_register_fundamental(param0 uint64, param1 string, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int) uint64 {
	cValue0 := (C.GType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GTypeInfo)(unsafe.Pointer(param2))

	cValue3 := (*C.GTypeFundamentalInfo)(unsafe.Pointer(param3))

	cValue4 := (C.GTypeFlags)(param4)

	ret := C.g_type_register_fundamental(cValue0, cValue1, cValue2, cValue3, cValue4)

	return (uint64)(ret)
}

func Fn_g_type_register_static(param0 uint64, param1 string, param2 unsafe.Pointer, param3 int) uint64 {
	cValue0 := (C.GType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GTypeInfo)(unsafe.Pointer(param2))

	cValue3 := (C.GTypeFlags)(param3)

	ret := C.g_type_register_static(cValue0, cValue1, cValue2, cValue3)

	return (uint64)(ret)
}

// UNSUPPORTED : type_register_static_simple : has callback

// UNSUPPORTED : type_remove_class_cache_func : has callback

// UNSUPPORTED : type_remove_interface_check : has callback

func Fn_g_type_set_qdata(param0 uint64, param1 uint32, param2 unsafe.Pointer) {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.GQuark)(param1)

	cValue2 := (C.gpointer)(param2)

	C.g_type_set_qdata(cValue0, cValue1, cValue2)
}

func Fn_g_type_test_flags(param0 uint64, param1 uint) bool {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.guint)(param1)

	ret := C.g_type_test_flags(cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_type_value_table_peek(param0 uint64) unsafe.Pointer {
	cValue0 := (C.GType)(param0)

	ret := C.g_type_value_table_peek(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : value_register_transform_func : has callback

func Fn_g_value_type_compatible(param0 uint64, param1 uint64) bool {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.GType)(param1)

	ret := C.g_value_type_compatible(cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_value_type_transformable(param0 uint64, param1 uint64) bool {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.GType)(param1)

	ret := C.g_value_type_transformable(cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_binding_get_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GBinding)(unsafe.Pointer(paramInstance))

	ret := C.g_binding_get_flags(cValueInstance)

	return (int)(ret)
}

func Fn_g_binding_get_source(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GBinding)(unsafe.Pointer(paramInstance))

	ret := C.g_binding_get_source(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_binding_get_source_property(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GBinding)(unsafe.Pointer(paramInstance))

	ret := C.g_binding_get_source_property(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_binding_get_target(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GBinding)(unsafe.Pointer(paramInstance))

	ret := C.g_binding_get_target(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_binding_get_target_property(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GBinding)(unsafe.Pointer(paramInstance))

	ret := C.g_binding_get_target_property(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : new : has varargs

// UNSUPPORTED : new_valist : has va_list

// UNSUPPORTED : new_with_properties : has non-string array param values

// UNSUPPORTED : newv : has non-string array param parameters

// UNSUPPORTED : add_toggle_ref : has callback

func Fn_g_object_add_weak_pointer(paramInstance unsafe.Pointer, param0 *unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))

	C.g_object_add_weak_pointer(cValueInstance, cValue0)
}

func Fn_g_object_bind_property(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 string, param3 int) unsafe.Pointer {
	cValueInstance := (C.gpointer)(paramInstance)

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gpointer)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.GBindingFlags)(param3)

	ret := C.g_object_bind_property(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : bind_property_full : has callback

func Fn_g_object_bind_property_with_closures(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 string, param3 int, param4 unsafe.Pointer, param5 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (C.gpointer)(paramInstance)

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gpointer)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.GBindingFlags)(param3)

	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))

	cValue5 := (*C.GClosure)(unsafe.Pointer(param5))

	ret := C.g_object_bind_property_with_closures(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : connect : has varargs

// UNSUPPORTED : disconnect : has varargs

// UNSUPPORTED : dup_data : has callback

// UNSUPPORTED : dup_qdata : has callback

func Fn_g_object_force_floating(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

	C.g_object_force_floating(cValueInstance)
}

func Fn_g_object_freeze_notify(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

	C.g_object_freeze_notify(cValueInstance)
}

// UNSUPPORTED : get : has varargs

func Fn_g_object_get_data(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_object_get_data(cValueInstance, cValue0)

	return (unsafe.Pointer)(ret)
}

func Fn_g_object_get_property(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	C.g_object_get_property(cValueInstance, cValue0, cValue1)
}

func Fn_g_object_get_qdata(paramInstance unsafe.Pointer, param0 uint32) unsafe.Pointer {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GQuark)(param0)

	ret := C.g_object_get_qdata(cValueInstance, cValue0)

	return (unsafe.Pointer)(ret)
}

// UNSUPPORTED : get_valist : has va_list

// UNSUPPORTED : getv : has non-string array param values

func Fn_g_object_is_floating(paramInstance unsafe.Pointer) bool {
	cValueInstance := (C.gpointer)(paramInstance)

	ret := C.g_object_is_floating(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_object_notify(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_object_notify(cValueInstance, cValue0)
}

func Fn_g_object_notify_by_pspec(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	C.g_object_notify_by_pspec(cValueInstance, cValue0)
}

func Fn_g_object_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (C.gpointer)(paramInstance)

	ret := C.g_object_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_object_ref_sink(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (C.gpointer)(paramInstance)

	ret := C.g_object_ref_sink(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : remove_toggle_ref : has callback

func Fn_g_object_remove_weak_pointer(paramInstance unsafe.Pointer, param0 *unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))

	C.g_object_remove_weak_pointer(cValueInstance, cValue0)
}

// UNSUPPORTED : replace_data : has callback

// UNSUPPORTED : replace_qdata : has callback

func Fn_g_object_run_dispose(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

	C.g_object_run_dispose(cValueInstance)
}

// UNSUPPORTED : set : has varargs

func Fn_g_object_set_data(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gpointer)(param1)

	C.g_object_set_data(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : set_data_full : has callback

func Fn_g_object_set_property(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	C.g_object_set_property(cValueInstance, cValue0, cValue1)
}

func Fn_g_object_set_qdata(paramInstance unsafe.Pointer, param0 uint32, param1 unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GQuark)(param0)

	cValue1 := (C.gpointer)(param1)

	C.g_object_set_qdata(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : set_qdata_full : has callback

// UNSUPPORTED : set_valist : has va_list

// UNSUPPORTED : setv : has non-string array param values

func Fn_g_object_steal_data(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_object_steal_data(cValueInstance, cValue0)

	return (unsafe.Pointer)(ret)
}

func Fn_g_object_steal_qdata(paramInstance unsafe.Pointer, param0 uint32) unsafe.Pointer {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GQuark)(param0)

	ret := C.g_object_steal_qdata(cValueInstance, cValue0)

	return (unsafe.Pointer)(ret)
}

func Fn_g_object_thaw_notify(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

	C.g_object_thaw_notify(cValueInstance)
}

func Fn_g_object_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (C.gpointer)(paramInstance)

	C.g_object_unref(cValueInstance)
}

func Fn_g_object_watch_closure(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	C.g_object_watch_closure(cValueInstance, cValue0)
}

// UNSUPPORTED : weak_ref : has callback

// UNSUPPORTED : weak_unref : has callback

func Fn_g_object_compat_control(param0 uint64, param1 unsafe.Pointer) uint64 {
	cValue0 := (C.gsize)(param0)

	cValue1 := (C.gpointer)(param1)

	ret := C.g_object_compat_control(cValue0, cValue1)

	return (uint64)(ret)
}

func Fn_g_object_interface_find_property(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_object_interface_find_property(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_object_interface_install_property(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (*C.GParamSpec)(unsafe.Pointer(param1))

	C.g_object_interface_install_property(cValue0, cValue1)
}

// UNSUPPORTED : interface_list_properties : has array return

func Fn_g_param_spec_get_blurb(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))

	ret := C.g_param_spec_get_blurb(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_param_spec_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))

	ret := C.g_param_spec_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_param_spec_get_nick(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))

	ret := C.g_param_spec_get_nick(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_param_spec_get_qdata(paramInstance unsafe.Pointer, param0 uint32) unsafe.Pointer {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GQuark)(param0)

	ret := C.g_param_spec_get_qdata(cValueInstance, cValue0)

	return (unsafe.Pointer)(ret)
}

func Fn_g_param_spec_get_redirect_target(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))

	ret := C.g_param_spec_get_redirect_target(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))

	ret := C.g_param_spec_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_ref_sink(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))

	ret := C.g_param_spec_ref_sink(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_param_spec_set_qdata(paramInstance unsafe.Pointer, param0 uint32, param1 unsafe.Pointer) {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GQuark)(param0)

	cValue1 := (C.gpointer)(param1)

	C.g_param_spec_set_qdata(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : set_qdata_full : has callback

func Fn_g_param_spec_sink(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))

	C.g_param_spec_sink(cValueInstance)
}

func Fn_g_param_spec_steal_qdata(paramInstance unsafe.Pointer, param0 uint32) unsafe.Pointer {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GQuark)(param0)

	ret := C.g_param_spec_steal_qdata(cValueInstance, cValue0)

	return (unsafe.Pointer)(ret)
}

func Fn_g_param_spec_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))

	C.g_param_spec_unref(cValueInstance)
}

func Fn_g_param_spec_internal(param0 uint64, param1 string, param2 string, param3 string, param4 int) unsafe.Pointer {
	cValue0 := (C.GType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	cValue4 := (C.GParamFlags)(param4)

	ret := C.g_param_spec_internal(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

func Fn_g_type_module_add_interface(paramInstance unsafe.Pointer, param0 uint64, param1 uint64, param2 unsafe.Pointer) {
	cValueInstance := (*C.GTypeModule)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	cValue1 := (C.GType)(param1)

	cValue2 := (*C.GInterfaceInfo)(unsafe.Pointer(param2))

	C.g_type_module_add_interface(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_type_module_register_enum(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) uint64 {
	cValueInstance := (*C.GTypeModule)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GEnumValue)(unsafe.Pointer(param1))

	ret := C.g_type_module_register_enum(cValueInstance, cValue0, cValue1)

	return (uint64)(ret)
}

func Fn_g_type_module_register_flags(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) uint64 {
	cValueInstance := (*C.GTypeModule)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GFlagsValue)(unsafe.Pointer(param1))

	ret := C.g_type_module_register_flags(cValueInstance, cValue0, cValue1)

	return (uint64)(ret)
}

func Fn_g_type_module_register_type(paramInstance unsafe.Pointer, param0 uint64, param1 string, param2 unsafe.Pointer, param3 int) uint64 {
	cValueInstance := (*C.GTypeModule)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GTypeInfo)(unsafe.Pointer(param2))

	cValue3 := (C.GTypeFlags)(param3)

	ret := C.g_type_module_register_type(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return (uint64)(ret)
}

func Fn_g_type_module_set_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GTypeModule)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_type_module_set_name(cValueInstance, cValue0)
}

func Fn_g_type_module_unuse(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTypeModule)(unsafe.Pointer(paramInstance))

	C.g_type_module_unuse(cValueInstance)
}

func Fn_g_type_module_use(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GTypeModule)(unsafe.Pointer(paramInstance))

	ret := C.g_type_module_use(cValueInstance)

	return toGoBool(ret)
}
