// +build gobject_2.4 gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.44 gobject_2.54

package gobject

// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// ParamSpecOverride is a wrapper around the C function g_param_spec_override.
func ParamSpecOverride(name int, overridden int) {}

// SignalAccumulatorTrueHandled is a wrapper around the C function g_signal_accumulator_true_handled.
func SignalAccumulatorTrueHandled(ihint int, returnAccu int, handlerReturn int, dummy int) {}

// TypeAddInterfaceCheck is a wrapper around the C function g_type_add_interface_check.
func TypeAddInterfaceCheck(checkData int, checkFunc int) {}

// TypeClassPeekStatic is a wrapper around the C function g_type_class_peek_static.
func TypeClassPeekStatic(type_ int) {}

// TypeDefaultInterfacePeek is a wrapper around the C function g_type_default_interface_peek.
func TypeDefaultInterfacePeek(gType int) {}

// TypeDefaultInterfaceRef is a wrapper around the C function g_type_default_interface_ref.
func TypeDefaultInterfaceRef(gType int) {}

// TypeDefaultInterfaceUnref is a wrapper around the C function g_type_default_interface_unref.
func TypeDefaultInterfaceUnref(gIface int) {}

// TypeRemoveInterfaceCheck is a wrapper around the C function g_type_remove_interface_check.
func TypeRemoveInterfaceCheck(checkData int, checkFunc int) {}
