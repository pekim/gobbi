// Code generated - DO NOT EDIT.

package javascriptcore

import (
	callback "github.com/pekim/gobbi/internal/cgo/callback"
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
)

var classObject *gi.Object
var classObject_Once sync.Once

func classObject_Set() error {
	var err error
	classObject_Once.Do(func() {
		classObject, err = gi.ObjectNew("JavaScriptCore", "Class")
	})
	return err
}

type Class struct {
	native unsafe.Pointer
}

func ClassNewFromNative(native unsafe.Pointer) *Class {
	instance := &Class{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Class) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToClass down casts any arbitrary Object to Class.
Exercise care, as this is a potentially dangerous function
if the Object is not a Class.
*/
func CastToClass(object *gobject.Object) *Class {
	return ClassNewFromNative(object.Native())
}

// Equals compares this Class with another Class, and returns true if they represent the same Object.
func (recv *Class) Equals(other *Class) bool {
	return other.Native() == recv.Native()
}

func (recv *Class) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'jsc_class_add_constructor' : parameter 'callback' of type 'GObject.Callback' not supported

// UNSUPPORTED : C value 'jsc_class_add_constructor_variadic' : parameter 'callback' of type 'Constructor' not supported

// UNSUPPORTED : C value 'jsc_class_add_constructorv' : parameter 'callback' of type 'GObject.Callback' not supported

// UNSUPPORTED : C value 'jsc_class_add_method' : parameter 'callback' of type 'GObject.Callback' not supported

// UNSUPPORTED : C value 'jsc_class_add_method_variadic' : parameter 'callback' of type 'ClassVariadicFunction' not supported

// UNSUPPORTED : C value 'jsc_class_add_methodv' : parameter 'callback' of type 'GObject.Callback' not supported

// UNSUPPORTED : C value 'jsc_class_add_property' : parameter 'property_type' of type 'GType' not supported

var classGetNameFunction *gi.Function
var classGetNameFunction_Once sync.Once

func classGetNameFunction_Set() error {
	var err error
	classGetNameFunction_Once.Do(func() {
		err = classObject_Set()
		if err != nil {
			return
		}
		classGetNameFunction, err = classObject.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type jsc_class_get_name.
func (recv *Class) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := classGetNameFunction_Set()
	if err == nil {
		ret = classGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var classGetParentFunction *gi.Function
var classGetParentFunction_Once sync.Once

func classGetParentFunction_Set() error {
	var err error
	classGetParentFunction_Once.Do(func() {
		err = classObject_Set()
		if err != nil {
			return
		}
		classGetParentFunction, err = classObject.InvokerNew("get_parent")
	})
	return err
}

// GetParent is a representation of the C type jsc_class_get_parent.
func (recv *Class) GetParent() *Class {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := classGetParentFunction_Set()
	if err == nil {
		ret = classGetParentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ClassNewFromNative(ret.Pointer())

	return retGo
}

var contextObject *gi.Object
var contextObject_Once sync.Once

func contextObject_Set() error {
	var err error
	contextObject_Once.Do(func() {
		contextObject, err = gi.ObjectNew("JavaScriptCore", "Context")
	})
	return err
}

type Context struct {
	native unsafe.Pointer
}

func ContextNewFromNative(native unsafe.Pointer) *Context {
	instance := &Context{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Context) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToContext down casts any arbitrary Object to Context.
Exercise care, as this is a potentially dangerous function
if the Object is not a Context.
*/
func CastToContext(object *gobject.Object) *Context {
	return ContextNewFromNative(object.Native())
}

// Equals compares this Context with another Context, and returns true if they represent the same Object.
func (recv *Context) Equals(other *Context) bool {
	return other.Native() == recv.Native()
}

func (recv *Context) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Context) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(contextObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Context) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(contextObject, recv.Native(), "parent", argValue)
}

var contextNewFunction *gi.Function
var contextNewFunction_Once sync.Once

func contextNewFunction_Set() error {
	var err error
	contextNewFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextNewFunction, err = contextObject.InvokerNew("new")
	})
	return err
}

// ContextNew is a representation of the C type jsc_context_new.
func ContextNew() *Context {

	var ret gi.Argument

	err := contextNewFunction_Set()
	if err == nil {
		ret = contextNewFunction.Invoke(nil, nil)
	}

	retGo := ContextNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var contextNewWithVirtualMachineFunction *gi.Function
var contextNewWithVirtualMachineFunction_Once sync.Once

func contextNewWithVirtualMachineFunction_Set() error {
	var err error
	contextNewWithVirtualMachineFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextNewWithVirtualMachineFunction, err = contextObject.InvokerNew("new_with_virtual_machine")
	})
	return err
}

// ContextNewWithVirtualMachine is a representation of the C type jsc_context_new_with_virtual_machine.
func ContextNewWithVirtualMachine(vm *VirtualMachine) *Context {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(vm.Native())

	var ret gi.Argument

	err := contextNewWithVirtualMachineFunction_Set()
	if err == nil {
		ret = contextNewWithVirtualMachineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ContextNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var contextCheckSyntaxFunction *gi.Function
var contextCheckSyntaxFunction_Once sync.Once

func contextCheckSyntaxFunction_Set() error {
	var err error
	contextCheckSyntaxFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextCheckSyntaxFunction, err = contextObject.InvokerNew("check_syntax")
	})
	return err
}

// CheckSyntax is a representation of the C type jsc_context_check_syntax.
func (recv *Context) CheckSyntax(code string, length int32, mode CheckSyntaxMode, uri string, lineNumber uint32) (CheckSyntaxResult, *Exception) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(code)
	inArgs[2].SetInt32(length)
	inArgs[3].SetInt32(int32(mode))
	inArgs[4].SetString(uri)
	inArgs[5].SetUint32(lineNumber)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := contextCheckSyntaxFunction_Set()
	if err == nil {
		ret = contextCheckSyntaxFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := CheckSyntaxResult(ret.Int32())
	out0 := ExceptionNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

var contextClearExceptionFunction *gi.Function
var contextClearExceptionFunction_Once sync.Once

func contextClearExceptionFunction_Set() error {
	var err error
	contextClearExceptionFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextClearExceptionFunction, err = contextObject.InvokerNew("clear_exception")
	})
	return err
}

// ClearException is a representation of the C type jsc_context_clear_exception.
func (recv *Context) ClearException() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := contextClearExceptionFunction_Set()
	if err == nil {
		contextClearExceptionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var contextEvaluateFunction *gi.Function
var contextEvaluateFunction_Once sync.Once

func contextEvaluateFunction_Set() error {
	var err error
	contextEvaluateFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextEvaluateFunction, err = contextObject.InvokerNew("evaluate")
	})
	return err
}

// Evaluate is a representation of the C type jsc_context_evaluate.
func (recv *Context) Evaluate(code string, length int32) *Value {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(code)
	inArgs[2].SetInt32(length)

	var ret gi.Argument

	err := contextEvaluateFunction_Set()
	if err == nil {
		ret = contextEvaluateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ValueNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'jsc_context_evaluate_in_object' : parameter 'object_instance' of type 'gpointer' not supported

var contextEvaluateWithSourceUriFunction *gi.Function
var contextEvaluateWithSourceUriFunction_Once sync.Once

func contextEvaluateWithSourceUriFunction_Set() error {
	var err error
	contextEvaluateWithSourceUriFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextEvaluateWithSourceUriFunction, err = contextObject.InvokerNew("evaluate_with_source_uri")
	})
	return err
}

// EvaluateWithSourceUri is a representation of the C type jsc_context_evaluate_with_source_uri.
func (recv *Context) EvaluateWithSourceUri(code string, length int32, uri string, lineNumber uint32) *Value {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(code)
	inArgs[2].SetInt32(length)
	inArgs[3].SetString(uri)
	inArgs[4].SetUint32(lineNumber)

	var ret gi.Argument

	err := contextEvaluateWithSourceUriFunction_Set()
	if err == nil {
		ret = contextEvaluateWithSourceUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ValueNewFromNative(ret.Pointer())

	return retGo
}

var contextGetExceptionFunction *gi.Function
var contextGetExceptionFunction_Once sync.Once

func contextGetExceptionFunction_Set() error {
	var err error
	contextGetExceptionFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextGetExceptionFunction, err = contextObject.InvokerNew("get_exception")
	})
	return err
}

// GetException is a representation of the C type jsc_context_get_exception.
func (recv *Context) GetException() *Exception {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := contextGetExceptionFunction_Set()
	if err == nil {
		ret = contextGetExceptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ExceptionNewFromNative(ret.Pointer())

	return retGo
}

var contextGetGlobalObjectFunction *gi.Function
var contextGetGlobalObjectFunction_Once sync.Once

func contextGetGlobalObjectFunction_Set() error {
	var err error
	contextGetGlobalObjectFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextGetGlobalObjectFunction, err = contextObject.InvokerNew("get_global_object")
	})
	return err
}

// GetGlobalObject is a representation of the C type jsc_context_get_global_object.
func (recv *Context) GetGlobalObject() *Value {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := contextGetGlobalObjectFunction_Set()
	if err == nil {
		ret = contextGetGlobalObjectFunction.Invoke(inArgs[:], nil)
	}

	retGo := ValueNewFromNative(ret.Pointer())

	return retGo
}

var contextGetValueFunction *gi.Function
var contextGetValueFunction_Once sync.Once

func contextGetValueFunction_Set() error {
	var err error
	contextGetValueFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextGetValueFunction, err = contextObject.InvokerNew("get_value")
	})
	return err
}

// GetValue is a representation of the C type jsc_context_get_value.
func (recv *Context) GetValue(name string) *Value {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := contextGetValueFunction_Set()
	if err == nil {
		ret = contextGetValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := ValueNewFromNative(ret.Pointer())

	return retGo
}

var contextGetVirtualMachineFunction *gi.Function
var contextGetVirtualMachineFunction_Once sync.Once

func contextGetVirtualMachineFunction_Set() error {
	var err error
	contextGetVirtualMachineFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextGetVirtualMachineFunction, err = contextObject.InvokerNew("get_virtual_machine")
	})
	return err
}

// GetVirtualMachine is a representation of the C type jsc_context_get_virtual_machine.
func (recv *Context) GetVirtualMachine() *VirtualMachine {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := contextGetVirtualMachineFunction_Set()
	if err == nil {
		ret = contextGetVirtualMachineFunction.Invoke(inArgs[:], nil)
	}

	retGo := VirtualMachineNewFromNative(ret.Pointer())

	return retGo
}

var contextPopExceptionHandlerFunction *gi.Function
var contextPopExceptionHandlerFunction_Once sync.Once

func contextPopExceptionHandlerFunction_Set() error {
	var err error
	contextPopExceptionHandlerFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextPopExceptionHandlerFunction, err = contextObject.InvokerNew("pop_exception_handler")
	})
	return err
}

// PopExceptionHandler is a representation of the C type jsc_context_pop_exception_handler.
func (recv *Context) PopExceptionHandler() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := contextPopExceptionHandlerFunction_Set()
	if err == nil {
		contextPopExceptionHandlerFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'jsc_context_push_exception_handler' : parameter 'handler' of type 'ExceptionHandler' not supported

// UNSUPPORTED : C value 'jsc_context_register_class' : parameter 'destroy_notify' of type 'GLib.DestroyNotify' not supported

var contextSetValueFunction *gi.Function
var contextSetValueFunction_Once sync.Once

func contextSetValueFunction_Set() error {
	var err error
	contextSetValueFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextSetValueFunction, err = contextObject.InvokerNew("set_value")
	})
	return err
}

// SetValue is a representation of the C type jsc_context_set_value.
func (recv *Context) SetValue(name string, value *Value) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)
	inArgs[2].SetPointer(value.Native())

	err := contextSetValueFunction_Set()
	if err == nil {
		contextSetValueFunction.Invoke(inArgs[:], nil)
	}

	return
}

var contextThrowFunction *gi.Function
var contextThrowFunction_Once sync.Once

func contextThrowFunction_Set() error {
	var err error
	contextThrowFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextThrowFunction, err = contextObject.InvokerNew("throw")
	})
	return err
}

// Throw is a representation of the C type jsc_context_throw.
func (recv *Context) Throw(errorMessage string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(errorMessage)

	err := contextThrowFunction_Set()
	if err == nil {
		contextThrowFunction.Invoke(inArgs[:], nil)
	}

	return
}

var contextThrowExceptionFunction *gi.Function
var contextThrowExceptionFunction_Once sync.Once

func contextThrowExceptionFunction_Set() error {
	var err error
	contextThrowExceptionFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextThrowExceptionFunction, err = contextObject.InvokerNew("throw_exception")
	})
	return err
}

// ThrowException is a representation of the C type jsc_context_throw_exception.
func (recv *Context) ThrowException(exception *Exception) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(exception.Native())

	err := contextThrowExceptionFunction_Set()
	if err == nil {
		contextThrowExceptionFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'jsc_context_throw_printf' : parameter '...' of type 'nil' not supported

var contextThrowWithNameFunction *gi.Function
var contextThrowWithNameFunction_Once sync.Once

func contextThrowWithNameFunction_Set() error {
	var err error
	contextThrowWithNameFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextThrowWithNameFunction, err = contextObject.InvokerNew("throw_with_name")
	})
	return err
}

// ThrowWithName is a representation of the C type jsc_context_throw_with_name.
func (recv *Context) ThrowWithName(errorName string, errorMessage string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(errorName)
	inArgs[2].SetString(errorMessage)

	err := contextThrowWithNameFunction_Set()
	if err == nil {
		contextThrowWithNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'jsc_context_throw_with_name_printf' : parameter '...' of type 'nil' not supported

var exceptionObject *gi.Object
var exceptionObject_Once sync.Once

func exceptionObject_Set() error {
	var err error
	exceptionObject_Once.Do(func() {
		exceptionObject, err = gi.ObjectNew("JavaScriptCore", "Exception")
	})
	return err
}

type Exception struct {
	native unsafe.Pointer
}

func ExceptionNewFromNative(native unsafe.Pointer) *Exception {
	instance := &Exception{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Exception) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToException down casts any arbitrary Object to Exception.
Exercise care, as this is a potentially dangerous function
if the Object is not a Exception.
*/
func CastToException(object *gobject.Object) *Exception {
	return ExceptionNewFromNative(object.Native())
}

// Equals compares this Exception with another Exception, and returns true if they represent the same Object.
func (recv *Exception) Equals(other *Exception) bool {
	return other.Native() == recv.Native()
}

func (recv *Exception) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Exception) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(exceptionObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Exception) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(exceptionObject, recv.Native(), "parent", argValue)
}

var exceptionNewFunction *gi.Function
var exceptionNewFunction_Once sync.Once

func exceptionNewFunction_Set() error {
	var err error
	exceptionNewFunction_Once.Do(func() {
		err = exceptionObject_Set()
		if err != nil {
			return
		}
		exceptionNewFunction, err = exceptionObject.InvokerNew("new")
	})
	return err
}

// ExceptionNew is a representation of the C type jsc_exception_new.
func ExceptionNew(context *Context, message string) *Exception {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(context.Native())
	inArgs[1].SetString(message)

	var ret gi.Argument

	err := exceptionNewFunction_Set()
	if err == nil {
		ret = exceptionNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := ExceptionNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'jsc_exception_new_printf' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'jsc_exception_new_vprintf' : parameter 'args' of type 'va_list' not supported

var exceptionNewWithNameFunction *gi.Function
var exceptionNewWithNameFunction_Once sync.Once

func exceptionNewWithNameFunction_Set() error {
	var err error
	exceptionNewWithNameFunction_Once.Do(func() {
		err = exceptionObject_Set()
		if err != nil {
			return
		}
		exceptionNewWithNameFunction, err = exceptionObject.InvokerNew("new_with_name")
	})
	return err
}

// ExceptionNewWithName is a representation of the C type jsc_exception_new_with_name.
func ExceptionNewWithName(context *Context, name string, message string) *Exception {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(context.Native())
	inArgs[1].SetString(name)
	inArgs[2].SetString(message)

	var ret gi.Argument

	err := exceptionNewWithNameFunction_Set()
	if err == nil {
		ret = exceptionNewWithNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ExceptionNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'jsc_exception_new_with_name_printf' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'jsc_exception_new_with_name_vprintf' : parameter 'args' of type 'va_list' not supported

var exceptionGetBacktraceStringFunction *gi.Function
var exceptionGetBacktraceStringFunction_Once sync.Once

func exceptionGetBacktraceStringFunction_Set() error {
	var err error
	exceptionGetBacktraceStringFunction_Once.Do(func() {
		err = exceptionObject_Set()
		if err != nil {
			return
		}
		exceptionGetBacktraceStringFunction, err = exceptionObject.InvokerNew("get_backtrace_string")
	})
	return err
}

// GetBacktraceString is a representation of the C type jsc_exception_get_backtrace_string.
func (recv *Exception) GetBacktraceString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := exceptionGetBacktraceStringFunction_Set()
	if err == nil {
		ret = exceptionGetBacktraceStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var exceptionGetColumnNumberFunction *gi.Function
var exceptionGetColumnNumberFunction_Once sync.Once

func exceptionGetColumnNumberFunction_Set() error {
	var err error
	exceptionGetColumnNumberFunction_Once.Do(func() {
		err = exceptionObject_Set()
		if err != nil {
			return
		}
		exceptionGetColumnNumberFunction, err = exceptionObject.InvokerNew("get_column_number")
	})
	return err
}

// GetColumnNumber is a representation of the C type jsc_exception_get_column_number.
func (recv *Exception) GetColumnNumber() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := exceptionGetColumnNumberFunction_Set()
	if err == nil {
		ret = exceptionGetColumnNumberFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var exceptionGetLineNumberFunction *gi.Function
var exceptionGetLineNumberFunction_Once sync.Once

func exceptionGetLineNumberFunction_Set() error {
	var err error
	exceptionGetLineNumberFunction_Once.Do(func() {
		err = exceptionObject_Set()
		if err != nil {
			return
		}
		exceptionGetLineNumberFunction, err = exceptionObject.InvokerNew("get_line_number")
	})
	return err
}

// GetLineNumber is a representation of the C type jsc_exception_get_line_number.
func (recv *Exception) GetLineNumber() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := exceptionGetLineNumberFunction_Set()
	if err == nil {
		ret = exceptionGetLineNumberFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var exceptionGetMessageFunction *gi.Function
var exceptionGetMessageFunction_Once sync.Once

func exceptionGetMessageFunction_Set() error {
	var err error
	exceptionGetMessageFunction_Once.Do(func() {
		err = exceptionObject_Set()
		if err != nil {
			return
		}
		exceptionGetMessageFunction, err = exceptionObject.InvokerNew("get_message")
	})
	return err
}

// GetMessage is a representation of the C type jsc_exception_get_message.
func (recv *Exception) GetMessage() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := exceptionGetMessageFunction_Set()
	if err == nil {
		ret = exceptionGetMessageFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var exceptionGetNameFunction *gi.Function
var exceptionGetNameFunction_Once sync.Once

func exceptionGetNameFunction_Set() error {
	var err error
	exceptionGetNameFunction_Once.Do(func() {
		err = exceptionObject_Set()
		if err != nil {
			return
		}
		exceptionGetNameFunction, err = exceptionObject.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type jsc_exception_get_name.
func (recv *Exception) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := exceptionGetNameFunction_Set()
	if err == nil {
		ret = exceptionGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var exceptionGetSourceUriFunction *gi.Function
var exceptionGetSourceUriFunction_Once sync.Once

func exceptionGetSourceUriFunction_Set() error {
	var err error
	exceptionGetSourceUriFunction_Once.Do(func() {
		err = exceptionObject_Set()
		if err != nil {
			return
		}
		exceptionGetSourceUriFunction, err = exceptionObject.InvokerNew("get_source_uri")
	})
	return err
}

// GetSourceUri is a representation of the C type jsc_exception_get_source_uri.
func (recv *Exception) GetSourceUri() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := exceptionGetSourceUriFunction_Set()
	if err == nil {
		ret = exceptionGetSourceUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var exceptionReportFunction *gi.Function
var exceptionReportFunction_Once sync.Once

func exceptionReportFunction_Set() error {
	var err error
	exceptionReportFunction_Once.Do(func() {
		err = exceptionObject_Set()
		if err != nil {
			return
		}
		exceptionReportFunction, err = exceptionObject.InvokerNew("report")
	})
	return err
}

// Report is a representation of the C type jsc_exception_report.
func (recv *Exception) Report() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := exceptionReportFunction_Set()
	if err == nil {
		ret = exceptionReportFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var exceptionToStringFunction *gi.Function
var exceptionToStringFunction_Once sync.Once

func exceptionToStringFunction_Set() error {
	var err error
	exceptionToStringFunction_Once.Do(func() {
		err = exceptionObject_Set()
		if err != nil {
			return
		}
		exceptionToStringFunction, err = exceptionObject.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type jsc_exception_to_string.
func (recv *Exception) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := exceptionToStringFunction_Set()
	if err == nil {
		ret = exceptionToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var valueObject *gi.Object
var valueObject_Once sync.Once

func valueObject_Set() error {
	var err error
	valueObject_Once.Do(func() {
		valueObject, err = gi.ObjectNew("JavaScriptCore", "Value")
	})
	return err
}

type Value struct {
	native unsafe.Pointer
}

func ValueNewFromNative(native unsafe.Pointer) *Value {
	instance := &Value{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Value) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToValue down casts any arbitrary Object to Value.
Exercise care, as this is a potentially dangerous function
if the Object is not a Value.
*/
func CastToValue(object *gobject.Object) *Value {
	return ValueNewFromNative(object.Native())
}

// Equals compares this Value with another Value, and returns true if they represent the same Object.
func (recv *Value) Equals(other *Value) bool {
	return other.Native() == recv.Native()
}

func (recv *Value) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Value) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(valueObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Value) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(valueObject, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'jsc_value_new_array' : parameter 'first_item_type' of type 'GType' not supported

// UNSUPPORTED : C value 'jsc_value_new_array_from_garray' : parameter 'array' of type 'nil' not supported

// UNSUPPORTED : C value 'jsc_value_new_array_from_strv' : parameter 'strv' of type 'nil' not supported

var valueNewBooleanFunction *gi.Function
var valueNewBooleanFunction_Once sync.Once

func valueNewBooleanFunction_Set() error {
	var err error
	valueNewBooleanFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueNewBooleanFunction, err = valueObject.InvokerNew("new_boolean")
	})
	return err
}

// ValueNewBoolean is a representation of the C type jsc_value_new_boolean.
func ValueNewBoolean(context *Context, value bool) *Value {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(context.Native())
	inArgs[1].SetBoolean(value)

	var ret gi.Argument

	err := valueNewBooleanFunction_Set()
	if err == nil {
		ret = valueNewBooleanFunction.Invoke(inArgs[:], nil)
	}

	retGo := ValueNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'jsc_value_new_function' : parameter 'callback' of type 'GObject.Callback' not supported

// UNSUPPORTED : C value 'jsc_value_new_function_variadic' : parameter 'callback' of type 'VariadicFunction' not supported

// UNSUPPORTED : C value 'jsc_value_new_functionv' : parameter 'callback' of type 'GObject.Callback' not supported

var valueNewNullFunction *gi.Function
var valueNewNullFunction_Once sync.Once

func valueNewNullFunction_Set() error {
	var err error
	valueNewNullFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueNewNullFunction, err = valueObject.InvokerNew("new_null")
	})
	return err
}

// ValueNewNull is a representation of the C type jsc_value_new_null.
func ValueNewNull(context *Context) *Value {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(context.Native())

	var ret gi.Argument

	err := valueNewNullFunction_Set()
	if err == nil {
		ret = valueNewNullFunction.Invoke(inArgs[:], nil)
	}

	retGo := ValueNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var valueNewNumberFunction *gi.Function
var valueNewNumberFunction_Once sync.Once

func valueNewNumberFunction_Set() error {
	var err error
	valueNewNumberFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueNewNumberFunction, err = valueObject.InvokerNew("new_number")
	})
	return err
}

// ValueNewNumber is a representation of the C type jsc_value_new_number.
func ValueNewNumber(context *Context, number float64) *Value {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(context.Native())
	inArgs[1].SetFloat64(number)

	var ret gi.Argument

	err := valueNewNumberFunction_Set()
	if err == nil {
		ret = valueNewNumberFunction.Invoke(inArgs[:], nil)
	}

	retGo := ValueNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'jsc_value_new_object' : parameter 'instance' of type 'gpointer' not supported

var valueNewStringFunction *gi.Function
var valueNewStringFunction_Once sync.Once

func valueNewStringFunction_Set() error {
	var err error
	valueNewStringFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueNewStringFunction, err = valueObject.InvokerNew("new_string")
	})
	return err
}

// ValueNewString is a representation of the C type jsc_value_new_string.
func ValueNewString(context *Context, string_ string) *Value {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(context.Native())
	inArgs[1].SetString(string_)

	var ret gi.Argument

	err := valueNewStringFunction_Set()
	if err == nil {
		ret = valueNewStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ValueNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var valueNewStringFromBytesFunction *gi.Function
var valueNewStringFromBytesFunction_Once sync.Once

func valueNewStringFromBytesFunction_Set() error {
	var err error
	valueNewStringFromBytesFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueNewStringFromBytesFunction, err = valueObject.InvokerNew("new_string_from_bytes")
	})
	return err
}

// ValueNewStringFromBytes is a representation of the C type jsc_value_new_string_from_bytes.
func ValueNewStringFromBytes(context *Context, bytes *glib.Bytes) *Value {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(context.Native())
	inArgs[1].SetPointer(bytes.Native())

	var ret gi.Argument

	err := valueNewStringFromBytesFunction_Set()
	if err == nil {
		ret = valueNewStringFromBytesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ValueNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var valueNewUndefinedFunction *gi.Function
var valueNewUndefinedFunction_Once sync.Once

func valueNewUndefinedFunction_Set() error {
	var err error
	valueNewUndefinedFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueNewUndefinedFunction, err = valueObject.InvokerNew("new_undefined")
	})
	return err
}

// ValueNewUndefined is a representation of the C type jsc_value_new_undefined.
func ValueNewUndefined(context *Context) *Value {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(context.Native())

	var ret gi.Argument

	err := valueNewUndefinedFunction_Set()
	if err == nil {
		ret = valueNewUndefinedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ValueNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'jsc_value_constructor_call' : parameter 'first_parameter_type' of type 'GType' not supported

// UNSUPPORTED : C value 'jsc_value_constructor_callv' : parameter 'parameters' of type 'nil' not supported

// UNSUPPORTED : C value 'jsc_value_function_call' : parameter 'first_parameter_type' of type 'GType' not supported

// UNSUPPORTED : C value 'jsc_value_function_callv' : parameter 'parameters' of type 'nil' not supported

var valueGetContextFunction *gi.Function
var valueGetContextFunction_Once sync.Once

func valueGetContextFunction_Set() error {
	var err error
	valueGetContextFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueGetContextFunction, err = valueObject.InvokerNew("get_context")
	})
	return err
}

// GetContext is a representation of the C type jsc_value_get_context.
func (recv *Value) GetContext() *Context {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := valueGetContextFunction_Set()
	if err == nil {
		ret = valueGetContextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ContextNewFromNative(ret.Pointer())

	return retGo
}

var valueIsArrayFunction *gi.Function
var valueIsArrayFunction_Once sync.Once

func valueIsArrayFunction_Set() error {
	var err error
	valueIsArrayFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueIsArrayFunction, err = valueObject.InvokerNew("is_array")
	})
	return err
}

// IsArray is a representation of the C type jsc_value_is_array.
func (recv *Value) IsArray() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := valueIsArrayFunction_Set()
	if err == nil {
		ret = valueIsArrayFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var valueIsBooleanFunction *gi.Function
var valueIsBooleanFunction_Once sync.Once

func valueIsBooleanFunction_Set() error {
	var err error
	valueIsBooleanFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueIsBooleanFunction, err = valueObject.InvokerNew("is_boolean")
	})
	return err
}

// IsBoolean is a representation of the C type jsc_value_is_boolean.
func (recv *Value) IsBoolean() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := valueIsBooleanFunction_Set()
	if err == nil {
		ret = valueIsBooleanFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var valueIsConstructorFunction *gi.Function
var valueIsConstructorFunction_Once sync.Once

func valueIsConstructorFunction_Set() error {
	var err error
	valueIsConstructorFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueIsConstructorFunction, err = valueObject.InvokerNew("is_constructor")
	})
	return err
}

// IsConstructor is a representation of the C type jsc_value_is_constructor.
func (recv *Value) IsConstructor() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := valueIsConstructorFunction_Set()
	if err == nil {
		ret = valueIsConstructorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var valueIsFunctionFunction *gi.Function
var valueIsFunctionFunction_Once sync.Once

func valueIsFunctionFunction_Set() error {
	var err error
	valueIsFunctionFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueIsFunctionFunction, err = valueObject.InvokerNew("is_function")
	})
	return err
}

// IsFunction is a representation of the C type jsc_value_is_function.
func (recv *Value) IsFunction() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := valueIsFunctionFunction_Set()
	if err == nil {
		ret = valueIsFunctionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var valueIsNullFunction *gi.Function
var valueIsNullFunction_Once sync.Once

func valueIsNullFunction_Set() error {
	var err error
	valueIsNullFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueIsNullFunction, err = valueObject.InvokerNew("is_null")
	})
	return err
}

// IsNull is a representation of the C type jsc_value_is_null.
func (recv *Value) IsNull() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := valueIsNullFunction_Set()
	if err == nil {
		ret = valueIsNullFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var valueIsNumberFunction *gi.Function
var valueIsNumberFunction_Once sync.Once

func valueIsNumberFunction_Set() error {
	var err error
	valueIsNumberFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueIsNumberFunction, err = valueObject.InvokerNew("is_number")
	})
	return err
}

// IsNumber is a representation of the C type jsc_value_is_number.
func (recv *Value) IsNumber() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := valueIsNumberFunction_Set()
	if err == nil {
		ret = valueIsNumberFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var valueIsObjectFunction *gi.Function
var valueIsObjectFunction_Once sync.Once

func valueIsObjectFunction_Set() error {
	var err error
	valueIsObjectFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueIsObjectFunction, err = valueObject.InvokerNew("is_object")
	})
	return err
}

// IsObject is a representation of the C type jsc_value_is_object.
func (recv *Value) IsObject() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := valueIsObjectFunction_Set()
	if err == nil {
		ret = valueIsObjectFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var valueIsStringFunction *gi.Function
var valueIsStringFunction_Once sync.Once

func valueIsStringFunction_Set() error {
	var err error
	valueIsStringFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueIsStringFunction, err = valueObject.InvokerNew("is_string")
	})
	return err
}

// IsString is a representation of the C type jsc_value_is_string.
func (recv *Value) IsString() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := valueIsStringFunction_Set()
	if err == nil {
		ret = valueIsStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var valueIsUndefinedFunction *gi.Function
var valueIsUndefinedFunction_Once sync.Once

func valueIsUndefinedFunction_Set() error {
	var err error
	valueIsUndefinedFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueIsUndefinedFunction, err = valueObject.InvokerNew("is_undefined")
	})
	return err
}

// IsUndefined is a representation of the C type jsc_value_is_undefined.
func (recv *Value) IsUndefined() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := valueIsUndefinedFunction_Set()
	if err == nil {
		ret = valueIsUndefinedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'jsc_value_object_define_property_accessor' : parameter 'flags' of type 'ValuePropertyFlags' not supported

// UNSUPPORTED : C value 'jsc_value_object_define_property_data' : parameter 'flags' of type 'ValuePropertyFlags' not supported

var valueObjectDeletePropertyFunction *gi.Function
var valueObjectDeletePropertyFunction_Once sync.Once

func valueObjectDeletePropertyFunction_Set() error {
	var err error
	valueObjectDeletePropertyFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueObjectDeletePropertyFunction, err = valueObject.InvokerNew("object_delete_property")
	})
	return err
}

// ObjectDeleteProperty is a representation of the C type jsc_value_object_delete_property.
func (recv *Value) ObjectDeleteProperty(name string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := valueObjectDeletePropertyFunction_Set()
	if err == nil {
		ret = valueObjectDeletePropertyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var valueObjectEnumeratePropertiesFunction *gi.Function
var valueObjectEnumeratePropertiesFunction_Once sync.Once

func valueObjectEnumeratePropertiesFunction_Set() error {
	var err error
	valueObjectEnumeratePropertiesFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueObjectEnumeratePropertiesFunction, err = valueObject.InvokerNew("object_enumerate_properties")
	})
	return err
}

// ObjectEnumerateProperties is a representation of the C type jsc_value_object_enumerate_properties.
func (recv *Value) ObjectEnumerateProperties() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := valueObjectEnumeratePropertiesFunction_Set()
	if err == nil {
		valueObjectEnumeratePropertiesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var valueObjectGetPropertyFunction *gi.Function
var valueObjectGetPropertyFunction_Once sync.Once

func valueObjectGetPropertyFunction_Set() error {
	var err error
	valueObjectGetPropertyFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueObjectGetPropertyFunction, err = valueObject.InvokerNew("object_get_property")
	})
	return err
}

// ObjectGetProperty is a representation of the C type jsc_value_object_get_property.
func (recv *Value) ObjectGetProperty(name string) *Value {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := valueObjectGetPropertyFunction_Set()
	if err == nil {
		ret = valueObjectGetPropertyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ValueNewFromNative(ret.Pointer())

	return retGo
}

var valueObjectGetPropertyAtIndexFunction *gi.Function
var valueObjectGetPropertyAtIndexFunction_Once sync.Once

func valueObjectGetPropertyAtIndexFunction_Set() error {
	var err error
	valueObjectGetPropertyAtIndexFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueObjectGetPropertyAtIndexFunction, err = valueObject.InvokerNew("object_get_property_at_index")
	})
	return err
}

// ObjectGetPropertyAtIndex is a representation of the C type jsc_value_object_get_property_at_index.
func (recv *Value) ObjectGetPropertyAtIndex(index uint32) *Value {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(index)

	var ret gi.Argument

	err := valueObjectGetPropertyAtIndexFunction_Set()
	if err == nil {
		ret = valueObjectGetPropertyAtIndexFunction.Invoke(inArgs[:], nil)
	}

	retGo := ValueNewFromNative(ret.Pointer())

	return retGo
}

var valueObjectHasPropertyFunction *gi.Function
var valueObjectHasPropertyFunction_Once sync.Once

func valueObjectHasPropertyFunction_Set() error {
	var err error
	valueObjectHasPropertyFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueObjectHasPropertyFunction, err = valueObject.InvokerNew("object_has_property")
	})
	return err
}

// ObjectHasProperty is a representation of the C type jsc_value_object_has_property.
func (recv *Value) ObjectHasProperty(name string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := valueObjectHasPropertyFunction_Set()
	if err == nil {
		ret = valueObjectHasPropertyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'jsc_value_object_invoke_method' : parameter 'first_parameter_type' of type 'GType' not supported

// UNSUPPORTED : C value 'jsc_value_object_invoke_methodv' : parameter 'parameters' of type 'nil' not supported

var valueObjectIsInstanceOfFunction *gi.Function
var valueObjectIsInstanceOfFunction_Once sync.Once

func valueObjectIsInstanceOfFunction_Set() error {
	var err error
	valueObjectIsInstanceOfFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueObjectIsInstanceOfFunction, err = valueObject.InvokerNew("object_is_instance_of")
	})
	return err
}

// ObjectIsInstanceOf is a representation of the C type jsc_value_object_is_instance_of.
func (recv *Value) ObjectIsInstanceOf(name string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := valueObjectIsInstanceOfFunction_Set()
	if err == nil {
		ret = valueObjectIsInstanceOfFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var valueObjectSetPropertyFunction *gi.Function
var valueObjectSetPropertyFunction_Once sync.Once

func valueObjectSetPropertyFunction_Set() error {
	var err error
	valueObjectSetPropertyFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueObjectSetPropertyFunction, err = valueObject.InvokerNew("object_set_property")
	})
	return err
}

// ObjectSetProperty is a representation of the C type jsc_value_object_set_property.
func (recv *Value) ObjectSetProperty(name string, property *Value) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)
	inArgs[2].SetPointer(property.Native())

	err := valueObjectSetPropertyFunction_Set()
	if err == nil {
		valueObjectSetPropertyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var valueObjectSetPropertyAtIndexFunction *gi.Function
var valueObjectSetPropertyAtIndexFunction_Once sync.Once

func valueObjectSetPropertyAtIndexFunction_Set() error {
	var err error
	valueObjectSetPropertyAtIndexFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueObjectSetPropertyAtIndexFunction, err = valueObject.InvokerNew("object_set_property_at_index")
	})
	return err
}

// ObjectSetPropertyAtIndex is a representation of the C type jsc_value_object_set_property_at_index.
func (recv *Value) ObjectSetPropertyAtIndex(index uint32, property *Value) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(index)
	inArgs[2].SetPointer(property.Native())

	err := valueObjectSetPropertyAtIndexFunction_Set()
	if err == nil {
		valueObjectSetPropertyAtIndexFunction.Invoke(inArgs[:], nil)
	}

	return
}

var valueToBooleanFunction *gi.Function
var valueToBooleanFunction_Once sync.Once

func valueToBooleanFunction_Set() error {
	var err error
	valueToBooleanFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueToBooleanFunction, err = valueObject.InvokerNew("to_boolean")
	})
	return err
}

// ToBoolean is a representation of the C type jsc_value_to_boolean.
func (recv *Value) ToBoolean() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := valueToBooleanFunction_Set()
	if err == nil {
		ret = valueToBooleanFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var valueToDoubleFunction *gi.Function
var valueToDoubleFunction_Once sync.Once

func valueToDoubleFunction_Set() error {
	var err error
	valueToDoubleFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueToDoubleFunction, err = valueObject.InvokerNew("to_double")
	})
	return err
}

// ToDouble is a representation of the C type jsc_value_to_double.
func (recv *Value) ToDouble() float64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := valueToDoubleFunction_Set()
	if err == nil {
		ret = valueToDoubleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo
}

var valueToInt32Function *gi.Function
var valueToInt32Function_Once sync.Once

func valueToInt32Function_Set() error {
	var err error
	valueToInt32Function_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueToInt32Function, err = valueObject.InvokerNew("to_int32")
	})
	return err
}

// ToInt32 is a representation of the C type jsc_value_to_int32.
func (recv *Value) ToInt32() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := valueToInt32Function_Set()
	if err == nil {
		ret = valueToInt32Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var valueToStringFunction *gi.Function
var valueToStringFunction_Once sync.Once

func valueToStringFunction_Set() error {
	var err error
	valueToStringFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueToStringFunction, err = valueObject.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type jsc_value_to_string.
func (recv *Value) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := valueToStringFunction_Set()
	if err == nil {
		ret = valueToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var valueToStringAsBytesFunction *gi.Function
var valueToStringAsBytesFunction_Once sync.Once

func valueToStringAsBytesFunction_Set() error {
	var err error
	valueToStringAsBytesFunction_Once.Do(func() {
		err = valueObject_Set()
		if err != nil {
			return
		}
		valueToStringAsBytesFunction, err = valueObject.InvokerNew("to_string_as_bytes")
	})
	return err
}

// ToStringAsBytes is a representation of the C type jsc_value_to_string_as_bytes.
func (recv *Value) ToStringAsBytes() *glib.Bytes {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := valueToStringAsBytesFunction_Set()
	if err == nil {
		ret = valueToStringAsBytesFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.BytesNewFromNative(ret.Pointer())

	return retGo
}

var virtualMachineObject *gi.Object
var virtualMachineObject_Once sync.Once

func virtualMachineObject_Set() error {
	var err error
	virtualMachineObject_Once.Do(func() {
		virtualMachineObject, err = gi.ObjectNew("JavaScriptCore", "VirtualMachine")
	})
	return err
}

type VirtualMachine struct {
	native unsafe.Pointer
}

func VirtualMachineNewFromNative(native unsafe.Pointer) *VirtualMachine {
	instance := &VirtualMachine{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *VirtualMachine) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToVirtualMachine down casts any arbitrary Object to VirtualMachine.
Exercise care, as this is a potentially dangerous function
if the Object is not a VirtualMachine.
*/
func CastToVirtualMachine(object *gobject.Object) *VirtualMachine {
	return VirtualMachineNewFromNative(object.Native())
}

// Equals compares this VirtualMachine with another VirtualMachine, and returns true if they represent the same Object.
func (recv *VirtualMachine) Equals(other *VirtualMachine) bool {
	return other.Native() == recv.Native()
}

func (recv *VirtualMachine) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *VirtualMachine) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(virtualMachineObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *VirtualMachine) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(virtualMachineObject, recv.Native(), "parent", argValue)
}

var virtualMachineNewFunction *gi.Function
var virtualMachineNewFunction_Once sync.Once

func virtualMachineNewFunction_Set() error {
	var err error
	virtualMachineNewFunction_Once.Do(func() {
		err = virtualMachineObject_Set()
		if err != nil {
			return
		}
		virtualMachineNewFunction, err = virtualMachineObject.InvokerNew("new")
	})
	return err
}

// VirtualMachineNew is a representation of the C type jsc_virtual_machine_new.
func VirtualMachineNew() *VirtualMachine {

	var ret gi.Argument

	err := virtualMachineNewFunction_Set()
	if err == nil {
		ret = virtualMachineNewFunction.Invoke(nil, nil)
	}

	retGo := VirtualMachineNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var weakValueObject *gi.Object
var weakValueObject_Once sync.Once

func weakValueObject_Set() error {
	var err error
	weakValueObject_Once.Do(func() {
		weakValueObject, err = gi.ObjectNew("JavaScriptCore", "WeakValue")
	})
	return err
}

type WeakValue struct {
	native unsafe.Pointer
}

func WeakValueNewFromNative(native unsafe.Pointer) *WeakValue {
	instance := &WeakValue{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *WeakValue) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToWeakValue down casts any arbitrary Object to WeakValue.
Exercise care, as this is a potentially dangerous function
if the Object is not a WeakValue.
*/
func CastToWeakValue(object *gobject.Object) *WeakValue {
	return WeakValueNewFromNative(object.Native())
}

// Equals compares this WeakValue with another WeakValue, and returns true if they represent the same Object.
func (recv *WeakValue) Equals(other *WeakValue) bool {
	return other.Native() == recv.Native()
}

func (recv *WeakValue) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *WeakValue) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(weakValueObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *WeakValue) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(weakValueObject, recv.Native(), "parent", argValue)
}

var weakValueNewFunction *gi.Function
var weakValueNewFunction_Once sync.Once

func weakValueNewFunction_Set() error {
	var err error
	weakValueNewFunction_Once.Do(func() {
		err = weakValueObject_Set()
		if err != nil {
			return
		}
		weakValueNewFunction, err = weakValueObject.InvokerNew("new")
	})
	return err
}

// WeakValueNew is a representation of the C type jsc_weak_value_new.
func WeakValueNew(value *Value) *WeakValue {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(value.Native())

	var ret gi.Argument

	err := weakValueNewFunction_Set()
	if err == nil {
		ret = weakValueNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := WeakValueNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var weakValueGetValueFunction *gi.Function
var weakValueGetValueFunction_Once sync.Once

func weakValueGetValueFunction_Set() error {
	var err error
	weakValueGetValueFunction_Once.Do(func() {
		err = weakValueObject_Set()
		if err != nil {
			return
		}
		weakValueGetValueFunction, err = weakValueObject.InvokerNew("get_value")
	})
	return err
}

// GetValue is a representation of the C type jsc_weak_value_get_value.
func (recv *WeakValue) GetValue() *Value {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := weakValueGetValueFunction_Set()
	if err == nil {
		ret = weakValueGetValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := ValueNewFromNative(ret.Pointer())

	return retGo
}

/*
ConnectCleared connects a callback to the 'cleared' signal of the WeakValue.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *WeakValue) ConnectCleared(handler func(instance *WeakValue)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := WeakValueNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "cleared", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *WeakValue) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}
