// This is a generated file - DO NOT EDIT
// +build gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Creates a new cursor by looking up @name in the current cursor
// theme.
//
// A recommended set of cursor names that will work across different
// platforms can be found in the CSS specification:
// - "none"
// - ![](default_cursor.png) "default"
// - ![](help_cursor.png) "help"
// - ![](pointer_cursor.png) "pointer"
// - ![](context_menu_cursor.png) "context-menu"
// - ![](progress_cursor.png) "progress"
// - ![](wait_cursor.png) "wait"
// - ![](cell_cursor.png) "cell"
// - ![](crosshair_cursor.png) "crosshair"
// - ![](text_cursor.png) "text"
// - ![](vertical_text_cursor.png) "vertical-text"
// - ![](alias_cursor.png) "alias"
// - ![](copy_cursor.png) "copy"
// - ![](no_drop_cursor.png) "no-drop"
// - ![](move_cursor.png) "move"
// - ![](not_allowed_cursor.png) "not-allowed"
// - ![](grab_cursor.png) "grab"
// - ![](grabbing_cursor.png) "grabbing"
// - ![](all_scroll_cursor.png) "all-scroll"
// - ![](col_resize_cursor.png) "col-resize"
// - ![](row_resize_cursor.png) "row-resize"
// - ![](n_resize_cursor.png) "n-resize"
// - ![](e_resize_cursor.png) "e-resize"
// - ![](s_resize_cursor.png) "s-resize"
// - ![](w_resize_cursor.png) "w-resize"
// - ![](ne_resize_cursor.png) "ne-resize"
// - ![](nw_resize_cursor.png) "nw-resize"
// - ![](sw_resize_cursor.png) "sw-resize"
// - ![](se_resize_cursor.png) "se-resize"
// - ![](ew_resize_cursor.png) "ew-resize"
// - ![](ns_resize_cursor.png) "ns-resize"
// - ![](nesw_resize_cursor.png) "nesw-resize"
// - ![](nwse_resize_cursor.png) "nwse-resize"
// - ![](zoom_in_cursor.png) "zoom-in"
// - ![](zoom_out_cursor.png) "zoom-out"
/*

C function : gdk_cursor_new_from_name
*/
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

	return retGo
}

// Returns a #GdkPixbuf with the image used to display the cursor.
//
// Note that depending on the capabilities of the windowing system and
// on the cursor, GDK may not be able to obtain the image data. In this
// case, %NULL is returned.
/*

C function : gdk_cursor_get_image
*/
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

// Warps the pointer of @display to the point @x,@y on
// the screen @screen, unless the pointer is confined
// to a window by a grab, in which case it will be moved
// as far as allowed by the grab. Warping the pointer
// creates events as if the user had moved the mouse
// instantaneously to the destination.
//
// Note that the pointer should normally be under the
// control of the user. This function was added to cover
// some rare use cases like keyboard navigation support
// for the color picker in the #GtkColorSelectionDialog.
/*

C function : gdk_display_warp_pointer
*/
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

// Gets a visual to use for creating windows with an alpha channel.
// The windowing system on which GTK+ is running
// may not support this capability, in which case %NULL will
// be returned. Even if a non-%NULL value is returned, its
// possible that the window’s alpha channel won’t be honored
// when displaying the window on the screen: in particular, for
// X an appropriate windowing manager and compositing manager
// must be running to provide appropriate display.
//
// This functionality is not implemented in the Windows backend.
//
// For setting an overall opacity for a top-level window, see
// gdk_window_set_opacity().
/*

C function : gdk_screen_get_rgba_visual
*/
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

// Move the part of @window indicated by @region by @dy pixels in the Y
// direction and @dx pixels in the X direction. The portions of @region
// that not covered by the new position of @region are invalidated.
//
// Child windows are not moved.
/*

C function : gdk_window_move_region
*/
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

// Toggles whether a window needs the user's
// urgent attention.
/*

C function : gdk_window_set_urgency_hint
*/
func (recv *Window) SetUrgencyHint(urgent bool) {
	c_urgent :=
		boolToGboolean(urgent)

	C.gdk_window_set_urgency_hint((*C.GdkWindow)(recv.native), c_urgent)

	return
}
