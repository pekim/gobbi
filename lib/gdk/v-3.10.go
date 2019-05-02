// This is a generated file - DO NOT EDIT
// +build gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

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

// CursorNewFromSurface is a wrapper around the C function gdk_cursor_new_from_surface.
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

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetSurface is a wrapper around the C function gdk_cursor_get_surface.
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

// GetPositionDouble is a wrapper around the C function gdk_device_get_position_double.
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

// GetMonitorScaleFactor is a wrapper around the C function gdk_screen_get_monitor_scale_factor.
func (recv *Screen) GetMonitorScaleFactor(monitorNum int32) int32 {
	c_monitor_num := (C.gint)(monitorNum)

	retC := C.gdk_screen_get_monitor_scale_factor((*C.GdkScreen)(recv.native), c_monitor_num)
	retGo := (int32)(retC)

	return retGo
}

// CreateSimilarImageSurface is a wrapper around the C function gdk_window_create_similar_image_surface.
func (recv *Window) CreateSimilarImageSurface(format int32, width int32, height int32, scale int32) *cairo.Surface {
	c_format := (C.cairo_format_t)(format)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	c_scale := (C.int)(scale)

	retC := C.gdk_window_create_similar_image_surface((*C.GdkWindow)(recv.native), c_format, c_width, c_height, c_scale)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_window_get_children_with_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// Unsupported : gdk_window_get_device_position_double : unsupported parameter mask : GdkModifierType* with indirection level of 1

// GetScaleFactor is a wrapper around the C function gdk_window_get_scale_factor.
func (recv *Window) GetScaleFactor() int32 {
	retC := C.gdk_window_get_scale_factor((*C.GdkWindow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gdk_window_set_invalidate_handler : unsupported parameter handler : no type generator for WindowInvalidateHandlerFunc (GdkWindowInvalidateHandlerFunc) for param handler

// SetOpaqueRegion is a wrapper around the C function gdk_window_set_opaque_region.
func (recv *Window) SetOpaqueRegion(region *cairo.Region) {
	c_region := (*C.cairo_region_t)(C.NULL)
	if region != nil {
		c_region = (*C.cairo_region_t)(region.ToC())
	}

	C.gdk_window_set_opaque_region((*C.GdkWindow)(recv.native), c_region)

	return
}

// CairoSurfaceCreateFromPixbuf is a wrapper around the C function gdk_cairo_surface_create_from_pixbuf.
func CairoSurfaceCreateFromPixbuf(pixbuf *gdkpixbuf.Pixbuf, scale int32, forWindow *Window) *cairo.Surface {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	c_scale := (C.int)(scale)

	c_for_window := (*C.GdkWindow)(C.NULL)
	if forWindow != nil {
		c_for_window = (*C.GdkWindow)(forWindow.ToC())
	}

	retC := C.gdk_cairo_surface_create_from_pixbuf(c_pixbuf, c_scale, c_for_window)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetAllowedBackends is a wrapper around the C function gdk_set_allowed_backends.
func SetAllowedBackends(backends string) {
	c_backends := C.CString(backends)
	defer C.free(unsafe.Pointer(c_backends))

	C.gdk_set_allowed_backends(c_backends)

	return
}
