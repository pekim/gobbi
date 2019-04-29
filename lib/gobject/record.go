// This is a generated file - DO NOT EDIT

package gobject

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// CClosure is a wrapper around the C record GCClosure.
type CClosure struct {
	native unsafe.Pointer
	// closure : record
	// callback : no type generator for gpointer, gpointer
}

// Closure is a wrapper around the C record GClosure.
type Closure struct {
	native unsafe.Pointer
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

// ClosureNotifyData is a wrapper around the C record GClosureNotifyData.
type ClosureNotifyData struct {
	native unsafe.Pointer
	// data : no type generator for gpointer, gpointer
	// notify : no type generator for ClosureNotify, GClosureNotify
}

// EnumClass is a wrapper around the C record GEnumClass.
type EnumClass struct {
	native unsafe.Pointer
	// g_type_class : record
	Minimum int32
	Maximum int32
	NValues uint32
	// values : record
}

// EnumValue is a wrapper around the C record GEnumValue.
type EnumValue struct {
	native    unsafe.Pointer
	Value     int32
	ValueName string
	ValueNick string
}

// FlagsClass is a wrapper around the C record GFlagsClass.
type FlagsClass struct {
	native unsafe.Pointer
	// g_type_class : record
	Mask    uint32
	NValues uint32
	// values : record
}

// FlagsValue is a wrapper around the C record GFlagsValue.
type FlagsValue struct {
	native    unsafe.Pointer
	Value     uint32
	ValueName string
	ValueNick string
}

// InitiallyUnownedClass is a wrapper around the C record GInitiallyUnownedClass.
type InitiallyUnownedClass struct {
	native unsafe.Pointer
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

// InterfaceInfo is a wrapper around the C record GInterfaceInfo.
type InterfaceInfo struct {
	native unsafe.Pointer
	// interface_init : no type generator for InterfaceInitFunc, GInterfaceInitFunc
	// interface_finalize : no type generator for InterfaceFinalizeFunc, GInterfaceFinalizeFunc
	// interface_data : no type generator for gpointer, gpointer
}

// ObjectClass is a wrapper around the C record GObjectClass.
type ObjectClass struct {
	native unsafe.Pointer
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

// ObjectConstructParam is a wrapper around the C record GObjectConstructParam.
type ObjectConstructParam struct {
	native unsafe.Pointer
	// pspec : record
	// value : record
}

// ParamSpecClass is a wrapper around the C record GParamSpecClass.
type ParamSpecClass struct {
	native unsafe.Pointer
	// g_type_class : record
	ValueType Type
	// no type for finalize
	// no type for value_set_default
	// no type for value_validate
	// no type for values_cmp
	// Private : dummy
}

// ParamSpecPool is a wrapper around the C record GParamSpecPool.
type ParamSpecPool struct {
	native unsafe.Pointer
}

// ParamSpecTypeInfo is a wrapper around the C record GParamSpecTypeInfo.
type ParamSpecTypeInfo struct {
	native       unsafe.Pointer
	InstanceSize uint16
	NPreallocs   uint16
	// no type for instance_init
	ValueType Type
	// no type for finalize
	// no type for value_set_default
	// no type for value_validate
	// no type for values_cmp
}

// Parameter is a wrapper around the C record GParameter.
type Parameter struct {
	native unsafe.Pointer
	Name   string
	// value : record
}

// SignalInvocationHint is a wrapper around the C record GSignalInvocationHint.
type SignalInvocationHint struct {
	native   unsafe.Pointer
	SignalId uint32
	Detail   glib.Quark
	RunType  SignalFlags
}

// SignalQuery is a wrapper around the C record GSignalQuery.
type SignalQuery struct {
	native      unsafe.Pointer
	SignalId    uint32
	SignalName  string
	Itype       Type
	SignalFlags SignalFlags
	ReturnType  Type
	NParams     uint32
	// no type for param_types
}

// TypeClass is a wrapper around the C record GTypeClass.
type TypeClass struct {
	native unsafe.Pointer
	// Private : g_type
}

// TypeFundamentalInfo is a wrapper around the C record GTypeFundamentalInfo.
type TypeFundamentalInfo struct {
	native    unsafe.Pointer
	TypeFlags TypeFundamentalFlags
}

// TypeInfo is a wrapper around the C record GTypeInfo.
type TypeInfo struct {
	native    unsafe.Pointer
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

// TypeInstance is a wrapper around the C record GTypeInstance.
type TypeInstance struct {
	native unsafe.Pointer
	// Private : g_class
}

// TypeInterface is a wrapper around the C record GTypeInterface.
type TypeInterface struct {
	native unsafe.Pointer
	// Private : g_type
	// Private : g_instance_type
}

// TypeModuleClass is a wrapper around the C record GTypeModuleClass.
type TypeModuleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for load
	// no type for unload
	// no type for reserved1
	// no type for reserved2
	// no type for reserved3
	// no type for reserved4
}

// TypePluginClass is a wrapper around the C record GTypePluginClass.
type TypePluginClass struct {
	native unsafe.Pointer
	// Private : base_iface
	// use_plugin : no type generator for TypePluginUse, GTypePluginUse
	// unuse_plugin : no type generator for TypePluginUnuse, GTypePluginUnuse
	// complete_type_info : no type generator for TypePluginCompleteTypeInfo, GTypePluginCompleteTypeInfo
	// complete_interface_info : no type generator for TypePluginCompleteInterfaceInfo, GTypePluginCompleteInterfaceInfo
}

// TypeQuery is a wrapper around the C record GTypeQuery.
type TypeQuery struct {
	native       unsafe.Pointer
	Type         Type
	TypeName     string
	ClassSize    uint32
	InstanceSize uint32
}

// TypeValueTable is a wrapper around the C record GTypeValueTable.
type TypeValueTable struct {
	native unsafe.Pointer
	// no type for value_init
	// no type for value_free
	// no type for value_copy
	// no type for value_peek_pointer
	CollectFormat string
	// no type for collect_value
	LcopyFormat string
	// no type for lcopy_value
}

// Value is a wrapper around the C record GValue.
type Value struct {
	native unsafe.Pointer
	// Private : g_type
	// no type for data
}

// ValueArray is a wrapper around the C record GValueArray.
type ValueArray struct {
	native  unsafe.Pointer
	NValues uint32
	// values : record
	// Private : n_prealloced
}

// WeakRef is a wrapper around the C record GWeakRef.
type WeakRef struct {
	native unsafe.Pointer
}
