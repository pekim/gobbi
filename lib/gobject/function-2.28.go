// +build gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.44 gobject_2.54

package gobject

// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// ClearObject is a wrapper around the C function g_clear_object.
func ClearObject(objectPtr int) {}

// SignalAccumulatorFirstWins is a wrapper around the C function g_signal_accumulator_first_wins.
func SignalAccumulatorFirstWins(ihint int, returnAccu int, handlerReturn int, dummy int) {}
