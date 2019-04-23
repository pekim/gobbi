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

// Equals compares this CClosure with another CClosure, and returns true if they represent the same GObject.
func (recv *CClosure) Equals(other *CClosure) bool {
	return other.ToC() == recv.ToC()
}

// g_cclosure_marshal_BOOLEAN__BOXED_BOXED : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_BOOLEAN__BOXED_BOXEDv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_BOOLEAN__FLAGS : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_BOOLEAN__FLAGSv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_STRING__OBJECT_POINTER : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_STRING__OBJECT_POINTERv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__BOOLEAN : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_VOID__BOOLEANv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__BOXED : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_VOID__BOXEDv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__CHAR : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_VOID__CHARv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__DOUBLE : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_VOID__DOUBLEv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__ENUM : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_VOID__ENUMv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__FLAGS : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_VOID__FLAGSv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__FLOAT : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_VOID__FLOATv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__INT : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_VOID__INTv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__LONG : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_VOID__LONGv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__OBJECT : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_VOID__OBJECTv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__PARAM : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_VOID__PARAMv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__POINTER : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_VOID__POINTERv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__STRING : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_VOID__STRINGv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__UCHAR : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_VOID__UCHARv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__UINT : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_VOID__UINT_POINTER : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_VOID__UINT_POINTERv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__UINTv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__ULONG : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_VOID__ULONGv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__VARIANT : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_VOID__VARIANTv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__VOID : unsupported parameter invocation_hint : no type generator for gpointer (gpointer) for param invocation_hint
// g_cclosure_marshal_VOID__VOIDv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_new : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func
// g_cclosure_new_object : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func
// g_cclosure_new_object_swap : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func
// g_cclosure_new_swap : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func
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

// Equals compares this Closure with another Closure, and returns true if they represent the same GObject.
func (recv *Closure) Equals(other *Closure) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_closure_new_object

// Unsupported : g_closure_new_simple : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_closure_add_finalize_notifier : unsupported parameter notify_data : no type generator for gpointer (gpointer) for param notify_data

// Unsupported : g_closure_add_invalidate_notifier : unsupported parameter notify_data : no type generator for gpointer (gpointer) for param notify_data

// Unsupported : g_closure_add_marshal_guards : unsupported parameter pre_marshal_data : no type generator for gpointer (gpointer) for param pre_marshal_data

// Blacklisted : g_closure_invalidate

// Unsupported : g_closure_invoke : unsupported parameter param_values :

// Blacklisted : g_closure_ref

// Unsupported : g_closure_remove_finalize_notifier : unsupported parameter notify_data : no type generator for gpointer (gpointer) for param notify_data

// Unsupported : g_closure_remove_invalidate_notifier : unsupported parameter notify_data : no type generator for gpointer (gpointer) for param notify_data

// Unsupported : g_closure_set_marshal : unsupported parameter marshal : no type generator for ClosureMarshal (GClosureMarshal) for param marshal

// Unsupported : g_closure_set_meta_marshal : unsupported parameter marshal_data : no type generator for gpointer (gpointer) for param marshal_data

// Blacklisted : g_closure_sink

// Blacklisted : g_closure_unref

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

// Equals compares this ClosureNotifyData with another ClosureNotifyData, and returns true if they represent the same GObject.
func (recv *ClosureNotifyData) Equals(other *ClosureNotifyData) bool {
	return other.ToC() == recv.ToC()
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
	recv.native.minimum =
		(C.gint)(recv.Minimum)
	recv.native.maximum =
		(C.gint)(recv.Maximum)
	recv.native.n_values =
		(C.guint)(recv.NValues)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EnumClass with another EnumClass, and returns true if they represent the same GObject.
func (recv *EnumClass) Equals(other *EnumClass) bool {
	return other.ToC() == recv.ToC()
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
	recv.native.value =
		(C.gint)(recv.Value)
	recv.native.value_name =
		C.CString(recv.ValueName)
	recv.native.value_nick =
		C.CString(recv.ValueNick)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EnumValue with another EnumValue, and returns true if they represent the same GObject.
func (recv *EnumValue) Equals(other *EnumValue) bool {
	return other.ToC() == recv.ToC()
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
	recv.native.mask =
		(C.guint)(recv.Mask)
	recv.native.n_values =
		(C.guint)(recv.NValues)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FlagsClass with another FlagsClass, and returns true if they represent the same GObject.
func (recv *FlagsClass) Equals(other *FlagsClass) bool {
	return other.ToC() == recv.ToC()
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
	recv.native.value =
		(C.guint)(recv.Value)
	recv.native.value_name =
		C.CString(recv.ValueName)
	recv.native.value_nick =
		C.CString(recv.ValueNick)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FlagsValue with another FlagsValue, and returns true if they represent the same GObject.
func (recv *FlagsValue) Equals(other *FlagsValue) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this InitiallyUnownedClass with another InitiallyUnownedClass, and returns true if they represent the same GObject.
func (recv *InitiallyUnownedClass) Equals(other *InitiallyUnownedClass) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this InterfaceInfo with another InterfaceInfo, and returns true if they represent the same GObject.
func (recv *InterfaceInfo) Equals(other *InterfaceInfo) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this ObjectClass with another ObjectClass, and returns true if they represent the same GObject.
func (recv *ObjectClass) Equals(other *ObjectClass) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_object_class_find_property

// Blacklisted : g_object_class_install_property

// Unsupported : g_object_class_list_properties : array return type :

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

// Equals compares this ObjectConstructParam with another ObjectConstructParam, and returns true if they represent the same GObject.
func (recv *ObjectConstructParam) Equals(other *ObjectConstructParam) bool {
	return other.ToC() == recv.ToC()
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
	recv.native.value_type =
		(C.GType)(recv.ValueType)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecClass with another ParamSpecClass, and returns true if they represent the same GObject.
func (recv *ParamSpecClass) Equals(other *ParamSpecClass) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this ParamSpecPool with another ParamSpecPool, and returns true if they represent the same GObject.
func (recv *ParamSpecPool) Equals(other *ParamSpecPool) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_param_spec_pool_new

// Blacklisted : g_param_spec_pool_insert

// Unsupported : g_param_spec_pool_list : array return type :

// Blacklisted : g_param_spec_pool_list_owned

// Blacklisted : g_param_spec_pool_lookup

// Blacklisted : g_param_spec_pool_remove

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
	recv.native.instance_size =
		(C.guint16)(recv.InstanceSize)
	recv.native.n_preallocs =
		(C.guint16)(recv.NPreallocs)
	recv.native.value_type =
		(C.GType)(recv.ValueType)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecTypeInfo with another ParamSpecTypeInfo, and returns true if they represent the same GObject.
func (recv *ParamSpecTypeInfo) Equals(other *ParamSpecTypeInfo) bool {
	return other.ToC() == recv.ToC()
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
	recv.native.name =
		C.CString(recv.Name)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Parameter with another Parameter, and returns true if they represent the same GObject.
func (recv *Parameter) Equals(other *Parameter) bool {
	return other.ToC() == recv.ToC()
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
	recv.native.signal_id =
		(C.guint)(recv.SignalId)
	recv.native.detail =
		(C.GQuark)(recv.Detail)
	recv.native.run_type =
		(C.GSignalFlags)(recv.RunType)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SignalInvocationHint with another SignalInvocationHint, and returns true if they represent the same GObject.
func (recv *SignalInvocationHint) Equals(other *SignalInvocationHint) bool {
	return other.ToC() == recv.ToC()
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
	recv.native.signal_id =
		(C.guint)(recv.SignalId)
	recv.native.signal_name =
		C.CString(recv.SignalName)
	recv.native.itype =
		(C.GType)(recv.Itype)
	recv.native.signal_flags =
		(C.GSignalFlags)(recv.SignalFlags)
	recv.native.return_type =
		(C.GType)(recv.ReturnType)
	recv.native.n_params =
		(C.guint)(recv.NParams)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SignalQuery with another SignalQuery, and returns true if they represent the same GObject.
func (recv *SignalQuery) Equals(other *SignalQuery) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this TypeClass with another TypeClass, and returns true if they represent the same GObject.
func (recv *TypeClass) Equals(other *TypeClass) bool {
	return other.ToC() == recv.ToC()
}

// g_type_class_adjust_private_offset : unsupported parameter g_class : no type generator for gpointer (gpointer) for param g_class
// Blacklisted : g_type_class_peek

// Blacklisted : g_type_class_ref

// Unsupported : g_type_class_get_private : no return generator

// Blacklisted : g_type_class_peek_parent

// Blacklisted : g_type_class_unref

// Blacklisted : g_type_class_unref_uncached

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
	recv.native.type_flags =
		(C.GTypeFundamentalFlags)(recv.TypeFlags)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypeFundamentalInfo with another TypeFundamentalInfo, and returns true if they represent the same GObject.
func (recv *TypeFundamentalInfo) Equals(other *TypeFundamentalInfo) bool {
	return other.ToC() == recv.ToC()
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
	recv.native.class_size =
		(C.guint16)(recv.ClassSize)
	recv.native.instance_size =
		(C.guint16)(recv.InstanceSize)
	recv.native.n_preallocs =
		(C.guint16)(recv.NPreallocs)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypeInfo with another TypeInfo, and returns true if they represent the same GObject.
func (recv *TypeInfo) Equals(other *TypeInfo) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this TypeInstance with another TypeInstance, and returns true if they represent the same GObject.
func (recv *TypeInstance) Equals(other *TypeInstance) bool {
	return other.ToC() == recv.ToC()
}

// Unsupported : g_type_instance_get_private : no return generator

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

// Equals compares this TypeInterface with another TypeInterface, and returns true if they represent the same GObject.
func (recv *TypeInterface) Equals(other *TypeInterface) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_type_interface_add_prerequisite

// Blacklisted : g_type_interface_get_plugin

// Blacklisted : g_type_interface_peek

// Blacklisted : g_type_interface_peek_parent

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

// Equals compares this TypeModuleClass with another TypeModuleClass, and returns true if they represent the same GObject.
func (recv *TypeModuleClass) Equals(other *TypeModuleClass) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this TypePluginClass with another TypePluginClass, and returns true if they represent the same GObject.
func (recv *TypePluginClass) Equals(other *TypePluginClass) bool {
	return other.ToC() == recv.ToC()
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
	recv.native._type =
		(C.GType)(recv.Type)
	recv.native.type_name =
		C.CString(recv.TypeName)
	recv.native.class_size =
		(C.guint)(recv.ClassSize)
	recv.native.instance_size =
		(C.guint)(recv.InstanceSize)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypeQuery with another TypeQuery, and returns true if they represent the same GObject.
func (recv *TypeQuery) Equals(other *TypeQuery) bool {
	return other.ToC() == recv.ToC()
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
	recv.native.collect_format =
		C.CString(recv.CollectFormat)
	recv.native.lcopy_format =
		C.CString(recv.LcopyFormat)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypeValueTable with another TypeValueTable, and returns true if they represent the same GObject.
func (recv *TypeValueTable) Equals(other *TypeValueTable) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_type_value_table_peek

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

// Equals compares this Value with another Value, and returns true if they represent the same GObject.
func (recv *Value) Equals(other *Value) bool {
	return other.ToC() == recv.ToC()
}

// g_value_register_transform_func : unsupported parameter transform_func : no type generator for ValueTransform (GValueTransform) for param transform_func
// Blacklisted : g_value_type_compatible

// Blacklisted : g_value_type_transformable

// Blacklisted : g_value_copy

// Unsupported : g_value_dup_boxed : no return generator

// Blacklisted : g_value_dup_object

// Blacklisted : g_value_dup_param

// Blacklisted : g_value_dup_string

// Blacklisted : g_value_fits_pointer

// Blacklisted : g_value_get_boolean

// Unsupported : g_value_get_boxed : no return generator

// Blacklisted : g_value_get_char

// Blacklisted : g_value_get_double

// Blacklisted : g_value_get_enum

// Blacklisted : g_value_get_flags

// Blacklisted : g_value_get_float

// Blacklisted : g_value_get_int

// Blacklisted : g_value_get_int64

// Blacklisted : g_value_get_long

// Blacklisted : g_value_get_object

// Blacklisted : g_value_get_param

// Unsupported : g_value_get_pointer : no return generator

// Blacklisted : g_value_get_string

// Blacklisted : g_value_get_uchar

// Blacklisted : g_value_get_uint

// Blacklisted : g_value_get_uint64

// Blacklisted : g_value_get_ulong

// Init is a wrapper around the C function g_value_init.
func (recv *Value) Init(gType Type) *Value {
	c_g_type := (C.GType)(gType)

	retC := C.g_value_init((*C.GValue)(recv.native), c_g_type)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_value_peek_pointer : no return generator

// Blacklisted : g_value_reset

// Blacklisted : g_value_set_boolean

// Unsupported : g_value_set_boxed : unsupported parameter v_boxed : no type generator for gpointer (gconstpointer) for param v_boxed

// Unsupported : g_value_set_boxed_take_ownership : unsupported parameter v_boxed : no type generator for gpointer (gconstpointer) for param v_boxed

// Blacklisted : g_value_set_char

// Blacklisted : g_value_set_double

// Blacklisted : g_value_set_enum

// Blacklisted : g_value_set_flags

// Blacklisted : g_value_set_float

// Unsupported : g_value_set_instance : unsupported parameter instance : no type generator for gpointer (gpointer) for param instance

// Blacklisted : g_value_set_int

// Blacklisted : g_value_set_int64

// Blacklisted : g_value_set_long

// Blacklisted : g_value_set_object

// Unsupported : g_value_set_object_take_ownership : unsupported parameter v_object : no type generator for gpointer (gpointer) for param v_object

// Blacklisted : g_value_set_param

// Blacklisted : g_value_set_param_take_ownership

// Unsupported : g_value_set_pointer : unsupported parameter v_pointer : no type generator for gpointer (gpointer) for param v_pointer

// Unsupported : g_value_set_static_boxed : unsupported parameter v_boxed : no type generator for gpointer (gconstpointer) for param v_boxed

// Blacklisted : g_value_set_static_string

// Blacklisted : g_value_set_string

// Blacklisted : g_value_set_string_take_ownership

// Blacklisted : g_value_set_uchar

// Blacklisted : g_value_set_uint

// Blacklisted : g_value_set_uint64

// Blacklisted : g_value_set_ulong

// Blacklisted : g_value_transform

// Blacklisted : g_value_unset

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
	recv.native.n_values =
		(C.guint)(recv.NValues)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ValueArray with another ValueArray, and returns true if they represent the same GObject.
func (recv *ValueArray) Equals(other *ValueArray) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_value_array_new

// Blacklisted : g_value_array_append

// Blacklisted : g_value_array_copy

// Blacklisted : g_value_array_free

// Blacklisted : g_value_array_get_nth

// Blacklisted : g_value_array_insert

// Blacklisted : g_value_array_prepend

// Blacklisted : g_value_array_remove

// Unsupported : g_value_array_sort : unsupported parameter compare_func : no type generator for GLib.CompareFunc (GCompareFunc) for param compare_func

// Unsupported : g_value_array_sort_with_data : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc (GCompareDataFunc) for param compare_func

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

// Equals compares this WeakRef with another WeakRef, and returns true if they represent the same GObject.
func (recv *WeakRef) Equals(other *WeakRef) bool {
	return other.ToC() == recv.ToC()
}
