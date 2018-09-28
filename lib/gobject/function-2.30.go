// This is a generated file - DO NOT EDIT
// +build gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.44 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// CclosureMarshalGeneric is a wrapper around the C function g_cclosure_marshal_generic.
func CclosureMarshalGeneric(closure *Closure, returnGvalue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(closure.ToC())

	c_return_gvalue := (*C.GValue)(returnGvalue.ToC())

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(paramValues.ToC())

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_generic(c_closure, c_return_gvalue, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}
