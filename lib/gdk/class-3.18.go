// This is a generated file - DO NOT EDIT
// +build gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	gio "github.com/pekim/gobbi/lib/gio"
	gobject "github.com/pekim/gobbi/lib/gobject"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

func (recv *AppLaunchContext) AppLaunchContext() *gio.AppLaunchContext {}

func (recv *Cursor) Object() *gobject.Object {}

func (recv *Device) Object() *gobject.Object {}

func (recv *DeviceManager) Object() *gobject.Object {}

func (recv *DeviceTool) Object() *gobject.Object {}

func (recv *Display) Object() *gobject.Object {}

func (recv *DisplayManager) Object() *gobject.Object {}

func (recv *DragContext) Object() *gobject.Object {}

func (recv *DrawingContext) Object() *gobject.Object {}

func (recv *FrameClock) Object() *gobject.Object {}

func (recv *GLContext) Object() *gobject.Object {}

// GetScrollLockState is a wrapper around the C function gdk_keymap_get_scroll_lock_state.
func (recv *Keymap) GetScrollLockState() bool {
	retC := C.gdk_keymap_get_scroll_lock_state((*C.GdkKeymap)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

func (recv *Keymap) Object() *gobject.Object {}

func (recv *Monitor) Object() *gobject.Object {}

func (recv *Screen) Object() *gobject.Object {}

func (recv *Seat) Object() *gobject.Object {}

func (recv *Visual) Object() *gobject.Object {}

// GetPassThrough is a wrapper around the C function gdk_window_get_pass_through.
func (recv *Window) GetPassThrough() bool {
	retC := C.gdk_window_get_pass_through((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetPassThrough is a wrapper around the C function gdk_window_set_pass_through.
func (recv *Window) SetPassThrough(passThrough bool) {
	c_pass_through :=
		boolToGboolean(passThrough)

	C.gdk_window_set_pass_through((*C.GdkWindow)(recv.native), c_pass_through)

	return
}

func (recv *Window) Object() *gobject.Object {}
