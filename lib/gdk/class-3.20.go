// This is a generated file - DO NOT EDIT
// +build gdk_3.20 gdk_3.22

package gdk

import (
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

func (recv *AppLaunchContext) AppLaunchContext() *gio.AppLaunchContext {}

func (recv *Cursor) Object() *gobject.Object {}

// GetSeat is a wrapper around the C function gdk_device_get_seat.
func (recv *Device) GetSeat() *Seat {
	retC := C.gdk_device_get_seat((*C.GdkDevice)(recv.native))
	retGo := SeatNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *Device) Object() *gobject.Object {}

func (recv *DeviceManager) Object() *gobject.Object {}

func (recv *DeviceTool) Object() *gobject.Object {}

// GetDefaultSeat is a wrapper around the C function gdk_display_get_default_seat.
func (recv *Display) GetDefaultSeat() *Seat {
	retC := C.gdk_display_get_default_seat((*C.GdkDisplay)(recv.native))
	retGo := SeatNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListSeats is a wrapper around the C function gdk_display_list_seats.
func (recv *Display) ListSeats() *glib.List {
	retC := C.gdk_display_list_seats((*C.GdkDisplay)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *Display) Object() *gobject.Object {}

func (recv *DisplayManager) Object() *gobject.Object {}

// GetDragWindow is a wrapper around the C function gdk_drag_context_get_drag_window.
func (recv *DragContext) GetDragWindow() *Window {
	retC := C.gdk_drag_context_get_drag_window((*C.GdkDragContext)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ManageDnd is a wrapper around the C function gdk_drag_context_manage_dnd.
func (recv *DragContext) ManageDnd(ipcWindow *Window, actions DragAction) bool {
	c_ipc_window := (*C.GdkWindow)(ipcWindow.ToC())

	c_actions := (C.GdkDragAction)(actions)

	retC := C.gdk_drag_context_manage_dnd((*C.GdkDragContext)(recv.native), c_ipc_window, c_actions)
	retGo := retC == C.TRUE

	return retGo
}

// SetHotspot is a wrapper around the C function gdk_drag_context_set_hotspot.
func (recv *DragContext) SetHotspot(hotX int32, hotY int32) {
	c_hot_x := (C.gint)(hotX)

	c_hot_y := (C.gint)(hotY)

	C.gdk_drag_context_set_hotspot((*C.GdkDragContext)(recv.native), c_hot_x, c_hot_y)

	return
}

func (recv *DragContext) Object() *gobject.Object {}

func (recv *DrawingContext) Object() *gobject.Object {}

func (recv *FrameClock) Object() *gobject.Object {}

// IsLegacy is a wrapper around the C function gdk_gl_context_is_legacy.
func (recv *GLContext) IsLegacy() bool {
	retC := C.gdk_gl_context_is_legacy((*C.GdkGLContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

func (recv *GLContext) Object() *gobject.Object {}

func (recv *Keymap) Object() *gobject.Object {}

func (recv *Monitor) Object() *gobject.Object {}

func (recv *Screen) Object() *gobject.Object {}

// GetCapabilities is a wrapper around the C function gdk_seat_get_capabilities.
func (recv *Seat) GetCapabilities() SeatCapabilities {
	retC := C.gdk_seat_get_capabilities((*C.GdkSeat)(recv.native))
	retGo := (SeatCapabilities)(retC)

	return retGo
}

// GetKeyboard is a wrapper around the C function gdk_seat_get_keyboard.
func (recv *Seat) GetKeyboard() *Device {
	retC := C.gdk_seat_get_keyboard((*C.GdkSeat)(recv.native))
	retGo := DeviceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPointer is a wrapper around the C function gdk_seat_get_pointer.
func (recv *Seat) GetPointer() *Device {
	retC := C.gdk_seat_get_pointer((*C.GdkSeat)(recv.native))
	retGo := DeviceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSlaves is a wrapper around the C function gdk_seat_get_slaves.
func (recv *Seat) GetSlaves(capabilities SeatCapabilities) *glib.List {
	c_capabilities := (C.GdkSeatCapabilities)(capabilities)

	retC := C.gdk_seat_get_slaves((*C.GdkSeat)(recv.native), c_capabilities)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_seat_grab : unsupported parameter event : no type generator for Event, const GdkEvent*

// Ungrab is a wrapper around the C function gdk_seat_ungrab.
func (recv *Seat) Ungrab() {
	C.gdk_seat_ungrab((*C.GdkSeat)(recv.native))

	return
}

func (recv *Seat) Object() *gobject.Object {}

func (recv *Visual) Object() *gobject.Object {}

func (recv *Window) Object() *gobject.Object {}
