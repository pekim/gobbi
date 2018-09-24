// This is a generated file - DO NOT EDIT
// +build gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.44 gobject_2.54

package gobject

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_clear_object : unsupported parameter object_ptr : no type generator for Object, volatile GObject**

// SignalAccumulatorFirstWins is a wrapper around the C function g_signal_accumulator_first_wins.
func SignalAccumulatorFirstWins(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value, dummy uintptr) bool {
	c_ihint := ihint.toC()

	c_return_accu := returnAccu.toC()

	c_handler_return := handlerReturn.toC()

	c_dummy := (C.gpointer)(dummy)

	retC := C.g_signal_accumulator_first_wins(c_ihint, c_return_accu, c_handler_return, c_dummy)
	retGo := (bool)(retC)

	return retGo
}
