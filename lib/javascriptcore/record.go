// Code generated - DO NOT EDIT.

package javascriptcore

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var globalContextRefStruct *gi.Struct
var globalContextRefStruct_Once sync.Once

func globalContextRefStruct_Set() {
	globalContextRefStruct_Once.Do(func() {
		globalContextRefStruct = gi.StructNew("JavaScriptCore", "GlobalContextRef")
	})
}

type GlobalContextRef struct {
	native uintptr
}

var globalContextRefRefFunction *gi.Function
var globalContextRefRefFunction_Once sync.Once

func globalContextRefRefFunction_Set() {
	globalContextRefRefFunction_Once.Do(func() {
		globalContextRefStruct_Set()
		globalContextRefRefFunction = globalContextRefStruct.InvokerNew("ref")
	})
}

// Ref is a representation of the C type JSGlobalContextRetain.
func (recv *GlobalContextRef) Ref() {
	globalContextRefRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	globalContextRefRefFunction.Invoke(inArgs[:], nil)

}

var globalContextRefUnrefFunction *gi.Function
var globalContextRefUnrefFunction_Once sync.Once

func globalContextRefUnrefFunction_Set() {
	globalContextRefUnrefFunction_Once.Do(func() {
		globalContextRefStruct_Set()
		globalContextRefUnrefFunction = globalContextRefStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type JSGlobalContextRelease.
func (recv *GlobalContextRef) Unref() {
	globalContextRefUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	globalContextRefUnrefFunction.Invoke(inArgs[:], nil)

}

var valueRefStruct *gi.Struct
var valueRefStruct_Once sync.Once

func valueRefStruct_Set() {
	valueRefStruct_Once.Do(func() {
		valueRefStruct = gi.StructNew("JavaScriptCore", "ValueRef")
	})
}

type ValueRef struct {
	native uintptr
}

var stringRefStruct *gi.Struct
var stringRefStruct_Once sync.Once

func stringRefStruct_Set() {
	stringRefStruct_Once.Do(func() {
		stringRefStruct = gi.StructNew("JavaScriptCore", "StringRef")
	})
}

type StringRef struct {
	native uintptr
}

var stringRefRefFunction *gi.Function
var stringRefRefFunction_Once sync.Once

func stringRefRefFunction_Set() {
	stringRefRefFunction_Once.Do(func() {
		stringRefStruct_Set()
		stringRefRefFunction = stringRefStruct.InvokerNew("ref")
	})
}

// Ref is a representation of the C type JSStringRetain.
func (recv *StringRef) Ref() {
	stringRefRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	stringRefRefFunction.Invoke(inArgs[:], nil)

}

var stringRefUnrefFunction *gi.Function
var stringRefUnrefFunction_Once sync.Once

func stringRefUnrefFunction_Set() {
	stringRefUnrefFunction_Once.Do(func() {
		stringRefStruct_Set()
		stringRefUnrefFunction = stringRefStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type JSStringRelease.
func (recv *StringRef) Unref() {
	stringRefUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	stringRefUnrefFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'JSStringGetMaximumUTF8CStringSize' : return type 'gsize' not supported

// UNSUPPORTED : C value 'JSStringGetUTF8CString' : parameter 'buffer_size' of type 'gsize' not supported

var classClassStruct *gi.Struct
var classClassStruct_Once sync.Once

func classClassStruct_Set() {
	classClassStruct_Once.Do(func() {
		classClassStruct = gi.StructNew("JavaScriptCore", "ClassClass")
	})
}

type ClassClass struct {
	native uintptr
}

var classVTableStruct *gi.Struct
var classVTableStruct_Once sync.Once

func classVTableStruct_Set() {
	classVTableStruct_Once.Do(func() {
		classVTableStruct = gi.StructNew("JavaScriptCore", "ClassVTable")
	})
}

type ClassVTable struct {
	native uintptr
	// UNSUPPORTED : C value 'get_property' : no Go type for 'ClassGetPropertyFunction'
	// UNSUPPORTED : C value 'set_property' : no Go type for 'ClassSetPropertyFunction'
	// UNSUPPORTED : C value 'has_property' : no Go type for 'ClassHasPropertyFunction'
	// UNSUPPORTED : C value 'delete_property' : no Go type for 'ClassDeletePropertyFunction'
	// UNSUPPORTED : C value 'enumerate_properties' : no Go type for 'ClassEnumeratePropertiesFunction'
	// UNSUPPORTED : C value '_jsc_reserved0' : missing Type
	// UNSUPPORTED : C value '_jsc_reserved1' : missing Type
	// UNSUPPORTED : C value '_jsc_reserved2' : missing Type
	// UNSUPPORTED : C value '_jsc_reserved3' : missing Type
}

var contextClassStruct *gi.Struct
var contextClassStruct_Once sync.Once

func contextClassStruct_Set() {
	contextClassStruct_Once.Do(func() {
		contextClassStruct = gi.StructNew("JavaScriptCore", "ContextClass")
	})
}

type ContextClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_jsc_reserved0' : missing Type
	// UNSUPPORTED : C value '_jsc_reserved1' : missing Type
	// UNSUPPORTED : C value '_jsc_reserved2' : missing Type
	// UNSUPPORTED : C value '_jsc_reserved3' : missing Type
}

var contextPrivateStruct *gi.Struct
var contextPrivateStruct_Once sync.Once

func contextPrivateStruct_Set() {
	contextPrivateStruct_Once.Do(func() {
		contextPrivateStruct = gi.StructNew("JavaScriptCore", "ContextPrivate")
	})
}

type ContextPrivate struct {
	native uintptr
}

var exceptionClassStruct *gi.Struct
var exceptionClassStruct_Once sync.Once

func exceptionClassStruct_Set() {
	exceptionClassStruct_Once.Do(func() {
		exceptionClassStruct = gi.StructNew("JavaScriptCore", "ExceptionClass")
	})
}

type ExceptionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_jsc_reserved0' : missing Type
	// UNSUPPORTED : C value '_jsc_reserved1' : missing Type
	// UNSUPPORTED : C value '_jsc_reserved2' : missing Type
	// UNSUPPORTED : C value '_jsc_reserved3' : missing Type
}

var exceptionPrivateStruct *gi.Struct
var exceptionPrivateStruct_Once sync.Once

func exceptionPrivateStruct_Set() {
	exceptionPrivateStruct_Once.Do(func() {
		exceptionPrivateStruct = gi.StructNew("JavaScriptCore", "ExceptionPrivate")
	})
}

type ExceptionPrivate struct {
	native uintptr
}

var valueClassStruct *gi.Struct
var valueClassStruct_Once sync.Once

func valueClassStruct_Set() {
	valueClassStruct_Once.Do(func() {
		valueClassStruct = gi.StructNew("JavaScriptCore", "ValueClass")
	})
}

type ValueClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_jsc_reserved0' : missing Type
	// UNSUPPORTED : C value '_jsc_reserved1' : missing Type
	// UNSUPPORTED : C value '_jsc_reserved2' : missing Type
	// UNSUPPORTED : C value '_jsc_reserved3' : missing Type
}

var valuePrivateStruct *gi.Struct
var valuePrivateStruct_Once sync.Once

func valuePrivateStruct_Set() {
	valuePrivateStruct_Once.Do(func() {
		valuePrivateStruct = gi.StructNew("JavaScriptCore", "ValuePrivate")
	})
}

type ValuePrivate struct {
	native uintptr
}

var virtualMachineClassStruct *gi.Struct
var virtualMachineClassStruct_Once sync.Once

func virtualMachineClassStruct_Set() {
	virtualMachineClassStruct_Once.Do(func() {
		virtualMachineClassStruct = gi.StructNew("JavaScriptCore", "VirtualMachineClass")
	})
}

type VirtualMachineClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_jsc_reserved0' : missing Type
	// UNSUPPORTED : C value '_jsc_reserved1' : missing Type
	// UNSUPPORTED : C value '_jsc_reserved2' : missing Type
	// UNSUPPORTED : C value '_jsc_reserved3' : missing Type
}

var virtualMachinePrivateStruct *gi.Struct
var virtualMachinePrivateStruct_Once sync.Once

func virtualMachinePrivateStruct_Set() {
	virtualMachinePrivateStruct_Once.Do(func() {
		virtualMachinePrivateStruct = gi.StructNew("JavaScriptCore", "VirtualMachinePrivate")
	})
}

type VirtualMachinePrivate struct {
	native uintptr
}

var weakValueClassStruct *gi.Struct
var weakValueClassStruct_Once sync.Once

func weakValueClassStruct_Set() {
	weakValueClassStruct_Once.Do(func() {
		weakValueClassStruct = gi.StructNew("JavaScriptCore", "WeakValueClass")
	})
}

type WeakValueClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_jsc_reserved0' : missing Type
	// UNSUPPORTED : C value '_jsc_reserved1' : missing Type
	// UNSUPPORTED : C value '_jsc_reserved2' : missing Type
	// UNSUPPORTED : C value '_jsc_reserved3' : missing Type
}

var weakValuePrivateStruct *gi.Struct
var weakValuePrivateStruct_Once sync.Once

func weakValuePrivateStruct_Set() {
	weakValuePrivateStruct_Once.Do(func() {
		weakValuePrivateStruct = gi.StructNew("JavaScriptCore", "WeakValuePrivate")
	})
}

type WeakValuePrivate struct {
	native uintptr
}
