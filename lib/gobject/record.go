// Code generated - DO NOT EDIT.

package gobject

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var cClosureStruct *gi.Struct
var cClosureStruct_Once sync.Once

func cClosureStruct_Set() error {
	var err error
	cClosureStruct_Once.Do(func() {
		cClosureStruct, err = gi.StructNew("GObject", "CClosure")
	})
	return err
}

type CClosure struct {
	native  uintptr
	Closure *Closure
	// UNSUPPORTED : C value 'callback' : no Go type for 'gpointer'
}

var closureStruct *gi.Struct
var closureStruct_Once sync.Once

func closureStruct_Set() error {
	var err error
	closureStruct_Once.Do(func() {
		closureStruct, err = gi.StructNew("GObject", "Closure")
	})
	return err
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

var closureInvalidateFunction *gi.Function
var closureInvalidateFunction_Once sync.Once

func closureInvalidateFunction_Set() {
	closureInvalidateFunction_Once.Do(func() {
		closureStruct_Set()
		closureInvalidateFunction = closureStruct.InvokerNew("invalidate")
	})
}

// Invalidate is a representation of the C type g_closure_invalidate.
func (recv *Closure) Invalidate() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	closureInvalidateFunction_Set()

	closureInvalidateFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_closure_invoke' : parameter 'return_value' of type 'Value' not supported

var closureRefFunction *gi.Function
var closureRefFunction_Once sync.Once

func closureRefFunction_Set() {
	closureRefFunction_Once.Do(func() {
		closureStruct_Set()
		closureRefFunction = closureStruct.InvokerNew("ref")
	})
}

// Ref is a representation of the C type g_closure_ref.
func (recv *Closure) Ref() *Closure {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	closureRefFunction_Set()

	ret = closureRefFunction.Invoke(inArgs[:], nil)

	retGo := &Closure{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_closure_remove_finalize_notifier' : parameter 'notify_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_closure_remove_invalidate_notifier' : parameter 'notify_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_closure_set_marshal' : parameter 'marshal' of type 'ClosureMarshal' not supported

// UNSUPPORTED : C value 'g_closure_set_meta_marshal' : parameter 'marshal_data' of type 'gpointer' not supported

var closureSinkFunction *gi.Function
var closureSinkFunction_Once sync.Once

func closureSinkFunction_Set() {
	closureSinkFunction_Once.Do(func() {
		closureStruct_Set()
		closureSinkFunction = closureStruct.InvokerNew("sink")
	})
}

// Sink is a representation of the C type g_closure_sink.
func (recv *Closure) Sink() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	closureSinkFunction_Set()

	closureSinkFunction.Invoke(inArgs[:], nil)

}

var closureUnrefFunction *gi.Function
var closureUnrefFunction_Once sync.Once

func closureUnrefFunction_Set() {
	closureUnrefFunction_Once.Do(func() {
		closureStruct_Set()
		closureUnrefFunction = closureStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type g_closure_unref.
func (recv *Closure) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	closureUnrefFunction_Set()

	closureUnrefFunction.Invoke(inArgs[:], nil)

}

var closureNotifyDataStruct *gi.Struct
var closureNotifyDataStruct_Once sync.Once

func closureNotifyDataStruct_Set() error {
	var err error
	closureNotifyDataStruct_Once.Do(func() {
		closureNotifyDataStruct, err = gi.StructNew("GObject", "ClosureNotifyData")
	})
	return err
}

type ClosureNotifyData struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'
	// UNSUPPORTED : C value 'notify' : no Go type for 'ClosureNotify'
}

var enumClassStruct *gi.Struct
var enumClassStruct_Once sync.Once

func enumClassStruct_Set() error {
	var err error
	enumClassStruct_Once.Do(func() {
		enumClassStruct, err = gi.StructNew("GObject", "EnumClass")
	})
	return err
}

type EnumClass struct {
	native     uintptr
	GTypeClass *TypeClass
	Minimum    int32
	Maximum    int32
	NValues    uint32
	Values     *EnumValue
}

var enumValueStruct *gi.Struct
var enumValueStruct_Once sync.Once

func enumValueStruct_Set() error {
	var err error
	enumValueStruct_Once.Do(func() {
		enumValueStruct, err = gi.StructNew("GObject", "EnumValue")
	})
	return err
}

type EnumValue struct {
	native    uintptr
	Value     int32
	ValueName string
	ValueNick string
}

var flagsClassStruct *gi.Struct
var flagsClassStruct_Once sync.Once

func flagsClassStruct_Set() error {
	var err error
	flagsClassStruct_Once.Do(func() {
		flagsClassStruct, err = gi.StructNew("GObject", "FlagsClass")
	})
	return err
}

type FlagsClass struct {
	native     uintptr
	GTypeClass *TypeClass
	Mask       uint32
	NValues    uint32
	Values     *FlagsValue
}

var flagsValueStruct *gi.Struct
var flagsValueStruct_Once sync.Once

func flagsValueStruct_Set() error {
	var err error
	flagsValueStruct_Once.Do(func() {
		flagsValueStruct, err = gi.StructNew("GObject", "FlagsValue")
	})
	return err
}

type FlagsValue struct {
	native    uintptr
	Value     uint32
	ValueName string
	ValueNick string
}

var initiallyUnownedClassStruct *gi.Struct
var initiallyUnownedClassStruct_Once sync.Once

func initiallyUnownedClassStruct_Set() error {
	var err error
	initiallyUnownedClassStruct_Once.Do(func() {
		initiallyUnownedClassStruct, err = gi.StructNew("GObject", "InitiallyUnownedClass")
	})
	return err
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

var interfaceInfoStruct *gi.Struct
var interfaceInfoStruct_Once sync.Once

func interfaceInfoStruct_Set() error {
	var err error
	interfaceInfoStruct_Once.Do(func() {
		interfaceInfoStruct, err = gi.StructNew("GObject", "InterfaceInfo")
	})
	return err
}

type InterfaceInfo struct {
	native uintptr
	// UNSUPPORTED : C value 'interface_init' : no Go type for 'InterfaceInitFunc'
	// UNSUPPORTED : C value 'interface_finalize' : no Go type for 'InterfaceFinalizeFunc'
	// UNSUPPORTED : C value 'interface_data' : no Go type for 'gpointer'
}

var objectClassStruct *gi.Struct
var objectClassStruct_Once sync.Once

func objectClassStruct_Set() error {
	var err error
	objectClassStruct_Once.Do(func() {
		objectClassStruct, err = gi.StructNew("GObject", "ObjectClass")
	})
	return err
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

var objectClassListPropertiesFunction *gi.Function
var objectClassListPropertiesFunction_Once sync.Once

func objectClassListPropertiesFunction_Set() {
	objectClassListPropertiesFunction_Once.Do(func() {
		objectClassStruct_Set()
		objectClassListPropertiesFunction = objectClassStruct.InvokerNew("list_properties")
	})
}

// ListProperties is a representation of the C type g_object_class_list_properties.
func (recv *ObjectClass) ListProperties() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	objectClassListPropertiesFunction_Set()

	objectClassListPropertiesFunction.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Uint32()

	return out0
}

var objectClassOverridePropertyFunction *gi.Function
var objectClassOverridePropertyFunction_Once sync.Once

func objectClassOverridePropertyFunction_Set() {
	objectClassOverridePropertyFunction_Once.Do(func() {
		objectClassStruct_Set()
		objectClassOverridePropertyFunction = objectClassStruct.InvokerNew("override_property")
	})
}

// OverrideProperty is a representation of the C type g_object_class_override_property.
func (recv *ObjectClass) OverrideProperty(propertyId uint32, name string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(propertyId)
	inArgs[2].SetString(name)

	objectClassOverridePropertyFunction_Set()

	objectClassOverridePropertyFunction.Invoke(inArgs[:], nil)

}

var objectConstructParamStruct *gi.Struct
var objectConstructParamStruct_Once sync.Once

func objectConstructParamStruct_Set() error {
	var err error
	objectConstructParamStruct_Once.Do(func() {
		objectConstructParamStruct, err = gi.StructNew("GObject", "ObjectConstructParam")
	})
	return err
}

type ObjectConstructParam struct {
	native uintptr
	// UNSUPPORTED : C value 'pspec' : no Go type for 'ParamSpec'
	Value *Value
}

var paramSpecClassStruct *gi.Struct
var paramSpecClassStruct_Once sync.Once

func paramSpecClassStruct_Set() error {
	var err error
	paramSpecClassStruct_Once.Do(func() {
		paramSpecClassStruct, err = gi.StructNew("GObject", "ParamSpecClass")
	})
	return err
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

var paramSpecPoolStruct *gi.Struct
var paramSpecPoolStruct_Once sync.Once

func paramSpecPoolStruct_Set() error {
	var err error
	paramSpecPoolStruct_Once.Do(func() {
		paramSpecPoolStruct, err = gi.StructNew("GObject", "ParamSpecPool")
	})
	return err
}

type ParamSpecPool struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_param_spec_pool_insert' : parameter 'pspec' of type 'ParamSpec' not supported

// UNSUPPORTED : C value 'g_param_spec_pool_list' : parameter 'owner_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_param_spec_pool_list_owned' : parameter 'owner_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_param_spec_pool_lookup' : parameter 'owner_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_param_spec_pool_remove' : parameter 'pspec' of type 'ParamSpec' not supported

var paramSpecTypeInfoStruct *gi.Struct
var paramSpecTypeInfoStruct_Once sync.Once

func paramSpecTypeInfoStruct_Set() error {
	var err error
	paramSpecTypeInfoStruct_Once.Do(func() {
		paramSpecTypeInfoStruct, err = gi.StructNew("GObject", "ParamSpecTypeInfo")
	})
	return err
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

var parameterStruct *gi.Struct
var parameterStruct_Once sync.Once

func parameterStruct_Set() error {
	var err error
	parameterStruct_Once.Do(func() {
		parameterStruct, err = gi.StructNew("GObject", "Parameter")
	})
	return err
}

type Parameter struct {
	native uintptr
	Name   string
	Value  *Value
}

var signalInvocationHintStruct *gi.Struct
var signalInvocationHintStruct_Once sync.Once

func signalInvocationHintStruct_Set() error {
	var err error
	signalInvocationHintStruct_Once.Do(func() {
		signalInvocationHintStruct, err = gi.StructNew("GObject", "SignalInvocationHint")
	})
	return err
}

type SignalInvocationHint struct {
	native   uintptr
	SignalId uint32
	// UNSUPPORTED : C value 'detail' : no Go type for 'GLib.Quark'
	// UNSUPPORTED : C value 'run_type' : no Go type for 'SignalFlags'
}

var signalQueryStruct *gi.Struct
var signalQueryStruct_Once sync.Once

func signalQueryStruct_Set() error {
	var err error
	signalQueryStruct_Once.Do(func() {
		signalQueryStruct, err = gi.StructNew("GObject", "SignalQuery")
	})
	return err
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

var typeClassStruct *gi.Struct
var typeClassStruct_Once sync.Once

func typeClassStruct_Set() error {
	var err error
	typeClassStruct_Once.Do(func() {
		typeClassStruct, err = gi.StructNew("GObject", "TypeClass")
	})
	return err
}

type TypeClass struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_type_class_add_private' : parameter 'private_size' of type 'gsize' not supported

var typeClassGetInstancePrivateOffsetFunction *gi.Function
var typeClassGetInstancePrivateOffsetFunction_Once sync.Once

func typeClassGetInstancePrivateOffsetFunction_Set() {
	typeClassGetInstancePrivateOffsetFunction_Once.Do(func() {
		typeClassStruct_Set()
		typeClassGetInstancePrivateOffsetFunction = typeClassStruct.InvokerNew("get_instance_private_offset")
	})
}

// GetInstancePrivateOffset is a representation of the C type g_type_class_get_instance_private_offset.
func (recv *TypeClass) GetInstancePrivateOffset() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	typeClassGetInstancePrivateOffsetFunction_Set()

	ret = typeClassGetInstancePrivateOffsetFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_type_class_get_private' : parameter 'private_type' of type 'GType' not supported

var typeClassPeekParentFunction *gi.Function
var typeClassPeekParentFunction_Once sync.Once

func typeClassPeekParentFunction_Set() {
	typeClassPeekParentFunction_Once.Do(func() {
		typeClassStruct_Set()
		typeClassPeekParentFunction = typeClassStruct.InvokerNew("peek_parent")
	})
}

// PeekParent is a representation of the C type g_type_class_peek_parent.
func (recv *TypeClass) PeekParent() *TypeClass {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	typeClassPeekParentFunction_Set()

	ret = typeClassPeekParentFunction.Invoke(inArgs[:], nil)

	retGo := &TypeClass{native: ret.Pointer()}

	return retGo
}

var typeClassUnrefFunction *gi.Function
var typeClassUnrefFunction_Once sync.Once

func typeClassUnrefFunction_Set() {
	typeClassUnrefFunction_Once.Do(func() {
		typeClassStruct_Set()
		typeClassUnrefFunction = typeClassStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type g_type_class_unref.
func (recv *TypeClass) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	typeClassUnrefFunction_Set()

	typeClassUnrefFunction.Invoke(inArgs[:], nil)

}

var typeClassUnrefUncachedFunction *gi.Function
var typeClassUnrefUncachedFunction_Once sync.Once

func typeClassUnrefUncachedFunction_Set() {
	typeClassUnrefUncachedFunction_Once.Do(func() {
		typeClassStruct_Set()
		typeClassUnrefUncachedFunction = typeClassStruct.InvokerNew("unref_uncached")
	})
}

// UnrefUncached is a representation of the C type g_type_class_unref_uncached.
func (recv *TypeClass) UnrefUncached() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	typeClassUnrefUncachedFunction_Set()

	typeClassUnrefUncachedFunction.Invoke(inArgs[:], nil)

}

var typeFundamentalInfoStruct *gi.Struct
var typeFundamentalInfoStruct_Once sync.Once

func typeFundamentalInfoStruct_Set() error {
	var err error
	typeFundamentalInfoStruct_Once.Do(func() {
		typeFundamentalInfoStruct, err = gi.StructNew("GObject", "TypeFundamentalInfo")
	})
	return err
}

type TypeFundamentalInfo struct {
	native uintptr
	// UNSUPPORTED : C value 'type_flags' : no Go type for 'TypeFundamentalFlags'
}

var typeInfoStruct *gi.Struct
var typeInfoStruct_Once sync.Once

func typeInfoStruct_Set() error {
	var err error
	typeInfoStruct_Once.Do(func() {
		typeInfoStruct, err = gi.StructNew("GObject", "TypeInfo")
	})
	return err
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

var typeInstanceStruct *gi.Struct
var typeInstanceStruct_Once sync.Once

func typeInstanceStruct_Set() error {
	var err error
	typeInstanceStruct_Once.Do(func() {
		typeInstanceStruct, err = gi.StructNew("GObject", "TypeInstance")
	})
	return err
}

type TypeInstance struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_type_instance_get_private' : parameter 'private_type' of type 'GType' not supported

var typeInterfaceStruct *gi.Struct
var typeInterfaceStruct_Once sync.Once

func typeInterfaceStruct_Set() error {
	var err error
	typeInterfaceStruct_Once.Do(func() {
		typeInterfaceStruct, err = gi.StructNew("GObject", "TypeInterface")
	})
	return err
}

type TypeInterface struct {
	native uintptr
}

var typeInterfacePeekParentFunction *gi.Function
var typeInterfacePeekParentFunction_Once sync.Once

func typeInterfacePeekParentFunction_Set() {
	typeInterfacePeekParentFunction_Once.Do(func() {
		typeInterfaceStruct_Set()
		typeInterfacePeekParentFunction = typeInterfaceStruct.InvokerNew("peek_parent")
	})
}

// PeekParent is a representation of the C type g_type_interface_peek_parent.
func (recv *TypeInterface) PeekParent() *TypeInterface {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	typeInterfacePeekParentFunction_Set()

	ret = typeInterfacePeekParentFunction.Invoke(inArgs[:], nil)

	retGo := &TypeInterface{native: ret.Pointer()}

	return retGo
}

var typeModuleClassStruct *gi.Struct
var typeModuleClassStruct_Once sync.Once

func typeModuleClassStruct_Set() error {
	var err error
	typeModuleClassStruct_Once.Do(func() {
		typeModuleClassStruct, err = gi.StructNew("GObject", "TypeModuleClass")
	})
	return err
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

var typePluginClassStruct *gi.Struct
var typePluginClassStruct_Once sync.Once

func typePluginClassStruct_Set() error {
	var err error
	typePluginClassStruct_Once.Do(func() {
		typePluginClassStruct, err = gi.StructNew("GObject", "TypePluginClass")
	})
	return err
}

type TypePluginClass struct {
	native uintptr
	// UNSUPPORTED : C value 'use_plugin' : no Go type for 'TypePluginUse'
	// UNSUPPORTED : C value 'unuse_plugin' : no Go type for 'TypePluginUnuse'
	// UNSUPPORTED : C value 'complete_type_info' : no Go type for 'TypePluginCompleteTypeInfo'
	// UNSUPPORTED : C value 'complete_interface_info' : no Go type for 'TypePluginCompleteInterfaceInfo'
}

var typeQueryStruct *gi.Struct
var typeQueryStruct_Once sync.Once

func typeQueryStruct_Set() error {
	var err error
	typeQueryStruct_Once.Do(func() {
		typeQueryStruct, err = gi.StructNew("GObject", "TypeQuery")
	})
	return err
}

type TypeQuery struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'GType'
	TypeName     string
	ClassSize    uint32
	InstanceSize uint32
}

var typeValueTableStruct *gi.Struct
var typeValueTableStruct_Once sync.Once

func typeValueTableStruct_Set() error {
	var err error
	typeValueTableStruct_Once.Do(func() {
		typeValueTableStruct, err = gi.StructNew("GObject", "TypeValueTable")
	})
	return err
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

var valueStruct *gi.Struct
var valueStruct_Once sync.Once

func valueStruct_Set() error {
	var err error
	valueStruct_Once.Do(func() {
		valueStruct, err = gi.StructNew("GObject", "Value")
	})
	return err
}

type Value struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : missing Type
}

// UNSUPPORTED : C value 'g_value_copy' : parameter 'dest_value' of type 'Value' not supported

// UNSUPPORTED : C value 'g_value_dup_boxed' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_value_dup_object' : return type 'Object' not supported

// UNSUPPORTED : C value 'g_value_dup_param' : return type 'ParamSpec' not supported

var valueDupStringFunction *gi.Function
var valueDupStringFunction_Once sync.Once

func valueDupStringFunction_Set() {
	valueDupStringFunction_Once.Do(func() {
		valueStruct_Set()
		valueDupStringFunction = valueStruct.InvokerNew("dup_string")
	})
}

// DupString is a representation of the C type g_value_dup_string.
func (recv *Value) DupString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	valueDupStringFunction_Set()

	ret = valueDupStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_value_dup_variant' : return type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_value_fits_pointer' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_value_get_boolean' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_value_get_boxed' : return type 'gpointer' not supported

var valueGetCharFunction *gi.Function
var valueGetCharFunction_Once sync.Once

func valueGetCharFunction_Set() {
	valueGetCharFunction_Once.Do(func() {
		valueStruct_Set()
		valueGetCharFunction = valueStruct.InvokerNew("get_char")
	})
}

// GetChar is a representation of the C type g_value_get_char.
func (recv *Value) GetChar() int8 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	valueGetCharFunction_Set()

	ret = valueGetCharFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int8()

	return retGo
}

// UNSUPPORTED : C value 'g_value_get_double' : return type 'gdouble' not supported

var valueGetEnumFunction *gi.Function
var valueGetEnumFunction_Once sync.Once

func valueGetEnumFunction_Set() {
	valueGetEnumFunction_Once.Do(func() {
		valueStruct_Set()
		valueGetEnumFunction = valueStruct.InvokerNew("get_enum")
	})
}

// GetEnum is a representation of the C type g_value_get_enum.
func (recv *Value) GetEnum() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	valueGetEnumFunction_Set()

	ret = valueGetEnumFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var valueGetFlagsFunction *gi.Function
var valueGetFlagsFunction_Once sync.Once

func valueGetFlagsFunction_Set() {
	valueGetFlagsFunction_Once.Do(func() {
		valueStruct_Set()
		valueGetFlagsFunction = valueStruct.InvokerNew("get_flags")
	})
}

// GetFlags is a representation of the C type g_value_get_flags.
func (recv *Value) GetFlags() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	valueGetFlagsFunction_Set()

	ret = valueGetFlagsFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_value_get_float' : return type 'gfloat' not supported

// UNSUPPORTED : C value 'g_value_get_gtype' : return type 'GType' not supported

var valueGetIntFunction *gi.Function
var valueGetIntFunction_Once sync.Once

func valueGetIntFunction_Set() {
	valueGetIntFunction_Once.Do(func() {
		valueStruct_Set()
		valueGetIntFunction = valueStruct.InvokerNew("get_int")
	})
}

// GetInt is a representation of the C type g_value_get_int.
func (recv *Value) GetInt() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	valueGetIntFunction_Set()

	ret = valueGetIntFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var valueGetInt64Function *gi.Function
var valueGetInt64Function_Once sync.Once

func valueGetInt64Function_Set() {
	valueGetInt64Function_Once.Do(func() {
		valueStruct_Set()
		valueGetInt64Function = valueStruct.InvokerNew("get_int64")
	})
}

// GetInt64 is a representation of the C type g_value_get_int64.
func (recv *Value) GetInt64() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	valueGetInt64Function_Set()

	ret = valueGetInt64Function.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var valueGetLongFunction *gi.Function
var valueGetLongFunction_Once sync.Once

func valueGetLongFunction_Set() {
	valueGetLongFunction_Once.Do(func() {
		valueStruct_Set()
		valueGetLongFunction = valueStruct.InvokerNew("get_long")
	})
}

// GetLong is a representation of the C type g_value_get_long.
func (recv *Value) GetLong() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	valueGetLongFunction_Set()

	ret = valueGetLongFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'g_value_get_object' : return type 'Object' not supported

// UNSUPPORTED : C value 'g_value_get_param' : return type 'ParamSpec' not supported

// UNSUPPORTED : C value 'g_value_get_pointer' : return type 'gpointer' not supported

var valueGetScharFunction *gi.Function
var valueGetScharFunction_Once sync.Once

func valueGetScharFunction_Set() {
	valueGetScharFunction_Once.Do(func() {
		valueStruct_Set()
		valueGetScharFunction = valueStruct.InvokerNew("get_schar")
	})
}

// GetSchar is a representation of the C type g_value_get_schar.
func (recv *Value) GetSchar() int8 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	valueGetScharFunction_Set()

	ret = valueGetScharFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int8()

	return retGo
}

var valueGetStringFunction *gi.Function
var valueGetStringFunction_Once sync.Once

func valueGetStringFunction_Set() {
	valueGetStringFunction_Once.Do(func() {
		valueStruct_Set()
		valueGetStringFunction = valueStruct.InvokerNew("get_string")
	})
}

// GetString is a representation of the C type g_value_get_string.
func (recv *Value) GetString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	valueGetStringFunction_Set()

	ret = valueGetStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var valueGetUcharFunction *gi.Function
var valueGetUcharFunction_Once sync.Once

func valueGetUcharFunction_Set() {
	valueGetUcharFunction_Once.Do(func() {
		valueStruct_Set()
		valueGetUcharFunction = valueStruct.InvokerNew("get_uchar")
	})
}

// GetUchar is a representation of the C type g_value_get_uchar.
func (recv *Value) GetUchar() uint8 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	valueGetUcharFunction_Set()

	ret = valueGetUcharFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint8()

	return retGo
}

var valueGetUintFunction *gi.Function
var valueGetUintFunction_Once sync.Once

func valueGetUintFunction_Set() {
	valueGetUintFunction_Once.Do(func() {
		valueStruct_Set()
		valueGetUintFunction = valueStruct.InvokerNew("get_uint")
	})
}

// GetUint is a representation of the C type g_value_get_uint.
func (recv *Value) GetUint() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	valueGetUintFunction_Set()

	ret = valueGetUintFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var valueGetUint64Function *gi.Function
var valueGetUint64Function_Once sync.Once

func valueGetUint64Function_Set() {
	valueGetUint64Function_Once.Do(func() {
		valueStruct_Set()
		valueGetUint64Function = valueStruct.InvokerNew("get_uint64")
	})
}

// GetUint64 is a representation of the C type g_value_get_uint64.
func (recv *Value) GetUint64() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	valueGetUint64Function_Set()

	ret = valueGetUint64Function.Invoke(inArgs[:], nil)

	retGo := ret.Uint64()

	return retGo
}

var valueGetUlongFunction *gi.Function
var valueGetUlongFunction_Once sync.Once

func valueGetUlongFunction_Set() {
	valueGetUlongFunction_Once.Do(func() {
		valueStruct_Set()
		valueGetUlongFunction = valueStruct.InvokerNew("get_ulong")
	})
}

// GetUlong is a representation of the C type g_value_get_ulong.
func (recv *Value) GetUlong() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	valueGetUlongFunction_Set()

	ret = valueGetUlongFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint64()

	return retGo
}

// UNSUPPORTED : C value 'g_value_get_variant' : return type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_value_init' : parameter 'g_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_value_init_from_instance' : parameter 'instance' of type 'TypeInstance' not supported

// UNSUPPORTED : C value 'g_value_peek_pointer' : return type 'gpointer' not supported

var valueResetFunction *gi.Function
var valueResetFunction_Once sync.Once

func valueResetFunction_Set() {
	valueResetFunction_Once.Do(func() {
		valueStruct_Set()
		valueResetFunction = valueStruct.InvokerNew("reset")
	})
}

// Reset is a representation of the C type g_value_reset.
func (recv *Value) Reset() *Value {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	valueResetFunction_Set()

	ret = valueResetFunction.Invoke(inArgs[:], nil)

	retGo := &Value{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_value_set_boolean' : parameter 'v_boolean' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_value_set_boxed' : parameter 'v_boxed' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_value_set_boxed_take_ownership' : parameter 'v_boxed' of type 'gpointer' not supported

var valueSetCharFunction *gi.Function
var valueSetCharFunction_Once sync.Once

func valueSetCharFunction_Set() {
	valueSetCharFunction_Once.Do(func() {
		valueStruct_Set()
		valueSetCharFunction = valueStruct.InvokerNew("set_char")
	})
}

// SetChar is a representation of the C type g_value_set_char.
func (recv *Value) SetChar(vChar int8) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt8(vChar)

	valueSetCharFunction_Set()

	valueSetCharFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_value_set_double' : parameter 'v_double' of type 'gdouble' not supported

var valueSetEnumFunction *gi.Function
var valueSetEnumFunction_Once sync.Once

func valueSetEnumFunction_Set() {
	valueSetEnumFunction_Once.Do(func() {
		valueStruct_Set()
		valueSetEnumFunction = valueStruct.InvokerNew("set_enum")
	})
}

// SetEnum is a representation of the C type g_value_set_enum.
func (recv *Value) SetEnum(vEnum int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(vEnum)

	valueSetEnumFunction_Set()

	valueSetEnumFunction.Invoke(inArgs[:], nil)

}

var valueSetFlagsFunction *gi.Function
var valueSetFlagsFunction_Once sync.Once

func valueSetFlagsFunction_Set() {
	valueSetFlagsFunction_Once.Do(func() {
		valueStruct_Set()
		valueSetFlagsFunction = valueStruct.InvokerNew("set_flags")
	})
}

// SetFlags is a representation of the C type g_value_set_flags.
func (recv *Value) SetFlags(vFlags uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(vFlags)

	valueSetFlagsFunction_Set()

	valueSetFlagsFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_value_set_float' : parameter 'v_float' of type 'gfloat' not supported

// UNSUPPORTED : C value 'g_value_set_gtype' : parameter 'v_gtype' of type 'GType' not supported

// UNSUPPORTED : C value 'g_value_set_instance' : parameter 'instance' of type 'gpointer' not supported

var valueSetIntFunction *gi.Function
var valueSetIntFunction_Once sync.Once

func valueSetIntFunction_Set() {
	valueSetIntFunction_Once.Do(func() {
		valueStruct_Set()
		valueSetIntFunction = valueStruct.InvokerNew("set_int")
	})
}

// SetInt is a representation of the C type g_value_set_int.
func (recv *Value) SetInt(vInt int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(vInt)

	valueSetIntFunction_Set()

	valueSetIntFunction.Invoke(inArgs[:], nil)

}

var valueSetInt64Function *gi.Function
var valueSetInt64Function_Once sync.Once

func valueSetInt64Function_Set() {
	valueSetInt64Function_Once.Do(func() {
		valueStruct_Set()
		valueSetInt64Function = valueStruct.InvokerNew("set_int64")
	})
}

// SetInt64 is a representation of the C type g_value_set_int64.
func (recv *Value) SetInt64(vInt64 int64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(vInt64)

	valueSetInt64Function_Set()

	valueSetInt64Function.Invoke(inArgs[:], nil)

}

var valueSetLongFunction *gi.Function
var valueSetLongFunction_Once sync.Once

func valueSetLongFunction_Set() {
	valueSetLongFunction_Once.Do(func() {
		valueStruct_Set()
		valueSetLongFunction = valueStruct.InvokerNew("set_long")
	})
}

// SetLong is a representation of the C type g_value_set_long.
func (recv *Value) SetLong(vLong int64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(vLong)

	valueSetLongFunction_Set()

	valueSetLongFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_value_set_object' : parameter 'v_object' of type 'Object' not supported

// UNSUPPORTED : C value 'g_value_set_object_take_ownership' : parameter 'v_object' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_value_set_param' : parameter 'param' of type 'ParamSpec' not supported

// UNSUPPORTED : C value 'g_value_set_param_take_ownership' : parameter 'param' of type 'ParamSpec' not supported

// UNSUPPORTED : C value 'g_value_set_pointer' : parameter 'v_pointer' of type 'gpointer' not supported

var valueSetScharFunction *gi.Function
var valueSetScharFunction_Once sync.Once

func valueSetScharFunction_Set() {
	valueSetScharFunction_Once.Do(func() {
		valueStruct_Set()
		valueSetScharFunction = valueStruct.InvokerNew("set_schar")
	})
}

// SetSchar is a representation of the C type g_value_set_schar.
func (recv *Value) SetSchar(vChar int8) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt8(vChar)

	valueSetScharFunction_Set()

	valueSetScharFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_value_set_static_boxed' : parameter 'v_boxed' of type 'gpointer' not supported

var valueSetStaticStringFunction *gi.Function
var valueSetStaticStringFunction_Once sync.Once

func valueSetStaticStringFunction_Set() {
	valueSetStaticStringFunction_Once.Do(func() {
		valueStruct_Set()
		valueSetStaticStringFunction = valueStruct.InvokerNew("set_static_string")
	})
}

// SetStaticString is a representation of the C type g_value_set_static_string.
func (recv *Value) SetStaticString(vString string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(vString)

	valueSetStaticStringFunction_Set()

	valueSetStaticStringFunction.Invoke(inArgs[:], nil)

}

var valueSetStringFunction *gi.Function
var valueSetStringFunction_Once sync.Once

func valueSetStringFunction_Set() {
	valueSetStringFunction_Once.Do(func() {
		valueStruct_Set()
		valueSetStringFunction = valueStruct.InvokerNew("set_string")
	})
}

// SetString is a representation of the C type g_value_set_string.
func (recv *Value) SetString(vString string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(vString)

	valueSetStringFunction_Set()

	valueSetStringFunction.Invoke(inArgs[:], nil)

}

var valueSetStringTakeOwnershipFunction *gi.Function
var valueSetStringTakeOwnershipFunction_Once sync.Once

func valueSetStringTakeOwnershipFunction_Set() {
	valueSetStringTakeOwnershipFunction_Once.Do(func() {
		valueStruct_Set()
		valueSetStringTakeOwnershipFunction = valueStruct.InvokerNew("set_string_take_ownership")
	})
}

// SetStringTakeOwnership is a representation of the C type g_value_set_string_take_ownership.
func (recv *Value) SetStringTakeOwnership(vString string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(vString)

	valueSetStringTakeOwnershipFunction_Set()

	valueSetStringTakeOwnershipFunction.Invoke(inArgs[:], nil)

}

var valueSetUcharFunction *gi.Function
var valueSetUcharFunction_Once sync.Once

func valueSetUcharFunction_Set() {
	valueSetUcharFunction_Once.Do(func() {
		valueStruct_Set()
		valueSetUcharFunction = valueStruct.InvokerNew("set_uchar")
	})
}

// SetUchar is a representation of the C type g_value_set_uchar.
func (recv *Value) SetUchar(vUchar uint8) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint8(vUchar)

	valueSetUcharFunction_Set()

	valueSetUcharFunction.Invoke(inArgs[:], nil)

}

var valueSetUintFunction *gi.Function
var valueSetUintFunction_Once sync.Once

func valueSetUintFunction_Set() {
	valueSetUintFunction_Once.Do(func() {
		valueStruct_Set()
		valueSetUintFunction = valueStruct.InvokerNew("set_uint")
	})
}

// SetUint is a representation of the C type g_value_set_uint.
func (recv *Value) SetUint(vUint uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(vUint)

	valueSetUintFunction_Set()

	valueSetUintFunction.Invoke(inArgs[:], nil)

}

var valueSetUint64Function *gi.Function
var valueSetUint64Function_Once sync.Once

func valueSetUint64Function_Set() {
	valueSetUint64Function_Once.Do(func() {
		valueStruct_Set()
		valueSetUint64Function = valueStruct.InvokerNew("set_uint64")
	})
}

// SetUint64 is a representation of the C type g_value_set_uint64.
func (recv *Value) SetUint64(vUint64 uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint64(vUint64)

	valueSetUint64Function_Set()

	valueSetUint64Function.Invoke(inArgs[:], nil)

}

var valueSetUlongFunction *gi.Function
var valueSetUlongFunction_Once sync.Once

func valueSetUlongFunction_Set() {
	valueSetUlongFunction_Once.Do(func() {
		valueStruct_Set()
		valueSetUlongFunction = valueStruct.InvokerNew("set_ulong")
	})
}

// SetUlong is a representation of the C type g_value_set_ulong.
func (recv *Value) SetUlong(vUlong uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint64(vUlong)

	valueSetUlongFunction_Set()

	valueSetUlongFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_value_set_variant' : parameter 'variant' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_value_take_boxed' : parameter 'v_boxed' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_value_take_object' : parameter 'v_object' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_value_take_param' : parameter 'param' of type 'ParamSpec' not supported

var valueTakeStringFunction *gi.Function
var valueTakeStringFunction_Once sync.Once

func valueTakeStringFunction_Set() {
	valueTakeStringFunction_Once.Do(func() {
		valueStruct_Set()
		valueTakeStringFunction = valueStruct.InvokerNew("take_string")
	})
}

// TakeString is a representation of the C type g_value_take_string.
func (recv *Value) TakeString(vString string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(vString)

	valueTakeStringFunction_Set()

	valueTakeStringFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_value_take_variant' : parameter 'variant' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_value_transform' : parameter 'dest_value' of type 'Value' not supported

var valueUnsetFunction *gi.Function
var valueUnsetFunction_Once sync.Once

func valueUnsetFunction_Set() {
	valueUnsetFunction_Once.Do(func() {
		valueStruct_Set()
		valueUnsetFunction = valueStruct.InvokerNew("unset")
	})
}

// Unset is a representation of the C type g_value_unset.
func (recv *Value) Unset() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	valueUnsetFunction_Set()

	valueUnsetFunction.Invoke(inArgs[:], nil)

}

var valueArrayStruct *gi.Struct
var valueArrayStruct_Once sync.Once

func valueArrayStruct_Set() error {
	var err error
	valueArrayStruct_Once.Do(func() {
		valueArrayStruct, err = gi.StructNew("GObject", "ValueArray")
	})
	return err
}

type ValueArray struct {
	native  uintptr
	NValues uint32
	Values  *Value
}

var valueArrayNewFunction *gi.Function
var valueArrayNewFunction_Once sync.Once

func valueArrayNewFunction_Set() {
	valueArrayNewFunction_Once.Do(func() {
		valueArrayStruct_Set()
		valueArrayNewFunction = valueArrayStruct.InvokerNew("new")
	})
}

// ValueArrayNew is a representation of the C type g_value_array_new.
func ValueArrayNew(nPrealloced uint32) *ValueArray {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(nPrealloced)

	var ret gi.Argument

	valueArrayNewFunction_Set()

	ret = valueArrayNewFunction.Invoke(inArgs[:], nil)

	retGo := &ValueArray{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_value_array_append' : parameter 'value' of type 'Value' not supported

var valueArrayCopyFunction *gi.Function
var valueArrayCopyFunction_Once sync.Once

func valueArrayCopyFunction_Set() {
	valueArrayCopyFunction_Once.Do(func() {
		valueArrayStruct_Set()
		valueArrayCopyFunction = valueArrayStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type g_value_array_copy.
func (recv *ValueArray) Copy() *ValueArray {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	valueArrayCopyFunction_Set()

	ret = valueArrayCopyFunction.Invoke(inArgs[:], nil)

	retGo := &ValueArray{native: ret.Pointer()}

	return retGo
}

var valueArrayFreeFunction *gi.Function
var valueArrayFreeFunction_Once sync.Once

func valueArrayFreeFunction_Set() {
	valueArrayFreeFunction_Once.Do(func() {
		valueArrayStruct_Set()
		valueArrayFreeFunction = valueArrayStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type g_value_array_free.
func (recv *ValueArray) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	valueArrayFreeFunction_Set()

	valueArrayFreeFunction.Invoke(inArgs[:], nil)

}

var valueArrayGetNthFunction *gi.Function
var valueArrayGetNthFunction_Once sync.Once

func valueArrayGetNthFunction_Set() {
	valueArrayGetNthFunction_Once.Do(func() {
		valueArrayStruct_Set()
		valueArrayGetNthFunction = valueArrayStruct.InvokerNew("get_nth")
	})
}

// GetNth is a representation of the C type g_value_array_get_nth.
func (recv *ValueArray) GetNth(index uint32) *Value {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(index)

	var ret gi.Argument

	valueArrayGetNthFunction_Set()

	ret = valueArrayGetNthFunction.Invoke(inArgs[:], nil)

	retGo := &Value{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_value_array_insert' : parameter 'value' of type 'Value' not supported

// UNSUPPORTED : C value 'g_value_array_prepend' : parameter 'value' of type 'Value' not supported

var valueArrayRemoveFunction *gi.Function
var valueArrayRemoveFunction_Once sync.Once

func valueArrayRemoveFunction_Set() {
	valueArrayRemoveFunction_Once.Do(func() {
		valueArrayStruct_Set()
		valueArrayRemoveFunction = valueArrayStruct.InvokerNew("remove")
	})
}

// Remove is a representation of the C type g_value_array_remove.
func (recv *ValueArray) Remove(index uint32) *ValueArray {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(index)

	var ret gi.Argument

	valueArrayRemoveFunction_Set()

	ret = valueArrayRemoveFunction.Invoke(inArgs[:], nil)

	retGo := &ValueArray{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_value_array_sort' : parameter 'compare_func' of type 'GLib.CompareFunc' not supported

// UNSUPPORTED : C value 'g_value_array_sort_with_data' : parameter 'compare_func' of type 'GLib.CompareDataFunc' not supported

var weakRefStruct *gi.Struct
var weakRefStruct_Once sync.Once

func weakRefStruct_Set() error {
	var err error
	weakRefStruct_Once.Do(func() {
		weakRefStruct, err = gi.StructNew("GObject", "WeakRef")
	})
	return err
}

type WeakRef struct {
	native uintptr
}

var weakRefClearFunction *gi.Function
var weakRefClearFunction_Once sync.Once

func weakRefClearFunction_Set() {
	weakRefClearFunction_Once.Do(func() {
		weakRefStruct_Set()
		weakRefClearFunction = weakRefStruct.InvokerNew("clear")
	})
}

// Clear is a representation of the C type g_weak_ref_clear.
func (recv *WeakRef) Clear() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	weakRefClearFunction_Set()

	weakRefClearFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_weak_ref_get' : return type 'Object' not supported

// UNSUPPORTED : C value 'g_weak_ref_init' : parameter 'object' of type 'Object' not supported

// UNSUPPORTED : C value 'g_weak_ref_set' : parameter 'object' of type 'Object' not supported
