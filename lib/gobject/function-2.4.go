// +build gobject_2.4 gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.44 gobject_2.54

package gobject

// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Unsupported function: g_param_spec_override : unsupported parameter overridden : type ParamSpec, GParamSpec*

// Unsupported function: g_signal_accumulator_true_handled : unsupported parameter ihint : type SignalInvocationHint, GSignalInvocationHint*

// Unsupported function: g_type_add_interface_check : unsupported parameter check_func : type TypeInterfaceCheckFunc, GTypeInterfaceCheckFunc

// Unsupported function: g_type_class_peek_static : unsupported parameter type : type GType, GType

// Unsupported function: g_type_default_interface_peek : unsupported parameter g_type : type GType, GType

// Unsupported function: g_type_default_interface_ref : unsupported parameter g_type : type GType, GType

// TypeDefaultInterfaceUnref is a wrapper around the C function g_type_default_interface_unref.
func TypeDefaultInterfaceUnref(gIface uintptr) {}

// Unsupported function: g_type_remove_interface_check : unsupported parameter check_func : type TypeInterfaceCheckFunc, GTypeInterfaceCheckFunc
