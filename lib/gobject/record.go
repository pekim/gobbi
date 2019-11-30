// Code generated - DO NOT EDIT.

package gobject

import (
	gi "github.com/pekim/gobbi/internal/gi"
	glib "github.com/pekim/gobbi/lib/glib"
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
	native uintptr
}

// Closure returns the C field 'closure'.
func (recv *CClosure) Closure() *Closure {
	argValue := gi.FieldGet(cClosureStruct, recv.native, "closure")
	value := &Closure{native: argValue.Pointer()}
	return value
}

// UNSUPPORTED : C value 'callback' : for field getter : no Go type for 'gpointer'

// CClosureStruct creates an uninitialised CClosure.
func CClosureStruct() *CClosure {
	err := cClosureStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CClosure{native: cClosureStruct.Alloc()}
	return structGo
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
	native uintptr
}

// InMarshal returns the C field 'in_marshal'.
func (recv *Closure) InMarshal() uint32 {
	argValue := gi.FieldGet(closureStruct, recv.native, "in_marshal")
	value := argValue.Uint32()
	return value
}

// IsInvalid returns the C field 'is_invalid'.
func (recv *Closure) IsInvalid() uint32 {
	argValue := gi.FieldGet(closureStruct, recv.native, "is_invalid")
	value := argValue.Uint32()
	return value
}

// UNSUPPORTED : C value 'marshal' : for field getter : missing Type

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
func (recv *Closure) Invalidate() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := closureInvalidateFunction_Set()
	if err == nil {
		closureInvalidateFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_closure_invoke' : parameter 'param_values' has no type

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
func (recv *Closure) Ref() *Closure {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := closureRefFunction_Set()
	if err == nil {
		ret = closureRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Closure{native: ret.Pointer()}

	return retGo
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
func (recv *Closure) Sink() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := closureSinkFunction_Set()
	if err == nil {
		closureSinkFunction.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *Closure) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := closureUnrefFunction_Set()
	if err == nil {
		closureUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
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
}

// UNSUPPORTED : C value 'data' : for field getter : no Go type for 'gpointer'

// UNSUPPORTED : C value 'notify' : for field getter : no Go type for 'ClosureNotify'

// ClosureNotifyDataStruct creates an uninitialised ClosureNotifyData.
func ClosureNotifyDataStruct() *ClosureNotifyData {
	err := closureNotifyDataStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ClosureNotifyData{native: closureNotifyDataStruct.Alloc()}
	return structGo
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
	native uintptr
}

// GTypeClass returns the C field 'g_type_class'.
func (recv *EnumClass) GTypeClass() *TypeClass {
	argValue := gi.FieldGet(enumClassStruct, recv.native, "g_type_class")
	value := &TypeClass{native: argValue.Pointer()}
	return value
}

// Minimum returns the C field 'minimum'.
func (recv *EnumClass) Minimum() int32 {
	argValue := gi.FieldGet(enumClassStruct, recv.native, "minimum")
	value := argValue.Int32()
	return value
}

// Maximum returns the C field 'maximum'.
func (recv *EnumClass) Maximum() int32 {
	argValue := gi.FieldGet(enumClassStruct, recv.native, "maximum")
	value := argValue.Int32()
	return value
}

// NValues returns the C field 'n_values'.
func (recv *EnumClass) NValues() uint32 {
	argValue := gi.FieldGet(enumClassStruct, recv.native, "n_values")
	value := argValue.Uint32()
	return value
}

// Values returns the C field 'values'.
func (recv *EnumClass) Values() *EnumValue {
	argValue := gi.FieldGet(enumClassStruct, recv.native, "values")
	value := &EnumValue{native: argValue.Pointer()}
	return value
}

// EnumClassStruct creates an uninitialised EnumClass.
func EnumClassStruct() *EnumClass {
	err := enumClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EnumClass{native: enumClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// Value returns the C field 'value'.
func (recv *EnumValue) Value() int32 {
	argValue := gi.FieldGet(enumValueStruct, recv.native, "value")
	value := argValue.Int32()
	return value
}

// ValueName returns the C field 'value_name'.
func (recv *EnumValue) ValueName() string {
	argValue := gi.FieldGet(enumValueStruct, recv.native, "value_name")
	value := argValue.String(false)
	return value
}

// ValueNick returns the C field 'value_nick'.
func (recv *EnumValue) ValueNick() string {
	argValue := gi.FieldGet(enumValueStruct, recv.native, "value_nick")
	value := argValue.String(false)
	return value
}

// EnumValueStruct creates an uninitialised EnumValue.
func EnumValueStruct() *EnumValue {
	err := enumValueStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EnumValue{native: enumValueStruct.Alloc()}
	return structGo
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
	native uintptr
}

// GTypeClass returns the C field 'g_type_class'.
func (recv *FlagsClass) GTypeClass() *TypeClass {
	argValue := gi.FieldGet(flagsClassStruct, recv.native, "g_type_class")
	value := &TypeClass{native: argValue.Pointer()}
	return value
}

// Mask returns the C field 'mask'.
func (recv *FlagsClass) Mask() uint32 {
	argValue := gi.FieldGet(flagsClassStruct, recv.native, "mask")
	value := argValue.Uint32()
	return value
}

// NValues returns the C field 'n_values'.
func (recv *FlagsClass) NValues() uint32 {
	argValue := gi.FieldGet(flagsClassStruct, recv.native, "n_values")
	value := argValue.Uint32()
	return value
}

// Values returns the C field 'values'.
func (recv *FlagsClass) Values() *FlagsValue {
	argValue := gi.FieldGet(flagsClassStruct, recv.native, "values")
	value := &FlagsValue{native: argValue.Pointer()}
	return value
}

// FlagsClassStruct creates an uninitialised FlagsClass.
func FlagsClassStruct() *FlagsClass {
	err := flagsClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FlagsClass{native: flagsClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// Value returns the C field 'value'.
func (recv *FlagsValue) Value() uint32 {
	argValue := gi.FieldGet(flagsValueStruct, recv.native, "value")
	value := argValue.Uint32()
	return value
}

// ValueName returns the C field 'value_name'.
func (recv *FlagsValue) ValueName() string {
	argValue := gi.FieldGet(flagsValueStruct, recv.native, "value_name")
	value := argValue.String(false)
	return value
}

// ValueNick returns the C field 'value_nick'.
func (recv *FlagsValue) ValueNick() string {
	argValue := gi.FieldGet(flagsValueStruct, recv.native, "value_nick")
	value := argValue.String(false)
	return value
}

// FlagsValueStruct creates an uninitialised FlagsValue.
func FlagsValueStruct() *FlagsValue {
	err := flagsValueStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FlagsValue{native: flagsValueStruct.Alloc()}
	return structGo
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
	native uintptr
}

// GTypeClass returns the C field 'g_type_class'.
func (recv *InitiallyUnownedClass) GTypeClass() *TypeClass {
	argValue := gi.FieldGet(initiallyUnownedClassStruct, recv.native, "g_type_class")
	value := &TypeClass{native: argValue.Pointer()}
	return value
}

// UNSUPPORTED : C value 'constructor' : for field getter : missing Type

// UNSUPPORTED : C value 'set_property' : for field getter : missing Type

// UNSUPPORTED : C value 'get_property' : for field getter : missing Type

// UNSUPPORTED : C value 'dispose' : for field getter : missing Type

// UNSUPPORTED : C value 'finalize' : for field getter : missing Type

// UNSUPPORTED : C value 'dispatch_properties_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'notify' : for field getter : missing Type

// UNSUPPORTED : C value 'constructed' : for field getter : missing Type

// InitiallyUnownedClassStruct creates an uninitialised InitiallyUnownedClass.
func InitiallyUnownedClassStruct() *InitiallyUnownedClass {
	err := initiallyUnownedClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &InitiallyUnownedClass{native: initiallyUnownedClassStruct.Alloc()}
	return structGo
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
}

// UNSUPPORTED : C value 'interface_init' : for field getter : no Go type for 'InterfaceInitFunc'

// UNSUPPORTED : C value 'interface_finalize' : for field getter : no Go type for 'InterfaceFinalizeFunc'

// UNSUPPORTED : C value 'interface_data' : for field getter : no Go type for 'gpointer'

// InterfaceInfoStruct creates an uninitialised InterfaceInfo.
func InterfaceInfoStruct() *InterfaceInfo {
	err := interfaceInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &InterfaceInfo{native: interfaceInfoStruct.Alloc()}
	return structGo
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
	native uintptr
}

// GTypeClass returns the C field 'g_type_class'.
func (recv *ObjectClass) GTypeClass() *TypeClass {
	argValue := gi.FieldGet(objectClassStruct, recv.native, "g_type_class")
	value := &TypeClass{native: argValue.Pointer()}
	return value
}

// UNSUPPORTED : C value 'constructor' : for field getter : missing Type

// UNSUPPORTED : C value 'set_property' : for field getter : missing Type

// UNSUPPORTED : C value 'get_property' : for field getter : missing Type

// UNSUPPORTED : C value 'dispose' : for field getter : missing Type

// UNSUPPORTED : C value 'finalize' : for field getter : missing Type

// UNSUPPORTED : C value 'dispatch_properties_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'notify' : for field getter : missing Type

// UNSUPPORTED : C value 'constructed' : for field getter : missing Type

// ObjectClassStruct creates an uninitialised ObjectClass.
func ObjectClassStruct() *ObjectClass {
	err := objectClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ObjectClass{native: objectClassStruct.Alloc()}
	return structGo
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
func (recv *ObjectClass) ListProperties() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	err := objectClassListPropertiesFunction_Set()
	if err == nil {
		objectClassListPropertiesFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint32()

	return out0
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
func (recv *ObjectClass) OverrideProperty(propertyId uint32, name string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(propertyId)
	inArgs[2].SetString(name)

	err := objectClassOverridePropertyFunction_Set()
	if err == nil {
		objectClassOverridePropertyFunction.Invoke(inArgs[:], nil)
	}

	return
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
}

// UNSUPPORTED : C value 'pspec' : for field getter : no Go type for 'ParamSpec'

// Value returns the C field 'value'.
func (recv *ObjectConstructParam) Value() *Value {
	argValue := gi.FieldGet(objectConstructParamStruct, recv.native, "value")
	value := &Value{native: argValue.Pointer()}
	return value
}

// ObjectConstructParamStruct creates an uninitialised ObjectConstructParam.
func ObjectConstructParamStruct() *ObjectConstructParam {
	err := objectConstructParamStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ObjectConstructParam{native: objectConstructParamStruct.Alloc()}
	return structGo
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
	native uintptr
}

// GTypeClass returns the C field 'g_type_class'.
func (recv *ParamSpecClass) GTypeClass() *TypeClass {
	argValue := gi.FieldGet(paramSpecClassStruct, recv.native, "g_type_class")
	value := &TypeClass{native: argValue.Pointer()}
	return value
}

// UNSUPPORTED : C value 'value_type' : for field getter : no Go type for 'GType'

// UNSUPPORTED : C value 'finalize' : for field getter : missing Type

// UNSUPPORTED : C value 'value_set_default' : for field getter : missing Type

// UNSUPPORTED : C value 'value_validate' : for field getter : missing Type

// UNSUPPORTED : C value 'values_cmp' : for field getter : missing Type

// ParamSpecClassStruct creates an uninitialised ParamSpecClass.
func ParamSpecClassStruct() *ParamSpecClass {
	err := paramSpecClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecClass{native: paramSpecClassStruct.Alloc()}
	return structGo
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

// ParamSpecPoolStruct creates an uninitialised ParamSpecPool.
func ParamSpecPoolStruct() *ParamSpecPool {
	err := paramSpecPoolStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecPool{native: paramSpecPoolStruct.Alloc()}
	return structGo
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
	native uintptr
}

// InstanceSize returns the C field 'instance_size'.
func (recv *ParamSpecTypeInfo) InstanceSize() uint16 {
	argValue := gi.FieldGet(paramSpecTypeInfoStruct, recv.native, "instance_size")
	value := argValue.Uint16()
	return value
}

// NPreallocs returns the C field 'n_preallocs'.
func (recv *ParamSpecTypeInfo) NPreallocs() uint16 {
	argValue := gi.FieldGet(paramSpecTypeInfoStruct, recv.native, "n_preallocs")
	value := argValue.Uint16()
	return value
}

// UNSUPPORTED : C value 'instance_init' : for field getter : missing Type

// UNSUPPORTED : C value 'value_type' : for field getter : no Go type for 'GType'

// UNSUPPORTED : C value 'finalize' : for field getter : missing Type

// UNSUPPORTED : C value 'value_set_default' : for field getter : missing Type

// UNSUPPORTED : C value 'value_validate' : for field getter : missing Type

// UNSUPPORTED : C value 'values_cmp' : for field getter : missing Type

// ParamSpecTypeInfoStruct creates an uninitialised ParamSpecTypeInfo.
func ParamSpecTypeInfoStruct() *ParamSpecTypeInfo {
	err := paramSpecTypeInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecTypeInfo{native: paramSpecTypeInfoStruct.Alloc()}
	return structGo
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
}

// Name returns the C field 'name'.
func (recv *Parameter) Name() string {
	argValue := gi.FieldGet(parameterStruct, recv.native, "name")
	value := argValue.String(false)
	return value
}

// Value returns the C field 'value'.
func (recv *Parameter) Value() *Value {
	argValue := gi.FieldGet(parameterStruct, recv.native, "value")
	value := &Value{native: argValue.Pointer()}
	return value
}

// ParameterStruct creates an uninitialised Parameter.
func ParameterStruct() *Parameter {
	err := parameterStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Parameter{native: parameterStruct.Alloc()}
	return structGo
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
	native uintptr
}

// SignalId returns the C field 'signal_id'.
func (recv *SignalInvocationHint) SignalId() uint32 {
	argValue := gi.FieldGet(signalInvocationHintStruct, recv.native, "signal_id")
	value := argValue.Uint32()
	return value
}

// Detail returns the C field 'detail'.
func (recv *SignalInvocationHint) Detail() glib.Quark {
	argValue := gi.FieldGet(signalInvocationHintStruct, recv.native, "detail")
	value := glib.Quark(argValue.Uint32())
	return value
}

// UNSUPPORTED : C value 'run_type' : for field getter : no Go type for 'SignalFlags'

// SignalInvocationHintStruct creates an uninitialised SignalInvocationHint.
func SignalInvocationHintStruct() *SignalInvocationHint {
	err := signalInvocationHintStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SignalInvocationHint{native: signalInvocationHintStruct.Alloc()}
	return structGo
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

type SignalQuery_ struct {
	native uintptr
}

// SignalId returns the C field 'signal_id'.
func (recv *SignalQuery_) SignalId() uint32 {
	argValue := gi.FieldGet(signalQueryStruct, recv.native, "signal_id")
	value := argValue.Uint32()
	return value
}

// SignalName returns the C field 'signal_name'.
func (recv *SignalQuery_) SignalName() string {
	argValue := gi.FieldGet(signalQueryStruct, recv.native, "signal_name")
	value := argValue.String(false)
	return value
}

// UNSUPPORTED : C value 'itype' : for field getter : no Go type for 'GType'

// UNSUPPORTED : C value 'signal_flags' : for field getter : no Go type for 'SignalFlags'

// UNSUPPORTED : C value 'return_type' : for field getter : no Go type for 'GType'

// NParams returns the C field 'n_params'.
func (recv *SignalQuery_) NParams() uint32 {
	argValue := gi.FieldGet(signalQueryStruct, recv.native, "n_params")
	value := argValue.Uint32()
	return value
}

// UNSUPPORTED : C value 'param_types' : for field getter : missing Type

// SignalQuery_Struct creates an uninitialised SignalQuery_.
func SignalQuery_Struct() *SignalQuery_ {
	err := signalQueryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SignalQuery_{native: signalQueryStruct.Alloc()}
	return structGo
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

// TypeClassStruct creates an uninitialised TypeClass.
func TypeClassStruct() *TypeClass {
	err := typeClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TypeClass{native: typeClassStruct.Alloc()}
	return structGo
}

var typeClassAddPrivateFunction *gi.Function
var typeClassAddPrivateFunction_Once sync.Once

func typeClassAddPrivateFunction_Set() error {
	var err error
	typeClassAddPrivateFunction_Once.Do(func() {
		err = typeClassStruct_Set()
		if err != nil {
			return
		}
		typeClassAddPrivateFunction, err = typeClassStruct.InvokerNew("add_private")
	})
	return err
}

// AddPrivate is a representation of the C type g_type_class_add_private.
func (recv *TypeClass) AddPrivate(privateSize uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint64(privateSize)

	err := typeClassAddPrivateFunction_Set()
	if err == nil {
		typeClassAddPrivateFunction.Invoke(inArgs[:], nil)
	}

	return
}

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
func (recv *TypeClass) GetInstancePrivateOffset() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := typeClassGetInstancePrivateOffsetFunction_Set()
	if err == nil {
		ret = typeClassGetInstancePrivateOffsetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
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
func (recv *TypeClass) PeekParent() *TypeClass {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := typeClassPeekParentFunction_Set()
	if err == nil {
		ret = typeClassPeekParentFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TypeClass{native: ret.Pointer()}

	return retGo
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
func (recv *TypeClass) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := typeClassUnrefFunction_Set()
	if err == nil {
		typeClassUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *TypeClass) UnrefUncached() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := typeClassUnrefUncachedFunction_Set()
	if err == nil {
		typeClassUnrefUncachedFunction.Invoke(inArgs[:], nil)
	}

	return
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
}

// UNSUPPORTED : C value 'type_flags' : for field getter : no Go type for 'TypeFundamentalFlags'

// TypeFundamentalInfoStruct creates an uninitialised TypeFundamentalInfo.
func TypeFundamentalInfoStruct() *TypeFundamentalInfo {
	err := typeFundamentalInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TypeFundamentalInfo{native: typeFundamentalInfoStruct.Alloc()}
	return structGo
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
	native uintptr
}

// ClassSize returns the C field 'class_size'.
func (recv *TypeInfo) ClassSize() uint16 {
	argValue := gi.FieldGet(typeInfoStruct, recv.native, "class_size")
	value := argValue.Uint16()
	return value
}

// UNSUPPORTED : C value 'base_init' : for field getter : no Go type for 'BaseInitFunc'

// UNSUPPORTED : C value 'base_finalize' : for field getter : no Go type for 'BaseFinalizeFunc'

// UNSUPPORTED : C value 'class_init' : for field getter : no Go type for 'ClassInitFunc'

// UNSUPPORTED : C value 'class_finalize' : for field getter : no Go type for 'ClassFinalizeFunc'

// UNSUPPORTED : C value 'class_data' : for field getter : no Go type for 'gpointer'

// InstanceSize returns the C field 'instance_size'.
func (recv *TypeInfo) InstanceSize() uint16 {
	argValue := gi.FieldGet(typeInfoStruct, recv.native, "instance_size")
	value := argValue.Uint16()
	return value
}

// NPreallocs returns the C field 'n_preallocs'.
func (recv *TypeInfo) NPreallocs() uint16 {
	argValue := gi.FieldGet(typeInfoStruct, recv.native, "n_preallocs")
	value := argValue.Uint16()
	return value
}

// UNSUPPORTED : C value 'instance_init' : for field getter : no Go type for 'InstanceInitFunc'

// ValueTable returns the C field 'value_table'.
func (recv *TypeInfo) ValueTable() *TypeValueTable {
	argValue := gi.FieldGet(typeInfoStruct, recv.native, "value_table")
	value := &TypeValueTable{native: argValue.Pointer()}
	return value
}

// TypeInfoStruct creates an uninitialised TypeInfo.
func TypeInfoStruct() *TypeInfo {
	err := typeInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TypeInfo{native: typeInfoStruct.Alloc()}
	return structGo
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

// TypeInstanceStruct creates an uninitialised TypeInstance.
func TypeInstanceStruct() *TypeInstance {
	err := typeInstanceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TypeInstance{native: typeInstanceStruct.Alloc()}
	return structGo
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

// TypeInterfaceStruct creates an uninitialised TypeInterface.
func TypeInterfaceStruct() *TypeInterface {
	err := typeInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TypeInterface{native: typeInterfaceStruct.Alloc()}
	return structGo
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
func (recv *TypeInterface) PeekParent() *TypeInterface {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := typeInterfacePeekParentFunction_Set()
	if err == nil {
		ret = typeInterfacePeekParentFunction.Invoke(inArgs[:], nil)
	}

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
	native uintptr
}

// ParentClass returns the C field 'parent_class'.
func (recv *TypeModuleClass) ParentClass() *ObjectClass {
	argValue := gi.FieldGet(typeModuleClassStruct, recv.native, "parent_class")
	value := &ObjectClass{native: argValue.Pointer()}
	return value
}

// UNSUPPORTED : C value 'load' : for field getter : missing Type

// UNSUPPORTED : C value 'unload' : for field getter : missing Type

// UNSUPPORTED : C value 'reserved1' : for field getter : missing Type

// UNSUPPORTED : C value 'reserved2' : for field getter : missing Type

// UNSUPPORTED : C value 'reserved3' : for field getter : missing Type

// UNSUPPORTED : C value 'reserved4' : for field getter : missing Type

// TypeModuleClassStruct creates an uninitialised TypeModuleClass.
func TypeModuleClassStruct() *TypeModuleClass {
	err := typeModuleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TypeModuleClass{native: typeModuleClassStruct.Alloc()}
	return structGo
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
}

// UNSUPPORTED : C value 'use_plugin' : for field getter : no Go type for 'TypePluginUse'

// UNSUPPORTED : C value 'unuse_plugin' : for field getter : no Go type for 'TypePluginUnuse'

// UNSUPPORTED : C value 'complete_type_info' : for field getter : no Go type for 'TypePluginCompleteTypeInfo'

// UNSUPPORTED : C value 'complete_interface_info' : for field getter : no Go type for 'TypePluginCompleteInterfaceInfo'

// TypePluginClassStruct creates an uninitialised TypePluginClass.
func TypePluginClassStruct() *TypePluginClass {
	err := typePluginClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TypePluginClass{native: typePluginClassStruct.Alloc()}
	return structGo
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'GType'

// TypeName returns the C field 'type_name'.
func (recv *TypeQuery) TypeName() string {
	argValue := gi.FieldGet(typeQueryStruct, recv.native, "type_name")
	value := argValue.String(false)
	return value
}

// ClassSize returns the C field 'class_size'.
func (recv *TypeQuery) ClassSize() uint32 {
	argValue := gi.FieldGet(typeQueryStruct, recv.native, "class_size")
	value := argValue.Uint32()
	return value
}

// InstanceSize returns the C field 'instance_size'.
func (recv *TypeQuery) InstanceSize() uint32 {
	argValue := gi.FieldGet(typeQueryStruct, recv.native, "instance_size")
	value := argValue.Uint32()
	return value
}

// TypeQueryStruct creates an uninitialised TypeQuery.
func TypeQueryStruct() *TypeQuery {
	err := typeQueryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TypeQuery{native: typeQueryStruct.Alloc()}
	return structGo
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
}

// UNSUPPORTED : C value 'value_init' : for field getter : missing Type

// UNSUPPORTED : C value 'value_free' : for field getter : missing Type

// UNSUPPORTED : C value 'value_copy' : for field getter : missing Type

// UNSUPPORTED : C value 'value_peek_pointer' : for field getter : missing Type

// CollectFormat returns the C field 'collect_format'.
func (recv *TypeValueTable) CollectFormat() string {
	argValue := gi.FieldGet(typeValueTableStruct, recv.native, "collect_format")
	value := argValue.String(false)
	return value
}

// UNSUPPORTED : C value 'collect_value' : for field getter : missing Type

// LcopyFormat returns the C field 'lcopy_format'.
func (recv *TypeValueTable) LcopyFormat() string {
	argValue := gi.FieldGet(typeValueTableStruct, recv.native, "lcopy_format")
	value := argValue.String(false)
	return value
}

// UNSUPPORTED : C value 'lcopy_value' : for field getter : missing Type

// TypeValueTableStruct creates an uninitialised TypeValueTable.
func TypeValueTableStruct() *TypeValueTable {
	err := typeValueTableStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TypeValueTable{native: typeValueTableStruct.Alloc()}
	return structGo
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
}

// UNSUPPORTED : C value 'data' : for field getter : missing Type

// ValueStruct creates an uninitialised Value.
func ValueStruct() *Value {
	err := valueStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Value{native: valueStruct.Alloc()}
	return structGo
}

var valueCopyFunction *gi.Function
var valueCopyFunction_Once sync.Once

func valueCopyFunction_Set() error {
	var err error
	valueCopyFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueCopyFunction, err = valueStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type g_value_copy.
func (recv *Value) Copy(destValue *Value) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(destValue.native)

	err := valueCopyFunction_Set()
	if err == nil {
		valueCopyFunction.Invoke(inArgs[:], nil)
	}

	return
}

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
func (recv *Value) DupString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueDupStringFunction_Set()
	if err == nil {
		ret = valueDupStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func (recv *Value) FitsPointer() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueFitsPointerFunction_Set()
	if err == nil {
		ret = valueFitsPointerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func (recv *Value) GetBoolean() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetBooleanFunction_Set()
	if err == nil {
		ret = valueGetBooleanFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func (recv *Value) GetChar() int8 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetCharFunction_Set()
	if err == nil {
		ret = valueGetCharFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int8()

	return retGo
}

var valueGetDoubleFunction *gi.Function
var valueGetDoubleFunction_Once sync.Once

func valueGetDoubleFunction_Set() error {
	var err error
	valueGetDoubleFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueGetDoubleFunction, err = valueStruct.InvokerNew("get_double")
	})
	return err
}

// GetDouble is a representation of the C type g_value_get_double.
func (recv *Value) GetDouble() float64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetDoubleFunction_Set()
	if err == nil {
		ret = valueGetDoubleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo
}

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
func (recv *Value) GetEnum() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetEnumFunction_Set()
	if err == nil {
		ret = valueGetEnumFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
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
func (recv *Value) GetFlags() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetFlagsFunction_Set()
	if err == nil {
		ret = valueGetFlagsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var valueGetFloatFunction *gi.Function
var valueGetFloatFunction_Once sync.Once

func valueGetFloatFunction_Set() error {
	var err error
	valueGetFloatFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueGetFloatFunction, err = valueStruct.InvokerNew("get_float")
	})
	return err
}

// GetFloat is a representation of the C type g_value_get_float.
func (recv *Value) GetFloat() float32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetFloatFunction_Set()
	if err == nil {
		ret = valueGetFloatFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float32()

	return retGo
}

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
func (recv *Value) GetInt() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetIntFunction_Set()
	if err == nil {
		ret = valueGetIntFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
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
func (recv *Value) GetInt64() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetInt64Function_Set()
	if err == nil {
		ret = valueGetInt64Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
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
func (recv *Value) GetLong() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetLongFunction_Set()
	if err == nil {
		ret = valueGetLongFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
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
func (recv *Value) GetSchar() int8 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetScharFunction_Set()
	if err == nil {
		ret = valueGetScharFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int8()

	return retGo
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
func (recv *Value) GetString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetStringFunction_Set()
	if err == nil {
		ret = valueGetStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
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
func (recv *Value) GetUchar() uint8 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetUcharFunction_Set()
	if err == nil {
		ret = valueGetUcharFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint8()

	return retGo
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
func (recv *Value) GetUint() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetUintFunction_Set()
	if err == nil {
		ret = valueGetUintFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
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
func (recv *Value) GetUint64() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetUint64Function_Set()
	if err == nil {
		ret = valueGetUint64Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
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
func (recv *Value) GetUlong() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueGetUlongFunction_Set()
	if err == nil {
		ret = valueGetUlongFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

// UNSUPPORTED : C value 'g_value_get_variant' : return type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_value_init' : parameter 'g_type' of type 'GType' not supported

var valueInitFromInstanceFunction *gi.Function
var valueInitFromInstanceFunction_Once sync.Once

func valueInitFromInstanceFunction_Set() error {
	var err error
	valueInitFromInstanceFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueInitFromInstanceFunction, err = valueStruct.InvokerNew("init_from_instance")
	})
	return err
}

// InitFromInstance is a representation of the C type g_value_init_from_instance.
func (recv *Value) InitFromInstance(instance *TypeInstance) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(instance.native)

	err := valueInitFromInstanceFunction_Set()
	if err == nil {
		valueInitFromInstanceFunction.Invoke(inArgs[:], nil)
	}

	return
}

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
func (recv *Value) Reset() *Value {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueResetFunction_Set()
	if err == nil {
		ret = valueResetFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Value{native: ret.Pointer()}

	return retGo
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
func (recv *Value) SetBoolean(vBoolean bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(vBoolean)

	err := valueSetBooleanFunction_Set()
	if err == nil {
		valueSetBooleanFunction.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *Value) SetChar(vChar int8) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt8(vChar)

	err := valueSetCharFunction_Set()
	if err == nil {
		valueSetCharFunction.Invoke(inArgs[:], nil)
	}

	return
}

var valueSetDoubleFunction *gi.Function
var valueSetDoubleFunction_Once sync.Once

func valueSetDoubleFunction_Set() error {
	var err error
	valueSetDoubleFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueSetDoubleFunction, err = valueStruct.InvokerNew("set_double")
	})
	return err
}

// SetDouble is a representation of the C type g_value_set_double.
func (recv *Value) SetDouble(vDouble float64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetFloat64(vDouble)

	err := valueSetDoubleFunction_Set()
	if err == nil {
		valueSetDoubleFunction.Invoke(inArgs[:], nil)
	}

	return
}

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
func (recv *Value) SetEnum(vEnum int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(vEnum)

	err := valueSetEnumFunction_Set()
	if err == nil {
		valueSetEnumFunction.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *Value) SetFlags(vFlags uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(vFlags)

	err := valueSetFlagsFunction_Set()
	if err == nil {
		valueSetFlagsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var valueSetFloatFunction *gi.Function
var valueSetFloatFunction_Once sync.Once

func valueSetFloatFunction_Set() error {
	var err error
	valueSetFloatFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueSetFloatFunction, err = valueStruct.InvokerNew("set_float")
	})
	return err
}

// SetFloat is a representation of the C type g_value_set_float.
func (recv *Value) SetFloat(vFloat float32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetFloat32(vFloat)

	err := valueSetFloatFunction_Set()
	if err == nil {
		valueSetFloatFunction.Invoke(inArgs[:], nil)
	}

	return
}

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
func (recv *Value) SetInt(vInt int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(vInt)

	err := valueSetIntFunction_Set()
	if err == nil {
		valueSetIntFunction.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *Value) SetInt64(vInt64 int64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(vInt64)

	err := valueSetInt64Function_Set()
	if err == nil {
		valueSetInt64Function.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *Value) SetLong(vLong int64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(vLong)

	err := valueSetLongFunction_Set()
	if err == nil {
		valueSetLongFunction.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *Value) SetSchar(vChar int8) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt8(vChar)

	err := valueSetScharFunction_Set()
	if err == nil {
		valueSetScharFunction.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *Value) SetStaticString(vString string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(vString)

	err := valueSetStaticStringFunction_Set()
	if err == nil {
		valueSetStaticStringFunction.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *Value) SetString(vString string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(vString)

	err := valueSetStringFunction_Set()
	if err == nil {
		valueSetStringFunction.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *Value) SetStringTakeOwnership(vString string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(vString)

	err := valueSetStringTakeOwnershipFunction_Set()
	if err == nil {
		valueSetStringTakeOwnershipFunction.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *Value) SetUchar(vUchar uint8) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint8(vUchar)

	err := valueSetUcharFunction_Set()
	if err == nil {
		valueSetUcharFunction.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *Value) SetUint(vUint uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(vUint)

	err := valueSetUintFunction_Set()
	if err == nil {
		valueSetUintFunction.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *Value) SetUint64(vUint64 uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint64(vUint64)

	err := valueSetUint64Function_Set()
	if err == nil {
		valueSetUint64Function.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *Value) SetUlong(vUlong uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint64(vUlong)

	err := valueSetUlongFunction_Set()
	if err == nil {
		valueSetUlongFunction.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *Value) TakeString(vString string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(vString)

	err := valueTakeStringFunction_Set()
	if err == nil {
		valueTakeStringFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_value_take_variant' : parameter 'variant' of type 'GLib.Variant' not supported

var valueTransformFunction *gi.Function
var valueTransformFunction_Once sync.Once

func valueTransformFunction_Set() error {
	var err error
	valueTransformFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueTransformFunction, err = valueStruct.InvokerNew("transform")
	})
	return err
}

// Transform is a representation of the C type g_value_transform.
func (recv *Value) Transform(destValue *Value) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(destValue.native)

	var ret gi.Argument

	err := valueTransformFunction_Set()
	if err == nil {
		ret = valueTransformFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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
func (recv *Value) Unset() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := valueUnsetFunction_Set()
	if err == nil {
		valueUnsetFunction.Invoke(inArgs[:], nil)
	}

	return
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
	native uintptr
}

// NValues returns the C field 'n_values'.
func (recv *ValueArray) NValues() uint32 {
	argValue := gi.FieldGet(valueArrayStruct, recv.native, "n_values")
	value := argValue.Uint32()
	return value
}

// Values returns the C field 'values'.
func (recv *ValueArray) Values() *Value {
	argValue := gi.FieldGet(valueArrayStruct, recv.native, "values")
	value := &Value{native: argValue.Pointer()}
	return value
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
func ValueArrayNew(nPrealloced uint32) *ValueArray {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(nPrealloced)

	var ret gi.Argument

	err := valueArrayNewFunction_Set()
	if err == nil {
		ret = valueArrayNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &ValueArray{native: ret.Pointer()}

	return retGo
}

var valueArrayAppendFunction *gi.Function
var valueArrayAppendFunction_Once sync.Once

func valueArrayAppendFunction_Set() error {
	var err error
	valueArrayAppendFunction_Once.Do(func() {
		err = valueArrayStruct_Set()
		if err != nil {
			return
		}
		valueArrayAppendFunction, err = valueArrayStruct.InvokerNew("append")
	})
	return err
}

// Append is a representation of the C type g_value_array_append.
func (recv *ValueArray) Append(value *Value) *ValueArray {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(value.native)

	var ret gi.Argument

	err := valueArrayAppendFunction_Set()
	if err == nil {
		ret = valueArrayAppendFunction.Invoke(inArgs[:], nil)
	}

	retGo := &ValueArray{native: ret.Pointer()}

	return retGo
}

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
func (recv *ValueArray) Copy() *ValueArray {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueArrayCopyFunction_Set()
	if err == nil {
		ret = valueArrayCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &ValueArray{native: ret.Pointer()}

	return retGo
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
func (recv *ValueArray) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := valueArrayFreeFunction_Set()
	if err == nil {
		valueArrayFreeFunction.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *ValueArray) GetNth(index uint32) *Value {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(index)

	var ret gi.Argument

	err := valueArrayGetNthFunction_Set()
	if err == nil {
		ret = valueArrayGetNthFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Value{native: ret.Pointer()}

	return retGo
}

var valueArrayInsertFunction *gi.Function
var valueArrayInsertFunction_Once sync.Once

func valueArrayInsertFunction_Set() error {
	var err error
	valueArrayInsertFunction_Once.Do(func() {
		err = valueArrayStruct_Set()
		if err != nil {
			return
		}
		valueArrayInsertFunction, err = valueArrayStruct.InvokerNew("insert")
	})
	return err
}

// Insert is a representation of the C type g_value_array_insert.
func (recv *ValueArray) Insert(index uint32, value *Value) *ValueArray {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(index)
	inArgs[2].SetPointer(value.native)

	var ret gi.Argument

	err := valueArrayInsertFunction_Set()
	if err == nil {
		ret = valueArrayInsertFunction.Invoke(inArgs[:], nil)
	}

	retGo := &ValueArray{native: ret.Pointer()}

	return retGo
}

var valueArrayPrependFunction *gi.Function
var valueArrayPrependFunction_Once sync.Once

func valueArrayPrependFunction_Set() error {
	var err error
	valueArrayPrependFunction_Once.Do(func() {
		err = valueArrayStruct_Set()
		if err != nil {
			return
		}
		valueArrayPrependFunction, err = valueArrayStruct.InvokerNew("prepend")
	})
	return err
}

// Prepend is a representation of the C type g_value_array_prepend.
func (recv *ValueArray) Prepend(value *Value) *ValueArray {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(value.native)

	var ret gi.Argument

	err := valueArrayPrependFunction_Set()
	if err == nil {
		ret = valueArrayPrependFunction.Invoke(inArgs[:], nil)
	}

	retGo := &ValueArray{native: ret.Pointer()}

	return retGo
}

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
func (recv *ValueArray) Remove(index uint32) *ValueArray {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(index)

	var ret gi.Argument

	err := valueArrayRemoveFunction_Set()
	if err == nil {
		ret = valueArrayRemoveFunction.Invoke(inArgs[:], nil)
	}

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

// WeakRefStruct creates an uninitialised WeakRef.
func WeakRefStruct() *WeakRef {
	err := weakRefStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WeakRef{native: weakRefStruct.Alloc()}
	return structGo
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
func (recv *WeakRef) Clear() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := weakRefClearFunction_Set()
	if err == nil {
		weakRefClearFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_weak_ref_get' : return type 'Object' not supported

// UNSUPPORTED : C value 'g_weak_ref_init' : parameter 'object' of type 'Object' not supported

// UNSUPPORTED : C value 'g_weak_ref_set' : parameter 'object' of type 'Object' not supported
