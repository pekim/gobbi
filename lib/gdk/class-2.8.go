// This is a generated file - DO NOT EDIT
// +build gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gio "github.com/pekim/gobbi/lib/gio"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

func (recv *AppLaunchContext) AppLaunchContext() *gio.AppLaunchContext {}

// CursorNewFromName is a wrapper around the C function gdk_cursor_new_from_name.
func CursorNewFromName(display *Display, name string) *Cursor {
	c_display := (*C.GdkDisplay)(display.ToC())

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gdk_cursor_new_from_name(c_display, c_name)
	retGo := CursorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetImage is a wrapper around the C function gdk_cursor_get_image.
func (recv *Cursor) GetImage() *gdkpixbuf.Pixbuf {
	retC := C.gdk_cursor_get_image((*C.GdkCursor)(recv.native))
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *Cursor) Object() *gobject.Object {}

func (recv *Device) Object() *gobject.Object {}

func (recv *DeviceManager) Object() *gobject.Object {}

func (recv *DeviceTool) Object() *gobject.Object {}

// WarpPointer is a wrapper around the C function gdk_display_warp_pointer.
func (recv *Display) WarpPointer(screen *Screen, x int32, y int32) {
	c_screen := (*C.GdkScreen)(screen.ToC())

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gdk_display_warp_pointer((*C.GdkDisplay)(recv.native), c_screen, c_x, c_y)

	return
}

func (recv *Display) Object() *gobject.Object {}

func (recv *DisplayManager) Object() *gobject.Object {}

func (recv *DragContext) Object() *gobject.Object {}

func (recv *DrawingContext) Object() *gobject.Object {}

func (recv *FrameClock) Object() *gobject.Object {}

func (recv *GLContext) Object() *gobject.Object {}

func (recv *Keymap) Object() *gobject.Object {}

func (recv *Monitor) Object() *gobject.Object {}

// GetRgbaVisual is a wrapper around the C function gdk_screen_get_rgba_visual.
func (recv *Screen) GetRgbaVisual() *Visual {
	retC := C.gdk_screen_get_rgba_visual((*C.GdkScreen)(recv.native))
	retGo := VisualNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *Screen) Object() *gobject.Object {}

func (recv *Seat) Object() *gobject.Object {}

func (recv *Visual) Object() *gobject.Object {}

// MoveRegion is a wrapper around the C function gdk_window_move_region.
func (recv *Window) MoveRegion(region *cairo.Region, dx int32, dy int32) {
	c_region := (*C.cairo_region_t)(region.ToC())

	c_dx := (C.gint)(dx)

	c_dy := (C.gint)(dy)

	C.gdk_window_move_region((*C.GdkWindow)(recv.native), c_region, c_dx, c_dy)

	return
}

// SetUrgencyHint is a wrapper around the C function gdk_window_set_urgency_hint.
func (recv *Window) SetUrgencyHint(urgent bool) {
	c_urgent :=
		boolToGboolean(urgent)

	C.gdk_window_set_urgency_hint((*C.GdkWindow)(recv.native), c_urgent)

	return
}

func (recv *Window) Object() *gobject.Object {}
