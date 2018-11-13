// This is a generated file - DO NOT EDIT
// +build gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Creates a new cursor from a cairo image surface.
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
// On the X backend, support for RGBA cursors requires a
// sufficently new version of the X Render extension.
/*

C function : gdk_cursor_new_from_surface
*/
func CursorNewFromSurface(display *Display, surface *cairo.Surface, x float64, y float64) *Cursor {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_surface := (*C.cairo_surface_t)(C.NULL)
	if surface != nil {
		c_surface = (*C.cairo_surface_t)(surface.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	retC := C.gdk_cursor_new_from_surface(c_display, c_surface, c_x, c_y)
	retGo := CursorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns a cairo image surface with the image used to display the cursor.
//
// Note that depending on the capabilities of the windowing system and
// on the cursor, GDK may not be able to obtain the image data. In this
// case, %NULL is returned.
/*

C function : gdk_cursor_get_surface
*/
func (recv *Cursor) GetSurface() (*cairo.Surface, float64, float64) {
	var c_x_hot C.gdouble

	var c_y_hot C.gdouble

	retC := C.gdk_cursor_get_surface((*C.GdkCursor)(recv.native), &c_x_hot, &c_y_hot)
	var retGo (*cairo.Surface)
	if retC == nil {
		retGo = nil
	} else {
		retGo = cairo.SurfaceNewFromC(unsafe.Pointer(retC))
	}

	xHot := (float64)(c_x_hot)

	yHot := (float64)(c_y_hot)

	return retGo, xHot, yHot
}

// Gets the current location of @device in double precision. As a slave device's
// coordinates are those of its master pointer, this function
// may not be called on devices of type %GDK_DEVICE_TYPE_SLAVE,
// unless there is an ongoing grab on them. See gdk_device_grab().
/*

C function : gdk_device_get_position_double
*/
func (recv *Device) GetPositionDouble() (*Screen, float64, float64) {
	var c_screen *C.GdkScreen

	var c_x C.gdouble

	var c_y C.gdouble

	C.gdk_device_get_position_double((*C.GdkDevice)(recv.native), &c_screen, &c_x, &c_y)

	screen := ScreenNewFromC(unsafe.Pointer(c_screen))

	x := (float64)(c_x)

	y := (float64)(c_y)

	return screen, x, y
}

// Returns the internal scale factor that maps from monitor coordinates
// to the actual device pixels. On traditional systems this is 1, but
// on very high density outputs this can be a higher value (often 2).
//
// This can be used if you want to create pixel based data for a
// particular monitor, but most of the time you’re drawing to a window
// where it is better to use gdk_window_get_scale_factor() instead.
/*

C function : gdk_screen_get_monitor_scale_factor
*/
func (recv *Screen) GetMonitorScaleFactor(monitorNum int32) int32 {
	c_monitor_num := (C.gint)(monitorNum)

	retC := C.gdk_screen_get_monitor_scale_factor((*C.GdkScreen)(recv.native), c_monitor_num)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gdk_window_create_similar_image_surface : unsupported parameter format : no type generator for gint (cairo_format_t) for param format

// Gets the list of children of @window known to GDK with a
// particular @user_data set on it.
//
// The returned list must be freed, but the elements in the
// list need not be.
//
// The list is returned in (relative) stacking order, i.e. the
// lowest window is first.
/*

C function : gdk_window_get_children_with_user_data
*/
func (recv *Window) GetChildrenWithUserData(userData uintptr) *glib.List {
	c_user_data := (C.gpointer)(userData)

	retC := C.gdk_window_get_children_with_user_data((*C.GdkWindow)(recv.native), c_user_data)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_window_get_device_position_double : unsupported parameter mask : GdkModifierType* with indirection level of 1

// Returns the internal scale factor that maps from window coordiantes
// to the actual device pixels. On traditional systems this is 1, but
// on very high density outputs this can be a higher value (often 2).
//
// A higher value means that drawing is automatically scaled up to
// a higher resolution, so any code doing drawing will automatically look
// nicer. However, if you are supplying pixel-based data the scale
// value can be used to determine whether to use a pixel resource
// with higher resolution data.
//
// The scale of a window may change during runtime, if this happens
// a configure event will be sent to the toplevel window.
/*

C function : gdk_window_get_scale_factor
*/
func (recv *Window) GetScaleFactor() int32 {
	retC := C.gdk_window_get_scale_factor((*C.GdkWindow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gdk_window_set_invalidate_handler : unsupported parameter handler : no type generator for WindowInvalidateHandlerFunc (GdkWindowInvalidateHandlerFunc) for param handler

// For optimisation purposes, compositing window managers may
// like to not draw obscured regions of windows, or turn off blending
// during for these regions. With RGB windows with no transparency,
// this is just the shape of the window, but with ARGB32 windows, the
// compositor does not know what regions of the window are transparent
// or not.
//
// This function only works for toplevel windows.
//
// GTK+ will update this property automatically if
// the @window background is opaque, as we know where the opaque regions
// are. If your window background is not opaque, please update this
// property in your #GtkWidget::style-updated handler.
/*

C function : gdk_window_set_opaque_region
*/
func (recv *Window) SetOpaqueRegion(region *cairo.Region) {
	c_region := (*C.cairo_region_t)(C.NULL)
	if region != nil {
		c_region = (*C.cairo_region_t)(region.ToC())
	}

	C.gdk_window_set_opaque_region((*C.GdkWindow)(recv.native), c_region)

	return
}
