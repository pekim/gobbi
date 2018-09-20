// This is a generated file - DO NOT EDIT

package gobject

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// InitiallyUnowned is a wrapper around the C record GInitiallyUnowned.
type InitiallyUnowned struct {
	native        *C.GInitiallyUnowned
	GTypeInstance *TypeInstance
	// ref_count : no type generator for guint, volatile guint
	// qdata : no type generator for GLib.Data, GData*
}

func initiallyUnownedNewFromC(c *C.GInitiallyUnowned) *InitiallyUnowned {
	if c == nil {
		return nil
	}

	g := &InitiallyUnowned{
		GTypeInstance: typeInstanceNewFromC(c.g_type_instance),
		native:        c,
	}

	return g
}

// Object is a wrapper around the C record GObject.
type Object struct {
	native        *C.GObject
	GTypeInstance *TypeInstance
	// ref_count : no type generator for guint, volatile guint
	// qdata : no type generator for GLib.Data, GData*
}

func objectNewFromC(c *C.GObject) *Object {
	if c == nil {
		return nil
	}

	g := &Object{
		GTypeInstance: typeInstanceNewFromC(c.g_type_instance),
		native:        c,
	}

	return g
}

// Unsupported : g_object_new : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_new_valist : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_new_with_properties : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_newv : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_add_toggle_ref : unsupported parameter notify : no type generator for ToggleNotify, GToggleNotify

// Unsupported : g_object_add_weak_pointer : unsupported parameter weak_pointer_location : no type generator for gpointer, gpointer*

// Unsupported : g_object_bind_property : unsupported parameter flags : no type generator for BindingFlags, GBindingFlags

// Unsupported : g_object_bind_property_full : unsupported parameter flags : no type generator for BindingFlags, GBindingFlags

// Unsupported : g_object_bind_property_with_closures : unsupported parameter flags : no type generator for BindingFlags, GBindingFlags

// Unsupported : g_object_connect : unsupported parameter ... : varargs

// Unsupported : g_object_disconnect : unsupported parameter ... : varargs

// Unsupported : g_object_dup_data : unsupported parameter dup_func : no type generator for GLib.DuplicateFunc, GDuplicateFunc

// Unsupported : g_object_dup_qdata : unsupported parameter quark : no type generator for GLib.Quark, GQuark

// Unsupported : g_object_force_floating : no return generator

// Unsupported : g_object_freeze_notify : no return generator

// Unsupported : g_object_get : unsupported parameter ... : varargs

// GetData is a wrapper around the C function g_object_get_data.
func (recv *Object) GetData(key string) uintptr {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_object_get_data(recv.native, c_key)
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_object_get_property : unsupported parameter value : record param - coming soon

// Unsupported : g_object_get_qdata : unsupported parameter quark : no type generator for GLib.Quark, GQuark

// Unsupported : g_object_get_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : g_object_getv : unsupported parameter names : no param type

// Unsupported : g_object_is_floating : no return generator

// Unsupported : g_object_notify : no return generator

// Unsupported : g_object_notify_by_pspec : unsupported parameter pspec : no type generator for ParamSpec, GParamSpec*

// Ref is a wrapper around the C function g_object_ref.
func (recv *Object) Ref() uintptr {
	retC := C.g_object_ref(recv.native)
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_object_remove_toggle_ref : unsupported parameter notify : no type generator for ToggleNotify, GToggleNotify

// Unsupported : g_object_remove_weak_pointer : unsupported parameter weak_pointer_location : no type generator for gpointer, gpointer*

// Unsupported : g_object_replace_data : unsupported parameter destroy : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_object_replace_qdata : unsupported parameter quark : no type generator for GLib.Quark, GQuark

// Unsupported : g_object_run_dispose : no return generator

// Unsupported : g_object_set : unsupported parameter ... : varargs

// Unsupported : g_object_set_data : no return generator

// Unsupported : g_object_set_data_full : unsupported parameter destroy : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_object_set_property : unsupported parameter value : record param - coming soon

// Unsupported : g_object_set_qdata : unsupported parameter quark : no type generator for GLib.Quark, GQuark

// Unsupported : g_object_set_qdata_full : unsupported parameter quark : no type generator for GLib.Quark, GQuark

// Unsupported : g_object_set_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : g_object_setv : unsupported parameter names : no param type

// StealData is a wrapper around the C function g_object_steal_data.
func (recv *Object) StealData(key string) uintptr {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_object_steal_data(recv.native, c_key)
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_object_steal_qdata : unsupported parameter quark : no type generator for GLib.Quark, GQuark

// Unsupported : g_object_thaw_notify : no return generator

// Unsupported : g_object_unref : no return generator

// Unsupported : g_object_watch_closure : unsupported parameter closure : Blacklisted record : GClosure

// Unsupported : g_object_weak_ref : unsupported parameter notify : no type generator for WeakNotify, GWeakNotify

// Unsupported : g_object_weak_unref : unsupported parameter notify : no type generator for WeakNotify, GWeakNotify

// ParamSpec is a wrapper around the C record GParamSpec.
type ParamSpec struct {
	native        *C.GParamSpec
	GTypeInstance *TypeInstance
	Name          string
	// flags : no type generator for ParamFlags, GParamFlags
	// value_type : no type generator for GType, GType
	// owner_type : no type generator for GType, GType
	Nick  string
	Blurb string
	// qdata : no type generator for GLib.Data, GData*
	RefCount uint32
	ParamId  uint32
}

func paramSpecNewFromC(c *C.GParamSpec) *ParamSpec {
	if c == nil {
		return nil
	}

	g := &ParamSpec{
		Blurb:         C.GoString(c._blurb),
		GTypeInstance: typeInstanceNewFromC(c.g_type_instance),
		Name:          C.GoString(c.name),
		Nick:          C.GoString(c._nick),
		ParamId:       (uint32)(c.param_id),
		RefCount:      (uint32)(c.ref_count),
		native:        c,
	}

	return g
}

// GetBlurb is a wrapper around the C function g_param_spec_get_blurb.
func (recv *ParamSpec) GetBlurb() string {
	retC := C.g_param_spec_get_blurb(recv.native)
	retGo := C.GoString(retC)

	return retGo
}

// GetName is a wrapper around the C function g_param_spec_get_name.
func (recv *ParamSpec) GetName() string {
	retC := C.g_param_spec_get_name(recv.native)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_param_spec_get_name_quark : no return generator

// GetNick is a wrapper around the C function g_param_spec_get_nick.
func (recv *ParamSpec) GetNick() string {
	retC := C.g_param_spec_get_nick(recv.native)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_param_spec_get_qdata : unsupported parameter quark : no type generator for GLib.Quark, GQuark

// Unsupported : g_param_spec_get_redirect_target : no return generator

// Unsupported : g_param_spec_ref : no return generator

// Unsupported : g_param_spec_ref_sink : no return generator

// Unsupported : g_param_spec_set_qdata : unsupported parameter quark : no type generator for GLib.Quark, GQuark

// Unsupported : g_param_spec_set_qdata_full : unsupported parameter quark : no type generator for GLib.Quark, GQuark

// Unsupported : g_param_spec_sink : no return generator

// Unsupported : g_param_spec_steal_qdata : unsupported parameter quark : no type generator for GLib.Quark, GQuark

// Unsupported : g_param_spec_unref : no return generator

// ParamSpecBoolean is a wrapper around the C record GParamSpecBoolean.
type ParamSpecBoolean struct {
	native *C.GParamSpecBoolean
	// parent_instance : no type generator for ParamSpec, GParamSpec
	// default_value : no type generator for gboolean, gboolean
}

func paramSpecBooleanNewFromC(c *C.GParamSpecBoolean) *ParamSpecBoolean {
	if c == nil {
		return nil
	}

	g := &ParamSpecBoolean{native: c}

	return g
}

// ParamSpecBoxed is a wrapper around the C record GParamSpecBoxed.
type ParamSpecBoxed struct {
	native *C.GParamSpecBoxed
	// parent_instance : no type generator for ParamSpec, GParamSpec
}

func paramSpecBoxedNewFromC(c *C.GParamSpecBoxed) *ParamSpecBoxed {
	if c == nil {
		return nil
	}

	g := &ParamSpecBoxed{native: c}

	return g
}

// ParamSpecChar is a wrapper around the C record GParamSpecChar.
type ParamSpecChar struct {
	native *C.GParamSpecChar
	// parent_instance : no type generator for ParamSpec, GParamSpec
	Minimum      int8
	Maximum      int8
	DefaultValue int8
}

func paramSpecCharNewFromC(c *C.GParamSpecChar) *ParamSpecChar {
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

// ParamSpecDouble is a wrapper around the C record GParamSpecDouble.
type ParamSpecDouble struct {
	native *C.GParamSpecDouble
	// parent_instance : no type generator for ParamSpec, GParamSpec
	Minimum      float64
	Maximum      float64
	DefaultValue float64
	Epsilon      float64
}

func paramSpecDoubleNewFromC(c *C.GParamSpecDouble) *ParamSpecDouble {
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

// ParamSpecEnum is a wrapper around the C record GParamSpecEnum.
type ParamSpecEnum struct {
	native *C.GParamSpecEnum
	// parent_instance : no type generator for ParamSpec, GParamSpec
	EnumClass    *EnumClass
	DefaultValue int32
}

func paramSpecEnumNewFromC(c *C.GParamSpecEnum) *ParamSpecEnum {
	if c == nil {
		return nil
	}

	g := &ParamSpecEnum{
		DefaultValue: (int32)(c.default_value),
		EnumClass:    enumClassNewFromC(c.enum_class),
		native:       c,
	}

	return g
}

// ParamSpecFlags is a wrapper around the C record GParamSpecFlags.
type ParamSpecFlags struct {
	native *C.GParamSpecFlags
	// parent_instance : no type generator for ParamSpec, GParamSpec
	FlagsClass   *FlagsClass
	DefaultValue uint32
}

func paramSpecFlagsNewFromC(c *C.GParamSpecFlags) *ParamSpecFlags {
	if c == nil {
		return nil
	}

	g := &ParamSpecFlags{
		DefaultValue: (uint32)(c.default_value),
		FlagsClass:   flagsClassNewFromC(c.flags_class),
		native:       c,
	}

	return g
}

// ParamSpecFloat is a wrapper around the C record GParamSpecFloat.
type ParamSpecFloat struct {
	native *C.GParamSpecFloat
	// parent_instance : no type generator for ParamSpec, GParamSpec
	Minimum      float32
	Maximum      float32
	DefaultValue float32
	Epsilon      float32
}

func paramSpecFloatNewFromC(c *C.GParamSpecFloat) *ParamSpecFloat {
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

// ParamSpecInt is a wrapper around the C record GParamSpecInt.
type ParamSpecInt struct {
	native *C.GParamSpecInt
	// parent_instance : no type generator for ParamSpec, GParamSpec
	Minimum      int32
	Maximum      int32
	DefaultValue int32
}

func paramSpecIntNewFromC(c *C.GParamSpecInt) *ParamSpecInt {
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

// ParamSpecInt64 is a wrapper around the C record GParamSpecInt64.
type ParamSpecInt64 struct {
	native *C.GParamSpecInt64
	// parent_instance : no type generator for ParamSpec, GParamSpec
	Minimum      int64
	Maximum      int64
	DefaultValue int64
}

func paramSpecInt64NewFromC(c *C.GParamSpecInt64) *ParamSpecInt64 {
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

// ParamSpecLong is a wrapper around the C record GParamSpecLong.
type ParamSpecLong struct {
	native *C.GParamSpecLong
	// parent_instance : no type generator for ParamSpec, GParamSpec
	Minimum      int64
	Maximum      int64
	DefaultValue int64
}

func paramSpecLongNewFromC(c *C.GParamSpecLong) *ParamSpecLong {
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

// ParamSpecObject is a wrapper around the C record GParamSpecObject.
type ParamSpecObject struct {
	native *C.GParamSpecObject
	// parent_instance : no type generator for ParamSpec, GParamSpec
}

func paramSpecObjectNewFromC(c *C.GParamSpecObject) *ParamSpecObject {
	if c == nil {
		return nil
	}

	g := &ParamSpecObject{native: c}

	return g
}

// ParamSpecParam is a wrapper around the C record GParamSpecParam.
type ParamSpecParam struct {
	native *C.GParamSpecParam
	// parent_instance : no type generator for ParamSpec, GParamSpec
}

func paramSpecParamNewFromC(c *C.GParamSpecParam) *ParamSpecParam {
	if c == nil {
		return nil
	}

	g := &ParamSpecParam{native: c}

	return g
}

// ParamSpecPointer is a wrapper around the C record GParamSpecPointer.
type ParamSpecPointer struct {
	native *C.GParamSpecPointer
	// parent_instance : no type generator for ParamSpec, GParamSpec
}

func paramSpecPointerNewFromC(c *C.GParamSpecPointer) *ParamSpecPointer {
	if c == nil {
		return nil
	}

	g := &ParamSpecPointer{native: c}

	return g
}

// ParamSpecString is a wrapper around the C record GParamSpecString.
type ParamSpecString struct {
	native *C.GParamSpecString
	// parent_instance : no type generator for ParamSpec, GParamSpec
	DefaultValue string
	CsetFirst    string
	CsetNth      string
	Substitutor  rune
	// Bitfield not supported :  1 null_fold_if_empty
	// Bitfield not supported :  1 ensure_non_null
}

func paramSpecStringNewFromC(c *C.GParamSpecString) *ParamSpecString {
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

// ParamSpecUChar is a wrapper around the C record GParamSpecUChar.
type ParamSpecUChar struct {
	native *C.GParamSpecUChar
	// parent_instance : no type generator for ParamSpec, GParamSpec
	Minimum      uint8
	Maximum      uint8
	DefaultValue uint8
}

func paramSpecUCharNewFromC(c *C.GParamSpecUChar) *ParamSpecUChar {
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

// ParamSpecUInt is a wrapper around the C record GParamSpecUInt.
type ParamSpecUInt struct {
	native *C.GParamSpecUInt
	// parent_instance : no type generator for ParamSpec, GParamSpec
	Minimum      uint32
	Maximum      uint32
	DefaultValue uint32
}

func paramSpecUIntNewFromC(c *C.GParamSpecUInt) *ParamSpecUInt {
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

// ParamSpecUInt64 is a wrapper around the C record GParamSpecUInt64.
type ParamSpecUInt64 struct {
	native *C.GParamSpecUInt64
	// parent_instance : no type generator for ParamSpec, GParamSpec
	Minimum      uint64
	Maximum      uint64
	DefaultValue uint64
}

func paramSpecUInt64NewFromC(c *C.GParamSpecUInt64) *ParamSpecUInt64 {
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

// ParamSpecULong is a wrapper around the C record GParamSpecULong.
type ParamSpecULong struct {
	native *C.GParamSpecULong
	// parent_instance : no type generator for ParamSpec, GParamSpec
	Minimum      uint64
	Maximum      uint64
	DefaultValue uint64
}

func paramSpecULongNewFromC(c *C.GParamSpecULong) *ParamSpecULong {
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

// ParamSpecUnichar is a wrapper around the C record GParamSpecUnichar.
type ParamSpecUnichar struct {
	native *C.GParamSpecUnichar
	// parent_instance : no type generator for ParamSpec, GParamSpec
	DefaultValue rune
}

func paramSpecUnicharNewFromC(c *C.GParamSpecUnichar) *ParamSpecUnichar {
	if c == nil {
		return nil
	}

	g := &ParamSpecUnichar{
		DefaultValue: (rune)(c.default_value),
		native:       c,
	}

	return g
}

// ParamSpecValueArray is a wrapper around the C record GParamSpecValueArray.
type ParamSpecValueArray struct {
	native *C.GParamSpecValueArray
	// parent_instance : no type generator for ParamSpec, GParamSpec
	// element_spec : no type generator for ParamSpec, GParamSpec*
	FixedNElements uint32
}

func paramSpecValueArrayNewFromC(c *C.GParamSpecValueArray) *ParamSpecValueArray {
	if c == nil {
		return nil
	}

	g := &ParamSpecValueArray{
		FixedNElements: (uint32)(c.fixed_n_elements),
		native:         c,
	}

	return g
}

// TypeModule is a wrapper around the C record GTypeModule.
type TypeModule struct {
	native *C.GTypeModule
	// parent_instance : no type generator for Object, GObject
	UseCount uint32
	// type_infos : no type generator for GLib.SList, GSList*
	// interface_infos : no type generator for GLib.SList, GSList*
	Name string
}

func typeModuleNewFromC(c *C.GTypeModule) *TypeModule {
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

// Unsupported : g_type_module_add_interface : unsupported parameter instance_type : no type generator for GType, GType

// Unsupported : g_type_module_register_enum : unsupported parameter const_static_values : record param - coming soon

// Unsupported : g_type_module_register_flags : unsupported parameter const_static_values : record param - coming soon

// Unsupported : g_type_module_register_type : unsupported parameter parent_type : no type generator for GType, GType

// Unsupported : g_type_module_set_name : no return generator

// Unsupported : g_type_module_unuse : no return generator

// Unsupported : g_type_module_use : no return generator