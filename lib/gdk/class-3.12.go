// This is a generated file - DO NOT EDIT
// +build gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	gio "github.com/pekim/gobbi/lib/gio"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

func (recv *AppLaunchContext) AppLaunchContext() *gio.AppLaunchContext {}

func (recv *Cursor) Object() *gobject.Object {}

// GetLastEventWindow is a wrapper around the C function gdk_device_get_last_event_window.
func (recv *Device) GetLastEventWindow() *Window {
	retC := C.gdk_device_get_last_event_window((*C.GdkDevice)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *Device) Object() *gobject.Object {}

func (recv *DeviceManager) Object() *gobject.Object {}

func (recv *DeviceTool) Object() *gobject.Object {}

func (recv *Display) Object() *gobject.Object {}

func (recv *DisplayManager) Object() *gobject.Object {}

func (recv *DragContext) Object() *gobject.Object {}

func (recv *DrawingContext) Object() *gobject.Object {}

func (recv *FrameClock) Object() *gobject.Object {}

func (recv *GLContext) Object() *gobject.Object {}

func (recv *Keymap) Object() *gobject.Object {}

func (recv *Monitor) Object() *gobject.Object {}

func (recv *Screen) Object() *gobject.Object {}

func (recv *Seat) Object() *gobject.Object {}

func (recv *Visual) Object() *gobject.Object {}

// GetEventCompression is a wrapper around the C function gdk_window_get_event_compression.
func (recv *Window) GetEventCompression() bool {
	retC := C.gdk_window_get_event_compression((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetEventCompression is a wrapper around the C function gdk_window_set_event_compression.
func (recv *Window) SetEventCompression(eventCompression bool) {
	c_event_compression :=
		boolToGboolean(eventCompression)

	C.gdk_window_set_event_compression((*C.GdkWindow)(recv.native), c_event_compression)

	return
}

// SetShadowWidth is a wrapper around the C function gdk_window_set_shadow_width.
func (recv *Window) SetShadowWidth(left int32, right int32, top int32, bottom int32) {
	c_left := (C.gint)(left)

	c_right := (C.gint)(right)

	c_top := (C.gint)(top)

	c_bottom := (C.gint)(bottom)

	C.gdk_window_set_shadow_width((*C.GdkWindow)(recv.native), c_left, c_right, c_top, c_bottom)

	return
}

func (recv *Window) Object() *gobject.Object {}
