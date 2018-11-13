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

// Creates a new cursor from a pixbuf.
//
// Not all GDK backends support RGBA cursors. If they are not
// supported, a monochrome approximation will be displayed.
// The functions gdk_display_supports_cursor_alpha() and
// gdk_display_supports_cursor_color() can be used to determine
// whether RGBA cursors are supported;
// gdk_display_get_default_cursor_size() and
// gdk_display_get_maximal_cursor_size() give information about
// cursor sizes.
//
// If @x or @y are `-1`, the pixbuf must have
// options named “x_hot” and “y_hot”, resp., containing
// integer values between `0` and the width resp. height of
// the pixbuf. (Since: 3.0)
//
// On the X backend, support for RGBA cursors requires a
// sufficently new version of the X Render extension.
/*

C function : gdk_cursor_new_from_pixbuf
*/
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

	return retGo
}

// Flushes any requests queued for the windowing system; this happens automatically
// when the main loop blocks waiting for new events, but if your application
// is drawing without returning control to the main loop, you may need
// to call this function explicitly. A common case where this function
// needs to be called is when an application is executing drawing commands
// from a thread other than the thread where the main loop is running.
//
// This is most useful for X11. On windowing systems where requests are
// handled synchronously, this function will do nothing.
/*

C function : gdk_display_flush
*/
func (recv *Display) Flush() {
	C.gdk_display_flush((*C.GdkDisplay)(recv.native))

	return
}

// Returns the default size to use for cursors on @display.
/*

C function : gdk_display_get_default_cursor_size
*/
func (recv *Display) GetDefaultCursorSize() uint32 {
	retC := C.gdk_display_get_default_cursor_size((*C.GdkDisplay)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Returns the default group leader window for all toplevel windows
// on @display. This window is implicitly created by GDK.
// See gdk_window_set_group().
/*

C function : gdk_display_get_default_group
*/
func (recv *Display) GetDefaultGroup() *Window {
	retC := C.gdk_display_get_default_group((*C.GdkDisplay)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the maximal size to use for cursors on @display.
/*

C function : gdk_display_get_maximal_cursor_size
*/
func (recv *Display) GetMaximalCursorSize() (uint32, uint32) {
	var c_width C.guint

	var c_height C.guint

	C.gdk_display_get_maximal_cursor_size((*C.GdkDisplay)(recv.native), &c_width, &c_height)

	width := (uint32)(c_width)

	height := (uint32)(c_height)

	return width, height
}

// Sets the double click distance (two clicks within this distance
// count as a double click and result in a #GDK_2BUTTON_PRESS event).
// See also gdk_display_set_double_click_time().
// Applications should not set this, it is a global
// user-configured setting.
/*

C function : gdk_display_set_double_click_distance
*/
func (recv *Display) SetDoubleClickDistance(distance uint32) {
	c_distance := (C.guint)(distance)

	C.gdk_display_set_double_click_distance((*C.GdkDisplay)(recv.native), c_distance)

	return
}

// Returns %TRUE if cursors can use an 8bit alpha channel
// on @display. Otherwise, cursors are restricted to bilevel
// alpha (i.e. a mask).
/*

C function : gdk_display_supports_cursor_alpha
*/
func (recv *Display) SupportsCursorAlpha() bool {
	retC := C.gdk_display_supports_cursor_alpha((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns %TRUE if multicolored cursors are supported
// on @display. Otherwise, cursors have only a forground
// and a background color.
/*

C function : gdk_display_supports_cursor_color
*/
func (recv *Display) SupportsCursorColor() bool {
	retC := C.gdk_display_supports_cursor_color((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the group leader window for @window. See gdk_window_set_group().
/*

C function : gdk_window_get_group
*/
func (recv *Window) GetGroup() *Window {
	retC := C.gdk_window_get_group((*C.GdkWindow)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Setting @accept_focus to %FALSE hints the desktop environment that the
// window doesn’t want to receive input focus.
//
// On X, it is the responsibility of the window manager to interpret this
// hint. ICCCM-compliant window manager usually respect it.
/*

C function : gdk_window_set_accept_focus
*/
func (recv *Window) SetAcceptFocus(acceptFocus bool) {
	c_accept_focus :=
		boolToGboolean(acceptFocus)

	C.gdk_window_set_accept_focus((*C.GdkWindow)(recv.native), c_accept_focus)

	return
}

// Set if @window must be kept above other windows. If the
// window was already above, then this function does nothing.
//
// On X11, asks the window manager to keep @window above, if the window
// manager supports this operation. Not all window managers support
// this, and some deliberately ignore it or don’t have a concept of
// “keep above”; so you can’t rely on the window being kept above.
// But it will happen with most standard window managers,
// and GDK makes a best effort to get it to happen.
/*

C function : gdk_window_set_keep_above
*/
func (recv *Window) SetKeepAbove(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gdk_window_set_keep_above((*C.GdkWindow)(recv.native), c_setting)

	return
}

// Set if @window must be kept below other windows. If the
// window was already below, then this function does nothing.
//
// On X11, asks the window manager to keep @window below, if the window
// manager supports this operation. Not all window managers support
// this, and some deliberately ignore it or don’t have a concept of
// “keep below”; so you can’t rely on the window being kept below.
// But it will happen with most standard window managers,
// and GDK makes a best effort to get it to happen.
/*

C function : gdk_window_set_keep_below
*/
func (recv *Window) SetKeepBelow(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gdk_window_set_keep_below((*C.GdkWindow)(recv.native), c_setting)

	return
}
