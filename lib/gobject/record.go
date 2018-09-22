// This is a generated file - DO NOT EDIT

package gobject

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// CClosure is a wrapper around the C record GCClosure.
type CClosure struct {
	native *C.GCClosure
	// closure : record
	Callback uintptr
}

func cClosureNewFromC(c *C.GCClosure) *CClosure {
	if c == nil {
		return nil
	}

	g := &CClosure{
		Callback: (uintptr)(c.callback),
		native:   c,
	}

	return g
}

func (recv *CClosure) toC() *C.GCClosure {
	recv.native.callback =
		(C.gpointer)(recv.Callback)

	return recv.native
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

func closureNewFromC(c *C.GClosure) *Closure {
	if c == nil {
		return nil
	}

	g := &Closure{native: c}

	return g
}

func (recv *Closure) toC() *C.GClosure {

	return recv.native
}

// Unsupported : g_closure_new_object : unsupported parameter object : no type generator for Object, GObject*

// ClosureNewSimple is a wrapper around the C function g_closure_new_simple.
func ClosureNewSimple(sizeofClosure uint32, data uintptr) *Closure {
	c_sizeof_closure := (C.guint)(sizeofClosure)

	c_data := (C.gpointer)(data)

	retC := C.g_closure_new_simple(c_sizeof_closure, c_data)
	retGo := closureNewFromC(retC)

	return retGo
}

// Unsupported : g_closure_add_finalize_notifier : unsupported parameter notify_func : no type generator for ClosureNotify, GClosureNotify

// Unsupported : g_closure_add_invalidate_notifier : unsupported parameter notify_func : no type generator for ClosureNotify, GClosureNotify

// Unsupported : g_closure_add_marshal_guards : unsupported parameter pre_marshal_notify : no type generator for ClosureNotify, GClosureNotify

// Unsupported : g_closure_invalidate : no return generator

// Unsupported : g_closure_invoke : unsupported parameter return_value : record param - coming soon

// Ref is a wrapper around the C function g_closure_ref.
func (recv *Closure) Ref() *Closure {
	retC := C.g_closure_ref((*C.GClosure)(recv.native))
	retGo := closureNewFromC(retC)

	return retGo
}

// Unsupported : g_closure_remove_finalize_notifier : unsupported parameter notify_func : no type generator for ClosureNotify, GClosureNotify

// Unsupported : g_closure_remove_invalidate_notifier : unsupported parameter notify_func : no type generator for ClosureNotify, GClosureNotify

// Unsupported : g_closure_set_marshal : unsupported parameter marshal : no type generator for ClosureMarshal, GClosureMarshal

// Unsupported : g_closure_set_meta_marshal : unsupported parameter meta_marshal : no type generator for ClosureMarshal, GClosureMarshal

// Unsupported : g_closure_sink : no return generator

// Unsupported : g_closure_unref : no return generator

// ClosureNotifyData is a wrapper around the C record GClosureNotifyData.
type ClosureNotifyData struct {
	native *C.GClosureNotifyData
	Data   uintptr
	// notify : no type generator for ClosureNotify, GClosureNotify
}

func closureNotifyDataNewFromC(c *C.GClosureNotifyData) *ClosureNotifyData {
	if c == nil {
		return nil
	}

	g := &ClosureNotifyData{
		Data:   (uintptr)(c.data),
		native: c,
	}

	return g
}

func (recv *ClosureNotifyData) toC() *C.GClosureNotifyData {
	recv.native.data =
		(C.gpointer)(recv.Data)

	return recv.native
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

func enumClassNewFromC(c *C.GEnumClass) *EnumClass {
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

func (recv *EnumClass) toC() *C.GEnumClass {
	recv.native.minimum =
		(C.gint)(recv.Minimum)
	recv.native.maximum =
		(C.gint)(recv.Maximum)
	recv.native.n_values =
		(C.guint)(recv.NValues)

	return recv.native
}

// EnumValue is a wrapper around the C record GEnumValue.
type EnumValue struct {
	native    *C.GEnumValue
	Value     int32
	ValueName string
	ValueNick string
}

func enumValueNewFromC(c *C.GEnumValue) *EnumValue {
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

func (recv *EnumValue) toC() *C.GEnumValue {
	recv.native.value =
		(C.gint)(recv.Value)
	recv.native.value_name =
		C.CString(recv.ValueName)
	recv.native.value_nick =
		C.CString(recv.ValueNick)

	return recv.native
}

// FlagsClass is a wrapper around the C record GFlagsClass.
type FlagsClass struct {
	native *C.GFlagsClass
	// g_type_class : record
	Mask    uint32
	NValues uint32
	// values : record
}

func flagsClassNewFromC(c *C.GFlagsClass) *FlagsClass {
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

func (recv *FlagsClass) toC() *C.GFlagsClass {
	recv.native.mask =
		(C.guint)(recv.Mask)
	recv.native.n_values =
		(C.guint)(recv.NValues)

	return recv.native
}

// FlagsValue is a wrapper around the C record GFlagsValue.
type FlagsValue struct {
	native    *C.GFlagsValue
	Value     uint32
	ValueName string
	ValueNick string
}

func flagsValueNewFromC(c *C.GFlagsValue) *FlagsValue {
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

func (recv *FlagsValue) toC() *C.GFlagsValue {
	recv.native.value =
		(C.guint)(recv.Value)
	recv.native.value_name =
		C.CString(recv.ValueName)
	recv.native.value_nick =
		C.CString(recv.ValueNick)

	return recv.native
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

func initiallyUnownedClassNewFromC(c *C.GInitiallyUnownedClass) *InitiallyUnownedClass {
	if c == nil {
		return nil
	}

	g := &InitiallyUnownedClass{native: c}

	return g
}

func (recv *InitiallyUnownedClass) toC() *C.GInitiallyUnownedClass {

	return recv.native
}

// InterfaceInfo is a wrapper around the C record GInterfaceInfo.
type InterfaceInfo struct {
	native *C.GInterfaceInfo
	// interface_init : no type generator for InterfaceInitFunc, GInterfaceInitFunc
	// interface_finalize : no type generator for InterfaceFinalizeFunc, GInterfaceFinalizeFunc
	InterfaceData uintptr
}

func interfaceInfoNewFromC(c *C.GInterfaceInfo) *InterfaceInfo {
	if c == nil {
		return nil
	}

	g := &InterfaceInfo{
		InterfaceData: (uintptr)(c.interface_data),
		native:        c,
	}

	return g
}

func (recv *InterfaceInfo) toC() *C.GInterfaceInfo {
	recv.native.interface_data =
		(C.gpointer)(recv.InterfaceData)

	return recv.native
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

func objectClassNewFromC(c *C.GObjectClass) *ObjectClass {
	if c == nil {
		return nil
	}

	g := &ObjectClass{native: c}

	return g
}

func (recv *ObjectClass) toC() *C.GObjectClass {

	return recv.native
}

// Unsupported : g_object_class_find_property : no return generator

// Unsupported : g_object_class_install_properties : unsupported parameter pspecs : no param type

// Unsupported : g_object_class_install_property : unsupported parameter pspec : no type generator for ParamSpec, GParamSpec*

// Unsupported : g_object_class_list_properties : unsupported parameter n_properties : no type generator for guint, guint*

// Unsupported : g_object_class_override_property : no return generator

// ObjectConstructParam is a wrapper around the C record GObjectConstructParam.
type ObjectConstructParam struct {
	native *C.GObjectConstructParam
	// pspec : no type generator for ParamSpec, GParamSpec*
	// value : record
}

func objectConstructParamNewFromC(c *C.GObjectConstructParam) *ObjectConstructParam {
	if c == nil {
		return nil
	}

	g := &ObjectConstructParam{native: c}

	return g
}

func (recv *ObjectConstructParam) toC() *C.GObjectConstructParam {

	return recv.native
}

// ParamSpecClass is a wrapper around the C record GParamSpecClass.
type ParamSpecClass struct {
	native *C.GParamSpecClass
	// g_type_class : record
	// value_type : no type generator for GType, GType
	// no type for finalize
	// no type for value_set_default
	// no type for value_validate
	// no type for values_cmp
	// Private : dummy
}

func paramSpecClassNewFromC(c *C.GParamSpecClass) *ParamSpecClass {
	if c == nil {
		return nil
	}

	g := &ParamSpecClass{native: c}

	return g
}

func (recv *ParamSpecClass) toC() *C.GParamSpecClass {

	return recv.native
}

// ParamSpecPool is a wrapper around the C record GParamSpecPool.
type ParamSpecPool struct {
	native *C.GParamSpecPool
}

func paramSpecPoolNewFromC(c *C.GParamSpecPool) *ParamSpecPool {
	if c == nil {
		return nil
	}

	g := &ParamSpecPool{native: c}

	return g
}

func (recv *ParamSpecPool) toC() *C.GParamSpecPool {

	return recv.native
}

// Unsupported : g_param_spec_pool_insert : unsupported parameter pspec : no type generator for ParamSpec, GParamSpec*

// Unsupported : g_param_spec_pool_list : unsupported parameter owner_type : no type generator for GType, GType

// Unsupported : g_param_spec_pool_list_owned : unsupported parameter owner_type : no type generator for GType, GType

// Unsupported : g_param_spec_pool_lookup : unsupported parameter owner_type : no type generator for GType, GType

// Unsupported : g_param_spec_pool_remove : unsupported parameter pspec : no type generator for ParamSpec, GParamSpec*

// ParamSpecTypeInfo is a wrapper around the C record GParamSpecTypeInfo.
type ParamSpecTypeInfo struct {
	native       *C.GParamSpecTypeInfo
	InstanceSize uint16
	NPreallocs   uint16
	// no type for instance_init
	// value_type : no type generator for GType, GType
	// no type for finalize
	// no type for value_set_default
	// no type for value_validate
	// no type for values_cmp
}

func paramSpecTypeInfoNewFromC(c *C.GParamSpecTypeInfo) *ParamSpecTypeInfo {
	if c == nil {
		return nil
	}

	g := &ParamSpecTypeInfo{
		InstanceSize: (uint16)(c.instance_size),
		NPreallocs:   (uint16)(c.n_preallocs),
		native:       c,
	}

	return g
}

func (recv *ParamSpecTypeInfo) toC() *C.GParamSpecTypeInfo {
	recv.native.instance_size =
		(C.guint16)(recv.InstanceSize)
	recv.native.n_preallocs =
		(C.guint16)(recv.NPreallocs)

	return recv.native
}

// Parameter is a wrapper around the C record GParameter.
type Parameter struct {
	native *C.GParameter
	Name   string
	// value : record
}

func parameterNewFromC(c *C.GParameter) *Parameter {
	if c == nil {
		return nil
	}

	g := &Parameter{
		Name:   C.GoString(c.name),
		native: c,
	}

	return g
}

func (recv *Parameter) toC() *C.GParameter {
	recv.native.name =
		C.CString(recv.Name)

	return recv.native
}

// SignalInvocationHint is a wrapper around the C record GSignalInvocationHint.
type SignalInvocationHint struct {
	native   *C.GSignalInvocationHint
	SignalId uint32
	// detail : no type generator for GLib.Quark, GQuark
	// run_type : no type generator for SignalFlags, GSignalFlags
}

func signalInvocationHintNewFromC(c *C.GSignalInvocationHint) *SignalInvocationHint {
	if c == nil {
		return nil
	}

	g := &SignalInvocationHint{
		SignalId: (uint32)(c.signal_id),
		native:   c,
	}

	return g
}

func (recv *SignalInvocationHint) toC() *C.GSignalInvocationHint {
	recv.native.signal_id =
		(C.guint)(recv.SignalId)

	return recv.native
}

// SignalQuery is a wrapper around the C record GSignalQuery.
type SignalQuery struct {
	native     *C.GSignalQuery
	SignalId   uint32
	SignalName string
	// itype : no type generator for GType, GType
	// signal_flags : no type generator for SignalFlags, GSignalFlags
	// return_type : no type generator for GType, GType
	NParams uint32
	// no type for param_types
}

func signalQueryNewFromC(c *C.GSignalQuery) *SignalQuery {
	if c == nil {
		return nil
	}

	g := &SignalQuery{
		NParams:    (uint32)(c.n_params),
		SignalId:   (uint32)(c.signal_id),
		SignalName: C.GoString(c.signal_name),
		native:     c,
	}

	return g
}

func (recv *SignalQuery) toC() *C.GSignalQuery {
	recv.native.signal_id =
		(C.guint)(recv.SignalId)
	recv.native.signal_name =
		C.CString(recv.SignalName)
	recv.native.n_params =
		(C.guint)(recv.NParams)

	return recv.native
}

// TypeClass is a wrapper around the C record GTypeClass.
type TypeClass struct {
	native *C.GTypeClass
	// Private : g_type
}

func typeClassNewFromC(c *C.GTypeClass) *TypeClass {
	if c == nil {
		return nil
	}

	g := &TypeClass{native: c}

	return g
}

func (recv *TypeClass) toC() *C.GTypeClass {

	return recv.native
}

// Unsupported : g_type_class_add_private : no return generator

// Unsupported : g_type_class_get_private : unsupported parameter private_type : no type generator for GType, GType

// PeekParent is a wrapper around the C function g_type_class_peek_parent.
func (recv *TypeClass) PeekParent() uintptr {
	retC := C.g_type_class_peek_parent((C.gpointer)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_type_class_unref : no return generator

// Unsupported : g_type_class_unref_uncached : no return generator

// TypeFundamentalInfo is a wrapper around the C record GTypeFundamentalInfo.
type TypeFundamentalInfo struct {
	native *C.GTypeFundamentalInfo
	// type_flags : no type generator for TypeFundamentalFlags, GTypeFundamentalFlags
}

func typeFundamentalInfoNewFromC(c *C.GTypeFundamentalInfo) *TypeFundamentalInfo {
	if c == nil {
		return nil
	}

	g := &TypeFundamentalInfo{native: c}

	return g
}

func (recv *TypeFundamentalInfo) toC() *C.GTypeFundamentalInfo {

	return recv.native
}

// TypeInfo is a wrapper around the C record GTypeInfo.
type TypeInfo struct {
	native    *C.GTypeInfo
	ClassSize uint16
	// base_init : no type generator for BaseInitFunc, GBaseInitFunc
	// base_finalize : no type generator for BaseFinalizeFunc, GBaseFinalizeFunc
	// class_init : no type generator for ClassInitFunc, GClassInitFunc
	// class_finalize : no type generator for ClassFinalizeFunc, GClassFinalizeFunc
	ClassData    uintptr
	InstanceSize uint16
	NPreallocs   uint16
	// instance_init : no type generator for InstanceInitFunc, GInstanceInitFunc
	// value_table : record
}

func typeInfoNewFromC(c *C.GTypeInfo) *TypeInfo {
	if c == nil {
		return nil
	}

	g := &TypeInfo{
		ClassData:    (uintptr)(c.class_data),
		ClassSize:    (uint16)(c.class_size),
		InstanceSize: (uint16)(c.instance_size),
		NPreallocs:   (uint16)(c.n_preallocs),
		native:       c,
	}

	return g
}

func (recv *TypeInfo) toC() *C.GTypeInfo {
	recv.native.class_size =
		(C.guint16)(recv.ClassSize)
	recv.native.class_data =
		(C.gconstpointer)(recv.ClassData)
	recv.native.instance_size =
		(C.guint16)(recv.InstanceSize)
	recv.native.n_preallocs =
		(C.guint16)(recv.NPreallocs)

	return recv.native
}

// TypeInstance is a wrapper around the C record GTypeInstance.
type TypeInstance struct {
	native *C.GTypeInstance
	// Private : g_class
}

func typeInstanceNewFromC(c *C.GTypeInstance) *TypeInstance {
	if c == nil {
		return nil
	}

	g := &TypeInstance{native: c}

	return g
}

func (recv *TypeInstance) toC() *C.GTypeInstance {

	return recv.native
}

// Unsupported : g_type_instance_get_private : unsupported parameter private_type : no type generator for GType, GType

// TypeInterface is a wrapper around the C record GTypeInterface.
type TypeInterface struct {
	native *C.GTypeInterface
	// Private : g_type
	// Private : g_instance_type
}

func typeInterfaceNewFromC(c *C.GTypeInterface) *TypeInterface {
	if c == nil {
		return nil
	}

	g := &TypeInterface{native: c}

	return g
}

func (recv *TypeInterface) toC() *C.GTypeInterface {

	return recv.native
}

// PeekParent is a wrapper around the C function g_type_interface_peek_parent.
func (recv *TypeInterface) PeekParent() uintptr {
	retC := C.g_type_interface_peek_parent((C.gpointer)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
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

func typeModuleClassNewFromC(c *C.GTypeModuleClass) *TypeModuleClass {
	if c == nil {
		return nil
	}

	g := &TypeModuleClass{native: c}

	return g
}

func (recv *TypeModuleClass) toC() *C.GTypeModuleClass {

	return recv.native
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

func typePluginClassNewFromC(c *C.GTypePluginClass) *TypePluginClass {
	if c == nil {
		return nil
	}

	g := &TypePluginClass{native: c}

	return g
}

func (recv *TypePluginClass) toC() *C.GTypePluginClass {

	return recv.native
}

// TypeQuery is a wrapper around the C record GTypeQuery.
type TypeQuery struct {
	native *C.GTypeQuery
	// _type : no type generator for GType, GType
	TypeName     string
	ClassSize    uint32
	InstanceSize uint32
}

func typeQueryNewFromC(c *C.GTypeQuery) *TypeQuery {
	if c == nil {
		return nil
	}

	g := &TypeQuery{
		ClassSize:    (uint32)(c.class_size),
		InstanceSize: (uint32)(c.instance_size),
		TypeName:     C.GoString(c.type_name),
		native:       c,
	}

	return g
}

func (recv *TypeQuery) toC() *C.GTypeQuery {
	recv.native.type_name =
		C.CString(recv.TypeName)
	recv.native.class_size =
		(C.guint)(recv.ClassSize)
	recv.native.instance_size =
		(C.guint)(recv.InstanceSize)

	return recv.native
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

func typeValueTableNewFromC(c *C.GTypeValueTable) *TypeValueTable {
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

func (recv *TypeValueTable) toC() *C.GTypeValueTable {
	recv.native.collect_format =
		C.CString(recv.CollectFormat)
	recv.native.lcopy_format =
		C.CString(recv.LcopyFormat)

	return recv.native
}

// Value is a wrapper around the C record GValue.
type Value struct {
	native *C.GValue
	// Private : g_type
	// no type for data
}

func valueNewFromC(c *C.GValue) *Value {
	if c == nil {
		return nil
	}

	g := &Value{native: c}

	return g
}

func (recv *Value) toC() *C.GValue {

	return recv.native
}

// Unsupported : g_value_copy : unsupported parameter dest_value : record param - coming soon

// DupBoxed is a wrapper around the C function g_value_dup_boxed.
func (recv *Value) DupBoxed() uintptr {
	retC := C.g_value_dup_boxed((*C.GValue)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// DupObject is a wrapper around the C function g_value_dup_object.
func (recv *Value) DupObject() uintptr {
	retC := C.g_value_dup_object((*C.GValue)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_value_dup_param : no return generator

// DupString is a wrapper around the C function g_value_dup_string.
func (recv *Value) DupString() string {
	retC := C.g_value_dup_string((*C.GValue)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_value_dup_variant : no return generator

// Unsupported : g_value_fits_pointer : no return generator

// Unsupported : g_value_get_boolean : no return generator

// GetBoxed is a wrapper around the C function g_value_get_boxed.
func (recv *Value) GetBoxed() uintptr {
	retC := C.g_value_get_boxed((*C.GValue)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// GetChar is a wrapper around the C function g_value_get_char.
func (recv *Value) GetChar() rune {
	retC := C.g_value_get_char((*C.GValue)(recv.native))
	retGo := (rune)(retC)

	return retGo
}

// GetDouble is a wrapper around the C function g_value_get_double.
func (recv *Value) GetDouble() float64 {
	retC := C.g_value_get_double((*C.GValue)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetEnum is a wrapper around the C function g_value_get_enum.
func (recv *Value) GetEnum() int32 {
	retC := C.g_value_get_enum((*C.GValue)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetFlags is a wrapper around the C function g_value_get_flags.
func (recv *Value) GetFlags() uint32 {
	retC := C.g_value_get_flags((*C.GValue)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetFloat is a wrapper around the C function g_value_get_float.
func (recv *Value) GetFloat() float32 {
	retC := C.g_value_get_float((*C.GValue)(recv.native))
	retGo := (float32)(retC)

	return retGo
}

// Unsupported : g_value_get_gtype : no return generator

// GetInt is a wrapper around the C function g_value_get_int.
func (recv *Value) GetInt() int32 {
	retC := C.g_value_get_int((*C.GValue)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetInt64 is a wrapper around the C function g_value_get_int64.
func (recv *Value) GetInt64() int64 {
	retC := C.g_value_get_int64((*C.GValue)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetLong is a wrapper around the C function g_value_get_long.
func (recv *Value) GetLong() int64 {
	retC := C.g_value_get_long((*C.GValue)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetObject is a wrapper around the C function g_value_get_object.
func (recv *Value) GetObject() uintptr {
	retC := C.g_value_get_object((*C.GValue)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_value_get_param : no return generator

// GetPointer is a wrapper around the C function g_value_get_pointer.
func (recv *Value) GetPointer() uintptr {
	retC := C.g_value_get_pointer((*C.GValue)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// GetString is a wrapper around the C function g_value_get_string.
func (recv *Value) GetString() string {
	retC := C.g_value_get_string((*C.GValue)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUchar is a wrapper around the C function g_value_get_uchar.
func (recv *Value) GetUchar() uint8 {
	retC := C.g_value_get_uchar((*C.GValue)(recv.native))
	retGo := (uint8)(retC)

	return retGo
}

// GetUint is a wrapper around the C function g_value_get_uint.
func (recv *Value) GetUint() uint32 {
	retC := C.g_value_get_uint((*C.GValue)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetUint64 is a wrapper around the C function g_value_get_uint64.
func (recv *Value) GetUint64() uint64 {
	retC := C.g_value_get_uint64((*C.GValue)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// GetUlong is a wrapper around the C function g_value_get_ulong.
func (recv *Value) GetUlong() uint64 {
	retC := C.g_value_get_ulong((*C.GValue)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported : g_value_get_variant : no return generator

// Unsupported : g_value_init : unsupported parameter g_type : no type generator for GType, GType

// Unsupported : g_value_init_from_instance : no return generator

// PeekPointer is a wrapper around the C function g_value_peek_pointer.
func (recv *Value) PeekPointer() uintptr {
	retC := C.g_value_peek_pointer((*C.GValue)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// Reset is a wrapper around the C function g_value_reset.
func (recv *Value) Reset() *Value {
	retC := C.g_value_reset((*C.GValue)(recv.native))
	retGo := valueNewFromC(retC)

	return retGo
}

// Unsupported : g_value_set_boolean : unsupported parameter v_boolean : no type generator for gboolean, gboolean

// Unsupported : g_value_set_boxed : no return generator

// Unsupported : g_value_set_boxed_take_ownership : no return generator

// Unsupported : g_value_set_char : no return generator

// Unsupported : g_value_set_double : no return generator

// Unsupported : g_value_set_enum : no return generator

// Unsupported : g_value_set_flags : no return generator

// Unsupported : g_value_set_float : no return generator

// Unsupported : g_value_set_gtype : unsupported parameter v_gtype : no type generator for GType, GType

// Unsupported : g_value_set_instance : no return generator

// Unsupported : g_value_set_int : no return generator

// Unsupported : g_value_set_int64 : no return generator

// Unsupported : g_value_set_long : no return generator

// Unsupported : g_value_set_object : no return generator

// Unsupported : g_value_set_object_take_ownership : no return generator

// Unsupported : g_value_set_param : unsupported parameter param : no type generator for ParamSpec, GParamSpec*

// Unsupported : g_value_set_param_take_ownership : unsupported parameter param : no type generator for ParamSpec, GParamSpec*

// Unsupported : g_value_set_pointer : no return generator

// Unsupported : g_value_set_schar : no return generator

// Unsupported : g_value_set_static_boxed : no return generator

// Unsupported : g_value_set_static_string : no return generator

// Unsupported : g_value_set_string : no return generator

// Unsupported : g_value_set_string_take_ownership : no return generator

// Unsupported : g_value_set_uchar : no return generator

// Unsupported : g_value_set_uint : no return generator

// Unsupported : g_value_set_uint64 : no return generator

// Unsupported : g_value_set_ulong : no return generator

// Unsupported : g_value_set_variant : unsupported parameter variant : no type generator for GLib.Variant, GVariant*

// Unsupported : g_value_take_boxed : no return generator

// Unsupported : g_value_take_object : no return generator

// Unsupported : g_value_take_param : unsupported parameter param : no type generator for ParamSpec, GParamSpec*

// Unsupported : g_value_take_string : no return generator

// Unsupported : g_value_take_variant : unsupported parameter variant : no type generator for GLib.Variant, GVariant*

// Unsupported : g_value_transform : unsupported parameter dest_value : record param - coming soon

// Unsupported : g_value_unset : no return generator

// ValueArray is a wrapper around the C record GValueArray.
type ValueArray struct {
	native  *C.GValueArray
	NValues uint32
	// values : record
	// Private : n_prealloced
}

func valueArrayNewFromC(c *C.GValueArray) *ValueArray {
	if c == nil {
		return nil
	}

	g := &ValueArray{
		NValues: (uint32)(c.n_values),
		native:  c,
	}

	return g
}

func (recv *ValueArray) toC() *C.GValueArray {
	recv.native.n_values =
		(C.guint)(recv.NValues)

	return recv.native
}

// ValueArrayNew is a wrapper around the C function g_value_array_new.
func ValueArrayNew(nPrealloced uint32) *ValueArray {
	c_n_prealloced := (C.guint)(nPrealloced)

	retC := C.g_value_array_new(c_n_prealloced)
	retGo := valueArrayNewFromC(retC)

	return retGo
}

// Unsupported : g_value_array_append : unsupported parameter value : record param - coming soon

// Copy is a wrapper around the C function g_value_array_copy.
func (recv *ValueArray) Copy() *ValueArray {
	retC := C.g_value_array_copy((*C.GValueArray)(recv.native))
	retGo := valueArrayNewFromC(retC)

	return retGo
}

// Unsupported : g_value_array_free : no return generator

// GetNth is a wrapper around the C function g_value_array_get_nth.
func (recv *ValueArray) GetNth(index uint32) *Value {
	c_index_ := (C.guint)(index)

	retC := C.g_value_array_get_nth((*C.GValueArray)(recv.native), c_index_)
	retGo := valueNewFromC(retC)

	return retGo
}

// Unsupported : g_value_array_insert : unsupported parameter value : record param - coming soon

// Unsupported : g_value_array_prepend : unsupported parameter value : record param - coming soon

// Remove is a wrapper around the C function g_value_array_remove.
func (recv *ValueArray) Remove(index uint32) *ValueArray {
	c_index_ := (C.guint)(index)

	retC := C.g_value_array_remove((*C.GValueArray)(recv.native), c_index_)
	retGo := valueArrayNewFromC(retC)

	return retGo
}

// Unsupported : g_value_array_sort : unsupported parameter compare_func : no type generator for GLib.CompareFunc, GCompareFunc

// Unsupported : g_value_array_sort_with_data : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc, GCompareDataFunc

// WeakRef is a wrapper around the C record GWeakRef.
type WeakRef struct {
	native *C.GWeakRef
}

func weakRefNewFromC(c *C.GWeakRef) *WeakRef {
	if c == nil {
		return nil
	}

	g := &WeakRef{native: c}

	return g
}

func (recv *WeakRef) toC() *C.GWeakRef {

	return recv.native
}

// Unsupported : g_weak_ref_clear : no return generator

// Unsupported : g_weak_ref_init : no return generator

// Unsupported : g_weak_ref_set : no return generator
