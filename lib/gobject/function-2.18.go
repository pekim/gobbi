// +build gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.44 gobject_2.54

package gobject

// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Unsupported function: g_signal_chain_from_overridden_handler : unsupported parameter ... : varargs

// Unsupported function: g_signal_new_class_handler : unsupported parameter ... : varargs

// SignalOverrideClassHandler is a wrapper around the C function g_signal_override_class_handler.
func SignalOverrideClassHandler(signalName int, instanceType int, classHandler int) {}
