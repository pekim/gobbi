// This is a generated file - DO NOT EDIT
// +build gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// GetModifierMask is a wrapper around the C function gdk_keymap_get_modifier_mask.
func (recv *Keymap) GetModifierMask(intent ModifierIntent) ModifierType {
	c_intent := (C.GdkModifierIntent)(intent)

	retC := C.gdk_keymap_get_modifier_mask((*C.GdkKeymap)(recv.native), c_intent)
	retGo := (ModifierType)(retC)

	return retGo
}

// GetModifierState is a wrapper around the C function gdk_keymap_get_modifier_state.
func (recv *Keymap) GetModifierState() uint32 {
	retC := C.gdk_keymap_get_modifier_state((*C.GdkKeymap)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetMonitorWorkarea is a wrapper around the C function gdk_screen_get_monitor_workarea.
func (recv *Screen) GetMonitorWorkarea(monitorNum int32) *Rectangle {
	c_monitor_num := (C.gint)(monitorNum)

	var c_dest C.GdkRectangle

	C.gdk_screen_get_monitor_workarea((*C.GdkScreen)(recv.native), c_monitor_num, &c_dest)

	dest := RectangleNewFromC(unsafe.Pointer(&c_dest))

	return dest
}

// BeginMoveDragForDevice is a wrapper around the C function gdk_window_begin_move_drag_for_device.
func (recv *Window) BeginMoveDragForDevice(device *Device, button int32, rootX int32, rootY int32, timestamp uint32) {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	c_button := (C.gint)(button)

	c_root_x := (C.gint)(rootX)

	c_root_y := (C.gint)(rootY)

	c_timestamp := (C.guint32)(timestamp)

	C.gdk_window_begin_move_drag_for_device((*C.GdkWindow)(recv.native), c_device, c_button, c_root_x, c_root_y, c_timestamp)

	return
}

// BeginResizeDragForDevice is a wrapper around the C function gdk_window_begin_resize_drag_for_device.
func (recv *Window) BeginResizeDragForDevice(edge WindowEdge, device *Device, button int32, rootX int32, rootY int32, timestamp uint32) {
	c_edge := (C.GdkWindowEdge)(edge)

	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	c_button := (C.gint)(button)

	c_root_x := (C.gint)(rootX)

	c_root_y := (C.gint)(rootY)

	c_timestamp := (C.guint32)(timestamp)

	C.gdk_window_begin_resize_drag_for_device((*C.GdkWindow)(recv.native), c_edge, c_device, c_button, c_root_x, c_root_y, c_timestamp)

	return
}
