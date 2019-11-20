// Code generated - DO NOT EDIT.

package javascriptcore

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var GlobalContextRefStruct *gi.Struct
var GlobalContextRefStructOnce sync.Once

func GlobalContextRefStructSet() {
	GlobalContextRefStructOnce.Do(func() {
		GlobalContextRefStruct = gi.StructNew("JavaScriptCore", "GlobalContextRef")
	})
}

type GlobalContextRef struct {
	native uintptr
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

var ValueRefStruct *gi.Struct
var ValueRefStructOnce sync.Once

func ValueRefStructSet() {
	ValueRefStructOnce.Do(func() {
		ValueRefStruct = gi.StructNew("JavaScriptCore", "ValueRef")
	})
}

type ValueRef struct {
	native uintptr
}

var StringRefStruct *gi.Struct
var StringRefStructOnce sync.Once

func StringRefStructSet() {
	StringRefStructOnce.Do(func() {
		StringRefStruct = gi.StructNew("JavaScriptCore", "StringRef")
	})
}

type StringRef struct {
	native uintptr
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

var ClassClassStruct *gi.Struct
var ClassClassStructOnce sync.Once

func ClassClassStructSet() {
	ClassClassStructOnce.Do(func() {
		ClassClassStruct = gi.StructNew("JavaScriptCore", "ClassClass")
	})
}

type ClassClass struct {
	native uintptr
}

var ClassVTableStruct *gi.Struct
var ClassVTableStructOnce sync.Once

func ClassVTableStructSet() {
	ClassVTableStructOnce.Do(func() {
		ClassVTableStruct = gi.StructNew("JavaScriptCore", "ClassVTable")
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

var ContextClassStruct *gi.Struct
var ContextClassStructOnce sync.Once

func ContextClassStructSet() {
	ContextClassStructOnce.Do(func() {
		ContextClassStruct = gi.StructNew("JavaScriptCore", "ContextClass")
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

var ContextPrivateStruct *gi.Struct
var ContextPrivateStructOnce sync.Once

func ContextPrivateStructSet() {
	ContextPrivateStructOnce.Do(func() {
		ContextPrivateStruct = gi.StructNew("JavaScriptCore", "ContextPrivate")
	})
}

type ContextPrivate struct {
	native uintptr
}

var ExceptionClassStruct *gi.Struct
var ExceptionClassStructOnce sync.Once

func ExceptionClassStructSet() {
	ExceptionClassStructOnce.Do(func() {
		ExceptionClassStruct = gi.StructNew("JavaScriptCore", "ExceptionClass")
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

var ExceptionPrivateStruct *gi.Struct
var ExceptionPrivateStructOnce sync.Once

func ExceptionPrivateStructSet() {
	ExceptionPrivateStructOnce.Do(func() {
		ExceptionPrivateStruct = gi.StructNew("JavaScriptCore", "ExceptionPrivate")
	})
}

type ExceptionPrivate struct {
	native uintptr
}

var ValueClassStruct *gi.Struct
var ValueClassStructOnce sync.Once

func ValueClassStructSet() {
	ValueClassStructOnce.Do(func() {
		ValueClassStruct = gi.StructNew("JavaScriptCore", "ValueClass")
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

var ValuePrivateStruct *gi.Struct
var ValuePrivateStructOnce sync.Once

func ValuePrivateStructSet() {
	ValuePrivateStructOnce.Do(func() {
		ValuePrivateStruct = gi.StructNew("JavaScriptCore", "ValuePrivate")
	})
}

type ValuePrivate struct {
	native uintptr
}

var VirtualMachineClassStruct *gi.Struct
var VirtualMachineClassStructOnce sync.Once

func VirtualMachineClassStructSet() {
	VirtualMachineClassStructOnce.Do(func() {
		VirtualMachineClassStruct = gi.StructNew("JavaScriptCore", "VirtualMachineClass")
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

var VirtualMachinePrivateStruct *gi.Struct
var VirtualMachinePrivateStructOnce sync.Once

func VirtualMachinePrivateStructSet() {
	VirtualMachinePrivateStructOnce.Do(func() {
		VirtualMachinePrivateStruct = gi.StructNew("JavaScriptCore", "VirtualMachinePrivate")
	})
}

type VirtualMachinePrivate struct {
	native uintptr
}

var WeakValueClassStruct *gi.Struct
var WeakValueClassStructOnce sync.Once

func WeakValueClassStructSet() {
	WeakValueClassStructOnce.Do(func() {
		WeakValueClassStruct = gi.StructNew("JavaScriptCore", "WeakValueClass")
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

var WeakValuePrivateStruct *gi.Struct
var WeakValuePrivateStructOnce sync.Once

func WeakValuePrivateStructSet() {
	WeakValuePrivateStructOnce.Do(func() {
		WeakValuePrivateStruct = gi.StructNew("JavaScriptCore", "WeakValuePrivate")
	})
}

type WeakValuePrivate struct {
	native uintptr
}
