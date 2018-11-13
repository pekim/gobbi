// This is a generated file - DO NOT EDIT
// +build gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Initializes and sets @value from an instantiatable type via the
// value_table's collect_value() function.
//
// Note: The @value will be initialised with the exact type of
// @instance.  If you wish to set the @value's type to a different GType
// (such as a parent class GType), you need to manually call
// g_value_init() and g_value_set_instance().
/*

C function : g_value_init_from_instance
*/
func (recv *Value) InitFromInstance(instance uintptr) {
	c_instance := (C.gpointer)(instance)

	C.g_value_init_from_instance((*C.GValue)(recv.native), c_instance)

	return
}
