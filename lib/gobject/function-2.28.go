// This is a generated file - DO NOT EDIT
// +build gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.44 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_clear_object : unsupported parameter object_ptr : in string with indirection level of 2

// SignalAccumulatorFirstWins is a wrapper around the C function g_signal_accumulator_first_wins.
func SignalAccumulatorFirstWins(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value, dummy uintptr) bool {
	c_ihint := (*C.GSignalInvocationHint)(ihint.ToC())

	c_return_accu := (*C.GValue)(returnAccu.ToC())

	c_handler_return := (*C.GValue)(handlerReturn.ToC())

	c_dummy := (C.gpointer)(dummy)

	retC := C.g_signal_accumulator_first_wins(c_ihint, c_return_accu, c_handler_return, c_dummy)
	retGo := retC == C.TRUE

	return retGo
}
