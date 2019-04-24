// This is a generated file - DO NOT EDIT
// +build gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import call "github.com/pekim/gobbi/lib/internal/call"

// TypeGetTypeRegistrationSerial is a wrapper around the C function g_type_get_type_registration_serial.
func TypeGetTypeRegistrationSerial() uint32 {
	data := call.Data{Return: call.Return{Type: call.RT_UINT}}
	call.Function(3572, &data)
	ret := data.Return.Uint32()

	return ret
}
