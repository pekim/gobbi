// Code generated - DO NOT EDIT.

package javascriptcore

import gi "github.com/pekim/gobbi/internal/gi"

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

type ValueRef struct {
	native uintptr
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

type ClassClass struct {
	native uintptr
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

type ContextClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_jsc_reserved0' : missing Type

	// UNSUPPORTED : C value '_jsc_reserved1' : missing Type

	// UNSUPPORTED : C value '_jsc_reserved2' : missing Type

	// UNSUPPORTED : C value '_jsc_reserved3' : missing Type

}

type ContextPrivate struct {
	native uintptr
}

type ExceptionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_jsc_reserved0' : missing Type

	// UNSUPPORTED : C value '_jsc_reserved1' : missing Type

	// UNSUPPORTED : C value '_jsc_reserved2' : missing Type

	// UNSUPPORTED : C value '_jsc_reserved3' : missing Type

}

type ExceptionPrivate struct {
	native uintptr
}

type ValueClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_jsc_reserved0' : missing Type

	// UNSUPPORTED : C value '_jsc_reserved1' : missing Type

	// UNSUPPORTED : C value '_jsc_reserved2' : missing Type

	// UNSUPPORTED : C value '_jsc_reserved3' : missing Type

}

type ValuePrivate struct {
	native uintptr
}

type VirtualMachineClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_jsc_reserved0' : missing Type

	// UNSUPPORTED : C value '_jsc_reserved1' : missing Type

	// UNSUPPORTED : C value '_jsc_reserved2' : missing Type

	// UNSUPPORTED : C value '_jsc_reserved3' : missing Type

}

type VirtualMachinePrivate struct {
	native uintptr
}

type WeakValueClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_jsc_reserved0' : missing Type

	// UNSUPPORTED : C value '_jsc_reserved1' : missing Type

	// UNSUPPORTED : C value '_jsc_reserved2' : missing Type

	// UNSUPPORTED : C value '_jsc_reserved3' : missing Type

}

type WeakValuePrivate struct {
	native uintptr
}
