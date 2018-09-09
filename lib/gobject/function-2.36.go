// +build gobject_2.36 gobject_2.44 gobject_2.54

package gobject

// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// TypeGetTypeRegistrationSerial is a wrapper around the C function g_type_get_type_registration_serial.
func TypeGetTypeRegistrationSerial() {}
