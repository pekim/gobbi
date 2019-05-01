// This is a generated file - DO NOT EDIT

package gobject

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Unsupported : alias has no type generator for ClosureMarshal, GClosureMarshal

// Unsupported : alias has no type generator for VaClosureMarshal, GVaClosureMarshal

// InitiallyUnowned is a wrapper around the C record GInitiallyUnowned.
type InitiallyUnowned struct {
	native *C.GInitiallyUnowned
	// All fields are private
}

func InitiallyUnownedNewFromC(u unsafe.Pointer) *InitiallyUnowned {
	c := (*C.GInitiallyUnowned)(u)
	if c == nil {
		return nil
	}

	g := &InitiallyUnowned{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *InitiallyUnowned) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *InitiallyUnowned) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object is a wrapper around the C record GObject.
type Object struct {
	native *C.GObject
	// All fields are private
}

func ObjectNewFromC(u unsafe.Pointer) *Object {
	c := (*C.GObject)(u)
	if c == nil {
		return nil
	}

	g := &Object{native: c}

	return g
}

func (recv *Object) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpec is a wrapper around the C record GParamSpec.
type ParamSpec struct {
	native *C.GParamSpec
	// g_type_instance : record
	Name      string
	Flags     ParamFlags
	ValueType Type
	OwnerType Type
	// Private : _nick
	// Private : _blurb
	// Private : qdata
	// Private : ref_count
	// Private : param_id
}

func ParamSpecNewFromC(u unsafe.Pointer) *ParamSpec {
	c := (*C.GParamSpec)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpec{
		Flags:     (ParamFlags)(c.flags),
		Name:      C.GoString(c.name),
		OwnerType: (Type)(c.owner_type),
		ValueType: (Type)(c.value_type),
		native:    c,
	}

	return g
}

func (recv *ParamSpec) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecBoolean is a wrapper around the C record GParamSpecBoolean.
type ParamSpecBoolean struct {
	native *C.GParamSpecBoolean
	// parent_instance : record
	DefaultValue bool
}

func ParamSpecBooleanNewFromC(u unsafe.Pointer) *ParamSpecBoolean {
	c := (*C.GParamSpecBoolean)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecBoolean{
		DefaultValue: c.default_value == C.TRUE,
		native:       c,
	}

	return g
}

func (recv *ParamSpecBoolean) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecBoxed is a wrapper around the C record GParamSpecBoxed.
type ParamSpecBoxed struct {
	native *C.GParamSpecBoxed
	// parent_instance : record
}

func ParamSpecBoxedNewFromC(u unsafe.Pointer) *ParamSpecBoxed {
	c := (*C.GParamSpecBoxed)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecBoxed{native: c}

	return g
}

func (recv *ParamSpecBoxed) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecChar is a wrapper around the C record GParamSpecChar.
type ParamSpecChar struct {
	native *C.GParamSpecChar
	// parent_instance : record
	Minimum      int8
	Maximum      int8
	DefaultValue int8
}

func ParamSpecCharNewFromC(u unsafe.Pointer) *ParamSpecChar {
	c := (*C.GParamSpecChar)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecChar{
		DefaultValue: (int8)(c.default_value),
		Maximum:      (int8)(c.maximum),
		Minimum:      (int8)(c.minimum),
		native:       c,
	}

	return g
}

func (recv *ParamSpecChar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecDouble is a wrapper around the C record GParamSpecDouble.
type ParamSpecDouble struct {
	native *C.GParamSpecDouble
	// parent_instance : record
	Minimum      float64
	Maximum      float64
	DefaultValue float64
	Epsilon      float64
}

func ParamSpecDoubleNewFromC(u unsafe.Pointer) *ParamSpecDouble {
	c := (*C.GParamSpecDouble)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecDouble{
		DefaultValue: (float64)(c.default_value),
		Epsilon:      (float64)(c.epsilon),
		Maximum:      (float64)(c.maximum),
		Minimum:      (float64)(c.minimum),
		native:       c,
	}

	return g
}

func (recv *ParamSpecDouble) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecEnum is a wrapper around the C record GParamSpecEnum.
type ParamSpecEnum struct {
	native *C.GParamSpecEnum
	// parent_instance : record
	// enum_class : record
	DefaultValue int32
}

func ParamSpecEnumNewFromC(u unsafe.Pointer) *ParamSpecEnum {
	c := (*C.GParamSpecEnum)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecEnum{
		DefaultValue: (int32)(c.default_value),
		native:       c,
	}

	return g
}

func (recv *ParamSpecEnum) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecFlags is a wrapper around the C record GParamSpecFlags.
type ParamSpecFlags struct {
	native *C.GParamSpecFlags
	// parent_instance : record
	// flags_class : record
	DefaultValue uint32
}

func ParamSpecFlagsNewFromC(u unsafe.Pointer) *ParamSpecFlags {
	c := (*C.GParamSpecFlags)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecFlags{
		DefaultValue: (uint32)(c.default_value),
		native:       c,
	}

	return g
}

func (recv *ParamSpecFlags) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecFloat is a wrapper around the C record GParamSpecFloat.
type ParamSpecFloat struct {
	native *C.GParamSpecFloat
	// parent_instance : record
	Minimum      float32
	Maximum      float32
	DefaultValue float32
	Epsilon      float32
}

func ParamSpecFloatNewFromC(u unsafe.Pointer) *ParamSpecFloat {
	c := (*C.GParamSpecFloat)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecFloat{
		DefaultValue: (float32)(c.default_value),
		Epsilon:      (float32)(c.epsilon),
		Maximum:      (float32)(c.maximum),
		Minimum:      (float32)(c.minimum),
		native:       c,
	}

	return g
}

func (recv *ParamSpecFloat) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecInt is a wrapper around the C record GParamSpecInt.
type ParamSpecInt struct {
	native *C.GParamSpecInt
	// parent_instance : record
	Minimum      int32
	Maximum      int32
	DefaultValue int32
}

func ParamSpecIntNewFromC(u unsafe.Pointer) *ParamSpecInt {
	c := (*C.GParamSpecInt)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecInt{
		DefaultValue: (int32)(c.default_value),
		Maximum:      (int32)(c.maximum),
		Minimum:      (int32)(c.minimum),
		native:       c,
	}

	return g
}

func (recv *ParamSpecInt) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecInt64 is a wrapper around the C record GParamSpecInt64.
type ParamSpecInt64 struct {
	native *C.GParamSpecInt64
	// parent_instance : record
	Minimum      int64
	Maximum      int64
	DefaultValue int64
}

func ParamSpecInt64NewFromC(u unsafe.Pointer) *ParamSpecInt64 {
	c := (*C.GParamSpecInt64)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecInt64{
		DefaultValue: (int64)(c.default_value),
		Maximum:      (int64)(c.maximum),
		Minimum:      (int64)(c.minimum),
		native:       c,
	}

	return g
}

func (recv *ParamSpecInt64) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecLong is a wrapper around the C record GParamSpecLong.
type ParamSpecLong struct {
	native *C.GParamSpecLong
	// parent_instance : record
	Minimum      int64
	Maximum      int64
	DefaultValue int64
}

func ParamSpecLongNewFromC(u unsafe.Pointer) *ParamSpecLong {
	c := (*C.GParamSpecLong)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecLong{
		DefaultValue: (int64)(c.default_value),
		Maximum:      (int64)(c.maximum),
		Minimum:      (int64)(c.minimum),
		native:       c,
	}

	return g
}

func (recv *ParamSpecLong) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecObject is a wrapper around the C record GParamSpecObject.
type ParamSpecObject struct {
	native *C.GParamSpecObject
	// parent_instance : record
}

func ParamSpecObjectNewFromC(u unsafe.Pointer) *ParamSpecObject {
	c := (*C.GParamSpecObject)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecObject{native: c}

	return g
}

func (recv *ParamSpecObject) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecParam is a wrapper around the C record GParamSpecParam.
type ParamSpecParam struct {
	native *C.GParamSpecParam
	// parent_instance : record
}

func ParamSpecParamNewFromC(u unsafe.Pointer) *ParamSpecParam {
	c := (*C.GParamSpecParam)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecParam{native: c}

	return g
}

func (recv *ParamSpecParam) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecPointer is a wrapper around the C record GParamSpecPointer.
type ParamSpecPointer struct {
	native *C.GParamSpecPointer
	// parent_instance : record
}

func ParamSpecPointerNewFromC(u unsafe.Pointer) *ParamSpecPointer {
	c := (*C.GParamSpecPointer)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecPointer{native: c}

	return g
}

func (recv *ParamSpecPointer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecString is a wrapper around the C record GParamSpecString.
type ParamSpecString struct {
	native *C.GParamSpecString
	// parent_instance : record
	DefaultValue string
	CsetFirst    string
	CsetNth      string
	Substitutor  rune
	// Bitfield not supported :  1 null_fold_if_empty
	// Bitfield not supported :  1 ensure_non_null
}

func ParamSpecStringNewFromC(u unsafe.Pointer) *ParamSpecString {
	c := (*C.GParamSpecString)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecString{
		CsetFirst:    C.GoString(c.cset_first),
		CsetNth:      C.GoString(c.cset_nth),
		DefaultValue: C.GoString(c.default_value),
		Substitutor:  (rune)(c.substitutor),
		native:       c,
	}

	return g
}

func (recv *ParamSpecString) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecUChar is a wrapper around the C record GParamSpecUChar.
type ParamSpecUChar struct {
	native *C.GParamSpecUChar
	// parent_instance : record
	Minimum      uint8
	Maximum      uint8
	DefaultValue uint8
}

func ParamSpecUCharNewFromC(u unsafe.Pointer) *ParamSpecUChar {
	c := (*C.GParamSpecUChar)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecUChar{
		DefaultValue: (uint8)(c.default_value),
		Maximum:      (uint8)(c.maximum),
		Minimum:      (uint8)(c.minimum),
		native:       c,
	}

	return g
}

func (recv *ParamSpecUChar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecUInt is a wrapper around the C record GParamSpecUInt.
type ParamSpecUInt struct {
	native *C.GParamSpecUInt
	// parent_instance : record
	Minimum      uint32
	Maximum      uint32
	DefaultValue uint32
}

func ParamSpecUIntNewFromC(u unsafe.Pointer) *ParamSpecUInt {
	c := (*C.GParamSpecUInt)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecUInt{
		DefaultValue: (uint32)(c.default_value),
		Maximum:      (uint32)(c.maximum),
		Minimum:      (uint32)(c.minimum),
		native:       c,
	}

	return g
}

func (recv *ParamSpecUInt) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecUInt64 is a wrapper around the C record GParamSpecUInt64.
type ParamSpecUInt64 struct {
	native *C.GParamSpecUInt64
	// parent_instance : record
	Minimum      uint64
	Maximum      uint64
	DefaultValue uint64
}

func ParamSpecUInt64NewFromC(u unsafe.Pointer) *ParamSpecUInt64 {
	c := (*C.GParamSpecUInt64)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecUInt64{
		DefaultValue: (uint64)(c.default_value),
		Maximum:      (uint64)(c.maximum),
		Minimum:      (uint64)(c.minimum),
		native:       c,
	}

	return g
}

func (recv *ParamSpecUInt64) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecULong is a wrapper around the C record GParamSpecULong.
type ParamSpecULong struct {
	native *C.GParamSpecULong
	// parent_instance : record
	Minimum      uint64
	Maximum      uint64
	DefaultValue uint64
}

func ParamSpecULongNewFromC(u unsafe.Pointer) *ParamSpecULong {
	c := (*C.GParamSpecULong)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecULong{
		DefaultValue: (uint64)(c.default_value),
		Maximum:      (uint64)(c.maximum),
		Minimum:      (uint64)(c.minimum),
		native:       c,
	}

	return g
}

func (recv *ParamSpecULong) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecUnichar is a wrapper around the C record GParamSpecUnichar.
type ParamSpecUnichar struct {
	native *C.GParamSpecUnichar
	// parent_instance : record
	DefaultValue rune
}

func ParamSpecUnicharNewFromC(u unsafe.Pointer) *ParamSpecUnichar {
	c := (*C.GParamSpecUnichar)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecUnichar{
		DefaultValue: (rune)(c.default_value),
		native:       c,
	}

	return g
}

func (recv *ParamSpecUnichar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecValueArray is a wrapper around the C record GParamSpecValueArray.
type ParamSpecValueArray struct {
	native *C.GParamSpecValueArray
	// parent_instance : record
	// element_spec : record
	FixedNElements uint32
}

func ParamSpecValueArrayNewFromC(u unsafe.Pointer) *ParamSpecValueArray {
	c := (*C.GParamSpecValueArray)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecValueArray{
		FixedNElements: (uint32)(c.fixed_n_elements),
		native:         c,
	}

	return g
}

func (recv *ParamSpecValueArray) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : GTypeModule

// Unsupported : type GLib.Type for TYPE_FLAG_RESERVED_ID_BIT

// Blacklisted : VALUE_COLLECT_FORMAT_MAX_LENGTH

// Unsupported : g_boxed_copy : unsupported parameter src_boxed : no type generator for gpointer (gconstpointer) for param src_boxed

// Unsupported : g_boxed_free : unsupported parameter boxed : no type generator for gpointer (gpointer) for param boxed

// Unsupported : g_boxed_type_register_static : unsupported parameter boxed_copy : no type generator for BoxedCopyFunc (GBoxedCopyFunc) for param boxed_copy

// Unsupported : g_cclosure_marshal_BOOLEAN__BOXED_BOXED : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_BOOLEAN__FLAGS : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_STRING__OBJECT_POINTER : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__BOOLEAN : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__BOXED : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__CHAR : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__DOUBLE : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__ENUM : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__FLAGS : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__FLOAT : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__INT : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__LONG : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__OBJECT : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__PARAM : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__POINTER : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__STRING : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__UCHAR : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__UINT : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__UINT_POINTER : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__ULONG : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__VARIANT : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_marshal_VOID__VOID : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint

// Unsupported : g_cclosure_new : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_object : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_object_swap : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_swap : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_enum_get_value : return type :

// Unsupported : g_enum_get_value_by_name : return type :

// Unsupported : g_enum_get_value_by_nick : return type :

// Unsupported : g_enum_register_static : return type :

// Unsupported : g_flags_get_first_value : return type :

// Unsupported : g_flags_get_value_by_name : return type :

// Unsupported : g_flags_get_value_by_nick : return type :

// Unsupported : g_flags_register_static : return type :

// Unsupported : g_gtype_get_type : return type :

// Unsupported : g_param_spec_boolean : return type :

// Unsupported : g_param_spec_boxed : return type :

// Unsupported : g_param_spec_char : return type :

// Unsupported : g_param_spec_double : return type :

// Unsupported : g_param_spec_enum : return type :

// Unsupported : g_param_spec_flags : return type :

// Unsupported : g_param_spec_float : return type :

// Unsupported : g_param_spec_int : return type :

// Unsupported : g_param_spec_int64 : return type :

// Unsupported : g_param_spec_long : return type :

// Unsupported : g_param_spec_object : return type :

// Unsupported : g_param_spec_param : return type :

// Unsupported : g_param_spec_pointer : return type :

// Unsupported : g_param_spec_pool_new : return type :

// Unsupported : g_param_spec_string : return type :

// Unsupported : g_param_spec_uchar : return type :

// Unsupported : g_param_spec_uint : return type :

// Unsupported : g_param_spec_uint64 : return type :

// Unsupported : g_param_spec_ulong : return type :

// Unsupported : g_param_spec_unichar : return type :

// Unsupported : g_param_spec_value_array : return type :

// Unsupported : g_param_type_register_static : return type :

// Unsupported : g_param_value_convert : return type :

// Unsupported : g_param_value_defaults : return type :

// Unsupported : g_param_value_validate : return type :

// Unsupported : g_pointer_type_register_static : return type :

// Unsupported : g_signal_add_emission_hook : unsupported parameter hook_func : no type generator for SignalEmissionHook (GSignalEmissionHook) for param hook_func

// Unsupported : g_signal_chain_from_overridden : unsupported parameter instance_and_params :

// Unsupported : g_signal_connect_data : unsupported parameter c_handler : no type generator for Callback (GCallback) for param c_handler

// Unsupported : g_signal_connect_object : unsupported parameter c_handler : no type generator for Callback (GCallback) for param c_handler

// Unsupported : g_signal_emit : unsupported parameter ... : varargs

// Unsupported : g_signal_emit_by_name : unsupported parameter ... : varargs

// Unsupported : g_signal_emit_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Unsupported : g_signal_emitv : unsupported parameter instance_and_params :

// Unsupported : g_signal_get_invocation_hint : return type :

// Unsupported : g_signal_handler_find : unsupported parameter func : no type generator for gpointer (gpointer) for param func

// Unsupported : g_signal_handler_is_connected : return type :

// Unsupported : g_signal_handlers_block_matched : unsupported parameter func : no type generator for gpointer (gpointer) for param func

// Unsupported : g_signal_handlers_disconnect_matched : unsupported parameter func : no type generator for gpointer (gpointer) for param func

// Unsupported : g_signal_handlers_unblock_matched : unsupported parameter func : no type generator for gpointer (gpointer) for param func

// Unsupported : g_signal_has_handler_pending : return type :

// Unsupported : g_signal_list_ids : array return type :

// Unsupported : g_signal_name : return type :

// Unsupported : g_signal_new : unsupported parameter accumulator : no type generator for SignalAccumulator (GSignalAccumulator) for param accumulator

// Unsupported : g_signal_new_valist : unsupported parameter accumulator : no type generator for SignalAccumulator (GSignalAccumulator) for param accumulator

// Unsupported : g_signal_newv : unsupported parameter accumulator : no type generator for SignalAccumulator (GSignalAccumulator) for param accumulator

// Unsupported : g_signal_parse_name : return type :

// Unsupported : g_signal_type_cclosure_new : return type :

// Unsupported : g_strdup_value_contents : return type :

// Unsupported : g_type_add_class_cache_func : unsupported parameter cache_data : no type generator for gpointer (gpointer) for param cache_data

// Unsupported : g_type_check_class_cast : return type :

// Unsupported : g_type_check_class_is_a : return type :

// Unsupported : g_type_check_instance : return type :

// Unsupported : g_type_check_instance_cast : return type :

// Unsupported : g_type_check_instance_is_a : return type :

// Unsupported : g_type_check_instance_is_fundamentally_a : return type :

// Unsupported : g_type_check_is_value_type : return type :

// Unsupported : g_type_check_value : return type :

// Unsupported : g_type_check_value_holds : return type :

// Unsupported : g_type_children : array return type :

// Unsupported : g_type_class_adjust_private_offset : unsupported parameter g_class : no type generator for gpointer (gpointer) for param g_class

// Unsupported : g_type_class_peek : return type :

// Unsupported : g_type_class_ref : return type :

// Unsupported : g_type_create_instance : return type :

// Unsupported : g_type_from_name : return type :

// Unsupported : g_type_fundamental : return type :

// Unsupported : g_type_fundamental_next : return type :

// Unsupported : g_type_get_plugin : return type :

// Unsupported : g_type_get_qdata : no return generator

// Unsupported : g_type_interface_get_plugin : return type :

// Unsupported : g_type_interface_peek : return type :

// Unsupported : g_type_interfaces : array return type :

// Unsupported : g_type_is_a : return type :

// Unsupported : g_type_name : return type :

// Unsupported : g_type_name_from_class : return type :

// Unsupported : g_type_name_from_instance : return type :

// Unsupported : g_type_next_base : return type :

// Unsupported : g_type_parent : return type :

// Unsupported : g_type_qname : return type :

// Unsupported : g_type_register_dynamic : return type :

// Unsupported : g_type_register_fundamental : return type :

// Unsupported : g_type_register_static : return type :

// Unsupported : g_type_remove_class_cache_func : unsupported parameter cache_data : no type generator for gpointer (gpointer) for param cache_data

// Unsupported : g_type_set_qdata : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_type_test_flags : return type :

// Unsupported : g_type_value_table_peek : return type :

// Unsupported : g_value_register_transform_func : unsupported parameter transform_func : no type generator for ValueTransform (GValueTransform) for param transform_func

// Unsupported : g_value_type_compatible : return type :

// Unsupported : g_value_type_transformable : return type :

// TypePlugin is a wrapper around the C record GTypePlugin.
type TypePlugin struct {
	native *C.GTypePlugin
}

func TypePluginNewFromC(u unsafe.Pointer) *TypePlugin {
	c := (*C.GTypePlugin)(u)
	if c == nil {
		return nil
	}

	g := &TypePlugin{native: c}

	return g
}

func (recv *TypePlugin) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CClosure is a wrapper around the C record GCClosure.
type CClosure struct {
	native *C.GCClosure
	// closure : record
	// callback : no type generator for gpointer, gpointer
}

func CClosureNewFromC(u unsafe.Pointer) *CClosure {
	c := (*C.GCClosure)(u)
	if c == nil {
		return nil
	}

	g := &CClosure{native: c}

	return g
}

func (recv *CClosure) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Closure is a wrapper around the C record GClosure.
type Closure struct {
	native *C.GClosure
	// Private : ref_count
	// Private : meta_marshal_nouse
	// Private : n_guards
	// Private : n_fnotifiers
	// Private : n_inotifiers
	// Private : in_inotify
	// Private : floating
	// Private : derivative_flag
	// Bitfield not supported :  1 in_marshal
	// Bitfield not supported :  1 is_invalid
	// no type for marshal
	// Private : data
	// Private : notifiers
}

func ClosureNewFromC(u unsafe.Pointer) *Closure {
	c := (*C.GClosure)(u)
	if c == nil {
		return nil
	}

	g := &Closure{native: c}

	return g
}

func (recv *Closure) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ClosureNotifyData is a wrapper around the C record GClosureNotifyData.
type ClosureNotifyData struct {
	native *C.GClosureNotifyData
	// data : no type generator for gpointer, gpointer
	// notify : no type generator for ClosureNotify, GClosureNotify
}

func ClosureNotifyDataNewFromC(u unsafe.Pointer) *ClosureNotifyData {
	c := (*C.GClosureNotifyData)(u)
	if c == nil {
		return nil
	}

	g := &ClosureNotifyData{native: c}

	return g
}

func (recv *ClosureNotifyData) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EnumClass is a wrapper around the C record GEnumClass.
type EnumClass struct {
	native *C.GEnumClass
	// g_type_class : record
	Minimum int32
	Maximum int32
	NValues uint32
	// values : record
}

func EnumClassNewFromC(u unsafe.Pointer) *EnumClass {
	c := (*C.GEnumClass)(u)
	if c == nil {
		return nil
	}

	g := &EnumClass{
		Maximum: (int32)(c.maximum),
		Minimum: (int32)(c.minimum),
		NValues: (uint32)(c.n_values),
		native:  c,
	}

	return g
}

func (recv *EnumClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EnumValue is a wrapper around the C record GEnumValue.
type EnumValue struct {
	native    *C.GEnumValue
	Value     int32
	ValueName string
	ValueNick string
}

func EnumValueNewFromC(u unsafe.Pointer) *EnumValue {
	c := (*C.GEnumValue)(u)
	if c == nil {
		return nil
	}

	g := &EnumValue{
		Value:     (int32)(c.value),
		ValueName: C.GoString(c.value_name),
		ValueNick: C.GoString(c.value_nick),
		native:    c,
	}

	return g
}

func (recv *EnumValue) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FlagsClass is a wrapper around the C record GFlagsClass.
type FlagsClass struct {
	native *C.GFlagsClass
	// g_type_class : record
	Mask    uint32
	NValues uint32
	// values : record
}

func FlagsClassNewFromC(u unsafe.Pointer) *FlagsClass {
	c := (*C.GFlagsClass)(u)
	if c == nil {
		return nil
	}

	g := &FlagsClass{
		Mask:    (uint32)(c.mask),
		NValues: (uint32)(c.n_values),
		native:  c,
	}

	return g
}

func (recv *FlagsClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FlagsValue is a wrapper around the C record GFlagsValue.
type FlagsValue struct {
	native    *C.GFlagsValue
	Value     uint32
	ValueName string
	ValueNick string
}

func FlagsValueNewFromC(u unsafe.Pointer) *FlagsValue {
	c := (*C.GFlagsValue)(u)
	if c == nil {
		return nil
	}

	g := &FlagsValue{
		Value:     (uint32)(c.value),
		ValueName: C.GoString(c.value_name),
		ValueNick: C.GoString(c.value_nick),
		native:    c,
	}

	return g
}

func (recv *FlagsValue) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InitiallyUnownedClass is a wrapper around the C record GInitiallyUnownedClass.
type InitiallyUnownedClass struct {
	native *C.GInitiallyUnownedClass
	// g_type_class : record
	// Private : construct_properties
	// no type for constructor
	// no type for set_property
	// no type for get_property
	// no type for dispose
	// no type for finalize
	// no type for dispatch_properties_changed
	// no type for notify
	// no type for constructed
	// Private : flags
	// Private : pdummy
}

func InitiallyUnownedClassNewFromC(u unsafe.Pointer) *InitiallyUnownedClass {
	c := (*C.GInitiallyUnownedClass)(u)
	if c == nil {
		return nil
	}

	g := &InitiallyUnownedClass{native: c}

	return g
}

func (recv *InitiallyUnownedClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InterfaceInfo is a wrapper around the C record GInterfaceInfo.
type InterfaceInfo struct {
	native *C.GInterfaceInfo
	// interface_init : no type generator for InterfaceInitFunc, GInterfaceInitFunc
	// interface_finalize : no type generator for InterfaceFinalizeFunc, GInterfaceFinalizeFunc
	// interface_data : no type generator for gpointer, gpointer
}

func InterfaceInfoNewFromC(u unsafe.Pointer) *InterfaceInfo {
	c := (*C.GInterfaceInfo)(u)
	if c == nil {
		return nil
	}

	g := &InterfaceInfo{native: c}

	return g
}

func (recv *InterfaceInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ObjectClass is a wrapper around the C record GObjectClass.
type ObjectClass struct {
	native *C.GObjectClass
	// g_type_class : record
	// Private : construct_properties
	// no type for constructor
	// no type for set_property
	// no type for get_property
	// no type for dispose
	// no type for finalize
	// no type for dispatch_properties_changed
	// no type for notify
	// no type for constructed
	// Private : flags
	// Private : pdummy
}

func ObjectClassNewFromC(u unsafe.Pointer) *ObjectClass {
	c := (*C.GObjectClass)(u)
	if c == nil {
		return nil
	}

	g := &ObjectClass{native: c}

	return g
}

func (recv *ObjectClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ObjectConstructParam is a wrapper around the C record GObjectConstructParam.
type ObjectConstructParam struct {
	native *C.GObjectConstructParam
	// pspec : record
	// value : record
}

func ObjectConstructParamNewFromC(u unsafe.Pointer) *ObjectConstructParam {
	c := (*C.GObjectConstructParam)(u)
	if c == nil {
		return nil
	}

	g := &ObjectConstructParam{native: c}

	return g
}

func (recv *ObjectConstructParam) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecClass is a wrapper around the C record GParamSpecClass.
type ParamSpecClass struct {
	native *C.GParamSpecClass
	// g_type_class : record
	ValueType Type
	// no type for finalize
	// no type for value_set_default
	// no type for value_validate
	// no type for values_cmp
	// Private : dummy
}

func ParamSpecClassNewFromC(u unsafe.Pointer) *ParamSpecClass {
	c := (*C.GParamSpecClass)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecClass{
		ValueType: (Type)(c.value_type),
		native:    c,
	}

	return g
}

func (recv *ParamSpecClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecPool is a wrapper around the C record GParamSpecPool.
type ParamSpecPool struct {
	native *C.GParamSpecPool
}

func ParamSpecPoolNewFromC(u unsafe.Pointer) *ParamSpecPool {
	c := (*C.GParamSpecPool)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecPool{native: c}

	return g
}

func (recv *ParamSpecPool) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecTypeInfo is a wrapper around the C record GParamSpecTypeInfo.
type ParamSpecTypeInfo struct {
	native       *C.GParamSpecTypeInfo
	InstanceSize uint16
	NPreallocs   uint16
	// no type for instance_init
	ValueType Type
	// no type for finalize
	// no type for value_set_default
	// no type for value_validate
	// no type for values_cmp
}

func ParamSpecTypeInfoNewFromC(u unsafe.Pointer) *ParamSpecTypeInfo {
	c := (*C.GParamSpecTypeInfo)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecTypeInfo{
		InstanceSize: (uint16)(c.instance_size),
		NPreallocs:   (uint16)(c.n_preallocs),
		ValueType:    (Type)(c.value_type),
		native:       c,
	}

	return g
}

func (recv *ParamSpecTypeInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Parameter is a wrapper around the C record GParameter.
type Parameter struct {
	native *C.GParameter
	Name   string
	// value : record
}

func ParameterNewFromC(u unsafe.Pointer) *Parameter {
	c := (*C.GParameter)(u)
	if c == nil {
		return nil
	}

	g := &Parameter{
		Name:   C.GoString(c.name),
		native: c,
	}

	return g
}

func (recv *Parameter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SignalInvocationHint is a wrapper around the C record GSignalInvocationHint.
type SignalInvocationHint struct {
	native   *C.GSignalInvocationHint
	SignalId uint32
	Detail   glib.Quark
	RunType  SignalFlags
}

func SignalInvocationHintNewFromC(u unsafe.Pointer) *SignalInvocationHint {
	c := (*C.GSignalInvocationHint)(u)
	if c == nil {
		return nil
	}

	g := &SignalInvocationHint{
		Detail:   (glib.Quark)(c.detail),
		RunType:  (SignalFlags)(c.run_type),
		SignalId: (uint32)(c.signal_id),
		native:   c,
	}

	return g
}

func (recv *SignalInvocationHint) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SignalQuery is a wrapper around the C record GSignalQuery.
type SignalQuery struct {
	native      *C.GSignalQuery
	SignalId    uint32
	SignalName  string
	Itype       Type
	SignalFlags SignalFlags
	ReturnType  Type
	NParams     uint32
	// no type for param_types
}

func SignalQueryNewFromC(u unsafe.Pointer) *SignalQuery {
	c := (*C.GSignalQuery)(u)
	if c == nil {
		return nil
	}

	g := &SignalQuery{
		Itype:       (Type)(c.itype),
		NParams:     (uint32)(c.n_params),
		ReturnType:  (Type)(c.return_type),
		SignalFlags: (SignalFlags)(c.signal_flags),
		SignalId:    (uint32)(c.signal_id),
		SignalName:  C.GoString(c.signal_name),
		native:      c,
	}

	return g
}

func (recv *SignalQuery) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TypeClass is a wrapper around the C record GTypeClass.
type TypeClass struct {
	native *C.GTypeClass
	// Private : g_type
}

func TypeClassNewFromC(u unsafe.Pointer) *TypeClass {
	c := (*C.GTypeClass)(u)
	if c == nil {
		return nil
	}

	g := &TypeClass{native: c}

	return g
}

func (recv *TypeClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TypeFundamentalInfo is a wrapper around the C record GTypeFundamentalInfo.
type TypeFundamentalInfo struct {
	native    *C.GTypeFundamentalInfo
	TypeFlags TypeFundamentalFlags
}

func TypeFundamentalInfoNewFromC(u unsafe.Pointer) *TypeFundamentalInfo {
	c := (*C.GTypeFundamentalInfo)(u)
	if c == nil {
		return nil
	}

	g := &TypeFundamentalInfo{
		TypeFlags: (TypeFundamentalFlags)(c.type_flags),
		native:    c,
	}

	return g
}

func (recv *TypeFundamentalInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TypeInfo is a wrapper around the C record GTypeInfo.
type TypeInfo struct {
	native    *C.GTypeInfo
	ClassSize uint16
	// base_init : no type generator for BaseInitFunc, GBaseInitFunc
	// base_finalize : no type generator for BaseFinalizeFunc, GBaseFinalizeFunc
	// class_init : no type generator for ClassInitFunc, GClassInitFunc
	// class_finalize : no type generator for ClassFinalizeFunc, GClassFinalizeFunc
	// class_data : no type generator for gpointer, gconstpointer
	InstanceSize uint16
	NPreallocs   uint16
	// instance_init : no type generator for InstanceInitFunc, GInstanceInitFunc
	// value_table : record
}

func TypeInfoNewFromC(u unsafe.Pointer) *TypeInfo {
	c := (*C.GTypeInfo)(u)
	if c == nil {
		return nil
	}

	g := &TypeInfo{
		ClassSize:    (uint16)(c.class_size),
		InstanceSize: (uint16)(c.instance_size),
		NPreallocs:   (uint16)(c.n_preallocs),
		native:       c,
	}

	return g
}

func (recv *TypeInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TypeInstance is a wrapper around the C record GTypeInstance.
type TypeInstance struct {
	native *C.GTypeInstance
	// Private : g_class
}

func TypeInstanceNewFromC(u unsafe.Pointer) *TypeInstance {
	c := (*C.GTypeInstance)(u)
	if c == nil {
		return nil
	}

	g := &TypeInstance{native: c}

	return g
}

func (recv *TypeInstance) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TypeInterface is a wrapper around the C record GTypeInterface.
type TypeInterface struct {
	native *C.GTypeInterface
	// Private : g_type
	// Private : g_instance_type
}

func TypeInterfaceNewFromC(u unsafe.Pointer) *TypeInterface {
	c := (*C.GTypeInterface)(u)
	if c == nil {
		return nil
	}

	g := &TypeInterface{native: c}

	return g
}

func (recv *TypeInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TypeModuleClass is a wrapper around the C record GTypeModuleClass.
type TypeModuleClass struct {
	native *C.GTypeModuleClass
	// parent_class : record
	// no type for load
	// no type for unload
	// no type for reserved1
	// no type for reserved2
	// no type for reserved3
	// no type for reserved4
}

func TypeModuleClassNewFromC(u unsafe.Pointer) *TypeModuleClass {
	c := (*C.GTypeModuleClass)(u)
	if c == nil {
		return nil
	}

	g := &TypeModuleClass{native: c}

	return g
}

func (recv *TypeModuleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TypePluginClass is a wrapper around the C record GTypePluginClass.
type TypePluginClass struct {
	native *C.GTypePluginClass
	// Private : base_iface
	// use_plugin : no type generator for TypePluginUse, GTypePluginUse
	// unuse_plugin : no type generator for TypePluginUnuse, GTypePluginUnuse
	// complete_type_info : no type generator for TypePluginCompleteTypeInfo, GTypePluginCompleteTypeInfo
	// complete_interface_info : no type generator for TypePluginCompleteInterfaceInfo, GTypePluginCompleteInterfaceInfo
}

func TypePluginClassNewFromC(u unsafe.Pointer) *TypePluginClass {
	c := (*C.GTypePluginClass)(u)
	if c == nil {
		return nil
	}

	g := &TypePluginClass{native: c}

	return g
}

func (recv *TypePluginClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TypeQuery is a wrapper around the C record GTypeQuery.
type TypeQuery struct {
	native       *C.GTypeQuery
	Type         Type
	TypeName     string
	ClassSize    uint32
	InstanceSize uint32
}

func TypeQueryNewFromC(u unsafe.Pointer) *TypeQuery {
	c := (*C.GTypeQuery)(u)
	if c == nil {
		return nil
	}

	g := &TypeQuery{
		ClassSize:    (uint32)(c.class_size),
		InstanceSize: (uint32)(c.instance_size),
		Type:         (Type)(c._type),
		TypeName:     C.GoString(c.type_name),
		native:       c,
	}

	return g
}

func (recv *TypeQuery) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TypeValueTable is a wrapper around the C record GTypeValueTable.
type TypeValueTable struct {
	native *C.GTypeValueTable
	// no type for value_init
	// no type for value_free
	// no type for value_copy
	// no type for value_peek_pointer
	CollectFormat string
	// no type for collect_value
	LcopyFormat string
	// no type for lcopy_value
}

func TypeValueTableNewFromC(u unsafe.Pointer) *TypeValueTable {
	c := (*C.GTypeValueTable)(u)
	if c == nil {
		return nil
	}

	g := &TypeValueTable{
		CollectFormat: C.GoString(c.collect_format),
		LcopyFormat:   C.GoString(c.lcopy_format),
		native:        c,
	}

	return g
}

func (recv *TypeValueTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Value is a wrapper around the C record GValue.
type Value struct {
	native *C.GValue
	// Private : g_type
	// no type for data
}

func ValueNewFromC(u unsafe.Pointer) *Value {
	c := (*C.GValue)(u)
	if c == nil {
		return nil
	}

	g := &Value{native: c}

	return g
}

func (recv *Value) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ValueArray is a wrapper around the C record GValueArray.
type ValueArray struct {
	native  *C.GValueArray
	NValues uint32
	// values : record
	// Private : n_prealloced
}

func ValueArrayNewFromC(u unsafe.Pointer) *ValueArray {
	c := (*C.GValueArray)(u)
	if c == nil {
		return nil
	}

	g := &ValueArray{
		NValues: (uint32)(c.n_values),
		native:  c,
	}

	return g
}

func (recv *ValueArray) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WeakRef is a wrapper around the C record GWeakRef.
type WeakRef struct {
	native *C.GWeakRef
}

func WeakRefNewFromC(u unsafe.Pointer) *WeakRef {
	c := (*C.GWeakRef)(u)
	if c == nil {
		return nil
	}

	g := &WeakRef{native: c}

	return g
}

func (recv *WeakRef) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
