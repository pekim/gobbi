// This is a generated file - DO NOT EDIT
// +build gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Returns an opaque serial number that represents the state of the set
// of registered types. Any time a type is registered this serial changes,
// which means you can cache information based on type lookups (such as
// g_type_from_name()) and know if the cache is still valid at a later
// time by comparing the current serial with the one at the type lookup.
/*

C function : g_type_get_type_registration_serial
*/
func TypeGetTypeRegistrationSerial() uint32 {
	retC := C.g_type_get_type_registration_serial()
	retGo := (uint32)(retC)

	return retGo
}
