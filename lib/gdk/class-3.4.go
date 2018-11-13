// This is a generated file - DO NOT EDIT
// +build gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Returns the modifier mask the @keymap’s windowing system backend
// uses for a particular purpose.
//
// Note that this function always returns real hardware modifiers, not
// virtual ones (e.g. it will return #GDK_MOD1_MASK rather than
// #GDK_META_MASK if the backend maps MOD1 to META), so there are use
// cases where the return value of this function has to be transformed
// by gdk_keymap_add_virtual_modifiers() in order to contain the
// expected result.
/*

C function : gdk_keymap_get_modifier_mask
*/
func (recv *Keymap) GetModifierMask(intent ModifierIntent) ModifierType {
	c_intent := (C.GdkModifierIntent)(intent)

	retC := C.gdk_keymap_get_modifier_mask((*C.GdkKeymap)(recv.native), c_intent)
	retGo := (ModifierType)(retC)

	return retGo
}

// Returns the current modifier state.
/*

C function : gdk_keymap_get_modifier_state
*/
func (recv *Keymap) GetModifierState() uint32 {
	retC := C.gdk_keymap_get_modifier_state((*C.GdkKeymap)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : gdk_screen_get_monitor_workarea : unsupported parameter dest : Blacklisted record : GdkRectangle

// Begins a window move operation (for a toplevel window).
// You might use this function to implement a “window move grip,” for
// example. The function works best with window managers that support the
// [Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec)
// but has a fallback implementation for other window managers.
/*

C function : gdk_window_begin_move_drag_for_device
*/
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

// Begins a window resize operation (for a toplevel window).
// You might use this function to implement a “window resize grip,” for
// example; in fact #GtkStatusbar uses it. The function works best
// with window managers that support the
// [Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec)
// but has a fallback implementation for other window managers.
/*

C function : gdk_window_begin_resize_drag_for_device
*/
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
