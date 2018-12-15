// This is a generated file - DO NOT EDIT
// +build gdk_2.4 gdk_2.6 gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// CursorNewFromPixbuf is a wrapper around the C function gdk_cursor_new_from_pixbuf.
func CursorNewFromPixbuf(display *Display, pixbuf *gdkpixbuf.Pixbuf, x int32, y int32) *Cursor {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gdk_cursor_new_from_pixbuf(c_display, c_pixbuf, c_x, c_y)
	retGo := CursorNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Flush is a wrapper around the C function gdk_display_flush.
func (recv *Display) Flush() {
	C.gdk_display_flush((*C.GdkDisplay)(recv.native))

	return
}

// GetDefaultCursorSize is a wrapper around the C function gdk_display_get_default_cursor_size.
func (recv *Display) GetDefaultCursorSize() uint32 {
	retC := C.gdk_display_get_default_cursor_size((*C.GdkDisplay)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetDefaultGroup is a wrapper around the C function gdk_display_get_default_group.
func (recv *Display) GetDefaultGroup() *Window {
	retC := C.gdk_display_get_default_group((*C.GdkDisplay)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMaximalCursorSize is a wrapper around the C function gdk_display_get_maximal_cursor_size.
func (recv *Display) GetMaximalCursorSize() (uint32, uint32) {
	var c_width C.guint

	var c_height C.guint

	C.gdk_display_get_maximal_cursor_size((*C.GdkDisplay)(recv.native), &c_width, &c_height)

	width := (uint32)(c_width)

	height := (uint32)(c_height)

	return width, height
}

// SetDoubleClickDistance is a wrapper around the C function gdk_display_set_double_click_distance.
func (recv *Display) SetDoubleClickDistance(distance uint32) {
	c_distance := (C.guint)(distance)

	C.gdk_display_set_double_click_distance((*C.GdkDisplay)(recv.native), c_distance)

	return
}

// SupportsCursorAlpha is a wrapper around the C function gdk_display_supports_cursor_alpha.
func (recv *Display) SupportsCursorAlpha() bool {
	retC := C.gdk_display_supports_cursor_alpha((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SupportsCursorColor is a wrapper around the C function gdk_display_supports_cursor_color.
func (recv *Display) SupportsCursorColor() bool {
	retC := C.gdk_display_supports_cursor_color((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetGroup is a wrapper around the C function gdk_window_get_group.
func (recv *Window) GetGroup() *Window {
	retC := C.gdk_window_get_group((*C.GdkWindow)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetAcceptFocus is a wrapper around the C function gdk_window_set_accept_focus.
func (recv *Window) SetAcceptFocus(acceptFocus bool) {
	c_accept_focus :=
		boolToGboolean(acceptFocus)

	C.gdk_window_set_accept_focus((*C.GdkWindow)(recv.native), c_accept_focus)

	return
}

// SetKeepAbove is a wrapper around the C function gdk_window_set_keep_above.
func (recv *Window) SetKeepAbove(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gdk_window_set_keep_above((*C.GdkWindow)(recv.native), c_setting)

	return
}

// SetKeepBelow is a wrapper around the C function gdk_window_set_keep_below.
func (recv *Window) SetKeepBelow(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gdk_window_set_keep_below((*C.GdkWindow)(recv.native), c_setting)

	return
}
