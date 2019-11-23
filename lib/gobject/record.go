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

func closureInvalidateFunction_Set() error {
	var err error
	closureInvalidateFunction_Once.Do(func() {
		err = closureStruct_Set()
		if err != nil {
			return
		}
		closureInvalidateFunction, err = closureStruct.InvokerNew("invalidate")
	})
	return err
}

// Invalidate is a representation of the C type g_closure_invalidate.
func (recv *Closure) Invalidate() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := closureInvalidateFunction_Set()
	if err == nil {
		closureInvalidateFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_closure_invoke' : parameter 'return_value' of type 'Value' not supported

var closureRefFunction *gi.Function
var closureRefFunction_Once sync.Once

func closureRefFunction_Set() error {
	var err error
	closureRefFunction_Once.Do(func() {
		err = closureStruct_Set()
		if err != nil {
			return
		}
		closureRefFunction, err = closureStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_closure_ref.
func (recv *Closure) Ref() (*Closure, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := closureRefFunction_Set()
	if err == nil {
		ret = closureRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Closure{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_closure_remove_finalize_notifier' : parameter 'notify_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_closure_remove_invalidate_notifier' : parameter 'notify_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_closure_set_marshal' : parameter 'marshal' of type 'ClosureMarshal' not supported

// UNSUPPORTED : C value 'g_closure_set_meta_marshal' : parameter 'marshal_data' of type 'gpointer' not supported

var closureSinkFunction *gi.Function
var closureSinkFunction_Once sync.Once

func closureSinkFunction_Set() error {
	var err error
	closureSinkFunction_Once.Do(func() {
		err = closureStruct_Set()
		if err != nil {
			return
		}
		closureSinkFunction, err = closureStruct.InvokerNew("sink")
	})
	return err
}

// Sink is a representation of the C type g_closure_sink.
func (recv *Closure) Sink() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := closureSinkFunction_Set()
	if err == nil {
		closureSinkFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var closureUnrefFunction *gi.Function
var closureUnrefFunction_Once sync.Once

func closureUnrefFunction_Set() error {
	var err error
	closureUnrefFunction_Once.Do(func() {
		err = closureStruct_Set()
		if err != nil {
			return
		}
		closureUnrefFunction, err = closureStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_closure_unref.
func (recv *Closure) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := closureUnrefFunction_Set()
	if err == nil {
		closureUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
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

func objectClassListPropertiesFunction_Set() error {
	var err error
	objectClassListPropertiesFunction_Once.Do(func() {
		err = objectClassStruct_Set()
		if err != nil {
			return
		}
		objectClassListPropertiesFunction, err = objectClassStruct.InvokerNew("list_properties")
	})
	return err
}

// ListProperties is a representation of the C type g_object_class_list_properties.
func (recv *ObjectClass) ListProperties() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	err := objectClassListPropertiesFunction_Set()
	if err == nil {
		objectClassListPropertiesFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint32()

	return out0, err
}

var objectClassOverridePropertyFunction *gi.Function
var objectClassOverridePropertyFunction_Once sync.Once

func objectClassOverridePropertyFunction_Set() error {
	var err error
	objectClassOverridePropertyFunction_Once.Do(func() {
		err = objectClassStruct_Set()
		if err != nil {
			return
		}
		objectClassOverridePropertyFunction, err = objectClassStruct.InvokerNew("override_property")
	})
	return err
}

// OverrideProperty is a representation of the C type g_object_class_override_property.
func (recv *ObjectClass) OverrideProperty(propertyId uint32, name string) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(propertyId)
	inArgs[2].SetString(name)

	err := objectClassOverridePropertyFunction_Set()
	if err == nil {
		objectClassOverridePropertyFunction.Invoke(inArgs[:], nil)
	}

	return err
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

func typeClassGetInstancePrivateOffsetFunction_Set() error {
	var err error
	typeClassGetInstancePrivateOffsetFunction_Once.Do(func() {
		err = typeClassStruct_Set()
		if err != nil {
			return
		}
		typeClassGetInstancePrivateOffsetFunction, err = typeClassStruct.InvokerNew("get_instance_private_offset")
	})
	return err
}

// GetInstancePrivateOffset is a representation of the C type g_type_class_get_instance_private_offset.
func (recv *TypeClass) GetInstancePrivateOffset() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := typeClassGetInstancePrivateOffsetFunction_Set()
	if err == nil {
		ret = typeClassGetInstancePrivateOffsetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_type_class_get_private' : parameter 'private_type' of type 'GType' not supported

var typeClassPeekParentFunction *gi.Function
var typeClassPeekParentFunction_Once sync.Once

func typeClassPeekParentFunction_Set() error {
	var err error
	typeClassPeekParentFunction_Once.Do(func() {
		err = typeClassStruct_Set()
		if err != nil {
			return
		}
		typeClassPeekParentFunction, err = typeClassStruct.InvokerNew("peek_parent")
	})
	return err
}

// PeekParent is a representation of the C type g_type_class_peek_parent.
func (recv *TypeClass) PeekParent() (*TypeClass, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := typeClassPeekParentFunction_Set()
	if err == nil {
		ret = typeClassPeekParentFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TypeClass{native: ret.Pointer()}

	return retGo, err
}

var typeClassUnrefFunction *gi.Function
var typeClassUnrefFunction_Once sync.Once

func typeClassUnrefFunction_Set() error {
	var err error
	typeClassUnrefFunction_Once.Do(func() {
		err = typeClassStruct_Set()
		if err != nil {
			return
		}
		typeClassUnrefFunction, err = typeClassStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_type_class_unref.
func (recv *TypeClass) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := typeClassUnrefFunction_Set()
	if err == nil {
		typeClassUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var typeClassUnrefUncachedFunction *gi.Function
var typeClassUnrefUncachedFunction_Once sync.Once

func typeClassUnrefUncachedFunction_Set() error {
	var err error
	typeClassUnrefUncachedFunction_Once.Do(func() {
		err = typeClassStruct_Set()
		if err != nil {
			return
		}
		typeClassUnrefUncachedFunction, err = typeClassStruct.InvokerNew("unref_uncached")
	})
	return err
}

// UnrefUncached is a representation of the C type g_type_class_unref_uncached.
func (recv *TypeClass) UnrefUncached() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := typeClassUnrefUncachedFunction_Set()
	if err == nil {
		typeClassUnrefUncachedFunction.Invoke(inArgs[:], nil)
	}

	return err
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

func typeInterfacePeekParentFunction_Set() error {
	var err error
	typeInterfacePeekParentFunction_Once.Do(func() {
		err = typeInterfaceStruct_Set()
		if err != nil {
			return
		}
		typeInterfacePeekParentFunction, err = typeInterfaceStruct.InvokerNew("peek_parent")
	})
	return err
}

// PeekParent is a representation of the C type g_type_interface_peek_parent.
func (recv *TypeInterface) PeekParent() (*TypeInterface, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := typeInterfacePeekParentFunction_Set()
	if err == nil {
		ret = typeInterfacePeekParentFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TypeInterface{native: ret.Pointer()}

	return retGo, err
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

func valueDupStringFunction_Set() error {
	var err error
	valueDupStringFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueDupStringFunction, err = valueStruct.InvokerNew("dup_string")
	})
	return err
}

// DupString is a representation of the C type g_value_dup_string.
func (recv *Value) DupString() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueDupStringFunction_Set()
	if err == nil {
		ret = valueDupStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_value_dup_variant' : return type 'GLib.Variant' not supported

var valueFitsPointerFunction *gi.Function
var valueFitsPointerFunction_Once sync.Once

func valueFitsPointerFunction_Set() error {
	var err error
	valueFitsPointerFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueFitsPointerFunction, err = valueStruct.InvokerNew("fits_pointer")
	})
	return err
}

// FitsPointer is a representation of the C type g_value_fits_pointer.
func (recv *Value) FitsPointer() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueFitsPointerFunction_Set()
	if err == nil {
		ret = valueFitsPointerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var valueGetBooleanFunction *gi.Function
var valueGetBooleanFunction_Once sync.Once

func valueGetBooleanFunction_Set() error {
	var err error
	valueGetBooleanFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueGetBooleanFunction, err = valueStruct.InvokerNew("get_boolean")
	})
	return err
}

// GetBoolean is a representation of the C type g_value_get_boolean.
func (recv *Value) GetBoolean() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetBooleanFunction_Set()
	if err == nil {
		ret = valueGetBooleanFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_value_get_boxed' : return type 'gpointer' not supported

var valueGetCharFunction *gi.Function
var valueGetCharFunction_Once sync.Once

func valueGetCharFunction_Set() error {
	var err error
	valueGetCharFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueGetCharFunction, err = valueStruct.InvokerNew("get_char")
	})
	return err
}

// GetChar is a representation of the C type g_value_get_char.
func (recv *Value) GetChar() (int8, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetCharFunction_Set()
	if err == nil {
		ret = valueGetCharFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int8()

	return retGo, err
}

// UNSUPPORTED : C value 'g_value_get_double' : return type 'gdouble' not supported

var valueGetEnumFunction *gi.Function
var valueGetEnumFunction_Once sync.Once

func valueGetEnumFunction_Set() error {
	var err error
	valueGetEnumFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueGetEnumFunction, err = valueStruct.InvokerNew("get_enum")
	})
	return err
}

// GetEnum is a representation of the C type g_value_get_enum.
func (recv *Value) GetEnum() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetEnumFunction_Set()
	if err == nil {
		ret = valueGetEnumFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var valueGetFlagsFunction *gi.Function
var valueGetFlagsFunction_Once sync.Once

func valueGetFlagsFunction_Set() error {
	var err error
	valueGetFlagsFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueGetFlagsFunction, err = valueStruct.InvokerNew("get_flags")
	})
	return err
}

// GetFlags is a representation of the C type g_value_get_flags.
func (recv *Value) GetFlags() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetFlagsFunction_Set()
	if err == nil {
		ret = valueGetFlagsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_value_get_float' : return type 'gfloat' not supported

// UNSUPPORTED : C value 'g_value_get_gtype' : return type 'GType' not supported

var valueGetIntFunction *gi.Function
var valueGetIntFunction_Once sync.Once

func valueGetIntFunction_Set() error {
	var err error
	valueGetIntFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueGetIntFunction, err = valueStruct.InvokerNew("get_int")
	})
	return err
}

// GetInt is a representation of the C type g_value_get_int.
func (recv *Value) GetInt() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetIntFunction_Set()
	if err == nil {
		ret = valueGetIntFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var valueGetInt64Function *gi.Function
var valueGetInt64Function_Once sync.Once

func valueGetInt64Function_Set() error {
	var err error
	valueGetInt64Function_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueGetInt64Function, err = valueStruct.InvokerNew("get_int64")
	})
	return err
}

// GetInt64 is a representation of the C type g_value_get_int64.
func (recv *Value) GetInt64() (int64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetInt64Function_Set()
	if err == nil {
		ret = valueGetInt64Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo, err
}

var valueGetLongFunction *gi.Function
var valueGetLongFunction_Once sync.Once

func valueGetLongFunction_Set() error {
	var err error
	valueGetLongFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueGetLongFunction, err = valueStruct.InvokerNew("get_long")
	})
	return err
}

// GetLong is a representation of the C type g_value_get_long.
func (recv *Value) GetLong() (int64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetLongFunction_Set()
	if err == nil {
		ret = valueGetLongFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo, err
}

// UNSUPPORTED : C value 'g_value_get_object' : return type 'Object' not supported

// UNSUPPORTED : C value 'g_value_get_param' : return type 'ParamSpec' not supported

// UNSUPPORTED : C value 'g_value_get_pointer' : return type 'gpointer' not supported

var valueGetScharFunction *gi.Function
var valueGetScharFunction_Once sync.Once

func valueGetScharFunction_Set() error {
	var err error
	valueGetScharFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueGetScharFunction, err = valueStruct.InvokerNew("get_schar")
	})
	return err
}

// GetSchar is a representation of the C type g_value_get_schar.
func (recv *Value) GetSchar() (int8, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetScharFunction_Set()
	if err == nil {
		ret = valueGetScharFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int8()

	return retGo, err
}

var valueGetStringFunction *gi.Function
var valueGetStringFunction_Once sync.Once

func valueGetStringFunction_Set() error {
	var err error
	valueGetStringFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueGetStringFunction, err = valueStruct.InvokerNew("get_string")
	})
	return err
}

// GetString is a representation of the C type g_value_get_string.
func (recv *Value) GetString() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetStringFunction_Set()
	if err == nil {
		ret = valueGetStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var valueGetUcharFunction *gi.Function
var valueGetUcharFunction_Once sync.Once

func valueGetUcharFunction_Set() error {
	var err error
	valueGetUcharFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueGetUcharFunction, err = valueStruct.InvokerNew("get_uchar")
	})
	return err
}

// GetUchar is a representation of the C type g_value_get_uchar.
func (recv *Value) GetUchar() (uint8, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetUcharFunction_Set()
	if err == nil {
		ret = valueGetUcharFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint8()

	return retGo, err
}

var valueGetUintFunction *gi.Function
var valueGetUintFunction_Once sync.Once

func valueGetUintFunction_Set() error {
	var err error
	valueGetUintFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueGetUintFunction, err = valueStruct.InvokerNew("get_uint")
	})
	return err
}

// GetUint is a representation of the C type g_value_get_uint.
func (recv *Value) GetUint() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetUintFunction_Set()
	if err == nil {
		ret = valueGetUintFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var valueGetUint64Function *gi.Function
var valueGetUint64Function_Once sync.Once

func valueGetUint64Function_Set() error {
	var err error
	valueGetUint64Function_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueGetUint64Function, err = valueStruct.InvokerNew("get_uint64")
	})
	return err
}

// GetUint64 is a representation of the C type g_value_get_uint64.
func (recv *Value) GetUint64() (uint64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetUint64Function_Set()
	if err == nil {
		ret = valueGetUint64Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo, err
}

var valueGetUlongFunction *gi.Function
var valueGetUlongFunction_Once sync.Once

func valueGetUlongFunction_Set() error {
	var err error
	valueGetUlongFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueGetUlongFunction, err = valueStruct.InvokerNew("get_ulong")
	})
	return err
}

// GetUlong is a representation of the C type g_value_get_ulong.
func (recv *Value) GetUlong() (uint64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetUlongFunction_Set()
	if err == nil {
		ret = valueGetUlongFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo, err
}

// UNSUPPORTED : C value 'g_value_get_variant' : return type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_value_init' : parameter 'g_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_value_init_from_instance' : parameter 'instance' of type 'TypeInstance' not supported

// UNSUPPORTED : C value 'g_value_peek_pointer' : return type 'gpointer' not supported

var valueResetFunction *gi.Function
var valueResetFunction_Once sync.Once

func valueResetFunction_Set() error {
	var err error
	valueResetFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueResetFunction, err = valueStruct.InvokerNew("reset")
	})
	return err
}

// Reset is a representation of the C type g_value_reset.
func (recv *Value) Reset() (*Value, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueResetFunction_Set()
	if err == nil {
		ret = valueResetFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Value{native: ret.Pointer()}

	return retGo, err
}

var valueSetBooleanFunction *gi.Function
var valueSetBooleanFunction_Once sync.Once

func valueSetBooleanFunction_Set() error {
	var err error
	valueSetBooleanFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueSetBooleanFunction, err = valueStruct.InvokerNew("set_boolean")
	})
	return err
}

// SetBoolean is a representation of the C type g_value_set_boolean.
func (recv *Value) SetBoolean(vBoolean bool) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(vBoolean)

	err := valueSetBooleanFunction_Set()
	if err == nil {
		valueSetBooleanFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_value_set_boxed' : parameter 'v_boxed' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_value_set_boxed_take_ownership' : parameter 'v_boxed' of type 'gpointer' not supported

var valueSetCharFunction *gi.Function
var valueSetCharFunction_Once sync.Once

func valueSetCharFunction_Set() error {
	var err error
	valueSetCharFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueSetCharFunction, err = valueStruct.InvokerNew("set_char")
	})
	return err
}

// SetChar is a representation of the C type g_value_set_char.
func (recv *Value) SetChar(vChar int8) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt8(vChar)

	err := valueSetCharFunction_Set()
	if err == nil {
		valueSetCharFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_value_set_double' : parameter 'v_double' of type 'gdouble' not supported

var valueSetEnumFunction *gi.Function
var valueSetEnumFunction_Once sync.Once

func valueSetEnumFunction_Set() error {
	var err error
	valueSetEnumFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueSetEnumFunction, err = valueStruct.InvokerNew("set_enum")
	})
	return err
}

// SetEnum is a representation of the C type g_value_set_enum.
func (recv *Value) SetEnum(vEnum int32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(vEnum)

	err := valueSetEnumFunction_Set()
	if err == nil {
		valueSetEnumFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var valueSetFlagsFunction *gi.Function
var valueSetFlagsFunction_Once sync.Once

func valueSetFlagsFunction_Set() error {
	var err error
	valueSetFlagsFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueSetFlagsFunction, err = valueStruct.InvokerNew("set_flags")
	})
	return err
}

// SetFlags is a representation of the C type g_value_set_flags.
func (recv *Value) SetFlags(vFlags uint32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(vFlags)

	err := valueSetFlagsFunction_Set()
	if err == nil {
		valueSetFlagsFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_value_set_float' : parameter 'v_float' of type 'gfloat' not supported

// UNSUPPORTED : C value 'g_value_set_gtype' : parameter 'v_gtype' of type 'GType' not supported

// UNSUPPORTED : C value 'g_value_set_instance' : parameter 'instance' of type 'gpointer' not supported

var valueSetIntFunction *gi.Function
var valueSetIntFunction_Once sync.Once

func valueSetIntFunction_Set() error {
	var err error
	valueSetIntFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueSetIntFunction, err = valueStruct.InvokerNew("set_int")
	})
	return err
}

// SetInt is a representation of the C type g_value_set_int.
func (recv *Value) SetInt(vInt int32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(vInt)

	err := valueSetIntFunction_Set()
	if err == nil {
		valueSetIntFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var valueSetInt64Function *gi.Function
var valueSetInt64Function_Once sync.Once

func valueSetInt64Function_Set() error {
	var err error
	valueSetInt64Function_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueSetInt64Function, err = valueStruct.InvokerNew("set_int64")
	})
	return err
}

// SetInt64 is a representation of the C type g_value_set_int64.
func (recv *Value) SetInt64(vInt64 int64) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(vInt64)

	err := valueSetInt64Function_Set()
	if err == nil {
		valueSetInt64Function.Invoke(inArgs[:], nil)
	}

	return err
}

var valueSetLongFunction *gi.Function
var valueSetLongFunction_Once sync.Once

func valueSetLongFunction_Set() error {
	var err error
	valueSetLongFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueSetLongFunction, err = valueStruct.InvokerNew("set_long")
	})
	return err
}

// SetLong is a representation of the C type g_value_set_long.
func (recv *Value) SetLong(vLong int64) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(vLong)

	err := valueSetLongFunction_Set()
	if err == nil {
		valueSetLongFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_value_set_object' : parameter 'v_object' of type 'Object' not supported

// UNSUPPORTED : C value 'g_value_set_object_take_ownership' : parameter 'v_object' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_value_set_param' : parameter 'param' of type 'ParamSpec' not supported

// UNSUPPORTED : C value 'g_value_set_param_take_ownership' : parameter 'param' of type 'ParamSpec' not supported

// UNSUPPORTED : C value 'g_value_set_pointer' : parameter 'v_pointer' of type 'gpointer' not supported

var valueSetScharFunction *gi.Function
var valueSetScharFunction_Once sync.Once

func valueSetScharFunction_Set() error {
	var err error
	valueSetScharFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueSetScharFunction, err = valueStruct.InvokerNew("set_schar")
	})
	return err
}

// SetSchar is a representation of the C type g_value_set_schar.
func (recv *Value) SetSchar(vChar int8) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt8(vChar)

	err := valueSetScharFunction_Set()
	if err == nil {
		valueSetScharFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_value_set_static_boxed' : parameter 'v_boxed' of type 'gpointer' not supported

var valueSetStaticStringFunction *gi.Function
var valueSetStaticStringFunction_Once sync.Once

func valueSetStaticStringFunction_Set() error {
	var err error
	valueSetStaticStringFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueSetStaticStringFunction, err = valueStruct.InvokerNew("set_static_string")
	})
	return err
}

// SetStaticString is a representation of the C type g_value_set_static_string.
func (recv *Value) SetStaticString(vString string) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(vString)

	err := valueSetStaticStringFunction_Set()
	if err == nil {
		valueSetStaticStringFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var valueSetStringFunction *gi.Function
var valueSetStringFunction_Once sync.Once

func valueSetStringFunction_Set() error {
	var err error
	valueSetStringFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueSetStringFunction, err = valueStruct.InvokerNew("set_string")
	})
	return err
}

// SetString is a representation of the C type g_value_set_string.
func (recv *Value) SetString(vString string) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(vString)

	err := valueSetStringFunction_Set()
	if err == nil {
		valueSetStringFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var valueSetStringTakeOwnershipFunction *gi.Function
var valueSetStringTakeOwnershipFunction_Once sync.Once

func valueSetStringTakeOwnershipFunction_Set() error {
	var err error
	valueSetStringTakeOwnershipFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueSetStringTakeOwnershipFunction, err = valueStruct.InvokerNew("set_string_take_ownership")
	})
	return err
}

// SetStringTakeOwnership is a representation of the C type g_value_set_string_take_ownership.
func (recv *Value) SetStringTakeOwnership(vString string) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(vString)

	err := valueSetStringTakeOwnershipFunction_Set()
	if err == nil {
		valueSetStringTakeOwnershipFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var valueSetUcharFunction *gi.Function
var valueSetUcharFunction_Once sync.Once

func valueSetUcharFunction_Set() error {
	var err error
	valueSetUcharFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueSetUcharFunction, err = valueStruct.InvokerNew("set_uchar")
	})
	return err
}

// SetUchar is a representation of the C type g_value_set_uchar.
func (recv *Value) SetUchar(vUchar uint8) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint8(vUchar)

	err := valueSetUcharFunction_Set()
	if err == nil {
		valueSetUcharFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var valueSetUintFunction *gi.Function
var valueSetUintFunction_Once sync.Once

func valueSetUintFunction_Set() error {
	var err error
	valueSetUintFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueSetUintFunction, err = valueStruct.InvokerNew("set_uint")
	})
	return err
}

// SetUint is a representation of the C type g_value_set_uint.
func (recv *Value) SetUint(vUint uint32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(vUint)

	err := valueSetUintFunction_Set()
	if err == nil {
		valueSetUintFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var valueSetUint64Function *gi.Function
var valueSetUint64Function_Once sync.Once

func valueSetUint64Function_Set() error {
	var err error
	valueSetUint64Function_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueSetUint64Function, err = valueStruct.InvokerNew("set_uint64")
	})
	return err
}

// SetUint64 is a representation of the C type g_value_set_uint64.
func (recv *Value) SetUint64(vUint64 uint64) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint64(vUint64)

	err := valueSetUint64Function_Set()
	if err == nil {
		valueSetUint64Function.Invoke(inArgs[:], nil)
	}

	return err
}

var valueSetUlongFunction *gi.Function
var valueSetUlongFunction_Once sync.Once

func valueSetUlongFunction_Set() error {
	var err error
	valueSetUlongFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueSetUlongFunction, err = valueStruct.InvokerNew("set_ulong")
	})
	return err
}

// SetUlong is a representation of the C type g_value_set_ulong.
func (recv *Value) SetUlong(vUlong uint64) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint64(vUlong)

	err := valueSetUlongFunction_Set()
	if err == nil {
		valueSetUlongFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_value_set_variant' : parameter 'variant' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_value_take_boxed' : parameter 'v_boxed' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_value_take_object' : parameter 'v_object' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_value_take_param' : parameter 'param' of type 'ParamSpec' not supported

var valueTakeStringFunction *gi.Function
var valueTakeStringFunction_Once sync.Once

func valueTakeStringFunction_Set() error {
	var err error
	valueTakeStringFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueTakeStringFunction, err = valueStruct.InvokerNew("take_string")
	})
	return err
}

// TakeString is a representation of the C type g_value_take_string.
func (recv *Value) TakeString(vString string) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(vString)

	err := valueTakeStringFunction_Set()
	if err == nil {
		valueTakeStringFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_value_take_variant' : parameter 'variant' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_value_transform' : parameter 'dest_value' of type 'Value' not supported

var valueUnsetFunction *gi.Function
var valueUnsetFunction_Once sync.Once

func valueUnsetFunction_Set() error {
	var err error
	valueUnsetFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueUnsetFunction, err = valueStruct.InvokerNew("unset")
	})
	return err
}

// Unset is a representation of the C type g_value_unset.
func (recv *Value) Unset() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := valueUnsetFunction_Set()
	if err == nil {
		valueUnsetFunction.Invoke(inArgs[:], nil)
	}

	return err
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

func valueArrayNewFunction_Set() error {
	var err error
	valueArrayNewFunction_Once.Do(func() {
		err = valueArrayStruct_Set()
		if err != nil {
			return
		}
		valueArrayNewFunction, err = valueArrayStruct.InvokerNew("new")
	})
	return err
}

// ValueArrayNew is a representation of the C type g_value_array_new.
func ValueArrayNew(nPrealloced uint32) (*ValueArray, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(nPrealloced)

	var ret gi.Argument

	err := valueArrayNewFunction_Set()
	if err == nil {
		ret = valueArrayNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &ValueArray{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_value_array_append' : parameter 'value' of type 'Value' not supported

var valueArrayCopyFunction *gi.Function
var valueArrayCopyFunction_Once sync.Once

func valueArrayCopyFunction_Set() error {
	var err error
	valueArrayCopyFunction_Once.Do(func() {
		err = valueArrayStruct_Set()
		if err != nil {
			return
		}
		valueArrayCopyFunction, err = valueArrayStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type g_value_array_copy.
func (recv *ValueArray) Copy() (*ValueArray, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueArrayCopyFunction_Set()
	if err == nil {
		ret = valueArrayCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &ValueArray{native: ret.Pointer()}

	return retGo, err
}

var valueArrayFreeFunction *gi.Function
var valueArrayFreeFunction_Once sync.Once

func valueArrayFreeFunction_Set() error {
	var err error
	valueArrayFreeFunction_Once.Do(func() {
		err = valueArrayStruct_Set()
		if err != nil {
			return
		}
		valueArrayFreeFunction, err = valueArrayStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_value_array_free.
func (recv *ValueArray) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := valueArrayFreeFunction_Set()
	if err == nil {
		valueArrayFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var valueArrayGetNthFunction *gi.Function
var valueArrayGetNthFunction_Once sync.Once

func valueArrayGetNthFunction_Set() error {
	var err error
	valueArrayGetNthFunction_Once.Do(func() {
		err = valueArrayStruct_Set()
		if err != nil {
			return
		}
		valueArrayGetNthFunction, err = valueArrayStruct.InvokerNew("get_nth")
	})
	return err
}

// GetNth is a representation of the C type g_value_array_get_nth.
func (recv *ValueArray) GetNth(index uint32) (*Value, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(index)

	var ret gi.Argument

	err := valueArrayGetNthFunction_Set()
	if err == nil {
		ret = valueArrayGetNthFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Value{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_value_array_insert' : parameter 'value' of type 'Value' not supported

// UNSUPPORTED : C value 'g_value_array_prepend' : parameter 'value' of type 'Value' not supported

var valueArrayRemoveFunction *gi.Function
var valueArrayRemoveFunction_Once sync.Once

func valueArrayRemoveFunction_Set() error {
	var err error
	valueArrayRemoveFunction_Once.Do(func() {
		err = valueArrayStruct_Set()
		if err != nil {
			return
		}
		valueArrayRemoveFunction, err = valueArrayStruct.InvokerNew("remove")
	})
	return err
}

// Remove is a representation of the C type g_value_array_remove.
func (recv *ValueArray) Remove(index uint32) (*ValueArray, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(index)

	var ret gi.Argument

	err := valueArrayRemoveFunction_Set()
	if err == nil {
		ret = valueArrayRemoveFunction.Invoke(inArgs[:], nil)
	}

	retGo := &ValueArray{native: ret.Pointer()}

	return retGo, err
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

func weakRefClearFunction_Set() error {
	var err error
	weakRefClearFunction_Once.Do(func() {
		err = weakRefStruct_Set()
		if err != nil {
			return
		}
		weakRefClearFunction, err = weakRefStruct.InvokerNew("clear")
	})
	return err
}

// Clear is a representation of the C type g_weak_ref_clear.
func (recv *WeakRef) Clear() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := weakRefClearFunction_Set()
	if err == nil {
		weakRefClearFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_weak_ref_get' : return type 'Object' not supported

// UNSUPPORTED : C value 'g_weak_ref_init' : parameter 'object' of type 'Object' not supported

// UNSUPPORTED : C value 'g_weak_ref_set' : parameter 'object' of type 'Object' not supported
