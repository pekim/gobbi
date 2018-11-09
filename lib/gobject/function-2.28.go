// This is a generated file - DO NOT EDIT
// +build gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// ClearObject is a wrapper around the C function g_clear_object.
func ClearObject(objectPtr *Object) {
	c_object_ptr := (**C.GObject)(C.NULL)
	if objectPtr != nil {
		c_object_ptr = (**C.GObject)(objectPtr.ToC())
	}

	C.g_clear_object(c_object_ptr)

	return
}

// SignalAccumulatorFirstWins is a wrapper around the C function g_signal_accumulator_first_wins.
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
