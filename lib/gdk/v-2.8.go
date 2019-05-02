// This is a generated file - DO NOT EDIT
// +build gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// CursorNewFromName is a wrapper around the C function gdk_cursor_new_from_name.
func CursorNewFromName(display *Display, name string) *Cursor {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gdk_cursor_new_from_name(c_display, c_name)
	var retGo (*Cursor)
	if retC == nil {
		retGo = nil
	} else {
		retGo = CursorNewFromC(unsafe.Pointer(retC))
	}

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetImage is a wrapper around the C function gdk_cursor_get_image.
func (recv *Cursor) GetImage() *gdkpixbuf.Pixbuf {
	retC := C.gdk_cursor_get_image((*C.GdkCursor)(recv.native))
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// WarpPointer is a wrapper around the C function gdk_display_warp_pointer.
func (recv *Display) WarpPointer(screen *Screen, x int32, y int32) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gdk_display_warp_pointer((*C.GdkDisplay)(recv.native), c_screen, c_x, c_y)

	return
}

// GetRgbaVisual is a wrapper around the C function gdk_screen_get_rgba_visual.
func (recv *Screen) GetRgbaVisual() *Visual {
	retC := C.gdk_screen_get_rgba_visual((*C.GdkScreen)(recv.native))
	var retGo (*Visual)
	if retC == nil {
		retGo = nil
	} else {
		retGo = VisualNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// MoveRegion is a wrapper around the C function gdk_window_move_region.
func (recv *Window) MoveRegion(region *cairo.Region, dx int32, dy int32) {
	c_region := (*C.cairo_region_t)(C.NULL)
	if region != nil {
		c_region = (*C.cairo_region_t)(region.ToC())
	}

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

// CairoCreate is a wrapper around the C function gdk_cairo_create.
func CairoCreate(window *Window) *cairo.Context {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	retC := C.gdk_cairo_create(c_window)
	retGo := cairo.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CairoRectangle is a wrapper around the C function gdk_cairo_rectangle.
func CairoRectangle(cr *cairo.Context, rectangle *Rectangle) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_rectangle := (*C.GdkRectangle)(C.NULL)
	if rectangle != nil {
		c_rectangle = (*C.GdkRectangle)(rectangle.ToC())
	}

	C.gdk_cairo_rectangle(c_cr, c_rectangle)

	return
}

// CairoRegion is a wrapper around the C function gdk_cairo_region.
func CairoRegion(cr *cairo.Context, region *cairo.Region) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_region := (*C.cairo_region_t)(C.NULL)
	if region != nil {
		c_region = (*C.cairo_region_t)(region.ToC())
	}

	C.gdk_cairo_region(c_cr, c_region)

	return
}

// CairoSetSourceColor is a wrapper around the C function gdk_cairo_set_source_color.
func CairoSetSourceColor(cr *cairo.Context, color *Color) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_color := (*C.GdkColor)(C.NULL)
	if color != nil {
		c_color = (*C.GdkColor)(color.ToC())
	}

	C.gdk_cairo_set_source_color(c_cr, c_color)

	return
}

// CairoSetSourcePixbuf is a wrapper around the C function gdk_cairo_set_source_pixbuf.
func CairoSetSourcePixbuf(cr *cairo.Context, pixbuf *gdkpixbuf.Pixbuf, pixbufX float64, pixbufY float64) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	c_pixbuf_x := (C.gdouble)(pixbufX)

	c_pixbuf_y := (C.gdouble)(pixbufY)

	C.gdk_cairo_set_source_pixbuf(c_cr, c_pixbuf, c_pixbuf_x, c_pixbuf_y)

	return
}

// EventGrabBroken is a wrapper around the C record GdkEventGrabBroken.
type EventGrabBroken struct {
	native *C.GdkEventGrabBroken
	Type   EventType
	// window : record
	SendEvent int8
	Keyboard  bool
	Implicit  bool
	// grab_window : record
}

func EventGrabBrokenNewFromC(u unsafe.Pointer) *EventGrabBroken {
	c := (*C.GdkEventGrabBroken)(u)
	if c == nil {
		return nil
	}

	g := &EventGrabBroken{
		Implicit:  c.implicit == C.TRUE,
		Keyboard:  c.keyboard == C.TRUE,
		SendEvent: (int8)(c.send_event),
		Type:      (EventType)(c._type),
		native:    c,
	}

	return g
}

func (recv *EventGrabBroken) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.keyboard =
		boolToGboolean(recv.Keyboard)
	recv.native.implicit =
		boolToGboolean(recv.Implicit)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventGrabBroken with another EventGrabBroken, and returns true if they represent the same GObject.
func (recv *EventGrabBroken) Equals(other *EventGrabBroken) bool {
	return other.ToC() == recv.ToC()
}
