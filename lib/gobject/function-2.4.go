// This is a generated file - DO NOT EDIT
// +build gobject_2.4 gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.44 gobject_2.54

package gobject

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_param_spec_override : unsupported parameter overridden : no type generator for ParamSpec, GParamSpec*

// SignalAccumulatorTrueHandled is a wrapper around the C function g_signal_accumulator_true_handled.
func SignalAccumulatorTrueHandled(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value, dummy uintptr) bool {
	c_ihint := (*C.GSignalInvocationHint)(ihint.ToC())

	c_return_accu := (*C.GValue)(returnAccu.ToC())

	c_handler_return := (*C.GValue)(handlerReturn.ToC())

	c_dummy := (C.gpointer)(dummy)

	retC := C.g_signal_accumulator_true_handled(c_ihint, c_return_accu, c_handler_return, c_dummy)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_type_add_interface_check : unsupported parameter check_func : no type generator for TypeInterfaceCheckFunc, GTypeInterfaceCheckFunc

// Unsupported : g_type_class_peek_static : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_type_default_interface_peek : unsupported parameter g_type : no type generator for GType, GType

// Unsupported : g_type_default_interface_ref : unsupported parameter g_type : no type generator for GType, GType

// Unsupported : g_type_default_interface_unref : no return generator

// Unsupported : g_type_remove_interface_check : unsupported parameter check_func : no type generator for TypeInterfaceCheckFunc, GTypeInterfaceCheckFunc
