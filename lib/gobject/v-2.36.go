// Code generated - DO NOT EDIT.
// +build gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// TypeGetTypeRegistrationSerial is a wrapper around the C function g_type_get_type_registration_serial.
func TypeGetTypeRegistrationSerial() uint32 {
	retC := C.g_type_get_type_registration_serial()
	retGo := (uint32)(retC)

	return retGo
}
