// Code generated - DO NOT EDIT.

package gobject

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var CClosureStruct *gi.Struct
var CClosureStructOnce sync.Once

func CClosureStructSet() {
	CClosureStructOnce.Do(func() {
		CClosureStruct = gi.StructNew("GObject", "CClosure")
	})
}

type CClosure struct {
	native  uintptr
	Closure *Closure
	// UNSUPPORTED : C value 'callback' : no Go type for 'gpointer'
}

var ClosureStruct *gi.Struct
var ClosureStructOnce sync.Once

func ClosureStructSet() {
	ClosureStructOnce.Do(func() {
		ClosureStruct = gi.StructNew("GObject", "Closure")
	})
}

type Closure struct {
	native    uintptr
	InMarshal uint32
	IsInvalid uint32
	// UNSUPPORTED : C value 'marshal' : missing Type
}

// UNSUPPORTED : C value 'g_closure_new_object' : parameter 'object' of type 'Object' not supported

// UNSUPPORTED : C value 'g_closure_new_simple' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_closure_add_finalize_notifier' : parameter 'notify_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_closure_add_invalidate_notifier' : parameter 'notify_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_closure_add_marshal_guards' : parameter 'pre_marshal_data' of type 'gpointer' not supported

var invalidateClosureInvoker *gi.Function

// Invalidate is a representation of the C type g_closure_invalidate.
func (recv *Closure) Invalidate() {
	if invalidateClosureInvoker == nil {
		invalidateClosureInvoker = gi.StructFunctionInvokerNew("GObject", "Closure", "invalidate")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	invalidateClosureInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_closure_invoke' : parameter 'return_value' of type 'Value' not supported

var refClosureInvoker *gi.Function

// Ref is a representation of the C type g_closure_ref.
func (recv *Closure) Ref() *Closure {
	if refClosureInvoker == nil {
		refClosureInvoker = gi.StructFunctionInvokerNew("GObject", "Closure", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refClosureInvoker.Invoke(inArgs[:], nil)

	retGo := &Closure{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_closure_remove_finalize_notifier' : parameter 'notify_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_closure_remove_invalidate_notifier' : parameter 'notify_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_closure_set_marshal' : parameter 'marshal' of type 'ClosureMarshal' not supported

// UNSUPPORTED : C value 'g_closure_set_meta_marshal' : parameter 'marshal_data' of type 'gpointer' not supported

var sinkClosureInvoker *gi.Function

// Sink is a representation of the C type g_closure_sink.
func (recv *Closure) Sink() {
	if sinkClosureInvoker == nil {
		sinkClosureInvoker = gi.StructFunctionInvokerNew("GObject", "Closure", "sink")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	sinkClosureInvoker.Invoke(inArgs[:], nil)

}

var unrefClosureInvoker *gi.Function

// Unref is a representation of the C type g_closure_unref.
func (recv *Closure) Unref() {
	if unrefClosureInvoker == nil {
		unrefClosureInvoker = gi.StructFunctionInvokerNew("GObject", "Closure", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefClosureInvoker.Invoke(inArgs[:], nil)

}

var ClosureNotifyDataStruct *gi.Struct
var ClosureNotifyDataStructOnce sync.Once

func ClosureNotifyDataStructSet() {
	ClosureNotifyDataStructOnce.Do(func() {
		ClosureNotifyDataStruct = gi.StructNew("GObject", "ClosureNotifyData")
	})
}

type ClosureNotifyData struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'
	// UNSUPPORTED : C value 'notify' : no Go type for 'ClosureNotify'
}

var EnumClassStruct *gi.Struct
var EnumClassStructOnce sync.Once

func EnumClassStructSet() {
	EnumClassStructOnce.Do(func() {
		EnumClassStruct = gi.StructNew("GObject", "EnumClass")
	})
}

type EnumClass struct {
	native     uintptr
	GTypeClass *TypeClass
	Minimum    int32
	Maximum    int32
	NValues    uint32
	Values     *EnumValue
}

var EnumValueStruct *gi.Struct
var EnumValueStructOnce sync.Once

func EnumValueStructSet() {
	EnumValueStructOnce.Do(func() {
		EnumValueStruct = gi.StructNew("GObject", "EnumValue")
	})
}

type EnumValue struct {
	native    uintptr
	Value     int32
	ValueName string
	ValueNick string
}

var FlagsClassStruct *gi.Struct
var FlagsClassStructOnce sync.Once

func FlagsClassStructSet() {
	FlagsClassStructOnce.Do(func() {
		FlagsClassStruct = gi.StructNew("GObject", "FlagsClass")
	})
}

type FlagsClass struct {
	native     uintptr
	GTypeClass *TypeClass
	Mask       uint32
	NValues    uint32
	Values     *FlagsValue
}

var FlagsValueStruct *gi.Struct
var FlagsValueStructOnce sync.Once

func FlagsValueStructSet() {
	FlagsValueStructOnce.Do(func() {
		FlagsValueStruct = gi.StructNew("GObject", "FlagsValue")
	})
}

type FlagsValue struct {
	native    uintptr
	Value     uint32
	ValueName string
	ValueNick string
}

var InitiallyUnownedClassStruct *gi.Struct
var InitiallyUnownedClassStructOnce sync.Once

func InitiallyUnownedClassStructSet() {
	InitiallyUnownedClassStructOnce.Do(func() {
		InitiallyUnownedClassStruct = gi.StructNew("GObject", "InitiallyUnownedClass")
	})
}

type InitiallyUnownedClass struct {
	native     uintptr
	GTypeClass *TypeClass
	// UNSUPPORTED : C value 'constructor' : missing Type
	// UNSUPPORTED : C value 'set_property' : missing Type
	// UNSUPPORTED : C value 'get_property' : missing Type
	// UNSUPPORTED : C value 'dispose' : missing Type
	// UNSUPPORTED : C value 'finalize' : missing Type
	// UNSUPPORTED : C value 'dispatch_properties_changed' : missing Type
	// UNSUPPORTED : C value 'notify' : missing Type
	// UNSUPPORTED : C value 'constructed' : missing Type
}

var InterfaceInfoStruct *gi.Struct
var InterfaceInfoStructOnce sync.Once

func InterfaceInfoStructSet() {
	InterfaceInfoStructOnce.Do(func() {
		InterfaceInfoStruct = gi.StructNew("GObject", "InterfaceInfo")
	})
}

type InterfaceInfo struct {
	native uintptr
	// UNSUPPORTED : C value 'interface_init' : no Go type for 'InterfaceInitFunc'
	// UNSUPPORTED : C value 'interface_finalize' : no Go type for 'InterfaceFinalizeFunc'
	// UNSUPPORTED : C value 'interface_data' : no Go type for 'gpointer'
}

var ObjectClassStruct *gi.Struct
var ObjectClassStructOnce sync.Once

func ObjectClassStructSet() {
	ObjectClassStructOnce.Do(func() {
		ObjectClassStruct = gi.StructNew("GObject", "ObjectClass")
	})
}

type ObjectClass struct {
	native     uintptr
	GTypeClass *TypeClass
	// UNSUPPORTED : C value 'constructor' : missing Type
	// UNSUPPORTED : C value 'set_property' : missing Type
	// UNSUPPORTED : C value 'get_property' : missing Type
	// UNSUPPORTED : C value 'dispose' : missing Type
	// UNSUPPORTED : C value 'finalize' : missing Type
	// UNSUPPORTED : C value 'dispatch_properties_changed' : missing Type
	// UNSUPPORTED : C value 'notify' : missing Type
	// UNSUPPORTED : C value 'constructed' : missing Type
}

// UNSUPPORTED : C value 'g_object_class_find_property' : return type 'ParamSpec' not supported

// UNSUPPORTED : C value 'g_object_class_install_properties' : parameter 'pspecs' has no type

// UNSUPPORTED : C value 'g_object_class_install_property' : parameter 'pspec' of type 'ParamSpec' not supported

var listPropertiesObjectClassInvoker *gi.Function

// ListProperties is a representation of the C type g_object_class_list_properties.
func (recv *ObjectClass) ListProperties() uint32 {
	if listPropertiesObjectClassInvoker == nil {
		listPropertiesObjectClassInvoker = gi.StructFunctionInvokerNew("GObject", "ObjectClass", "list_properties")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	listPropertiesObjectClassInvoker.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Uint32()

	return out0
}

var overridePropertyObjectClassInvoker *gi.Function

// OverrideProperty is a representation of the C type g_object_class_override_property.
func (recv *ObjectClass) OverrideProperty(propertyId uint32, name string) {
	if overridePropertyObjectClassInvoker == nil {
		overridePropertyObjectClassInvoker = gi.StructFunctionInvokerNew("GObject", "ObjectClass", "override_property")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(propertyId)
	inArgs[2].SetString(name)

	overridePropertyObjectClassInvoker.Invoke(inArgs[:], nil)

}

var ObjectConstructParamStruct *gi.Struct
var ObjectConstructParamStructOnce sync.Once

func ObjectConstructParamStructSet() {
	ObjectConstructParamStructOnce.Do(func() {
		ObjectConstructParamStruct = gi.StructNew("GObject", "ObjectConstructParam")
	})
}

type ObjectConstructParam struct {
	native uintptr
	// UNSUPPORTED : C value 'pspec' : no Go type for 'ParamSpec'
	Value *Value
}

var ParamSpecClassStruct *gi.Struct
var ParamSpecClassStructOnce sync.Once

func ParamSpecClassStructSet() {
	ParamSpecClassStructOnce.Do(func() {
		ParamSpecClassStruct = gi.StructNew("GObject", "ParamSpecClass")
	})
}

type ParamSpecClass struct {
	native     uintptr
	GTypeClass *TypeClass
	// UNSUPPORTED : C value 'value_type' : no Go type for 'GType'
	// UNSUPPORTED : C value 'finalize' : missing Type
	// UNSUPPORTED : C value 'value_set_default' : missing Type
	// UNSUPPORTED : C value 'value_validate' : missing Type
	// UNSUPPORTED : C value 'values_cmp' : missing Type
}

var ParamSpecPoolStruct *gi.Struct
var ParamSpecPoolStructOnce sync.Once

func ParamSpecPoolStructSet() {
	ParamSpecPoolStructOnce.Do(func() {
		ParamSpecPoolStruct = gi.StructNew("GObject", "ParamSpecPool")
	})
}

type ParamSpecPool struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_param_spec_pool_insert' : parameter 'pspec' of type 'ParamSpec' not supported

// UNSUPPORTED : C value 'g_param_spec_pool_list' : parameter 'owner_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_param_spec_pool_list_owned' : parameter 'owner_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_param_spec_pool_lookup' : parameter 'owner_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_param_spec_pool_remove' : parameter 'pspec' of type 'ParamSpec' not supported

var ParamSpecTypeInfoStruct *gi.Struct
var ParamSpecTypeInfoStructOnce sync.Once

func ParamSpecTypeInfoStructSet() {
	ParamSpecTypeInfoStructOnce.Do(func() {
		ParamSpecTypeInfoStruct = gi.StructNew("GObject", "ParamSpecTypeInfo")
	})
}

type ParamSpecTypeInfo struct {
	native       uintptr
	InstanceSize uint16
	NPreallocs   uint16
	// UNSUPPORTED : C value 'instance_init' : missing Type
	// UNSUPPORTED : C value 'value_type' : no Go type for 'GType'
	// UNSUPPORTED : C value 'finalize' : missing Type
	// UNSUPPORTED : C value 'value_set_default' : missing Type
	// UNSUPPORTED : C value 'value_validate' : missing Type
	// UNSUPPORTED : C value 'values_cmp' : missing Type
}

var ParameterStruct *gi.Struct
var ParameterStructOnce sync.Once

func ParameterStructSet() {
	ParameterStructOnce.Do(func() {
		ParameterStruct = gi.StructNew("GObject", "Parameter")
	})
}

type Parameter struct {
	native uintptr
	Name   string
	Value  *Value
}

var SignalInvocationHintStruct *gi.Struct
var SignalInvocationHintStructOnce sync.Once

func SignalInvocationHintStructSet() {
	SignalInvocationHintStructOnce.Do(func() {
		SignalInvocationHintStruct = gi.StructNew("GObject", "SignalInvocationHint")
	})
}

type SignalInvocationHint struct {
	native   uintptr
	SignalId uint32
	// UNSUPPORTED : C value 'detail' : no Go type for 'GLib.Quark'
	// UNSUPPORTED : C value 'run_type' : no Go type for 'SignalFlags'
}

var SignalQueryStruct *gi.Struct
var SignalQueryStructOnce sync.Once

func SignalQueryStructSet() {
	SignalQueryStructOnce.Do(func() {
		SignalQueryStruct = gi.StructNew("GObject", "SignalQuery")
	})
}

type SignalQuery struct {
	native     uintptr
	SignalId   uint32
	SignalName string
	// UNSUPPORTED : C value 'itype' : no Go type for 'GType'
	// UNSUPPORTED : C value 'signal_flags' : no Go type for 'SignalFlags'
	// UNSUPPORTED : C value 'return_type' : no Go type for 'GType'
	NParams uint32
	// UNSUPPORTED : C value 'param_types' : missing Type
}

var TypeClassStruct *gi.Struct
var TypeClassStructOnce sync.Once

func TypeClassStructSet() {
	TypeClassStructOnce.Do(func() {
		TypeClassStruct = gi.StructNew("GObject", "TypeClass")
	})
}

type TypeClass struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_type_class_add_private' : parameter 'private_size' of type 'gsize' not supported

var getInstancePrivateOffsetTypeClassInvoker *gi.Function

// GetInstancePrivateOffset is a representation of the C type g_type_class_get_instance_private_offset.
func (recv *TypeClass) GetInstancePrivateOffset() int32 {
	if getInstancePrivateOffsetTypeClassInvoker == nil {
		getInstancePrivateOffsetTypeClassInvoker = gi.StructFunctionInvokerNew("GObject", "TypeClass", "get_instance_private_offset")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getInstancePrivateOffsetTypeClassInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_type_class_get_private' : parameter 'private_type' of type 'GType' not supported

var peekParentTypeClassInvoker *gi.Function

// PeekParent is a representation of the C type g_type_class_peek_parent.
func (recv *TypeClass) PeekParent() *TypeClass {
	if peekParentTypeClassInvoker == nil {
		peekParentTypeClassInvoker = gi.StructFunctionInvokerNew("GObject", "TypeClass", "peek_parent")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := peekParentTypeClassInvoker.Invoke(inArgs[:], nil)

	retGo := &TypeClass{native: ret.Pointer()}

	return retGo
}

var unrefTypeClassInvoker *gi.Function

// Unref is a representation of the C type g_type_class_unref.
func (recv *TypeClass) Unref() {
	if unrefTypeClassInvoker == nil {
		unrefTypeClassInvoker = gi.StructFunctionInvokerNew("GObject", "TypeClass", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefTypeClassInvoker.Invoke(inArgs[:], nil)

}

var unrefUncachedTypeClassInvoker *gi.Function

// UnrefUncached is a representation of the C type g_type_class_unref_uncached.
func (recv *TypeClass) UnrefUncached() {
	if unrefUncachedTypeClassInvoker == nil {
		unrefUncachedTypeClassInvoker = gi.StructFunctionInvokerNew("GObject", "TypeClass", "unref_uncached")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefUncachedTypeClassInvoker.Invoke(inArgs[:], nil)

}

var TypeFundamentalInfoStruct *gi.Struct
var TypeFundamentalInfoStructOnce sync.Once

func TypeFundamentalInfoStructSet() {
	TypeFundamentalInfoStructOnce.Do(func() {
		TypeFundamentalInfoStruct = gi.StructNew("GObject", "TypeFundamentalInfo")
	})
}

type TypeFundamentalInfo struct {
	native uintptr
	// UNSUPPORTED : C value 'type_flags' : no Go type for 'TypeFundamentalFlags'
}

var TypeInfoStruct *gi.Struct
var TypeInfoStructOnce sync.Once

func TypeInfoStructSet() {
	TypeInfoStructOnce.Do(func() {
		TypeInfoStruct = gi.StructNew("GObject", "TypeInfo")
	})
}

type TypeInfo struct {
	native    uintptr
	ClassSize uint16
	// UNSUPPORTED : C value 'base_init' : no Go type for 'BaseInitFunc'
	// UNSUPPORTED : C value 'base_finalize' : no Go type for 'BaseFinalizeFunc'
	// UNSUPPORTED : C value 'class_init' : no Go type for 'ClassInitFunc'
	// UNSUPPORTED : C value 'class_finalize' : no Go type for 'ClassFinalizeFunc'
	// UNSUPPORTED : C value 'class_data' : no Go type for 'gpointer'
	InstanceSize uint16
	NPreallocs   uint16
	// UNSUPPORTED : C value 'instance_init' : no Go type for 'InstanceInitFunc'
	ValueTable *TypeValueTable
}

var TypeInstanceStruct *gi.Struct
var TypeInstanceStructOnce sync.Once

func TypeInstanceStructSet() {
	TypeInstanceStructOnce.Do(func() {
		TypeInstanceStruct = gi.StructNew("GObject", "TypeInstance")
	})
}

type TypeInstance struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_type_instance_get_private' : parameter 'private_type' of type 'GType' not supported

var TypeInterfaceStruct *gi.Struct
var TypeInterfaceStructOnce sync.Once

func TypeInterfaceStructSet() {
	TypeInterfaceStructOnce.Do(func() {
		TypeInterfaceStruct = gi.StructNew("GObject", "TypeInterface")
	})
}

type TypeInterface struct {
	native uintptr
}

var peekParentTypeInterfaceInvoker *gi.Function

// PeekParent is a representation of the C type g_type_interface_peek_parent.
func (recv *TypeInterface) PeekParent() *TypeInterface {
	if peekParentTypeInterfaceInvoker == nil {
		peekParentTypeInterfaceInvoker = gi.StructFunctionInvokerNew("GObject", "TypeInterface", "peek_parent")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := peekParentTypeInterfaceInvoker.Invoke(inArgs[:], nil)

	retGo := &TypeInterface{native: ret.Pointer()}

	return retGo
}

var TypeModuleClassStruct *gi.Struct
var TypeModuleClassStructOnce sync.Once

func TypeModuleClassStructSet() {
	TypeModuleClassStructOnce.Do(func() {
		TypeModuleClassStruct = gi.StructNew("GObject", "TypeModuleClass")
	})
}

type TypeModuleClass struct {
	native      uintptr
	ParentClass *ObjectClass
	// UNSUPPORTED : C value 'load' : missing Type
	// UNSUPPORTED : C value 'unload' : missing Type
	// UNSUPPORTED : C value 'reserved1' : missing Type
	// UNSUPPORTED : C value 'reserved2' : missing Type
	// UNSUPPORTED : C value 'reserved3' : missing Type
	// UNSUPPORTED : C value 'reserved4' : missing Type
}

var TypePluginClassStruct *gi.Struct
var TypePluginClassStructOnce sync.Once

func TypePluginClassStructSet() {
	TypePluginClassStructOnce.Do(func() {
		TypePluginClassStruct = gi.StructNew("GObject", "TypePluginClass")
	})
}

type TypePluginClass struct {
	native uintptr
	// UNSUPPORTED : C value 'use_plugin' : no Go type for 'TypePluginUse'
	// UNSUPPORTED : C value 'unuse_plugin' : no Go type for 'TypePluginUnuse'
	// UNSUPPORTED : C value 'complete_type_info' : no Go type for 'TypePluginCompleteTypeInfo'
	// UNSUPPORTED : C value 'complete_interface_info' : no Go type for 'TypePluginCompleteInterfaceInfo'
}

var TypeQueryStruct *gi.Struct
var TypeQueryStructOnce sync.Once

func TypeQueryStructSet() {
	TypeQueryStructOnce.Do(func() {
		TypeQueryStruct = gi.StructNew("GObject", "TypeQuery")
	})
}

type TypeQuery struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'GType'
	TypeName     string
	ClassSize    uint32
	InstanceSize uint32
}

var TypeValueTableStruct *gi.Struct
var TypeValueTableStructOnce sync.Once

func TypeValueTableStructSet() {
	TypeValueTableStructOnce.Do(func() {
		TypeValueTableStruct = gi.StructNew("GObject", "TypeValueTable")
	})
}

type TypeValueTable struct {
	native uintptr
	// UNSUPPORTED : C value 'value_init' : missing Type
	// UNSUPPORTED : C value 'value_free' : missing Type
	// UNSUPPORTED : C value 'value_copy' : missing Type
	// UNSUPPORTED : C value 'value_peek_pointer' : missing Type
	CollectFormat string
	// UNSUPPORTED : C value 'collect_value' : missing Type
	LcopyFormat string
	// UNSUPPORTED : C value 'lcopy_value' : missing Type
}

var ValueStruct *gi.Struct
var ValueStructOnce sync.Once

func ValueStructSet() {
	ValueStructOnce.Do(func() {
		ValueStruct = gi.StructNew("GObject", "Value")
	})
}

type Value struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : missing Type
}

// UNSUPPORTED : C value 'g_value_copy' : parameter 'dest_value' of type 'Value' not supported

// UNSUPPORTED : C value 'g_value_dup_boxed' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_value_dup_object' : return type 'Object' not supported

// UNSUPPORTED : C value 'g_value_dup_param' : return type 'ParamSpec' not supported

var dupStringValueInvoker *gi.Function

// DupString is a representation of the C type g_value_dup_string.
func (recv *Value) DupString() string {
	if dupStringValueInvoker == nil {
		dupStringValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "dup_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dupStringValueInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_value_dup_variant' : return type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_value_fits_pointer' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_value_get_boolean' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_value_get_boxed' : return type 'gpointer' not supported

var getCharValueInvoker *gi.Function

// GetChar is a representation of the C type g_value_get_char.
func (recv *Value) GetChar() int8 {
	if getCharValueInvoker == nil {
		getCharValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "get_char")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getCharValueInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int8()

	return retGo
}

// UNSUPPORTED : C value 'g_value_get_double' : return type 'gdouble' not supported

var getEnumValueInvoker *gi.Function

// GetEnum is a representation of the C type g_value_get_enum.
func (recv *Value) GetEnum() int32 {
	if getEnumValueInvoker == nil {
		getEnumValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "get_enum")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getEnumValueInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getFlagsValueInvoker *gi.Function

// GetFlags is a representation of the C type g_value_get_flags.
func (recv *Value) GetFlags() uint32 {
	if getFlagsValueInvoker == nil {
		getFlagsValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "get_flags")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getFlagsValueInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_value_get_float' : return type 'gfloat' not supported

// UNSUPPORTED : C value 'g_value_get_gtype' : return type 'GType' not supported

var getIntValueInvoker *gi.Function

// GetInt is a representation of the C type g_value_get_int.
func (recv *Value) GetInt() int32 {
	if getIntValueInvoker == nil {
		getIntValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "get_int")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getIntValueInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getInt64ValueInvoker *gi.Function

// GetInt64 is a representation of the C type g_value_get_int64.
func (recv *Value) GetInt64() int64 {
	if getInt64ValueInvoker == nil {
		getInt64ValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "get_int64")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getInt64ValueInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var getLongValueInvoker *gi.Function

// GetLong is a representation of the C type g_value_get_long.
func (recv *Value) GetLong() int64 {
	if getLongValueInvoker == nil {
		getLongValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "get_long")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getLongValueInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'g_value_get_object' : return type 'Object' not supported

// UNSUPPORTED : C value 'g_value_get_param' : return type 'ParamSpec' not supported

// UNSUPPORTED : C value 'g_value_get_pointer' : return type 'gpointer' not supported

var getScharValueInvoker *gi.Function

// GetSchar is a representation of the C type g_value_get_schar.
func (recv *Value) GetSchar() int8 {
	if getScharValueInvoker == nil {
		getScharValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "get_schar")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getScharValueInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int8()

	return retGo
}

var getStringValueInvoker *gi.Function

// GetString is a representation of the C type g_value_get_string.
func (recv *Value) GetString() string {
	if getStringValueInvoker == nil {
		getStringValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "get_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getStringValueInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getUcharValueInvoker *gi.Function

// GetUchar is a representation of the C type g_value_get_uchar.
func (recv *Value) GetUchar() uint8 {
	if getUcharValueInvoker == nil {
		getUcharValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "get_uchar")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getUcharValueInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint8()

	return retGo
}

var getUintValueInvoker *gi.Function

// GetUint is a representation of the C type g_value_get_uint.
func (recv *Value) GetUint() uint32 {
	if getUintValueInvoker == nil {
		getUintValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "get_uint")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getUintValueInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var getUint64ValueInvoker *gi.Function

// GetUint64 is a representation of the C type g_value_get_uint64.
func (recv *Value) GetUint64() uint64 {
	if getUint64ValueInvoker == nil {
		getUint64ValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "get_uint64")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getUint64ValueInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint64()

	return retGo
}

var getUlongValueInvoker *gi.Function

// GetUlong is a representation of the C type g_value_get_ulong.
func (recv *Value) GetUlong() uint64 {
	if getUlongValueInvoker == nil {
		getUlongValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "get_ulong")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getUlongValueInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint64()

	return retGo
}

// UNSUPPORTED : C value 'g_value_get_variant' : return type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_value_init' : parameter 'g_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_value_init_from_instance' : parameter 'instance' of type 'TypeInstance' not supported

// UNSUPPORTED : C value 'g_value_peek_pointer' : return type 'gpointer' not supported

var resetValueInvoker *gi.Function

// Reset is a representation of the C type g_value_reset.
func (recv *Value) Reset() *Value {
	if resetValueInvoker == nil {
		resetValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "reset")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := resetValueInvoker.Invoke(inArgs[:], nil)

	retGo := &Value{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_value_set_boolean' : parameter 'v_boolean' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_value_set_boxed' : parameter 'v_boxed' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_value_set_boxed_take_ownership' : parameter 'v_boxed' of type 'gpointer' not supported

var setCharValueInvoker *gi.Function

// SetChar is a representation of the C type g_value_set_char.
func (recv *Value) SetChar(vChar int8) {
	if setCharValueInvoker == nil {
		setCharValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "set_char")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt8(vChar)

	setCharValueInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_value_set_double' : parameter 'v_double' of type 'gdouble' not supported

var setEnumValueInvoker *gi.Function

// SetEnum is a representation of the C type g_value_set_enum.
func (recv *Value) SetEnum(vEnum int32) {
	if setEnumValueInvoker == nil {
		setEnumValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "set_enum")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(vEnum)

	setEnumValueInvoker.Invoke(inArgs[:], nil)

}

var setFlagsValueInvoker *gi.Function

// SetFlags is a representation of the C type g_value_set_flags.
func (recv *Value) SetFlags(vFlags uint32) {
	if setFlagsValueInvoker == nil {
		setFlagsValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "set_flags")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(vFlags)

	setFlagsValueInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_value_set_float' : parameter 'v_float' of type 'gfloat' not supported

// UNSUPPORTED : C value 'g_value_set_gtype' : parameter 'v_gtype' of type 'GType' not supported

// UNSUPPORTED : C value 'g_value_set_instance' : parameter 'instance' of type 'gpointer' not supported

var setIntValueInvoker *gi.Function

// SetInt is a representation of the C type g_value_set_int.
func (recv *Value) SetInt(vInt int32) {
	if setIntValueInvoker == nil {
		setIntValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "set_int")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(vInt)

	setIntValueInvoker.Invoke(inArgs[:], nil)

}

var setInt64ValueInvoker *gi.Function

// SetInt64 is a representation of the C type g_value_set_int64.
func (recv *Value) SetInt64(vInt64 int64) {
	if setInt64ValueInvoker == nil {
		setInt64ValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "set_int64")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(vInt64)

	setInt64ValueInvoker.Invoke(inArgs[:], nil)

}

var setLongValueInvoker *gi.Function

// SetLong is a representation of the C type g_value_set_long.
func (recv *Value) SetLong(vLong int64) {
	if setLongValueInvoker == nil {
		setLongValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "set_long")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(vLong)

	setLongValueInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_value_set_object' : parameter 'v_object' of type 'Object' not supported

// UNSUPPORTED : C value 'g_value_set_object_take_ownership' : parameter 'v_object' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_value_set_param' : parameter 'param' of type 'ParamSpec' not supported

// UNSUPPORTED : C value 'g_value_set_param_take_ownership' : parameter 'param' of type 'ParamSpec' not supported

// UNSUPPORTED : C value 'g_value_set_pointer' : parameter 'v_pointer' of type 'gpointer' not supported

var setScharValueInvoker *gi.Function

// SetSchar is a representation of the C type g_value_set_schar.
func (recv *Value) SetSchar(vChar int8) {
	if setScharValueInvoker == nil {
		setScharValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "set_schar")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt8(vChar)

	setScharValueInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_value_set_static_boxed' : parameter 'v_boxed' of type 'gpointer' not supported

var setStaticStringValueInvoker *gi.Function

// SetStaticString is a representation of the C type g_value_set_static_string.
func (recv *Value) SetStaticString(vString string) {
	if setStaticStringValueInvoker == nil {
		setStaticStringValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "set_static_string")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(vString)

	setStaticStringValueInvoker.Invoke(inArgs[:], nil)

}

var setStringValueInvoker *gi.Function

// SetString is a representation of the C type g_value_set_string.
func (recv *Value) SetString(vString string) {
	if setStringValueInvoker == nil {
		setStringValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "set_string")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(vString)

	setStringValueInvoker.Invoke(inArgs[:], nil)

}

var setStringTakeOwnershipValueInvoker *gi.Function

// SetStringTakeOwnership is a representation of the C type g_value_set_string_take_ownership.
func (recv *Value) SetStringTakeOwnership(vString string) {
	if setStringTakeOwnershipValueInvoker == nil {
		setStringTakeOwnershipValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "set_string_take_ownership")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(vString)

	setStringTakeOwnershipValueInvoker.Invoke(inArgs[:], nil)

}

var setUcharValueInvoker *gi.Function

// SetUchar is a representation of the C type g_value_set_uchar.
func (recv *Value) SetUchar(vUchar uint8) {
	if setUcharValueInvoker == nil {
		setUcharValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "set_uchar")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint8(vUchar)

	setUcharValueInvoker.Invoke(inArgs[:], nil)

}

var setUintValueInvoker *gi.Function

// SetUint is a representation of the C type g_value_set_uint.
func (recv *Value) SetUint(vUint uint32) {
	if setUintValueInvoker == nil {
		setUintValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "set_uint")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(vUint)

	setUintValueInvoker.Invoke(inArgs[:], nil)

}

var setUint64ValueInvoker *gi.Function

// SetUint64 is a representation of the C type g_value_set_uint64.
func (recv *Value) SetUint64(vUint64 uint64) {
	if setUint64ValueInvoker == nil {
		setUint64ValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "set_uint64")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint64(vUint64)

	setUint64ValueInvoker.Invoke(inArgs[:], nil)

}

var setUlongValueInvoker *gi.Function

// SetUlong is a representation of the C type g_value_set_ulong.
func (recv *Value) SetUlong(vUlong uint64) {
	if setUlongValueInvoker == nil {
		setUlongValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "set_ulong")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint64(vUlong)

	setUlongValueInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_value_set_variant' : parameter 'variant' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_value_take_boxed' : parameter 'v_boxed' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_value_take_object' : parameter 'v_object' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_value_take_param' : parameter 'param' of type 'ParamSpec' not supported

var takeStringValueInvoker *gi.Function

// TakeString is a representation of the C type g_value_take_string.
func (recv *Value) TakeString(vString string) {
	if takeStringValueInvoker == nil {
		takeStringValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "take_string")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(vString)

	takeStringValueInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_value_take_variant' : parameter 'variant' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_value_transform' : parameter 'dest_value' of type 'Value' not supported

var unsetValueInvoker *gi.Function

// Unset is a representation of the C type g_value_unset.
func (recv *Value) Unset() {
	if unsetValueInvoker == nil {
		unsetValueInvoker = gi.StructFunctionInvokerNew("GObject", "Value", "unset")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unsetValueInvoker.Invoke(inArgs[:], nil)

}

var ValueArrayStruct *gi.Struct
var ValueArrayStructOnce sync.Once

func ValueArrayStructSet() {
	ValueArrayStructOnce.Do(func() {
		ValueArrayStruct = gi.StructNew("GObject", "ValueArray")
	})
}

type ValueArray struct {
	native  uintptr
	NValues uint32
	Values  *Value
}

var newValueArrayInvoker *gi.Function

// ValueArrayNew is a representation of the C type g_value_array_new.
func ValueArrayNew(nPrealloced uint32) *ValueArray {
	if newValueArrayInvoker == nil {
		newValueArrayInvoker = gi.StructFunctionInvokerNew("GObject", "ValueArray", "new")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(nPrealloced)

	ret := newValueArrayInvoker.Invoke(inArgs[:], nil)

	retGo := &ValueArray{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_value_array_append' : parameter 'value' of type 'Value' not supported

var copyValueArrayInvoker *gi.Function

// Copy is a representation of the C type g_value_array_copy.
func (recv *ValueArray) Copy() *ValueArray {
	if copyValueArrayInvoker == nil {
		copyValueArrayInvoker = gi.StructFunctionInvokerNew("GObject", "ValueArray", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyValueArrayInvoker.Invoke(inArgs[:], nil)

	retGo := &ValueArray{native: ret.Pointer()}

	return retGo
}

var freeValueArrayInvoker *gi.Function

// Free is a representation of the C type g_value_array_free.
func (recv *ValueArray) Free() {
	if freeValueArrayInvoker == nil {
		freeValueArrayInvoker = gi.StructFunctionInvokerNew("GObject", "ValueArray", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeValueArrayInvoker.Invoke(inArgs[:], nil)

}

var getNthValueArrayInvoker *gi.Function

// GetNth is a representation of the C type g_value_array_get_nth.
func (recv *ValueArray) GetNth(index uint32) *Value {
	if getNthValueArrayInvoker == nil {
		getNthValueArrayInvoker = gi.StructFunctionInvokerNew("GObject", "ValueArray", "get_nth")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(index)

	ret := getNthValueArrayInvoker.Invoke(inArgs[:], nil)

	retGo := &Value{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_value_array_insert' : parameter 'value' of type 'Value' not supported

// UNSUPPORTED : C value 'g_value_array_prepend' : parameter 'value' of type 'Value' not supported

var removeValueArrayInvoker *gi.Function

// Remove is a representation of the C type g_value_array_remove.
func (recv *ValueArray) Remove(index uint32) *ValueArray {
	if removeValueArrayInvoker == nil {
		removeValueArrayInvoker = gi.StructFunctionInvokerNew("GObject", "ValueArray", "remove")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(index)

	ret := removeValueArrayInvoker.Invoke(inArgs[:], nil)

	retGo := &ValueArray{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_value_array_sort' : parameter 'compare_func' of type 'GLib.CompareFunc' not supported

// UNSUPPORTED : C value 'g_value_array_sort_with_data' : parameter 'compare_func' of type 'GLib.CompareDataFunc' not supported

var WeakRefStruct *gi.Struct
var WeakRefStructOnce sync.Once

func WeakRefStructSet() {
	WeakRefStructOnce.Do(func() {
		WeakRefStruct = gi.StructNew("GObject", "WeakRef")
	})
}

type WeakRef struct {
	native uintptr
}

var clearWeakRefInvoker *gi.Function

// Clear is a representation of the C type g_weak_ref_clear.
func (recv *WeakRef) Clear() {
	if clearWeakRefInvoker == nil {
		clearWeakRefInvoker = gi.StructFunctionInvokerNew("GObject", "WeakRef", "clear")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	clearWeakRefInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_weak_ref_get' : return type 'Object' not supported

// UNSUPPORTED : C value 'g_weak_ref_init' : parameter 'object' of type 'Object' not supported

// UNSUPPORTED : C value 'g_weak_ref_set' : parameter 'object' of type 'Object' not supported
