// Code generated - DO NOT EDIT.

package javascriptcore

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"runtime"
	"sync"
)

var classStruct *gi.Struct
var classStruct_Once sync.Once

func classStruct_Set() error {
	var err error
	classStruct_Once.Do(func() {
		classStruct, err = gi.StructNew("JavaScriptCore", "Class")
	})
	return err
}

type Class struct {
	native uintptr
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
		err = classStruct_Set()
		if err != nil {
			return
		}
		classGetNameFunction, err = classStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type jsc_class_get_name.
func (recv *Class) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := classGetNameFunction_Set()
	if err == nil {
		ret = classGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'jsc_class_get_parent' : return type 'Class' not supported

// ClassStruct creates an uninitialised Class.
func ClassStruct() *Class {
	err := classStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Class{native: classStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeClass)
	return structGo
}
func finalizeClass(obj *Class) {
	classStruct.Free(obj.native)
}

var contextStruct *gi.Struct
var contextStruct_Once sync.Once

func contextStruct_Set() error {
	var err error
	contextStruct_Once.Do(func() {
		contextStruct, err = gi.StructNew("JavaScriptCore", "Context")
	})
	return err
}

type Context struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'jsc_context_new' : return type 'Context' not supported

// UNSUPPORTED : C value 'jsc_context_new_with_virtual_machine' : parameter 'vm' of type 'VirtualMachine' not supported

// UNSUPPORTED : C value 'jsc_context_check_syntax' : parameter 'mode' of type 'CheckSyntaxMode' not supported

var contextClearExceptionFunction *gi.Function
var contextClearExceptionFunction_Once sync.Once

func contextClearExceptionFunction_Set() error {
	var err error
	contextClearExceptionFunction_Once.Do(func() {
		err = contextStruct_Set()
		if err != nil {
			return
		}
		contextClearExceptionFunction, err = contextStruct.InvokerNew("clear_exception")
	})
	return err
}

// ClearException is a representation of the C type jsc_context_clear_exception.
func (recv *Context) ClearException() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := contextClearExceptionFunction_Set()
	if err == nil {
		contextClearExceptionFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'jsc_context_evaluate' : return type 'Value' not supported

// UNSUPPORTED : C value 'jsc_context_evaluate_in_object' : parameter 'object_instance' of type 'gpointer' not supported

// UNSUPPORTED : C value 'jsc_context_evaluate_with_source_uri' : return type 'Value' not supported

// UNSUPPORTED : C value 'jsc_context_get_exception' : return type 'Exception' not supported

// UNSUPPORTED : C value 'jsc_context_get_global_object' : return type 'Value' not supported

// UNSUPPORTED : C value 'jsc_context_get_value' : return type 'Value' not supported

// UNSUPPORTED : C value 'jsc_context_get_virtual_machine' : return type 'VirtualMachine' not supported

var contextPopExceptionHandlerFunction *gi.Function
var contextPopExceptionHandlerFunction_Once sync.Once

func contextPopExceptionHandlerFunction_Set() error {
	var err error
	contextPopExceptionHandlerFunction_Once.Do(func() {
		err = contextStruct_Set()
		if err != nil {
			return
		}
		contextPopExceptionHandlerFunction, err = contextStruct.InvokerNew("pop_exception_handler")
	})
	return err
}

// PopExceptionHandler is a representation of the C type jsc_context_pop_exception_handler.
func (recv *Context) PopExceptionHandler() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := contextPopExceptionHandlerFunction_Set()
	if err == nil {
		contextPopExceptionHandlerFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'jsc_context_push_exception_handler' : parameter 'handler' of type 'ExceptionHandler' not supported

// UNSUPPORTED : C value 'jsc_context_register_class' : parameter 'parent_class' of type 'Class' not supported

// UNSUPPORTED : C value 'jsc_context_set_value' : parameter 'value' of type 'Value' not supported

var contextThrowFunction *gi.Function
var contextThrowFunction_Once sync.Once

func contextThrowFunction_Set() error {
	var err error
	contextThrowFunction_Once.Do(func() {
		err = contextStruct_Set()
		if err != nil {
			return
		}
		contextThrowFunction, err = contextStruct.InvokerNew("throw")
	})
	return err
}

// Throw is a representation of the C type jsc_context_throw.
func (recv *Context) Throw(errorMessage string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(errorMessage)

	err := contextThrowFunction_Set()
	if err == nil {
		contextThrowFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'jsc_context_throw_exception' : parameter 'exception' of type 'Exception' not supported

// UNSUPPORTED : C value 'jsc_context_throw_printf' : parameter '...' of type 'nil' not supported

var contextThrowWithNameFunction *gi.Function
var contextThrowWithNameFunction_Once sync.Once

func contextThrowWithNameFunction_Set() error {
	var err error
	contextThrowWithNameFunction_Once.Do(func() {
		err = contextStruct_Set()
		if err != nil {
			return
		}
		contextThrowWithNameFunction, err = contextStruct.InvokerNew("throw_with_name")
	})
	return err
}

// ThrowWithName is a representation of the C type jsc_context_throw_with_name.
func (recv *Context) ThrowWithName(errorName string, errorMessage string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(errorName)
	inArgs[2].SetString(errorMessage)

	err := contextThrowWithNameFunction_Set()
	if err == nil {
		contextThrowWithNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'jsc_context_throw_with_name_printf' : parameter '...' of type 'nil' not supported

var exceptionStruct *gi.Struct
var exceptionStruct_Once sync.Once

func exceptionStruct_Set() error {
	var err error
	exceptionStruct_Once.Do(func() {
		exceptionStruct, err = gi.StructNew("JavaScriptCore", "Exception")
	})
	return err
}

type Exception struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'jsc_exception_new' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'jsc_exception_new_printf' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'jsc_exception_new_vprintf' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'jsc_exception_new_with_name' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'jsc_exception_new_with_name_printf' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'jsc_exception_new_with_name_vprintf' : parameter 'context' of type 'Context' not supported

var exceptionGetBacktraceStringFunction *gi.Function
var exceptionGetBacktraceStringFunction_Once sync.Once

func exceptionGetBacktraceStringFunction_Set() error {
	var err error
	exceptionGetBacktraceStringFunction_Once.Do(func() {
		err = exceptionStruct_Set()
		if err != nil {
			return
		}
		exceptionGetBacktraceStringFunction, err = exceptionStruct.InvokerNew("get_backtrace_string")
	})
	return err
}

// GetBacktraceString is a representation of the C type jsc_exception_get_backtrace_string.
func (recv *Exception) GetBacktraceString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = exceptionStruct_Set()
		if err != nil {
			return
		}
		exceptionGetColumnNumberFunction, err = exceptionStruct.InvokerNew("get_column_number")
	})
	return err
}

// GetColumnNumber is a representation of the C type jsc_exception_get_column_number.
func (recv *Exception) GetColumnNumber() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = exceptionStruct_Set()
		if err != nil {
			return
		}
		exceptionGetLineNumberFunction, err = exceptionStruct.InvokerNew("get_line_number")
	})
	return err
}

// GetLineNumber is a representation of the C type jsc_exception_get_line_number.
func (recv *Exception) GetLineNumber() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = exceptionStruct_Set()
		if err != nil {
			return
		}
		exceptionGetMessageFunction, err = exceptionStruct.InvokerNew("get_message")
	})
	return err
}

// GetMessage is a representation of the C type jsc_exception_get_message.
func (recv *Exception) GetMessage() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = exceptionStruct_Set()
		if err != nil {
			return
		}
		exceptionGetNameFunction, err = exceptionStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type jsc_exception_get_name.
func (recv *Exception) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = exceptionStruct_Set()
		if err != nil {
			return
		}
		exceptionGetSourceUriFunction, err = exceptionStruct.InvokerNew("get_source_uri")
	})
	return err
}

// GetSourceUri is a representation of the C type jsc_exception_get_source_uri.
func (recv *Exception) GetSourceUri() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = exceptionStruct_Set()
		if err != nil {
			return
		}
		exceptionReportFunction, err = exceptionStruct.InvokerNew("report")
	})
	return err
}

// Report is a representation of the C type jsc_exception_report.
func (recv *Exception) Report() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = exceptionStruct_Set()
		if err != nil {
			return
		}
		exceptionToStringFunction, err = exceptionStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type jsc_exception_to_string.
func (recv *Exception) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := exceptionToStringFunction_Set()
	if err == nil {
		ret = exceptionToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var valueStruct *gi.Struct
var valueStruct_Once sync.Once

func valueStruct_Set() error {
	var err error
	valueStruct_Once.Do(func() {
		valueStruct, err = gi.StructNew("JavaScriptCore", "Value")
	})
	return err
}

type Value struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'jsc_value_new_array' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'jsc_value_new_array_from_garray' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'jsc_value_new_array_from_strv' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'jsc_value_new_boolean' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'jsc_value_new_function' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'jsc_value_new_function_variadic' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'jsc_value_new_functionv' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'jsc_value_new_null' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'jsc_value_new_number' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'jsc_value_new_object' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'jsc_value_new_string' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'jsc_value_new_string_from_bytes' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'jsc_value_new_undefined' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'jsc_value_constructor_call' : parameter 'first_parameter_type' of type 'GType' not supported

// UNSUPPORTED : C value 'jsc_value_constructor_callv' : parameter 'parameters' of type 'nil' not supported

// UNSUPPORTED : C value 'jsc_value_function_call' : parameter 'first_parameter_type' of type 'GType' not supported

// UNSUPPORTED : C value 'jsc_value_function_callv' : parameter 'parameters' of type 'nil' not supported

// UNSUPPORTED : C value 'jsc_value_get_context' : return type 'Context' not supported

var valueIsArrayFunction *gi.Function
var valueIsArrayFunction_Once sync.Once

func valueIsArrayFunction_Set() error {
	var err error
	valueIsArrayFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueIsArrayFunction, err = valueStruct.InvokerNew("is_array")
	})
	return err
}

// IsArray is a representation of the C type jsc_value_is_array.
func (recv *Value) IsArray() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueIsBooleanFunction, err = valueStruct.InvokerNew("is_boolean")
	})
	return err
}

// IsBoolean is a representation of the C type jsc_value_is_boolean.
func (recv *Value) IsBoolean() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueIsConstructorFunction, err = valueStruct.InvokerNew("is_constructor")
	})
	return err
}

// IsConstructor is a representation of the C type jsc_value_is_constructor.
func (recv *Value) IsConstructor() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueIsFunctionFunction, err = valueStruct.InvokerNew("is_function")
	})
	return err
}

// IsFunction is a representation of the C type jsc_value_is_function.
func (recv *Value) IsFunction() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueIsNullFunction, err = valueStruct.InvokerNew("is_null")
	})
	return err
}

// IsNull is a representation of the C type jsc_value_is_null.
func (recv *Value) IsNull() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueIsNumberFunction, err = valueStruct.InvokerNew("is_number")
	})
	return err
}

// IsNumber is a representation of the C type jsc_value_is_number.
func (recv *Value) IsNumber() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueIsObjectFunction, err = valueStruct.InvokerNew("is_object")
	})
	return err
}

// IsObject is a representation of the C type jsc_value_is_object.
func (recv *Value) IsObject() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueIsStringFunction, err = valueStruct.InvokerNew("is_string")
	})
	return err
}

// IsString is a representation of the C type jsc_value_is_string.
func (recv *Value) IsString() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueIsUndefinedFunction, err = valueStruct.InvokerNew("is_undefined")
	})
	return err
}

// IsUndefined is a representation of the C type jsc_value_is_undefined.
func (recv *Value) IsUndefined() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueObjectDeletePropertyFunction, err = valueStruct.InvokerNew("object_delete_property")
	})
	return err
}

// ObjectDeleteProperty is a representation of the C type jsc_value_object_delete_property.
func (recv *Value) ObjectDeleteProperty(name string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueObjectEnumeratePropertiesFunction, err = valueStruct.InvokerNew("object_enumerate_properties")
	})
	return err
}

// ObjectEnumerateProperties is a representation of the C type jsc_value_object_enumerate_properties.
func (recv *Value) ObjectEnumerateProperties() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := valueObjectEnumeratePropertiesFunction_Set()
	if err == nil {
		valueObjectEnumeratePropertiesFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'jsc_value_object_get_property' : return type 'Value' not supported

// UNSUPPORTED : C value 'jsc_value_object_get_property_at_index' : return type 'Value' not supported

var valueObjectHasPropertyFunction *gi.Function
var valueObjectHasPropertyFunction_Once sync.Once

func valueObjectHasPropertyFunction_Set() error {
	var err error
	valueObjectHasPropertyFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueObjectHasPropertyFunction, err = valueStruct.InvokerNew("object_has_property")
	})
	return err
}

// ObjectHasProperty is a representation of the C type jsc_value_object_has_property.
func (recv *Value) ObjectHasProperty(name string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueObjectIsInstanceOfFunction, err = valueStruct.InvokerNew("object_is_instance_of")
	})
	return err
}

// ObjectIsInstanceOf is a representation of the C type jsc_value_object_is_instance_of.
func (recv *Value) ObjectIsInstanceOf(name string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := valueObjectIsInstanceOfFunction_Set()
	if err == nil {
		ret = valueObjectIsInstanceOfFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'jsc_value_object_set_property' : parameter 'property' of type 'Value' not supported

// UNSUPPORTED : C value 'jsc_value_object_set_property_at_index' : parameter 'property' of type 'Value' not supported

var valueToBooleanFunction *gi.Function
var valueToBooleanFunction_Once sync.Once

func valueToBooleanFunction_Set() error {
	var err error
	valueToBooleanFunction_Once.Do(func() {
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueToBooleanFunction, err = valueStruct.InvokerNew("to_boolean")
	})
	return err
}

// ToBoolean is a representation of the C type jsc_value_to_boolean.
func (recv *Value) ToBoolean() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueToDoubleFunction, err = valueStruct.InvokerNew("to_double")
	})
	return err
}

// ToDouble is a representation of the C type jsc_value_to_double.
func (recv *Value) ToDouble() float64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueToInt32Function, err = valueStruct.InvokerNew("to_int32")
	})
	return err
}

// ToInt32 is a representation of the C type jsc_value_to_int32.
func (recv *Value) ToInt32() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = valueStruct_Set()
		if err != nil {
			return
		}
		valueToStringFunction, err = valueStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type jsc_value_to_string.
func (recv *Value) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := valueToStringFunction_Set()
	if err == nil {
		ret = valueToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'jsc_value_to_string_as_bytes' : return type 'GLib.Bytes' not supported

var virtualMachineStruct *gi.Struct
var virtualMachineStruct_Once sync.Once

func virtualMachineStruct_Set() error {
	var err error
	virtualMachineStruct_Once.Do(func() {
		virtualMachineStruct, err = gi.StructNew("JavaScriptCore", "VirtualMachine")
	})
	return err
}

type VirtualMachine struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'jsc_virtual_machine_new' : return type 'VirtualMachine' not supported

var weakValueStruct *gi.Struct
var weakValueStruct_Once sync.Once

func weakValueStruct_Set() error {
	var err error
	weakValueStruct_Once.Do(func() {
		weakValueStruct, err = gi.StructNew("JavaScriptCore", "WeakValue")
	})
	return err
}

type WeakValue struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'jsc_weak_value_new' : parameter 'value' of type 'Value' not supported

// UNSUPPORTED : C value 'jsc_weak_value_get_value' : return type 'Value' not supported
