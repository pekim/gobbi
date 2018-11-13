// This is a generated file - DO NOT EDIT
// +build gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Returns the axis use for @index_.
/*

C function

gdk_device_get_axis_use
*/
func (recv *Device) GetAxisUse(index uint32) AxisUse {
	c_index_ := (C.guint)(index)

	retC := C.gdk_device_get_axis_use((*C.GdkDevice)(recv.native), c_index_)
	retGo := (AxisUse)(retC)

	return retGo
}

// Determines whether the pointer follows device motion.
// This is not meaningful for keyboard devices, which don't have a pointer.
/*

C function

gdk_device_get_has_cursor
*/
func (recv *Device) GetHasCursor() bool {
	retC := C.gdk_device_get_has_cursor((*C.GdkDevice)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_device_get_key : unsupported parameter modifiers : GdkModifierType* with indirection level of 1

// Determines the mode of the device.
/*

C function

gdk_device_get_mode
*/
func (recv *Device) GetMode() InputMode {
	retC := C.gdk_device_get_mode((*C.GdkDevice)(recv.native))
	retGo := (InputMode)(retC)

	return retGo
}

// Determines the name of the device.
/*

C function

gdk_device_get_name
*/
func (recv *Device) GetName() string {
	retC := C.gdk_device_get_name((*C.GdkDevice)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Determines the type of the device.
/*

C function

gdk_device_get_source
*/
func (recv *Device) GetSource() InputSource {
	retC := C.gdk_device_get_source((*C.GdkDevice)(recv.native))
	retGo := (InputSource)(retC)

	return retGo
}

// Unsupported : gdk_keymap_add_virtual_modifiers : unsupported parameter state : GdkModifierType* with indirection level of 1

// Unsupported : gdk_keymap_map_virtual_modifiers : unsupported parameter state : GdkModifierType* with indirection level of 1

// Gets the primary monitor for @screen.  The primary monitor
// is considered the monitor where the “main desktop” lives.
// While normal application windows typically allow the window
// manager to place the windows, specialized desktop applications
// such as panels should place themselves on the primary monitor.
//
// If no primary monitor is configured by the user, the return value
// will be 0, defaulting to the first monitor.
/*

C function

gdk_screen_get_primary_monitor
*/
func (recv *Screen) GetPrimaryMonitor() int32 {
	retC := C.gdk_screen_get_primary_monitor((*C.GdkScreen)(recv.native))
	retGo := (int32)(retC)

	return retGo
}
