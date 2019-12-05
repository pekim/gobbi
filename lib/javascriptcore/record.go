// Code generated - DO NOT EDIT.

package javascriptcore

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"runtime"
	"sync"
)

var globalContextRefStruct *gi.Struct
var globalContextRefStruct_Once sync.Once

func globalContextRefStruct_Set() error {
	var err error
	globalContextRefStruct_Once.Do(func() {
		globalContextRefStruct, err = gi.StructNew("JavaScriptCore", "GlobalContextRef")
	})
	return err
}

type GlobalContextRef struct {
	native uintptr
}

func GlobalContextRefNewFromNative(native uintptr) *GlobalContextRef {
	return &GlobalContextRef{native: native}
}

var globalContextRefRefFunction *gi.Function
var globalContextRefRefFunction_Once sync.Once

func globalContextRefRefFunction_Set() error {
	var err error
	globalContextRefRefFunction_Once.Do(func() {
		err = globalContextRefStruct_Set()
		if err != nil {
			return
		}
		globalContextRefRefFunction, err = globalContextRefStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type JSGlobalContextRetain.
func (recv *GlobalContextRef) Ref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := globalContextRefRefFunction_Set()
	if err == nil {
		globalContextRefRefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var globalContextRefUnrefFunction *gi.Function
var globalContextRefUnrefFunction_Once sync.Once

func globalContextRefUnrefFunction_Set() error {
	var err error
	globalContextRefUnrefFunction_Once.Do(func() {
		err = globalContextRefStruct_Set()
		if err != nil {
			return
		}
		globalContextRefUnrefFunction, err = globalContextRefStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type JSGlobalContextRelease.
func (recv *GlobalContextRef) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := globalContextRefUnrefFunction_Set()
	if err == nil {
		globalContextRefUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

// GlobalContextRefStruct creates an uninitialised GlobalContextRef.
func GlobalContextRefStruct() *GlobalContextRef {
	err := globalContextRefStruct_Set()
	if err != nil {
		return nil
	}

	structGo := GlobalContextRefNewFromNative(globalContextRefStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeGlobalContextRef)
	return structGo
}
func finalizeGlobalContextRef(obj *GlobalContextRef) {
	globalContextRefStruct.Free(obj.native)
}

var valueRefStruct *gi.Struct
var valueRefStruct_Once sync.Once

func valueRefStruct_Set() error {
	var err error
	valueRefStruct_Once.Do(func() {
		valueRefStruct, err = gi.StructNew("JavaScriptCore", "ValueRef")
	})
	return err
}

type ValueRef struct {
	native uintptr
}

func ValueRefNewFromNative(native uintptr) *ValueRef {
	return &ValueRef{native: native}
}

// ValueRefStruct creates an uninitialised ValueRef.
func ValueRefStruct() *ValueRef {
	err := valueRefStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ValueRefNewFromNative(valueRefStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeValueRef)
	return structGo
}
func finalizeValueRef(obj *ValueRef) {
	valueRefStruct.Free(obj.native)
}

var stringRefStruct *gi.Struct
var stringRefStruct_Once sync.Once

func stringRefStruct_Set() error {
	var err error
	stringRefStruct_Once.Do(func() {
		stringRefStruct, err = gi.StructNew("JavaScriptCore", "StringRef")
	})
	return err
}

type StringRef struct {
	native uintptr
}

func StringRefNewFromNative(native uintptr) *StringRef {
	return &StringRef{native: native}
}

var stringRefRefFunction *gi.Function
var stringRefRefFunction_Once sync.Once

func stringRefRefFunction_Set() error {
	var err error
	stringRefRefFunction_Once.Do(func() {
		err = stringRefStruct_Set()
		if err != nil {
			return
		}
		stringRefRefFunction, err = stringRefStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type JSStringRetain.
func (recv *StringRef) Ref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := stringRefRefFunction_Set()
	if err == nil {
		stringRefRefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var stringRefUnrefFunction *gi.Function
var stringRefUnrefFunction_Once sync.Once

func stringRefUnrefFunction_Set() error {
	var err error
	stringRefUnrefFunction_Once.Do(func() {
		err = stringRefStruct_Set()
		if err != nil {
			return
		}
		stringRefUnrefFunction, err = stringRefStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type JSStringRelease.
func (recv *StringRef) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := stringRefUnrefFunction_Set()
	if err == nil {
		stringRefUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var stringRefGetMaximumUTF8CStringSizeFunction *gi.Function
var stringRefGetMaximumUTF8CStringSizeFunction_Once sync.Once

func stringRefGetMaximumUTF8CStringSizeFunction_Set() error {
	var err error
	stringRefGetMaximumUTF8CStringSizeFunction_Once.Do(func() {
		err = stringRefStruct_Set()
		if err != nil {
			return
		}
		stringRefGetMaximumUTF8CStringSizeFunction, err = stringRefStruct.InvokerNew("GetMaximumUTF8CStringSize")
	})
	return err
}

// GetMaximumUTF8CStringSize is a representation of the C type JSStringGetMaximumUTF8CStringSize.
func (recv *StringRef) GetMaximumUTF8CStringSize() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := stringRefGetMaximumUTF8CStringSizeFunction_Set()
	if err == nil {
		ret = stringRefGetMaximumUTF8CStringSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

var stringRefGetUTF8CStringJSStringGetUTF8CStringFunction *gi.Function
var stringRefGetUTF8CStringJSStringGetUTF8CStringFunction_Once sync.Once

func stringRefGetUTF8CStringJSStringGetUTF8CStringFunction_Set() error {
	var err error
	stringRefGetUTF8CStringJSStringGetUTF8CStringFunction_Once.Do(func() {
		err = stringRefStruct_Set()
		if err != nil {
			return
		}
		stringRefGetUTF8CStringJSStringGetUTF8CStringFunction, err = stringRefStruct.InvokerNew("GetUTF8CStringJSStringGetUTF8CString")
	})
	return err
}

// GetUTF8CStringJSStringGetUTF8CString is a representation of the C type JSStringGetUTF8CString.
func (recv *StringRef) GetUTF8CStringJSStringGetUTF8CString(buffer string, bufferSize uint64) (uint64, string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(buffer)
	inArgs[2].SetUint64(bufferSize)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := stringRefGetUTF8CStringJSStringGetUTF8CStringFunction_Set()
	if err == nil {
		ret = stringRefGetUTF8CStringJSStringGetUTF8CStringFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Uint64()
	out0 := outArgs[0].String(false)

	return retGo, out0
}

// StringRefStruct creates an uninitialised StringRef.
func StringRefStruct() *StringRef {
	err := stringRefStruct_Set()
	if err != nil {
		return nil
	}

	structGo := StringRefNewFromNative(stringRefStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeStringRef)
	return structGo
}
func finalizeStringRef(obj *StringRef) {
	stringRefStruct.Free(obj.native)
}

var classClassStruct *gi.Struct
var classClassStruct_Once sync.Once

func classClassStruct_Set() error {
	var err error
	classClassStruct_Once.Do(func() {
		classClassStruct, err = gi.StructNew("JavaScriptCore", "ClassClass")
	})
	return err
}

type ClassClass struct {
	native uintptr
}

func ClassClassNewFromNative(native uintptr) *ClassClass {
	return &ClassClass{native: native}
}

// ClassClassStruct creates an uninitialised ClassClass.
func ClassClassStruct() *ClassClass {
	err := classClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ClassClassNewFromNative(classClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeClassClass)
	return structGo
}
func finalizeClassClass(obj *ClassClass) {
	classClassStruct.Free(obj.native)
}

var classVTableStruct *gi.Struct
var classVTableStruct_Once sync.Once

func classVTableStruct_Set() error {
	var err error
	classVTableStruct_Once.Do(func() {
		classVTableStruct, err = gi.StructNew("JavaScriptCore", "ClassVTable")
	})
	return err
}

type ClassVTable struct {
	native uintptr
}

func ClassVTableNewFromNative(native uintptr) *ClassVTable {
	return &ClassVTable{native: native}
}

// UNSUPPORTED : C value 'get_property' : for field getter : no Go type for 'ClassGetPropertyFunction'

// UNSUPPORTED : C value 'get_property' : for field setter : no Go type for 'ClassGetPropertyFunction'

// UNSUPPORTED : C value 'set_property' : for field getter : no Go type for 'ClassSetPropertyFunction'

// UNSUPPORTED : C value 'set_property' : for field setter : no Go type for 'ClassSetPropertyFunction'

// UNSUPPORTED : C value 'has_property' : for field getter : no Go type for 'ClassHasPropertyFunction'

// UNSUPPORTED : C value 'has_property' : for field setter : no Go type for 'ClassHasPropertyFunction'

// UNSUPPORTED : C value 'delete_property' : for field getter : no Go type for 'ClassDeletePropertyFunction'

// UNSUPPORTED : C value 'delete_property' : for field setter : no Go type for 'ClassDeletePropertyFunction'

// UNSUPPORTED : C value 'enumerate_properties' : for field getter : no Go type for 'ClassEnumeratePropertiesFunction'

// UNSUPPORTED : C value 'enumerate_properties' : for field setter : no Go type for 'ClassEnumeratePropertiesFunction'

// UNSUPPORTED : C value '_jsc_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_jsc_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_jsc_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_jsc_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved3' : for field setter : missing Type

// ClassVTableStruct creates an uninitialised ClassVTable.
func ClassVTableStruct() *ClassVTable {
	err := classVTableStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ClassVTableNewFromNative(classVTableStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeClassVTable)
	return structGo
}
func finalizeClassVTable(obj *ClassVTable) {
	classVTableStruct.Free(obj.native)
}

var contextClassStruct *gi.Struct
var contextClassStruct_Once sync.Once

func contextClassStruct_Set() error {
	var err error
	contextClassStruct_Once.Do(func() {
		contextClassStruct, err = gi.StructNew("JavaScriptCore", "ContextClass")
	})
	return err
}

type ContextClass struct {
	native uintptr
}

func ContextClassNewFromNative(native uintptr) *ContextClass {
	return &ContextClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_jsc_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_jsc_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_jsc_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_jsc_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved3' : for field setter : missing Type

// ContextClassStruct creates an uninitialised ContextClass.
func ContextClassStruct() *ContextClass {
	err := contextClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ContextClassNewFromNative(contextClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeContextClass)
	return structGo
}
func finalizeContextClass(obj *ContextClass) {
	contextClassStruct.Free(obj.native)
}

var contextPrivateStruct *gi.Struct
var contextPrivateStruct_Once sync.Once

func contextPrivateStruct_Set() error {
	var err error
	contextPrivateStruct_Once.Do(func() {
		contextPrivateStruct, err = gi.StructNew("JavaScriptCore", "ContextPrivate")
	})
	return err
}

type ContextPrivate struct {
	native uintptr
}

func ContextPrivateNewFromNative(native uintptr) *ContextPrivate {
	return &ContextPrivate{native: native}
}

// ContextPrivateStruct creates an uninitialised ContextPrivate.
func ContextPrivateStruct() *ContextPrivate {
	err := contextPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ContextPrivateNewFromNative(contextPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeContextPrivate)
	return structGo
}
func finalizeContextPrivate(obj *ContextPrivate) {
	contextPrivateStruct.Free(obj.native)
}

var exceptionClassStruct *gi.Struct
var exceptionClassStruct_Once sync.Once

func exceptionClassStruct_Set() error {
	var err error
	exceptionClassStruct_Once.Do(func() {
		exceptionClassStruct, err = gi.StructNew("JavaScriptCore", "ExceptionClass")
	})
	return err
}

type ExceptionClass struct {
	native uintptr
}

func ExceptionClassNewFromNative(native uintptr) *ExceptionClass {
	return &ExceptionClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_jsc_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_jsc_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_jsc_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_jsc_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved3' : for field setter : missing Type

// ExceptionClassStruct creates an uninitialised ExceptionClass.
func ExceptionClassStruct() *ExceptionClass {
	err := exceptionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ExceptionClassNewFromNative(exceptionClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeExceptionClass)
	return structGo
}
func finalizeExceptionClass(obj *ExceptionClass) {
	exceptionClassStruct.Free(obj.native)
}

var exceptionPrivateStruct *gi.Struct
var exceptionPrivateStruct_Once sync.Once

func exceptionPrivateStruct_Set() error {
	var err error
	exceptionPrivateStruct_Once.Do(func() {
		exceptionPrivateStruct, err = gi.StructNew("JavaScriptCore", "ExceptionPrivate")
	})
	return err
}

type ExceptionPrivate struct {
	native uintptr
}

func ExceptionPrivateNewFromNative(native uintptr) *ExceptionPrivate {
	return &ExceptionPrivate{native: native}
}

// ExceptionPrivateStruct creates an uninitialised ExceptionPrivate.
func ExceptionPrivateStruct() *ExceptionPrivate {
	err := exceptionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ExceptionPrivateNewFromNative(exceptionPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeExceptionPrivate)
	return structGo
}
func finalizeExceptionPrivate(obj *ExceptionPrivate) {
	exceptionPrivateStruct.Free(obj.native)
}

var valueClassStruct *gi.Struct
var valueClassStruct_Once sync.Once

func valueClassStruct_Set() error {
	var err error
	valueClassStruct_Once.Do(func() {
		valueClassStruct, err = gi.StructNew("JavaScriptCore", "ValueClass")
	})
	return err
}

type ValueClass struct {
	native uintptr
}

func ValueClassNewFromNative(native uintptr) *ValueClass {
	return &ValueClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_jsc_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_jsc_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_jsc_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_jsc_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved3' : for field setter : missing Type

// ValueClassStruct creates an uninitialised ValueClass.
func ValueClassStruct() *ValueClass {
	err := valueClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ValueClassNewFromNative(valueClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeValueClass)
	return structGo
}
func finalizeValueClass(obj *ValueClass) {
	valueClassStruct.Free(obj.native)
}

var valuePrivateStruct *gi.Struct
var valuePrivateStruct_Once sync.Once

func valuePrivateStruct_Set() error {
	var err error
	valuePrivateStruct_Once.Do(func() {
		valuePrivateStruct, err = gi.StructNew("JavaScriptCore", "ValuePrivate")
	})
	return err
}

type ValuePrivate struct {
	native uintptr
}

func ValuePrivateNewFromNative(native uintptr) *ValuePrivate {
	return &ValuePrivate{native: native}
}

// ValuePrivateStruct creates an uninitialised ValuePrivate.
func ValuePrivateStruct() *ValuePrivate {
	err := valuePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ValuePrivateNewFromNative(valuePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeValuePrivate)
	return structGo
}
func finalizeValuePrivate(obj *ValuePrivate) {
	valuePrivateStruct.Free(obj.native)
}

var virtualMachineClassStruct *gi.Struct
var virtualMachineClassStruct_Once sync.Once

func virtualMachineClassStruct_Set() error {
	var err error
	virtualMachineClassStruct_Once.Do(func() {
		virtualMachineClassStruct, err = gi.StructNew("JavaScriptCore", "VirtualMachineClass")
	})
	return err
}

type VirtualMachineClass struct {
	native uintptr
}

func VirtualMachineClassNewFromNative(native uintptr) *VirtualMachineClass {
	return &VirtualMachineClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_jsc_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_jsc_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_jsc_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_jsc_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved3' : for field setter : missing Type

// VirtualMachineClassStruct creates an uninitialised VirtualMachineClass.
func VirtualMachineClassStruct() *VirtualMachineClass {
	err := virtualMachineClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := VirtualMachineClassNewFromNative(virtualMachineClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeVirtualMachineClass)
	return structGo
}
func finalizeVirtualMachineClass(obj *VirtualMachineClass) {
	virtualMachineClassStruct.Free(obj.native)
}

var virtualMachinePrivateStruct *gi.Struct
var virtualMachinePrivateStruct_Once sync.Once

func virtualMachinePrivateStruct_Set() error {
	var err error
	virtualMachinePrivateStruct_Once.Do(func() {
		virtualMachinePrivateStruct, err = gi.StructNew("JavaScriptCore", "VirtualMachinePrivate")
	})
	return err
}

type VirtualMachinePrivate struct {
	native uintptr
}

func VirtualMachinePrivateNewFromNative(native uintptr) *VirtualMachinePrivate {
	return &VirtualMachinePrivate{native: native}
}

// VirtualMachinePrivateStruct creates an uninitialised VirtualMachinePrivate.
func VirtualMachinePrivateStruct() *VirtualMachinePrivate {
	err := virtualMachinePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := VirtualMachinePrivateNewFromNative(virtualMachinePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeVirtualMachinePrivate)
	return structGo
}
func finalizeVirtualMachinePrivate(obj *VirtualMachinePrivate) {
	virtualMachinePrivateStruct.Free(obj.native)
}

var weakValueClassStruct *gi.Struct
var weakValueClassStruct_Once sync.Once

func weakValueClassStruct_Set() error {
	var err error
	weakValueClassStruct_Once.Do(func() {
		weakValueClassStruct, err = gi.StructNew("JavaScriptCore", "WeakValueClass")
	})
	return err
}

type WeakValueClass struct {
	native uintptr
}

func WeakValueClassNewFromNative(native uintptr) *WeakValueClass {
	return &WeakValueClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_jsc_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_jsc_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_jsc_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_jsc_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_jsc_reserved3' : for field setter : missing Type

// WeakValueClassStruct creates an uninitialised WeakValueClass.
func WeakValueClassStruct() *WeakValueClass {
	err := weakValueClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WeakValueClassNewFromNative(weakValueClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWeakValueClass)
	return structGo
}
func finalizeWeakValueClass(obj *WeakValueClass) {
	weakValueClassStruct.Free(obj.native)
}

var weakValuePrivateStruct *gi.Struct
var weakValuePrivateStruct_Once sync.Once

func weakValuePrivateStruct_Set() error {
	var err error
	weakValuePrivateStruct_Once.Do(func() {
		weakValuePrivateStruct, err = gi.StructNew("JavaScriptCore", "WeakValuePrivate")
	})
	return err
}

type WeakValuePrivate struct {
	native uintptr
}

func WeakValuePrivateNewFromNative(native uintptr) *WeakValuePrivate {
	return &WeakValuePrivate{native: native}
}

// WeakValuePrivateStruct creates an uninitialised WeakValuePrivate.
func WeakValuePrivateStruct() *WeakValuePrivate {
	err := weakValuePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WeakValuePrivateNewFromNative(weakValuePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWeakValuePrivate)
	return structGo
}
func finalizeWeakValuePrivate(obj *WeakValuePrivate) {
	weakValuePrivateStruct.Free(obj.native)
}
