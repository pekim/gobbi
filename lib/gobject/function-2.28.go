// This is a generated file - DO NOT EDIT
// +build gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

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

// Unsupported : g_signal_accumulator_first_wins : unsupported parameter dummy : no type generator for gpointer (gpointer) for param dummy
