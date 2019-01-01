// This is a generated file - DO NOT EDIT

package gobject

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// CClosure is a wrapper around the C record GCClosure.
type CClosure struct {
	native *C.GCClosure
	// closure : record
	Callback uintptr
}

func CClosureNewFromC(u unsafe.Pointer) *CClosure {
	c := (*C.GCClosure)(u)
	if c == nil {
		return nil
	}

	g := &CClosure{
		Callback: (uintptr)(c.callback),
		native:   c,
	}

	return g
}

func (recv *CClosure) ToC() unsafe.Pointer {
	recv.native.callback =
		(C.gpointer)(recv.Callback)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CClosure with another CClosure, and returns true if they represent the same GObject.
func (recv *CClosure) Equals(other *CClosure) bool {
	return other.ToC() == recv.ToC()
}

// CClosureMarshalBooleanBoxedBoxed is a wrapper around the C function g_cclosure_marshal_BOOLEAN__BOXED_BOXED.
func CClosureMarshalBooleanBoxedBoxed(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_BOOLEAN__BOXED_BOXEDv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalBooleanFlags is a wrapper around the C function g_cclosure_marshal_BOOLEAN__FLAGS.
func CClosureMarshalBooleanFlags(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_BOOLEAN__FLAGSv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalStringObjectPointer is a wrapper around the C function g_cclosure_marshal_STRING__OBJECT_POINTER.
func CClosureMarshalStringObjectPointer(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_STRING__OBJECT_POINTERv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidBoolean is a wrapper around the C function g_cclosure_marshal_VOID__BOOLEAN.
func CClosureMarshalVoidBoolean(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_VOID__BOOLEANv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidBoxed is a wrapper around the C function g_cclosure_marshal_VOID__BOXED.
func CClosureMarshalVoidBoxed(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_VOID__BOXEDv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidChar is a wrapper around the C function g_cclosure_marshal_VOID__CHAR.
func CClosureMarshalVoidChar(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_VOID__CHARv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidDouble is a wrapper around the C function g_cclosure_marshal_VOID__DOUBLE.
func CClosureMarshalVoidDouble(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_VOID__DOUBLEv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidEnum is a wrapper around the C function g_cclosure_marshal_VOID__ENUM.
func CClosureMarshalVoidEnum(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_VOID__ENUMv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidFlags is a wrapper around the C function g_cclosure_marshal_VOID__FLAGS.
func CClosureMarshalVoidFlags(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_VOID__FLAGSv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidFloat is a wrapper around the C function g_cclosure_marshal_VOID__FLOAT.
func CClosureMarshalVoidFloat(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_VOID__FLOATv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidInt is a wrapper around the C function g_cclosure_marshal_VOID__INT.
func CClosureMarshalVoidInt(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_VOID__INTv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidLong is a wrapper around the C function g_cclosure_marshal_VOID__LONG.
func CClosureMarshalVoidLong(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_VOID__LONGv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidObject is a wrapper around the C function g_cclosure_marshal_VOID__OBJECT.
func CClosureMarshalVoidObject(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_VOID__OBJECTv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidParam is a wrapper around the C function g_cclosure_marshal_VOID__PARAM.
func CClosureMarshalVoidParam(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_VOID__PARAMv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidPointer is a wrapper around the C function g_cclosure_marshal_VOID__POINTER.
func CClosureMarshalVoidPointer(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_VOID__POINTERv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidString is a wrapper around the C function g_cclosure_marshal_VOID__STRING.
func CClosureMarshalVoidString(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_VOID__STRINGv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidUchar is a wrapper around the C function g_cclosure_marshal_VOID__UCHAR.
func CClosureMarshalVoidUchar(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_VOID__UCHARv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidUint is a wrapper around the C function g_cclosure_marshal_VOID__UINT.
func CClosureMarshalVoidUint(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// CClosureMarshalVoidUintPointer is a wrapper around the C function g_cclosure_marshal_VOID__UINT_POINTER.
func CClosureMarshalVoidUintPointer(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_VOID__UINT_POINTERv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__UINTv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidUlong is a wrapper around the C function g_cclosure_marshal_VOID__ULONG.
func CClosureMarshalVoidUlong(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_VOID__ULONGv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidVariant is a wrapper around the C function g_cclosure_marshal_VOID__VARIANT.
func CClosureMarshalVoidVariant(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_VOID__VARIANTv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidVoid is a wrapper around the C function g_cclosure_marshal_VOID__VOID.
func CClosureMarshalVoidVoid(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
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

// g_cclosure_marshal_VOID__VOIDv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_new : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func
// g_cclosure_new_object : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func
// g_cclosure_new_object_swap : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func
// g_cclosure_new_swap : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func
// Closure is a wrapper around the C record GClosure.
type Closure struct {
	native *C.GClosure
	// Private : ref_count
	// Private : meta_marshal_nouse
	// Private : n_guards
	// Private : n_fnotifiers
	// Private : n_inotifiers
	// Private : in_inotify
	// Private : floating
	// Private : derivative_flag
	// Bitfield not supported :  1 in_marshal
	// Bitfield not supported :  1 is_invalid
	// no type for marshal
	// Private : data
	// Private : notifiers
}

func ClosureNewFromC(u unsafe.Pointer) *Closure {
	c := (*C.GClosure)(u)
	if c == nil {
		return nil
	}

	g := &Closure{native: c}

	return g
}

func (recv *Closure) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Closure with another Closure, and returns true if they represent the same GObject.
func (recv *Closure) Equals(other *Closure) bool {
	return other.ToC() == recv.ToC()
}

// ClosureNewObject is a wrapper around the C function g_closure_new_object.
func ClosureNewObject(sizeofClosure uint32, object *Object) *Closure {
	c_sizeof_closure := (C.guint)(sizeofClosure)

	c_object := (*C.GObject)(C.NULL)
	if object != nil {
		c_object = (*C.GObject)(object.ToC())
	}

	retC := C.g_closure_new_object(c_sizeof_closure, c_object)
	retGo := ClosureNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ClosureNewSimple is a wrapper around the C function g_closure_new_simple.
func ClosureNewSimple(sizeofClosure uint32, data uintptr) *Closure {
	c_sizeof_closure := (C.guint)(sizeofClosure)

	c_data := (C.gpointer)(data)

	retC := C.g_closure_new_simple(c_sizeof_closure, c_data)
	retGo := ClosureNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_closure_add_finalize_notifier : unsupported parameter notify_func : no type generator for ClosureNotify (GClosureNotify) for param notify_func

// Unsupported : g_closure_add_invalidate_notifier : unsupported parameter notify_func : no type generator for ClosureNotify (GClosureNotify) for param notify_func

// Unsupported : g_closure_add_marshal_guards : unsupported parameter pre_marshal_notify : no type generator for ClosureNotify (GClosureNotify) for param pre_marshal_notify

// Invalidate is a wrapper around the C function g_closure_invalidate.
func (recv *Closure) Invalidate() {
	C.g_closure_invalidate((*C.GClosure)(recv.native))

	return
}

// Unsupported : g_closure_invoke : unsupported parameter param_values :

// Ref is a wrapper around the C function g_closure_ref.
func (recv *Closure) Ref() *Closure {
	retC := C.g_closure_ref((*C.GClosure)(recv.native))
	retGo := ClosureNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_closure_remove_finalize_notifier : unsupported parameter notify_func : no type generator for ClosureNotify (GClosureNotify) for param notify_func

// Unsupported : g_closure_remove_invalidate_notifier : unsupported parameter notify_func : no type generator for ClosureNotify (GClosureNotify) for param notify_func

// Unsupported : g_closure_set_marshal : unsupported parameter marshal : no type generator for ClosureMarshal (GClosureMarshal) for param marshal

// Unsupported : g_closure_set_meta_marshal : unsupported parameter meta_marshal : no type generator for ClosureMarshal (GClosureMarshal) for param meta_marshal

// Sink is a wrapper around the C function g_closure_sink.
func (recv *Closure) Sink() {
	C.g_closure_sink((*C.GClosure)(recv.native))

	return
}

// Unref is a wrapper around the C function g_closure_unref.
func (recv *Closure) Unref() {
	C.g_closure_unref((*C.GClosure)(recv.native))

	return
}

// ClosureNotifyData is a wrapper around the C record GClosureNotifyData.
type ClosureNotifyData struct {
	native *C.GClosureNotifyData
	Data   uintptr
	// notify : no type generator for ClosureNotify, GClosureNotify
}

func ClosureNotifyDataNewFromC(u unsafe.Pointer) *ClosureNotifyData {
	c := (*C.GClosureNotifyData)(u)
	if c == nil {
		return nil
	}

	g := &ClosureNotifyData{
		Data:   (uintptr)(c.data),
		native: c,
	}

	return g
}

func (recv *ClosureNotifyData) ToC() unsafe.Pointer {
	recv.native.data =
		(C.gpointer)(recv.Data)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ClosureNotifyData with another ClosureNotifyData, and returns true if they represent the same GObject.
func (recv *ClosureNotifyData) Equals(other *ClosureNotifyData) bool {
	return other.ToC() == recv.ToC()
}

// EnumClass is a wrapper around the C record GEnumClass.
type EnumClass struct {
	native *C.GEnumClass
	// g_type_class : record
	Minimum int32
	Maximum int32
	NValues uint32
	// values : record
}

func EnumClassNewFromC(u unsafe.Pointer) *EnumClass {
	c := (*C.GEnumClass)(u)
	if c == nil {
		return nil
	}

	g := &EnumClass{
		Maximum: (int32)(c.maximum),
		Minimum: (int32)(c.minimum),
		NValues: (uint32)(c.n_values),
		native:  c,
	}

	return g
}

func (recv *EnumClass) ToC() unsafe.Pointer {
	recv.native.minimum =
		(C.gint)(recv.Minimum)
	recv.native.maximum =
		(C.gint)(recv.Maximum)
	recv.native.n_values =
		(C.guint)(recv.NValues)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EnumClass with another EnumClass, and returns true if they represent the same GObject.
func (recv *EnumClass) Equals(other *EnumClass) bool {
	return other.ToC() == recv.ToC()
}

// EnumValue is a wrapper around the C record GEnumValue.
type EnumValue struct {
	native    *C.GEnumValue
	Value     int32
	ValueName string
	ValueNick string
}

func EnumValueNewFromC(u unsafe.Pointer) *EnumValue {
	c := (*C.GEnumValue)(u)
	if c == nil {
		return nil
	}

	g := &EnumValue{
		Value:     (int32)(c.value),
		ValueName: C.GoString(c.value_name),
		ValueNick: C.GoString(c.value_nick),
		native:    c,
	}

	return g
}

func (recv *EnumValue) ToC() unsafe.Pointer {
	recv.native.value =
		(C.gint)(recv.Value)
	recv.native.value_name =
		C.CString(recv.ValueName)
	recv.native.value_nick =
		C.CString(recv.ValueNick)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EnumValue with another EnumValue, and returns true if they represent the same GObject.
func (recv *EnumValue) Equals(other *EnumValue) bool {
	return other.ToC() == recv.ToC()
}

// FlagsClass is a wrapper around the C record GFlagsClass.
type FlagsClass struct {
	native *C.GFlagsClass
	// g_type_class : record
	Mask    uint32
	NValues uint32
	// values : record
}

func FlagsClassNewFromC(u unsafe.Pointer) *FlagsClass {
	c := (*C.GFlagsClass)(u)
	if c == nil {
		return nil
	}

	g := &FlagsClass{
		Mask:    (uint32)(c.mask),
		NValues: (uint32)(c.n_values),
		native:  c,
	}

	return g
}

func (recv *FlagsClass) ToC() unsafe.Pointer {
	recv.native.mask =
		(C.guint)(recv.Mask)
	recv.native.n_values =
		(C.guint)(recv.NValues)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FlagsClass with another FlagsClass, and returns true if they represent the same GObject.
func (recv *FlagsClass) Equals(other *FlagsClass) bool {
	return other.ToC() == recv.ToC()
}

// FlagsValue is a wrapper around the C record GFlagsValue.
type FlagsValue struct {
	native    *C.GFlagsValue
	Value     uint32
	ValueName string
	ValueNick string
}

func FlagsValueNewFromC(u unsafe.Pointer) *FlagsValue {
	c := (*C.GFlagsValue)(u)
	if c == nil {
		return nil
	}

	g := &FlagsValue{
		Value:     (uint32)(c.value),
		ValueName: C.GoString(c.value_name),
		ValueNick: C.GoString(c.value_nick),
		native:    c,
	}

	return g
}

func (recv *FlagsValue) ToC() unsafe.Pointer {
	recv.native.value =
		(C.guint)(recv.Value)
	recv.native.value_name =
		C.CString(recv.ValueName)
	recv.native.value_nick =
		C.CString(recv.ValueNick)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FlagsValue with another FlagsValue, and returns true if they represent the same GObject.
func (recv *FlagsValue) Equals(other *FlagsValue) bool {
	return other.ToC() == recv.ToC()
}

// InitiallyUnownedClass is a wrapper around the C record GInitiallyUnownedClass.
type InitiallyUnownedClass struct {
	native *C.GInitiallyUnownedClass
	// g_type_class : record
	// Private : construct_properties
	// no type for constructor
	// no type for set_property
	// no type for get_property
	// no type for dispose
	// no type for finalize
	// no type for dispatch_properties_changed
	// no type for notify
	// no type for constructed
	// Private : flags
	// Private : pdummy
}

func InitiallyUnownedClassNewFromC(u unsafe.Pointer) *InitiallyUnownedClass {
	c := (*C.GInitiallyUnownedClass)(u)
	if c == nil {
		return nil
	}

	g := &InitiallyUnownedClass{native: c}

	return g
}

func (recv *InitiallyUnownedClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InitiallyUnownedClass with another InitiallyUnownedClass, and returns true if they represent the same GObject.
func (recv *InitiallyUnownedClass) Equals(other *InitiallyUnownedClass) bool {
	return other.ToC() == recv.ToC()
}

// InterfaceInfo is a wrapper around the C record GInterfaceInfo.
type InterfaceInfo struct {
	native *C.GInterfaceInfo
	// interface_init : no type generator for InterfaceInitFunc, GInterfaceInitFunc
	// interface_finalize : no type generator for InterfaceFinalizeFunc, GInterfaceFinalizeFunc
	InterfaceData uintptr
}

func InterfaceInfoNewFromC(u unsafe.Pointer) *InterfaceInfo {
	c := (*C.GInterfaceInfo)(u)
	if c == nil {
		return nil
	}

	g := &InterfaceInfo{
		InterfaceData: (uintptr)(c.interface_data),
		native:        c,
	}

	return g
}

func (recv *InterfaceInfo) ToC() unsafe.Pointer {
	recv.native.interface_data =
		(C.gpointer)(recv.InterfaceData)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InterfaceInfo with another InterfaceInfo, and returns true if they represent the same GObject.
func (recv *InterfaceInfo) Equals(other *InterfaceInfo) bool {
	return other.ToC() == recv.ToC()
}

// ObjectClass is a wrapper around the C record GObjectClass.
type ObjectClass struct {
	native *C.GObjectClass
	// g_type_class : record
	// Private : construct_properties
	// no type for constructor
	// no type for set_property
	// no type for get_property
	// no type for dispose
	// no type for finalize
	// no type for dispatch_properties_changed
	// no type for notify
	// no type for constructed
	// Private : flags
	// Private : pdummy
}

func ObjectClassNewFromC(u unsafe.Pointer) *ObjectClass {
	c := (*C.GObjectClass)(u)
	if c == nil {
		return nil
	}

	g := &ObjectClass{native: c}

	return g
}

func (recv *ObjectClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ObjectClass with another ObjectClass, and returns true if they represent the same GObject.
func (recv *ObjectClass) Equals(other *ObjectClass) bool {
	return other.ToC() == recv.ToC()
}

// FindProperty is a wrapper around the C function g_object_class_find_property.
func (recv *ObjectClass) FindProperty(propertyName string) *ParamSpec {
	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	retC := C.g_object_class_find_property((*C.GObjectClass)(recv.native), c_property_name)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InstallProperty is a wrapper around the C function g_object_class_install_property.
func (recv *ObjectClass) InstallProperty(propertyId uint32, pspec *ParamSpec) {
	c_property_id := (C.guint)(propertyId)

	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	C.g_object_class_install_property((*C.GObjectClass)(recv.native), c_property_id, c_pspec)

	return
}

// Unsupported : g_object_class_list_properties : array return type :

// ObjectConstructParam is a wrapper around the C record GObjectConstructParam.
type ObjectConstructParam struct {
	native *C.GObjectConstructParam
	// pspec : record
	// value : record
}

func ObjectConstructParamNewFromC(u unsafe.Pointer) *ObjectConstructParam {
	c := (*C.GObjectConstructParam)(u)
	if c == nil {
		return nil
	}

	g := &ObjectConstructParam{native: c}

	return g
}

func (recv *ObjectConstructParam) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ObjectConstructParam with another ObjectConstructParam, and returns true if they represent the same GObject.
func (recv *ObjectConstructParam) Equals(other *ObjectConstructParam) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpecClass is a wrapper around the C record GParamSpecClass.
type ParamSpecClass struct {
	native *C.GParamSpecClass
	// g_type_class : record
	ValueType Type
	// no type for finalize
	// no type for value_set_default
	// no type for value_validate
	// no type for values_cmp
	// Private : dummy
}

func ParamSpecClassNewFromC(u unsafe.Pointer) *ParamSpecClass {
	c := (*C.GParamSpecClass)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecClass{
		ValueType: (Type)(c.value_type),
		native:    c,
	}

	return g
}

func (recv *ParamSpecClass) ToC() unsafe.Pointer {
	recv.native.value_type =
		(C.GType)(recv.ValueType)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecClass with another ParamSpecClass, and returns true if they represent the same GObject.
func (recv *ParamSpecClass) Equals(other *ParamSpecClass) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpecPool is a wrapper around the C record GParamSpecPool.
type ParamSpecPool struct {
	native *C.GParamSpecPool
}

func ParamSpecPoolNewFromC(u unsafe.Pointer) *ParamSpecPool {
	c := (*C.GParamSpecPool)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecPool{native: c}

	return g
}

func (recv *ParamSpecPool) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecPool with another ParamSpecPool, and returns true if they represent the same GObject.
func (recv *ParamSpecPool) Equals(other *ParamSpecPool) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpecPoolNew is a wrapper around the C function g_param_spec_pool_new.
func ParamSpecPoolNew(typePrefixing bool) *ParamSpecPool {
	c_type_prefixing :=
		boolToGboolean(typePrefixing)

	retC := C.g_param_spec_pool_new(c_type_prefixing)
	retGo := ParamSpecPoolNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Insert is a wrapper around the C function g_param_spec_pool_insert.
func (recv *ParamSpecPool) Insert(pspec *ParamSpec, ownerType Type) {
	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	c_owner_type := (C.GType)(ownerType)

	C.g_param_spec_pool_insert((*C.GParamSpecPool)(recv.native), c_pspec, c_owner_type)

	return
}

// Unsupported : g_param_spec_pool_list : array return type :

// ListOwned is a wrapper around the C function g_param_spec_pool_list_owned.
func (recv *ParamSpecPool) ListOwned(ownerType Type) *glib.List {
	c_owner_type := (C.GType)(ownerType)

	retC := C.g_param_spec_pool_list_owned((*C.GParamSpecPool)(recv.native), c_owner_type)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Lookup is a wrapper around the C function g_param_spec_pool_lookup.
func (recv *ParamSpecPool) Lookup(paramName string, ownerType Type, walkAncestors bool) *ParamSpec {
	c_param_name := C.CString(paramName)
	defer C.free(unsafe.Pointer(c_param_name))

	c_owner_type := (C.GType)(ownerType)

	c_walk_ancestors :=
		boolToGboolean(walkAncestors)

	retC := C.g_param_spec_pool_lookup((*C.GParamSpecPool)(recv.native), c_param_name, c_owner_type, c_walk_ancestors)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Remove is a wrapper around the C function g_param_spec_pool_remove.
func (recv *ParamSpecPool) Remove(pspec *ParamSpec) {
	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	C.g_param_spec_pool_remove((*C.GParamSpecPool)(recv.native), c_pspec)

	return
}

// ParamSpecTypeInfo is a wrapper around the C record GParamSpecTypeInfo.
type ParamSpecTypeInfo struct {
	native       *C.GParamSpecTypeInfo
	InstanceSize uint16
	NPreallocs   uint16
	// no type for instance_init
	ValueType Type
	// no type for finalize
	// no type for value_set_default
	// no type for value_validate
	// no type for values_cmp
}

func ParamSpecTypeInfoNewFromC(u unsafe.Pointer) *ParamSpecTypeInfo {
	c := (*C.GParamSpecTypeInfo)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecTypeInfo{
		InstanceSize: (uint16)(c.instance_size),
		NPreallocs:   (uint16)(c.n_preallocs),
		ValueType:    (Type)(c.value_type),
		native:       c,
	}

	return g
}

func (recv *ParamSpecTypeInfo) ToC() unsafe.Pointer {
	recv.native.instance_size =
		(C.guint16)(recv.InstanceSize)
	recv.native.n_preallocs =
		(C.guint16)(recv.NPreallocs)
	recv.native.value_type =
		(C.GType)(recv.ValueType)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecTypeInfo with another ParamSpecTypeInfo, and returns true if they represent the same GObject.
func (recv *ParamSpecTypeInfo) Equals(other *ParamSpecTypeInfo) bool {
	return other.ToC() == recv.ToC()
}

// Parameter is a wrapper around the C record GParameter.
type Parameter struct {
	native *C.GParameter
	Name   string
	// value : record
}

func ParameterNewFromC(u unsafe.Pointer) *Parameter {
	c := (*C.GParameter)(u)
	if c == nil {
		return nil
	}

	g := &Parameter{
		Name:   C.GoString(c.name),
		native: c,
	}

	return g
}

func (recv *Parameter) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Parameter with another Parameter, and returns true if they represent the same GObject.
func (recv *Parameter) Equals(other *Parameter) bool {
	return other.ToC() == recv.ToC()
}

// SignalInvocationHint is a wrapper around the C record GSignalInvocationHint.
type SignalInvocationHint struct {
	native   *C.GSignalInvocationHint
	SignalId uint32
	Detail   glib.Quark
	RunType  SignalFlags
}

func SignalInvocationHintNewFromC(u unsafe.Pointer) *SignalInvocationHint {
	c := (*C.GSignalInvocationHint)(u)
	if c == nil {
		return nil
	}

	g := &SignalInvocationHint{
		Detail:   (glib.Quark)(c.detail),
		RunType:  (SignalFlags)(c.run_type),
		SignalId: (uint32)(c.signal_id),
		native:   c,
	}

	return g
}

func (recv *SignalInvocationHint) ToC() unsafe.Pointer {
	recv.native.signal_id =
		(C.guint)(recv.SignalId)
	recv.native.detail =
		(C.GQuark)(recv.Detail)
	recv.native.run_type =
		(C.GSignalFlags)(recv.RunType)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SignalInvocationHint with another SignalInvocationHint, and returns true if they represent the same GObject.
func (recv *SignalInvocationHint) Equals(other *SignalInvocationHint) bool {
	return other.ToC() == recv.ToC()
}

// SignalQuery is a wrapper around the C record GSignalQuery.
type SignalQuery struct {
	native      *C.GSignalQuery
	SignalId    uint32
	SignalName  string
	Itype       Type
	SignalFlags SignalFlags
	ReturnType  Type
	NParams     uint32
	// no type for param_types
}

func SignalQueryNewFromC(u unsafe.Pointer) *SignalQuery {
	c := (*C.GSignalQuery)(u)
	if c == nil {
		return nil
	}

	g := &SignalQuery{
		Itype:       (Type)(c.itype),
		NParams:     (uint32)(c.n_params),
		ReturnType:  (Type)(c.return_type),
		SignalFlags: (SignalFlags)(c.signal_flags),
		SignalId:    (uint32)(c.signal_id),
		SignalName:  C.GoString(c.signal_name),
		native:      c,
	}

	return g
}

func (recv *SignalQuery) ToC() unsafe.Pointer {
	recv.native.signal_id =
		(C.guint)(recv.SignalId)
	recv.native.signal_name =
		C.CString(recv.SignalName)
	recv.native.itype =
		(C.GType)(recv.Itype)
	recv.native.signal_flags =
		(C.GSignalFlags)(recv.SignalFlags)
	recv.native.return_type =
		(C.GType)(recv.ReturnType)
	recv.native.n_params =
		(C.guint)(recv.NParams)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SignalQuery with another SignalQuery, and returns true if they represent the same GObject.
func (recv *SignalQuery) Equals(other *SignalQuery) bool {
	return other.ToC() == recv.ToC()
}

// TypeClass is a wrapper around the C record GTypeClass.
type TypeClass struct {
	native *C.GTypeClass
	// Private : g_type
}

func TypeClassNewFromC(u unsafe.Pointer) *TypeClass {
	c := (*C.GTypeClass)(u)
	if c == nil {
		return nil
	}

	g := &TypeClass{native: c}

	return g
}

func (recv *TypeClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypeClass with another TypeClass, and returns true if they represent the same GObject.
func (recv *TypeClass) Equals(other *TypeClass) bool {
	return other.ToC() == recv.ToC()
}

// TypeClassAdjustPrivateOffset is a wrapper around the C function g_type_class_adjust_private_offset.
func TypeClassAdjustPrivateOffset(gClass uintptr, privateSizeOrOffset int32) {
	c_g_class := (C.gpointer)(gClass)

	c_private_size_or_offset := (C.gint)(privateSizeOrOffset)

	C.g_type_class_adjust_private_offset(c_g_class, &c_private_size_or_offset)

	return
}

// TypeClassPeek is a wrapper around the C function g_type_class_peek.
func TypeClassPeek(type_ Type) uintptr {
	c_type := (C.GType)(type_)

	retC := C.g_type_class_peek(c_type)
	retGo := (uintptr)(retC)

	return retGo
}

// TypeClassRef is a wrapper around the C function g_type_class_ref.
func TypeClassRef(type_ Type) uintptr {
	c_type := (C.GType)(type_)

	retC := C.g_type_class_ref(c_type)
	retGo := (uintptr)(retC)

	return retGo
}

// GetPrivate is a wrapper around the C function g_type_class_get_private.
func (recv *TypeClass) GetPrivate(privateType Type) uintptr {
	c_private_type := (C.GType)(privateType)

	retC := C.g_type_class_get_private((*C.GTypeClass)(recv.native), c_private_type)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// PeekParent is a wrapper around the C function g_type_class_peek_parent.
func (recv *TypeClass) PeekParent() uintptr {
	retC := C.g_type_class_peek_parent((C.gpointer)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// Unref is a wrapper around the C function g_type_class_unref.
func (recv *TypeClass) Unref() {
	C.g_type_class_unref((C.gpointer)(recv.native))

	return
}

// UnrefUncached is a wrapper around the C function g_type_class_unref_uncached.
func (recv *TypeClass) UnrefUncached() {
	C.g_type_class_unref_uncached((C.gpointer)(recv.native))

	return
}

// TypeFundamentalInfo is a wrapper around the C record GTypeFundamentalInfo.
type TypeFundamentalInfo struct {
	native    *C.GTypeFundamentalInfo
	TypeFlags TypeFundamentalFlags
}

func TypeFundamentalInfoNewFromC(u unsafe.Pointer) *TypeFundamentalInfo {
	c := (*C.GTypeFundamentalInfo)(u)
	if c == nil {
		return nil
	}

	g := &TypeFundamentalInfo{
		TypeFlags: (TypeFundamentalFlags)(c.type_flags),
		native:    c,
	}

	return g
}

func (recv *TypeFundamentalInfo) ToC() unsafe.Pointer {
	recv.native.type_flags =
		(C.GTypeFundamentalFlags)(recv.TypeFlags)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypeFundamentalInfo with another TypeFundamentalInfo, and returns true if they represent the same GObject.
func (recv *TypeFundamentalInfo) Equals(other *TypeFundamentalInfo) bool {
	return other.ToC() == recv.ToC()
}

// TypeInfo is a wrapper around the C record GTypeInfo.
type TypeInfo struct {
	native    *C.GTypeInfo
	ClassSize uint16
	// base_init : no type generator for BaseInitFunc, GBaseInitFunc
	// base_finalize : no type generator for BaseFinalizeFunc, GBaseFinalizeFunc
	// class_init : no type generator for ClassInitFunc, GClassInitFunc
	// class_finalize : no type generator for ClassFinalizeFunc, GClassFinalizeFunc
	ClassData    uintptr
	InstanceSize uint16
	NPreallocs   uint16
	// instance_init : no type generator for InstanceInitFunc, GInstanceInitFunc
	// value_table : record
}

func TypeInfoNewFromC(u unsafe.Pointer) *TypeInfo {
	c := (*C.GTypeInfo)(u)
	if c == nil {
		return nil
	}

	g := &TypeInfo{
		ClassData:    (uintptr)(c.class_data),
		ClassSize:    (uint16)(c.class_size),
		InstanceSize: (uint16)(c.instance_size),
		NPreallocs:   (uint16)(c.n_preallocs),
		native:       c,
	}

	return g
}

func (recv *TypeInfo) ToC() unsafe.Pointer {
	recv.native.class_size =
		(C.guint16)(recv.ClassSize)
	recv.native.class_data =
		(C.gconstpointer)(recv.ClassData)
	recv.native.instance_size =
		(C.guint16)(recv.InstanceSize)
	recv.native.n_preallocs =
		(C.guint16)(recv.NPreallocs)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypeInfo with another TypeInfo, and returns true if they represent the same GObject.
func (recv *TypeInfo) Equals(other *TypeInfo) bool {
	return other.ToC() == recv.ToC()
}

// TypeInstance is a wrapper around the C record GTypeInstance.
type TypeInstance struct {
	native *C.GTypeInstance
	// Private : g_class
}

func TypeInstanceNewFromC(u unsafe.Pointer) *TypeInstance {
	c := (*C.GTypeInstance)(u)
	if c == nil {
		return nil
	}

	g := &TypeInstance{native: c}

	return g
}

func (recv *TypeInstance) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypeInstance with another TypeInstance, and returns true if they represent the same GObject.
func (recv *TypeInstance) Equals(other *TypeInstance) bool {
	return other.ToC() == recv.ToC()
}

// GetPrivate is a wrapper around the C function g_type_instance_get_private.
func (recv *TypeInstance) GetPrivate(privateType Type) uintptr {
	c_private_type := (C.GType)(privateType)

	retC := C.g_type_instance_get_private((*C.GTypeInstance)(recv.native), c_private_type)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// TypeInterface is a wrapper around the C record GTypeInterface.
type TypeInterface struct {
	native *C.GTypeInterface
	// Private : g_type
	// Private : g_instance_type
}

func TypeInterfaceNewFromC(u unsafe.Pointer) *TypeInterface {
	c := (*C.GTypeInterface)(u)
	if c == nil {
		return nil
	}

	g := &TypeInterface{native: c}

	return g
}

func (recv *TypeInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypeInterface with another TypeInterface, and returns true if they represent the same GObject.
func (recv *TypeInterface) Equals(other *TypeInterface) bool {
	return other.ToC() == recv.ToC()
}

// TypeInterfaceAddPrerequisite is a wrapper around the C function g_type_interface_add_prerequisite.
func TypeInterfaceAddPrerequisite(interfaceType Type, prerequisiteType Type) {
	c_interface_type := (C.GType)(interfaceType)

	c_prerequisite_type := (C.GType)(prerequisiteType)

	C.g_type_interface_add_prerequisite(c_interface_type, c_prerequisite_type)

	return
}

// TypeInterfaceGetPlugin is a wrapper around the C function g_type_interface_get_plugin.
func TypeInterfaceGetPlugin(instanceType Type, interfaceType Type) *TypePlugin {
	c_instance_type := (C.GType)(instanceType)

	c_interface_type := (C.GType)(interfaceType)

	retC := C.g_type_interface_get_plugin(c_instance_type, c_interface_type)
	retGo := TypePluginNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeInterfacePeek is a wrapper around the C function g_type_interface_peek.
func TypeInterfacePeek(instanceClass uintptr, ifaceType Type) uintptr {
	c_instance_class := (C.gpointer)(instanceClass)

	c_iface_type := (C.GType)(ifaceType)

	retC := C.g_type_interface_peek(c_instance_class, c_iface_type)
	retGo := (uintptr)(retC)

	return retGo
}

// PeekParent is a wrapper around the C function g_type_interface_peek_parent.
func (recv *TypeInterface) PeekParent() uintptr {
	retC := C.g_type_interface_peek_parent((C.gpointer)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// TypeModuleClass is a wrapper around the C record GTypeModuleClass.
type TypeModuleClass struct {
	native *C.GTypeModuleClass
	// parent_class : record
	// no type for load
	// no type for unload
	// no type for reserved1
	// no type for reserved2
	// no type for reserved3
	// no type for reserved4
}

func TypeModuleClassNewFromC(u unsafe.Pointer) *TypeModuleClass {
	c := (*C.GTypeModuleClass)(u)
	if c == nil {
		return nil
	}

	g := &TypeModuleClass{native: c}

	return g
}

func (recv *TypeModuleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypeModuleClass with another TypeModuleClass, and returns true if they represent the same GObject.
func (recv *TypeModuleClass) Equals(other *TypeModuleClass) bool {
	return other.ToC() == recv.ToC()
}

// TypePluginClass is a wrapper around the C record GTypePluginClass.
type TypePluginClass struct {
	native *C.GTypePluginClass
	// Private : base_iface
	// use_plugin : no type generator for TypePluginUse, GTypePluginUse
	// unuse_plugin : no type generator for TypePluginUnuse, GTypePluginUnuse
	// complete_type_info : no type generator for TypePluginCompleteTypeInfo, GTypePluginCompleteTypeInfo
	// complete_interface_info : no type generator for TypePluginCompleteInterfaceInfo, GTypePluginCompleteInterfaceInfo
}

func TypePluginClassNewFromC(u unsafe.Pointer) *TypePluginClass {
	c := (*C.GTypePluginClass)(u)
	if c == nil {
		return nil
	}

	g := &TypePluginClass{native: c}

	return g
}

func (recv *TypePluginClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypePluginClass with another TypePluginClass, and returns true if they represent the same GObject.
func (recv *TypePluginClass) Equals(other *TypePluginClass) bool {
	return other.ToC() == recv.ToC()
}

// TypeQuery is a wrapper around the C record GTypeQuery.
type TypeQuery struct {
	native       *C.GTypeQuery
	Type         Type
	TypeName     string
	ClassSize    uint32
	InstanceSize uint32
}

func TypeQueryNewFromC(u unsafe.Pointer) *TypeQuery {
	c := (*C.GTypeQuery)(u)
	if c == nil {
		return nil
	}

	g := &TypeQuery{
		ClassSize:    (uint32)(c.class_size),
		InstanceSize: (uint32)(c.instance_size),
		Type:         (Type)(c._type),
		TypeName:     C.GoString(c.type_name),
		native:       c,
	}

	return g
}

func (recv *TypeQuery) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GType)(recv.Type)
	recv.native.type_name =
		C.CString(recv.TypeName)
	recv.native.class_size =
		(C.guint)(recv.ClassSize)
	recv.native.instance_size =
		(C.guint)(recv.InstanceSize)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypeQuery with another TypeQuery, and returns true if they represent the same GObject.
func (recv *TypeQuery) Equals(other *TypeQuery) bool {
	return other.ToC() == recv.ToC()
}

// TypeValueTable is a wrapper around the C record GTypeValueTable.
type TypeValueTable struct {
	native *C.GTypeValueTable
	// no type for value_init
	// no type for value_free
	// no type for value_copy
	// no type for value_peek_pointer
	CollectFormat string
	// no type for collect_value
	LcopyFormat string
	// no type for lcopy_value
}

func TypeValueTableNewFromC(u unsafe.Pointer) *TypeValueTable {
	c := (*C.GTypeValueTable)(u)
	if c == nil {
		return nil
	}

	g := &TypeValueTable{
		CollectFormat: C.GoString(c.collect_format),
		LcopyFormat:   C.GoString(c.lcopy_format),
		native:        c,
	}

	return g
}

func (recv *TypeValueTable) ToC() unsafe.Pointer {
	recv.native.collect_format =
		C.CString(recv.CollectFormat)
	recv.native.lcopy_format =
		C.CString(recv.LcopyFormat)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypeValueTable with another TypeValueTable, and returns true if they represent the same GObject.
func (recv *TypeValueTable) Equals(other *TypeValueTable) bool {
	return other.ToC() == recv.ToC()
}

// TypeValueTablePeek is a wrapper around the C function g_type_value_table_peek.
func TypeValueTablePeek(type_ Type) *TypeValueTable {
	c_type := (C.GType)(type_)

	retC := C.g_type_value_table_peek(c_type)
	retGo := TypeValueTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Value is a wrapper around the C record GValue.
type Value struct {
	native *C.GValue
	// Private : g_type
	// no type for data
}

func ValueNewFromC(u unsafe.Pointer) *Value {
	c := (*C.GValue)(u)
	if c == nil {
		return nil
	}

	g := &Value{native: c}

	return g
}

func (recv *Value) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Value with another Value, and returns true if they represent the same GObject.
func (recv *Value) Equals(other *Value) bool {
	return other.ToC() == recv.ToC()
}

// g_value_register_transform_func : unsupported parameter transform_func : no type generator for ValueTransform (GValueTransform) for param transform_func
// ValueTypeCompatible is a wrapper around the C function g_value_type_compatible.
func ValueTypeCompatible(srcType Type, destType Type) bool {
	c_src_type := (C.GType)(srcType)

	c_dest_type := (C.GType)(destType)

	retC := C.g_value_type_compatible(c_src_type, c_dest_type)
	retGo := retC == C.TRUE

	return retGo
}

// ValueTypeTransformable is a wrapper around the C function g_value_type_transformable.
func ValueTypeTransformable(srcType Type, destType Type) bool {
	c_src_type := (C.GType)(srcType)

	c_dest_type := (C.GType)(destType)

	retC := C.g_value_type_transformable(c_src_type, c_dest_type)
	retGo := retC == C.TRUE

	return retGo
}

// Copy is a wrapper around the C function g_value_copy.
func (recv *Value) Copy(destValue *Value) {
	c_dest_value := (*C.GValue)(C.NULL)
	if destValue != nil {
		c_dest_value = (*C.GValue)(destValue.ToC())
	}

	C.g_value_copy((*C.GValue)(recv.native), c_dest_value)

	return
}

// DupBoxed is a wrapper around the C function g_value_dup_boxed.
func (recv *Value) DupBoxed() uintptr {
	retC := C.g_value_dup_boxed((*C.GValue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// DupObject is a wrapper around the C function g_value_dup_object.
func (recv *Value) DupObject() uintptr {
	retC := C.g_value_dup_object((*C.GValue)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// DupParam is a wrapper around the C function g_value_dup_param.
func (recv *Value) DupParam() *ParamSpec {
	retC := C.g_value_dup_param((*C.GValue)(recv.native))
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DupString is a wrapper around the C function g_value_dup_string.
func (recv *Value) DupString() string {
	retC := C.g_value_dup_string((*C.GValue)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// FitsPointer is a wrapper around the C function g_value_fits_pointer.
func (recv *Value) FitsPointer() bool {
	retC := C.g_value_fits_pointer((*C.GValue)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetBoolean is a wrapper around the C function g_value_get_boolean.
func (recv *Value) GetBoolean() bool {
	retC := C.g_value_get_boolean((*C.GValue)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetBoxed is a wrapper around the C function g_value_get_boxed.
func (recv *Value) GetBoxed() uintptr {
	retC := C.g_value_get_boxed((*C.GValue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// GetChar is a wrapper around the C function g_value_get_char.
func (recv *Value) GetChar() rune {
	retC := C.g_value_get_char((*C.GValue)(recv.native))
	retGo := (rune)(retC)

	return retGo
}

// GetDouble is a wrapper around the C function g_value_get_double.
func (recv *Value) GetDouble() float64 {
	retC := C.g_value_get_double((*C.GValue)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetEnum is a wrapper around the C function g_value_get_enum.
func (recv *Value) GetEnum() int32 {
	retC := C.g_value_get_enum((*C.GValue)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetFlags is a wrapper around the C function g_value_get_flags.
func (recv *Value) GetFlags() uint32 {
	retC := C.g_value_get_flags((*C.GValue)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetFloat is a wrapper around the C function g_value_get_float.
func (recv *Value) GetFloat() float32 {
	retC := C.g_value_get_float((*C.GValue)(recv.native))
	retGo := (float32)(retC)

	return retGo
}

// GetInt is a wrapper around the C function g_value_get_int.
func (recv *Value) GetInt() int32 {
	retC := C.g_value_get_int((*C.GValue)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetInt64 is a wrapper around the C function g_value_get_int64.
func (recv *Value) GetInt64() int64 {
	retC := C.g_value_get_int64((*C.GValue)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetLong is a wrapper around the C function g_value_get_long.
func (recv *Value) GetLong() int64 {
	retC := C.g_value_get_long((*C.GValue)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetObject is a wrapper around the C function g_value_get_object.
func (recv *Value) GetObject() uintptr {
	retC := C.g_value_get_object((*C.GValue)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// GetParam is a wrapper around the C function g_value_get_param.
func (recv *Value) GetParam() *ParamSpec {
	retC := C.g_value_get_param((*C.GValue)(recv.native))
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPointer is a wrapper around the C function g_value_get_pointer.
func (recv *Value) GetPointer() uintptr {
	retC := C.g_value_get_pointer((*C.GValue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// GetString is a wrapper around the C function g_value_get_string.
func (recv *Value) GetString() string {
	retC := C.g_value_get_string((*C.GValue)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUchar is a wrapper around the C function g_value_get_uchar.
func (recv *Value) GetUchar() uint8 {
	retC := C.g_value_get_uchar((*C.GValue)(recv.native))
	retGo := (uint8)(retC)

	return retGo
}

// GetUint is a wrapper around the C function g_value_get_uint.
func (recv *Value) GetUint() uint32 {
	retC := C.g_value_get_uint((*C.GValue)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetUint64 is a wrapper around the C function g_value_get_uint64.
func (recv *Value) GetUint64() uint64 {
	retC := C.g_value_get_uint64((*C.GValue)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// GetUlong is a wrapper around the C function g_value_get_ulong.
func (recv *Value) GetUlong() uint64 {
	retC := C.g_value_get_ulong((*C.GValue)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Init is a wrapper around the C function g_value_init.
func (recv *Value) Init(gType Type) *Value {
	c_g_type := (C.GType)(gType)

	retC := C.g_value_init((*C.GValue)(recv.native), c_g_type)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PeekPointer is a wrapper around the C function g_value_peek_pointer.
func (recv *Value) PeekPointer() uintptr {
	retC := C.g_value_peek_pointer((*C.GValue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Reset is a wrapper around the C function g_value_reset.
func (recv *Value) Reset() *Value {
	retC := C.g_value_reset((*C.GValue)(recv.native))
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetBoolean is a wrapper around the C function g_value_set_boolean.
func (recv *Value) SetBoolean(vBoolean bool) {
	c_v_boolean :=
		boolToGboolean(vBoolean)

	C.g_value_set_boolean((*C.GValue)(recv.native), c_v_boolean)

	return
}

// SetBoxed is a wrapper around the C function g_value_set_boxed.
func (recv *Value) SetBoxed(vBoxed uintptr) {
	c_v_boxed := (C.gconstpointer)(vBoxed)

	C.g_value_set_boxed((*C.GValue)(recv.native), c_v_boxed)

	return
}

// SetBoxedTakeOwnership is a wrapper around the C function g_value_set_boxed_take_ownership.
func (recv *Value) SetBoxedTakeOwnership(vBoxed uintptr) {
	c_v_boxed := (C.gconstpointer)(vBoxed)

	C.g_value_set_boxed_take_ownership((*C.GValue)(recv.native), c_v_boxed)

	return
}

// SetChar is a wrapper around the C function g_value_set_char.
func (recv *Value) SetChar(vChar rune) {
	c_v_char := (C.gchar)(vChar)

	C.g_value_set_char((*C.GValue)(recv.native), c_v_char)

	return
}

// SetDouble is a wrapper around the C function g_value_set_double.
func (recv *Value) SetDouble(vDouble float64) {
	c_v_double := (C.gdouble)(vDouble)

	C.g_value_set_double((*C.GValue)(recv.native), c_v_double)

	return
}

// SetEnum is a wrapper around the C function g_value_set_enum.
func (recv *Value) SetEnum(vEnum int32) {
	c_v_enum := (C.gint)(vEnum)

	C.g_value_set_enum((*C.GValue)(recv.native), c_v_enum)

	return
}

// SetFlags is a wrapper around the C function g_value_set_flags.
func (recv *Value) SetFlags(vFlags uint32) {
	c_v_flags := (C.guint)(vFlags)

	C.g_value_set_flags((*C.GValue)(recv.native), c_v_flags)

	return
}

// SetFloat is a wrapper around the C function g_value_set_float.
func (recv *Value) SetFloat(vFloat float32) {
	c_v_float := (C.gfloat)(vFloat)

	C.g_value_set_float((*C.GValue)(recv.native), c_v_float)

	return
}

// SetInstance is a wrapper around the C function g_value_set_instance.
func (recv *Value) SetInstance(instance uintptr) {
	c_instance := (C.gpointer)(instance)

	C.g_value_set_instance((*C.GValue)(recv.native), c_instance)

	return
}

// SetInt is a wrapper around the C function g_value_set_int.
func (recv *Value) SetInt(vInt int32) {
	c_v_int := (C.gint)(vInt)

	C.g_value_set_int((*C.GValue)(recv.native), c_v_int)

	return
}

// SetInt64 is a wrapper around the C function g_value_set_int64.
func (recv *Value) SetInt64(vInt64 int64) {
	c_v_int64 := (C.gint64)(vInt64)

	C.g_value_set_int64((*C.GValue)(recv.native), c_v_int64)

	return
}

// SetLong is a wrapper around the C function g_value_set_long.
func (recv *Value) SetLong(vLong int64) {
	c_v_long := (C.glong)(vLong)

	C.g_value_set_long((*C.GValue)(recv.native), c_v_long)

	return
}

// SetObject is a wrapper around the C function g_value_set_object.
func (recv *Value) SetObject(vObject uintptr) {
	c_v_object := (C.gpointer)(vObject)

	C.g_value_set_object((*C.GValue)(recv.native), c_v_object)

	return
}

// SetObjectTakeOwnership is a wrapper around the C function g_value_set_object_take_ownership.
func (recv *Value) SetObjectTakeOwnership(vObject uintptr) {
	c_v_object := (C.gpointer)(vObject)

	C.g_value_set_object_take_ownership((*C.GValue)(recv.native), c_v_object)

	return
}

// SetParam is a wrapper around the C function g_value_set_param.
func (recv *Value) SetParam(param *ParamSpec) {
	c_param := (*C.GParamSpec)(C.NULL)
	if param != nil {
		c_param = (*C.GParamSpec)(param.ToC())
	}

	C.g_value_set_param((*C.GValue)(recv.native), c_param)

	return
}

// SetParamTakeOwnership is a wrapper around the C function g_value_set_param_take_ownership.
func (recv *Value) SetParamTakeOwnership(param *ParamSpec) {
	c_param := (*C.GParamSpec)(C.NULL)
	if param != nil {
		c_param = (*C.GParamSpec)(param.ToC())
	}

	C.g_value_set_param_take_ownership((*C.GValue)(recv.native), c_param)

	return
}

// SetPointer is a wrapper around the C function g_value_set_pointer.
func (recv *Value) SetPointer(vPointer uintptr) {
	c_v_pointer := (C.gpointer)(vPointer)

	C.g_value_set_pointer((*C.GValue)(recv.native), c_v_pointer)

	return
}

// SetStaticBoxed is a wrapper around the C function g_value_set_static_boxed.
func (recv *Value) SetStaticBoxed(vBoxed uintptr) {
	c_v_boxed := (C.gconstpointer)(vBoxed)

	C.g_value_set_static_boxed((*C.GValue)(recv.native), c_v_boxed)

	return
}

// SetStaticString is a wrapper around the C function g_value_set_static_string.
func (recv *Value) SetStaticString(vString string) {
	c_v_string := C.CString(vString)
	defer C.free(unsafe.Pointer(c_v_string))

	C.g_value_set_static_string((*C.GValue)(recv.native), c_v_string)

	return
}

// SetString is a wrapper around the C function g_value_set_string.
func (recv *Value) SetString(vString string) {
	c_v_string := C.CString(vString)
	defer C.free(unsafe.Pointer(c_v_string))

	C.g_value_set_string((*C.GValue)(recv.native), c_v_string)

	return
}

// SetStringTakeOwnership is a wrapper around the C function g_value_set_string_take_ownership.
func (recv *Value) SetStringTakeOwnership(vString string) {
	c_v_string := C.CString(vString)
	defer C.free(unsafe.Pointer(c_v_string))

	C.g_value_set_string_take_ownership((*C.GValue)(recv.native), c_v_string)

	return
}

// SetUchar is a wrapper around the C function g_value_set_uchar.
func (recv *Value) SetUchar(vUchar uint8) {
	c_v_uchar := (C.guchar)(vUchar)

	C.g_value_set_uchar((*C.GValue)(recv.native), c_v_uchar)

	return
}

// SetUint is a wrapper around the C function g_value_set_uint.
func (recv *Value) SetUint(vUint uint32) {
	c_v_uint := (C.guint)(vUint)

	C.g_value_set_uint((*C.GValue)(recv.native), c_v_uint)

	return
}

// SetUint64 is a wrapper around the C function g_value_set_uint64.
func (recv *Value) SetUint64(vUint64 uint64) {
	c_v_uint64 := (C.guint64)(vUint64)

	C.g_value_set_uint64((*C.GValue)(recv.native), c_v_uint64)

	return
}

// SetUlong is a wrapper around the C function g_value_set_ulong.
func (recv *Value) SetUlong(vUlong uint64) {
	c_v_ulong := (C.gulong)(vUlong)

	C.g_value_set_ulong((*C.GValue)(recv.native), c_v_ulong)

	return
}

// Transform is a wrapper around the C function g_value_transform.
func (recv *Value) Transform(destValue *Value) bool {
	c_dest_value := (*C.GValue)(C.NULL)
	if destValue != nil {
		c_dest_value = (*C.GValue)(destValue.ToC())
	}

	retC := C.g_value_transform((*C.GValue)(recv.native), c_dest_value)
	retGo := retC == C.TRUE

	return retGo
}

// Unset is a wrapper around the C function g_value_unset.
func (recv *Value) Unset() {
	C.g_value_unset((*C.GValue)(recv.native))

	return
}

// ValueArray is a wrapper around the C record GValueArray.
type ValueArray struct {
	native  *C.GValueArray
	NValues uint32
	// values : record
	// Private : n_prealloced
}

func ValueArrayNewFromC(u unsafe.Pointer) *ValueArray {
	c := (*C.GValueArray)(u)
	if c == nil {
		return nil
	}

	g := &ValueArray{
		NValues: (uint32)(c.n_values),
		native:  c,
	}

	return g
}

func (recv *ValueArray) ToC() unsafe.Pointer {
	recv.native.n_values =
		(C.guint)(recv.NValues)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ValueArray with another ValueArray, and returns true if they represent the same GObject.
func (recv *ValueArray) Equals(other *ValueArray) bool {
	return other.ToC() == recv.ToC()
}

// ValueArrayNew is a wrapper around the C function g_value_array_new.
func ValueArrayNew(nPrealloced uint32) *ValueArray {
	c_n_prealloced := (C.guint)(nPrealloced)

	retC := C.g_value_array_new(c_n_prealloced)
	retGo := ValueArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Append is a wrapper around the C function g_value_array_append.
func (recv *ValueArray) Append(value *Value) *ValueArray {
	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.g_value_array_append((*C.GValueArray)(recv.native), c_value)
	retGo := ValueArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function g_value_array_copy.
func (recv *ValueArray) Copy() *ValueArray {
	retC := C.g_value_array_copy((*C.GValueArray)(recv.native))
	retGo := ValueArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function g_value_array_free.
func (recv *ValueArray) Free() {
	C.g_value_array_free((*C.GValueArray)(recv.native))

	return
}

// GetNth is a wrapper around the C function g_value_array_get_nth.
func (recv *ValueArray) GetNth(index uint32) *Value {
	c_index_ := (C.guint)(index)

	retC := C.g_value_array_get_nth((*C.GValueArray)(recv.native), c_index_)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Insert is a wrapper around the C function g_value_array_insert.
func (recv *ValueArray) Insert(index uint32, value *Value) *ValueArray {
	c_index_ := (C.guint)(index)

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.g_value_array_insert((*C.GValueArray)(recv.native), c_index_, c_value)
	retGo := ValueArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Prepend is a wrapper around the C function g_value_array_prepend.
func (recv *ValueArray) Prepend(value *Value) *ValueArray {
	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.g_value_array_prepend((*C.GValueArray)(recv.native), c_value)
	retGo := ValueArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Remove is a wrapper around the C function g_value_array_remove.
func (recv *ValueArray) Remove(index uint32) *ValueArray {
	c_index_ := (C.guint)(index)

	retC := C.g_value_array_remove((*C.GValueArray)(recv.native), c_index_)
	retGo := ValueArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_value_array_sort : unsupported parameter compare_func : no type generator for GLib.CompareFunc (GCompareFunc) for param compare_func

// Unsupported : g_value_array_sort_with_data : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc (GCompareDataFunc) for param compare_func

// WeakRef is a wrapper around the C record GWeakRef.
type WeakRef struct {
	native *C.GWeakRef
}

func WeakRefNewFromC(u unsafe.Pointer) *WeakRef {
	c := (*C.GWeakRef)(u)
	if c == nil {
		return nil
	}

	g := &WeakRef{native: c}

	return g
}

func (recv *WeakRef) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WeakRef with another WeakRef, and returns true if they represent the same GObject.
func (recv *WeakRef) Equals(other *WeakRef) bool {
	return other.ToC() == recv.ToC()
}
