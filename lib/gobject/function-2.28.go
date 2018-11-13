// This is a generated file - DO NOT EDIT
// +build gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Clears a reference to a #GObject.
//
// @object_ptr must not be %NULL.
//
// If the reference is %NULL then this function does nothing.
// Otherwise, the reference count of the object is decreased and the
// pointer is set to %NULL.
//
// A macro is also included that allows this function to be used without
// pointer casts.
/*

C function : g_clear_object
*/
func ClearObject(objectPtr *Object) {
	c_object_ptr := (**C.GObject)(C.NULL)
	if objectPtr != nil {
		c_object_ptr = (**C.GObject)(objectPtr.ToC())
	}

	C.g_clear_object(c_object_ptr)

	return
}

// A predefined #GSignalAccumulator for signals intended to be used as a
// hook for application code to provide a particular value.  Usually
// only one such value is desired and multiple handlers for the same
// signal don't make much sense (except for the case of the default
// handler defined in the class structure, in which case you will
// usually want the signal connection to override the class handler).
//
// This accumulator will use the return value from the first signal
// handler that is run as the return value for the signal and not run
// any further handlers (ie: the first handler "wins").
/*

C function : g_signal_accumulator_first_wins
*/
func SignalAccumulatorFirstWins(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value, dummy uintptr) bool {
	c_ihint := (*C.GSignalInvocationHint)(C.NULL)
	if ihint != nil {
		c_ihint = (*C.GSignalInvocationHint)(ihint.ToC())
	}

	c_return_accu := (*C.GValue)(C.NULL)
	if returnAccu != nil {
		c_return_accu = (*C.GValue)(returnAccu.ToC())
	}

	c_handler_return := (*C.GValue)(C.NULL)
	if handlerReturn != nil {
		c_handler_return = (*C.GValue)(handlerReturn.ToC())
	}

	c_dummy := (C.gpointer)(dummy)

	retC := C.g_signal_accumulator_first_wins(c_ihint, c_return_accu, c_handler_return, c_dummy)
	retGo := retC == C.TRUE

	return retGo
}
