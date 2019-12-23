// Code generated - DO NOT EDIT.
// +build gobject_2.4

package gobject

import "unsafe"

// #include <glib-object.h>
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.TRUE
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

func Fn_g_boxed_copy(param0 uint64, param1 unsafe.Pointer) {
	cValue0 := (C.GType)(param0)
	cValue1 := (C.gconstpointer)(param1)

}

func Fn_g_boxed_free(param0 uint64, param1 *unsafe.Pointer) {
	cValue0 := (C.GType)(param0)
	cValue1 := (*C.gpointer)(unsafe.Pointer(param1))

}

// UNSUPPORTED : boxed_type_register_static : has callback

func Fn_g_cclosure_marshal_BOOLEAN__BOXED_BOXED(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

func Fn_g_cclosure_marshal_BOOLEAN__FLAGS(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

func Fn_g_cclosure_marshal_STRING__OBJECT_POINTER(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

func Fn_g_cclosure_marshal_VOID__BOOLEAN(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

func Fn_g_cclosure_marshal_VOID__BOXED(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

func Fn_g_cclosure_marshal_VOID__CHAR(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

func Fn_g_cclosure_marshal_VOID__DOUBLE(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

func Fn_g_cclosure_marshal_VOID__ENUM(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

func Fn_g_cclosure_marshal_VOID__FLAGS(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

func Fn_g_cclosure_marshal_VOID__FLOAT(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

func Fn_g_cclosure_marshal_VOID__INT(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

func Fn_g_cclosure_marshal_VOID__LONG(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

func Fn_g_cclosure_marshal_VOID__OBJECT(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

func Fn_g_cclosure_marshal_VOID__PARAM(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

func Fn_g_cclosure_marshal_VOID__POINTER(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

func Fn_g_cclosure_marshal_VOID__STRING(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

func Fn_g_cclosure_marshal_VOID__UCHAR(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

func Fn_g_cclosure_marshal_VOID__UINT(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

func Fn_g_cclosure_marshal_VOID__UINT_POINTER(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

func Fn_g_cclosure_marshal_VOID__ULONG(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

func Fn_g_cclosure_marshal_VOID__VOID(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint, param3 unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (C.guint)(param2)
	cValue3 := (*C.GValue)(unsafe.Pointer(param3))
	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))

}

// UNSUPPORTED : cclosure_new : has callback

// UNSUPPORTED : cclosure_new_object : has callback

// UNSUPPORTED : cclosure_new_object_swap : has callback

// UNSUPPORTED : cclosure_new_swap : has callback

func Fn_g_enum_complete_type_info(param0 uint64, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValue0 := (C.GType)(param0)
	cValue1 := (*C.GTypeInfo)(unsafe.Pointer(param1))
	cValue2 := (*C.GEnumValue)(unsafe.Pointer(param2))

}

func Fn_g_enum_get_value(param0 unsafe.Pointer, param1 int) {
	cValue0 := (*C.GEnumClass)(unsafe.Pointer(param0))
	cValue1 := (C.gint)(param1)

}

func Fn_g_enum_get_value_by_name(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GEnumClass)(unsafe.Pointer(param0))
	cValue1 := 42

}

func Fn_g_enum_get_value_by_nick(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GEnumClass)(unsafe.Pointer(param0))
	cValue1 := 42

}

func Fn_g_enum_register_static(param0 string, param1 unsafe.Pointer) {
	cValue0 := 42
	cValue1 := (*C.GEnumValue)(unsafe.Pointer(param1))

}

func Fn_g_flags_complete_type_info(param0 uint64, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValue0 := (C.GType)(param0)
	cValue1 := (*C.GTypeInfo)(unsafe.Pointer(param1))
	cValue2 := (*C.GFlagsValue)(unsafe.Pointer(param2))

}

func Fn_g_flags_get_first_value(param0 unsafe.Pointer, param1 uint) {
	cValue0 := (*C.GFlagsClass)(unsafe.Pointer(param0))
	cValue1 := (C.guint)(param1)

}

func Fn_g_flags_get_value_by_name(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GFlagsClass)(unsafe.Pointer(param0))
	cValue1 := 42

}

func Fn_g_flags_get_value_by_nick(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GFlagsClass)(unsafe.Pointer(param0))
	cValue1 := 42

}

func Fn_g_flags_register_static(param0 string, param1 unsafe.Pointer) {
	cValue0 := 42
	cValue1 := (*C.GFlagsValue)(unsafe.Pointer(param1))

}

func Fn_g_gtype_get_type() {

}

func Fn_g_param_spec_boolean(param0 string, param1 string, param2 string, param3 bool, param4 int) {
	cValue0 := 42
	cValue1 := 42
	cValue2 := 42
	cValue3 := toCBool(param3)
	cValue4 := (C.GParamFlags)(param4)

}

func Fn_g_param_spec_boxed(param0 string, param1 string, param2 string, param3 uint64, param4 int) {
	cValue0 := 42
	cValue1 := 42
	cValue2 := 42
	cValue3 := (C.GType)(param3)
	cValue4 := (C.GParamFlags)(param4)

}

func Fn_g_param_spec_char(param0 string, param1 string, param2 string, param3 int8, param4 int8, param5 int8, param6 int) {
	cValue0 := 42
	cValue1 := 42
	cValue2 := 42
	cValue3 := (C.gint8)(param3)
	cValue4 := (C.gint8)(param4)
	cValue5 := (C.gint8)(param5)
	cValue6 := (C.GParamFlags)(param6)

}

func Fn_g_param_spec_double(param0 string, param1 string, param2 string, param3 float64, param4 float64, param5 float64, param6 int) {
	cValue0 := 42
	cValue1 := 42
	cValue2 := 42
	cValue3 := (C.gdouble)(param3)
	cValue4 := (C.gdouble)(param4)
	cValue5 := (C.gdouble)(param5)
	cValue6 := (C.GParamFlags)(param6)

}

func Fn_g_param_spec_enum(param0 string, param1 string, param2 string, param3 uint64, param4 int, param5 int) {
	cValue0 := 42
	cValue1 := 42
	cValue2 := 42
	cValue3 := (C.GType)(param3)
	cValue4 := (C.gint)(param4)
	cValue5 := (C.GParamFlags)(param5)

}

func Fn_g_param_spec_flags(param0 string, param1 string, param2 string, param3 uint64, param4 uint, param5 int) {
	cValue0 := 42
	cValue1 := 42
	cValue2 := 42
	cValue3 := (C.GType)(param3)
	cValue4 := (C.guint)(param4)
	cValue5 := (C.GParamFlags)(param5)

}

func Fn_g_param_spec_float(param0 string, param1 string, param2 string, param3 float32, param4 float32, param5 float32, param6 int) {
	cValue0 := 42
	cValue1 := 42
	cValue2 := 42
	cValue3 := (C.gfloat)(param3)
	cValue4 := (C.gfloat)(param4)
	cValue5 := (C.gfloat)(param5)
	cValue6 := (C.GParamFlags)(param6)

}

func Fn_g_param_spec_int(param0 string, param1 string, param2 string, param3 int, param4 int, param5 int, param6 int) {
	cValue0 := 42
	cValue1 := 42
	cValue2 := 42
	cValue3 := (C.gint)(param3)
	cValue4 := (C.gint)(param4)
	cValue5 := (C.gint)(param5)
	cValue6 := (C.GParamFlags)(param6)

}

func Fn_g_param_spec_int64(param0 string, param1 string, param2 string, param3 int64, param4 int64, param5 int64, param6 int) {
	cValue0 := 42
	cValue1 := 42
	cValue2 := 42
	cValue3 := (C.gint64)(param3)
	cValue4 := (C.gint64)(param4)
	cValue5 := (C.gint64)(param5)
	cValue6 := (C.GParamFlags)(param6)

}

func Fn_g_param_spec_long(param0 string, param1 string, param2 string, param3 int64, param4 int64, param5 int64, param6 int) {
	cValue0 := 42
	cValue1 := 42
	cValue2 := 42
	cValue3 := (C.glong)(param3)
	cValue4 := (C.glong)(param4)
	cValue5 := (C.glong)(param5)
	cValue6 := (C.GParamFlags)(param6)

}

func Fn_g_param_spec_object(param0 string, param1 string, param2 string, param3 uint64, param4 int) {
	cValue0 := 42
	cValue1 := 42
	cValue2 := 42
	cValue3 := (C.GType)(param3)
	cValue4 := (C.GParamFlags)(param4)

}

func Fn_g_param_spec_override(param0 string, param1 unsafe.Pointer) {
	cValue0 := 42
	cValue1 := (*C.GParamSpec)(unsafe.Pointer(param1))

}

func Fn_g_param_spec_param(param0 string, param1 string, param2 string, param3 uint64, param4 int) {
	cValue0 := 42
	cValue1 := 42
	cValue2 := 42
	cValue3 := (C.GType)(param3)
	cValue4 := (C.GParamFlags)(param4)

}

func Fn_g_param_spec_pointer(param0 string, param1 string, param2 string, param3 int) {
	cValue0 := 42
	cValue1 := 42
	cValue2 := 42
	cValue3 := (C.GParamFlags)(param3)

}

func Fn_g_param_spec_pool_new(param0 bool) {
	cValue0 := toCBool(param0)

}

func Fn_g_param_spec_string(param0 string, param1 string, param2 string, param3 string, param4 int) {
	cValue0 := 42
	cValue1 := 42
	cValue2 := 42
	cValue3 := 42
	cValue4 := (C.GParamFlags)(param4)

}

func Fn_g_param_spec_uchar(param0 string, param1 string, param2 string, param3 uint8, param4 uint8, param5 uint8, param6 int) {
	cValue0 := 42
	cValue1 := 42
	cValue2 := 42
	cValue3 := (C.guint8)(param3)
	cValue4 := (C.guint8)(param4)
	cValue5 := (C.guint8)(param5)
	cValue6 := (C.GParamFlags)(param6)

}

func Fn_g_param_spec_uint(param0 string, param1 string, param2 string, param3 uint, param4 uint, param5 uint, param6 int) {
	cValue0 := 42
	cValue1 := 42
	cValue2 := 42
	cValue3 := (C.guint)(param3)
	cValue4 := (C.guint)(param4)
	cValue5 := (C.guint)(param5)
	cValue6 := (C.GParamFlags)(param6)

}

func Fn_g_param_spec_uint64(param0 string, param1 string, param2 string, param3 uint64, param4 uint64, param5 uint64, param6 int) {
	cValue0 := 42
	cValue1 := 42
	cValue2 := 42
	cValue3 := (C.guint64)(param3)
	cValue4 := (C.guint64)(param4)
	cValue5 := (C.guint64)(param5)
	cValue6 := (C.GParamFlags)(param6)

}

func Fn_g_param_spec_ulong(param0 string, param1 string, param2 string, param3 uint64, param4 uint64, param5 uint64, param6 int) {
	cValue0 := 42
	cValue1 := 42
	cValue2 := 42
	cValue3 := (C.gulong)(param3)
	cValue4 := (C.gulong)(param4)
	cValue5 := (C.gulong)(param5)
	cValue6 := (C.GParamFlags)(param6)

}

func Fn_g_param_spec_unichar(param0 string, param1 string, param2 string, param3 rune, param4 int) {
	cValue0 := 42
	cValue1 := 42
	cValue2 := 42
	cValue3 := (C.gunichar)(param3)
	cValue4 := (C.GParamFlags)(param4)

}

func Fn_g_param_spec_value_array(param0 string, param1 string, param2 string, param3 unsafe.Pointer, param4 int) {
	cValue0 := 42
	cValue1 := 42
	cValue2 := 42
	cValue3 := (*C.GParamSpec)(unsafe.Pointer(param3))
	cValue4 := (C.GParamFlags)(param4)

}

func Fn_g_param_type_register_static(param0 string, param1 unsafe.Pointer) {
	cValue0 := 42
	cValue1 := (*C.GParamSpecTypeInfo)(unsafe.Pointer(param1))

}

func Fn_g_param_value_convert(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 bool) {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (*C.GValue)(unsafe.Pointer(param2))
	cValue3 := toCBool(param3)

}

func Fn_g_param_value_defaults(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

}

func Fn_g_param_value_set_default(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

}

func Fn_g_param_value_validate(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

}

func Fn_g_param_values_cmp(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

}

func Fn_g_pointer_type_register_static(param0 string) {
	cValue0 := 42

}

func Fn_g_signal_accumulator_true_handled(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 *unsafe.Pointer) {
	cValue0 := (*C.GSignalInvocationHint)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))
	cValue2 := (*C.GValue)(unsafe.Pointer(param2))
	cValue3 := (*C.gpointer)(unsafe.Pointer(param3))

}

// UNSUPPORTED : signal_add_emission_hook : has callback

func Fn_g_signal_chain_from_overridden(param0 []Value, param1 unsafe.Pointer) {
	// has array param
}

// UNSUPPORTED : signal_chain_from_overridden_handler : has varargs

func Fn_g_signal_connect_closure(param0 *unsafe.Pointer, param1 string, param2 unsafe.Pointer, param3 bool) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))
	cValue1 := 42
	cValue2 := (*C.GClosure)(unsafe.Pointer(param2))
	cValue3 := toCBool(param3)

}

func Fn_g_signal_connect_closure_by_id(param0 *unsafe.Pointer, param1 uint, param2 uint32, param3 unsafe.Pointer, param4 bool) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))
	cValue1 := (C.guint)(param1)
	cValue2 := (C.GQuark)(param2)
	cValue3 := (*C.GClosure)(unsafe.Pointer(param3))
	cValue4 := toCBool(param4)

}

// UNSUPPORTED : signal_connect_data : has callback

// UNSUPPORTED : signal_connect_object : has callback

// UNSUPPORTED : signal_emit : has varargs

// UNSUPPORTED : signal_emit_by_name : has varargs

// UNSUPPORTED : signal_emit_valist : has va_list

func Fn_g_signal_emitv(param0 []Value, param1 uint, param2 uint32, param3 unsafe.Pointer) {
	// has array param
}

func Fn_g_signal_get_invocation_hint(param0 *unsafe.Pointer) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))

}

func Fn_g_signal_handler_block(param0 *unsafe.Pointer, param1 uint64) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))
	cValue1 := (C.gulong)(param1)

}

func Fn_g_signal_handler_disconnect(param0 *unsafe.Pointer, param1 uint64) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))
	cValue1 := (C.gulong)(param1)

}

func Fn_g_signal_handler_find(param0 *unsafe.Pointer, param1 int, param2 uint, param3 uint32, param4 unsafe.Pointer, param5 *unsafe.Pointer, param6 *unsafe.Pointer) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))
	cValue1 := (C.GSignalMatchType)(param1)
	cValue2 := (C.guint)(param2)
	cValue3 := (C.GQuark)(param3)
	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))
	cValue6 := (*C.gpointer)(unsafe.Pointer(param6))

}

func Fn_g_signal_handler_is_connected(param0 *unsafe.Pointer, param1 uint64) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))
	cValue1 := (C.gulong)(param1)

}

func Fn_g_signal_handler_unblock(param0 *unsafe.Pointer, param1 uint64) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))
	cValue1 := (C.gulong)(param1)

}

func Fn_g_signal_handlers_block_matched(param0 *unsafe.Pointer, param1 int, param2 uint, param3 uint32, param4 unsafe.Pointer, param5 *unsafe.Pointer, param6 *unsafe.Pointer) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))
	cValue1 := (C.GSignalMatchType)(param1)
	cValue2 := (C.guint)(param2)
	cValue3 := (C.GQuark)(param3)
	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))
	cValue6 := (*C.gpointer)(unsafe.Pointer(param6))

}

func Fn_g_signal_handlers_destroy(param0 *unsafe.Pointer) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))

}

func Fn_g_signal_handlers_disconnect_matched(param0 *unsafe.Pointer, param1 int, param2 uint, param3 uint32, param4 unsafe.Pointer, param5 *unsafe.Pointer, param6 *unsafe.Pointer) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))
	cValue1 := (C.GSignalMatchType)(param1)
	cValue2 := (C.guint)(param2)
	cValue3 := (C.GQuark)(param3)
	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))
	cValue6 := (*C.gpointer)(unsafe.Pointer(param6))

}

func Fn_g_signal_handlers_unblock_matched(param0 *unsafe.Pointer, param1 int, param2 uint, param3 uint32, param4 unsafe.Pointer, param5 *unsafe.Pointer, param6 *unsafe.Pointer) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))
	cValue1 := (C.GSignalMatchType)(param1)
	cValue2 := (C.guint)(param2)
	cValue3 := (C.GQuark)(param3)
	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))
	cValue5 := (*C.gpointer)(unsafe.Pointer(param5))
	cValue6 := (*C.gpointer)(unsafe.Pointer(param6))

}

func Fn_g_signal_has_handler_pending(param0 *unsafe.Pointer, param1 uint, param2 uint32, param3 bool) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))
	cValue1 := (C.guint)(param1)
	cValue2 := (C.GQuark)(param2)
	cValue3 := toCBool(param3)

}

func Fn_g_signal_list_ids(param0 uint64, param1 *uint) {
	cValue0 := (C.GType)(param0)
	cValue1 := (*C.guint)(unsafe.Pointer(param1))

}

func Fn_g_signal_lookup(param0 string, param1 uint64) {
	cValue0 := 42
	cValue1 := (C.GType)(param1)

}

func Fn_g_signal_name(param0 uint) {
	cValue0 := (C.guint)(param0)

}

// UNSUPPORTED : signal_new : has varargs

// UNSUPPORTED : signal_new_class_handler : has varargs

// UNSUPPORTED : signal_new_valist : has va_list

// UNSUPPORTED : signal_newv : has callback

func Fn_g_signal_override_class_closure(param0 uint, param1 uint64, param2 unsafe.Pointer) {
	cValue0 := (C.guint)(param0)
	cValue1 := (C.GType)(param1)
	cValue2 := (*C.GClosure)(unsafe.Pointer(param2))

}

// UNSUPPORTED : signal_override_class_handler : has callback

func Fn_g_signal_parse_name(param0 string, param1 uint64, param2 *uint, param3 *uint32, param4 bool) {
	cValue0 := 42
	cValue1 := (C.GType)(param1)
	cValue2 := (*C.guint)(unsafe.Pointer(param2))
	cValue3 := (*C.GQuark)(unsafe.Pointer(param3))
	cValue4 := toCBool(param4)

}

func Fn_g_signal_query(param0 uint, param1 unsafe.Pointer) {
	cValue0 := (C.guint)(param0)
	cValue1 := (*C.GSignalQuery)(unsafe.Pointer(param1))

}

func Fn_g_signal_remove_emission_hook(param0 uint, param1 uint64) {
	cValue0 := (C.guint)(param0)
	cValue1 := (C.gulong)(param1)

}

func Fn_g_signal_stop_emission(param0 *unsafe.Pointer, param1 uint, param2 uint32) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))
	cValue1 := (C.guint)(param1)
	cValue2 := (C.GQuark)(param2)

}

func Fn_g_signal_stop_emission_by_name(param0 *unsafe.Pointer, param1 string) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))
	cValue1 := 42

}

func Fn_g_signal_type_cclosure_new(param0 uint64, param1 uint) {
	cValue0 := (C.GType)(param0)
	cValue1 := (C.guint)(param1)

}

func Fn_g_source_set_closure(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GSource)(unsafe.Pointer(param0))
	cValue1 := (*C.GClosure)(unsafe.Pointer(param1))

}

func Fn_g_source_set_dummy_callback(param0 unsafe.Pointer) {
	cValue0 := (*C.GSource)(unsafe.Pointer(param0))

}

func Fn_g_strdup_value_contents(param0 unsafe.Pointer) {
	cValue0 := (*C.GValue)(unsafe.Pointer(param0))

}

// UNSUPPORTED : type_add_class_cache_func : has callback

func Fn_g_type_add_instance_private(param0 uint64, param1 uint64) {
	cValue0 := (C.GType)(param0)
	cValue1 := (C.gsize)(param1)

}

// UNSUPPORTED : type_add_interface_check : has callback

func Fn_g_type_add_interface_dynamic(param0 uint64, param1 uint64, param2 unsafe.Pointer) {
	cValue0 := (C.GType)(param0)
	cValue1 := (C.GType)(param1)
	cValue2 := (*C.GTypePlugin)(unsafe.Pointer(param2))

}

func Fn_g_type_add_interface_static(param0 uint64, param1 uint64, param2 unsafe.Pointer) {
	cValue0 := (C.GType)(param0)
	cValue1 := (C.GType)(param1)
	cValue2 := (*C.GInterfaceInfo)(unsafe.Pointer(param2))

}

func Fn_g_type_check_class_cast(param0 unsafe.Pointer, param1 uint64) {
	cValue0 := (*C.GTypeClass)(unsafe.Pointer(param0))
	cValue1 := (C.GType)(param1)

}

func Fn_g_type_check_class_is_a(param0 unsafe.Pointer, param1 uint64) {
	cValue0 := (*C.GTypeClass)(unsafe.Pointer(param0))
	cValue1 := (C.GType)(param1)

}

func Fn_g_type_check_instance(param0 unsafe.Pointer) {
	cValue0 := (*C.GTypeInstance)(unsafe.Pointer(param0))

}

func Fn_g_type_check_instance_cast(param0 unsafe.Pointer, param1 uint64) {
	cValue0 := (*C.GTypeInstance)(unsafe.Pointer(param0))
	cValue1 := (C.GType)(param1)

}

func Fn_g_type_check_instance_is_a(param0 unsafe.Pointer, param1 uint64) {
	cValue0 := (*C.GTypeInstance)(unsafe.Pointer(param0))
	cValue1 := (C.GType)(param1)

}

func Fn_g_type_check_instance_is_fundamentally_a(param0 unsafe.Pointer, param1 uint64) {
	cValue0 := (*C.GTypeInstance)(unsafe.Pointer(param0))
	cValue1 := (C.GType)(param1)

}

func Fn_g_type_check_is_value_type(param0 uint64) {
	cValue0 := (C.GType)(param0)

}

func Fn_g_type_check_value(param0 unsafe.Pointer) {
	cValue0 := (*C.GValue)(unsafe.Pointer(param0))

}

func Fn_g_type_check_value_holds(param0 unsafe.Pointer, param1 uint64) {
	cValue0 := (*C.GValue)(unsafe.Pointer(param0))
	cValue1 := (C.GType)(param1)

}

func Fn_g_type_children(param0 uint64, param1 *uint) {
	cValue0 := (C.GType)(param0)
	cValue1 := (*C.guint)(unsafe.Pointer(param1))

}

func Fn_g_type_class_adjust_private_offset(param0 *unsafe.Pointer, param1 *int) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))
	cValue1 := (*C.gint)(unsafe.Pointer(param1))

}

func Fn_g_type_class_peek(param0 uint64) {
	cValue0 := (C.GType)(param0)

}

func Fn_g_type_class_peek_static(param0 uint64) {
	cValue0 := (C.GType)(param0)

}

func Fn_g_type_class_ref(param0 uint64) {
	cValue0 := (C.GType)(param0)

}

func Fn_g_type_create_instance(param0 uint64) {
	cValue0 := (C.GType)(param0)

}

func Fn_g_type_default_interface_peek(param0 uint64) {
	cValue0 := (C.GType)(param0)

}

func Fn_g_type_default_interface_ref(param0 uint64) {
	cValue0 := (C.GType)(param0)

}

func Fn_g_type_default_interface_unref(param0 *unsafe.Pointer) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))

}

func Fn_g_type_depth(param0 uint64) {
	cValue0 := (C.GType)(param0)

}

func Fn_g_type_free_instance(param0 unsafe.Pointer) {
	cValue0 := (*C.GTypeInstance)(unsafe.Pointer(param0))

}

func Fn_g_type_from_name(param0 string) {
	cValue0 := 42

}

func Fn_g_type_fundamental(param0 uint64) {
	cValue0 := (C.GType)(param0)

}

func Fn_g_type_fundamental_next() {

}

func Fn_g_type_get_plugin(param0 uint64) {
	cValue0 := (C.GType)(param0)

}

func Fn_g_type_get_qdata(param0 uint64, param1 uint32) {
	cValue0 := (C.GType)(param0)
	cValue1 := (C.GQuark)(param1)

}

func Fn_g_type_init() {

}

func Fn_g_type_init_with_debug_flags(param0 int) {
	cValue0 := (C.GTypeDebugFlags)(param0)

}

func Fn_g_type_interface_add_prerequisite(param0 uint64, param1 uint64) {
	cValue0 := (C.GType)(param0)
	cValue1 := (C.GType)(param1)

}

func Fn_g_type_interface_get_plugin(param0 uint64, param1 uint64) {
	cValue0 := (C.GType)(param0)
	cValue1 := (C.GType)(param1)

}

func Fn_g_type_interface_peek(param0 *unsafe.Pointer, param1 uint64) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))
	cValue1 := (C.GType)(param1)

}

func Fn_g_type_interface_prerequisites(param0 uint64, param1 *uint) {
	cValue0 := (C.GType)(param0)
	cValue1 := (*C.guint)(unsafe.Pointer(param1))

}

func Fn_g_type_interfaces(param0 uint64, param1 *uint) {
	cValue0 := (C.GType)(param0)
	cValue1 := (*C.guint)(unsafe.Pointer(param1))

}

func Fn_g_type_is_a(param0 uint64, param1 uint64) {
	cValue0 := (C.GType)(param0)
	cValue1 := (C.GType)(param1)

}

func Fn_g_type_name(param0 uint64) {
	cValue0 := (C.GType)(param0)

}

func Fn_g_type_name_from_class(param0 unsafe.Pointer) {
	cValue0 := (*C.GTypeClass)(unsafe.Pointer(param0))

}

func Fn_g_type_name_from_instance(param0 unsafe.Pointer) {
	cValue0 := (*C.GTypeInstance)(unsafe.Pointer(param0))

}

func Fn_g_type_next_base(param0 uint64, param1 uint64) {
	cValue0 := (C.GType)(param0)
	cValue1 := (C.GType)(param1)

}

func Fn_g_type_parent(param0 uint64) {
	cValue0 := (C.GType)(param0)

}

func Fn_g_type_qname(param0 uint64) {
	cValue0 := (C.GType)(param0)

}

func Fn_g_type_query(param0 uint64, param1 unsafe.Pointer) {
	cValue0 := (C.GType)(param0)
	cValue1 := (*C.GTypeQuery)(unsafe.Pointer(param1))

}

func Fn_g_type_register_dynamic(param0 uint64, param1 string, param2 unsafe.Pointer, param3 int) {
	cValue0 := (C.GType)(param0)
	cValue1 := 42
	cValue2 := (*C.GTypePlugin)(unsafe.Pointer(param2))
	cValue3 := (C.GTypeFlags)(param3)

}

func Fn_g_type_register_fundamental(param0 uint64, param1 string, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int) {
	cValue0 := (C.GType)(param0)
	cValue1 := 42
	cValue2 := (*C.GTypeInfo)(unsafe.Pointer(param2))
	cValue3 := (*C.GTypeFundamentalInfo)(unsafe.Pointer(param3))
	cValue4 := (C.GTypeFlags)(param4)

}

func Fn_g_type_register_static(param0 uint64, param1 string, param2 unsafe.Pointer, param3 int) {
	cValue0 := (C.GType)(param0)
	cValue1 := 42
	cValue2 := (*C.GTypeInfo)(unsafe.Pointer(param2))
	cValue3 := (C.GTypeFlags)(param3)

}

// UNSUPPORTED : type_register_static_simple : has callback

// UNSUPPORTED : type_remove_class_cache_func : has callback

// UNSUPPORTED : type_remove_interface_check : has callback

func Fn_g_type_set_qdata(param0 uint64, param1 uint32, param2 *unsafe.Pointer) {
	cValue0 := (C.GType)(param0)
	cValue1 := (C.GQuark)(param1)
	cValue2 := (*C.gpointer)(unsafe.Pointer(param2))

}

func Fn_g_type_test_flags(param0 uint64, param1 uint) {
	cValue0 := (C.GType)(param0)
	cValue1 := (C.guint)(param1)

}

func Fn_g_type_value_table_peek(param0 uint64) {
	cValue0 := (C.GType)(param0)

}

// UNSUPPORTED : value_register_transform_func : has callback

func Fn_g_value_type_compatible(param0 uint64, param1 uint64) {
	cValue0 := (C.GType)(param0)
	cValue1 := (C.GType)(param1)

}

func Fn_g_value_type_transformable(param0 uint64, param1 uint64) {
	cValue0 := (C.GType)(param0)
	cValue1 := (C.GType)(param1)

}

// UNSUPPORTED : new : has varargs

// UNSUPPORTED : new_valist : has va_list

func Fn_g_object_newv(param0 uint64, param1 uint, param2 []Parameter) {
	// has array param
}

// UNSUPPORTED : add_toggle_ref : has callback

func Fn_g_object_add_weak_pointer(paramInstance unsafe.Pointer, param0 *unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))

}

// UNSUPPORTED : bind_property_full : has callback

// UNSUPPORTED : connect : has varargs

// UNSUPPORTED : disconnect : has varargs

// UNSUPPORTED : dup_data : has callback

// UNSUPPORTED : dup_qdata : has callback

func Fn_g_object_freeze_notify(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

}

// UNSUPPORTED : get : has varargs

func Fn_g_object_get_data(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))
	cValue0 := 42

}

func Fn_g_object_get_property(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))
	cValue0 := 42
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

}

func Fn_g_object_get_qdata(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GQuark)(param0)

}

// UNSUPPORTED : get_valist : has va_list

func Fn_g_object_notify(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))
	cValue0 := 42

}

func Fn_g_object_ref(paramInstance *unsafe.Pointer) {
	cValueInstance := (*C.gpointer)(unsafe.Pointer(paramInstance))

}

// UNSUPPORTED : remove_toggle_ref : has callback

func Fn_g_object_remove_weak_pointer(paramInstance unsafe.Pointer, param0 *unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))

}

// UNSUPPORTED : replace_data : has callback

// UNSUPPORTED : replace_qdata : has callback

func Fn_g_object_run_dispose(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

}

// UNSUPPORTED : set : has varargs

func Fn_g_object_set_data(paramInstance unsafe.Pointer, param0 string, param1 *unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))
	cValue0 := 42
	cValue1 := (*C.gpointer)(unsafe.Pointer(param1))

}

// UNSUPPORTED : set_data_full : has callback

func Fn_g_object_set_property(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))
	cValue0 := 42
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

}

func Fn_g_object_set_qdata(paramInstance unsafe.Pointer, param0 uint32, param1 *unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GQuark)(param0)
	cValue1 := (*C.gpointer)(unsafe.Pointer(param1))

}

// UNSUPPORTED : set_qdata_full : has callback

// UNSUPPORTED : set_valist : has va_list

func Fn_g_object_steal_data(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))
	cValue0 := 42

}

func Fn_g_object_steal_qdata(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GQuark)(param0)

}

func Fn_g_object_thaw_notify(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))

}

func Fn_g_object_unref(paramInstance *unsafe.Pointer) {
	cValueInstance := (*C.gpointer)(unsafe.Pointer(paramInstance))

}

func Fn_g_object_watch_closure(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GObject)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

}

// UNSUPPORTED : weak_ref : has callback

// UNSUPPORTED : weak_unref : has callback

func Fn_g_object_compat_control(param0 uint64, param1 *unsafe.Pointer) {
	cValue0 := (C.gsize)(param0)
	cValue1 := (*C.gpointer)(unsafe.Pointer(param1))

}

func Fn_g_object_interface_find_property(param0 *unsafe.Pointer, param1 string) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))
	cValue1 := 42

}

func Fn_g_object_interface_install_property(param0 *unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))
	cValue1 := (*C.GParamSpec)(unsafe.Pointer(param1))

}

func Fn_g_object_interface_list_properties(param0 *unsafe.Pointer, param1 *uint) {
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))
	cValue1 := (*C.guint)(unsafe.Pointer(param1))

}

func Fn_g_param_spec_get_blurb(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))

}

func Fn_g_param_spec_get_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))

}

func Fn_g_param_spec_get_nick(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))

}

func Fn_g_param_spec_get_qdata(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GQuark)(param0)

}

func Fn_g_param_spec_get_redirect_target(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))

}

func Fn_g_param_spec_ref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))

}

func Fn_g_param_spec_set_qdata(paramInstance unsafe.Pointer, param0 uint32, param1 *unsafe.Pointer) {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GQuark)(param0)
	cValue1 := (*C.gpointer)(unsafe.Pointer(param1))

}

// UNSUPPORTED : set_qdata_full : has callback

func Fn_g_param_spec_sink(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))

}

func Fn_g_param_spec_steal_qdata(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GQuark)(param0)

}

func Fn_g_param_spec_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GParamSpec)(unsafe.Pointer(paramInstance))

}

func Fn_g_param_spec_internal(param0 uint64, param1 string, param2 string, param3 string, param4 int) {
	cValue0 := (C.GType)(param0)
	cValue1 := 42
	cValue2 := 42
	cValue3 := 42
	cValue4 := (C.GParamFlags)(param4)

}

func Fn_g_type_module_add_interface(paramInstance unsafe.Pointer, param0 uint64, param1 uint64, param2 unsafe.Pointer) {
	cValueInstance := (*C.GTypeModule)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GType)(param0)
	cValue1 := (C.GType)(param1)
	cValue2 := (*C.GInterfaceInfo)(unsafe.Pointer(param2))

}

func Fn_g_type_module_register_type(paramInstance unsafe.Pointer, param0 uint64, param1 string, param2 unsafe.Pointer, param3 int) {
	cValueInstance := (*C.GTypeModule)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GType)(param0)
	cValue1 := 42
	cValue2 := (*C.GTypeInfo)(unsafe.Pointer(param2))
	cValue3 := (C.GTypeFlags)(param3)

}

func Fn_g_type_module_set_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GTypeModule)(unsafe.Pointer(paramInstance))
	cValue0 := 42

}

func Fn_g_type_module_unuse(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTypeModule)(unsafe.Pointer(paramInstance))

}

func Fn_g_type_module_use(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTypeModule)(unsafe.Pointer(paramInstance))

}
