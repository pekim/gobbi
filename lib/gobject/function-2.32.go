// +build gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.44 gobject_2.54

package gobject

// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// SignalSetVaMarshaller is a wrapper around the C function g_signal_set_va_marshaller.
func SignalSetVaMarshaller(signalId int, instanceType int, vaMarshaller int) {}
