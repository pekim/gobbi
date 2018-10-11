// This is a generated file - DO NOT EDIT
// +build gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

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

// GetNKeys is a wrapper around the C function gdk_device_get_n_keys.
func (recv *Device) GetNKeys() int32 {
	retC := C.gdk_device_get_n_keys((*C.GdkDevice)(recv.native))
	retGo := (int32)(retC)

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

// GetDisplay is a wrapper around the C function gdk_window_get_display.
func (recv *Window) GetDisplay() *Display {
	retC := C.gdk_window_get_display((*C.GdkWindow)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHeight is a wrapper around the C function gdk_window_get_height.
func (recv *Window) GetHeight() int32 {
	retC := C.gdk_window_get_height((*C.GdkWindow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetScreen is a wrapper around the C function gdk_window_get_screen.
func (recv *Window) GetScreen() *Screen {
	retC := C.gdk_window_get_screen((*C.GdkWindow)(recv.native))
	retGo := ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetVisual is a wrapper around the C function gdk_window_get_visual.
func (recv *Window) GetVisual() *Visual {
	retC := C.gdk_window_get_visual((*C.GdkWindow)(recv.native))
	retGo := VisualNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWidth is a wrapper around the C function gdk_window_get_width.
func (recv *Window) GetWidth() int32 {
	retC := C.gdk_window_get_width((*C.GdkWindow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

func (recv *Window) Object() *gobject.Object {}
