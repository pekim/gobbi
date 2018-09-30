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

// CursorNewFromSurface is a wrapper around the C function gdk_cursor_new_from_surface.
func CursorNewFromSurface(display *Display, surface *cairo.Surface, x float64, y float64) *Cursor {
	c_display := (*C.GdkDisplay)(display.ToC())

	c_surface := (*C.cairo_surface_t)(surface.ToC())

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	retC := C.gdk_cursor_new_from_surface(c_display, c_surface, c_x, c_y)
	retGo := CursorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSurface is a wrapper around the C function gdk_cursor_get_surface.
func (recv *Cursor) GetSurface() (*cairo.Surface, *float64, *float64) {
	var c_x_hot C.gdouble

	var c_y_hot C.gdouble

	retC := C.gdk_cursor_get_surface((*C.GdkCursor)(recv.native), &c_x_hot, &c_y_hot)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	xHot := (*float64)(&c_x_hot)

	yHot := (*float64)(&c_y_hot)

	return retGo, xHot, yHot
}

// Unsupported : gdk_device_get_position_double : unsupported parameter screen : record with indirection level of 2

// GetMonitorScaleFactor is a wrapper around the C function gdk_screen_get_monitor_scale_factor.
func (recv *Screen) GetMonitorScaleFactor(monitorNum int32) int32 {
	c_monitor_num := (C.gint)(monitorNum)

	retC := C.gdk_screen_get_monitor_scale_factor((*C.GdkScreen)(recv.native), c_monitor_num)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gdk_window_create_similar_image_surface : unsupported parameter format : no type generator for gint, cairo_format_t

// GetChildrenWithUserData is a wrapper around the C function gdk_window_get_children_with_user_data.
func (recv *Window) GetChildrenWithUserData(userData uintptr) *glib.List {
	c_user_data := (C.gpointer)(userData)

	retC := C.gdk_window_get_children_with_user_data((*C.GdkWindow)(recv.native), c_user_data)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_window_get_device_position_double : unsupported parameter mask : GdkModifierType* with indirection level of 1

// GetScaleFactor is a wrapper around the C function gdk_window_get_scale_factor.
func (recv *Window) GetScaleFactor() int32 {
	retC := C.gdk_window_get_scale_factor((*C.GdkWindow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gdk_window_set_invalidate_handler : unsupported parameter handler : no type generator for WindowInvalidateHandlerFunc, GdkWindowInvalidateHandlerFunc

// SetOpaqueRegion is a wrapper around the C function gdk_window_set_opaque_region.
func (recv *Window) SetOpaqueRegion(region *cairo.Region) {
	c_region := (*C.cairo_region_t)(region.ToC())

	C.gdk_window_set_opaque_region((*C.GdkWindow)(recv.native), c_region)

	return
}
