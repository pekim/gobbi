// This is a generated file - DO NOT EDIT

package gobject

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Provide a copy of a boxed structure @src_boxed which is of type @boxed_type.
/*

C function

g_boxed_copy
*/
func BoxedCopy(boxedType Type, srcBoxed uintptr) uintptr {
	c_boxed_type := (C.GType)(boxedType)

	c_src_boxed := (C.gconstpointer)(srcBoxed)

	retC := C.g_boxed_copy(c_boxed_type, c_src_boxed)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Free the boxed structure @boxed which is of type @boxed_type.
/*

C function

g_boxed_free
*/
func BoxedFree(boxedType Type, boxed uintptr) {
	c_boxed_type := (C.GType)(boxedType)

	c_boxed := (C.gpointer)(boxed)

	C.g_boxed_free(c_boxed_type, c_boxed)

	return
}

// Unsupported : g_boxed_type_register_static : unsupported parameter boxed_copy : no type generator for BoxedCopyFunc (GBoxedCopyFunc) for param boxed_copy

// A #GClosureMarshal function for use with signals with handlers that
// take two boxed pointers as arguments and return a boolean.  If you
// have such a signal, you will probably also need to use an
// accumulator, such as g_signal_accumulator_true_handled().
/*

C function

g_cclosure_marshal_BOOLEAN__BOXED_BOXED
*/
func CclosureMarshalBooleanBoxedBoxed(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_BOOLEAN__BOXED_BOXED(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with handlers that
// take a flags type as an argument and return a boolean.  If you have
// such a signal, you will probably also need to use an accumulator,
// such as g_signal_accumulator_true_handled().
/*

C function

g_cclosure_marshal_BOOLEAN__FLAGS
*/
func CclosureMarshalBooleanFlags(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_BOOLEAN__FLAGS(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with handlers that
// take a #GObject and a pointer and produce a string.  It is highly
// unlikely that your signal handler fits this description.
/*

C function

g_cclosure_marshal_STRING__OBJECT_POINTER
*/
func CclosureMarshalStringObjectPointer(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_STRING__OBJECT_POINTER(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with a single
// boolean argument.
/*

C function

g_cclosure_marshal_VOID__BOOLEAN
*/
func CclosureMarshalVoidBoolean(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__BOOLEAN(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with a single
// argument which is any boxed pointer type.
/*

C function

g_cclosure_marshal_VOID__BOXED
*/
func CclosureMarshalVoidBoxed(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__BOXED(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with a single
// character argument.
/*

C function

g_cclosure_marshal_VOID__CHAR
*/
func CclosureMarshalVoidChar(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__CHAR(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with one
// double-precision floating point argument.
/*

C function

g_cclosure_marshal_VOID__DOUBLE
*/
func CclosureMarshalVoidDouble(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__DOUBLE(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with a single
// argument with an enumerated type.
/*

C function

g_cclosure_marshal_VOID__ENUM
*/
func CclosureMarshalVoidEnum(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__ENUM(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with a single
// argument with a flags types.
/*

C function

g_cclosure_marshal_VOID__FLAGS
*/
func CclosureMarshalVoidFlags(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__FLAGS(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with one
// single-precision floating point argument.
/*

C function

g_cclosure_marshal_VOID__FLOAT
*/
func CclosureMarshalVoidFloat(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__FLOAT(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with a single
// integer argument.
/*

C function

g_cclosure_marshal_VOID__INT
*/
func CclosureMarshalVoidInt(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__INT(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with with a single
// long integer argument.
/*

C function

g_cclosure_marshal_VOID__LONG
*/
func CclosureMarshalVoidLong(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__LONG(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with a single
// #GObject argument.
/*

C function

g_cclosure_marshal_VOID__OBJECT
*/
func CclosureMarshalVoidObject(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__OBJECT(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with a single
// argument of type #GParamSpec.
/*

C function

g_cclosure_marshal_VOID__PARAM
*/
func CclosureMarshalVoidParam(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__PARAM(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with a single raw
// pointer argument type.
//
// If it is possible, it is better to use one of the more specific
// functions such as g_cclosure_marshal_VOID__OBJECT() or
// g_cclosure_marshal_VOID__OBJECT().
/*

C function

g_cclosure_marshal_VOID__POINTER
*/
func CclosureMarshalVoidPointer(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__POINTER(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with a single string
// argument.
/*

C function

g_cclosure_marshal_VOID__STRING
*/
func CclosureMarshalVoidString(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__STRING(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with a single
// unsigned character argument.
/*

C function

g_cclosure_marshal_VOID__UCHAR
*/
func CclosureMarshalVoidUchar(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__UCHAR(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with with a single
// unsigned integer argument.
/*

C function

g_cclosure_marshal_VOID__UINT
*/
func CclosureMarshalVoidUint(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__UINT(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with a unsigned int
// and a pointer as arguments.
/*

C function

g_cclosure_marshal_VOID__UINT_POINTER
*/
func CclosureMarshalVoidUintPointer(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__UINT_POINTER(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with a single
// unsigned long integer argument.
/*

C function

g_cclosure_marshal_VOID__ULONG
*/
func CclosureMarshalVoidUlong(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__ULONG(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with a single
// #GVariant argument.
/*

C function

g_cclosure_marshal_VOID__VARIANT
*/
func CclosureMarshalVoidVariant(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__VARIANT(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// A #GClosureMarshal function for use with signals with no arguments.
/*

C function

g_cclosure_marshal_VOID__VOID
*/
func CclosureMarshalVoidVoid(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__VOID(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// Unsupported : g_cclosure_new : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_object : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_object_swap : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_swap : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// This function is meant to be called from the `complete_type_info`
// function of a #GTypePlugin implementation, as in the following
// example:
//
// |[<!-- language="C" -->
// static void
// my_enum_complete_type_info (GTypePlugin     *plugin,
// GType            g_type,
// GTypeInfo       *info,
// GTypeValueTable *value_table)
// {
// static const GEnumValue values[] = {
// { MY_ENUM_FOO, "MY_ENUM_FOO", "foo" },
// { MY_ENUM_BAR, "MY_ENUM_BAR", "bar" },
// { 0, NULL, NULL }
// };
//
// g_enum_complete_type_info (type, info, values);
// }
// ]|
/*

C function

g_enum_complete_type_info
*/
func EnumCompleteTypeInfo(gEnumType Type, constValues *EnumValue) *TypeInfo {
	c_g_enum_type := (C.GType)(gEnumType)

	var c_info C.GTypeInfo

	c_const_values := (*C.GEnumValue)(C.NULL)
	if constValues != nil {
		c_const_values = (*C.GEnumValue)(constValues.ToC())
	}

	C.g_enum_complete_type_info(c_g_enum_type, &c_info, c_const_values)

	info := TypeInfoNewFromC(unsafe.Pointer(&c_info))

	return info
}

// Returns the #GEnumValue for a value.
/*

C function

g_enum_get_value
*/
func EnumGetValue(enumClass *EnumClass, value int32) *EnumValue {
	c_enum_class := (*C.GEnumClass)(C.NULL)
	if enumClass != nil {
		c_enum_class = (*C.GEnumClass)(enumClass.ToC())
	}

	c_value := (C.gint)(value)

	retC := C.g_enum_get_value(c_enum_class, c_value)
	retGo := EnumValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Looks up a #GEnumValue by name.
/*

C function

g_enum_get_value_by_name
*/
func EnumGetValueByName(enumClass *EnumClass, name string) *EnumValue {
	c_enum_class := (*C.GEnumClass)(C.NULL)
	if enumClass != nil {
		c_enum_class = (*C.GEnumClass)(enumClass.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_enum_get_value_by_name(c_enum_class, c_name)
	retGo := EnumValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Looks up a #GEnumValue by nickname.
/*

C function

g_enum_get_value_by_nick
*/
func EnumGetValueByNick(enumClass *EnumClass, nick string) *EnumValue {
	c_enum_class := (*C.GEnumClass)(C.NULL)
	if enumClass != nil {
		c_enum_class = (*C.GEnumClass)(enumClass.ToC())
	}

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	retC := C.g_enum_get_value_by_nick(c_enum_class, c_nick)
	retGo := EnumValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Registers a new static enumeration type with the name @name.
//
// It is normally more convenient to let [glib-mkenums][glib-mkenums],
// generate a my_enum_get_type() function from a usual C enumeration
// definition  than to write one yourself using g_enum_register_static().
/*

C function

g_enum_register_static
*/
func EnumRegisterStatic(name string, constStaticValues *EnumValue) Type {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_const_static_values := (*C.GEnumValue)(C.NULL)
	if constStaticValues != nil {
		c_const_static_values = (*C.GEnumValue)(constStaticValues.ToC())
	}

	retC := C.g_enum_register_static(c_name, c_const_static_values)
	retGo := (Type)(retC)

	return retGo
}

// This function is meant to be called from the complete_type_info()
// function of a #GTypePlugin implementation, see the example for
// g_enum_complete_type_info() above.
/*

C function

g_flags_complete_type_info
*/
func FlagsCompleteTypeInfo(gFlagsType Type, constValues *FlagsValue) *TypeInfo {
	c_g_flags_type := (C.GType)(gFlagsType)

	var c_info C.GTypeInfo

	c_const_values := (*C.GFlagsValue)(C.NULL)
	if constValues != nil {
		c_const_values = (*C.GFlagsValue)(constValues.ToC())
	}

	C.g_flags_complete_type_info(c_g_flags_type, &c_info, c_const_values)

	info := TypeInfoNewFromC(unsafe.Pointer(&c_info))

	return info
}

// Returns the first #GFlagsValue which is set in @value.
/*

C function

g_flags_get_first_value
*/
func FlagsGetFirstValue(flagsClass *FlagsClass, value uint32) *FlagsValue {
	c_flags_class := (*C.GFlagsClass)(C.NULL)
	if flagsClass != nil {
		c_flags_class = (*C.GFlagsClass)(flagsClass.ToC())
	}

	c_value := (C.guint)(value)

	retC := C.g_flags_get_first_value(c_flags_class, c_value)
	retGo := FlagsValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Looks up a #GFlagsValue by name.
/*

C function

g_flags_get_value_by_name
*/
func FlagsGetValueByName(flagsClass *FlagsClass, name string) *FlagsValue {
	c_flags_class := (*C.GFlagsClass)(C.NULL)
	if flagsClass != nil {
		c_flags_class = (*C.GFlagsClass)(flagsClass.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_flags_get_value_by_name(c_flags_class, c_name)
	retGo := FlagsValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Looks up a #GFlagsValue by nickname.
/*

C function

g_flags_get_value_by_nick
*/
func FlagsGetValueByNick(flagsClass *FlagsClass, nick string) *FlagsValue {
	c_flags_class := (*C.GFlagsClass)(C.NULL)
	if flagsClass != nil {
		c_flags_class = (*C.GFlagsClass)(flagsClass.ToC())
	}

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	retC := C.g_flags_get_value_by_nick(c_flags_class, c_nick)
	retGo := FlagsValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Registers a new static flags type with the name @name.
//
// It is normally more convenient to let [glib-mkenums][glib-mkenums]
// generate a my_flags_get_type() function from a usual C enumeration
// definition than to write one yourself using g_flags_register_static().
/*

C function

g_flags_register_static
*/
func FlagsRegisterStatic(name string, constStaticValues *FlagsValue) Type {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_const_static_values := (*C.GFlagsValue)(C.NULL)
	if constStaticValues != nil {
		c_const_static_values = (*C.GFlagsValue)(constStaticValues.ToC())
	}

	retC := C.g_flags_register_static(c_name, c_const_static_values)
	retGo := (Type)(retC)

	return retGo
}

/*

C function

g_gtype_get_type
*/
func GtypeGetType() Type {
	retC := C.g_gtype_get_type()
	retGo := (Type)(retC)

	return retGo
}

// Unsupported : g_param_spec_boolean : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_boxed : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_char : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_double : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_enum : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_flags : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_float : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_int : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_int64 : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_long : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_object : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_param : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_pointer : return type : Blacklisted record : GParamSpec

// Creates a new #GParamSpecPool.
//
// If @type_prefixing is %TRUE, lookups in the newly created pool will
// allow to specify the owner as a colon-separated prefix of the
// property name, like "GtkContainer:border-width". This feature is
// deprecated, so you should always set @type_prefixing to %FALSE.
/*

C function

g_param_spec_pool_new
*/
func ParamSpecPoolNew(typePrefixing bool) *ParamSpecPool {
	c_type_prefixing :=
		boolToGboolean(typePrefixing)

	retC := C.g_param_spec_pool_new(c_type_prefixing)
	retGo := ParamSpecPoolNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_param_spec_string : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_uchar : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_uint : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_uint64 : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_ulong : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_unichar : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_value_array : unsupported parameter element_spec : Blacklisted record : GParamSpec

// Registers @name as the name of a new static type derived from
// #G_TYPE_PARAM. The type system uses the information contained in
// the #GParamSpecTypeInfo structure pointed to by @info to manage the
// #GParamSpec type and its instances.
/*

C function

g_param_type_register_static
*/
func ParamTypeRegisterStatic(name string, pspecInfo *ParamSpecTypeInfo) Type {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_pspec_info := (*C.GParamSpecTypeInfo)(C.NULL)
	if pspecInfo != nil {
		c_pspec_info = (*C.GParamSpecTypeInfo)(pspecInfo.ToC())
	}

	retC := C.g_param_type_register_static(c_name, c_pspec_info)
	retGo := (Type)(retC)

	return retGo
}

// Unsupported : g_param_value_convert : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : g_param_value_defaults : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : g_param_value_set_default : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : g_param_value_validate : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : g_param_values_cmp : unsupported parameter pspec : Blacklisted record : GParamSpec

// Creates a new %G_TYPE_POINTER derived type id for a new
// pointer type with name @name.
/*

C function

g_pointer_type_register_static
*/
func PointerTypeRegisterStatic(name string) Type {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_pointer_type_register_static(c_name)
	retGo := (Type)(retC)

	return retGo
}

// Unsupported : g_signal_add_emission_hook : unsupported parameter hook_func : no type generator for SignalEmissionHook (GSignalEmissionHook) for param hook_func

// Unsupported : g_signal_chain_from_overridden : unsupported parameter instance_and_params :

// Connects a closure to a signal for a particular object.
/*

C function

g_signal_connect_closure
*/
func SignalConnectClosure(instance uintptr, detailedSignal string, closure *Closure, after bool) uint64 {
	c_instance := (C.gpointer)(instance)

	c_detailed_signal := C.CString(detailedSignal)
	defer C.free(unsafe.Pointer(c_detailed_signal))

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_after :=
		boolToGboolean(after)

	retC := C.g_signal_connect_closure(c_instance, c_detailed_signal, c_closure, c_after)
	retGo := (uint64)(retC)

	return retGo
}

// Connects a closure to a signal for a particular object.
/*

C function

g_signal_connect_closure_by_id
*/
func SignalConnectClosureById(instance uintptr, signalId uint32, detail glib.Quark, closure *Closure, after bool) uint64 {
	c_instance := (C.gpointer)(instance)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_after :=
		boolToGboolean(after)

	retC := C.g_signal_connect_closure_by_id(c_instance, c_signal_id, c_detail, c_closure, c_after)
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported : g_signal_connect_data : unsupported parameter c_handler : no type generator for Callback (GCallback) for param c_handler

// Unsupported : g_signal_connect_object : unsupported parameter c_handler : no type generator for Callback (GCallback) for param c_handler

// Unsupported : g_signal_emit : unsupported parameter ... : varargs

// Unsupported : g_signal_emit_by_name : unsupported parameter ... : varargs

// Unsupported : g_signal_emit_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Unsupported : g_signal_emitv : unsupported parameter instance_and_params :

// Returns the invocation hint of the innermost signal emission of instance.
/*

C function

g_signal_get_invocation_hint
*/
func SignalGetInvocationHint(instance uintptr) *SignalInvocationHint {
	c_instance := (C.gpointer)(instance)

	retC := C.g_signal_get_invocation_hint(c_instance)
	retGo := SignalInvocationHintNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blocks a handler of an instance so it will not be called during any
// signal emissions unless it is unblocked again. Thus "blocking" a
// signal handler means to temporarily deactive it, a signal handler
// has to be unblocked exactly the same amount of times it has been
// blocked before to become active again.
//
// The @handler_id has to be a valid signal handler id, connected to a
// signal of @instance.
/*

C function

g_signal_handler_block
*/
func SignalHandlerBlock(instance uintptr, handlerId uint64) {
	c_instance := (C.gpointer)(instance)

	c_handler_id := (C.gulong)(handlerId)

	C.g_signal_handler_block(c_instance, c_handler_id)

	return
}

// Disconnects a handler from an instance so it will not be called during
// any future or currently ongoing emissions of the signal it has been
// connected to. The @handler_id becomes invalid and may be reused.
//
// The @handler_id has to be a valid signal handler id, connected to a
// signal of @instance.
/*

C function

g_signal_handler_disconnect
*/
func SignalHandlerDisconnect(instance uintptr, handlerId uint64) {
	c_instance := (C.gpointer)(instance)

	c_handler_id := (C.gulong)(handlerId)

	C.g_signal_handler_disconnect(c_instance, c_handler_id)

	return
}

// Finds the first signal handler that matches certain selection criteria.
// The criteria mask is passed as an OR-ed combination of #GSignalMatchType
// flags, and the criteria values are passed as arguments.
// The match @mask has to be non-0 for successful matches.
// If no handler was found, 0 is returned.
/*

C function

g_signal_handler_find
*/
func SignalHandlerFind(instance uintptr, mask SignalMatchType, signalId uint32, detail glib.Quark, closure *Closure, func_ uintptr, data uintptr) uint64 {
	c_instance := (C.gpointer)(instance)

	c_mask := (C.GSignalMatchType)(mask)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_func := (C.gpointer)(func_)

	c_data := (C.gpointer)(data)

	retC := C.g_signal_handler_find(c_instance, c_mask, c_signal_id, c_detail, c_closure, c_func, c_data)
	retGo := (uint64)(retC)

	return retGo
}

// Returns whether @handler_id is the ID of a handler connected to @instance.
/*

C function

g_signal_handler_is_connected
*/
func SignalHandlerIsConnected(instance uintptr, handlerId uint64) bool {
	c_instance := (C.gpointer)(instance)

	c_handler_id := (C.gulong)(handlerId)

	retC := C.g_signal_handler_is_connected(c_instance, c_handler_id)
	retGo := retC == C.TRUE

	return retGo
}

// Undoes the effect of a previous g_signal_handler_block() call.  A
// blocked handler is skipped during signal emissions and will not be
// invoked, unblocking it (for exactly the amount of times it has been
// blocked before) reverts its "blocked" state, so the handler will be
// recognized by the signal system and is called upon future or
// currently ongoing signal emissions (since the order in which
// handlers are called during signal emissions is deterministic,
// whether the unblocked handler in question is called as part of a
// currently ongoing emission depends on how far that emission has
// proceeded yet).
//
// The @handler_id has to be a valid id of a signal handler that is
// connected to a signal of @instance and is currently blocked.
/*

C function

g_signal_handler_unblock
*/
func SignalHandlerUnblock(instance uintptr, handlerId uint64) {
	c_instance := (C.gpointer)(instance)

	c_handler_id := (C.gulong)(handlerId)

	C.g_signal_handler_unblock(c_instance, c_handler_id)

	return
}

// Blocks all handlers on an instance that match a certain selection criteria.
// The criteria mask is passed as an OR-ed combination of #GSignalMatchType
// flags, and the criteria values are passed as arguments.
// Passing at least one of the %G_SIGNAL_MATCH_CLOSURE, %G_SIGNAL_MATCH_FUNC
// or %G_SIGNAL_MATCH_DATA match flags is required for successful matches.
// If no handlers were found, 0 is returned, the number of blocked handlers
// otherwise.
/*

C function

g_signal_handlers_block_matched
*/
func SignalHandlersBlockMatched(instance uintptr, mask SignalMatchType, signalId uint32, detail glib.Quark, closure *Closure, func_ uintptr, data uintptr) uint32 {
	c_instance := (C.gpointer)(instance)

	c_mask := (C.GSignalMatchType)(mask)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_func := (C.gpointer)(func_)

	c_data := (C.gpointer)(data)

	retC := C.g_signal_handlers_block_matched(c_instance, c_mask, c_signal_id, c_detail, c_closure, c_func, c_data)
	retGo := (uint32)(retC)

	return retGo
}

// Destroy all signal handlers of a type instance. This function is
// an implementation detail of the #GObject dispose implementation,
// and should not be used outside of the type system.
/*

C function

g_signal_handlers_destroy
*/
func SignalHandlersDestroy(instance uintptr) {
	c_instance := (C.gpointer)(instance)

	C.g_signal_handlers_destroy(c_instance)

	return
}

// Disconnects all handlers on an instance that match a certain
// selection criteria. The criteria mask is passed as an OR-ed
// combination of #GSignalMatchType flags, and the criteria values are
// passed as arguments.  Passing at least one of the
// %G_SIGNAL_MATCH_CLOSURE, %G_SIGNAL_MATCH_FUNC or
// %G_SIGNAL_MATCH_DATA match flags is required for successful
// matches.  If no handlers were found, 0 is returned, the number of
// disconnected handlers otherwise.
/*

C function

g_signal_handlers_disconnect_matched
*/
func SignalHandlersDisconnectMatched(instance uintptr, mask SignalMatchType, signalId uint32, detail glib.Quark, closure *Closure, func_ uintptr, data uintptr) uint32 {
	c_instance := (C.gpointer)(instance)

	c_mask := (C.GSignalMatchType)(mask)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_func := (C.gpointer)(func_)

	c_data := (C.gpointer)(data)

	retC := C.g_signal_handlers_disconnect_matched(c_instance, c_mask, c_signal_id, c_detail, c_closure, c_func, c_data)
	retGo := (uint32)(retC)

	return retGo
}

// Unblocks all handlers on an instance that match a certain selection
// criteria. The criteria mask is passed as an OR-ed combination of
// #GSignalMatchType flags, and the criteria values are passed as arguments.
// Passing at least one of the %G_SIGNAL_MATCH_CLOSURE, %G_SIGNAL_MATCH_FUNC
// or %G_SIGNAL_MATCH_DATA match flags is required for successful matches.
// If no handlers were found, 0 is returned, the number of unblocked handlers
// otherwise. The match criteria should not apply to any handlers that are
// not currently blocked.
/*

C function

g_signal_handlers_unblock_matched
*/
func SignalHandlersUnblockMatched(instance uintptr, mask SignalMatchType, signalId uint32, detail glib.Quark, closure *Closure, func_ uintptr, data uintptr) uint32 {
	c_instance := (C.gpointer)(instance)

	c_mask := (C.GSignalMatchType)(mask)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_func := (C.gpointer)(func_)

	c_data := (C.gpointer)(data)

	retC := C.g_signal_handlers_unblock_matched(c_instance, c_mask, c_signal_id, c_detail, c_closure, c_func, c_data)
	retGo := (uint32)(retC)

	return retGo
}

// Returns whether there are any handlers connected to @instance for the
// given signal id and detail.
//
// If @detail is 0 then it will only match handlers that were connected
// without detail.  If @detail is non-zero then it will match handlers
// connected both without detail and with the given detail.  This is
// consistent with how a signal emitted with @detail would be delivered
// to those handlers.
//
// Since 2.46 this also checks for a non-default class closure being
// installed, as this is basically always what you want.
//
// One example of when you might use this is when the arguments to the
// signal are difficult to compute. A class implementor may opt to not
// emit the signal if no one is attached anyway, thus saving the cost
// of building the arguments.
/*

C function

g_signal_has_handler_pending
*/
func SignalHasHandlerPending(instance uintptr, signalId uint32, detail glib.Quark, mayBeBlocked bool) bool {
	c_instance := (C.gpointer)(instance)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_may_be_blocked :=
		boolToGboolean(mayBeBlocked)

	retC := C.g_signal_has_handler_pending(c_instance, c_signal_id, c_detail, c_may_be_blocked)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_signal_list_ids : no return type

// Given the name of the signal and the type of object it connects to, gets
// the signal's identifying integer. Emitting the signal by number is
// somewhat faster than using the name each time.
//
// Also tries the ancestors of the given type.
//
// See g_signal_new() for details on allowed signal names.
/*

C function

g_signal_lookup
*/
func SignalLookup(name string, itype Type) uint32 {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_itype := (C.GType)(itype)

	retC := C.g_signal_lookup(c_name, c_itype)
	retGo := (uint32)(retC)

	return retGo
}

// Given the signal's identifier, finds its name.
//
// Two different signals may have the same name, if they have differing types.
/*

C function

g_signal_name
*/
func SignalName(signalId uint32) string {
	c_signal_id := (C.guint)(signalId)

	retC := C.g_signal_name(c_signal_id)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_signal_new : unsupported parameter accumulator : no type generator for SignalAccumulator (GSignalAccumulator) for param accumulator

// Unsupported : g_signal_new_valist : unsupported parameter accumulator : no type generator for SignalAccumulator (GSignalAccumulator) for param accumulator

// Unsupported : g_signal_newv : unsupported parameter accumulator : no type generator for SignalAccumulator (GSignalAccumulator) for param accumulator

// Overrides the class closure (i.e. the default handler) for the given signal
// for emissions on instances of @instance_type. @instance_type must be derived
// from the type to which the signal belongs.
//
// See g_signal_chain_from_overridden() and
// g_signal_chain_from_overridden_handler() for how to chain up to the
// parent class closure from inside the overridden one.
/*

C function

g_signal_override_class_closure
*/
func SignalOverrideClassClosure(signalId uint32, instanceType Type, classClosure *Closure) {
	c_signal_id := (C.guint)(signalId)

	c_instance_type := (C.GType)(instanceType)

	c_class_closure := (*C.GClosure)(C.NULL)
	if classClosure != nil {
		c_class_closure = (*C.GClosure)(classClosure.ToC())
	}

	C.g_signal_override_class_closure(c_signal_id, c_instance_type, c_class_closure)

	return
}

// Internal function to parse a signal name into its @signal_id
// and @detail quark.
/*

C function

g_signal_parse_name
*/
func SignalParseName(detailedSignal string, itype Type, forceDetailQuark bool) (bool, uint32, glib.Quark) {
	c_detailed_signal := C.CString(detailedSignal)
	defer C.free(unsafe.Pointer(c_detailed_signal))

	c_itype := (C.GType)(itype)

	var c_signal_id_p C.guint

	var c_detail_p C.GQuark

	c_force_detail_quark :=
		boolToGboolean(forceDetailQuark)

	retC := C.g_signal_parse_name(c_detailed_signal, c_itype, &c_signal_id_p, &c_detail_p, c_force_detail_quark)
	retGo := retC == C.TRUE

	signalIdP := (uint32)(c_signal_id_p)

	detailP := (glib.Quark)(c_detail_p)

	return retGo, signalIdP, detailP
}

// Queries the signal system for in-depth information about a
// specific signal. This function will fill in a user-provided
// structure to hold signal-specific information. If an invalid
// signal id is passed in, the @signal_id member of the #GSignalQuery
// is 0. All members filled into the #GSignalQuery structure should
// be considered constant and have to be left untouched.
/*

C function

g_signal_query
*/
func SignalQuery_(signalId uint32) *SignalQuery {
	c_signal_id := (C.guint)(signalId)

	var c_query C.GSignalQuery

	C.g_signal_query(c_signal_id, &c_query)

	query := SignalQueryNewFromC(unsafe.Pointer(&c_query))

	return query
}

// Deletes an emission hook.
/*

C function

g_signal_remove_emission_hook
*/
func SignalRemoveEmissionHook(signalId uint32, hookId uint64) {
	c_signal_id := (C.guint)(signalId)

	c_hook_id := (C.gulong)(hookId)

	C.g_signal_remove_emission_hook(c_signal_id, c_hook_id)

	return
}

// Stops a signal's current emission.
//
// This will prevent the default method from running, if the signal was
// %G_SIGNAL_RUN_LAST and you connected normally (i.e. without the "after"
// flag).
//
// Prints a warning if used on a signal which isn't being emitted.
/*

C function

g_signal_stop_emission
*/
func SignalStopEmission(instance uintptr, signalId uint32, detail glib.Quark) {
	c_instance := (C.gpointer)(instance)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	C.g_signal_stop_emission(c_instance, c_signal_id, c_detail)

	return
}

// Stops a signal's current emission.
//
// This is just like g_signal_stop_emission() except it will look up the
// signal id for you.
/*

C function

g_signal_stop_emission_by_name
*/
func SignalStopEmissionByName(instance uintptr, detailedSignal string) {
	c_instance := (C.gpointer)(instance)

	c_detailed_signal := C.CString(detailedSignal)
	defer C.free(unsafe.Pointer(c_detailed_signal))

	C.g_signal_stop_emission_by_name(c_instance, c_detailed_signal)

	return
}

// Creates a new closure which invokes the function found at the offset
// @struct_offset in the class structure of the interface or classed type
// identified by @itype.
/*

C function

g_signal_type_cclosure_new
*/
func SignalTypeCclosureNew(itype Type, structOffset uint32) *Closure {
	c_itype := (C.GType)(itype)

	c_struct_offset := (C.guint)(structOffset)

	retC := C.g_signal_type_cclosure_new(c_itype, c_struct_offset)
	retGo := ClosureNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Set the callback for a source as a #GClosure.
//
// If the source is not one of the standard GLib types, the @closure_callback
// and @closure_marshal fields of the #GSourceFuncs structure must have been
// filled in with pointers to appropriate functions.
/*

C function

g_source_set_closure
*/
func SourceSetClosure(source *glib.Source, closure *Closure) {
	c_source := (*C.GSource)(C.NULL)
	if source != nil {
		c_source = (*C.GSource)(source.ToC())
	}

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	C.g_source_set_closure(c_source, c_closure)

	return
}

// Sets a dummy callback for @source. The callback will do nothing, and
// if the source expects a #gboolean return value, it will return %TRUE.
// (If the source expects any other type of return value, it will return
// a 0/%NULL value; whatever g_value_init() initializes a #GValue to for
// that type.)
//
// If the source is not one of the standard GLib types, the
// @closure_callback and @closure_marshal fields of the #GSourceFuncs
// structure must have been filled in with pointers to appropriate
// functions.
/*

C function

g_source_set_dummy_callback
*/
func SourceSetDummyCallback(source *glib.Source) {
	c_source := (*C.GSource)(C.NULL)
	if source != nil {
		c_source = (*C.GSource)(source.ToC())
	}

	C.g_source_set_dummy_callback(c_source)

	return
}

// Return a newly allocated string, which describes the contents of a
// #GValue.  The main purpose of this function is to describe #GValue
// contents for debugging output, the way in which the contents are
// described may change between different GLib versions.
/*

C function

g_strdup_value_contents
*/
func StrdupValueContents(value *Value) string {
	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.g_strdup_value_contents(c_value)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_type_add_class_cache_func : unsupported parameter cache_func : no type generator for TypeClassCacheFunc (GTypeClassCacheFunc) for param cache_func

/*

C function

g_type_add_instance_private
*/
func TypeAddInstancePrivate(classType Type, privateSize uint64) int32 {
	c_class_type := (C.GType)(classType)

	c_private_size := (C.gsize)(privateSize)

	retC := C.g_type_add_instance_private(c_class_type, c_private_size)
	retGo := (int32)(retC)

	return retGo
}

// Adds the dynamic @interface_type to @instantiable_type. The information
// contained in the #GTypePlugin structure pointed to by @plugin
// is used to manage the relationship.
/*

C function

g_type_add_interface_dynamic
*/
func TypeAddInterfaceDynamic(instanceType Type, interfaceType Type, plugin *TypePlugin) {
	c_instance_type := (C.GType)(instanceType)

	c_interface_type := (C.GType)(interfaceType)

	c_plugin := (*C.GTypePlugin)(plugin.ToC())

	C.g_type_add_interface_dynamic(c_instance_type, c_interface_type, c_plugin)

	return
}

// Adds the static @interface_type to @instantiable_type.
// The information contained in the #GInterfaceInfo structure
// pointed to by @info is used to manage the relationship.
/*

C function

g_type_add_interface_static
*/
func TypeAddInterfaceStatic(instanceType Type, interfaceType Type, info *InterfaceInfo) {
	c_instance_type := (C.GType)(instanceType)

	c_interface_type := (C.GType)(interfaceType)

	c_info := (*C.GInterfaceInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GInterfaceInfo)(info.ToC())
	}

	C.g_type_add_interface_static(c_instance_type, c_interface_type, c_info)

	return
}

/*

C function

g_type_check_class_cast
*/
func TypeCheckClassCast(gClass *TypeClass, isAType Type) *TypeClass {
	c_g_class := (*C.GTypeClass)(C.NULL)
	if gClass != nil {
		c_g_class = (*C.GTypeClass)(gClass.ToC())
	}

	c_is_a_type := (C.GType)(isAType)

	retC := C.g_type_check_class_cast(c_g_class, c_is_a_type)
	retGo := TypeClassNewFromC(unsafe.Pointer(retC))

	return retGo
}

/*

C function

g_type_check_class_is_a
*/
func TypeCheckClassIsA(gClass *TypeClass, isAType Type) bool {
	c_g_class := (*C.GTypeClass)(C.NULL)
	if gClass != nil {
		c_g_class = (*C.GTypeClass)(gClass.ToC())
	}

	c_is_a_type := (C.GType)(isAType)

	retC := C.g_type_check_class_is_a(c_g_class, c_is_a_type)
	retGo := retC == C.TRUE

	return retGo
}

// Private helper function to aid implementation of the
// G_TYPE_CHECK_INSTANCE() macro.
/*

C function

g_type_check_instance
*/
func TypeCheckInstance(instance *TypeInstance) bool {
	c_instance := (*C.GTypeInstance)(C.NULL)
	if instance != nil {
		c_instance = (*C.GTypeInstance)(instance.ToC())
	}

	retC := C.g_type_check_instance(c_instance)
	retGo := retC == C.TRUE

	return retGo
}

/*

C function

g_type_check_instance_cast
*/
func TypeCheckInstanceCast(instance *TypeInstance, ifaceType Type) *TypeInstance {
	c_instance := (*C.GTypeInstance)(C.NULL)
	if instance != nil {
		c_instance = (*C.GTypeInstance)(instance.ToC())
	}

	c_iface_type := (C.GType)(ifaceType)

	retC := C.g_type_check_instance_cast(c_instance, c_iface_type)
	retGo := TypeInstanceNewFromC(unsafe.Pointer(retC))

	return retGo
}

/*

C function

g_type_check_instance_is_a
*/
func TypeCheckInstanceIsA(instance *TypeInstance, ifaceType Type) bool {
	c_instance := (*C.GTypeInstance)(C.NULL)
	if instance != nil {
		c_instance = (*C.GTypeInstance)(instance.ToC())
	}

	c_iface_type := (C.GType)(ifaceType)

	retC := C.g_type_check_instance_is_a(c_instance, c_iface_type)
	retGo := retC == C.TRUE

	return retGo
}

/*

C function

g_type_check_instance_is_fundamentally_a
*/
func TypeCheckInstanceIsFundamentallyA(instance *TypeInstance, fundamentalType Type) bool {
	c_instance := (*C.GTypeInstance)(C.NULL)
	if instance != nil {
		c_instance = (*C.GTypeInstance)(instance.ToC())
	}

	c_fundamental_type := (C.GType)(fundamentalType)

	retC := C.g_type_check_instance_is_fundamentally_a(c_instance, c_fundamental_type)
	retGo := retC == C.TRUE

	return retGo
}

/*

C function

g_type_check_is_value_type
*/
func TypeCheckIsValueType(type_ Type) bool {
	c_type := (C.GType)(type_)

	retC := C.g_type_check_is_value_type(c_type)
	retGo := retC == C.TRUE

	return retGo
}

/*

C function

g_type_check_value
*/
func TypeCheckValue(value *Value) bool {
	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.g_type_check_value(c_value)
	retGo := retC == C.TRUE

	return retGo
}

/*

C function

g_type_check_value_holds
*/
func TypeCheckValueHolds(value *Value, type_ Type) bool {
	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	c_type := (C.GType)(type_)

	retC := C.g_type_check_value_holds(c_value, c_type)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_type_children : no return type

/*

C function

g_type_class_adjust_private_offset
*/
func TypeClassAdjustPrivateOffset(gClass uintptr, privateSizeOrOffset int32) {
	c_g_class := (C.gpointer)(gClass)

	c_private_size_or_offset := (C.gint)(privateSizeOrOffset)

	C.g_type_class_adjust_private_offset(c_g_class, &c_private_size_or_offset)

	return
}

// This function is essentially the same as g_type_class_ref(),
// except that the classes reference count isn't incremented.
// As a consequence, this function may return %NULL if the class
// of the type passed in does not currently exist (hasn't been
// referenced before).
/*

C function

g_type_class_peek
*/
func TypeClassPeek(type_ Type) uintptr {
	c_type := (C.GType)(type_)

	retC := C.g_type_class_peek(c_type)
	retGo := (uintptr)(retC)

	return retGo
}

// Increments the reference count of the class structure belonging to
// @type. This function will demand-create the class if it doesn't
// exist already.
/*

C function

g_type_class_ref
*/
func TypeClassRef(type_ Type) uintptr {
	c_type := (C.GType)(type_)

	retC := C.g_type_class_ref(c_type)
	retGo := (uintptr)(retC)

	return retGo
}

// Creates and initializes an instance of @type if @type is valid and
// can be instantiated. The type system only performs basic allocation
// and structure setups for instances: actual instance creation should
// happen through functions supplied by the type's fundamental type
// implementation.  So use of g_type_create_instance() is reserved for
// implementators of fundamental types only. E.g. instances of the
// #GObject hierarchy should be created via g_object_new() and never
// directly through g_type_create_instance() which doesn't handle things
// like singleton objects or object construction.
//
// The extended members of the returned instance are guaranteed to be filled
// with zeros.
//
// Note: Do not use this function, unless you're implementing a
// fundamental type. Also language bindings should not use this
// function, but g_object_new() instead.
/*

C function

g_type_create_instance
*/
func TypeCreateInstance(type_ Type) *TypeInstance {
	c_type := (C.GType)(type_)

	retC := C.g_type_create_instance(c_type)
	retGo := TypeInstanceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the length of the ancestry of the passed in type. This
// includes the type itself, so that e.g. a fundamental type has depth 1.
/*

C function

g_type_depth
*/
func TypeDepth(type_ Type) uint32 {
	c_type := (C.GType)(type_)

	retC := C.g_type_depth(c_type)
	retGo := (uint32)(retC)

	return retGo
}

// Frees an instance of a type, returning it to the instance pool for
// the type, if there is one.
//
// Like g_type_create_instance(), this function is reserved for
// implementors of fundamental types.
/*

C function

g_type_free_instance
*/
func TypeFreeInstance(instance *TypeInstance) {
	c_instance := (*C.GTypeInstance)(C.NULL)
	if instance != nil {
		c_instance = (*C.GTypeInstance)(instance.ToC())
	}

	C.g_type_free_instance(c_instance)

	return
}

// Lookup the type ID from a given type name, returning 0 if no type
// has been registered under this name (this is the preferred method
// to find out by name whether a specific type has been registered
// yet).
/*

C function

g_type_from_name
*/
func TypeFromName(name string) Type {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_type_from_name(c_name)
	retGo := (Type)(retC)

	return retGo
}

// Internal function, used to extract the fundamental type ID portion.
// Use G_TYPE_FUNDAMENTAL() instead.
/*

C function

g_type_fundamental
*/
func TypeFundamental(typeId Type) Type {
	c_type_id := (C.GType)(typeId)

	retC := C.g_type_fundamental(c_type_id)
	retGo := (Type)(retC)

	return retGo
}

// Returns the next free fundamental type id which can be used to
// register a new fundamental type with g_type_register_fundamental().
// The returned type ID represents the highest currently registered
// fundamental type identifier.
/*

C function

g_type_fundamental_next
*/
func TypeFundamentalNext() Type {
	retC := C.g_type_fundamental_next()
	retGo := (Type)(retC)

	return retGo
}

// Returns the #GTypePlugin structure for @type.
/*

C function

g_type_get_plugin
*/
func TypeGetPlugin(type_ Type) *TypePlugin {
	c_type := (C.GType)(type_)

	retC := C.g_type_get_plugin(c_type)
	retGo := TypePluginNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Obtains data which has previously been attached to @type
// with g_type_set_qdata().
//
// Note that this does not take subtyping into account; data
// attached to one type with g_type_set_qdata() cannot
// be retrieved from a subtype using g_type_get_qdata().
/*

C function

g_type_get_qdata
*/
func TypeGetQdata(type_ Type, quark glib.Quark) uintptr {
	c_type := (C.GType)(type_)

	c_quark := (C.GQuark)(quark)

	retC := C.g_type_get_qdata(c_type, c_quark)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// This function used to initialise the type system.  Since GLib 2.36,
// the type system is initialised automatically and this function does
// nothing.
/*

C function

g_type_init
*/
func TypeInit() {
	C.g_type_init()

	return
}

// This function used to initialise the type system with debugging
// flags.  Since GLib 2.36, the type system is initialised automatically
// and this function does nothing.
//
// If you need to enable debugging features, use the GOBJECT_DEBUG
// environment variable.
/*

C function

g_type_init_with_debug_flags
*/
func TypeInitWithDebugFlags(debugFlags TypeDebugFlags) {
	c_debug_flags := (C.GTypeDebugFlags)(debugFlags)

	C.g_type_init_with_debug_flags(c_debug_flags)

	return
}

// Adds @prerequisite_type to the list of prerequisites of @interface_type.
// This means that any type implementing @interface_type must also implement
// @prerequisite_type. Prerequisites can be thought of as an alternative to
// interface derivation (which GType doesn't support). An interface can have
// at most one instantiatable prerequisite type.
/*

C function

g_type_interface_add_prerequisite
*/
func TypeInterfaceAddPrerequisite(interfaceType Type, prerequisiteType Type) {
	c_interface_type := (C.GType)(interfaceType)

	c_prerequisite_type := (C.GType)(prerequisiteType)

	C.g_type_interface_add_prerequisite(c_interface_type, c_prerequisite_type)

	return
}

// Returns the #GTypePlugin structure for the dynamic interface
// @interface_type which has been added to @instance_type, or %NULL
// if @interface_type has not been added to @instance_type or does
// not have a #GTypePlugin structure. See g_type_add_interface_dynamic().
/*

C function

g_type_interface_get_plugin
*/
func TypeInterfaceGetPlugin(instanceType Type, interfaceType Type) *TypePlugin {
	c_instance_type := (C.GType)(instanceType)

	c_interface_type := (C.GType)(interfaceType)

	retC := C.g_type_interface_get_plugin(c_instance_type, c_interface_type)
	retGo := TypePluginNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the #GTypeInterface structure of an interface to which the
// passed in class conforms.
/*

C function

g_type_interface_peek
*/
func TypeInterfacePeek(instanceClass uintptr, ifaceType Type) uintptr {
	c_instance_class := (C.gpointer)(instanceClass)

	c_iface_type := (C.GType)(ifaceType)

	retC := C.g_type_interface_peek(c_instance_class, c_iface_type)
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_type_interfaces : no return type

// If @is_a_type is a derivable type, check whether @type is a
// descendant of @is_a_type. If @is_a_type is an interface, check
// whether @type conforms to it.
/*

C function

g_type_is_a
*/
func TypeIsA(type_ Type, isAType Type) bool {
	c_type := (C.GType)(type_)

	c_is_a_type := (C.GType)(isAType)

	retC := C.g_type_is_a(c_type, c_is_a_type)
	retGo := retC == C.TRUE

	return retGo
}

// Get the unique name that is assigned to a type ID.  Note that this
// function (like all other GType API) cannot cope with invalid type
// IDs. %G_TYPE_INVALID may be passed to this function, as may be any
// other validly registered type ID, but randomized type IDs should
// not be passed in and will most likely lead to a crash.
/*

C function

g_type_name
*/
func TypeName(type_ Type) string {
	c_type := (C.GType)(type_)

	retC := C.g_type_name(c_type)
	retGo := C.GoString(retC)

	return retGo
}

/*

C function

g_type_name_from_class
*/
func TypeNameFromClass(gClass *TypeClass) string {
	c_g_class := (*C.GTypeClass)(C.NULL)
	if gClass != nil {
		c_g_class = (*C.GTypeClass)(gClass.ToC())
	}

	retC := C.g_type_name_from_class(c_g_class)
	retGo := C.GoString(retC)

	return retGo
}

/*

C function

g_type_name_from_instance
*/
func TypeNameFromInstance(instance *TypeInstance) string {
	c_instance := (*C.GTypeInstance)(C.NULL)
	if instance != nil {
		c_instance = (*C.GTypeInstance)(instance.ToC())
	}

	retC := C.g_type_name_from_instance(c_instance)
	retGo := C.GoString(retC)

	return retGo
}

// Given a @leaf_type and a @root_type which is contained in its
// anchestry, return the type that @root_type is the immediate parent
// of. In other words, this function determines the type that is
// derived directly from @root_type which is also a base class of
// @leaf_type.  Given a root type and a leaf type, this function can
// be used to determine the types and order in which the leaf type is
// descended from the root type.
/*

C function

g_type_next_base
*/
func TypeNextBase(leafType Type, rootType Type) Type {
	c_leaf_type := (C.GType)(leafType)

	c_root_type := (C.GType)(rootType)

	retC := C.g_type_next_base(c_leaf_type, c_root_type)
	retGo := (Type)(retC)

	return retGo
}

// Return the direct parent type of the passed in type. If the passed
// in type has no parent, i.e. is a fundamental type, 0 is returned.
/*

C function

g_type_parent
*/
func TypeParent(type_ Type) Type {
	c_type := (C.GType)(type_)

	retC := C.g_type_parent(c_type)
	retGo := (Type)(retC)

	return retGo
}

// Get the corresponding quark of the type IDs name.
/*

C function

g_type_qname
*/
func TypeQname(type_ Type) glib.Quark {
	c_type := (C.GType)(type_)

	retC := C.g_type_qname(c_type)
	retGo := (glib.Quark)(retC)

	return retGo
}

// Queries the type system for information about a specific type.
// This function will fill in a user-provided structure to hold
// type-specific information. If an invalid #GType is passed in, the
// @type member of the #GTypeQuery is 0. All members filled into the
// #GTypeQuery structure should be considered constant and have to be
// left untouched.
/*

C function

g_type_query
*/
func TypeQuery_(type_ Type) *TypeQuery {
	c_type := (C.GType)(type_)

	var c_query C.GTypeQuery

	C.g_type_query(c_type, &c_query)

	query := TypeQueryNewFromC(unsafe.Pointer(&c_query))

	return query
}

// Registers @type_name as the name of a new dynamic type derived from
// @parent_type.  The type system uses the information contained in the
// #GTypePlugin structure pointed to by @plugin to manage the type and its
// instances (if not abstract).  The value of @flags determines the nature
// (e.g. abstract or not) of the type.
/*

C function

g_type_register_dynamic
*/
func TypeRegisterDynamic(parentType Type, typeName string, plugin *TypePlugin, flags TypeFlags) Type {
	c_parent_type := (C.GType)(parentType)

	c_type_name := C.CString(typeName)
	defer C.free(unsafe.Pointer(c_type_name))

	c_plugin := (*C.GTypePlugin)(plugin.ToC())

	c_flags := (C.GTypeFlags)(flags)

	retC := C.g_type_register_dynamic(c_parent_type, c_type_name, c_plugin, c_flags)
	retGo := (Type)(retC)

	return retGo
}

// Registers @type_id as the predefined identifier and @type_name as the
// name of a fundamental type. If @type_id is already registered, or a
// type named @type_name is already registered, the behaviour is undefined.
// The type system uses the information contained in the #GTypeInfo structure
// pointed to by @info and the #GTypeFundamentalInfo structure pointed to by
// @finfo to manage the type and its instances. The value of @flags determines
// additional characteristics of the fundamental type.
/*

C function

g_type_register_fundamental
*/
func TypeRegisterFundamental(typeId Type, typeName string, info *TypeInfo, finfo *TypeFundamentalInfo, flags TypeFlags) Type {
	c_type_id := (C.GType)(typeId)

	c_type_name := C.CString(typeName)
	defer C.free(unsafe.Pointer(c_type_name))

	c_info := (*C.GTypeInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GTypeInfo)(info.ToC())
	}

	c_finfo := (*C.GTypeFundamentalInfo)(C.NULL)
	if finfo != nil {
		c_finfo = (*C.GTypeFundamentalInfo)(finfo.ToC())
	}

	c_flags := (C.GTypeFlags)(flags)

	retC := C.g_type_register_fundamental(c_type_id, c_type_name, c_info, c_finfo, c_flags)
	retGo := (Type)(retC)

	return retGo
}

// Registers @type_name as the name of a new static type derived from
// @parent_type. The type system uses the information contained in the
// #GTypeInfo structure pointed to by @info to manage the type and its
// instances (if not abstract). The value of @flags determines the nature
// (e.g. abstract or not) of the type.
/*

C function

g_type_register_static
*/
func TypeRegisterStatic(parentType Type, typeName string, info *TypeInfo, flags TypeFlags) Type {
	c_parent_type := (C.GType)(parentType)

	c_type_name := C.CString(typeName)
	defer C.free(unsafe.Pointer(c_type_name))

	c_info := (*C.GTypeInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GTypeInfo)(info.ToC())
	}

	c_flags := (C.GTypeFlags)(flags)

	retC := C.g_type_register_static(c_parent_type, c_type_name, c_info, c_flags)
	retGo := (Type)(retC)

	return retGo
}

// Unsupported : g_type_remove_class_cache_func : unsupported parameter cache_func : no type generator for TypeClassCacheFunc (GTypeClassCacheFunc) for param cache_func

// Attaches arbitrary data to a type.
/*

C function

g_type_set_qdata
*/
func TypeSetQdata(type_ Type, quark glib.Quark, data uintptr) {
	c_type := (C.GType)(type_)

	c_quark := (C.GQuark)(quark)

	c_data := (C.gpointer)(data)

	C.g_type_set_qdata(c_type, c_quark, c_data)

	return
}

/*

C function

g_type_test_flags
*/
func TypeTestFlags(type_ Type, flags uint32) bool {
	c_type := (C.GType)(type_)

	c_flags := (C.guint)(flags)

	retC := C.g_type_test_flags(c_type, c_flags)
	retGo := retC == C.TRUE

	return retGo
}

// Returns the location of the #GTypeValueTable associated with @type.
//
// Note that this function should only be used from source code
// that implements or has internal knowledge of the implementation of
// @type.
/*

C function

g_type_value_table_peek
*/
func TypeValueTablePeek(type_ Type) *TypeValueTable {
	c_type := (C.GType)(type_)

	retC := C.g_type_value_table_peek(c_type)
	retGo := TypeValueTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_value_register_transform_func : unsupported parameter transform_func : no type generator for ValueTransform (GValueTransform) for param transform_func

// Returns whether a #GValue of type @src_type can be copied into
// a #GValue of type @dest_type.
/*

C function

g_value_type_compatible
*/
func ValueTypeCompatible(srcType Type, destType Type) bool {
	c_src_type := (C.GType)(srcType)

	c_dest_type := (C.GType)(destType)

	retC := C.g_value_type_compatible(c_src_type, c_dest_type)
	retGo := retC == C.TRUE

	return retGo
}

// Check whether g_value_transform() is able to transform values
// of type @src_type into values of type @dest_type. Note that for
// the types to be transformable, they must be compatible or a
// transformation function must be registered.
/*

C function

g_value_type_transformable
*/
func ValueTypeTransformable(srcType Type, destType Type) bool {
	c_src_type := (C.GType)(srcType)

	c_dest_type := (C.GType)(destType)

	retC := C.g_value_type_transformable(c_src_type, c_dest_type)
	retGo := retC == C.TRUE

	return retGo
}
