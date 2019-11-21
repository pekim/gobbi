// Code generated - DO NOT EDIT.

package javascriptcore

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var globalContextRefStruct *gi.Struct
var globalContextRefStructOnce sync.Once

func globalContextRefStructSet() {
	globalContextRefStructOnce.Do(func() {
		globalContextRefStruct = gi.StructNew("JavaScriptCore", "GlobalContextRef")
	})
}

type GlobalContextRef struct {
	native uintptr
}

var globalContextRefRefFunction *gi.Function
var globalContextRefRefFunctionOnce sync.Once

func globalContextRefRefFunctionSet() {
	globalContextRefRefFunctionOnce.Do(func() {
		globalContextRefRefFunction = gi.FunctionInvokerNew("JavaScriptCore", "ref")
	})
}

var refGlobalContextRefInvoker *gi.Function

// Ref is a representation of the C type JSGlobalContextRetain.
func (recv *GlobalContextRef) Ref() {
	if refGlobalContextRefInvoker == nil {
		refGlobalContextRefInvoker = gi.StructFunctionInvokerNew("JavaScriptCore", "GlobalContextRef", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	refGlobalContextRefInvoker.Invoke(inArgs[:], nil)

}

var globalContextRefUnrefFunction *gi.Function
var globalContextRefUnrefFunctionOnce sync.Once

func globalContextRefUnrefFunctionSet() {
	globalContextRefUnrefFunctionOnce.Do(func() {
		globalContextRefUnrefFunction = gi.FunctionInvokerNew("JavaScriptCore", "unref")
	})
}

var unrefGlobalContextRefInvoker *gi.Function

// Unref is a representation of the C type JSGlobalContextRelease.
func (recv *GlobalContextRef) Unref() {
	if unrefGlobalContextRefInvoker == nil {
		unrefGlobalContextRefInvoker = gi.StructFunctionInvokerNew("JavaScriptCore", "GlobalContextRef", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefGlobalContextRefInvoker.Invoke(inArgs[:], nil)

}

var valueRefStruct *gi.Struct
var valueRefStructOnce sync.Once

func valueRefStructSet() {
	valueRefStructOnce.Do(func() {
		valueRefStruct = gi.StructNew("JavaScriptCore", "ValueRef")
	})
}

type ValueRef struct {
	native uintptr
}

var stringRefStruct *gi.Struct
var stringRefStructOnce sync.Once

func stringRefStructSet() {
	stringRefStructOnce.Do(func() {
		stringRefStruct = gi.StructNew("JavaScriptCore", "StringRef")
	})
}

type StringRef struct {
	native uintptr
}

var stringRefRefFunction *gi.Function
var stringRefRefFunctionOnce sync.Once

func stringRefRefFunctionSet() {
	stringRefRefFunctionOnce.Do(func() {
		stringRefRefFunction = gi.FunctionInvokerNew("JavaScriptCore", "ref")
	})
}

var refStringRefInvoker *gi.Function

// Ref is a representation of the C type JSStringRetain.
func (recv *StringRef) Ref() {
	if refStringRefInvoker == nil {
		refStringRefInvoker = gi.StructFunctionInvokerNew("JavaScriptCore", "StringRef", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	refStringRefInvoker.Invoke(inArgs[:], nil)

}

var stringRefUnrefFunction *gi.Function
var stringRefUnrefFunctionOnce sync.Once

func stringRefUnrefFunctionSet() {
	stringRefUnrefFunctionOnce.Do(func() {
		stringRefUnrefFunction = gi.FunctionInvokerNew("JavaScriptCore", "unref")
	})
}

var unrefStringRefInvoker *gi.Function

// Unref is a representation of the C type JSStringRelease.
func (recv *StringRef) Unref() {
	if unrefStringRefInvoker == nil {
		unrefStringRefInvoker = gi.StructFunctionInvokerNew("JavaScriptCore", "StringRef", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefStringRefInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'JSStringGetMaximumUTF8CStringSize' : return type 'gsize' not supported

// UNSUPPORTED : C value 'JSStringGetUTF8CString' : parameter 'buffer_size' of type 'gsize' not supported

var classClassStruct *gi.Struct
var classClassStructOnce sync.Once

func classClassStructSet() {
	classClassStructOnce.Do(func() {
		classClassStruct = gi.StructNew("JavaScriptCore", "ClassClass")
	})
}

type ClassClass struct {
	native uintptr
}

var classVTableStruct *gi.Struct
var classVTableStructOnce sync.Once

func classVTableStructSet() {
	classVTableStructOnce.Do(func() {
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
var contextClassStructOnce sync.Once

func contextClassStructSet() {
	contextClassStructOnce.Do(func() {
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
var contextPrivateStructOnce sync.Once

func contextPrivateStructSet() {
	contextPrivateStructOnce.Do(func() {
		contextPrivateStruct = gi.StructNew("JavaScriptCore", "ContextPrivate")
	})
}

type ContextPrivate struct {
	native uintptr
}

var exceptionClassStruct *gi.Struct
var exceptionClassStructOnce sync.Once

func exceptionClassStructSet() {
	exceptionClassStructOnce.Do(func() {
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
var exceptionPrivateStructOnce sync.Once

func exceptionPrivateStructSet() {
	exceptionPrivateStructOnce.Do(func() {
		exceptionPrivateStruct = gi.StructNew("JavaScriptCore", "ExceptionPrivate")
	})
}

type ExceptionPrivate struct {
	native uintptr
}

var valueClassStruct *gi.Struct
var valueClassStructOnce sync.Once

func valueClassStructSet() {
	valueClassStructOnce.Do(func() {
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
var valuePrivateStructOnce sync.Once

func valuePrivateStructSet() {
	valuePrivateStructOnce.Do(func() {
		valuePrivateStruct = gi.StructNew("JavaScriptCore", "ValuePrivate")
	})
}

type ValuePrivate struct {
	native uintptr
}

var virtualMachineClassStruct *gi.Struct
var virtualMachineClassStructOnce sync.Once

func virtualMachineClassStructSet() {
	virtualMachineClassStructOnce.Do(func() {
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
var virtualMachinePrivateStructOnce sync.Once

func virtualMachinePrivateStructSet() {
	virtualMachinePrivateStructOnce.Do(func() {
		virtualMachinePrivateStruct = gi.StructNew("JavaScriptCore", "VirtualMachinePrivate")
	})
}

type VirtualMachinePrivate struct {
	native uintptr
}

var weakValueClassStruct *gi.Struct
var weakValueClassStructOnce sync.Once

func weakValueClassStructSet() {
	weakValueClassStructOnce.Do(func() {
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
var weakValuePrivateStructOnce sync.Once

func weakValuePrivateStructSet() {
	weakValuePrivateStructOnce.Do(func() {
		weakValuePrivateStruct = gi.StructNew("JavaScriptCore", "WeakValuePrivate")
	})
}

type WeakValuePrivate struct {
	native uintptr
}
