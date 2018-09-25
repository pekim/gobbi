// This is a generated file - DO NOT EDIT

package gobject

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

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

// Unsupported : g_object_new : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_new_valist : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_new_with_properties : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_newv : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_add_toggle_ref : unsupported parameter notify : no type generator for ToggleNotify, GToggleNotify

// Unsupported : g_object_add_weak_pointer : unsupported parameter weak_pointer_location : no type generator for gpointer, gpointer*

// Unsupported : g_object_bind_property_full : unsupported parameter transform_to : no type generator for BindingTransformFunc, GBindingTransformFunc

// Unsupported : g_object_connect : unsupported parameter ... : varargs

// Unsupported : g_object_disconnect : unsupported parameter ... : varargs

// Unsupported : g_object_dup_data : unsupported parameter dup_func : no type generator for GLib.DuplicateFunc, GDuplicateFunc

// Unsupported : g_object_dup_qdata : unsupported parameter dup_func : no type generator for GLib.DuplicateFunc, GDuplicateFunc

// Unsupported : g_object_force_floating : no return generator

// Unsupported : g_object_freeze_notify : no return generator

// Unsupported : g_object_get : unsupported parameter ... : varargs

// GetData is a wrapper around the C function g_object_get_data.
func (recv *Object) GetData(key string) uintptr {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_object_get_data((*C.GObject)(recv.native), c_key)
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_object_get_property : no return generator

// GetQdata is a wrapper around the C function g_object_get_qdata.
func (recv *Object) GetQdata(quark glib.Quark) uintptr {
	c_quark := (C.GQuark)(quark)

	retC := C.g_object_get_qdata((*C.GObject)(recv.native), c_quark)
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_object_get_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : g_object_getv : unsupported parameter names : no param type

// Unsupported : g_object_notify : no return generator

// Unsupported : g_object_notify_by_pspec : no return generator

// Ref is a wrapper around the C function g_object_ref.
func (recv *Object) Ref() uintptr {
	retC := C.g_object_ref((C.gpointer)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_object_remove_toggle_ref : unsupported parameter notify : no type generator for ToggleNotify, GToggleNotify

// Unsupported : g_object_remove_weak_pointer : unsupported parameter weak_pointer_location : no type generator for gpointer, gpointer*

// Unsupported : g_object_replace_data : unsupported parameter destroy : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_object_replace_qdata : unsupported parameter destroy : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_object_run_dispose : no return generator

// Unsupported : g_object_set : unsupported parameter ... : varargs

// Unsupported : g_object_set_data : no return generator

// Unsupported : g_object_set_data_full : unsupported parameter destroy : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_object_set_property : no return generator

// Unsupported : g_object_set_qdata : no return generator

// Unsupported : g_object_set_qdata_full : unsupported parameter destroy : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_object_set_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : g_object_setv : unsupported parameter names : no param type

// StealData is a wrapper around the C function g_object_steal_data.
func (recv *Object) StealData(key string) uintptr {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_object_steal_data((*C.GObject)(recv.native), c_key)
	retGo := (uintptr)(retC)

	return retGo
}

// StealQdata is a wrapper around the C function g_object_steal_qdata.
func (recv *Object) StealQdata(quark glib.Quark) uintptr {
	c_quark := (C.GQuark)(quark)

	retC := C.g_object_steal_qdata((*C.GObject)(recv.native), c_quark)
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_object_thaw_notify : no return generator

// Unsupported : g_object_unref : no return generator

// Unsupported : g_object_watch_closure : no return generator

// Unsupported : g_object_weak_ref : unsupported parameter notify : no type generator for WeakNotify, GWeakNotify

// Unsupported : g_object_weak_unref : unsupported parameter notify : no type generator for WeakNotify, GWeakNotify

// ParamSpec is a wrapper around the C record GParamSpec.
type ParamSpec struct {
	native *C.GParamSpec
	// g_type_instance : record
	Name  string
	Flags ParamFlags
	// value_type : no type generator for GType, GType
	// owner_type : no type generator for GType, GType
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
		Flags:  (ParamFlags)(c.flags),
		Name:   C.GoString(c.name),
		native: c,
	}

	return g
}

func (recv *ParamSpec) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)
	recv.native.flags =
		(C.GParamFlags)(recv.Flags)

	return (unsafe.Pointer)(recv.native)
}

// GetBlurb is a wrapper around the C function g_param_spec_get_blurb.
func (recv *ParamSpec) GetBlurb() string {
	retC := C.g_param_spec_get_blurb((*C.GParamSpec)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetName is a wrapper around the C function g_param_spec_get_name.
func (recv *ParamSpec) GetName() string {
	retC := C.g_param_spec_get_name((*C.GParamSpec)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetNick is a wrapper around the C function g_param_spec_get_nick.
func (recv *ParamSpec) GetNick() string {
	retC := C.g_param_spec_get_nick((*C.GParamSpec)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetQdata is a wrapper around the C function g_param_spec_get_qdata.
func (recv *ParamSpec) GetQdata(quark glib.Quark) uintptr {
	c_quark := (C.GQuark)(quark)

	retC := C.g_param_spec_get_qdata((*C.GParamSpec)(recv.native), c_quark)
	retGo := (uintptr)(retC)

	return retGo
}

// Ref is a wrapper around the C function g_param_spec_ref.
func (recv *ParamSpec) Ref() *ParamSpec {
	retC := C.g_param_spec_ref((*C.GParamSpec)(recv.native))
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_param_spec_set_qdata : no return generator

// Unsupported : g_param_spec_set_qdata_full : unsupported parameter destroy : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_param_spec_sink : no return generator

// StealQdata is a wrapper around the C function g_param_spec_steal_qdata.
func (recv *ParamSpec) StealQdata(quark glib.Quark) uintptr {
	c_quark := (C.GQuark)(quark)

	retC := C.g_param_spec_steal_qdata((*C.GParamSpec)(recv.native), c_quark)
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_param_spec_unref : no return generator

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
	recv.native.default_value =
		boolToGboolean(recv.DefaultValue)

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
	recv.native.minimum =
		(C.gint8)(recv.Minimum)
	recv.native.maximum =
		(C.gint8)(recv.Maximum)
	recv.native.default_value =
		(C.gint8)(recv.DefaultValue)

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
	recv.native.minimum =
		(C.gdouble)(recv.Minimum)
	recv.native.maximum =
		(C.gdouble)(recv.Maximum)
	recv.native.default_value =
		(C.gdouble)(recv.DefaultValue)
	recv.native.epsilon =
		(C.gdouble)(recv.Epsilon)

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
	recv.native.default_value =
		(C.gint)(recv.DefaultValue)

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
	recv.native.default_value =
		(C.guint)(recv.DefaultValue)

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
	recv.native.minimum =
		(C.gfloat)(recv.Minimum)
	recv.native.maximum =
		(C.gfloat)(recv.Maximum)
	recv.native.default_value =
		(C.gfloat)(recv.DefaultValue)
	recv.native.epsilon =
		(C.gfloat)(recv.Epsilon)

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
	recv.native.minimum =
		(C.gint)(recv.Minimum)
	recv.native.maximum =
		(C.gint)(recv.Maximum)
	recv.native.default_value =
		(C.gint)(recv.DefaultValue)

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
	recv.native.minimum =
		(C.gint64)(recv.Minimum)
	recv.native.maximum =
		(C.gint64)(recv.Maximum)
	recv.native.default_value =
		(C.gint64)(recv.DefaultValue)

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
	recv.native.minimum =
		(C.glong)(recv.Minimum)
	recv.native.maximum =
		(C.glong)(recv.Maximum)
	recv.native.default_value =
		(C.glong)(recv.DefaultValue)

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
	recv.native.default_value =
		C.CString(recv.DefaultValue)
	recv.native.cset_first =
		C.CString(recv.CsetFirst)
	recv.native.cset_nth =
		C.CString(recv.CsetNth)
	recv.native.substitutor =
		(C.gchar)(recv.Substitutor)

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
	recv.native.minimum =
		(C.guint8)(recv.Minimum)
	recv.native.maximum =
		(C.guint8)(recv.Maximum)
	recv.native.default_value =
		(C.guint8)(recv.DefaultValue)

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
	recv.native.minimum =
		(C.guint)(recv.Minimum)
	recv.native.maximum =
		(C.guint)(recv.Maximum)
	recv.native.default_value =
		(C.guint)(recv.DefaultValue)

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
	recv.native.minimum =
		(C.guint64)(recv.Minimum)
	recv.native.maximum =
		(C.guint64)(recv.Maximum)
	recv.native.default_value =
		(C.guint64)(recv.DefaultValue)

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
	recv.native.minimum =
		(C.gulong)(recv.Minimum)
	recv.native.maximum =
		(C.gulong)(recv.Maximum)
	recv.native.default_value =
		(C.gulong)(recv.DefaultValue)

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
	recv.native.default_value =
		(C.gunichar)(recv.DefaultValue)

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
	recv.native.fixed_n_elements =
		(C.guint)(recv.FixedNElements)

	return (unsafe.Pointer)(recv.native)
}

// TypeModule is a wrapper around the C record GTypeModule.
type TypeModule struct {
	native *C.GTypeModule
	// parent_instance : record
	UseCount uint32
	// type_infos : record
	// interface_infos : record
	Name string
}

func TypeModuleNewFromC(u unsafe.Pointer) *TypeModule {
	c := (*C.GTypeModule)(u)
	if c == nil {
		return nil
	}

	g := &TypeModule{
		Name:     C.GoString(c.name),
		UseCount: (uint32)(c.use_count),
		native:   c,
	}

	return g
}

func (recv *TypeModule) ToC() unsafe.Pointer {
	recv.native.use_count =
		(C.guint)(recv.UseCount)
	recv.native.name =
		C.CString(recv.Name)

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_type_module_add_interface : unsupported parameter instance_type : no type generator for GType, GType

// Unsupported : g_type_module_register_enum : no return generator

// Unsupported : g_type_module_register_flags : no return generator

// Unsupported : g_type_module_register_type : unsupported parameter parent_type : no type generator for GType, GType

// Unsupported : g_type_module_set_name : no return generator

// Unsupported : g_type_module_unuse : no return generator

// Use is a wrapper around the C function g_type_module_use.
func (recv *TypeModule) Use() bool {
	retC := C.g_type_module_use((*C.GTypeModule)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
