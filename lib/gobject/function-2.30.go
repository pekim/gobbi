// This is a generated file - DO NOT EDIT
// +build gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// A generic marshaller function implemented via
// [libffi](http://sourceware.org/libffi/).
//
// Normally this function is not passed explicitly to g_signal_new(),
// but used automatically by GLib when specifying a %NULL marshaller.
/*

C function : g_cclosure_marshal_generic
*/
func CclosureMarshalGeneric(closure *Closure, returnGvalue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_gvalue := (*C.GValue)(C.NULL)
	if returnGvalue != nil {
		c_return_gvalue = (*C.GValue)(returnGvalue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_generic(c_closure, c_return_gvalue, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}
