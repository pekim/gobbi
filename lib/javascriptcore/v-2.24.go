// Code generated - DO NOT EDIT.
// +build javascriptcore_2.24

package javascriptcore

import (
	"fmt"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <jsc/jsc.h>
// #include <stdlib.h>
/*

	static void _jsc_context_throw_printf(JSCContext* context, const char* format) {
		return jsc_context_throw_printf(context, format);
    }
*/
/*

	static void _jsc_context_throw_with_name_printf(JSCContext* context, const char* error_name, const char* format) {
		return jsc_context_throw_with_name_printf(context, error_name, format);
    }
*/
/*

	static JSCException* _jsc_exception_new_printf(JSCContext* context, const char* format) {
		return jsc_exception_new_printf(context, format);
    }
*/
/*

	static JSCException* _jsc_exception_new_with_name_printf(JSCContext* context, const char* name, const char* format) {
		return jsc_exception_new_with_name_printf(context, name, format);
    }
*/
/*

	void weakvalue_clearedHandler(GObject *, gpointer);

	static gulong WeakValue_signal_connect_cleared(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "cleared", G_CALLBACK(weakvalue_clearedHandler), data);
	}

*/
import "C"

type ValuePropertyFlags C.JSCValuePropertyFlags

const (
	VALUE_PROPERTY_CONFIGURABLE ValuePropertyFlags = 1
	VALUE_PROPERTY_ENUMERABLE   ValuePropertyFlags = 2
	VALUE_PROPERTY_WRITABLE     ValuePropertyFlags = 4
)

// Class is a wrapper around the C record JSCClass.
type Class struct {
	native *C.JSCClass
}

func ClassNewFromC(u unsafe.Pointer) *Class {
	c := (*C.JSCClass)(u)
	if c == nil {
		return nil
	}

	g := &Class{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Class) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Class) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Class with another Class, and returns true if they represent the same GObject.
func (recv *Class) Equals(other *Class) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Class) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Class.
// Exercise care, as this is a potentially dangerous function if the Object is not a Class.
func CastToClass(object *gobject.Object) *Class {
	return ClassNewFromC(object.ToC())
}

// Unsupported : jsc_class_add_constructor : unsupported parameter callback : no type generator for GObject.Callback (GCallback) for param callback

// Unsupported : jsc_class_add_constructor_variadic : unsupported parameter callback : no type generator for GObject.Callback (GCallback) for param callback

// Unsupported : jsc_class_add_constructorv : unsupported parameter callback : no type generator for GObject.Callback (GCallback) for param callback

// Unsupported : jsc_class_add_method : unsupported parameter callback : no type generator for GObject.Callback (GCallback) for param callback

// Unsupported : jsc_class_add_method_variadic : unsupported parameter callback : no type generator for GObject.Callback (GCallback) for param callback

// Unsupported : jsc_class_add_methodv : unsupported parameter callback : no type generator for GObject.Callback (GCallback) for param callback

// Unsupported : jsc_class_add_property : unsupported parameter getter : no type generator for GObject.Callback (GCallback) for param getter

// GetName is a wrapper around the C function jsc_class_get_name.
func (recv *Class) GetName() string {
	retC := C.jsc_class_get_name((*C.JSCClass)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetParent is a wrapper around the C function jsc_class_get_parent.
func (recv *Class) GetParent() *Class {
	retC := C.jsc_class_get_parent((*C.JSCClass)(recv.native))
	retGo := ClassNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Context is a wrapper around the C record JSCContext.
type Context struct {
	native *C.JSCContext
	// parent : record
	// Private : priv
}

func ContextNewFromC(u unsafe.Pointer) *Context {
	c := (*C.JSCContext)(u)
	if c == nil {
		return nil
	}

	g := &Context{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Context) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Context) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Context with another Context, and returns true if they represent the same GObject.
func (recv *Context) Equals(other *Context) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Context) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Context.
// Exercise care, as this is a potentially dangerous function if the Object is not a Context.
func CastToContext(object *gobject.Object) *Context {
	return ContextNewFromC(object.ToC())
}

// ContextNew is a wrapper around the C function jsc_context_new.
func ContextNew() *Context {
	retC := C.jsc_context_new()
	retGo := ContextNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ContextNewWithVirtualMachine is a wrapper around the C function jsc_context_new_with_virtual_machine.
func ContextNewWithVirtualMachine(vm *VirtualMachine) *Context {
	c_vm := (*C.JSCVirtualMachine)(C.NULL)
	if vm != nil {
		c_vm = (*C.JSCVirtualMachine)(vm.ToC())
	}

	retC := C.jsc_context_new_with_virtual_machine(c_vm)
	retGo := ContextNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ContextGetCurrent is a wrapper around the C function jsc_context_get_current.
func ContextGetCurrent() *Context {
	retC := C.jsc_context_get_current()
	var retGo (*Context)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ContextNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// CheckSyntax is a wrapper around the C function jsc_context_check_syntax.
func (recv *Context) CheckSyntax(code string, mode CheckSyntaxMode, uri string, lineNumber uint32) (CheckSyntaxResult, *Exception) {
	c_code := C.CString(code)
	defer C.free(unsafe.Pointer(c_code))

	c_length := (C.gssize)(len(code))

	c_mode := (C.JSCCheckSyntaxMode)(mode)

	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_line_number := (C.unsigned)(lineNumber)

	var c_exception *C.JSCException

	retC := C.jsc_context_check_syntax((*C.JSCContext)(recv.native), c_code, c_length, c_mode, c_uri, c_line_number, &c_exception)
	retGo := (CheckSyntaxResult)(retC)

	exception := ExceptionNewFromC(unsafe.Pointer(c_exception))

	return retGo, exception
}

// ClearException is a wrapper around the C function jsc_context_clear_exception.
func (recv *Context) ClearException() {
	C.jsc_context_clear_exception((*C.JSCContext)(recv.native))

	return
}

// Evaluate is a wrapper around the C function jsc_context_evaluate.
func (recv *Context) Evaluate(code string) *Value {
	c_code := C.CString(code)
	defer C.free(unsafe.Pointer(c_code))

	c_length := (C.gssize)(len(code))

	retC := C.jsc_context_evaluate((*C.JSCContext)(recv.native), c_code, c_length)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EvaluateInObject is a wrapper around the C function jsc_context_evaluate_in_object.
func (recv *Context) EvaluateInObject(code string, objectInstance uintptr, objectClass *Class, uri string, lineNumber uint32) (*Value, *Value) {
	c_code := C.CString(code)
	defer C.free(unsafe.Pointer(c_code))

	c_length := (C.gssize)(len(code))

	c_object_instance := (C.gpointer)(objectInstance)

	c_object_class := (*C.JSCClass)(C.NULL)
	if objectClass != nil {
		c_object_class = (*C.JSCClass)(objectClass.ToC())
	}

	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_line_number := (C.guint)(lineNumber)

	var c_object *C.JSCValue

	retC := C.jsc_context_evaluate_in_object((*C.JSCContext)(recv.native), c_code, c_length, c_object_instance, c_object_class, c_uri, c_line_number, &c_object)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	object := ValueNewFromC(unsafe.Pointer(c_object))

	return retGo, object
}

// EvaluateWithSourceUri is a wrapper around the C function jsc_context_evaluate_with_source_uri.
func (recv *Context) EvaluateWithSourceUri(code string, uri string, lineNumber uint32) *Value {
	c_code := C.CString(code)
	defer C.free(unsafe.Pointer(c_code))

	c_length := (C.gssize)(len(code))

	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_line_number := (C.guint)(lineNumber)

	retC := C.jsc_context_evaluate_with_source_uri((*C.JSCContext)(recv.native), c_code, c_length, c_uri, c_line_number)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetException is a wrapper around the C function jsc_context_get_exception.
func (recv *Context) GetException() *Exception {
	retC := C.jsc_context_get_exception((*C.JSCContext)(recv.native))
	var retGo (*Exception)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ExceptionNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetGlobalObject is a wrapper around the C function jsc_context_get_global_object.
func (recv *Context) GetGlobalObject() *Value {
	retC := C.jsc_context_get_global_object((*C.JSCContext)(recv.native))
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetValue is a wrapper around the C function jsc_context_get_value.
func (recv *Context) GetValue(name string) *Value {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.jsc_context_get_value((*C.JSCContext)(recv.native), c_name)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetVirtualMachine is a wrapper around the C function jsc_context_get_virtual_machine.
func (recv *Context) GetVirtualMachine() *VirtualMachine {
	retC := C.jsc_context_get_virtual_machine((*C.JSCContext)(recv.native))
	retGo := VirtualMachineNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PopExceptionHandler is a wrapper around the C function jsc_context_pop_exception_handler.
func (recv *Context) PopExceptionHandler() {
	C.jsc_context_pop_exception_handler((*C.JSCContext)(recv.native))

	return
}

// Unsupported : jsc_context_push_exception_handler : unsupported parameter handler : no type generator for ExceptionHandler (JSCExceptionHandler) for param handler

// Unsupported : jsc_context_register_class : unsupported parameter destroy_notify : no type generator for GLib.DestroyNotify (GDestroyNotify) for param destroy_notify

// SetValue is a wrapper around the C function jsc_context_set_value.
func (recv *Context) SetValue(name string, value *Value) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := (*C.JSCValue)(C.NULL)
	if value != nil {
		c_value = (*C.JSCValue)(value.ToC())
	}

	C.jsc_context_set_value((*C.JSCContext)(recv.native), c_name, c_value)

	return
}

// Throw is a wrapper around the C function jsc_context_throw.
func (recv *Context) Throw(errorMessage string) {
	c_error_message := C.CString(errorMessage)
	defer C.free(unsafe.Pointer(c_error_message))

	C.jsc_context_throw((*C.JSCContext)(recv.native), c_error_message)

	return
}

// ThrowException is a wrapper around the C function jsc_context_throw_exception.
func (recv *Context) ThrowException(exception *Exception) {
	c_exception := (*C.JSCException)(C.NULL)
	if exception != nil {
		c_exception = (*C.JSCException)(exception.ToC())
	}

	C.jsc_context_throw_exception((*C.JSCContext)(recv.native), c_exception)

	return
}

// ThrowPrintf is a wrapper around the C function jsc_context_throw_printf.
func (recv *Context) ThrowPrintf(format string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._jsc_context_throw_printf((*C.JSCContext)(recv.native), c_format)

	return
}

// ThrowWithName is a wrapper around the C function jsc_context_throw_with_name.
func (recv *Context) ThrowWithName(errorName string, errorMessage string) {
	c_error_name := C.CString(errorName)
	defer C.free(unsafe.Pointer(c_error_name))

	c_error_message := C.CString(errorMessage)
	defer C.free(unsafe.Pointer(c_error_message))

	C.jsc_context_throw_with_name((*C.JSCContext)(recv.native), c_error_name, c_error_message)

	return
}

// ThrowWithNamePrintf is a wrapper around the C function jsc_context_throw_with_name_printf.
func (recv *Context) ThrowWithNamePrintf(errorName string, format string, args ...interface{}) {
	c_error_name := C.CString(errorName)
	defer C.free(unsafe.Pointer(c_error_name))

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._jsc_context_throw_with_name_printf((*C.JSCContext)(recv.native), c_error_name, c_format)

	return
}

// Exception is a wrapper around the C record JSCException.
type Exception struct {
	native *C.JSCException
	// parent : record
	// Private : priv
}

func ExceptionNewFromC(u unsafe.Pointer) *Exception {
	c := (*C.JSCException)(u)
	if c == nil {
		return nil
	}

	g := &Exception{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Exception) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Exception) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Exception with another Exception, and returns true if they represent the same GObject.
func (recv *Exception) Equals(other *Exception) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Exception) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Exception.
// Exercise care, as this is a potentially dangerous function if the Object is not a Exception.
func CastToException(object *gobject.Object) *Exception {
	return ExceptionNewFromC(object.ToC())
}

// ExceptionNew is a wrapper around the C function jsc_exception_new.
func ExceptionNew(context *Context, message string) *Exception {
	c_context := (*C.JSCContext)(C.NULL)
	if context != nil {
		c_context = (*C.JSCContext)(context.ToC())
	}

	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	retC := C.jsc_exception_new(c_context, c_message)
	retGo := ExceptionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ExceptionNewPrintf is a wrapper around the C function jsc_exception_new_printf.
func ExceptionNewPrintf(context *Context, format string, args ...interface{}) *Exception {
	c_context := (*C.JSCContext)(C.NULL)
	if context != nil {
		c_context = (*C.JSCContext)(context.ToC())
	}

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	retC := C._jsc_exception_new_printf(c_context, c_format)
	retGo := ExceptionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : jsc_exception_new_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// ExceptionNewWithName is a wrapper around the C function jsc_exception_new_with_name.
func ExceptionNewWithName(context *Context, name string, message string) *Exception {
	c_context := (*C.JSCContext)(C.NULL)
	if context != nil {
		c_context = (*C.JSCContext)(context.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	retC := C.jsc_exception_new_with_name(c_context, c_name, c_message)
	retGo := ExceptionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ExceptionNewWithNamePrintf is a wrapper around the C function jsc_exception_new_with_name_printf.
func ExceptionNewWithNamePrintf(context *Context, name string, format string, args ...interface{}) *Exception {
	c_context := (*C.JSCContext)(C.NULL)
	if context != nil {
		c_context = (*C.JSCContext)(context.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	retC := C._jsc_exception_new_with_name_printf(c_context, c_name, c_format)
	retGo := ExceptionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : jsc_exception_new_with_name_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// GetBacktraceString is a wrapper around the C function jsc_exception_get_backtrace_string.
func (recv *Exception) GetBacktraceString() string {
	retC := C.jsc_exception_get_backtrace_string((*C.JSCException)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetColumnNumber is a wrapper around the C function jsc_exception_get_column_number.
func (recv *Exception) GetColumnNumber() uint32 {
	retC := C.jsc_exception_get_column_number((*C.JSCException)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetLineNumber is a wrapper around the C function jsc_exception_get_line_number.
func (recv *Exception) GetLineNumber() uint32 {
	retC := C.jsc_exception_get_line_number((*C.JSCException)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetMessage is a wrapper around the C function jsc_exception_get_message.
func (recv *Exception) GetMessage() string {
	retC := C.jsc_exception_get_message((*C.JSCException)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetName is a wrapper around the C function jsc_exception_get_name.
func (recv *Exception) GetName() string {
	retC := C.jsc_exception_get_name((*C.JSCException)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetSourceUri is a wrapper around the C function jsc_exception_get_source_uri.
func (recv *Exception) GetSourceUri() string {
	retC := C.jsc_exception_get_source_uri((*C.JSCException)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Report is a wrapper around the C function jsc_exception_report.
func (recv *Exception) Report() string {
	retC := C.jsc_exception_report((*C.JSCException)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ToString is a wrapper around the C function jsc_exception_to_string.
func (recv *Exception) ToString() string {
	retC := C.jsc_exception_to_string((*C.JSCException)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Value is a wrapper around the C record JSCValue.
type Value struct {
	native *C.JSCValue
	// parent : record
	// Private : priv
}

func ValueNewFromC(u unsafe.Pointer) *Value {
	c := (*C.JSCValue)(u)
	if c == nil {
		return nil
	}

	g := &Value{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Value) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Value) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Value with another Value, and returns true if they represent the same GObject.
func (recv *Value) Equals(other *Value) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Value) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Value.
// Exercise care, as this is a potentially dangerous function if the Object is not a Value.
func CastToValue(object *gobject.Object) *Value {
	return ValueNewFromC(object.ToC())
}

// Unsupported : jsc_value_new_array : unsupported parameter ... : varargs

// Unsupported : jsc_value_new_array_from_garray : unsupported parameter array :

// ValueNewArrayFromStrv is a wrapper around the C function jsc_value_new_array_from_strv.
func ValueNewArrayFromStrv(context *Context, strv []string) *Value {
	c_context := (*C.JSCContext)(C.NULL)
	if context != nil {
		c_context = (*C.JSCContext)(context.ToC())
	}

	c_strv_array := make([]*C.gchar, len(strv)+1, len(strv)+1)
	for i, item := range strv {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_strv_array[i] = c
	}
	c_strv_array[len(strv)] = nil
	c_strv_arrayPtr := &c_strv_array[0]
	c_strv := (**C.char)(unsafe.Pointer(c_strv_arrayPtr))

	retC := C.jsc_value_new_array_from_strv(c_context, c_strv)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ValueNewBoolean is a wrapper around the C function jsc_value_new_boolean.
func ValueNewBoolean(context *Context, value bool) *Value {
	c_context := (*C.JSCContext)(C.NULL)
	if context != nil {
		c_context = (*C.JSCContext)(context.ToC())
	}

	c_value :=
		boolToGboolean(value)

	retC := C.jsc_value_new_boolean(c_context, c_value)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : jsc_value_new_function : unsupported parameter callback : no type generator for GObject.Callback (GCallback) for param callback

// Unsupported : jsc_value_new_function_variadic : unsupported parameter callback : no type generator for GObject.Callback (GCallback) for param callback

// Unsupported : jsc_value_new_functionv : unsupported parameter callback : no type generator for GObject.Callback (GCallback) for param callback

// ValueNewNull is a wrapper around the C function jsc_value_new_null.
func ValueNewNull(context *Context) *Value {
	c_context := (*C.JSCContext)(C.NULL)
	if context != nil {
		c_context = (*C.JSCContext)(context.ToC())
	}

	retC := C.jsc_value_new_null(c_context)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ValueNewNumber is a wrapper around the C function jsc_value_new_number.
func ValueNewNumber(context *Context, number float64) *Value {
	c_context := (*C.JSCContext)(C.NULL)
	if context != nil {
		c_context = (*C.JSCContext)(context.ToC())
	}

	c_number := (C.double)(number)

	retC := C.jsc_value_new_number(c_context, c_number)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ValueNewObject is a wrapper around the C function jsc_value_new_object.
func ValueNewObject(context *Context, instance uintptr, jscClass *Class) *Value {
	c_context := (*C.JSCContext)(C.NULL)
	if context != nil {
		c_context = (*C.JSCContext)(context.ToC())
	}

	c_instance := (C.gpointer)(instance)

	c_jsc_class := (*C.JSCClass)(C.NULL)
	if jscClass != nil {
		c_jsc_class = (*C.JSCClass)(jscClass.ToC())
	}

	retC := C.jsc_value_new_object(c_context, c_instance, c_jsc_class)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ValueNewString is a wrapper around the C function jsc_value_new_string.
func ValueNewString(context *Context, string_ string) *Value {
	c_context := (*C.JSCContext)(C.NULL)
	if context != nil {
		c_context = (*C.JSCContext)(context.ToC())
	}

	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.jsc_value_new_string(c_context, c_string)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ValueNewUndefined is a wrapper around the C function jsc_value_new_undefined.
func ValueNewUndefined(context *Context) *Value {
	c_context := (*C.JSCContext)(C.NULL)
	if context != nil {
		c_context = (*C.JSCContext)(context.ToC())
	}

	retC := C.jsc_value_new_undefined(c_context)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : jsc_value_constructor_call : unsupported parameter ... : varargs

// Unsupported : jsc_value_constructor_callv : unsupported parameter parameters :

// Unsupported : jsc_value_function_call : unsupported parameter ... : varargs

// Unsupported : jsc_value_function_callv : unsupported parameter parameters :

// GetContext is a wrapper around the C function jsc_value_get_context.
func (recv *Value) GetContext() *Context {
	retC := C.jsc_value_get_context((*C.JSCValue)(recv.native))
	retGo := ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsArray is a wrapper around the C function jsc_value_is_array.
func (recv *Value) IsArray() bool {
	retC := C.jsc_value_is_array((*C.JSCValue)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsBoolean is a wrapper around the C function jsc_value_is_boolean.
func (recv *Value) IsBoolean() bool {
	retC := C.jsc_value_is_boolean((*C.JSCValue)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsConstructor is a wrapper around the C function jsc_value_is_constructor.
func (recv *Value) IsConstructor() bool {
	retC := C.jsc_value_is_constructor((*C.JSCValue)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsFunction is a wrapper around the C function jsc_value_is_function.
func (recv *Value) IsFunction() bool {
	retC := C.jsc_value_is_function((*C.JSCValue)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsNull is a wrapper around the C function jsc_value_is_null.
func (recv *Value) IsNull() bool {
	retC := C.jsc_value_is_null((*C.JSCValue)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsNumber is a wrapper around the C function jsc_value_is_number.
func (recv *Value) IsNumber() bool {
	retC := C.jsc_value_is_number((*C.JSCValue)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsObject is a wrapper around the C function jsc_value_is_object.
func (recv *Value) IsObject() bool {
	retC := C.jsc_value_is_object((*C.JSCValue)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsString is a wrapper around the C function jsc_value_is_string.
func (recv *Value) IsString() bool {
	retC := C.jsc_value_is_string((*C.JSCValue)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsUndefined is a wrapper around the C function jsc_value_is_undefined.
func (recv *Value) IsUndefined() bool {
	retC := C.jsc_value_is_undefined((*C.JSCValue)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : jsc_value_object_define_property_accessor : unsupported parameter getter : no type generator for GObject.Callback (GCallback) for param getter

// ObjectDefinePropertyData is a wrapper around the C function jsc_value_object_define_property_data.
func (recv *Value) ObjectDefinePropertyData(propertyName string, flags ValuePropertyFlags, propertyValue *Value) {
	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	c_flags := (C.JSCValuePropertyFlags)(flags)

	c_property_value := (*C.JSCValue)(C.NULL)
	if propertyValue != nil {
		c_property_value = (*C.JSCValue)(propertyValue.ToC())
	}

	C.jsc_value_object_define_property_data((*C.JSCValue)(recv.native), c_property_name, c_flags, c_property_value)

	return
}

// ObjectDeleteProperty is a wrapper around the C function jsc_value_object_delete_property.
func (recv *Value) ObjectDeleteProperty(name string) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.jsc_value_object_delete_property((*C.JSCValue)(recv.native), c_name)
	retGo := retC == C.TRUE

	return retGo
}

// ObjectEnumerateProperties is a wrapper around the C function jsc_value_object_enumerate_properties.
func (recv *Value) ObjectEnumerateProperties() []string {
	retC := C.jsc_value_object_enumerate_properties((*C.JSCValue)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// ObjectGetProperty is a wrapper around the C function jsc_value_object_get_property.
func (recv *Value) ObjectGetProperty(name string) *Value {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.jsc_value_object_get_property((*C.JSCValue)(recv.native), c_name)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ObjectGetPropertyAtIndex is a wrapper around the C function jsc_value_object_get_property_at_index.
func (recv *Value) ObjectGetPropertyAtIndex(index uint32) *Value {
	c_index := (C.guint)(index)

	retC := C.jsc_value_object_get_property_at_index((*C.JSCValue)(recv.native), c_index)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ObjectHasProperty is a wrapper around the C function jsc_value_object_has_property.
func (recv *Value) ObjectHasProperty(name string) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.jsc_value_object_has_property((*C.JSCValue)(recv.native), c_name)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : jsc_value_object_invoke_method : unsupported parameter ... : varargs

// Unsupported : jsc_value_object_invoke_methodv : unsupported parameter parameters :

// ObjectIsInstanceOf is a wrapper around the C function jsc_value_object_is_instance_of.
func (recv *Value) ObjectIsInstanceOf(name string) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.jsc_value_object_is_instance_of((*C.JSCValue)(recv.native), c_name)
	retGo := retC == C.TRUE

	return retGo
}

// ObjectSetProperty is a wrapper around the C function jsc_value_object_set_property.
func (recv *Value) ObjectSetProperty(name string, property *Value) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_property := (*C.JSCValue)(C.NULL)
	if property != nil {
		c_property = (*C.JSCValue)(property.ToC())
	}

	C.jsc_value_object_set_property((*C.JSCValue)(recv.native), c_name, c_property)

	return
}

// ObjectSetPropertyAtIndex is a wrapper around the C function jsc_value_object_set_property_at_index.
func (recv *Value) ObjectSetPropertyAtIndex(index uint32, property *Value) {
	c_index := (C.guint)(index)

	c_property := (*C.JSCValue)(C.NULL)
	if property != nil {
		c_property = (*C.JSCValue)(property.ToC())
	}

	C.jsc_value_object_set_property_at_index((*C.JSCValue)(recv.native), c_index, c_property)

	return
}

// ToBoolean is a wrapper around the C function jsc_value_to_boolean.
func (recv *Value) ToBoolean() bool {
	retC := C.jsc_value_to_boolean((*C.JSCValue)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ToDouble is a wrapper around the C function jsc_value_to_double.
func (recv *Value) ToDouble() float64 {
	retC := C.jsc_value_to_double((*C.JSCValue)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// ToInt32 is a wrapper around the C function jsc_value_to_int32.
func (recv *Value) ToInt32() int32 {
	retC := C.jsc_value_to_int32((*C.JSCValue)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// ToString is a wrapper around the C function jsc_value_to_string.
func (recv *Value) ToString() string {
	retC := C.jsc_value_to_string((*C.JSCValue)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ToStringAsBytes is a wrapper around the C function jsc_value_to_string_as_bytes.
func (recv *Value) ToStringAsBytes() *glib.Bytes {
	retC := C.jsc_value_to_string_as_bytes((*C.JSCValue)(recv.native))
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VirtualMachine is a wrapper around the C record JSCVirtualMachine.
type VirtualMachine struct {
	native *C.JSCVirtualMachine
	// parent : record
	// Private : priv
}

func VirtualMachineNewFromC(u unsafe.Pointer) *VirtualMachine {
	c := (*C.JSCVirtualMachine)(u)
	if c == nil {
		return nil
	}

	g := &VirtualMachine{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *VirtualMachine) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *VirtualMachine) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this VirtualMachine with another VirtualMachine, and returns true if they represent the same GObject.
func (recv *VirtualMachine) Equals(other *VirtualMachine) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *VirtualMachine) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to VirtualMachine.
// Exercise care, as this is a potentially dangerous function if the Object is not a VirtualMachine.
func CastToVirtualMachine(object *gobject.Object) *VirtualMachine {
	return VirtualMachineNewFromC(object.ToC())
}

// VirtualMachineNew is a wrapper around the C function jsc_virtual_machine_new.
func VirtualMachineNew() *VirtualMachine {
	retC := C.jsc_virtual_machine_new()
	retGo := VirtualMachineNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// WeakValue is a wrapper around the C record JSCWeakValue.
type WeakValue struct {
	native *C.JSCWeakValue
	// parent : record
	// Private : priv
}

func WeakValueNewFromC(u unsafe.Pointer) *WeakValue {
	c := (*C.JSCWeakValue)(u)
	if c == nil {
		return nil
	}

	g := &WeakValue{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *WeakValue) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *WeakValue) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WeakValue with another WeakValue, and returns true if they represent the same GObject.
func (recv *WeakValue) Equals(other *WeakValue) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *WeakValue) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to WeakValue.
// Exercise care, as this is a potentially dangerous function if the Object is not a WeakValue.
func CastToWeakValue(object *gobject.Object) *WeakValue {
	return WeakValueNewFromC(object.ToC())
}

type signalWeakValueClearedDetail struct {
	callback  WeakValueSignalClearedCallback
	handlerID C.gulong
}

var signalWeakValueClearedId int
var signalWeakValueClearedMap = make(map[int]signalWeakValueClearedDetail)
var signalWeakValueClearedLock sync.RWMutex

// WeakValueSignalClearedCallback is a callback function for a 'cleared' signal emitted from a WeakValue.
type WeakValueSignalClearedCallback func(targetObject *WeakValue)

/*
ConnectCleared connects the callback to the 'cleared' signal for the WeakValue.

The returned value represents the connection, and may be passed to DisconnectCleared to remove it.
*/
func (recv *WeakValue) ConnectCleared(callback WeakValueSignalClearedCallback) int {
	signalWeakValueClearedLock.Lock()
	defer signalWeakValueClearedLock.Unlock()

	signalWeakValueClearedId++
	instance := C.gpointer(recv.native)
	handlerID := C.WeakValue_signal_connect_cleared(instance, C.gpointer(uintptr(signalWeakValueClearedId)))

	detail := signalWeakValueClearedDetail{callback, handlerID}
	signalWeakValueClearedMap[signalWeakValueClearedId] = detail

	return signalWeakValueClearedId
}

/*
DisconnectCleared disconnects a callback from the 'cleared' signal for the WeakValue.

The connectionID should be a value returned from a call to ConnectCleared.
*/
func (recv *WeakValue) DisconnectCleared(connectionID int) {
	signalWeakValueClearedLock.Lock()
	defer signalWeakValueClearedLock.Unlock()

	detail, exists := signalWeakValueClearedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWeakValueClearedMap, connectionID)
}

//export weakvalue_clearedHandler
func weakvalue_clearedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalWeakValueClearedLock.RLock()
	defer signalWeakValueClearedLock.RUnlock()

	targetObject := WeakValueNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWeakValueClearedMap[index].callback
	callback(targetObject)
}

// WeakValueNew is a wrapper around the C function jsc_weak_value_new.
func WeakValueNew(value *Value) *WeakValue {
	c_value := (*C.JSCValue)(C.NULL)
	if value != nil {
		c_value = (*C.JSCValue)(value.ToC())
	}

	retC := C.jsc_weak_value_new(c_value)
	retGo := WeakValueNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetValue is a wrapper around the C function jsc_weak_value_get_value.
func (recv *WeakValue) GetValue() *Value {
	retC := C.jsc_weak_value_get_value((*C.JSCWeakValue)(recv.native))
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

const MAJOR_VERSION int32 = C.JSC_MAJOR_VERSION
const MICRO_VERSION int32 = C.JSC_MICRO_VERSION
const MINOR_VERSION int32 = C.JSC_MINOR_VERSION
const OPTIONS_USE_DFG string = C.JSC_OPTIONS_USE_DFG
const OPTIONS_USE_FTL string = C.JSC_OPTIONS_USE_FTL
const OPTIONS_USE_JIT string = C.JSC_OPTIONS_USE_JIT
const OPTIONS_USE_LLINT string = C.JSC_OPTIONS_USE_LLINT

type CheckSyntaxMode C.JSCCheckSyntaxMode

const (
	CHECK_SYNTAX_MODE_SCRIPT CheckSyntaxMode = 0
	CHECK_SYNTAX_MODE_MODULE CheckSyntaxMode = 1
)

type CheckSyntaxResult C.JSCCheckSyntaxResult

const (
	CHECK_SYNTAX_RESULT_SUCCESS                    CheckSyntaxResult = 0
	CHECK_SYNTAX_RESULT_RECOVERABLE_ERROR          CheckSyntaxResult = 1
	CHECK_SYNTAX_RESULT_IRRECOVERABLE_ERROR        CheckSyntaxResult = 2
	CHECK_SYNTAX_RESULT_UNTERMINATED_LITERAL_ERROR CheckSyntaxResult = 3
	CHECK_SYNTAX_RESULT_OUT_OF_MEMORY_ERROR        CheckSyntaxResult = 4
	CHECK_SYNTAX_RESULT_STACK_OVERFLOW_ERROR       CheckSyntaxResult = 5
)

type OptionType C.JSCOptionType

const (
	OPTION_BOOLEAN      OptionType = 0
	OPTION_INT          OptionType = 1
	OPTION_UINT         OptionType = 2
	OPTION_SIZE         OptionType = 3
	OPTION_DOUBLE       OptionType = 4
	OPTION_STRING       OptionType = 5
	OPTION_RANGE_STRING OptionType = 6
)

// GetMajorVersion is a wrapper around the C function jsc_get_major_version.
func GetMajorVersion() uint32 {
	retC := C.jsc_get_major_version()
	retGo := (uint32)(retC)

	return retGo
}

// GetMicroVersion is a wrapper around the C function jsc_get_micro_version.
func GetMicroVersion() uint32 {
	retC := C.jsc_get_micro_version()
	retGo := (uint32)(retC)

	return retGo
}

// GetMinorVersion is a wrapper around the C function jsc_get_minor_version.
func GetMinorVersion() uint32 {
	retC := C.jsc_get_minor_version()
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : jsc_options_foreach : unsupported parameter function : no type generator for OptionsFunc (JSCOptionsFunc) for param function

// OptionsGetBoolean is a wrapper around the C function jsc_options_get_boolean.
func OptionsGetBoolean(option string) (bool, bool) {
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	var c_value C.gboolean

	retC := C.jsc_options_get_boolean(c_option, &c_value)
	retGo := retC == C.TRUE

	value := c_value == C.TRUE

	return retGo, value
}

// OptionsGetDouble is a wrapper around the C function jsc_options_get_double.
func OptionsGetDouble(option string) (bool, float64) {
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	var c_value C.gdouble

	retC := C.jsc_options_get_double(c_option, &c_value)
	retGo := retC == C.TRUE

	value := (float64)(c_value)

	return retGo, value
}

// OptionsGetInt is a wrapper around the C function jsc_options_get_int.
func OptionsGetInt(option string) (bool, int32) {
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	var c_value C.gint

	retC := C.jsc_options_get_int(c_option, &c_value)
	retGo := retC == C.TRUE

	value := (int32)(c_value)

	return retGo, value
}

// OptionsGetOptionGroup is a wrapper around the C function jsc_options_get_option_group.
func OptionsGetOptionGroup() *glib.OptionGroup {
	retC := C.jsc_options_get_option_group()
	retGo := glib.OptionGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// OptionsGetRangeString is a wrapper around the C function jsc_options_get_range_string.
func OptionsGetRangeString(option string) (bool, string) {
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	var c_value *C.char

	retC := C.jsc_options_get_range_string(c_option, &c_value)
	retGo := retC == C.TRUE

	value := C.GoString(c_value)
	defer C.free(unsafe.Pointer(c_value))

	return retGo, value
}

// OptionsGetSize is a wrapper around the C function jsc_options_get_size.
func OptionsGetSize(option string) (bool, uint64) {
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	var c_value C.gsize

	retC := C.jsc_options_get_size(c_option, &c_value)
	retGo := retC == C.TRUE

	value := (uint64)(c_value)

	return retGo, value
}

// OptionsGetString is a wrapper around the C function jsc_options_get_string.
func OptionsGetString(option string) (bool, string) {
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	var c_value *C.char

	retC := C.jsc_options_get_string(c_option, &c_value)
	retGo := retC == C.TRUE

	value := C.GoString(c_value)
	defer C.free(unsafe.Pointer(c_value))

	return retGo, value
}

// OptionsGetUint is a wrapper around the C function jsc_options_get_uint.
func OptionsGetUint(option string) (bool, uint32) {
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	var c_value C.guint

	retC := C.jsc_options_get_uint(c_option, &c_value)
	retGo := retC == C.TRUE

	value := (uint32)(c_value)

	return retGo, value
}

// OptionsSetBoolean is a wrapper around the C function jsc_options_set_boolean.
func OptionsSetBoolean(option string, value bool) bool {
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	c_value :=
		boolToGboolean(value)

	retC := C.jsc_options_set_boolean(c_option, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// OptionsSetDouble is a wrapper around the C function jsc_options_set_double.
func OptionsSetDouble(option string, value float64) bool {
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	c_value := (C.gdouble)(value)

	retC := C.jsc_options_set_double(c_option, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// OptionsSetInt is a wrapper around the C function jsc_options_set_int.
func OptionsSetInt(option string, value int32) bool {
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	c_value := (C.gint)(value)

	retC := C.jsc_options_set_int(c_option, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// OptionsSetRangeString is a wrapper around the C function jsc_options_set_range_string.
func OptionsSetRangeString(option string, value string) bool {
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	retC := C.jsc_options_set_range_string(c_option, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// OptionsSetSize is a wrapper around the C function jsc_options_set_size.
func OptionsSetSize(option string, value uint64) bool {
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	c_value := (C.gsize)(value)

	retC := C.jsc_options_set_size(c_option, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// OptionsSetString is a wrapper around the C function jsc_options_set_string.
func OptionsSetString(option string, value string) bool {
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	retC := C.jsc_options_set_string(c_option, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// OptionsSetUint is a wrapper around the C function jsc_options_set_uint.
func OptionsSetUint(option string, value uint32) bool {
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	c_value := (C.guint)(value)

	retC := C.jsc_options_set_uint(c_option, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// ClassClass is a wrapper around the C record JSCClassClass.
type ClassClass struct {
	native *C.JSCClassClass
}

func ClassClassNewFromC(u unsafe.Pointer) *ClassClass {
	c := (*C.JSCClassClass)(u)
	if c == nil {
		return nil
	}

	g := &ClassClass{native: c}

	return g
}

func (recv *ClassClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ClassClass with another ClassClass, and returns true if they represent the same GObject.
func (recv *ClassClass) Equals(other *ClassClass) bool {
	return other.ToC() == recv.ToC()
}

// ClassVTable is a wrapper around the C record JSCClassVTable.
type ClassVTable struct {
	native *C.JSCClassVTable
	// get_property : no type generator for ClassGetPropertyFunction, JSCClassGetPropertyFunction
	// set_property : no type generator for ClassSetPropertyFunction, JSCClassSetPropertyFunction
	// has_property : no type generator for ClassHasPropertyFunction, JSCClassHasPropertyFunction
	// delete_property : no type generator for ClassDeletePropertyFunction, JSCClassDeletePropertyFunction
	// enumerate_properties : no type generator for ClassEnumeratePropertiesFunction, JSCClassEnumeratePropertiesFunction
	// no type for _jsc_reserved0
	// no type for _jsc_reserved1
	// no type for _jsc_reserved2
	// no type for _jsc_reserved3
}

func ClassVTableNewFromC(u unsafe.Pointer) *ClassVTable {
	c := (*C.JSCClassVTable)(u)
	if c == nil {
		return nil
	}

	g := &ClassVTable{native: c}

	return g
}

func (recv *ClassVTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ClassVTable with another ClassVTable, and returns true if they represent the same GObject.
func (recv *ClassVTable) Equals(other *ClassVTable) bool {
	return other.ToC() == recv.ToC()
}

// ContextClass is a wrapper around the C record JSCContextClass.
type ContextClass struct {
	native *C.JSCContextClass
	// parent_class : record
	// no type for _jsc_reserved0
	// no type for _jsc_reserved1
	// no type for _jsc_reserved2
	// no type for _jsc_reserved3
}

func ContextClassNewFromC(u unsafe.Pointer) *ContextClass {
	c := (*C.JSCContextClass)(u)
	if c == nil {
		return nil
	}

	g := &ContextClass{native: c}

	return g
}

func (recv *ContextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ContextClass with another ContextClass, and returns true if they represent the same GObject.
func (recv *ContextClass) Equals(other *ContextClass) bool {
	return other.ToC() == recv.ToC()
}

// ContextPrivate is a wrapper around the C record JSCContextPrivate.
type ContextPrivate struct {
	native *C.JSCContextPrivate
}

func ContextPrivateNewFromC(u unsafe.Pointer) *ContextPrivate {
	c := (*C.JSCContextPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ContextPrivate{native: c}

	return g
}

func (recv *ContextPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ContextPrivate with another ContextPrivate, and returns true if they represent the same GObject.
func (recv *ContextPrivate) Equals(other *ContextPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ExceptionClass is a wrapper around the C record JSCExceptionClass.
type ExceptionClass struct {
	native *C.JSCExceptionClass
	// parent_class : record
	// no type for _jsc_reserved0
	// no type for _jsc_reserved1
	// no type for _jsc_reserved2
	// no type for _jsc_reserved3
}

func ExceptionClassNewFromC(u unsafe.Pointer) *ExceptionClass {
	c := (*C.JSCExceptionClass)(u)
	if c == nil {
		return nil
	}

	g := &ExceptionClass{native: c}

	return g
}

func (recv *ExceptionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ExceptionClass with another ExceptionClass, and returns true if they represent the same GObject.
func (recv *ExceptionClass) Equals(other *ExceptionClass) bool {
	return other.ToC() == recv.ToC()
}

// ExceptionPrivate is a wrapper around the C record JSCExceptionPrivate.
type ExceptionPrivate struct {
	native *C.JSCExceptionPrivate
}

func ExceptionPrivateNewFromC(u unsafe.Pointer) *ExceptionPrivate {
	c := (*C.JSCExceptionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ExceptionPrivate{native: c}

	return g
}

func (recv *ExceptionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ExceptionPrivate with another ExceptionPrivate, and returns true if they represent the same GObject.
func (recv *ExceptionPrivate) Equals(other *ExceptionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ValueClass is a wrapper around the C record JSCValueClass.
type ValueClass struct {
	native *C.JSCValueClass
	// parent_class : record
	// no type for _jsc_reserved0
	// no type for _jsc_reserved1
	// no type for _jsc_reserved2
	// no type for _jsc_reserved3
}

func ValueClassNewFromC(u unsafe.Pointer) *ValueClass {
	c := (*C.JSCValueClass)(u)
	if c == nil {
		return nil
	}

	g := &ValueClass{native: c}

	return g
}

func (recv *ValueClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ValueClass with another ValueClass, and returns true if they represent the same GObject.
func (recv *ValueClass) Equals(other *ValueClass) bool {
	return other.ToC() == recv.ToC()
}

// ValuePrivate is a wrapper around the C record JSCValuePrivate.
type ValuePrivate struct {
	native *C.JSCValuePrivate
}

func ValuePrivateNewFromC(u unsafe.Pointer) *ValuePrivate {
	c := (*C.JSCValuePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ValuePrivate{native: c}

	return g
}

func (recv *ValuePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ValuePrivate with another ValuePrivate, and returns true if they represent the same GObject.
func (recv *ValuePrivate) Equals(other *ValuePrivate) bool {
	return other.ToC() == recv.ToC()
}

// VirtualMachineClass is a wrapper around the C record JSCVirtualMachineClass.
type VirtualMachineClass struct {
	native *C.JSCVirtualMachineClass
	// parent_class : record
	// no type for _jsc_reserved0
	// no type for _jsc_reserved1
	// no type for _jsc_reserved2
	// no type for _jsc_reserved3
}

func VirtualMachineClassNewFromC(u unsafe.Pointer) *VirtualMachineClass {
	c := (*C.JSCVirtualMachineClass)(u)
	if c == nil {
		return nil
	}

	g := &VirtualMachineClass{native: c}

	return g
}

func (recv *VirtualMachineClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this VirtualMachineClass with another VirtualMachineClass, and returns true if they represent the same GObject.
func (recv *VirtualMachineClass) Equals(other *VirtualMachineClass) bool {
	return other.ToC() == recv.ToC()
}

// VirtualMachinePrivate is a wrapper around the C record JSCVirtualMachinePrivate.
type VirtualMachinePrivate struct {
	native *C.JSCVirtualMachinePrivate
}

func VirtualMachinePrivateNewFromC(u unsafe.Pointer) *VirtualMachinePrivate {
	c := (*C.JSCVirtualMachinePrivate)(u)
	if c == nil {
		return nil
	}

	g := &VirtualMachinePrivate{native: c}

	return g
}

func (recv *VirtualMachinePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this VirtualMachinePrivate with another VirtualMachinePrivate, and returns true if they represent the same GObject.
func (recv *VirtualMachinePrivate) Equals(other *VirtualMachinePrivate) bool {
	return other.ToC() == recv.ToC()
}

// WeakValueClass is a wrapper around the C record JSCWeakValueClass.
type WeakValueClass struct {
	native *C.JSCWeakValueClass
	// parent_class : record
	// no type for _jsc_reserved0
	// no type for _jsc_reserved1
	// no type for _jsc_reserved2
	// no type for _jsc_reserved3
}

func WeakValueClassNewFromC(u unsafe.Pointer) *WeakValueClass {
	c := (*C.JSCWeakValueClass)(u)
	if c == nil {
		return nil
	}

	g := &WeakValueClass{native: c}

	return g
}

func (recv *WeakValueClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WeakValueClass with another WeakValueClass, and returns true if they represent the same GObject.
func (recv *WeakValueClass) Equals(other *WeakValueClass) bool {
	return other.ToC() == recv.ToC()
}

// WeakValuePrivate is a wrapper around the C record JSCWeakValuePrivate.
type WeakValuePrivate struct {
	native *C.JSCWeakValuePrivate
}

func WeakValuePrivateNewFromC(u unsafe.Pointer) *WeakValuePrivate {
	c := (*C.JSCWeakValuePrivate)(u)
	if c == nil {
		return nil
	}

	g := &WeakValuePrivate{native: c}

	return g
}

func (recv *WeakValuePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WeakValuePrivate with another WeakValuePrivate, and returns true if they represent the same GObject.
func (recv *WeakValuePrivate) Equals(other *WeakValuePrivate) bool {
	return other.ToC() == recv.ToC()
}
